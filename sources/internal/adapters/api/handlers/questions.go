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

type questionsHandler struct {
	cfg              *configs.Configs
	questionsService ports.QuestionsService
}

func (h *questionsHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.CreateQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("questionsHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("questionsHandler:Create (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Create question
		if err := h.questionsService.Create(question); err != nil {
			log.Printf("questionsHandler:Create (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}

func (h *questionsHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.UpdateQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("questionsHandler:Update (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// GET ID
		questionID := ctx.Param("id")
		if questionID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(questionID)

		// FIND QUESTION
		question, err := h.questionsService.FindByID(id)
		if err != nil {
			log.Printf("questionsHandler:Update (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// UPDATE QUESTION
		if err := h.questionsService.Update(question); err != nil {
			log.Printf("questionsHandler:Update (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *questionsHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// GET ID
		questionID := ctx.Param("id")
		if questionID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(questionID)

		// FIND QUESTION
		question, err := h.questionsService.FindByID(id)
		if err != nil {
			log.Printf("questionsHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// DELETE QUESTION
		err = h.questionsService.Delete(id)
		if err != nil {
			log.Printf("questionsHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(question)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *questionsHandler) FindQuestion() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Bind Request
		req := new(dto.FindQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("questionsHandler:Find (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("questionsHandler:Find (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND QUESTION
		final_q, err := h.questionsService.FindQuestionByNumber(question)
		if err != nil {
			log.Printf("questionsHandler:Find (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.QuestionResponse)
		resp.FromDomain(final_q)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *questionsHandler) FindBajin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Bind Request
		req := new(dto.FindQuestionRequest)
		if err := ctx.BindJSON(req); err != nil {
			log.Printf("questionsHandler:Find (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		question := new(domain.Question)
		if err := req.ToDomain(question); err != nil {
			log.Printf("questionsHandler:Find (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// FIND QUESTION
		final_q, err := h.questionsService.FindBajin(question)
		if err != nil {
			log.Printf("questionsHandler:Find (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.BajinResponse)
		resp.SliceFromDomain(final_q)
		dto.WriteResponse(ctx, resp)
	}
}

func NewQuestionsHandler(
	cfg *configs.Configs,
	questionsService ports.QuestionsService,
) *questionsHandler {
	return &questionsHandler{
		cfg,
		questionsService,
	}
}
