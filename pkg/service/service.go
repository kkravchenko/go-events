package service

import (
	"context"
	_ "context"
	"events/pkg/models"
	_ "events/pkg/models"
	"events/pkg/repository"
	"time"
	_ "time"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(ctx context.Context, event string, logType string) error {
	e := &models.Event{
		DateTime: time.Now().UTC(),
		Event:    event,
		LogType:  logType,
	}

	return s.repo.Create(ctx, e)
}
