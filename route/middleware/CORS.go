package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// * ►─═ Middleware CORS ═─►
func CORS() fiber.Handler {
	return cors.New()
}
