package usecase

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
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
	err := uc.authRepo.FindAndDeleteToken()
	if err != nil {
		return nil, nil
	}

	return nil, nil
}

func (uc *authUseCase) RefreshTokens() {}
