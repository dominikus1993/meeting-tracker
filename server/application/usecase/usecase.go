package usecase

import (
	"context"
	"server/application/dto"
	"server/application/service"
	"server/domain/model"
	"server/domain/repository"

	"github.com/google/uuid"
)

type MeetingsUseCase interface {
	StartNew(leader string, context context.Context) (*dto.StartedMeetingDto, error)
	GetAll(context context.Context) ([]dto.StartedMeetingDto, error)
}

type meetingsUseCase struct {
	repo    repository.MeetingsRepository
	service service.MeetingsService
}

func (u *meetingsUseCase) GetAll(context context.Context) ([]*dto.StartedMeetingDto, error) {
	result, err := u.service.GetAll(context)
	if err != nil {
		return nil, err
	}
	resultCount := len(result)
	res := make([]*dto.StartedMeetingDto, resultCount)

	for i := 0; i < resultCount; i++ {
		el := result[i]
		res[i] = &dto.StartedMeetingDto{MeetingID: el.MeetingID, Leader: el.Leader, Start: el.Start, Finished: el.Finished}
	}

	return res, nil
}

func NewMeetingsUseCase(repo repository.MeetingsRepository, service service.MeetingsService) *meetingsUseCase {
	return &meetingsUseCase{repo: repo, service: service}
}

func (u *meetingsUseCase) StartNew(leader string, context context.Context) (*dto.StartedMeetingDto, error) {
	id := uuid.New()
	result, err := u.service.Start(model.NewMeeting(id.String(), leader), context)
	if err != nil {
		return nil, err
	}
	return &dto.StartedMeetingDto{MeetingID: result.MeetingID, Leader: result.Leader, Start: result.Start, Finished: false}, nil
}
