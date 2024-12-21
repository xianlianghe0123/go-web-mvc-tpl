package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/xianlianghe0123/go-web-mvc-tpl/docs"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/handler"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/middleware"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/jwt"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

type HttpServer struct {
	*gin.Engine
	logger *log.Logger
	srv    *http.Server
}

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
) *HttpServer {
	gin.SetMode(gin.DebugMode)
	s := &HttpServer{
		Engine: gin.Default(),
		logger: logger,
	}
	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.GetString("http.host"), conf.GetInt("http.port")),
		Handler: s.Engine,
	}

	// swagger doc
	docs.SwaggerInfo.BasePath = "/g"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	g := s.Group("/xxx")
	{
		// No route group has permission
		noAuthRouter := g.Group("/user")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}

		authRouter := g.Group("", middleware.NoStrictAuth(jwt, logger))
		{
			userRouter := authRouter.Group("/user")
			{
				userRouter.POST("/profile", userHandler.Login)
			}
		}
	}

	return s
}

func (s *HttpServer) Start(ctx context.Context) error {
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}
func (s *HttpServer) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatal("HttpServer forced to shutdown: ", err)
	}

	s.logger.Sugar().Info("HttpServer exiting")
	return nil
}
