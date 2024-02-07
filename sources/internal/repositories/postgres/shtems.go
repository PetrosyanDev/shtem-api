// Erik Petrosyan ©
package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	postgresclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/core/domain"
)

var shtemsTableName = "shtems"

type shtemsTable struct {
	id          string
	name        string
	description string
	author      string
	link_name   string
	image       string
	pdf         string
	category    string
	keywords    string
}

var shtemsTableComponents = shtemsTable{
	id:          shtemsTableName + ".id",
	name:        shtemsTableName + ".name",
	description: shtemsTableName + ".description",
	author:      shtemsTableName + ".author",
	link_name:   shtemsTableName + ".link_name",
	image:       shtemsTableName + ".image",
	pdf:         shtemsTableName + ".pdf",
	keywords:    shtemsTableName + ".keywords",
	category:    shtemsTableName + ".category",
}
var shtemsTableComponentsNon = shtemsTable{
	id:          "id",
	name:        "name",
	description: "description",
	author:      "author",
	link_name:   "link_name",
	image:       "image",
	pdf:         "pdf",
	keywords:    "keywords",
	category:    "category",
}

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) Create(shtemaran *domain.Shtemaran) domain.Error {

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING %s`,
		shtemsTableName,
		// INTO
		shtemsTableComponentsNon.name,
		shtemsTableComponentsNon.description,
		shtemsTableComponentsNon.author,
		shtemsTableComponentsNon.link_name,
		shtemsTableComponentsNon.image,
		shtemsTableComponentsNon.pdf,
		shtemsTableComponentsNon.keywords,
		shtemsTableComponentsNon.category,
		// RETURNING
		shtemsTableComponentsNon.id,
	)

	err := q.db.QueryRow(q.ctx, query,
		shtemaran.Name,
		shtemaran.Description,
		shtemaran.Author,
		shtemaran.LinkName,
		shtemaran.Image,
		shtemaran.PDF,
		shtemaran.Keywords,
		shtemaran.Category,
	).Scan(&shtemaran.Id)
	if err != nil {
		log.Println(err)
		return domain.NewError().SetError(err)
	}

	return nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *shtemsDB) Update(shtemaran *domain.Shtemaran) domain.Error {

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s=$1, %s=$2, %s=$3, %s=$4, %s=$5, %s=$6, %s=$7, %s=$8 
		WHERE %s=$9`,
		shtemsTableName, // TABLE NAME
		shtemsTableComponentsNon.name,
		shtemsTableComponentsNon.description,
		shtemsTableComponentsNon.author,
		shtemsTableComponentsNon.link_name,
		shtemsTableComponentsNon.image,
		shtemsTableComponentsNon.pdf,
		shtemsTableComponentsNon.keywords,
		shtemsTableComponentsNon.category,
		shtemsTableComponents.id, // for identifying the question to update
	)
	_, err := q.db.Exec(q.ctx, query,
		shtemaran.Name,
		shtemaran.Description,
		shtemaran.Author,
		shtemaran.LinkName,
		shtemaran.Image,
		shtemaran.PDF,
		shtemaran.Keywords,
		shtemaran.Category,
		shtemaran.Id, // for identifying the question to update
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	return nil
}

// DELETE!
// DELETE!
// DELETE!
func (q *shtemsDB) Delete(id int64) domain.Error {
	// DELETE!
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE %s=$1`,
		shtemsTableName,
		shtemsTableComponents.id,
	)
	_, err := q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

func (q *shtemsDB) GetShtemByLinkName(name string) (*domain.Shtemaran, domain.Error) {

	var result *domain.Shtemaran

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s 
		FROM %s 
		WHERE %s=$1
		LIMIT 1`,
		shtemsTableComponents.id,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableName,                 // TABLE NAME
		shtemsTableComponents.link_name, // LINK NAME
	)

	rows, err := q.db.Query(q.ctx, query, name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		var name, description, author, linkName, image, pdf sql.NullString

		if err := rows.Scan(
			&id,
			&name,
			&description,
			&author,
			&linkName,
			&image,
			&pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		result = &domain.Shtemaran{
			Id:          id,
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
		}
	}

	return result, nil
}

func (q *shtemsDB) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s 
		FROM %s`,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableComponents.pdf,
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, description, author, linkName, image, pdf sql.NullString

		if err := rows.Scan(
			&name,
			&author,
			&description,
			&linkName,
			&image,
			&pdf,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		shtemarans = append(shtemarans, &domain.Shtemaran{
			Name:        name.String,
			Description: description.String,
			Author:      author.String,
			LinkName:    linkName.String,
			Image:       image.String,
			PDF:         pdf.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func (q *shtemsDB) GetShtemLinkNames() ([]string, domain.Error) {
	var linkNames []string

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT DISTINCT %s
		FROM %s`,
		shtemsTableComponents.link_name,
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var linkName sql.NullString

		if err := rows.Scan(
			&linkName,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		linkNames = append(linkNames, linkName.String)
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return linkNames, nil
}

func (q *shtemsDB) GetShtemsByCategoryId(c_id int64) ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s
		FROM %s
		JOIN %s
		ON %s = $1`,
		shtemsTableComponents.name,
		shtemsTableComponents.description,
		shtemsTableComponents.author,
		shtemsTableComponents.link_name,
		shtemsTableComponents.image,
		shtemsTableName,                // TABLE NAME
		categoriesTableName,            // JOIN TABLE NAME
		shtemsTableComponents.category, // MATCH
	)

	rows, err := q.db.Query(q.ctx, query, c_id)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, description, author, linkName, image sql.NullString

		if err := rows.Scan(
			&name,
			&author,
			&description,
			&linkName,
			&image,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		shtemarans = append(shtemarans, &domain.Shtemaran{
			Name:        name.String,
			Author:      author.String,
			Description: description.String,
			LinkName:    linkName.String,
			Image:       image.String,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func NewShtemsDB(ctx context.Context, db *postgresclient.PostgresDB) *shtemsDB {
	return &shtemsDB{ctx, db}
}
