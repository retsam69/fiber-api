package main

import (
	"github.com/attapon-th/go-pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

var (
	AppName string
	Version string
	Build   string
)

// @title Indev API
// @version 1.1
// @description Get Vaccine Immunization And Vaccine Inventory Stock
// @contact.name Indev-ICT@MOPH Team
// @contact.url https://indev.moph.go.th/blog/
// @contact.email researchmoph@gmail.com
// @host indev.moph.go.th
// @BasePath /api
// @securityDefinitions.basic BasicAuth

func init() {
	// Default Logger `github.com/attapon-th/go-pkg/logger` BaseBy: `github.com/phuslu/log`
	logger.SetDefaultlogger(logger.GetLogger(log.DebugLevel))
}

func main() {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	fConfig := fiber.Config{}
	// production mode
	if !viper.GetBool("dev") {
		fConfig.DisableStartupMessage = true
		log.DefaultLogger.Caller = 0
	}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	app := fiber.New(fConfig)

	// Startup router
	// route.Init(app)

	// Start Server Listener
	if err := app.Listen(viper.GetString("listen")); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
