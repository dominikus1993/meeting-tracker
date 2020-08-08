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
	StartNew() (*dto.StartedMeetingDto, error)
}

type meetingsUseCase struct {
	repo    repository.MeetingsRepository
	service service.MeetingsService
}

func (u *meetingsUseCase) StartNew(leader string) (*dto.StartedMeetingDto, error) {
	id := uuid.New()
	ctx := context.Background()
	result, err := u.service.Start(model.NewMeeting(id.String(), leader), ctx)
	if err != nil {
		return nil, err
	}
	return &dto.StartedMeetingDto{MeetingID: result.MeetingID, Leader: result.Leader, Start: result.Start, Finished: false}, nil
}
