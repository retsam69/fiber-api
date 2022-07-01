package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

type RestAPIGeter interface {
	Get(*fiber.Ctx) error
	GetByID(*fiber.Ctx, string) error
}

type RestAPIAdder interface {
	Add(c *fiber.Ctx) error
}

type RestAPIUpdater interface {
	Update(*fiber.Ctx, string) error
}

type RestAPIPatcher interface {
	Patch(*fiber.Ctx, string) error
}

type RestAPIDeleter interface {
	Detele(*fiber.Ctx, string) error
}

func NewRestApi(fiberRoute fiber.Router, prefixPath string, fn any, fiberHandler ...fiber.Handler) fiber.Router {
	log.Info().Str("path", prefixPath).Msgf("Register Route: %s", prefixPath)
	r := fiberRoute.Group(prefixPath, fiberHandler...)
	if f, ok := fn.(RestAPIGeter); ok {
		r.Get("/", f.Get)
		r.Get("/*", handlerGetParam(f.GetByID))
	}

	if f, ok := fn.(RestAPIAdder); ok {
		r.Post("/", f.Add)
	}

	if f, ok := fn.(RestAPIUpdater); ok {
		r.Put("/*", handlerGetParam(f.Update))
	}

	if f, ok := fn.(RestAPIPatcher); ok {
		r.Put("/*", handlerGetParam(f.Patch))
	}

	if f, ok := fn.(RestAPIDeleter); ok {
		r.Put("/*", handlerGetParam(f.Detele))
	}
	return r
}

func handlerGetParam(f func(*fiber.Ctx, string) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := c.Params("*", "")
		return f(c, p)
	}
}
