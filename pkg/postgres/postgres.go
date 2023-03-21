package postgres

import (
	"context"
	"errors"
	"fmt"
	"test-tasks/tg-bot/config"
	"test-tasks/tg-bot/internal/entity"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	PgError(err error) error
}

type DB struct {
	PostgresDB
	Builder *builder
}

type postgresDB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) (*DB, error) {
	pool, err := pgxpool.Connect(ctx, generateConnString(cfg))
	if err != nil {
		return nil, err
	}

	return &DB{
		Builder:    newBuilder(),
		PostgresDB: &postgresDB{pool: pool},
	}, nil
}

func generateConnString(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)
}

func (p *postgresDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	r, err := p.pool.Query(ctx, sql, args...)
	return r, p.PgError(err)
}

func (p *postgresDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	r := p.pool.QueryRow(ctx, sql, args...)
	return r
}

func (p *postgresDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	c, err := p.pool.Exec(ctx, sql, args...)
	return c, p.PgError(err)
}

func (p *postgresDB) PgError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.ErrorNotFound
	}
	return err
}
