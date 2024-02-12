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

type adminCategoriesHandler struct {
	cfg               *configs.Configs
	shtemsService     ports.ShtemsService
	categoriesService ports.CategoriesService
	adminService      ports.AdminService
}

func (h *adminCategoriesHandler) All() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// GET Categories
		categories, err := h.categoriesService.GetCategories()
		if err != nil {
			log.Printf("adminShtemHandler:Get2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.CategoriesResponse)
		resp.SliceFromDomain(categories)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *adminCategoriesHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.CreateCategoryRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminShtemHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to category
		category := new(domain.Category)
		if err := req.ToDomain(category); err != nil {
			log.Printf("adminShtemHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Try to find category
		if _, err := h.categoriesService.GetCategoryByLinkName(category.LinkName); err != nil {
			log.Println("adminShtemHandler:Exists")
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("ALREADY EXISTS"))
			return
		}

		// Create category
		if err := h.categoriesService.Create(category); err != nil {
			log.Printf("adminQuestionHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.CategoryResponse)
		resp.FromDomain(category)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}

func (h *adminCategoriesHandler) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		userID := ctx.Param("id")
		if userID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(userID)

		// FIND SHTEM
		final_c, err := h.categoriesService.FindById(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Get2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		final_c.C_id = int64(id)

		resp := new(dto.CategoryResponse)
		resp.FromDomain(final_c)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *adminCategoriesHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.UpdateCategoryRequest)
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
		ct, err := h.categoriesService.FindById(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Update2 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("DOESNT EXIST"))
			return
		}

		// Convert to shtem
		category := new(domain.Category)
		if err := req.ToDomain(category, ct); err != nil {
			log.Printf("adminShtemHandler:Update1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		category.C_id = int64(id)

		// UPDATE SHTEM
		if err := h.categoriesService.Update(category); err != nil {
			log.Printf("adminShtemHandler:Update3 (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		resp := new(dto.CategoryResponse)
		resp.FromDomain(category)
		dto.WriteResponse(ctx, resp)
	}
}

func (h *adminCategoriesHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		shtemID := ctx.Param("id")
		if shtemID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(shtemID)

		// FIND SHTEM
		categ, err := h.categoriesService.FindById(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// DELETE SHTEM
		err = h.categoriesService.Delete(int64(id))
		if err != nil {
			log.Printf("adminShtemHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		categ.C_id = int64(id)

		resp := new(dto.CategoryResponse)
		resp.FromDomain(categ)
		dto.WriteResponse(ctx, resp)
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
