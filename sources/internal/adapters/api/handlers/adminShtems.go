package handlers

import (
	"log"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
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
		// Bind Request
		req := new(dto.CreateShtemRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminShtemHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to shtem
		shtem := new(domain.Shtemaran)
		if err := req.ToDomain(shtem); err != nil {
			log.Printf("adminShtemHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Try to find shtem
		if _, err := h.shtemsService.GetShtemByLinkName(shtem.LinkName); err == nil {
			log.Println("adminShtemHandler:Exists")
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("ALREADY EXISTS"))
			return
		}

		// Create shtem
		if err := h.shtemsService.Create(shtem); err != nil {
			log.Printf("adminQuestionHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}
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
