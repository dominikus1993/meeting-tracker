package routes

import (
	"server/application/usecase"

	"github.com/gin-gonic/gin"
)

type MeetingRouter struct {
	meetingUseCase usecase.MeetingsUseCase
}

func (r *MeetingRouter) CreateNewMeeting(c *gin.Context) {
	dto, err := r.meetingUseCase.StartNew()
	if err != nil {
		c.Status(500)
		c.Error(err)
	} else {
		c.JSON(201, dto)
	}
}
