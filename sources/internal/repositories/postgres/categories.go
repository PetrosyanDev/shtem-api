package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	postgresclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/core/domain"
)

var categoriesTableName = "categories"

type categoriesTable struct {
	c_id        string
	name        string
	description string
	link_name   string
}

var categoriesTableComponents = categoriesTable{
	c_id:        categoriesTableName + ".c_id",
	name:        categoriesTableName + ".name",
	description: categoriesTableName + ".description",
	link_name:   categoriesTableName + ".link_name",
}
var categoriesTableComponentsNon = categoriesTable{
	c_id:        "c_id",
	name:        "name",
	description: "description",
	link_name:   "link_name",
}

type categoriesDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

// CREATE
// CREATE
// CREATE
func (q *categoriesDB) Create(category *domain.Category) domain.Error {

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s,%s) 
		VALUES ($1, $2, $3)
		RETURNING %s`,
		categoriesTableName,
		// INTO
		categoriesTableComponentsNon.name,
		categoriesTableComponentsNon.description,
		categoriesTableComponentsNon.link_name,
		// RETURNING
		categoriesTableComponentsNon.c_id,
	)

	err := q.db.QueryRow(q.ctx, query,
		category.Name,
		category.Description,
		category.LinkName,
	).Scan(&category.C_id)
	if err != nil {
		log.Println(err)
		return domain.NewError().SetError(err)
	}

	return nil
}

func (q *categoriesDB) FindById(id int64) (*domain.Category, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s
		FROM %s 
		WHERE %s=$1`,
		// SELECT
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		// FROM
		categoriesTableName,
		// WHERE
		categoriesTableComponents.c_id, // shtems
	)

	res := domain.Category{}

	err := q.db.QueryRow(q.ctx, query,
		id,
	).Scan(
		&res.Name,
		&res.Description,
		&res.LinkName,
	)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &res, nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *categoriesDB) Update(category *domain.Category) domain.Error {

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s=$1, %s=$2, %s=$3 
		WHERE %s=$4`,
		categoriesTableName, // TABLE NAME
		categoriesTableComponentsNon.name,
		categoriesTableComponentsNon.description,
		categoriesTableComponentsNon.link_name,
		categoriesTableComponents.c_id, // for identifying the question to update
	)
	_, err := q.db.Exec(q.ctx, query,
		category.Name,
		category.Description,
		category.LinkName,
		category.C_id, // for identifying the question to update
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	return nil
}

// DELETE!
// DELETE!
// DELETE!
func (q *categoriesDB) Delete(id int64) domain.Error {
	// DELETE!
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE %s=$1`,
		categoriesTableName,
		categoriesTableComponents.c_id,
	)
	_, err := q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

func (q *categoriesDB) GetCategories() ([]*domain.Category, domain.Error) {
	var categories []*domain.Category

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s
		FROM %s`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		categoriesTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, description, link_name sql.NullString

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&link_name,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		categories = append(categories, &domain.Category{
			C_id:        id,
			Name:        name.String,
			Description: description.String,
			LinkName:    link_name.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return categories, nil
}
func (q *categoriesDB) GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error) {
	var category *domain.Category

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s
		FROM %s
		WHERE %s = $1
		LIMIT 1`,
		categoriesTableComponents.c_id,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		categoriesTableName,                 // TABLE NAME
		categoriesTableComponents.link_name, // WHERE
	)

	rows, err := q.db.Query(q.ctx, query, c_link_name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, description, link_name sql.NullString

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&link_name,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		category = &domain.Category{
			C_id:        id,
			Name:        name.String,
			Description: description.String,
			LinkName:    link_name.String,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return category, nil
}

func (q *categoriesDB) GetCategoriesWithShtems() (domain.Categories, domain.Error) {
	categories := make(domain.Categories)

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT
			COUNT(%s) AS arraysCount,
			%s AS category,
			%s AS c_description,
			%s AS c_link_name,
			ARRAY_AGG(%s) AS names,
			ARRAY_AGG(%s) AS descriptions,
			ARRAY_AGG(%s) AS link_names,
			ARRAY_AGG(%s) AS images,
			ARRAY_AGG(%s) AS authors
		FROM %s
		LEFT JOIN %s
		ON %s = %s
		GROUP BY %s;`,
		// ARRAY_ARG
		categoriesTableComponents.name,
		categoriesTableComponents.name,
		categoriesTableComponents.description,
		categoriesTableComponents.link_name,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.author,
		// FROM TABLE NAME
		categoriesTableName,
		// LEFT JOIN TABLE NAME
		shtemsTableName,
		// ON
		categoriesTableComponents.c_id,
		shtemsTableComponents.category,
		// GROUP BY
		categoriesTableComponents.c_id,
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var arraysCount int
		var category, c_link_name string
		var c_description sql.NullString
		var names, descriptions, link_names, images, authors []sql.NullString

		if err := rows.Scan(
			&arraysCount,
			&category,
			&c_description,
			&c_link_name,
			&names,
			&descriptions,
			&link_names,
			&images,
			&authors,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		for i := 0; i < arraysCount; i++ {
			c := domain.Category{
				Name:        category,
				Description: c_description.String,
				LinkName:    c_link_name,
			}
			s := &domain.Shtemaran{
				Name:        names[i].String,
				Description: descriptions[i].String,
				LinkName:    link_names[i].String,
				Image:       images[i].String,
				Author:      authors[i].String,
			}

			categories[c] = append(categories[c], s)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return categories, nil
}

func (q *categoriesDB) GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error) {

	var result []*domain.Shtemaran

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
		FROM %s
		JOIN %s
		ON %s = %s
		WHERE %s = $1`,
		shtemsTableComponents.id,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableComponents.keywords,
		shtemsTableComponents.category,
		shtemsTableComponents.has_quiz,
		shtemsTableComponents.has_pdf,
		// FROM TABLE NAME
		shtemsTableName,
		// JOIN TABLE NAME
		categoriesTableName,
		// ON
		categoriesTableComponents.c_id,
		shtemsTableComponents.category,
		// LINK NAME
		categoriesTableComponents.link_name,
	)

	rows, err := q.db.Query(q.ctx, query, c_linkName)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, category int64
		var name, description, author, linkName, image, pdf sql.NullString
		var keywords []string
		var has_quiz, has_pdf bool

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
			&keywords,
			&category,
			&has_quiz,
			&has_pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		s := &domain.Shtemaran{
			Id:          id,
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
			Keywords:    keywords,
			Category:    category,
			HasQuiz:     has_quiz,
			HasPDF:      has_pdf,
		}

		result = append(result, s)
	}

	return result, nil
}

func NewCategoriesDB(ctx context.Context, db *postgresclient.PostgresDB) *categoriesDB {
	return &categoriesDB{ctx, db}
}
