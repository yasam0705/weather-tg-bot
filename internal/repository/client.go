package repository

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
)

type ClientRepo interface {
	Update(ctx context.Context, client *entity.Client) error
	Create(ctx context.Context, client *entity.Client) error
	FindOne(ctx context.Context, filter map[string]string) (*entity.Client, error)
}
