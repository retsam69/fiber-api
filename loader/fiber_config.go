package loader

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

var (
	Jsoniter    jsoniter.API
	fiberConfig fiber.Config
	FiberApp    *fiber.App
)

func InitFiber() {
	Jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	fiberConfig = fiber.Config{
		JSONEncoder:   Jsoniter.Marshal,
		JSONDecoder:   Jsoniter.Unmarshal,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "FiberV2",
		ErrorHandler:  FiberErrorHandler,
	}

	maxProcess := viper.GetInt("app.maxprocs")
	if maxProcess > 0 {
		runtime.GOMAXPROCS(maxProcess)
	} else if maxProcess == 0 {
		fiberConfig.Prefork = false
	}
	FiberApp = fiber.New(fiberConfig)
}
