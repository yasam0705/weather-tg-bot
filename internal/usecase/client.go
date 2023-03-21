package usecase

import (
	"context"
	"fmt"
	"test-tasks/tg-bot/internal/entity"
	"test-tasks/tg-bot/internal/repository"
	"time"

	"github.com/google/uuid"
)

type client struct {
	ctxTimeout time.Duration
	repo       repository.ClientRepo
}

type ClientUseCase interface {
	Update(ctx context.Context, client *entity.Client) error
	Create(ctx context.Context, client *entity.Client) error
	FindOne(ctx context.Context, filter map[string]string) (*entity.Client, error)
	CreateOrUpdate(ctx context.Context, client *entity.Client) (*entity.Client, error)
}

func NewClient(ctxTimeout time.Duration, repo repository.ClientRepo) *client {
	return &client{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (c *client) beforeCreate(client *entity.Client) {
	client.Guid = uuid.New().String()
	client.CreatedAt = time.Now()
	client.UpdatedAt = client.CreatedAt
}

func (c *client) beforeUpdate(client *entity.Client) {
	client.UpdatedAt = time.Now()
}

func (c *client) Update(ctx context.Context, client *entity.Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	c.beforeUpdate(client)
	return c.repo.Update(ctx, client)
}

func (c *client) Create(ctx context.Context, client *entity.Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	c.beforeCreate(client)
	return c.repo.Create(ctx, client)
}

func (c *client) FindOne(ctx context.Context, filter map[string]string) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	return c.repo.FindOne(ctx, filter)
}

func (c *client) CreateOrUpdate(ctx context.Context, m *entity.Client) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	client, err := c.FindOne(ctx, map[string]string{
		"client_id": fmt.Sprintf("%d", m.ClientId),
	})
	if err != nil && err != entity.ErrorNotFound {
		return nil, err
	}
	if client == nil {
		err = c.Create(ctx, m)
		return m, err
	}
	client.FirstName = m.FirstName
	client.LastName = m.LastName
	client.Username = m.Username
	return client, c.Update(ctx, client)
}
