package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
}

func main() {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	fConfig := fiber.Config{
		DisableStartupMessage: !viper.GetBool("dev"),
	}
	_ = viper.UnmarshalKey("fiber", &fConfig)
	app := fiber.New(fConfig)

	// Startup router
	// route.Init(app)

	// Start Server Listener
	if err := app.Listen(viper.GetString("listen")); err != nil {
		log.Panicln(err)
	}
}
