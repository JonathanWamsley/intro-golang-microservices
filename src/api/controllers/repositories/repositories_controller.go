package repositories

import (
	"net/http"

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/domain/repositories"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/services"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		c.JSON((apiErr.Status()), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}