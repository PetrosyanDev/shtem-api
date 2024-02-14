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

func (q *questionsService) Delete(id int64) domain.Error {
	err := q.questionsRepository.Delete(id)
	return err
}

func (q *questionsService) FindQuestion(question *domain.Question) (*domain.Question, domain.Error) {
	return q.questionsRepository.FindQuestion(question)
}

func (q *questionsService) FindBajin(question *domain.Question) ([]*domain.Question, domain.Error) {
	return q.questionsRepository.FindBajin(question)
}

func (q *questionsService) FindAllByShtem(shtemId int64) ([]*domain.Question, domain.Error) {
	return q.questionsRepository.FindAllByShtem(shtemId)
}

func (q *questionsService) FindByID(id int64) (*domain.Question, domain.Error) {
	return q.questionsRepository.FindByID(id)
}

func NewQuestionsService(questionsRepository repositories.QuestionsRepository) *questionsService {
	return &questionsService{questionsRepository}
}
