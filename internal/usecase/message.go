package usecase

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	"test-tasks/tg-bot/internal/repository"
	"time"

	"github.com/google/uuid"
)

type message struct {
	ctxTimeout time.Duration
	repo       repository.MessageRepo
}

type MessageUseCase interface {
	Create(ctx context.Context, message *entity.Message) error
	FindOne(ctx context.Context, filter map[string]string) (*entity.Message, error)
	CountRequests(ctx context.Context, clientID int64) (uint64, error)
	RequestsData(ctx context.Context, clientID int64) (time.Time, time.Time, error)
}

func NewMessage(ctxTimeout time.Duration, repo repository.MessageRepo) *message {
	return &message{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (c *message) beforeCreate(message *entity.Message) {
	message.Guid = uuid.New().String()
	message.CreatedAt = time.Now()
}

func (c *message) Create(ctx context.Context, message *entity.Message) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	c.beforeCreate(message)
	return c.repo.Create(ctx, message)
}

func (c *message) FindOne(ctx context.Context, filter map[string]string) (*entity.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	return c.repo.FindOne(ctx, filter)
}

func (c *message) CountRequests(ctx context.Context, clientID int64) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	return c.repo.CountRequests(ctx, clientID)
}

func (c *message) RequestsData(ctx context.Context, clientID int64) (time.Time, time.Time, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	return c.repo.RequestsData(ctx, clientID)
}
