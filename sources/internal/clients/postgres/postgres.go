package postgresclient

import (
	"context"
	"fmt"
	"log"
	"shtem-api/sources/internal/configs"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	dbScheme = "postgresql://"
)

type PostgresDB struct {
	*pgx.Conn
}

func (p *PostgresDB) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := p.Conn.Close(ctx)
	if err == nil {
		log.Println("disconnected form PostgresDB")
	}
	return err
}

func NewPostgresDBConn(ctx context.Context, cfg *configs.Configs) (*PostgresDB, error) {
	log.Println("connecting to PostgreSQL")
	conn, err := connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	err = ping(conn)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{Conn: conn}, nil
}

func connect(ctx context.Context, cfg *configs.Configs) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(fmt.Sprintf(
		"%s%s:%s/%s?user=%s&password=%s",
		dbScheme, cfg.PostgresDB.Address, cfg.PostgresDB.Port, cfg.PostgresDB.DB, cfg.PostgresDB.User, cfg.PostgresDB.Pass,
	))
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ping(client *pgx.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := client.Ping(ctx); err != nil {
		return err
	}
	return nil
}
