package postgresclient

import (
	"context"
	"fmt"
	"log"
	"shtem-api/sources/internal/configs"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbScheme = "postgresql://"
)

type PostgresDB struct {
	*pgxpool.Pool
}

func (p *PostgresDB) Stop() error {

	p.Close()
	log.Println("disconnected from PostgresDB")

	return nil

}

func NewPostgresDBConn(ctx context.Context, cfg *configs.Configs) (*PostgresDB, error) {
	log.Println("connecting to PostgreSQL")
	pool, err := connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	err = ping(pool)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{Pool: pool}, nil
}

func connect(ctx context.Context, cfg *configs.Configs) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(fmt.Sprintf(
		"%s%s:%s/%s?user=%s&password=%s",
		dbScheme, cfg.PostgresDB.Address, cfg.PostgresDB.Port, cfg.PostgresDB.DB, cfg.PostgresDB.User, cfg.PostgresDB.Pass,
	))
	if err != nil {
		return nil, err
	}

	connConfig.MaxConns = 20 // Set the maximum number of connections as needed

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func ping(pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := conn.Conn().Ping(ctx); err != nil {
		return err
	}

	return nil
}
