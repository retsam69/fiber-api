package bootloader

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
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
	if viper.GetBool("logger.outfile") {
		SetLoggerProduction()
	}

	SetupFiberApp()

	return FiberApp
}
