package common

import "github.com/gofiber/fiber/v2"

type (
	BaseHandler struct {
	}
	BaseResponseDTO[T any] struct {
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
		Error      string      `json:"error,omitempty"`
		StatusCode int         `json:"status_code,omitempty"`
	}
)

func NewBaseResponseDTO[T any](message string, data *T, errorStr string, statusCode int) *BaseResponseDTO[T] {

	return &BaseResponseDTO[T]{
		Message:    message,
		Data:       data,
		Error:      errorStr,
		StatusCode: statusCode,
	}
}

func NewSuccessResponseDTO[T any](data *T) *BaseResponseDTO[T] {
	return NewBaseResponseDTO("Амжилттай", data, "", fiber.StatusOK)
}

func NewErrorResponseDTO(error string) *BaseResponseDTO[any] {
	return NewBaseResponseDTO[any]("Алдаа гарлаа", nil, error, fiber.StatusBadRequest)
}

func (h *BaseHandler) Success(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(NewSuccessResponseDTO[any](&data))
}

func (h *BaseHandler) Error(c *fiber.Ctx, error error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponseDTO(error.Error()))
}
