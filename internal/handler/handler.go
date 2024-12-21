package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/jwt"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(
	logger *log.Logger,
) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*jwt.MyCustomClaims).UserId
}
