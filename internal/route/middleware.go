package route

import (
	"os"
	"strings"
	"time"

	mlogger "github.com/gofiber/fiber/v2/middleware/logger"

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
	prefixUser := prefix
	Users := make(map[string]string)
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, prefixUser) {
			sp := strings.SplitN(e, "=", 1)
			if len(sp) == 2 {
				username := sp[0]
				username = strings.TrimPrefix(username, prefixUser)
				username = strings.ToLower(username)
				password := sp[1]
				if username != "" && password != "" {
					Users[username] = password
				}
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
