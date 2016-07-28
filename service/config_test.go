package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configFilename = "dummy_config.json"
	fakeEnvName    = "dummy"
	fakeDomain     = "q.bogus.com"
	defaultEnvName = "integration"
	defaultDomain  = "qa.domain.com"
)

func AddFakeEnvironment() Config {
	c, _ := LoadConfig(configFilename)
	env := NewEnvironment(fakeEnvName, fakeDomain)
	c.AddConfigEnvironment(env)

	reload, _ := LoadConfig(configFilename)
	return reload
}

func TestLoadConfigReturnsNoError(t *testing.T) {
	c, e := LoadConfig(configFilename)
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

func TestNewConfigReturnsNoError(t *testing.T) {
	e := NewEnvironment("bogus", "qa.domain.com")
	c := NewConfig(e, "config.json")
	defer os.Remove("./config.json")
	assert.Nil(t, c)
}

func TestAddConfigEnvironmentReturnsPlusOneEnvironmentsSize(t *testing.T) {
	c, _ := LoadConfig(configFilename)
	env := NewEnvironment(fakeEnvName, fakeDomain)
	c.AddConfigEnvironment(env)

	reload, err := LoadConfig(configFilename)
	defer NewConfig(NewEnvironment(defaultEnvName, defaultDomain), configFilename)

	assert.Nil(t, err)
	assert.Equal(t, c.Path, configFilename)
	assert.Equal(t, len(reload.Environments), len(c.Environments)+1)
}

func TestRemoveConfigEnvironmentReturnsMinusOneEnvironmentsSize(t *testing.T) {
	AddFakeEnvironment()
	c, err := LoadConfig(configFilename)
	c.RemoveConfigEnvironment(fakeEnvName)

	reload, err := LoadConfig(configFilename)
	defer NewConfig(NewEnvironment(defaultEnvName, defaultDomain), configFilename)

	assert.Nil(t, err)
	assert.Equal(t, len(reload.Environments), len(c.Environments)-1)
}

func TestChangeUserReturnsUpdatedUsername(t *testing.T) {
	c, err := LoadConfig(configFilename)
	c.ChangeUser("violett")

	reload, err := LoadConfig(configFilename)
	defer NewConfig(NewEnvironment(defaultEnvName, defaultDomain), configFilename)

	assert.Nil(t, err)
	assert.NotEqual(t, reload.User, c.User)
}

func TestSwitchEnvironmentReturnsUpdatedDefaultEnvironment(t *testing.T) {
	c := AddFakeEnvironment()
	err := c.UseEnvironment(fakeEnvName)

	reload, err := LoadConfig(configFilename)
	defer NewConfig(NewEnvironment(defaultEnvName, defaultDomain), configFilename)

	assert.Nil(t, err)
	assert.NotEqual(t, reload.Default, c.Default)
}

func TestGetEnvironmentReturnsEnvironment(t *testing.T) {
	c, err := LoadConfig(configFilename)
	env, err := c.GetEnvironment(defaultEnvName)
	defer NewConfig(NewEnvironment(defaultEnvName, defaultDomain), configFilename)

	assert.Nil(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, defaultEnvName, env.Name)
}
