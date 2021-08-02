package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/attapon-th/go-pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
// @description Get Vaccine Immunization And Vaccine Inventory Stock
// @contact.name Indev-ICT@MOPH Team
// @contact.url https://indev.moph.go.th/blog/
// @contact.email researchmoph@gmail.com
// @host indev.moph.go.th
// @BasePath /
// @securityDefinitions.basic BasicAuth

func init() {
	fmt.Printf("AppName: %s\nVersion: %s\nBuild: %s\n", AppName, Version, Build)
	// Default Logger `github.com/attapon-th/go-pkg/logger` BaseBy: `github.com/phuslu/log`
	logger.SetDefaultlogger(logger.GetLogger(log.DebugLevel))

	viper.SetDefault("version", Version)
	viper.SetDefault("build", Build)
}

func main() {
	pflag.Parse()
	loadConfigByFile(*ConfigFile)

	// load os env from prefix env
	loadEnvByPrefix("APP_", true)
	loadEnvByPrefix("USER_", false)
	viper.SetEnvPrefix(ENV_PREFIX)
	viper.AutomaticEnv()

	printConfig()

	fConfig := fiber.Config{}
	// production mode
	if !viper.GetBool("dev") {
		fConfig.DisableStartupMessage = true
		log.DefaultLogger = logger.GetLoggerJson(log.DebugLevel)
		log.DefaultLogger.Caller = 0
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

func loadConfigByFile(filename string) {
	v := viper.New()
	v.SetConfigFile(filename)
	if err := v.ReadInConfig(); err != nil {
		log.Warn().Err(err).Msg("defautl config error.")
	}
	for _, k := range v.AllKeys() {
		viper.SetDefault(k, v.Get(k))
	}

}

func loadEnvByPrefix(pf string, isTrim bool) {
	if len(pf) == 0 {
		return
	}
	for _, kv := range os.Environ() {
		if strings.HasPrefix(kv, pf) {
			if isTrim {
				kv = strings.TrimPrefix(kv, pf)
			}
			s := strings.SplitN(kv, "=", 2)
			if len(s) == 2 {
				viper.SetDefault(s[0], s[1])
			}

		}
	}
}

func printConfig() {
	for _, k := range viper.AllKeys() {
		log.Debug().Msgf("Config: %s => %v", k, viper.Get(k))
	}
}
