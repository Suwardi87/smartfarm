package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// 🧱 Struktur dasar response global
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// 📦 Struktur untuk error detail (optional)
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

// ✅ Response sukses (200, 201, dll)
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ✅ Response sukses dengan pagination/meta info
func SuccessWithMeta(c *fiber.Ctx, message string, data interface{}, meta interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// ❌ Response error umum
func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(APIResponse{
		Status:  "error",
		Message: message,
	})
}

// ❌ Response error dengan detail validasi
func ValidationError(c *fiber.Ctx, message string, details []ErrorDetail) error {
	return c.Status(fiber.StatusBadRequest).JSON(APIResponse{
		Status:  "error",
		Message: message,
		Data:    details,
	})
}

// ⚙️ Fungsi helper untuk membentuk detail validasi dari validator
func MapValidationErrors(errs error) []ErrorDetail {
	var details []ErrorDetail

	if errs == nil {
		return details
	}

	for _, e := range errs.(validator.ValidationErrors) {
		detail := ErrorDetail{
			Field:   e.Field(),
			Message: "Invalid value for " + e.Field(),
		}
		details = append(details, detail)
	}

	return details
}
