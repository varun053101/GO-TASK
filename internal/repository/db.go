package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/varun053101/GO-TASK/config"
)

// shared database connection pool
var DB *pgxpool.Pool

// connect to postgres running in docker
func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	// verify connection
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	DB = pool
	return nil
}
