package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"gitlab.com/indev-moph/fiber-api/model/api_response"
	"gitlab.com/indev-moph/fiber-api/route/regisroute"
)

// ? https://github.com/attapon-th/go-valid-struct   --> Validate struct

//!  For Conntroller Template

func init() {
	// ? ADD Register Route
	regisroute.AddRoute(
		func(r fiber.Router) {
			log.Info().Str("path", "/ping").Msg("Register Route: /ping ")
			r.Get("/ping", EndpointPing)
		},
	)
}

// @Summary     Ping
// @Description Ping Api Server
// @Tags        General
// @Success     200     {string} string status
// @Failure     default {string} string
// @security    BasicAuth
// @Router      /ping [get]
func EndpointPing(c *fiber.Ctx) error {
	res := api_response.NewAPIResponse()
	res.OK = true
	res.Msg = "Ping OK!"
	return c.JSON(res)
}
