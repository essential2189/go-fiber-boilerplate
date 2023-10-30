package app

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewRouter(app *fiber.App, routes []Route) fiber.Router {
	for _, route := range routes {
		for _, mapping := range route.Table() {
			switch mapping.Method {
			case "GET":
				app.Get(mapping.Path, mapping.Handler)
			case "POST":
				app.Post(mapping.Path, mapping.Handler)
			case "PUT":
				app.Put(mapping.Path, mapping.Handler)
			case "PATCH":
				app.Patch(mapping.Path, mapping.Handler)
			case "DELETE":
				app.Delete(mapping.Path, mapping.Handler)
			}
		}
	}

	return nil
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

type Route interface {
	Table() []Mapping
}

type Mapping struct {
	Method  string
	Path    string
	Handler func(ctx *fiber.Ctx) error
}
