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
var shtemBajinsTableName = "shtem_bajins"

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
	has_quiz    string
	has_pdf     string
}

type shtemBajinsTable struct {
	id       string
	shtem_id string
	name     string
	number   string
	is_ready string
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
	has_quiz:    shtemsTableName + ".has_quiz",
	has_pdf:     shtemsTableName + ".has_pdf",
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
	has_quiz:    "has_quiz",
	has_pdf:     "has_pdf",
}

var shtemBajinsTableComponents = shtemBajinsTable{
	id:       shtemBajinsTableName + ".id",
	shtem_id: shtemBajinsTableName + ".shtem_id",
	name:     shtemBajinsTableName + ".name",
	number:   shtemBajinsTableName + ".number",
	is_ready: shtemBajinsTableName + ".is_ready",
}

type shtemsDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (q *shtemsDB) Create(shtemaran *domain.Shtemaran) domain.Error {

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		shtemsTableComponentsNon.has_quiz,
		shtemsTableComponentsNon.has_pdf,
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
		shtemaran.HasQuiz,
		shtemaran.HasPDF,
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
func (q *shtemsDB) FindById(id int64) (*domain.Shtemaran, domain.Error) {

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
		FROM %s 
		WHERE %s=$1`,
		// SELECT
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
		// FROM
		shtemsTableName,
		// WHERE
		shtemsTableComponents.id, // shtems
	)

	res := domain.Shtemaran{}

	err := q.db.QueryRow(q.ctx, query,
		id,
	).Scan(
		&res.Id,
		&res.Name,
		&res.Description,
		&res.Author,
		&res.LinkName,
		&res.Image,
		&res.PDF,
		&res.Keywords,
		&res.Category,
		&res.HasQuiz,
		&res.HasPDF,
	)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &res, nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *shtemsDB) Update(shtemaran *domain.Shtemaran) domain.Error {

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s=$1, %s=$2, %s=$3, %s=$4, %s=$5, %s=$6, %s=$7, %s=$8, %s=$9, %s=$10
		WHERE %s=$11`,
		shtemsTableName, // TABLE NAME
		shtemsTableComponentsNon.name,
		shtemsTableComponentsNon.description,
		shtemsTableComponentsNon.author,
		shtemsTableComponentsNon.link_name,
		shtemsTableComponentsNon.image,
		shtemsTableComponentsNon.pdf,
		shtemsTableComponentsNon.keywords,
		shtemsTableComponentsNon.category,
		shtemsTableComponentsNon.has_quiz,
		shtemsTableComponentsNon.has_pdf,
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
		shtemaran.HasQuiz,
		shtemaran.HasPDF,
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
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
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
		shtemsTableComponents.keywords,
		shtemsTableComponents.category,
		shtemsTableComponents.has_quiz,
		shtemsTableComponents.has_pdf,
		shtemsTableName,                 // TABLE NAME
		shtemsTableComponents.link_name, // LINK NAME
	)

	rows, err := q.db.Query(q.ctx, query, name)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	if rows.Next() {
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

		result = &domain.Shtemaran{
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
	}

	return result, nil
}

func (q *shtemsDB) GetShtems() ([]*domain.Shtemaran, domain.Error) {
	var shtemarans []*domain.Shtemaran

	// FIND DISTINCT SHTEMARAN NAMES
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s 
		FROM %s`,
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
		shtemsTableName, // TABLE NAME
	)

	rows, err := q.db.Query(q.ctx, query)
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

		shtemarans = append(shtemarans, &domain.Shtemaran{
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
		SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s
		FROM %s
		JOIN %s
		ON %s = $1`,
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

		shtemarans = append(shtemarans, &domain.Shtemaran{
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
		})
	}

	if err := rows.Err(); err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return shtemarans, nil
}

func (q *shtemsDB) GetShtemBajinsByLinkName(link string) ([]*domain.ShtemBajin, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s
		FROM %s
		JOIN %s
		ON %s = $1 AND %s=%s`,
		shtemBajinsTableComponents.id,
		shtemBajinsTableComponents.shtem_id,
		shtemBajinsTableComponents.name,
		shtemBajinsTableComponents.number,
		shtemBajinsTableComponents.is_ready,
		shtemBajinsTableName,                // TABLE NAME
		shtemsTableName,                     // JOIN TABLE NAME
		shtemsTableComponents.link_name,     // MATCH
		shtemsTableComponents.id,            // MATCH
		shtemBajinsTableComponents.shtem_id, // MATCH
	)

	rows, err := q.db.Query(q.ctx, query, link)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}
	defer rows.Close()

	var result = []*domain.ShtemBajin{}

	for rows.Next() {
		var id, shtem_id int64
		var name string
		var number int
		var isReady bool

		if err := rows.Scan(
			&id,
			&shtem_id,
			&name,
			&number,
			&isReady,
		); err != nil {
			return nil, domain.NewError().SetError(err)
		}

		result = append(result, &domain.ShtemBajin{
			Id:      id,
			ShtemId: shtem_id,
			Name:    name,
			Number:  number,
			IsReady: isReady,
		})
	}

	return result, nil
}

func NewShtemsDB(ctx context.Context, db *postgresclient.PostgresDB) *shtemsDB {
	return &shtemsDB{ctx, db}
}
