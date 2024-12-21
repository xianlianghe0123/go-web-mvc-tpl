package service

import (
	"context"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/model/do"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/model/dto"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/repository"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/jwt"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*do.UserDO, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, error)
}

func NewUserService(
	service *Service,
	jwt *jwt.JWT,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		jwt:      jwt,
		Service:  service,
	}
}

type userService struct {
	*Service
	jwt      *jwt.JWT
	userRepo repository.UserRepository
}

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (*do.UserDO, error) {
	panic("implement me")
}

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	panic("implement me")
}
