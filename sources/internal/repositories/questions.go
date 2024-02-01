package repositories

import (
	"shtem-api/sources/internal/core/domain"
	postgresrepository "shtem-api/sources/internal/repositories/postgres"
)

type questionsRepository struct {
	db postgresrepository.QuestionsDB
}

func (p *questionsRepository) Create(question *domain.Question) domain.Error {
	return p.db.Create(question)
}

func (p *questionsRepository) Update(question *domain.Question) domain.Error {
	return p.db.Update(question)
}

func (p *questionsRepository) Delete(id int64) domain.Error {
	return p.db.Delete(id)
}

func (p *questionsRepository) FindBajin(question *domain.Question) ([]*domain.Question, domain.Error) {
	return p.db.FindBajin(question)
}

func (p *questionsRepository) FindQuestion(question *domain.Question) (*domain.Question, domain.Error) {
	return p.db.FindQuestion(question)
}

func (p *questionsRepository) FindByID(id int64) (*domain.Question, domain.Error) {
	return p.db.FindByID(id)
}

func NewQuestionsRepository(db postgresrepository.QuestionsDB) *questionsRepository {
	return &questionsRepository{db}
}
