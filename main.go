package main

import (
	"fmt"

	logger "github.com/attapon-th/phuslulogger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gitlab.com/indev-moph/fiber-api/cmd/cmd_api"
	"gitlab.com/indev-moph/fiber-api/route"
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

// @title API
// @version 1
// @description API Service
// @contact.name
// @contact.url
// @contact.email
// @host localhost:8888
// @schemes http https
// @BasePath /api
// @securityDefinitions.basic BasicAuth
func main() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	logger.SetDefaultlogger()
	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// ---- Plaase Uncommant ----
	// // load config or etc.
	cmd_api.Init()
	// start http server
	cmd_api.StartFiberServer(route.Init)
}
