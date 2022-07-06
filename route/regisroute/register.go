package regisroute

import (
	"path"

	"github.com/gofiber/fiber/v2"
)

var (
	PathPrefix  string = "/"
	RegisRoutes []func(fiber.Router)
)

func AddRoute(fn func(fiber.Router)) {
	RegisRoutes = append(RegisRoutes, fn)
}

func JoinPath(p ...string) string {
	if len(p) > 0 {
		p = append([]string{PathPrefix}, p...)
		return path.Join(p...)
	}
	return PathPrefix
}
