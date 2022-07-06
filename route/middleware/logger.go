package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
)

// * ►─═ Middleware Log Access ═─►
func LoggerAccess() fiber.Handler {
	return mlogger.New(mlogger.Config{
		Format:       "${time} - [${pid}][${latency}][${method}][${status}] ${url}\n",
		TimeFormat:   "2006-01-02T15:04:05.999Z07",
		TimeZone:     "Asia/Bangkok",
		TimeInterval: time.Second,
	})
}
