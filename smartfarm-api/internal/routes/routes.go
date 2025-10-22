package routes

import (
	"github.com/gofiber/fiber/v2"
	"smartfarm-api/internal/middleware"
	"smartfarm-api/internal/petani"
	"smartfarm-api/internal/lahan"
	"smartfarm-api/internal/user"
)

// RegisterRoutes berfungsi mendaftarkan semua route global
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api") // prefix global API

	// =============================
	// 🔐 AUTH MODULE (REGISTER / LOGIN)
	// =============================

	// ✅ inisialisasi dependency chain
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	// ✅ route group
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Get("/users", handler.GetAll)

	// =============================
	// 🧱 ADMIN / PETANI MODULE
	// =============================
	petaniRepo := petani.NewPetaniRepository()
    petaniService := petani.NewPetaniService(petaniRepo) // ✅ return interface
	petaniHandler := petani.NewPetaniHandler(petaniService)
		// =============================
	// 🧱 ADMIN / LAHAN MODULE
	// =============================
	lahanRepo := lahan.NewLahanRepository()
    lahanService := lahan.NewLahanService(lahanRepo) // ✅ return interface
	lahanHandler := lahan.NewLahanHandler(lahanService)

	// ✅ Route lengkap
	admin := api.Group("/admin")
	// petani routes
	admin.Get("/petani", petaniHandler.GetAllPetani)
	admin.Post("/petani", petaniHandler.CreatePetani)
	admin.Get("/petani/:id", petaniHandler.GetPetaniByID)
	admin.Put("/petani/:id", petaniHandler.UpdatePetani)
	admin.Delete("/petani/:id", petaniHandler.DeletePetani)
	// lahan routes
	admin.Get("/petani", lahanHandler.GetAllLahan)
	// admin.Post("/petani", lahanHandler.CreateLahan)
	// admin.Get("/petani/:id", lahanHandler.GetLahanByID)
	// admin.Put("/petani/:id", lahanHandler.UpdateLahan)
	// admin.Delete("/petani/:id", lahanHandler.DeleteLahan)

	// =============================
	// 🧱 PROTECTED ROUTES (JWT)
	// =============================

	protected := api.Group("/protected", middleware.JWTProtected())

	protected.Get("/me", func(c *fiber.Ctx) error {
		claims := c.Locals("user")
		return c.JSON(fiber.Map{
			"message": "token valid",
			"user":    claims,
		})
	})

	// =============================
	// 📋 NANTI TAMBAHKAN MODUL LAIN DI SINI
	// =============================
	// contoh:
	// petaniRepo := petani.NewRepository()
	// petaniService := petani.NewService(petaniRepo)
	// petaniHandler := petani.NewHandler(petaniService)
	//
	// petaniRoutes := api.Group("/petani", middleware.JWTProtected([]byte(os.Getenv("JWT_SECRET"))))
	// petaniRoutes.Get("/", petaniHandler.GetAll)
	// petaniRoutes.Post("/", petaniHandler.Create)
}
