package route

import (
	"fmt"
	"time"

	rapidoc "github.com/attapon-th/gofiber-rapidoc"
	mlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New()
}

// "github.com/gofiber/fiber/v2/middleware/compress"
// func Compress() fiber.Handler {
// 	return compress.New(compress.Config{
// 		Level: compress.LevelBestSpeed,
// 	})
// }

// "github.com/gofiber/fiber/v2/middleware/basicauth"
func BasicAuth() fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: viper.GetStringMapString("auth.basic"),
	})
}

func LoggerAccess() fiber.Handler {
	return mlogger.New(mlogger.Config{
		Format:       "${time} - [${latency}][${method}][${status}] ${url}\n",
		TimeFormat:   "2006-01-02T15:04:05.999",
		TimeZone:     "Asia/Bangkok",
		TimeInterval: time.Second,
	})
}

func EndpointSwagger(r fiber.Router) {
	UrlPrefix := fmt.Sprintf("%s/swagger", viper.GetString("app.prefix"))
	UrlSwaggerFile := fmt.Sprintf("%s/docs/openapi-%s.json", UrlPrefix, viper.GetString("version"))
	r.Get(UrlPrefix+"/*", rapidoc.New(rapidoc.Config{
		Title:       "Service API",
		HeaderText:  "Service API",
		RenderStyle: rapidoc.RenderStyle_View,
		SchemaStyle: rapidoc.SchemaStyle_Table,
		SpecURL:     UrlSwaggerFile,
		LogoURL:     `https://cdn-icons-png.flaticon.com/512/2165/2165004.png`,
	}))

}
