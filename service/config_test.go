package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigReturnsNoError(t *testing.T) {
	c, e := LoadConfig("dummy_config.json")
	assert.Nil(t, e)
	assert.NotNil(t, c.Default)
	assert.NotNil(t, c.User)
	assert.NotNil(t, c.Environments)
}

func TestLoadConfigWithNonExistingFileReturnsError(t *testing.T) {
	c, e := LoadConfig("non-existing.json")
	assert.NotNil(t, e)
	assert.Equal(t, c, Config{})
}

// func TestAddConfigEnvironmentReturnsPlusOneEnvironmentsSize(t *testing.T) {
// 	c, _ := LoadConfig("dummy_config.json")
// 	original := c.Environments
// 	env := NewEnvironment("dummy", "q.bogus.com")
// 	err := c.AddConfigEnvironment(env)
//
// 	assert.Nil(t, err)
// 	assert.NotEqual(t, len(c.Environments), len(original))
// }

func TestNewConfigReturnsNoError(t *testing.T) {
	e := NewEnvironment("bogus", "qa.domain.com")
	c := NewConfig(e)
	defer os.Remove("./config.json")
	assert.Nil(t, c)
}
