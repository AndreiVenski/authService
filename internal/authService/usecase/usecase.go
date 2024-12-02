package usecase

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
	"authService/pkg/utils"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "authUseCase.GetNewTokens.GenerateTokens")
	}

	refreshTokenRecord := models.NewRefreshTokenRecord(tokens, uc.cfg.Server.RefreshTokenExpiresHourInt, userInfo.IP)
	if err = refreshTokenRecord.HashToken(tokens.RefreshToken); err != nil {
		return nil, errors.Wrap(err, "authUseCase.GetNewTokens.HashToken")
	}

	if err = uc.authRepo.WriteRefreshTokenRecord(ctx, refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (uc *authUseCase) RefreshAccessToken(ctx context.Context, refreshToken string, refreshTokenID uuid.UUID, ipAddr string) (string, string, error) {
	tokenData, err := uc.authRepo.GetRefreshTokenData(ctx, refreshTokenID)
	if err != nil {
		return "", "", err
	}

	if time.Now().After(tokenData.Expires) {
		return "", "", errors.New("token expires")
	}

	if !tokenData.VerifyRefreshToken(refreshToken) {
		return "", "", errors.New("token is incorrect")
	}

	if ipAddr != tokenData.IPAddr {
		if err = uc.ipMismatch(ctx, tokenData.UserID, ipAddr); err != nil {
			return "", "", err
		}
	}

	userInfo := &models.UserInfo{UserID: tokenData.UserID, IP: ipAddr}
	newRefreshTokenID := uuid.New()
	accessToken, err := utils.GenerateAccessToken(uc.cfg, userInfo, newRefreshTokenID)
	if err != nil {
		return "", "", errors.Wrap(err, "authUseCase.RefreshAccessToken.GenerateAcceessToken")
	}

	if err = uc.authRepo.UpdateRefreshTokenID(ctx, refreshTokenID, newRefreshTokenID); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshTokenID.String(), nil
}

func (uc *authUseCase) ipMismatch(ctx context.Context, userID uuid.UUID, ipAddr string) error {
	user, err := uc.authRepo.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	err = uc.authEmail.SendWarningIPEmail(user, ipAddr)
	return err
}
