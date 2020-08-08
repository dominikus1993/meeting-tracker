package repository

import (
	"context"
	"errors"
	"server/domain/model"
)

type QueryError struct {
	Id  string
	Err error
}

var ErrNotFound = errors.New("Meeting not found")

type MeetingsRepository interface {
	FindById(id string, ctx context.Context) (*model.Meeting, error)
	FindByIds(ids []string, ctx context.Context) ([]*model.Meeting, error)
	Create(states []*model.Meeting, ctx context.Context) error
	Count(ctx context.Context) (int64, error)
}
