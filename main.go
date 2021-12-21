package main

import (
	"fmt"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/indev-moph/fiber-api/controller"
	"gitlab.com/indev-moph/fiber-api/loader"
	"gitlab.com/indev-moph/fiber-api/route"
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

	// // ---- Plaase Uncommant ----
	loader.Init()                     // <---- Uncommend Line
	Serv(controller.Init, route.Init) // <---- Uncommend Line
}

func Serv(ctl func() []func(fiber.Router), rt func(fiber.Router, ...func(fiber.Router))) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	fConfig := fiber.Config{
		CaseSensitive: true,
		AppName:       AppName,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		// ErrorHandler:  loader.FiberErrorHandler,
	}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	// production mode
	if !viper.GetBool("app.dev") {
		fConfig.DisableStartupMessage = true
		fConfig.Prefork = true
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
