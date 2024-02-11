package postgresrepository

import (
	"context"
	"errors"
	"fmt"
	"log"
	postgresclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/core/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var adminTableName = "admin"

const PassCost = 12

type adminTable struct {
	id        string
	createdAt string
	updatedAt string
	username  string
	password  string
}

var adminTableComponents = adminTable{
	id:        adminTableName + ".id",
	createdAt: adminTableName + ".created_at",
	updatedAt: adminTableName + ".updated_at",
	username:  adminTableName + ".username",
	password:  adminTableName + ".password",
}
var adminTableComponentsNon = adminTable{
	id:        "id",
	createdAt: "created_at",
	updatedAt: "updated_at",
	username:  "username",
	password:  "password",
}

type adminDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

// CREATE
// CREATE
// CREATE
func (a *adminDB) Create(username, password string) (*domain.Admin, domain.Error) {

	adm := domain.Admin{}
	adm.Username = username
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), PassCost)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	adm.Password = string(hashedPass)

	query := fmt.Sprintf(`
	INSERT INTO %s (%s, %s) 
	VALUES ($1, $2)
	RETURNING %s`,
		adminTableName,
		adminTableComponentsNon.username,
		adminTableComponentsNon.password,
		adminTableComponentsNon.id,
	)

	err = a.db.QueryRow(a.ctx, query, adm.Username, adm.Password).Scan(&adm.ID)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &adm, nil
}

func (q *adminDB) GetById(id int64) (*domain.Admin, domain.Error) {
	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s
		FROM %s 
		WHERE %s=$1`,
		// SELECT
		adminTableComponents.id,
		adminTableComponents.username,
		adminTableComponents.password,
		adminTableComponents.createdAt,
		adminTableComponents.updatedAt,
		// FROM
		adminTableName,
		// WHERE
		adminTableComponents.id, // shtems
	)

	res := domain.Admin{}

	err := q.db.QueryRow(q.ctx, query,
		id,
	).Scan(
		&res.ID,
		&res.Username,
		&res.Password,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, domain.NewError().SetError(err)
	}

	return &res, nil
}

// GET
// GET
// GET
func (a *adminDB) GetByToken(token string) (*domain.Admin, domain.Error) {

	query := fmt.Sprintf(`
		SELECT %s, %s 
		FROM %s
		JOIN %s
		ON %s=%s
		WHERE %s=$1
		`,
		// SELECT
		adminTableComponents.username,
		adminTableComponents.password,
		// FROM
		adminTableName,
		// JOIN
		adminTokenTableName,
		// ON
		adminTableComponents.id,
		adminTokenTableComponents.admin_id,
		// WHERE
		adminTokenTableComponents.token,
	)

	adm := domain.Admin{}

	err := a.db.QueryRow(a.ctx, query,
		token,
	).Scan(
		&adm.Username,
		&adm.Password,
	)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	return &adm, nil
}

// UPDATE!
// UPDATE!
// UPDATE!
func (q *adminDB) Update(adm *domain.Admin) domain.Error {

	adm.UpdatedAt = time.Now()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(adm.Password), PassCost)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	adm.Password = string(hashedPass)

	query := fmt.Sprintf(`
		UPDATE %s 
		SET %s=$1, %s=$2, %s=$3
		WHERE %s=$4`,
		adminTableName, // TABLE NAME
		adminTableComponentsNon.updatedAt,
		adminTableComponentsNon.username,
		adminTableComponentsNon.password,
		adminTableComponents.id, // for identifying the question to update
	)
	_, err = q.db.Exec(q.ctx, query,
		adm.UpdatedAt,
		adm.Username,
		adm.Password,
		adm.ID,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}
	return nil
}

// DELETE!
// DELETE!
// DELETE!
func (q *adminDB) Delete(id int64) domain.Error {
	// DELETE!
	query_token := fmt.Sprintf(`
		DELETE FROM %s
		WHERE %s=$1`,
		adminTokenTableName,
		adminTokenTableComponents.admin_id,
	)
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE %s=$1`,
		adminTableName,
		adminTableComponents.id,
	)
	_, err := q.db.Exec(q.ctx, query_token,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	_, err = q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

// OTHER
// OTHER
// OTHER
func (a *adminDB) GetByUsername(username string) (*domain.Admin, domain.Error) {

	query := fmt.Sprintf(`
		SELECT %s, %s, %s
		FROM %s
		WHERE %s=$1
		`,
		// SELECT
		adminTableComponents.id,
		adminTableComponents.username,
		adminTableComponents.password,
		// FROM
		adminTableName,
		// WHERE
		adminTableComponents.username,
	)

	adm := domain.Admin{}

	err := a.db.QueryRow(a.ctx, query,
		username,
	).Scan(
		&adm.ID,
		&adm.Username,
		&adm.Password,
	)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	return &adm, nil
}

func (a *adminDB) GetAdmins() (*[]*domain.Admin, domain.Error) {

	var users = []*domain.Admin{}

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s
		FROM %s`,
		// SELECT
		adminTableComponents.id,
		adminTableComponents.createdAt,
		adminTableComponents.updatedAt,
		adminTableComponents.username,
		adminTableComponents.password,
		// FROM
		adminTableName,
	)

	rows, err := a.db.Query(a.ctx, query)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	for rows.Next() {

		adm := domain.Admin{}

		if err := rows.Scan(
			&adm.ID,
			&adm.CreatedAt,
			&adm.UpdatedAt,
			&adm.Username,
			&adm.Password,
		); err != nil {
			log.Println(err)
			return nil, domain.NewError().SetError(err)
		}

		users = append(users, &adm)
	}

	return &users, nil
}

func (a *adminDB) PasswordMatches(usr domain.Admin, plainText string) (bool, domain.Error) {
	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, domain.NewError().SetError(err)
		}
	}

	return true, nil
}

func NewAdminDB(ctx context.Context, db *postgresclient.PostgresDB) *adminDB {
	return &adminDB{ctx, db}
}
