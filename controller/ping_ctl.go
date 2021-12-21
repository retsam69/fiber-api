package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

//!  For Conntroller Template

func init() {
	RegisRoutes = append(RegisRoutes, func(r fiber.Router) {
		log.Debug().Str("prefix", "/ping").Msg("Register Route: /ping ")
		r.Get("/ping", EndpointPing)
	})
}

// @Summary Ping
// @Description Ping Api Server
// @Tags      General
// @Success   200      {string}  string  status
// @Failure   default  {string}  string
// @security  BasicAuth
// @Router    /ping [get]
func EndpointPing(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
