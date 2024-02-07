package handlers

import (
	"log"
	"net/http"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/ports"
	"strconv"

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

		// Responce
		resp := new(dto.ShtemResponse)
		resp.FromDomain(shtem)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}

func (h *adminShtemsHandler) FindByLinkName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.FindShtemRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminQuestionHandler:Get (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		shtem := new(domain.Shtemaran)
		if err := req.ToDomain(shtem); err != nil {
			log.Printf("adminQuestionHandler:Get1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND SHTEM
		final_s, err := h.shtemsService.GetShtemByLinkName(shtem.LinkName)
		if err != nil {
			log.Printf("adminQuestionHandler:Get2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.ShtemResponse)
		resp.FromDomain(final_s)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *adminShtemsHandler) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// GET ID
		userID := ctx.Param("id")
		if userID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(userID)

		// FIND SHTEM
		final_s, err := h.shtemsService.FindById(int64(id))
		if err != nil {
			log.Printf("adminQuestionHandler:Get2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.ShtemResponse)
		resp.FromDomain(final_s)
		dto.WriteResponse(ctx, resp)
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
