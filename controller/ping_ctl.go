package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

// ? https://github.com/attapon-th/go-valid-struct   --> Validate struct

//!  For Conntroller Template

func init() {

	EndpointRegistor := func(r fiber.Router) {
		log.Info().Str("path", "/ping").Msg("Register Route: /ping ")
		r.Get("/ping", EndpointPing)
	}

	// ? ADD Register Route
	RegisRoutes = append(RegisRoutes, EndpointRegistor)
}

// @Summary      Ping
// @Description  Ping Api Server
// @Tags         General
// @Success      200      {string}  string  status
// @Failure      default  {string}  string
// @security     BasicAuth
// @Router       /ping [get]
func EndpointPing(c *fiber.Ctx) error {
	type Response struct {
		OK bool `json:"ok"`
	}
	return c.JSON(Response{OK: true})
}
