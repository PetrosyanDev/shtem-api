// HRACH_DEV © iMed Cloud Services, Inc.
package repositories

import "shtem-api/sources/internal/core/domain"

type QuestionsRepository interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(id int) domain.Error
	FindByShtem(question *domain.Question) (*domain.Question, domain.Error)
	FindByID(id int) (*domain.Question, domain.Error)
}
