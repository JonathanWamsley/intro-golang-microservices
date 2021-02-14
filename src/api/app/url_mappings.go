package app

import (
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/controllers/polo"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}