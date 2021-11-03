package cmd_api

import (
	"net/url"
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
	ParseBaseURL()
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

func ParseBaseURL() {
	if u, err := url.Parse(viper.GetString("app.baseurl")); err != nil {
		log.Fatal().Err(err).Msgf("error: app.baseurl")
	} else {
		viper.SetDefault("app.hostname", u.Hostname())
		if u.Port() == "" {
			viper.SetDefault("app.port", 80)
		} else {
			viper.SetDefault("app.port", u.Port())
		}

		viper.SetDefault("app.schema", u.Scheme)
		viper.SetDefault("app.prefix", u.Path)
		if viper.GetBool("app.dev") {
			viper.SetDefault("app.listen", "127.0.0.1")
		} else {
			viper.SetDefault("app.listen", "0.0.0.0")
		}
	}
}

func printConfig() {
	for _, k := range viper.AllKeys() {
		log.Debug().Msgf("Config: %s => %v", k, viper.Get(k))
	}
}
