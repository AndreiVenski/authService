package usecase

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
	"authService/pkg/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

type authUseCase struct {
	cfg       *config.Config
	logger    logger.Logger
	authRepo  authService.Repository
	authEmail authService.Email
}

func NewAuthUseCase(cfg *config.Config, logger logger.Logger, authRepo authService.Repository, authEmail authService.Email) authService.UseCase {
	return &authUseCase{
		cfg:       cfg,
		logger:    logger,
		authRepo:  authRepo,
		authEmail: authEmail,
	}
}

func (uc *authUseCase) GetNewTokens(ctx context.Context, userInfo *models.UserInfo) (*models.Tokens, error) {
	tokens, err := utils.GenerateTokens(uc.cfg, userInfo)
	if err != nil {
		return nil, err
	}

	refreshTokenRecord := models.NewRefreshTokenRecord(tokens, uc.cfg.Server.RefreshTokenExpiresHourInt, userInfo.IP)
	if err = refreshTokenRecord.HashToken(tokens.RefreshToken); err != nil {
		return nil, err
	}

	if err = uc.authRepo.WriteRefreshTokenRecord(ctx, refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (uc *authUseCase) RefreshAccessToken(ctx context.Context, refreshToken, ipAddr string) (string, error) {
	refreshTokenRecord := &models.RefreshTokenRecord{}
	if err := refreshTokenRecord.HashToken(refreshToken); err != nil {
		return "", err
	}

	tokenData, err := uc.authRepo.GetRefreshTokenData(ctx, refreshTokenRecord.GetHashedToken())
	if err != nil {
		return "", err
	}

	if time.Now().After(tokenData.Expires) {
		return "", errors.New("token expires")
	}

	if ipAddr != tokenData.IPAddr {

		if err := uc.ipMismatch(ctx, tokenData.UserID); err != nil {

			return "", err
		}
	}

	userInfo := &models.UserInfo{UserID: tokenData.UserID, IP: ipAddr}
	newRefreshTokenID := uuid.New()
	accessToken, err := utils.GenerateAccessToken(uc.cfg, userInfo, newRefreshTokenID)
	if err != nil {
		return "", err
	}

	if err = uc.authRepo.UpdateRefreshTokenID(ctx, refreshTokenRecord.GetHashedToken(), newRefreshTokenID); err != nil {
		return "", err
	}

	return accessToken, nil
}

func (uc *authUseCase) ipMismatch(ctx context.Context, UserID uuid.UUID) error {
	user, err := uc.authRepo.GetUser(ctx, UserID)
	if err != nil {
		return err
	}

	err = uc.authEmail.SendWarningIPEmail(user)
	return err
}
