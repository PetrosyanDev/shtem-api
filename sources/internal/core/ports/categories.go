package ports

import "shtem-api/sources/internal/core/domain"

type CategoriesService interface {
	GetCategories() ([]*domain.Category, domain.Error)
	GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error)
	GetCategoriesWithShtems() (domain.Categories, domain.Error)
	GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error)
}
