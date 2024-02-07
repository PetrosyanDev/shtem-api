package ports

import "shtem-api/sources/internal/core/domain"

type CategoriesService interface {
	Create(category *domain.Category) domain.Error
	FindById(id int64) (*domain.Category, domain.Error)
	Update(category *domain.Category) domain.Error
	Delete(id int64) domain.Error

	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
	GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error)
}
