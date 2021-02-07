package services

import (
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/domain"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
