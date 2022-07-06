package restful

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
)

type RestfulGetter interface {
	Get(*fiber.Ctx) error
	GetByID(*fiber.Ctx, string) error
}

type RestfulAdder interface {
	Add(c *fiber.Ctx) error
}

type RestfulUpdater interface {
	Update(*fiber.Ctx, string) error
}

type RestfulPatcher interface {
	Patch(*fiber.Ctx, string) error
}

type RestfulDeleter interface {
	Detele(*fiber.Ctx, string) error
}

func NewRestful(fiberRoute fiber.Router, prefixPath string, fn any, fiberHandler ...fiber.Handler) fiber.Router {
	log.Info().Str("path", prefixPath).Msgf("Register Route: %s", prefixPath)
	r := fiberRoute.Group(prefixPath, fiberHandler...)
	if f, ok := fn.(RestfulGetter); ok {
		r.Get("/", f.Get)
		r.Get("/*", handlerGetParam(f.GetByID))
	}

	if f, ok := fn.(RestfulAdder); ok {
		r.Post("/", f.Add)
	}

	if f, ok := fn.(RestfulUpdater); ok {
		r.Put("/*", handlerGetParam(f.Update))
	}

	if f, ok := fn.(RestfulPatcher); ok {
		r.Put("/*", handlerGetParam(f.Patch))
	}

	if f, ok := fn.(RestfulDeleter); ok {
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
