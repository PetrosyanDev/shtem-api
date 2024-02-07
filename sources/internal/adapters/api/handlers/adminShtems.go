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
		if _, err := h.shtemsService.GetShtemByLinkName(shtem.LinkName); err != nil {
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
			log.Printf("adminShtemHandler:Get (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		shtem := new(domain.Shtemaran)
		if err := req.ToDomain(shtem); err != nil {
			log.Printf("adminShtemHandler:Get1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND SHTEM
		final_s, err := h.shtemsService.GetShtemByLinkName(shtem.LinkName)
		if err != nil {
			log.Printf("adminShtemHandler:Get2 (%v)", err.RawError())
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
			log.Printf("adminShtemHandler:Get2 (%v)", err.RawError())
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
		// Bind Request
		req := new(dto.UpdateShtemRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminShtemHandler:Update (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// GET ID ADD TO QUESTION
		shtemID := ctx.Param("id")
		if shtemID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}
		id, _ := strconv.Atoi(shtemID)

		// CHECK IF SHTEM EXISTS
		shtemaran, err := h.shtemsService.FindById(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Update2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("DOESNT EXIST"))
			return
		}

		// Convert to shtem
		shtem := new(domain.Shtemaran)
		if err := req.ToDomain(shtem, shtemaran); err != nil {
			log.Printf("adminShtemHandler:Update1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		shtem.Id = int64(id)

		// UPDATE SHTEM
		if err := h.shtemsService.Update(shtem); err != nil {
			log.Printf("adminShtemHandler:Update3 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.ShtemResponse)
		resp.FromDomain(shtem)
		dto.WriteResponse(ctx, resp)
	}
}
func (h *adminShtemsHandler) Cover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(dto.UploadCoverShtemRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminQuestionHandler:UploadBlockMedia (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// GET ID
		shtemID := ctx.Param("id")
		if shtemID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}
		id, _ := strconv.Atoi(shtemID)

		_, err := h.shtemsService.FindById(int64(id))
		if err != nil {
			log.Printf("adminQuestionHandler:UploadBlockMedia (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminQuestionHandler:UploadBlockMedia (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// file, err := h.filesService.CreateFileFromBase64(int64(id), false, req.Data64)
		// TODO
	}
}
func (h *adminShtemsHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		shtemID := ctx.Param("id")
		if shtemID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(shtemID)

		// FIND SHTEM
		shtem, err := h.shtemsService.FindById(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// DELETE SHTEM
		err = h.questionsService.Delete(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.ShtemResponse)
		resp.FromDomain(shtem)
		dto.WriteResponse(ctx, resp)
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
