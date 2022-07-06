package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gitlab.com/indev-moph/fiber-api/route/middleware"
	"gitlab.com/indev-moph/fiber-api/route/regisroute"
)

func Init(app fiber.Router) {

	// Set Url Prefix in ENV: APP_Prefix
	regisroute.PathPrefix = viper.GetString("app.prefix")

	//***** Register Routes Milddleware  *****//
	middleware.EndpointSwagger(app, regisroute.JoinPath("/swagger"))
	// EndpointMonitor(app,  regisroute.JoinPath("/dashboard"))

	rg := app.Group(
		regisroute.JoinPath(),
		middleware.CORS(),
		middleware.LoggerAccess(),
		// BasicAuth(),
	)

	for _, fn := range regisroute.RegisRoutes {
		fn(rg)
	}
}
