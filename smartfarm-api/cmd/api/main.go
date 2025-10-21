package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	
	"github.com/gofiber/swagger"                  // ✅ Swagger handler

	"smartfarm-api/internal/middleware"
	"smartfarm-api/internal/routes"
	"smartfarm-api/internal/user"
	"smartfarm-api/internal/petani"
	"smartfarm-api/pkg/database"

	_ "smartfarm-api/internal/docs" // ✅ Hasil dari `swag init`
)

// @title SmartFarm API
// @version 1.0
// @description API dokumentasi untuk sistem SmartFarm berbasis Fiber + GORM
// @contact.name Suwardi Developer
// @contact.email suwardi87@gmail.com
// @host localhost:8083
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan format: Bearer {token}

func main() {
	// === 1️⃣ Koneksi ke database ===
	database.Connect()

	// === 2️⃣ Migrasi model ===
	if err := database.DB.AutoMigrate(&user.User{}, &petani.Petani{}); err != nil {
		log.Fatal("❌ Gagal migrasi:", err)
	}
	log.Println("✅ Migrasi user sukses")

	// === 3️⃣ Inisialisasi Fiber ===
	app := fiber.New()

	// === 4️⃣ Middleware Global ===
	app.Use(middleware.CORSConfig()) // ✅ Panggil CORS dari folder middleware

	// === 5️⃣ Swagger UI ===
	app.Get("/swagger/*", swagger.HandlerDefault)

	// === 6️⃣ Register semua route global ===
	routes.RegisterRoutes(app)

	// === 7️⃣ Root redirect ke Swagger ===
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	// === 8️⃣ Jalankan server ===
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 SmartFarm API running at http://localhost%s\n", addr)
	log.Printf("📚 Swagger docs available at http://localhost%s/swagger/index.html\n", addr)

	log.Fatal(app.Listen(addr))
}
