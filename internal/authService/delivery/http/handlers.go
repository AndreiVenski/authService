package http

import (
	"authService/config"
	"authService/internal/authService"
	"authService/internal/models"
	"authService/pkg/httpErrors"
	"authService/pkg/logger"
	"authService/pkg/utils"
	"github.com/gofiber/fiber/v2"
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

// @Summary Get New Tokens
// @Description Get new tokens : access and refresh
// @ID getNewTokens
// @Accept json
// @Produce json
// @Param X-Forwarded-For header string false "User's IP Address"
// @Param request body models.GetNewTokenData true "User Info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tokens [post]
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
	tokens, err := h.authUC.GetNewTokens(ctx.UserContext(), userInfo)
	if err != nil {
		h.logger.Errorf("Failed to get new tokens: %v", err)
		return respondWithError(ctx, fiber.StatusInternalServerError, "Could not generate tokens")
	}

	return ctx.JSON(fiber.Map{"tokens": tokens})
}

// @Summary Refresh Tokens
// @Description Refresh access token
// @ID refreshToken
// @Accept json
// @Produce json
// @Param X-Forwarded-For header string false "User's IP Address"
// @Param request body models.RefreshData true "Refresh Token Data"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tokens/refresh [post]
func (h *authHandler) RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken := &models.RefreshData{}

	err := utils.ReadFromRequest(ctx, refreshToken)
	if err != nil {
		h.logger.Errorf("Invalid refresh token dataL %v", err)
		return respondWithError(ctx, fiber.StatusBadRequest, "Invalid refresh token data")
	}

	ip := ctx.Get("X-Forwarded-For")
	if ip == "" {
		ip = ctx.IP()
	}

	token, newRefreshTokenID, err := h.authUC.RefreshAccessToken(ctx.UserContext(), refreshToken.Token, refreshToken.TokenID, ip)
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
