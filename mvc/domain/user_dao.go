package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDao struct{}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
