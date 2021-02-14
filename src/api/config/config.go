package config

import "os"

const(
	// export SECRET_GITHUB_ACCESS_TOKEN=your_access_token
	secretGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var(
	githubAccessToken = os.Getenv(secretGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}  