package loader

import (
	"github.com/spf13/viper"
)

var (
	Dev bool = false
)

func Init() {
	loadDefaultConfig()
	if fileConfig := viper.GetString("config"); fileConfig != "" {
		LoadConfigByFile(fileConfig, false)
	}
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()
	ParseBaseURL()
	printConfig()

	if !viper.GetBool("app.dev") {
		Dev = false
	} else {
		Dev = true
	}

	if viper.GetBool("logger.outfile") {
		SetLoggerProduction()
	}

}
