package app

import (

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}