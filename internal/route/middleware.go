package route

import (
	"fmt"
	"strings"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
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

func EndpointSwagger(r fiber.Router) {
	UrlSwaggerFile := fmt.Sprintf("%s/docs/openapi-%s.yaml", UrlPrefix, viper.GetString("version"))
	r.Get("/docs/*", rapidoc.New(rapidoc.Config{
		Title:       "Service API",
		HeaderText:  "Service API",
		RenderStyle: rapidoc.RenderStyle_View,
		SchemaStyle: rapidoc.SchemaStyle_Table,
		SpecURL:     UrlSwaggerFile,
		LogoURL:     `https://cdn-icons-png.flaticon.com/512/2165/2165004.png`,
	}))

	r.Get("/swagger/*", swagger.New(swagger.Config{
		DocExpansion: "full",
		URL:          UrlSwaggerFile,
		// Prefill OAuth ClientId on Authorize popup
		// OAuth: &swagger.OAuthConfig{
		// 	AppName:  "OAuth Provider",
		// 	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		// },
		// // Ability to change OAuth2 redirect uri location
		// OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

}
