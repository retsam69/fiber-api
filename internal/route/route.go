package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/viper"
	"gitlab.com/indev-moph/fiber-api/internal/controller"
)

var (
	UrlPrefix string
)

func Init(app fiber.Router) {

	// controller.Init()

	// Set Url Prefix in ENV: APP_Prefix
	UrlPrefix = viper.GetString("app.prefix")

	app.Use(UrlPrefix+"/",
		CORS(),
		LoggerAccess(),
		BasicAuth(),
	)

	EndpointSwagger(app)

	rg := app.Group(UrlPrefix)
	for i, v := range controller.RegisRoutes {
		log.Info().Msgf("Registor Endpoint: %d", i+1)
		v(rg)
	}

}
