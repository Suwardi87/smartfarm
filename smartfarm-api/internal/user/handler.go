package user

import (
	"net/http"
	"smartfarm-api/pkg/database"
	"smartfarm-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service userService
}

func NewHandler(service userService) *Handler {
	return &Handler{service: service}
}



// Register godoc
// @Summary Register user baru
// @Description Mendaftarkan user baru dengan hashing password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body user.RegisterRequest true "User data"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var input RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}

	user := User{
		NamaLengkap: input.NamaLengkap,
		Username:    input.Username,
		Password:    hashed,
		Level:       input.Level,
	}

	// 🔍 ubah bagian ini biar tahu error detailnya
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(), // tampilkan error asli
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "register success",
		"data": fiber.Map{
			"id":           user.ID,
			"nama_lengkap": user.NamaLengkap,
			"username":     user.Username,
			"level":        user.Level,
		},
	})
}


// Login godoc
// @Summary Login user
// @Description Login menggunakan username dan password
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body user.LoginRequest true "username & password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var input LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	token, user, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"level":    user.Level,
		},
	})
}


// GetAll godoc
// @Summary Mendapatkan semua user
// @Description Menampilkan seluruh data user dari database
// @Tags User
// @Produce json
// @Success 200 {array} user.User
// @Router /auth/users [get]
// di Handler
func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAllUser()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}