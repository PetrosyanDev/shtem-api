package ports

import "shtem-api/sources/internal/core/domain"

type QuestionsService interface {
	Create(question *domain.Question) domain.Error
	Update(question *domain.Question) domain.Error
	Delete(id int64) domain.Error
	FindBajin(question *domain.Question) ([]*domain.Question, domain.Error)
	FindQuestion(question *domain.Question) (*domain.Question, domain.Error)
	FindByID(id int64) (*domain.Question, domain.Error)
}
