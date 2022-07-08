package controller

import (
	"github.com/phuslu/log"
	"gitlab.com/indev-moph/fiber-api/controller/ctl"
	"gitlab.com/indev-moph/fiber-api/controller/upload_endpoint"
)

func Init() {
	log.Info().Msg("Init Controller")
	//
	// * Code init any more
	ctl.NewConnectionMysql()

	// ! Initialization Endpoint by other package
	upload_endpoint.Init()

}
