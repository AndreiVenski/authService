package http

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/logger"
	"authService/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authUC authService.UseCase
	cfg    *config.Config
	logger logger.Logger
}

func NewAuthHandler(cfg *config.Config, logger logger.Logger, authUC authService.UseCase) authService.Handler {
	return &authHandler{
		cfg:    cfg,
		logger: logger,
		authUC: authUC,
	}
}

func (h *authHandler) GetNewTokens(ctx *fiber.Ctx) error {
	var userInfo *models.UserInfo
	err := utils.ReadInfoFromRequest(ctx, userInfo)
	if err != nil {
		return ctx.JSON(fiber.Map{})
	}

	tokens, err := h.authUC.GetNewTokens(userInfo)
	if err != nil {

	}

	return ctx.JSON(fiber.Map{"tokens": tokens})
}

func (h *authHandler) RefreshTokens(ctx *fiber.Ctx) error {
	return nil
}
