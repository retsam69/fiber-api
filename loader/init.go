package loader

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	Dev bool = false
)

func Init() *fiber.App {
	loadDefaultConfig()
	if fileConfig := viper.GetString("config"); fileConfig != "" {
		LoadConfigByFile(fileConfig, false)
	}
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()
	ParseBaseURL()
	printConfig()
	Dev = viper.GetBool("app.dev")
	if viper.GetBool("logger.outfile") {
		SetLoggerProduction()
	}

	InitFiber()

	return FiberApp
}
