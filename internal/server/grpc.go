package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

type GrpcServer struct {
	*grpc.Server
	logger *log.Logger
	host   string
	port   int
}

func NewGrpcServer(logger *log.Logger, conf *viper.Viper) *GrpcServer {
	s := &GrpcServer{
		Server: grpc.NewServer(),
		logger: logger,
		host:   conf.GetString("grpc.host"),
		port:   conf.GetInt("grpc.port"),
	}
	return s
}

func (s *GrpcServer) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		s.logger.Sugar().Fatalf("Failed to listen: %v", err)
	}
	if err = s.Server.Serve(lis); err != nil {
		s.logger.Sugar().Fatalf("Failed to serve: %v", err)
	}
	return nil

}
func (s *GrpcServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	s.Server.GracefulStop()

	s.logger.Info("GrpcServer exiting")

	return nil
}
