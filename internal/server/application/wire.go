//go:build wireinject
// +build wireinject

package application

import (
	"ego-demo/internal/server/controller"
	"ego-demo/internal/server/service"
	"github.com/google/wire"
	"github.com/gotomicro/ego"
)

func InitApp(app *ego.Ego) error {
	panic(wire.Build(
		wire.Struct(new(Options), "*"),
		controller.ProviderSet,
		service.ProviderSet,
		ProviderSet,
		initApp,
	))
}
