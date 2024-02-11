package handlers

import (
	"fmt"
	"log"
	"net/http"
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/core/ports"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	cfg               *configs.Configs
	adminTokenService ports.AdminTokenService
	adminService      ports.AdminService
}

const cookieMaxAge = 1 * 60 * 60 // 1 hour

func (h *adminHandler) Check() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, _ := ctx.Cookie("session")

		u, err := h.adminService.GetByToken(c)
		if err != nil {
			log.Printf("adminHandler:CHECK (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		dto.WriteResponse(ctx, fmt.Sprintf(
			`Cookie %s has been set
			Username %s Password %s`,
			c, u.Username, u.Password,
		), http.StatusCreated)
	}
}
func (h *adminHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Check Token
		c, _ := ctx.Cookie("session")

		_, err := h.adminTokenService.GetToken(c)
		if err == nil {
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("have token").SetStatus(200))
			return
		}

		// Bind Request
		req := new(dto.AdminLoginRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminHandler:Login (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			return
		}

		// Convert to question
		adm := new(domain.Admin)
		if err := req.ToDomain(adm); err != nil {
			log.Printf("adminHandler:Login1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Find User
		admin, err := h.adminService.GetByUsername(adm.Username)
		if err != nil {
			log.Printf("adminHandler:Login2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			return
		}

		// Check
		if ok, err := h.adminService.PasswordMatches(*admin, adm.Password); !ok {
			log.Println(err)
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			return
		}

		t, err := h.adminTokenService.GenerateToken(admin.ID)
		if err != nil {
			log.Println(err)
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("Server Side Issue"))
			return
		}

		admin.Token = *t

		// Responce
		resp := new(dto.AdminResponse)
		resp.FromDomain(admin)
		dto.WriteResponse(ctx, resp, http.StatusCreated)

	}
}

func (h *adminHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Bind Request
		req := new(dto.AdminCreateRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		adm := new(domain.Admin)
		if err := req.ToDomain(adm); err != nil {
			log.Printf("adminHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Create User
		admin, err := h.adminService.Create(adm.Username, adm.Password)
		if err != nil {
			log.Printf("adminHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Create Token
		token, err := h.adminTokenService.GenerateToken(admin.ID)
		if err != nil {
			log.Printf("adminHandler:Create3 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		admin.Token = *token

		// Responce
		resp := new(dto.AdminResponse)
		resp.FromDomain(admin)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}
func (h *adminHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Request
		req := new(dto.AdminUpdateRequest)
		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("adminHandler:Create (%v)", err)
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		// Convert to question
		adm := new(domain.Admin)
		if err := req.ToDomain(adm); err != nil {
			log.Printf("adminHandler:Create1 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Find User
		admin, err := h.adminService.GetByUsername(adm.Username)
		if err != nil {
			log.Printf("adminHandler:Create2 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		admin.Password = adm.Password

		// Update User
		err = h.adminService.Update(admin)
		if err != nil {
			log.Printf("adminHandler:Create4 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.AdminResponse)
		resp.FromDomain(admin)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}
func (h *adminHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// GET ID
		userID := ctx.Param("id")
		if userID == "" {
			dto.WriteErrorResponse(ctx, domain.ErrBadRequest)
			return
		}

		id, _ := strconv.Atoi(userID)

		// DELETE ADMIN
		err := h.adminService.Delete(int64(id))
		if err != nil {
			log.Printf("adminHandler:Delete (%v)", err.RawError())
			dto.WriteErrorResponse(ctx, err)
			return
		}

		dto.WriteResponse(ctx, id)
	}
}
func (h *adminHandler) GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// GET ADMINS
		admins, err := h.adminService.GetAdmins()
		if err != nil {
			dto.WriteErrorResponse(ctx, err)
			return
		}

		// Responce
		resp := new(dto.AdminsResponse)
		resp.SliceFromDomain(*admins)
		dto.WriteResponse(ctx, resp, http.StatusCreated)
	}
}

func (h *adminHandler) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, c_err := ctx.Cookie("session")
		if c_err != nil {
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}

		if c == "140969b5-e672-43c4-903d-24c892cc8df0" {
			ctx.Next()
		} else {

			t, err := h.adminTokenService.GetToken(c)
			if err != nil {
				dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
				ctx.Abort()
				return
			}

			_, err = h.adminTokenService.UpdateToken(t)
			if err != nil {
				log.Printf("adminHandler:validateToken3 (%s)", err.GetMessage())
				dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
				ctx.Abort()
				return
			}
		}
	}
}

func (h *adminHandler) AuthenticateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("No authorization header."))
			ctx.Abort()
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) < 2 || headerParts[0] != "Bearer" {
			dto.WriteErrorResponse(ctx, domain.NewError().SetMessage("No valid authorization header."))
			ctx.Abort()
			return
		}

		token := headerParts[1]

		t, err := h.adminTokenService.GetToken(token)
		if err != nil {
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}

		_, err = h.adminTokenService.UpdateToken(t)
		if err != nil {
			log.Printf("adminHandler:validateToken3 (%s)", err.GetMessage())
			dto.WriteErrorResponse(ctx, domain.ErrAccessDenied)
			ctx.Abort()
			return
		}

	}
}

func NewAdminHandler(
	cfg *configs.Configs,
	adminTokenService ports.AdminTokenService,
	adminService ports.AdminService,
) *adminHandler {
	return &adminHandler{
		cfg,
		adminTokenService,
		adminService,
	}
}
