package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	postgresclient "shtem-api/sources/internal/clients/postgres"
	"shtem-api/sources/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

var adminTokenTableName = "admin_token"

type adminTokenTable struct {
	id        string
	token     string
	tokenHash string
	createdAt string
	updatedAt string
	expiry    string
	admin_id  string
}

var adminTokenTableComponents = adminTokenTable{
	id:        adminTokenTableName + ".id",
	token:     adminTokenTableName + ".token",
	tokenHash: adminTokenTableName + ".token_hash",
	createdAt: adminTokenTableName + ".created_at",
	updatedAt: adminTokenTableName + ".updated_at",
	expiry:    adminTokenTableName + ".expiry",
	admin_id:  adminTokenTableName + ".admin_id",
}
var adminTokenTableComponentsNon = adminTokenTable{
	id:        "id",
	token:     "token",
	tokenHash: "token_hash",
	createdAt: "created_at",
	updatedAt: "updated_at",
	expiry:    "expiry",
	admin_id:  "admin_id",
}

type adminTokenDB struct {
	ctx context.Context
	db  *postgresclient.PostgresDB
}

func (a *adminTokenDB) GenerateToken(id int64) (*domain.AdminToken, domain.Error) {

	t := domain.AdminToken{}
	t.AdminId = id
	t.Token = uuid.NewString()
	t.CreatedAt = time.Now()
	t.Expiry = time.Now().Add(1 * time.Hour)

	log.Println(t)

	query := fmt.Sprintf(`
		INSERT INTO %s (%s,%s,%s,%s) 
		VALUES ($1, $2, $3, $4)`,
		// INSERT
		adminTokenTableName,
		// ()
		adminTokenTableComponentsNon.token,
		adminTokenTableComponentsNon.createdAt,
		adminTokenTableComponentsNon.expiry,
		adminTokenTableComponentsNon.admin_id,
	)

	_, err := a.db.Exec(a.ctx, query,
		t.Token,
		t.CreatedAt,
		t.Expiry,
		t.AdminId,
	)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	return &t, nil
}

func (a *adminTokenDB) GetToken(token string) (*domain.AdminToken, domain.Error) {

	query := fmt.Sprintf(`
		SELECT %s, %s, %s, %s, %s, %s
		FROM %s 
		WHERE %s=$1
		LIMIT 1`,
		// SELECT
		adminTokenTableComponents.id,
		adminTokenTableComponents.token,
		adminTokenTableComponents.tokenHash,
		adminTokenTableComponents.createdAt,
		adminTokenTableComponents.updatedAt,
		adminTokenTableComponents.expiry,
		// adminTokenTableComponents.admin_id
		// FROM
		adminTokenTableName,
		// WHERE
		adminTokenTableComponents.token,
	)

	t := new(domain.AdminToken)

	var created_at, updated_at, expiry sql.NullTime

	err := a.db.QueryRow(a.ctx, query,
		token,
	).Scan(
		&t.Id,
		&t.Token,
		&t.TokenHash,
		&created_at,
		&updated_at,
		&expiry,
	)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	t.CreatedAt = created_at.Time
	t.UpdatedAt = updated_at.Time
	t.Expiry = expiry.Time

	if t.Expiry.Before(time.Now()) {
		return nil, domain.ErrAccessDenied
	}

	return t, nil
}

func (a *adminTokenDB) UpdateToken(t *domain.AdminToken) (*domain.AdminToken, domain.Error) {

	query := fmt.Sprintf(`
		UPDATE %s
		SET %s=$1, %s=$2
		WHERE %s=$3`,
		// UPDATE
		adminTokenTableName,
		// SET
		adminTokenTableComponentsNon.updatedAt,
		adminTokenTableComponentsNon.expiry,
		// WHERE
		adminTokenTableComponentsNon.token,
	)

	t.UpdatedAt = time.Now()
	t.Expiry = time.Now().Add(1 * time.Hour)

	_, err := a.db.Exec(a.ctx, query,
		t.UpdatedAt,
		t.Expiry,
		t.Token,
	)
	if err != nil {
		log.Println(err)
		return nil, domain.NewError().SetError(err)
	}

	return t, nil
}

func NewAdminTokenDB(ctx context.Context, db *postgresclient.PostgresDB) *adminTokenDB {
	return &adminTokenDB{ctx, db}
}
