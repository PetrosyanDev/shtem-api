// Erik Petrosyan ©
package api

import (
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/adapters/api/middlewares"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

// TODO: Auth middleware
func NewAPIRouter(cfg *configs.Configs, apiHandler ports.APIHandler, adminHandler ports.AdminHandler) *gin.Engine {

	r := gin.Default()
	middlewares.ApplyCommonMiddlewares(r, cfg)

	apiV1 := r.Group("/api/v1")
	{
		posts := apiV1.Group("/questions")

		posts.POST("/create", apiHandler.Create())
		posts.POST("/:id/update", apiHandler.Update())
		posts.DELETE("/:id/delete", apiHandler.Delete())

		posts.POST("/getShtems", apiHandler.GetShtems())
		posts.POST("/find", apiHandler.FindQuestion())
		posts.POST("/findBajin", apiHandler.FindBajin())
	}

	// ADMIN
	admin := apiV1.Group("/admin").Use(adminHandler.ValidateToken())
	{
		admin.GET("/check", adminHandler.Check())

		admin.POST("/create", adminHandler.Create())
		admin.POST("/:id/update", adminHandler.Update())
		admin.POST("/:id/delete", adminHandler.Delete())
	}

	// OTHER
	apiV1.GET("/create", adminHandler.GenerateToken())

	r.NoRoute(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	r.NoMethod(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	return r
}
