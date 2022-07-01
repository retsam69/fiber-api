package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

var (
	UrlPrefix string
)

func Init(app fiber.Router, RegisRoutes ...func(fiber.Router)) {

	// Set Url Prefix in ENV: APP_Prefix
	UrlPrefix = viper.GetString("app.prefix")

	//***** Register Routes Buildin  *****//
	EndpointSwagger(app, UrlPrefix+"/swagger")
	// EndpointMonitor(app, UrlPrefix+"/dashboard")

	rg := app.Group(UrlPrefix,
		CORS(),
		LoggerAccess(),
		// BasicAuth(),
	)

	for i, v := range RegisRoutes {
		log.Info().Msgf("Registor Endpoint: %d", i+1)
		v(rg)
	}
}
