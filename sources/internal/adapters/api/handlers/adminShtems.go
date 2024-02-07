package handlers

import (
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type adminShtemsHandler struct {
	cfg              *configs.Configs
	questionsService ports.QuestionsService
	shtemsService    ports.ShtemsService
	adminService     ports.AdminService
}

func (h *adminShtemsHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func (h *adminShtemsHandler) Find() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func (h *adminShtemsHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func (h *adminShtemsHandler) Cover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func (h *adminShtemsHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func NewAdminShtemsHandler(
	cfg *configs.Configs,
	questionsService ports.QuestionsService,
	shtemsService ports.ShtemsService,
	adminService ports.AdminService,
) *adminShtemsHandler {
	return &adminShtemsHandler{
		cfg,
		questionsService,
		shtemsService,
		adminService,
	}
}
