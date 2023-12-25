// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"fmt"
	postgreclient "shtem-api/sources/internal/clients/postgresql"
	"shtem-api/sources/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

var questionsTableComponents = struct {
	q_id      string
	shtemaran string
	bajin     string
	mas       string
	q_number  string
	text      string
	options   string
	answers   string
}{
	q_id:      "q_id",
	shtemaran: "shtemaran",
	bajin:     "bajin",
	mas:       "mas",
	q_number:  "q_number",
	text:      "text",
	options:   "options",
	answers:   "answers",
}

var questionsTableName = "questions"

type questionsDB struct {
	ctx context.Context
	db  *postgreclient.PostgresDB
}

// CREATE!
// CREATE!
// CREATE!
func (q *questionsDB) Create(question *domain.Question) domain.Error {

	query := fmt.Sprintf("INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s) VALUES ($1, $2, $3, $4, $5, $6, $7)", questionsTableName,
		questionsTableComponents.shtemaran,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
	)
	_, err := q.db.Exec(q.ctx, query,
		question.ShtemName,
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

	return nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *questionsDB) Update(question *domain.Question) domain.Error {

	query := fmt.Sprintf("UPDATE %s SET %s=$1, %s=$2, %s=$3, %s=$4, %s=$5, %s=$6, %s=$7 WHERE %s=$8", questionsTableName,
		questionsTableComponents.shtemaran,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableComponents.q_id, // for identifying the question to update
	)
	_, err := q.db.Exec(q.ctx, query,
		question.ShtemName,
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
	return nil
}

// DELETE!
// DELETE!
// DELETE!
func (q *questionsDB) Delete(id int) domain.Error {
	// DELETE!
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", questionsTableName,
		questionsTableComponents.q_id,
	)
	_, err := q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

// FINDBYSHTEM
// FINDBYSHTEM
// FINDBYSHTEM
func (q *questionsDB) FindQuestionByNumber(question *domain.Question) (*domain.Question, domain.Error) {

	var result domain.Question

	// FIND!
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s=$1 AND %s=$2 AND %s=$3 AND %s=$4",
		questionsTableComponents.shtemaran,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableName,                 // TABLE NAME
		questionsTableComponents.shtemaran, // WHERE BAJIN =
		questionsTableComponents.bajin,     // WHERE BAJIN =
		questionsTableComponents.mas,       // WHERE MAS =
		questionsTableComponents.q_number,  // WHERE Q_NUMBER =
	)
	err := q.db.QueryRow(q.ctx, query, question.ShtemName, question.Bajin, question.Mas, question.Number).
		Scan(
			&result.ShtemName,
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

	return &result, nil
}

func (q *questionsDB) GetShtemNames() ([]string, domain.Error) {
	var shtemaranNames []string

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s",
		questionsTableComponents.shtemaran,
		questionsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var shtemaranName string
		if err := rows.Scan(&shtemaranName); err != nil {
			return nil, domain.NewError().SetError(err)
		}
		shtemaranNames = append(shtemaranNames, shtemaranName)
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemaranNames, nil
}

func (q *questionsDB) FindBajin(question *domain.Question) ([]*domain.Question, domain.Error) {

	// FIND!
	query := fmt.Sprintf(`
			SELECT %s, %s, %s, %s, %s, %s, %s 
			FROM %s 
			WHERE %s=$1 AND %s=$2
			ORDER BY %s, %s`,
		questionsTableComponents.shtemaran,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableName,                 // TABLE NAME
		questionsTableComponents.shtemaran, // WHERE SHTEM =
		questionsTableComponents.bajin,     // WHERE BAJIN =
		questionsTableComponents.mas,       // Sort by
		questionsTableComponents.q_number,  // Sort by
	)
	rows, err := q.db.Query(q.ctx, query,
		question.ShtemName,
		question.Bajin,
	)
	if err == pgx.ErrNoRows {
		return nil, domain.ErrNoRows
	} else if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	var outputQuestions []*domain.Question

	for rows.Next() {
		var result domain.Question
		// Scan the row data into the result struct
		if err := rows.Scan(
			&result.ShtemName,
			&result.Bajin,
			&result.Mas,
			&result.Number,
			&result.Text,
			&result.Options,
			&result.Answers,
		); err != nil {
			return nil, domain.ErrBadRequest
		}

		outputQuestions = append(outputQuestions, &result)
	}

	return outputQuestions, nil
}

// FINDBYID!
// FINDBYID!
// FINDBYID!
func (q *questionsDB) FindByID(id int) (*domain.Question, domain.Error) {

	var result domain.Question

	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s=$1",
		questionsTableComponents.q_id,
		questionsTableComponents.shtemaran,
		questionsTableComponents.bajin,
		questionsTableComponents.mas,
		questionsTableComponents.q_number,
		questionsTableComponents.text,
		questionsTableComponents.options,
		questionsTableComponents.answers,
		questionsTableName,            // SHTEM NAME
		questionsTableComponents.q_id, // Identifyer
	)
	err := q.db.QueryRow(q.ctx, query, id).
		Scan(
			&result.ID,
			&result.ShtemName,
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

	return &result, nil
}

func NewQuestionsDB(ctx context.Context, db *postgreclient.PostgresDB) *questionsDB {
	return &questionsDB{ctx, db}
}
