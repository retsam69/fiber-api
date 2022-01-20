package main

import (
	"fmt"
	"os"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
// @BasePath                   /
// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
	_ = pflag.StringP("config", "c", "", "config file path")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// ---- Plaase Uncommant ---- //
	// Serv(
	// 	loader.Init(),
	// 	controller.Init,
	// 	route.Init)
}

func Serv(app *fiber.App, ctl func() []func(fiber.Router), rt func(fiber.Router, ...func(fiber.Router))) {
	// production mode
	// if !viper.GetBool("app.dev") {

	// }

	var RegisRoutes = ctl()
	rt(app, RegisRoutes...)

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
