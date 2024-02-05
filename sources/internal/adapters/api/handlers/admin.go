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

		// u, err := h.adminService.Create("Erik", "pass")
		// if err != nil {
		// 	log.Printf("adminHandler:generateToken (%s)", err.GetMessage())
		// 	dto.WriteErrorResponse(ctx, err)
		// 	return
		// }

		t, err := h.adminTokenService.GenerateToken(1)
		if err != nil {
			log.Printf("adminHandler:generateToken (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		ctx.SetCookie("session", t.Token, cookieMaxAge, "/", h.cfg.API.Addr, false, true)
		dto.WriteResponse(ctx, fmt.Sprintf("Cookie %s has been set", t.Token), http.StatusCreated)
	}
}

func (h *adminHandler) Check() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, _ := ctx.Cookie("session")

		u, err := h.adminService.GetByToken(c)
		if err != nil {
			log.Printf("adminHandler:CHECK (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		dto.WriteResponse(ctx, fmt.Sprintf(
			`Cookie %s has been set
			Username %s Password %s`,
			c, u.Username, u.Password,
		), http.StatusCreated)
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
