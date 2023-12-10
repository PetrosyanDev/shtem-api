package dto

import (
	"net/http"
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Data  *T     `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteResponse(ctx *gin.Context, r any, status ...int) {
	if len(status) > 0 {
		ctx.SecureJSON(status[0], r)
		return
	}
	ctx.SecureJSON(http.StatusOK, r)
}

func WriteErrorResponse(ctx *gin.Context, e domain.Error) {
	ctx.SecureJSON(e.GetStatus(), Response[any]{Error: e.GetMessage()})
}
