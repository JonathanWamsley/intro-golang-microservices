package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TextConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", secretGithubAccessToken)
}

func TestGetAccessToken(t *testing.T) {
	assert.EqualValues(t, "", GetGithubAccessToken())
}