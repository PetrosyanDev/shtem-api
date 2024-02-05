package postgresrepository

import (
	"context"
	"fmt"
	"log"
	postgresclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/core/domain"
	"time"
)

var adminTableName = "admin"

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
func (a *adminDB) Create(username, password, token string) (*domain.Admin, domain.Error) {

	adm := domain.Admin{}
	adm.Username = username
	adm.Password = password

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s) 
		VALUES ($1, $2)`,
		// INSERT
		adminTableName,
		// ()
		adminTableComponentsNon.username,
		adminTableComponentsNon.password,
	)

	_, err := a.db.Exec(a.ctx, query,
		adm.Username,
		adm.Password,
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
	_, err := q.db.Exec(q.ctx, query,
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
	query := fmt.Sprintf(`
		DELETE FROM %s 
		WHERE %s=$1`,
		adminTableName,
		adminTableComponents.id,
	)
	_, err := q.db.Exec(q.ctx, query,
		id,
	)
	if err != nil {
		return domain.NewError().SetError(err)
	}

	return nil
}

func NewAdminDB(ctx context.Context, db *postgresclient.PostgresDB) *adminDB {
	return &adminDB{ctx, db}
}