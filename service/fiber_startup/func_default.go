package fiber_startup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

const TYPE_CONFIG_DEFAULT = "ini"

var CONFIG_DEFAULT string = `
[app]
dev=true
baseurl=https://localhost:80
maxproces=2
listen=127.0.0.1
[logger]
outfile=false
log=./logs/log.log
error=./logs/error.log
`

type ErrorHandlerJsoner interface {
	SetError(c *fiber.Ctx, code int, err error)
}

func NewErrorHandlerJson(errResponse ErrorHandlerJsoner) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		log.Error().Err(err).Msgf("API Access Error.")
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		if errResponse != nil {
			var code int
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				errResponse.SetError(c, e.Code, err)
			} else {
				code = c.Context().Response.StatusCode()
				errResponse.SetError(c, code, err)
			}
			return c.Status(code).JSON(errResponse)
		} else {
			e := fiber.ErrInternalServerError
			return c.Status(e.Code).JSON(fiber.Map{
				"ok":  false,
				"msg": e.Error(),
			})
		}
	}
}
