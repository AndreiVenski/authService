package http

import (
	"authService/internal/authService"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(api fiber.Router, handlers authService.Handler) {
	api.Post("/tokens", handlers.GetNewTokens)
	api.Post("/tokens/refresh", handlers.RefreshTokens)
}
