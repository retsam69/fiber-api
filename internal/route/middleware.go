package route

import (
	"strings"
	"time"

	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New()
}

// func Compress() fiber.Handler {
// 	return compress.New(compress.Config{
// 		Level: compress.LevelBestSpeed,
// 	})

// }
func GetUsersByEnv(prefix string) map[string]string {
	prefixUser := strings.ToLower(prefix)
	Users := make(map[string]string)
	for _, e := range viper.AllKeys() {
		if strings.HasPrefix(e, prefixUser) {
			username := strings.TrimPrefix(e, prefixUser)
			password := viper.GetString(e)
			if username != "" && password != "" {
				Users[username] = password
			}
		}
	}

	return Users
}
func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: GetUsersByEnv("USER_"),
	})
}

func LoggerAccess() fiber.Handler {
	return mlogger.New(mlogger.Config{
		Format:       "${time} - [${latency}][${method}][${status}] ${url}\n",
		TimeFormat:   "2006-01-02T15:04:05.999Z07:00",
		TimeZone:     "Asia/Bangkok",
		TimeInterval: time.Second,
	})
}
