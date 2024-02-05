package handlers

import (
	"log"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (h *adminHandler) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, c_err := ctx.Cookie("session")
		if c_err != nil {
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}

		t, err := h.adminTokenService.GetToken(c)
		if err != nil {
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}

		_, err = h.adminTokenService.UpdateToken(t)
		if err != nil {
			log.Printf("adminHandler:validateToken3 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}
	}
}
