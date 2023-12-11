// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"fmt"
	"log"
	postgreclient "shtem-api/sources/internal/clients/postgresql"
	"shtem-api/sources/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

type questionsDB struct {
	ctx context.Context
	db  *postgreclient.PostgresDB
}

func (q *questionsDB) Create(question *domain.Question) domain.Error {

	// INSERT!
	query := fmt.Sprintf("INSERT INTO %s (bajin,mas,q_number,text,options,answer) VALUES ($1, $2, $3, $4, $5, $6)", question.ShtemName)
	res, err := q.db.Exec(q.ctx, query,
		question.Bajin,
		question.Mas,
		question.Number,
		question.Text,
		question.Options,
		question.Answers,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	rowsAffected := res.RowsAffected()
	log.Printf("Inserted %d rows\n", rowsAffected)
	return nil
}

func (q *questionsDB) Update(question *domain.Question) domain.Error {

	// UPDATE!
	query := fmt.Sprintf("UPDATE %s SET bajin=$1, mas=$2, q_number=$3, text=$4, options=$5, answer=$6 WHERE id=$7", question.ShtemName)
	res, err := q.db.Exec(q.ctx, query,
		question.Bajin,
		question.Mas,
		question.Number,
		question.Text,
		question.Options,
		question.Answers,
		question.ID, // for identifying the question to update
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	rowsAffected := res.RowsAffected()
	log.Printf("Updated %d rows\n", rowsAffected)
	return nil
}

func (q *questionsDB) Delete(question *domain.Question) domain.Error {
	// DELETE!
	query := fmt.Sprintf("DELETE FROM %s WHERE q_id=$1", question.ShtemName)
	_, err := q.db.Exec(q.ctx, query, question.ID)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

func (q *questionsDB) FindByShtem(question *domain.Question) (*domain.Question, domain.Error) {

	var result domain.Question

	// FIND!
	query := fmt.Sprintf("SELECT q_id, bajin, mas, q_number, text, options, answer FROM %s WHERE bajin=$1 AND mas=$2 AND q_number=$3", question.ShtemName)
	err := q.db.QueryRow(q.ctx, query, question.Bajin, question.Mas, question.Number).
		Scan(
			&result.ID,
			&result.Bajin,
			&result.Mas,
			&result.Number,
			&result.Text,
			&result.Options,
			&result.Answers,
		)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNoRows
	} else if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	result.ShtemName = question.ShtemName

	return &result, nil
}

func NewQuestionsDB(ctx context.Context, db *postgreclient.PostgresDB) *questionsDB {
	return &questionsDB{ctx, db}
}
