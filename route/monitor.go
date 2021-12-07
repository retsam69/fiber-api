package route

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

func EndpointMonitor(r fiber.Router) {
	UrlPrefix := fmt.Sprintf("%s/dashboard/*", viper.GetString("app.prefix"))
	r.Get(UrlPrefix, monitor.New(monitor.Config{
		APIOnly: false,
	}))
}
