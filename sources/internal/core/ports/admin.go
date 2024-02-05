package ports

import (
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	GenerateToken() gin.HandlerFunc
	ValidateToken() gin.HandlerFunc

	Check() gin.HandlerFunc
}

type AdminService interface {
	Create(username, password string) (*domain.Admin, domain.Error)
	GetByToken(token string) (*domain.Admin, domain.Error)
	Update(adm *domain.Admin) domain.Error
	Delete(id int64) domain.Error
}
