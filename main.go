package main

import (
	"fmt"
	"os"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.com/indev-moph/fiber-api/controller"
	"gitlab.com/indev-moph/fiber-api/fiber_startup"
	"gitlab.com/indev-moph/fiber-api/model/api_response"
	"gitlab.com/indev-moph/fiber-api/route"
)

var (
	AppName string
	Version string
	Build   string
)

// @title        API
// @version      1
// @description  API Service
// @contact.name
// @contact.url
// @contact.email
// @schemes                    http
// @host                       localhost:8888
// @BasePath                   /api
// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
	_ = pflag.StringP("config", "c", "", "config file path")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// * ---- Plaase Uncommant ---- //
	fiber_startup.FiberConfig.ErrorHandler = fiber_startup.NewErrorHandlerJson(api_response.NewAPIError())
	app := fiber_startup.NewFiberApp()
	StartService(app, controller.Init, route.Init)
	// * ---- END Uncommant ---- //
}

func StartService(app *fiber.App, controllerInit func(), routeCreator func(fiber.Router)) {
	controllerInit()
	routeCreator(app)

	if !fiber.IsChild() {
		log.Info().Msg("Parent process")
	} else {
		log.Info().Msgf("Child process pid: %d", os.Getpid())
	}
	// Start Server Listener
	ServerLister := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Msgf("Server listener: %s", ServerLister)
	if err := app.Listen(ServerLister); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
