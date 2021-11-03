package main

import (
	"fmt"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gitlab.com/indev-moph/fiber-api/cmd/cmd_api"
	"gitlab.com/indev-moph/fiber-api/internal/route"
)

const (
	ENV_PREFIX = "APP"
)

var (
	AppName    string
	Version    string
	Build      string
	ConfigFile = pflag.StringP("config", "c", "", "config file path")
)

// @title Indev API
// @version 1.1
// @description API Service
// @contact.name Indev-ICT@MOPH Team
// @contact.url https://indev.moph.go.th/blog/
// @contact.email researchmoph@gmail.com
// @host localhost:8888
// @schemes http https
// @BasePath /api/v2
// @securityDefinitions.basic BasicAuth
func main() {
	SetLogger()
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// load config or etc.
	cmd_api.Init()
	// start http server
	StartFiberServer(route.Init)
}

func SetLogger() {
	// Default Logger `github.com/attapon-th/go-pkg/logger` BaseBy: `github.com/phuslu/log`
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)

}

func StartFiberServer(initFunc ...func(fiber.Router)) {
	fConfig := fiber.Config{}
	// production mode
	if !viper.GetBool("app.dev") {
		fConfig.DisableStartupMessage = true
		log.DefaultLogger = logger.GetLoggerFileAndConsole(
			"logs/log.log", "logs/error.log", log.DebugLevel, 0)
		go logger.RunLogFileRotation()
		log.DefaultLogger = logger.GetLoggerConsole(log.DebugLevel, 0)
	}

	_ = viper.UnmarshalKey("fiber", &fConfig)
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
