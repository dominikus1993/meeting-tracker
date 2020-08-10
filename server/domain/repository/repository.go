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
	Start(states *model.Meeting, ctx context.Context) (*model.Meeting, error)
	Finish(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error)
	GetAll(ctx context.Context) ([]*model.Meeting, error)
	GetById(id string, ctx context.Context) (*model.Meeting, error)
	GetId() (*string, error)
}
