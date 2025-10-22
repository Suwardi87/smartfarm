package lahan

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"smartfarm-api/pkg/response" // pastikan ini sesuai path package kamu
)

var validate = validator.New()

type LahanHandler struct {
	service lahanService
}

func NewLahanHandler(service lahanService) *LahanHandler {
	return &LahanHandler{service: service}
}

// ✅ Get All Lahan (dengan pagination)
func (h *LahanHandler) GetAllLahan(c *fiber.Ctx) error {
	// 🔹 Ambil query param page & limit
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// 🔹 Panggil service
	lahans, total, err := h.service.GetAllLahanPaginated(page, limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data lahan")
	}

	// 🔹 Hitung total halaman
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// 🔹 Buat metadata untuk pagination
	meta := fiber.Map{
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	}

	// 🔹 Kirim response sukses dengan meta info
	return response.SuccessWithMeta(c, "Data lahan ditemukan", lahans, meta)
}
