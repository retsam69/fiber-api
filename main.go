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
// @host indev.moph.go.th
// @BasePath /
// @securityDefinitions.basic BasicAuth

func init() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	// Default Logger `github.com/attapon-th/go-pkg/logger` BaseBy: `github.com/phuslu/log`
	logger.SetDefaultlogger()

	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
}

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	cmd_api.Init()
	fConfig := fiber.Config{}
	// production mode
	if !viper.GetBool("dev") {
		fConfig.DisableStartupMessage = true
		log.DefaultLogger = logger.GetLoggerConsole(log.DebugLevel, 0)
	}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	app := fiber.New(fConfig)

	// Startup router
	// route.Init(app)

	log.Info().Msg("start server listener...")
	// Start Server Listener
	if err := app.Listen(viper.GetString("listen")); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
