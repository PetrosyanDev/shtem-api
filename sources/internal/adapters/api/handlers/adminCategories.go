package handlers

import (
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type adminCategoriesHandler struct {
	cfg               *configs.Configs
	shtemsService     ports.ShtemsService
	categoriesService ports.CategoriesService
	adminService      ports.AdminService
}

func (h *adminCategoriesHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func (h *adminCategoriesHandler) Find() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func (h *adminCategoriesHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
func (h *adminCategoriesHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func NewAdminCategoriesHandler(
	cfg *configs.Configs,
	shtemsService ports.ShtemsService,
	categoriesService ports.CategoriesService,
	adminService ports.AdminService,
) *adminCategoriesHandler {
	return &adminCategoriesHandler{
		cfg,
		shtemsService,
		categoriesService,
		adminService,
	}
}
