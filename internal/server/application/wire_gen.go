// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"ego-demo/internal/server/controller"
	"ego-demo/internal/server/controller/http"
	"ego-demo/internal/server/service/example"
	"github.com/gotomicro/ego"
)

// Injectors from wire.go:

func InitApp(app *ego.Ego) error {
	option := &example.Option{}
	service := example.NewService(option)
	httpOption := &http.Option{
		Service: service,
	}
	handler := http.NewHandler(httpOption)
	options := &controller.Options{
		HttpHandler: handler,
	}
	httpServer := NewHttpServer(options)
	applicationOptions := Options{
		http: httpServer,
	}
	error2 := initApp(app, applicationOptions)
	return error2
}