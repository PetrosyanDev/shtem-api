package handlers

import (
	"fmt"
	"log"
	"net/http"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	cfg               *configs.Configs
	adminTokenService ports.AdminTokenService
	adminService      ports.AdminService
}

const cookieMaxAge = 1 * 60 * 60 // 1 hour

func (h *adminHandler) GenerateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		t, err := h.adminTokenService.GenerateToken()
		if err != nil {
			log.Printf("adminHandler:generateToken (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		ctx.SetCookie("session", t.Token, cookieMaxAge, "/", h.cfg.API.Addr, false, true)
		dto.WriteResponse(ctx, fmt.Sprintf("Cookie %s has been set", t.Token), http.StatusCreated)
	}
}

func NewAdminHandler(
	cfg *configs.Configs,
	adminTokenService ports.AdminTokenService,
	adminService ports.AdminService,
) *adminHandler {
	return &adminHandler{
		cfg,
		adminTokenService,
		adminService,
	}
}
