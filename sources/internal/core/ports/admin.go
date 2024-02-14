package ports

import (
	"shtem-api/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type AdminService interface {
	// CRUD
	Create(username, password string) (*domain.Admin, domain.Error)
	GetByToken(token string) (*domain.Admin, domain.Error)
	GetById(id int64) (*domain.Admin, domain.Error)
	Update(adm *domain.Admin) (*domain.Admin, domain.Error)
	Delete(id int64) domain.Error

	PasswordMatches(usr domain.Admin, plainText string) (bool, domain.Error)

	GetByUsername(username string) (*domain.Admin, domain.Error)
	GetAdmins() (*[]*domain.Admin, domain.Error)
}

type AdminHandler interface {
	AuthenticateToken() gin.HandlerFunc

	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc

	Create() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type AdminQuestionsHandler interface {
	Create() gin.HandlerFunc
	Find() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc

	All() gin.HandlerFunc
	FindBajin() gin.HandlerFunc
}

type AdminShtemsHandler interface {
	Create() gin.HandlerFunc
	FindByLinkName() gin.HandlerFunc
	FindById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Cover() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type AdminCategoriesHandler interface {
	All() gin.HandlerFunc
	GetShtems() gin.HandlerFunc

	Create() gin.HandlerFunc
	FindById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
