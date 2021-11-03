package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	UrlPrefix string
)

func Init(app fiber.Router) {

	// controller.Init()

	// Set Url Prefix in ENV: APP_Prefix
	UrlPrefix = viper.GetString("app.prefix")

	EndpointSwagger(app)

	app.Use(UrlPrefix+"/",
		CORS(),
		LoggerAccess(),
	)

	// // ---- Plaase Uncommant ----
	// rg := app.Group(UrlPrefix, BasicAuth())
	// for i, v := range controller.RegisRoutes {
	// 	log.Info().Msgf("Registor Endpoint: %d", i+1)
	// 	v(rg)
	// }

}
