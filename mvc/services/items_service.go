package services

import (
	"net/http"

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/domain"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "something went wrong",
	}
}
