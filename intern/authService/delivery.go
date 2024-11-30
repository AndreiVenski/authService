package authService

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetNewTokens(ctx *fiber.Ctx) error
	RefreshTokens(ctx *fiber.Ctx) error
}
