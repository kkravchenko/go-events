package repository

import (
	"context"
	"events/pkg/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(ctx context.Context, e *models.Event) error {
	return r.db.WithContext(ctx).Create(e).Error
}
