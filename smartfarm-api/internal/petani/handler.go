package petani

import (
	"strconv"

	"smartfarm-api/pkg/response" // ✅ import response global kamu

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type PetaniHandler struct {
    service PetaniService
}


func NewPetaniHandler(service PetaniService) *PetaniHandler {
	return &PetaniHandler{service: service}
}

// ✅ Create Petani
func (h *PetaniHandler) CreatePetani(c *fiber.Ctx) error {
	var req CreatePetaniRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid JSON body")
	}

	// 🔍 Validasi otomatis dari tag di DTO
	if err := validate.Struct(req); err != nil {
		details := response.MapValidationErrors(err)
		return response.ValidationError(c, "Validation failed", details)
	}

	petani, err := h.service.CreatePetani(req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal membuat data petani")
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:   petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}

	return response.Success(c, "Petani berhasil dibuat", resp)
}

// ✅ Get All Petani
func (h *PetaniHandler) GetAllPetani(c *fiber.Ctx) error {
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
	petanis, total, err := h.service.GetAllPetaniPaginated(page, limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data petani")
	}

	// 🔹 Hitung total halaman
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// 🔹 Bentuk meta data
	meta := fiber.Map{
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPages,
	}

	// 🔹 Kirim response sukses dengan meta info
	return response.SuccessWithMeta(c, "Data petani ditemukan", petanis, meta)
}


// ✅ Get Petani by ID
func (h *PetaniHandler) GetPetaniByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	petani, err := h.service.GetPetaniByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Petani tidak ditemukan")
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:   petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}

	return response.Success(c, "Data petani ditemukan", resp)
}

// ✅ Update Petani
func (h *PetaniHandler) UpdatePetani(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var req UpdatePetaniRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid JSON body")
	}

	// Validasi
	if err := validate.Struct(req); err != nil {
		details := response.MapValidationErrors(err)
		return response.ValidationError(c, "Validation failed", details)
	}

	petani, err := h.service.UpdatePetani(uint(id), req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal memperbarui data petani")
	}

	resp := PetaniResponse{
		ID:        petani.ID,
		Nama:      petani.Nama,
		Alamat:   petani.Alamat,
		CreatedAt: petani.CreatedAt,
		UpdatedAt: petani.UpdatedAt,
	}

	return response.Success(c, "Data petani berhasil diperbarui", resp)
}

// ✅ Delete Petani
func (h *PetaniHandler) DeletePetani(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	if err := h.service.DeletePetani(uint(id)); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Gagal menghapus data petani")
	}

	return response.Success(c, "Petani berhasil dihapus", nil)
}
