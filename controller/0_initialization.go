package controller

import (
	"github.com/phuslu/log"
	"gitlab.com/indev-moph/fiber-api/controller/ctl"
	"gitlab.com/indev-moph/fiber-api/controller/ep_upload"
)

func Init() {
	log.Info().Msg("Init Controller")
	//
	// * Code init any more
	ctl.NewConnectionMysql()

	// ! Initialization Endpoint by other package
	ep_upload.Init()

}
