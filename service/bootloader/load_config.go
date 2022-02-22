package bootloader

import (
	"net/url"
	"strings"

	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

func loadDefaultConfig() {
	sr := strings.NewReader(CONFIG_DEFAULT)
	v := viper.New()
	v.SetConfigType(TYPE_CONFIG_DEFAULT)
	if err := v.ReadConfig(sr); err != nil {
		log.Fatal().Err(err).Msg("load default configs error.")
	}
	for _, name := range v.AllKeys() {
		viper.SetDefault(name, v.Get(name))
	}

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
		// log.Debug().Str("port", u.Port()).Msg("")
		viper.SetDefault("app.hostname", u.Hostname())
		if u.Port() == "" || u.Port() == "0" {
			viper.SetDefault("app.port", "80")
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
