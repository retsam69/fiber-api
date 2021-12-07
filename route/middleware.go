package route

import (
	"strings"
	"time"

	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/phuslu/log"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/compress"
	// "github.com/gofiber/fiber/v2/middleware/cache"
	// "github.com/gofiber/fiber/v2/middleware/timeout"
)

// * ░▒▓█►─═ Middleware CORS ═─►█▓▒░
func CORS() fiber.Handler {
	return cors.New()
}

// * ░▒▓█►─═ Middleware Http Basic Auth ═─►█▓▒░
func BasicAuth() fiber.Handler {
	users := viper.GetString("auth.basic")
	var u = make(map[string]string)
	for _, v := range strings.Split(users, ",") {
		// log.Debug().Msg(v)
		t := strings.SplitN(v, ":", 2)
		if len(t) == 2 {
			u[strings.TrimSpace(t[0])] = strings.TrimSpace(t[1])
			log.Debug().Str("Username", strings.TrimSpace(t[0])).Msgf("BasicAuth: User Registered")
		}
	}
	return basicauth.New(basicauth.Config{
		Users: u,
	})
}

// * ░▒▓█►─═ Middleware Log Access ═─►█▓▒░
func LoggerAccess() fiber.Handler {
	return mlogger.New(mlogger.Config{
		Format:       "${time} - [${pid}][${latency}][${method}][${status}] ${url}\n",
		TimeFormat:   "2006-01-02T15:04:05.999Z07",
		TimeZone:     "Asia/Bangkok",
		TimeInterval: time.Second,
	})
}

// * ░▒▓█►─═ Middleware Compress ═─►█▓▒░
// func Compress() fiber.Handler {
// 	return compress.New(compress.Config{
// 		Level: compress.LevelBestSpeed,
// 	})
// }

// * ░▒▓█►─═ Middleware Cache ═─►█▓▒░
// func FiberCache() fiber.Handler {
// 	return cache.New(cache.Config{
// 		Next: func(c *fiber.Ctx) bool {
// 			return c.Query("cache") == "false"
// 		},
// 		Expiration:   30 * time.Minute,
// 		CacheControl: true,
// 	})
// }

// * ░▒▓█►─═ Middleware Timeout ═─►█▓▒░
// var SetTimeout = timeout.New
