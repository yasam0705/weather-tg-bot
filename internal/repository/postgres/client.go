package postgres

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	postgres_pkg "test-tasks/tg-bot/pkg/postgres"
)

type clientRepo struct {
	tableName string
	db        *postgres_pkg.DB
}

func NewClientRepo(db *postgres_pkg.DB) *clientRepo {
	return &clientRepo{
		tableName: "client",
		db:        db,
	}
}

func (c *clientRepo) FindOne(ctx context.Context, filter map[string]string) (*entity.Client, error) {
	query := c.db.Builder.Select("guid", "client_id", "first_name", "last_name", "username", "created_at", "updated_at").From(c.tableName)

	// c.db.
	for k, v := range filter {
		switch k {
		case "guid", "client_id":
			query.Where(c.db.Builder.Equal(k, v))
		}
	}

	sql, args, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	row := c.db.QueryRow(ctx, sql, args...)

	client := &entity.Client{}
	err = row.Scan(
		&client.Guid,
		&client.ClientId,
		&client.FirstName,
		&client.LastName,
		&client.Username,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, c.db.PgError(err)
	}

	return client, nil
}

func (c *clientRepo) Create(ctx context.Context, client *entity.Client) error {
	m := map[string]interface{}{
		"guid":       client.Guid,
		"client_id":  client.ClientId,
		"first_name": client.FirstName,
		"last_name":  client.LastName,
		"username":   client.Username,
		"created_at": client.CreatedAt,
		"updated_at": client.UpdatedAt,
	}

	query := c.db.Builder.Insert(c.tableName).Rows(m)

	sql, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = c.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *clientRepo) Update(ctx context.Context, client *entity.Client) error {
	m := map[string]interface{}{
		"first_name": client.FirstName,
		"last_name":  client.LastName,
		"username":   client.Username,
		"updated_at": client.UpdatedAt,
	}

	query := c.db.Builder.Update(c.tableName).Set(m).Where(c.db.Builder.Equal("guid", client.Guid))

	sql, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = c.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
