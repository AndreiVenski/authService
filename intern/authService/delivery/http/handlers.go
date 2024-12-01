package http

import (
	"authService/config"
	"authService/intern/authService"
	"authService/intern/models"
	"authService/pkg/httpErrors"
	"authService/pkg/logger"
	"authService/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func respondWithError(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(fiber.Map{
		"error": message,
	})
}

func (h *authHandler) GetNewTokens(ctx *fiber.Ctx) error {
	userInfo := &models.UserInfo{}
	err := utils.ReadFromRequest(ctx, userInfo)
	if err != nil {
		h.logger.Errorf("Invalid request data: %v")
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid request data")
	}
	ip := ctx.Get("X-Forwarded-For")
	if ip == "" {
		ip = ctx.IP()
	}
	userInfo.IP = ip
	tokens, err := h.authUC.GetNewTokens(ctx.Context(), userInfo)
	if err != nil {
		h.logger.Errorf("Failed to get new tokens: %v", err)
		return respondWithError(ctx, fiber.StatusInternalServerError, "Could not generate tokens")
	}

	return ctx.JSON(fiber.Map{"tokens": tokens})
}

func (h *authHandler) RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken := &struct {
		Token   string    `json:"refresh_token" validate:"required,base64,len=44"`
		TokenID uuid.UUID `json:"refresh_token_id" validate:"required,uuid4"`
	}{}

	err := utils.ReadFromRequest(ctx, refreshToken)
	if err != nil {
		h.logger.Errorf("Invalid refresh token dataL %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid refresh token data")
	}

	ip := ctx.Get("X-Forwarded-For")
	if ip == "" {
		ip = ctx.IP()
	}

	token, newRefreshTokenID, err := h.authUC.RefreshAccessToken(ctx.Context(), refreshToken.Token, refreshToken.TokenID, ip)
	if err != nil {
		switch {
		case errors.Is(err, httpErrors.ErrRefreshTokenExpires),
			errors.Is(err, httpErrors.ErrUserNotFound),
			errors.Is(err, httpErrors.ErrRefreshTokenNotFound),
			errors.Is(err, httpErrors.ErrRefreshTokenIncorrect):
			return respondWithError(ctx, fiber.StatusUnauthorized, err.Error())
		default:
			h.logger.Errorf("Failed to refresh access token: %v", err)
			return respondWithError(ctx, fiber.StatusInternalServerError, "Could not refresh access token")
		}
	}

	return ctx.JSON(fiber.Map{"access_token": token, "refresh_token_id": newRefreshTokenID})
}
