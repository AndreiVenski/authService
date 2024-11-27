package http

import (
	"authService/intern/authService"
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(api fiber.Router, handlers authService.Handler) {
	api.Post("/tokens")
	api.Post("/tokens/refresh")
}
