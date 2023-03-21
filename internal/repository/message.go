package repository

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	"time"
)

type MessageRepo interface {
	Create(ctx context.Context, client *entity.Message) error
	FindOne(ctx context.Context, filter map[string]string) (*entity.Message, error)
	CountRequests(ctx context.Context, clientID int64) (uint64, error)
	RequestsData(ctx context.Context, clientID int64) (time.Time, time.Time, error)
}
