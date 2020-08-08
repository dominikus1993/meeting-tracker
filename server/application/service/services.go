package service

import (
	"context"
	"server/domain/model"
)

type MeetingsService interface {
	Start(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error)
}
