package main

import (
	"log"
	"server/api/routes"
	"server/application/service"
	"server/application/usecase"
	"server/common/env"
	"server/infrastructure/mongodb"
	"server/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *routes.MeetingRouter {
	r := repository.NewMeetingRepository(mongodb.NewClient(env.GetEnvOrDefault("MEETINGS_DB", "mongodb://root:rootpassword@127.0.0.1:27017")))
	s := service.NewMyMeetingService(r)
	useCase := usecase.NewMeetingsUseCase(s, r)
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
