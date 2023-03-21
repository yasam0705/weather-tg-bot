package usecase

import (
	"context"
	"test-tasks/tg-bot/internal/entity"
	"time"
)

type stat struct {
	messageUseCase MessageUseCase
	ctxTimeout     time.Duration
}

type Stat interface {
	GetAllStats(ctx context.Context, clientID int64) (*entity.Stat, error)
}

func NewStat(messageUseCase MessageUseCase, ctxTimeout time.Duration) *stat {
	return &stat{
		messageUseCase: messageUseCase,
		ctxTimeout:     ctxTimeout,
	}
}

func (s *stat) GetAllStats(ctx context.Context, clientID int64) (*entity.Stat, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	countRequest, err := s.messageUseCase.CountRequests(ctx, clientID)
	if err != nil {
		return nil, err
	}

	firstDate, lastDate, err := s.messageUseCase.RequestsData(ctx, clientID)
	if err != nil {
		return nil, err
	}

	return &entity.Stat{
		CountRequests: countRequest,
		FirstRequest:  firstDate.Format("2006-01-02 15:04:05"),
		LastRequest:   lastDate.Format("2006-01-02 15:04:05"),
	}, nil
}
