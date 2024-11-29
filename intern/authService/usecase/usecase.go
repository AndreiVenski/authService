package usecase

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
	"authService/pkg/utils"
)

type authUseCase struct {
	cfg      *config.Config
	logger   logger.Logger
	authRepo authService.Repository
}

func NewAuthUseCase(cfg *config.Config, logger logger.Logger, authRepo authService.Repository) authService.UseCase {
	return &authUseCase{
		cfg:      cfg,
		logger:   logger,
		authRepo: authRepo,
	}
}

func (uc *authUseCase) GetNewTokens(userInfo *models.UserInfo) (*models.Tokens, error) {
	err := uc.authRepo.FindAndDeleteToken(userInfo)
	if err != nil {
		return nil, nil
	}

	tokens, err := utils.GenerateTokens(uc.cfg, userInfo)
	if err != nil {
		return nil, err
	}

	refreshTokenRecord := models.NewRefreshTokenRecord(tokens, uc.cfg.Server.RefreshTokenExpiresHourInt)
	if err = refreshTokenRecord.HashToken(tokens.RefreshToken); err != nil {
		return nil, err
	}

	if err = uc.authRepo.WriteRefreshTokenRecord(refreshTokenRecord); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (uc *authUseCase) RefreshTokens() {

}
