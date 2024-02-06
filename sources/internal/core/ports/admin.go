package ports

import (
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AdminService interface {
	// CRUD
	Create(username, password string) (*domain.Admin, domain.Error)
	GetByToken(token string) (*domain.Admin, domain.Error)
	Update(adm *domain.Admin) domain.Error
	Delete(id int64) domain.Error

	GetByUsername(username string) (*domain.Admin, domain.Error)
	GetAdmins() (*[]*domain.Admin, domain.Error)
}

type AdminHandler interface {
	GenerateToken() gin.HandlerFunc
	ValidateToken() gin.HandlerFunc

	Check() gin.HandlerFunc
	Login() gin.HandlerFunc

	Create() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type AdminQuestionsHandler interface {
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc

	FindBajin() gin.HandlerFunc
}
