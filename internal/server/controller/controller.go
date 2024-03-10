package controller

import (
	"ego-demo/internal/server/controller/http"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	http.NewHandler,
	wire.Struct(new(http.Option), "*"),
)

// Options .
type Options struct {
	HttpHandler *http.Handler
}
