package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type shtemsService struct {
	shtemsRepository repositories.ShtemsRepository
}

func (p *shtemsService) Create(shtemaran *domain.Shtemaran) domain.Error {
	return p.shtemsRepository.Create(shtemaran)
}
func (p *shtemsService) FindById(id int64) (*domain.Shtemaran, domain.Error) {
	return p.shtemsRepository.FindById(id)
}
func (p *shtemsService) Update(shtemaran *domain.Shtemaran) domain.Error {
	return p.shtemsRepository.Update(shtemaran)
}
func (p *shtemsService) Delete(id int64) domain.Error {
	return p.shtemsRepository.Delete(id)
}
func (q *shtemsService) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtems()
}
func (q *shtemsService) GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemsByCategoryId(c_id)
}
func (q *shtemsService) GetShtemLinkNames() ([]string, domain.Error) {
	return q.shtemsRepository.GetShtemLinkNames()
}
func (q *shtemsService) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {
	return q.shtemsRepository.GetShtemByLinkName(name)
}
func (q *shtemsService) GetShtemBajinsByLinkName(link string) ([]*domain.ShtemBajin, domain.Error) {
	return q.shtemsRepository.GetShtemBajinsByLinkName(link)
}

func NewShtemsService(shtemsRepository repositories.ShtemsRepository) *shtemsService {
	return &shtemsService{shtemsRepository}
}
