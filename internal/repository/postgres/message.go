package postgres

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	postgres_pkg "test-tasks/tg-bot/pkg/postgres"
	"time"
)

type messageRepo struct {
	tableName string
	db        *postgres_pkg.DB
}

func NewMessageRepo(db *postgres_pkg.DB) *messageRepo {
	return &messageRepo{
		tableName: "message",
		db:        db,
	}
}

func (c *messageRepo) FindOne(ctx context.Context, filter map[string]string) (*entity.Message, error) {
	query := c.db.Builder.Select("guid", "client_id", "text", "created_at").From(c.tableName)

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

	client := &entity.Message{}
	err = row.Scan(
		&client.Guid,
		&client.ClientId,
		&client.Text,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, c.db.PgError(err)
	}

	return client, nil
}

func (c *messageRepo) Create(ctx context.Context, client *entity.Message) error {
	m := map[string]interface{}{
		"guid":       client.Guid,
		"client_id":  client.ClientId,
		"text":       client.Text,
		"created_at": client.CreatedAt,
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

func (c *messageRepo) CountRequests(ctx context.Context, clientID int64) (uint64, error) {
	query := c.db.Builder.Select(c.db.Builder.Count(1)).From(c.tableName).Where(c.db.Builder.Equal("client_id", clientID))

	sql, args, err := query.ToSQL()
	if err != nil {
		return 0, err
	}
	var res uint64

	row := c.db.QueryRow(ctx, sql, args...)
	err = row.Scan(&res)
	if err != nil {
		return 0, c.db.PgError(err)
	}

	return res, nil
}

func (c *messageRepo) RequestsData(ctx context.Context, clientID int64) (time.Time, time.Time, error) {
	var firstDate, lastDate time.Time
	query := c.db.Builder.
		Select(
			"created_at",
		).
		From(c.tableName).
		Where(c.db.Builder.Equal("client_id", clientID)).
		Limit(1).
		Order(
			c.db.Builder.OrderByAsc("created_at"),
		)

	// query.Order(c.db.Builder.OrderByAsc("created_at"))
	query = query.UnionAll(query.ClearOrder().Order(c.db.Builder.OrderByDesc("created_at")))

	sql, args, err := query.ToSQL()
	if err != nil {
		return firstDate, lastDate, err
	}

	rows, err := c.db.Query(ctx, sql, args...)
	if err != nil {
		return firstDate, lastDate, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&firstDate)
	if err != nil {
		return firstDate, lastDate, err
	}
	rows.Next()
	err = rows.Scan(&lastDate)
	if err != nil {
		return firstDate, lastDate, c.db.PgError(err)
	}

	return firstDate, lastDate, nil
}
