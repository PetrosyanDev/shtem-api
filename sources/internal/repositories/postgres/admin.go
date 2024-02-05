package postgresrepository

import (
	"context"
	postgresclient "shtem-api/sources/internal/clients/postgres"
)

var adminTableName = "admin"

var adminTableComponents = struct {
	id        string
	createdAt string
	updatedAt string
	username  string
	password  string
}{
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
