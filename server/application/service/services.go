package service

import (
	"context"
	"server/domain/model"
	"server/domain/repository"
)

type MeetingsService interface {
	Start(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error)
	GetAll(ctx context.Context) ([]*model.Meeting, error)
	GetMeeting(id string, ctx context.Context) (*model.Meeting, error)
	Finish(id string, ctx context.Context) (*model.Meeting, error)
}

type myMeetingService struct {
	repo repository.MeetingsRepository
}

func (m myMeetingService) Start(meeting *model.Meeting, ctx context.Context) (*model.Meeting, error) {
	return m.repo.Start(meeting, ctx)
}

func (m myMeetingService) GetAll(ctx context.Context) ([]*model.Meeting, error) {
	return m.repo.GetAll(ctx)
}

func (m myMeetingService) GetMeeting(id string, ctx context.Context) (*model.Meeting, error) {
	return m.repo.GetById(id, ctx)
}

func (m myMeetingService) Finish(id string, ctx context.Context) (*model.Meeting, error) {
	model, err := m.repo.GetById(id, ctx)
	if err != nil {
		return nil, err
	}

	return m.repo.Finish(model, ctx)
}

func NewMyMeetingService(repo repository.MeetingsRepository) *myMeetingService {
	return &myMeetingService{repo: repo}
}
