//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/app"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/server"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

var serverSet = wire.NewSet(
	server.NewTask,
)

// build App
func newApp(
	task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		newApp,
	))
}
