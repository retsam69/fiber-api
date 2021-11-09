package controller

import "github.com/gofiber/fiber/v2"

var (
	RegisRoutes []func(fiber.Router)
)

func init() {
	RegisRoutes = append(RegisRoutes, EndpointPing)
}

func Init() {

}

// @Summary
// @Description
// @Success 200 {string} string status
// @Failure default {string} string
// @security BasicAuth
// @Router /ping [get]
func EndpointPing(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}
