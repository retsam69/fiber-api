package route

import (
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

}
