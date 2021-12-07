package loader

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	Dev bool = false
)

func Init() {
	loadDefaultConfig()
	if fileConfig := viper.GetString("config"); fileConfig != "" {
		LoadConfigByFile(fileConfig, false)
	}
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv()
	ParseBaseURL()
	printConfig()

	if !viper.GetBool("app.dev") {
		Dev = false
		SetLoggerProduction()
	} else {
		Dev = true
	}

}

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(code).JSON(APIError{
		APIResponse: APIResponse{
			IsError: true,
			Msg:     err.Error(),
		},
		Detail: nil,
	})

}
