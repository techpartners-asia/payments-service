package fiberPkg

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return c.Status(fiber.StatusBadRequest).JSON(sharedDTO.NewBaseResponseDTO[any](err.Error(), nil, err.Error(), fiber.StatusBadRequest))
		// },
	})
}
