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
	Finish(id string, ctx context.Context) (*model.Meeting, error)
}
