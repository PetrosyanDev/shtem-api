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

type adminQuestionHandler struct {
	cfg              *configs.Configs
	questionsService ports.QuestionsService
	shtemsService    ports.ShtemsService
	adminService     ports.AdminService
}

func (h *adminQuestionHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.CreateQuestionRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminQuestionHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("adminQuestionHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Try to find question
		if _, err := h.questionsService.FindQuestion(question); err == nil {
			log.Println("adminQuestionHandler:Exists")
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("ALREADY EXISTS"))
			return
		}

		// Create question
		if err := h.questionsService.Create(question); err != nil {
			log.Printf("adminQuestionHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}
func (h *adminQuestionHandler) Find() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.FindQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminQuestionHandler:Get (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question, h.shtemsService); err != nil {
			log.Printf("adminQuestionHandler:Get1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND QUESTION
		final_q, err := h.questionsService.FindQuestion(question)
		if err != nil {
			log.Printf("adminQuestionHandler:Get2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(final_q)
		dto.WriteResponse(ctx, resp)
	}
}
func (h *adminQuestionHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.UpdateQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("adminQuestionHandler:Update (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("adminQuestionHandler:Update1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// GET ID ADD TO QUESTION
		questionID := ctx.Param("id")
		if questionID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}
		id, _ := strconv.Atoi(questionID)

		// GET QUESTION
		if _, err := h.questionsService.FindByID(int64(id)); err != nil {
			log.Printf("adminQuestionHandler:Update2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("NO ROWS"))
			return
		}

		q, err := h.questionsService.FindQuestion(question)
		if err != nil {
			log.Printf("adminQuestionHandler:Update2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("NO ROWS1"))
			return
		}

		if q.Q_id != int64(id) {
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("IDs DOESN'T MATCH"))
			return
		}

		// UPDATE QUESTION
		if err := h.questionsService.Update(question); err != nil {
			log.Printf("adminQuestionHandler:Update3 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp)
	}
}
func (h *adminQuestionHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		questionID := ctx.Param("id")
		if questionID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(questionID)

		// FIND QUESTION
		question, err := h.questionsService.FindByID(int64(id))
		if err != nil {
			log.Printf("apiHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// DELETE QUESTION
		err = h.questionsService.Delete(int64(id))
		if err != nil {
			log.Printf("apiHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *adminQuestionHandler) FindBajin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.FindQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("apiHandler:Find (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question, h.shtemsService); err != nil {
			log.Printf("apiHandler:Find (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND QUESTION
		final_q, err := h.questionsService.FindBajin(question)
		if err != nil {
			log.Printf("apiHandler:Find (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.BajinResponse)
		resp.SliceFromDomain(final_q)
		dto.WriteResponse(ctx, resp)
	}
}

func NewAdminQuestionHandler(
	cfg *configs.Configs,
	questionsService ports.QuestionsService,
	shtemsService ports.ShtemsService,
	adminService ports.AdminService,
) *adminQuestionHandler {
	return &adminQuestionHandler{
		cfg,
		questionsService,
		shtemsService,
		adminService,
	}
}
