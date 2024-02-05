// HRACH_DEV © iMed Cloud Services, Inc.
package repositories

import "shtem-api/sources/internal/core/domain"

type QuestionsRepository interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(id int64) domain.Error
	FindBajin(question *domain.Question) ([]*domain.Question, domain.Error)
	FindQuestion(question *domain.Question) (*domain.Question, domain.Error)
	FindByID(id int64) (*domain.Question, domain.Error)
}

type ShtemsRepository interface {
	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}

type CategoriesRepository interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
	GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error)
}

type AdminRepository interface {
	// CRUD
	Create(username, password string) (*domain.Admin, domain.Error)
	GetByToken(token string) (*domain.Admin, domain.Error)
	Update(adm *domain.Admin) domain.Error
	Delete(id int64) domain.Error

	GetByUsername(username string) (*domain.Admin, domain.Error)
	GetAdmins() (*[]*domain.Admin, domain.Error)
}

type AdminTokenRepository interface {
	GenerateToken(id int64) (*domain.AdminToken, domain.Error)
	GetToken(token string) (*domain.AdminToken, domain.Error)
	UpdateToken(t *domain.AdminToken) (*domain.AdminToken, domain.Error)
}
