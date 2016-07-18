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

func NewConfig(env Environment) error {
	c := Config{}
	c.Default = env.Name
	c.Environments = []Environment{env}
	return c.writeToJSON()
}

func (c Config) AddConfigEnvironment(env Environment) error {
	c.Environments = append(c.Environments, env)
	return c.updateConfig()
}

func (c Config) RemoveConfigEnvironment(environmentName string) error {
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

func (c Config) SwitchEnvironment(env Environment) error {
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

	file, err := os.Create("./config.json")
	if err != nil {
		return err
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	w.Write(data)
	w.Flush()

	return nil
}
