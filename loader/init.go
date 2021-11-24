package loader

import "github.com/spf13/viper"

func Init() {
	loadDefaultConfig()
	if fileConfig := viper.GetString("config"); fileConfig != "" {
		LoadConfigByFile(fileConfig, false)
	}
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()
	ParseBaseURL()
	printConfig()
}
