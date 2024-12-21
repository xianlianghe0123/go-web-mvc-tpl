//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/app"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/handler"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/repository"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/server"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/service"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/jwt"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/sid"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *server.HttpServer,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
