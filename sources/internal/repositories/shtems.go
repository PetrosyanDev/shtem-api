package repositories

import (
	"shtem-api/sources/internal/core/domain"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
)

type shtemsRepository struct {
	db postgresrepository.ShtemsDB
}

func (p *shtemsRepository) Create(shtemaran *domain.Shtemaran) domain.Error {
	return p.db.Create(shtemaran)
}
func (p *shtemsRepository) FindById(id int64) (*domain.Shtemaran, domain.Error) {
	return p.db.FindById(id)
}
func (p *shtemsRepository) Update(shtemaran *domain.Shtemaran) domain.Error {
	return p.db.Update(shtemaran)
}
func (p *shtemsRepository) Delete(id int64) domain.Error {
	return p.db.Delete(id)
}
func (p *shtemsRepository) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	return p.db.GetShtems()
}
func (p *shtemsRepository) GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error) {
	return p.db.GetShtemsByCategoryId(c_id)
}
func (p *shtemsRepository) GetShtemLinkNames() ([]string, domain.Error) {
	return p.db.GetShtemLinkNames()
}
func (p *shtemsRepository) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return p.db.GetShtemByLinkName(name)
}

func NewShtemsRepository(db postgresrepository.ShtemsDB) *shtemsRepository {
	return &shtemsRepository{db}
}
