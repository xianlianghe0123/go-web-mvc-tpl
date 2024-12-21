package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data any) {
	if data == nil {
		data = map[string]any{}
	}
	resp := &Response{Code: 0, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data any) {
	if data == nil {
		data = map[string]string{}
	}
	e := new(Error)
	ok := errors.As(err, &e)
	if !ok {
		e = ErrUnknownError
	}
	resp := &Response{Code: e.Code(), Message: e.Message(), Data: data}
	ctx.JSON(httpCode, resp)
}
