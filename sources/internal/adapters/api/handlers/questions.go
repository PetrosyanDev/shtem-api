package handlers

import (
	"log"
	"net/http"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/ports"

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

// func (h *questionsHandler) Update() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// Bind Request
// 		req := new(dto.UpdateQuestionRequest)
// 		if err := ctx.BindJSON(req); err != nil {
// 			log.Printf("questionsHandler:Update (%v)", err)
// 			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
// 			return
// 		}

// 		// GET ID
// 		questionID := ctx.Param("id")
// 		if questionID == "" {
// 			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
// 			return
// 		}

// 		question, err := h.questionsService.FindByID(questionID)
// 		if err != nil {
// 			log.Printf("questionsHandler:Update (%v)", err.RawError())
// 			dto.WriteErrorResponse(ctx, err)
// 			return
// 		}
// 		if err := req.ToDomain(question); err != nil {
// 			log.Printf("questionsHandler:Update (%v)", err.RawError())
// 			dto.WriteErrorResponse(ctx, err)
// 			return
// 		}
// 		if err := h.questionsService.Update(question); err != nil {
// 			log.Printf("questionsHandler:Update (%v)", err.RawError())
// 			dto.WriteErrorResponse(ctx, err)
// 			return
// 		}
// 		resp := new(dto.questionResponse)
// 		resp.FromDomain(question)
// 		resp.FormatURLs(h.cfg, question)
// 		dto.WriteResponse(ctx, resp)
// 	}
// }
