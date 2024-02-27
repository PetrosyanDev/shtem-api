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

type apiHandler struct {
	cfg               *configs.Configs
	questionsService  ports.QuestionsService
	shtemsService     ports.ShtemsService
	categoriesService ports.CategoriesService
}

func (h *apiHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.CreateQuestionRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("apiHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("apiHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Create question
		if err := h.questionsService.Create(question); err != nil {
			log.Printf("apiHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}

func (h *apiHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.UpdateQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("apiHandler:Update (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("apiHandler:Create1 (%s)", err.GetMessage())
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
		question.Q_id = int64(id)

		// UPDATE QUESTION
		if err := h.questionsService.Update(question); err != nil {
			log.Printf("apiHandler:Update (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *apiHandler) Delete() gin.HandlerFunc {
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

func (h *apiHandler) FindQuestion() gin.HandlerFunc {
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
		final_q, err := h.questionsService.FindQuestion(question)
		if err != nil {
			log.Printf("apiHandler:Find (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(final_q)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *apiHandler) FindBajin() gin.HandlerFunc {
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

func (h *apiHandler) GetShtems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// FIND shtems
		shtems, err := h.shtemsService.GetShtemLinkNames()
		if err != nil {
			log.Printf("apiHandler:GetShtems (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.ShtemsResponce)
		resp.SliceFromDomain(shtems)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *apiHandler) GetShtemBajins() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		shtemLink := ctx.Param("shtem")
		if shtemLink == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		bajins, err := h.shtemsService.GetShtemBajinsByLinkName(shtemLink)
		if err != nil {
			log.Printf("apiHandler:GetBajins (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		if len(bajins) == 0 {
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("NO ROWS FOUND"))
			return
		}

		resp := new(dto.FullShtemBajinResponse)
		resp.SliceFromDomain(bajins)
		dto.WriteResponse(ctx, resp)
	}
}

func NewAPIHandler(
	cfg *configs.Configs,
	questionsService ports.QuestionsService,
	shtemsService ports.ShtemsService,
	categoriesService ports.CategoriesService,
) *apiHandler {
	return &apiHandler{
		cfg,
		questionsService,
		shtemsService,
		categoriesService,
	}
}
