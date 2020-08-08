package repository

import (
	"errors"
	"server/domain/model"
)

type QueryError struct {
	Id  string
	Err error
}

var ErrNotFound = errors.New("Meeting not found")

type MeetingsRepository interface {
	FindById(id string) (*model.Meeting, error)
	FindByIds(ids []string) ([]*model.Meeting, error)
	Create(states []*model.Meeting) error
	Count() (int64, error)
}
