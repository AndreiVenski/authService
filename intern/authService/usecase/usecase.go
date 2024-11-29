package usecase

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
	"authService/pkg/utils"
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

func (uc *authUseCase) GetNewTokens(userInfo *models.UserInfo) (*models.Tokens, error) {
	tokens, err := utils.GenerateTokens(uc.cfg, userInfo)
	if err != nil {
		return nil, err
	}

	refreshTokenRecord := models.NewRefreshTokenRecord(tokens, uc.cfg.Server.RefreshTokenExpiresHourInt, userInfo.IP)
	if err = refreshTokenRecord.HashToken(tokens.RefreshToken); err != nil {
		return nil, err
	}

	if err = uc.authRepo.WriteRefreshTokenRecord(refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (uc *authUseCase) RefreshAccessToken(refreshToken, ipAddr string) (*models.Tokens, error) {
	tokenData, err := uc.authRepo.GetRefreshTokenData(refreshToken)
	if err != nil {
		return nil, err
	}

	if time.Now().After(tokenData.Expires) {
		return nil, errors.New("token expires")
	}

	if ipAddr != tokenData.IPAddr {

		if err := uc.ipMismatch(tokenData.UserID); err != nil {

			return nil, err
		}
	}

	userInfo := &models.UserInfo{UserID: tokenData.UserID, IP: ipAddr}
	tokens, err := utils.GenerateTokens(uc.cfg, userInfo)
	if err != nil {
		return nil, err
	}
	refreshTokenRecord := models.NewRefreshTokenRecord(tokens, uc.cfg.Server.RefreshTokenExpiresHourInt, ipAddr)
	if err = refreshTokenRecord.HashToken(tokens.RefreshToken); err != nil {
		return nil, err
	}

	if err = uc.authRepo.UpdateRefreshTokenRecord(refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (uc *authUseCase) ipMismatch(UserID uuid.UUID) error {
	user, err := uc.authRepo.GetUser(UserID)
	if err != nil {
		return err
	}

	err = uc.authEmail.SendWarningIPEmail(user)
	return err
}
