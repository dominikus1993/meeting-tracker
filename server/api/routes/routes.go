package routes

import (
	"server/application/usecase"

	"github.com/gin-gonic/gin"
)

type MeetingRouter struct {
	meetingUseCase usecase.MeetingsUseCase
}

func NewMeetingRouter(meetingUseCase usecase.MeetingsUseCase) *MeetingRouter {
	return &MeetingRouter{meetingUseCase: meetingUseCase}
}

func (router *MeetingRouter) createNewMeeting(c *gin.Context) {
	leader := c.Param("leader")
	dto, err := router.meetingUseCase.StartNew(leader, c)
	if err != nil {
		c.Status(500)
		_ = c.Error(err)
	} else {
		c.JSON(201, dto)
	}
}

func (router *MeetingRouter) getAllMeetings(c *gin.Context) {
	dtos, err := router.meetingUseCase.GetAll(c)
	if err != nil {
		c.Status(500)
		_ = c.Error(err)
	} else {
		c.JSON(200, dtos)
	}
}

func (router *MeetingRouter) SetUpRoutes(r *gin.Engine) {
	r.GET("/meetings", router.getAllMeetings)
	r.POST("/meetings/:leader", router.createNewMeeting)
}
