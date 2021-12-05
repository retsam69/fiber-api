package main

import (
	"fmt"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gitlab.com/indev-moph/fiber-api/loader"
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

// @title        API
// @version      1
// @description  API Service
// @contact.name
// @contact.url
// @contact.email
// @host                       localhost:8888
// @schemes                    http
// @BasePath                   /api
// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// ---- Plaase Uncommant ----
	// load config or etc.
	loader.Init()

	// start http server
	// Serv(controller.Init, route.Init) // <---- Uncommend Line
}

func Serv(ctl func() []func(fiber.Router), rt func(fiber.Router, ...func(fiber.Router))) {
	fConfig := fiber.Config{}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	// production mode
	if !viper.GetBool("app.dev") {
		fConfig.DisableStartupMessage = true
		loader.SetLoggerProduction()
	}

	app := fiber.New(fConfig)
	var RegisRoutes = ctl()
	rt(app, RegisRoutes...)

	log.Info().Msg("start server listener...")
	// Start Server Listener
	ServerLister := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Msgf("Server listener: %s", ServerLister)
	if err := app.Listen(ServerLister); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
