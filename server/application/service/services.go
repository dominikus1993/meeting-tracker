package service

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
)

type MeetingsService interface {
	Start(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error)
	GetAll(ctx context.Context) ([]*model.Meeting, error)
	GetById(id string, ctx context.Context) (*model.Meeting, error)
	Finish(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error)
}

type myMeetingService struct {
	repo repository.MeetingsRepository
}

func (m myMeetingService) Start(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	panic("implement me")
}

func (m myMeetingService) GetAll(ctx context.Context) ([]*model.Meeting, error) {
	panic("implement me")
}

func (m myMeetingService) GetById(id string, ctx context.Context) (*model.Meeting, error) {
	panic("implement me")
}

func (m myMeetingService) Finish(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	panic("implement me")
}

func NewMyMeetingService(repo repository.MeetingsRepository) *myMeetingService {
	return &myMeetingService{repo: repo}
}
