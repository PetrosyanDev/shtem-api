package ports

import "shtem-api/sources/internal/core/domain"

type ShtemsService interface {
	Create(shtemaran *domain.Shtemaran) domain.Error
	Update(shtemaran *domain.Shtemaran) domain.Error
	Delete(id int64) domain.Error

	GetShtems() ([]*domain.Shtemaran, domain.Error)
	GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error)
	GetShtemLinkNames() ([]string, domain.Error)
	GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error)
}
