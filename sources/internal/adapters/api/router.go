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
func NewAPIRouter(
	cfg *configs.Configs,
	apiHandler ports.APIHandler,
	adminHandler ports.AdminHandler,
	adminQuestionHandler ports.AdminQuestionsHandler,
	adminShtemsHandler ports.AdminShtemsHandler,
) *gin.Engine {

	r := gin.Default()
	middlewares.ApplyCommonMiddlewares(r, cfg)

	apiV1 := r.Group("/api/v1")
	{
		posts := apiV1.Group("/questions")

		posts.POST("/create", apiHandler.Create())
		posts.POST("/:id/update", apiHandler.Update())
		posts.DELETE("/:id/delete", apiHandler.Delete())

		posts.POST("/getShtems", apiHandler.GetShtems())
		posts.POST("/findBajin", apiHandler.FindBajin())

		// LOGIN
		apiV1.POST("/login", adminHandler.Login())
	}

	// OTHER

	// ADMIN
	admin := apiV1.Group("/admin")
	// SECURITY
	admin.Use(adminHandler.ValidateToken())
	{
		admin.GET("/check", adminHandler.Check())
	}
	{
		admins := admin.Group("/admins")

		admins.POST("/create", adminHandler.Create())
		admins.POST("/get", adminHandler.GetUsers())
		admins.POST("/update", adminHandler.Update())
		admins.DELETE("/:id/delete", adminHandler.Delete())
	}
	{
		questions := admin.Group("/questions")

		questions.POST("/create", adminQuestionHandler.Create())
		questions.POST("/:id", adminQuestionHandler.Find())
		questions.POST("/:id/update", adminQuestionHandler.Update())
		questions.DELETE("/:id/delete", adminQuestionHandler.Delete())

		questions.POST("/find-bajin", adminQuestionHandler.FindBajin())
	}
	{
		shtems := admin.Group("/shtems")

		shtems.POST("/create", adminShtemsHandler.Create())
		shtems.POST("/:id", adminShtemsHandler.Find())
		shtems.POST("/:id/update", adminShtemsHandler.Update())
		shtems.POST("/:id/cover/upload", adminShtemsHandler.Cover())
		shtems.DELETE("/:id/delete", adminShtemsHandler.Delete())
	}

	r.NoRoute(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	r.NoMethod(func(ctx *gin.Context) { dto.WriteErrorResponse(ctx, domain.ErrRequestPath) })
	return r
}
