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
	userInfo := &models.UserInfo{}
	err := utils.ReadFromRequest(ctx, userInfo)
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	tokens, err := h.authUC.GetNewTokens(ctx.Context(), userInfo)
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"tokens": tokens})
}

func (h *authHandler) RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken := &struct {
		Token string `json:"refresh_token"`
	}{}

	err := utils.ReadFromRequest(ctx, refreshToken)
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	tokens, err := h.authUC.RefreshAccessToken(ctx.Context(), refreshToken.Token, ctx.IP())
	if err != nil {
		return ctx.JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"tokens": tokens})
}
