package middleware

// import (
// "github.com/gofiber/fiber/v2/middleware/cache"
// )

// * ►─═ Middleware Cache ═─►
// func FiberCache() fiber.Handler {
// 	return cache.New(cache.Config{
// 		Next: func(c *fiber.Ctx) bool {
// 			return c.Query("cache") == "false"
// 		},
// 		Expiration:   30 * time.Minute,
// 		CacheControl: true,
// 	})
// }
