// Erik Petrosyan ©
package api

import (
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/adapters/api/handlers"
	"shtem-api/sources/internal/adapters/api/middlewares"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// TODO: Auth middleware
func NewAPIRouter(cfg *configs.Configs, questionsHandler handlers.QuestionsHandler) *gin.Engine {

	r := gin.Default()
	middlewares.ApplyCommonMiddlewares(r, cfg)
	apiV1 := r.Group("/api/v1")
	{
		posts := apiV1.Group("/questions")
		posts.POST("/create", questionsHandler.Create())
		posts.POST("/find", questionsHandler.Find())
		posts.POST("/:id/update", questionsHandler.Update())
		posts.DELETE("/:id/delete", questionsHandler.Delete())
	}

	r.NoRoute(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	r.NoMethod(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	return r
}
