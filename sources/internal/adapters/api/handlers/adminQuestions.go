package handlers

import (
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type adminQuestionHandler struct {
	cfg              *configs.Configs
	questionsService ports.QuestionsService
	adminService     ports.AdminService
}

func (h *adminQuestionHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func NewAdminQuestionHandler(
	cfg *configs.Configs,
	questionsService ports.QuestionsService,
	adminService ports.AdminService,
) *adminQuestionHandler {
	return &adminQuestionHandler{
		cfg,
		questionsService,
		adminService,
	}
}
