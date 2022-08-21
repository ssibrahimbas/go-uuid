package server

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/go-uuid/tools"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*tools.Result); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(e)
	}
	if e, ok := err.(*tools.DataResult); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(e)
	}
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	err = ctx.Status(code).JSON(tools.ErrorResult(err.Error()))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(tools.ErrorResult("Something went wrong"))
	}
	return nil
}
