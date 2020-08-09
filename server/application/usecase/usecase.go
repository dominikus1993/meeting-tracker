package usecase

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/application/dto"
	"server/application/service"
	"server/domain/model"
)

type MeetingsUseCase interface {
	StartNew(leader string, context context.Context) (*dto.StartedMeetingDto, error)
	GetAll(context context.Context) ([]*dto.StartedMeetingDto, error)
	Finish(id string, c context.Context) (*dto.FinishedMeetingDto, error)
	GetMeeting(id string, c *gin.Context) (*dto.MeetingDto, error)
}

type meetingsUseCase struct {
	service service.MeetingsService
}

func (u *meetingsUseCase) GetMeeting(id string, c *gin.Context) (*dto.MeetingDto, error) {
	result, err := u.service.GetById(id, c)
	if err != nil {
		return nil, err
	}
	return &dto.MeetingDto{MeetingID: result.MeetingID, Leader: result.Leader, Start: result.Start, End: result.Finish, Finished: true}, nil
}

func (u *meetingsUseCase) Finish(id string, c context.Context) (*dto.FinishedMeetingDto, error) {
	result, err := u.service.Finish(id, c)
	if err != nil {
		return nil, err
	}
	return &dto.FinishedMeetingDto{MeetingID: result.MeetingID, Leader: result.Leader, Start: result.Start, End: result.Finish, Finished: true}, nil
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

func (u *meetingsUseCase) StartNew(leader string, context context.Context) (*dto.StartedMeetingDto, error) {
	id := uuid.New()
	result, err := u.service.Start(model.NewMeeting(id.String(), leader), context)
	if err != nil {
		return nil, err
	}
	return &dto.StartedMeetingDto{MeetingID: result.MeetingID, Leader: result.Leader, Start: result.Start, Finished: false}, nil
}

func NewMeetingsUseCase(service service.MeetingsService) *meetingsUseCase {
	return &meetingsUseCase{service: service}
}
