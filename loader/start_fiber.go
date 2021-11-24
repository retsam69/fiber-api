package loader

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

func StartFiberServer(initFunc ...func(fiber.Router)) {
	fConfig := fiber.Config{}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	// production mode
	if !viper.GetBool("app.dev") {
		fConfig.DisableStartupMessage = true
		SetLoggerProduction()
	}

	app := fiber.New(fConfig)

	for _, n := range initFunc {
		n(app)
	}

	log.Info().Msg("start server listener...")
	// Start Server Listener
	ServerLister := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Msgf("Server listener: %s", ServerLister)
	if err := app.Listen(ServerLister); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
