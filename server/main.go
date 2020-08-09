package main

import (
	"log"
	"server/api/routes"
	"server/application/service"
	"server/application/usecase"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *routes.MeetingRouter {
	s := service.NewMyMeetingService(nil)
	useCase := usecase.NewMeetingsUseCase(s)
	return routes.NewMeetingRouter(useCase)
}

func main() {
	router := setUpRouter()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.SetUpRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
