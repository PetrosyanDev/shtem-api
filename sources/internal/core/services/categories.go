package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type categoriesService struct {
	categoriesRepository repositories.CategoriesRepository
}

func (q *categoriesService) GetCategories() ([]*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategories()
}
func (q *categoriesService) GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategoryByLinkName(c_link_name)
}
func (q *categoriesService) GetCategoriesWithShtems() (domain.Categories, domain.Error) {
	return q.categoriesRepository.GetCategoriesWithShtems()
}
func (q *categoriesService) GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error) {
	return q.categoriesRepository.GetShtemsByCategoryLinkName(c_linkName)
}

func NewCategoriesService(categoriesRepository repositories.CategoriesRepository) *categoriesService {
	return &categoriesService{categoriesRepository}
}
