package fiber_startup

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

var (
	Json        jsoniter.API
	FiberConfig fiber.Config
	FiberApp    *fiber.App
)

func init() {
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
	FiberConfig = fiber.Config{
		JSONEncoder:   Json.Marshal,
		JSONDecoder:   Json.Unmarshal,
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "FiberV2",
	}
}

func SetupFiberApp() {
	maxProcess := viper.GetInt("app.maxprocs")
	if maxProcess > 0 {
		runtime.GOMAXPROCS(maxProcess)
	} else if maxProcess == 0 {
		FiberConfig.Prefork = false
	}
	FiberApp = fiber.New(FiberConfig)
}
