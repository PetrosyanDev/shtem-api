// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	postgreclient "shtem-api/sources/internal/clients/postgresql"
	"shtem-api/sources/internal/core/domain"
	"shtem-api/sources/internal/repositories/postgres/models"
)

type questionsDB struct {
	ctx context.Context
}

func (s *questionsDB) Create(shtem *domain.Question) domain.Error {
	model := new(models.Question)
	err := model.FromDomain(shtem)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	// INSERT!
	return nil
}

func (s *questionsDB) Update(shtem *domain.Question) domain.Error {
	model := new(models.Question)
	if err := model.FromDomain(shtem); err != nil {
		return domain.NewError().SetError(err)
	}
	// UPDATE!
	return nil
}

func (s *questionsDB) Delete(shtem *domain.Question) domain.Error {
	model := new(models.Question)
	if err := model.FromDomain(shtem); err != nil {
		return domain.NewError().SetError(err)
	}
	// DELETE!
	return nil
}

func NewQuestionsDB(ctx context.Context, db *postgreclient.PostgresDB) *questionsDB {
	return &questionsDB{ctx}
}
