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

func (q *questionsService) Delete(id int) domain.Error {
	err := q.questionsRepository.Delete(id)
	return err
}

/*
Convert shtemName, Bajin, Mas, Number to full domain.Question Struct
*/
func (q *questionsService) FindQuestionByNumber(question *domain.Question) (*domain.Question, domain.Error) {
	return q.questionsRepository.FindQuestionByNumber(question)
}

func (q *questionsService) FindBajin(question *domain.Question) ([]*domain.Question, domain.Error) {
	return q.questionsRepository.FindBajin(question)
}

func (q *questionsService) FindByID(id int) (*domain.Question, domain.Error) {
	return q.questionsRepository.FindByID(id)
}
func (q *questionsService) GetShtemNames() ([]string, domain.Error) {
	return q.questionsRepository.GetShtemNames()
}

func NewQuestionsService(questionsRepository repositories.QuestionsRepository) *questionsService {
	return &questionsService{questionsRepository}
}
