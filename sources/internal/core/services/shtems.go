package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type shtemsService struct {
	shtemsRepository repositories.ShtemsRepository
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

func NewShtemsService(shtemsRepository repositories.ShtemsRepository) *shtemsService {
	return &shtemsService{shtemsRepository}
}
