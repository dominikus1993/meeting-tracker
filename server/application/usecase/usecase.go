package usecase

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/application/dto"
	"server/application/service"
	"server/domain/model"
	"server/domain/repository"
	"time"
)

type MeetingsUseCase interface {
	StartNew(leaders []string, context context.Context) (*dto.StartedMeetingDto, error)
	GetAll(context context.Context) ([]*dto.StartedMeetingDto, error)
	Finish(id string, c context.Context) (*dto.FinishedMeetingDto, error)
	GetMeeting(id string, c *gin.Context) (*dto.MeetingDto, error)
}

type meetingsUseCase struct {
	service service.MeetingsService
	repo    repository.MeetingsRepository
}

func NewMeetingsUseCase(service service.MeetingsService, repo repository.MeetingsRepository) *meetingsUseCase {
	return &meetingsUseCase{service: service, repo: repo}
}

func (u *meetingsUseCase) GetMeeting(id string, c *gin.Context) (*dto.MeetingDto, error) {
	result, err := u.service.GetMeeting(id, c)
	if err != nil {
		return nil, err
	}
	var finish *time.Time
	if !result.Finish.IsZero() {
		finish = &result.Finish
	}
	return &dto.MeetingDto{MeetingID: result.MeetingID, Leaders: result.Leaders, Start: result.Start, End: finish, Finished: true}, nil
}

func (u *meetingsUseCase) Finish(id string, c context.Context) (*dto.FinishedMeetingDto, error) {
	result, err := u.service.Finish(id, c)
	if err != nil {
		return nil, err
	}
	return &dto.FinishedMeetingDto{MeetingID: result.MeetingID, Leaders: result.Leaders, Start: result.Start, End: result.Finish, Finished: true}, nil
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
		res[i] = &dto.StartedMeetingDto{MeetingID: el.MeetingID, Leaders: el.Leaders, Start: el.Start, Finished: el.Finished}
	}

	return res, nil
}

func (u *meetingsUseCase) StartNew(leaders []string, context context.Context) (*dto.StartedMeetingDto, error) {
	id, err := u.repo.GetId()
	if err != nil {
		return nil, err
	}
	result, err := u.service.Start(model.NewMeeting(id, leaders), context)
	if err != nil {
		return nil, err
	}
	return &dto.StartedMeetingDto{MeetingID: result.MeetingID, Leaders: result.Leaders, Start: result.Start, Finished: false}, nil
}
