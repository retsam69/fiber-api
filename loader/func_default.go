package loader

import "github.com/gofiber/fiber/v2"

const TYPE_CONFIG_DEFAULT = "ini"

var CONFIG_DEFAULT string = `
[app]
dev=1
baseurl=https://localhost:80
maxprocs=2
listen=127.0.0.1

[logger]
outfile=0
log=./logs/log.log
error=./logs/error.log

`

type APIResponse struct {
	IsError bool   `json:"error"` // Response is Error
	Msg     string `json:"msg"`   // Success Message
} // @name APISuccess

type APIError struct {
	IsError bool        `json:"error"`            // Response is Error
	Msg     string      `json:"msg"`              // Success Message
	Detail  interface{} `json:"detail,omitempty"` // Eror Detail or ETC.
} // @name APIError

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(code).JSON(APIError{
		IsError: true,
		Msg:     err.Error(),
		Detail:  nil,
	})

}
