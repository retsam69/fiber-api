package cmd_api

import (
	"strings"

	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

func Init() {
	LoadConfigByFile("default.yaml", true)
	if fileConfig := viper.GetString("config"); fileConfig != "" {
		LoadConfigByFile(fileConfig, false)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()
	printConfig()
}

func LoadConfigByFile(filename string, isDefault bool) {
	v := viper.New()
	v.SetConfigFile(filename)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("defautl config error.")
	}
	for _, name := range v.AllKeys() {
		if isDefault {
			viper.SetDefault(name, v.Get(name))
		} else {
			viper.Set(name, v.Get(name))
		}

	}

}

func printConfig() {
	for _, k := range viper.AllKeys() {
		log.Debug().Msgf("Config: %s => %v", k, viper.Get(k))
	}
}
