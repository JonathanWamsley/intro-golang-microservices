package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/services"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/mvc/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		// handle the err and return to client
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(apiErr.Message))
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
