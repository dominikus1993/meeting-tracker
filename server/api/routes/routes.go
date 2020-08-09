package routes

import (
	"server/application/usecase"

	"github.com/gin-gonic/gin"
)

type MeetingRouter struct {
	meetingUseCase usecase.MeetingsUseCase
}

func CreateNewMeeting(c *gin.Context) {

}
