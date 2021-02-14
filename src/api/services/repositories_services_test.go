package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/clients/restclient"
	"github.com/JonathanWamsley/courses/federico/intro-golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidName(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name: "",
		Description: "test description",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid repository name", err.Message())
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
}

func TestCreateRepoIErrorFromGithub(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication",
				"documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name: "test name",
		Description: "test description",
	}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Requires authentication", err.Message())
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "test name", "owner": {"login": "login owner"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "test name"}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)	
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "test name", result.Name)
	assert.EqualValues(t, "login owner", result.Owner)
}