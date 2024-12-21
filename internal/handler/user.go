package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/service"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body helper.RegisterRequest true "params"
// @Success 200 {object} helper.Response
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	panic("implement me")
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body helper.LoginRequest true "params"
// @Success 200 {object} helper.LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	panic("implement me")
}
