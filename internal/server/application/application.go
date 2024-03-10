package application

import (
	"ego-demo/internal/server/controller"
	"github.com/google/wire"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewHttpServer,
	wire.Struct(new(controller.Options), "*"),
)

// Options is server options.
type Options struct {
	http *HttpServer
}

func initApp(app *ego.Ego, opts Options) error {

	// http
	app.Serve(opts.http.ginServer)

	elog.Info("initApp finished")
	return nil
}
