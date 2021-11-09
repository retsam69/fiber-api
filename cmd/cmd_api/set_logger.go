package cmd_api

import (
	"github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

func SetLoggerProduction() {
	log.DefaultLogger = phuslulogger.GetLoggerFileAndConsole(
		viper.GetString("logger.log"),
		viper.GetString("logger.error"),
		log.DebugLevel,
		0)
	go phuslulogger.RunLogFileRotation()
}
