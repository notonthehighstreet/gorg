package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	Default      string
	Environments []Environment
	Path         string
	User         string
}

func LoadConfig(dataFile string) (Config, error) {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}

func NewConfig(env Environment, path string) error {
	c := Config{}
	c.Default = env.Name
	c.Environments = []Environment{env}
	c.Path = path
	return c.writeToJSON()
}

func (c Config) AddConfigEnvironment(env Environment) error {
	for _, environment := range c.Environments {
		if environment == env {
			return errors.New("environment has already exists")
		}
	}

	c.Environments = append(c.Environments, env)
	return c.updateConfig()
}

func (c Config) RemoveConfigEnvironment(environmentName string) error {
	if c.Default == environmentName {
		return errors.New("can not remove default environment")
	}

	for i, env := range c.Environments {
		if env.Name == environmentName {
			c.Environments = append(c.Environments[:i], c.Environments[i+1:]...)
			return c.updateConfig()
		}
	}

	return nil
}

func (c Config) ChangeUser(username string) error {
	c.User = username
	return c.updateConfig()
}

func (c Config) UseEnvironment(environmentName string) error {
	env, err := c.GetEnvironment(environmentName)
	if err != nil {
		return err
	}

	err = c.switchEnvironment(env)
	if err != nil {
		return err
	}

	return nil
}

func (c Config) switchEnvironment(env Environment) error {
	c.Default = env.Name
	return c.updateConfig()
}

func (c Config) GetEnvironment(environmentName string) (Environment, error) {
	for _, env := range c.Environments {
		if env.Name == environmentName {
			return env, nil
		}
	}

	return Environment{}, errors.New("unknown environment name")
}

func (c Config) updateConfig() error {
	err := c.writeToJSON()
	if err != nil {
		return err
	}

	return nil
}

func (c Config) writeToJSON() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	file, err := os.Create(c.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	w.Write(data)
	w.Flush()

	return nil
}
