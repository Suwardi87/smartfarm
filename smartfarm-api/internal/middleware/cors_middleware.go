package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORSConfig berfungsi mengatur cross-origin request
func CORSConfig() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // ganti nanti dengan domain FE
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, // aktif kalau kamu pakai JWT/Cookie
	})
}
