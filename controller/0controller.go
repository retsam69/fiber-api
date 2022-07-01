package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

var (
	RegisRoutes []func(fiber.Router)
)

func Init() []func(fiber.Router) {
	log.Info().Msg("Init Controller")
	//
	// * Code init any more
	NewDatabase()

	return RegisRoutes
}

func NewDatabase() {

}
