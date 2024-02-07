package repositories

import (
	"shtem-api/sources/internal/core/domain"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
)

type categoriesRepository struct {
	db postgresrepository.CategoriesDB
}

func (p *categoriesRepository) Create(category *domain.Category) domain.Error {
	return p.db.Create(category)
}
func (p *categoriesRepository) FindById(id int64) (*domain.Category, domain.Error) {
	return p.db.FindById(id)
}
func (p *categoriesRepository) Update(category *domain.Category) domain.Error {
	return p.db.Update(category)
}
func (p *categoriesRepository) Delete(id int64) domain.Error {
	return p.db.Delete(id)
}
func (p *categoriesRepository) GetCategories() ([]*domain.Category, domain.Error) {
	return p.db.GetCategories()
}
func (p *categoriesRepository) GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error) {
	return p.db.GetCategoryByLinkName(c_link_name)
}
func (p *categoriesRepository) GetCategoriesWithShtems() (domain.Categories, domain.Error) {
	return p.db.GetCategoriesWithShtems()
}
func (p *categoriesRepository) GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error) {
	return p.db.GetShtemsByCategoryLinkName(c_linkName)
}

func NewCategoriesRepository(db postgresrepository.CategoriesDB) *categoriesRepository {
	return &categoriesRepository{db}
}
