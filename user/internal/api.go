package internal

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/go-uuid/entity"
	"github.com/ssibrahimbas/go-uuid/tools"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	u := &entity.CreateUserDto{}
	c.BodyParser(u)
	usr, err := h.srv.NewUser(u)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.ErrorResult(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(tools.SuccessDataResult("User created successfully", usr))
}
func (h *Handler) getUserDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	usr, err := h.srv.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(tools.ErrorResult("something went wrong"))
	}
	return c.Status(fiber.StatusOK).JSON(tools.SuccessDataResult("User fetched successfully", usr))
}
func (h *Handler) getAllUsers(c *fiber.Ctx) error {
	arr, err := h.srv.GetAllUsers()
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusBadRequest).JSON(tools.ErrorDataResult("something went wrong", err))
	}
	return c.Status(fiber.StatusOK).JSON(tools.SuccessDataResult("Users listed successfully", arr))
}
