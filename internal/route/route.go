package route

import (
	"fmt"

	rapidoc "github.com/attapon-th/gofiber-rapidoc"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	RegisRoutes []func(fiber.Router) error
	UrlPrefix   string
)

func Init(app fiber.Router) {

	// controller.Init()

	// Set Url Prefix in ENV: APP_Prefix
	UrlPrefix = viper.GetString("prefix")

	app.Use(UrlPrefix+"/",
		CORS(),
		LoggerAccess(),
	)

	app.Get(UrlPrefix+"/docs/*", rapidoc.New(rapidoc.Config{
		Title:       "Service API",
		HeaderText:  "Service API",
		RenderStyle: rapidoc.RenderStyle_View,
		SchemaStyle: rapidoc.SchemaStyle_Table,
		SpecURL:     fmt.Sprintf("%s/docs/docs/openapi-%s.yaml", UrlPrefix, viper.GetString("version")),
		LogoURL:     "https://indev.moph.go.th/blog/wp-content/uploads/2021/03/logo.png",
	}))

}
