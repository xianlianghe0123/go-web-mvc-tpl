package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/helper"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/jwt"
	"github.com/xianlianghe0123/go-web-mvc-tpl/internal/utils/log"

	"go.uber.org/zap"
)

func StrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			logger.WithContext(ctx).Warn("No token", zap.Any("data", map[string]any{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			helper.HandleError(ctx, http.StatusUnauthorized, helper.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("token error", zap.Any("data", map[string]any{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}), zap.Error(err))
			helper.HandleError(ctx, http.StatusUnauthorized, helper.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func NoStrictAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	if userInfo, ok := ctx.MustGet("claims").(*jwt.MyCustomClaims); ok {
		logger.WithValue(ctx, zap.String("UserId", userInfo.UserId))
	}
}
