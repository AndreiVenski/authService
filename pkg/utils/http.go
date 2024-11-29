package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ReadFromRequest(ctx *fiber.Ctx, model interface{}) error {
	if err := ctx.BodyParser(model); err != nil {
		return err
	}
	return validate.Struct(model)
}
