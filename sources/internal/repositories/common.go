// HRACH_DEV © iMed Cloud Services, Inc.
package repositories

import "shtem-api/sources/internal/core/domain"

type QuestionsRepository interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(question *domain.Question) domain.Error
	FindByID(id string) (*domain.Question, domain.Error)
}
