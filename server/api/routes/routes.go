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
	request := struct {
		Leaders []string `json:"leaders"`
	}{}
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(400)
	}
	dto, err := router.meetingUseCase.StartNew(request.Leaders, c)
	if err != nil {
		c.Status(500)
		_ = c.Error(err)
	} else {
		c.JSON(201, dto)
	}
}

func (router *MeetingRouter) finishMeeting(c *gin.Context) {
	id := c.Param("id")
	dto, err := router.meetingUseCase.Finish(id, c)
	if err != nil {
		c.JSON(500, err.Error())
	} else {
		c.JSON(201, dto)
	}
}

func (router *MeetingRouter) getById(c *gin.Context) {
	id := c.Param("id")
	dto, err := router.meetingUseCase.GetMeeting(id, c)
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
	r.GET("/meetings/:id", router.getById)
	r.GET("/meetings", router.getAllMeetings)
	r.POST("/meetings", router.createNewMeeting)
	r.PUT("/meetings/:id", router.finishMeeting)
}
