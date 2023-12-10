package services

import (
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories"
)

type questionsService struct {
	questionsRepository repositories.QuestionsRepository
}

func (q *questionsService) Create(question *domain.Question) domain.Error {
	err := q.questionsRepository.Create(question)
	return err
}

func (q *questionsService) Update(question *domain.Question) domain.Error {
	err := q.questionsRepository.Update(question)
	return err
}

func (q *questionsService) Delete(question *domain.Question) domain.Error {
	err := q.questionsRepository.Delete(question)
	return err
}

func (q *questionsService) FindByShtem(question *domain.Question) (*domain.Question, domain.Error) {
	return q.questionsRepository.FindByShtem(question)
}

func NewQuestionsService(questionsRepository repositories.QuestionsRepository) *questionsService {
	return &questionsService{questionsRepository}
}
