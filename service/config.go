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

func NewConfig(name string, domain string, environments []Environment) Config {
	c := Config{}
	c.Default = name
	c.Environments = environments
	return c
}

func NewConfigFile(name string, domain string) error {
	env := NewEnvironment(name, domain)
	config := NewConfig(name, domain, []Environment{env})
	err := config.writeToJSON()
	if err != nil {
		return err
	}

	return nil
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

func (c Config) AddConfigEnvironment(environmentName string, domain string) error {
	c.Environments = append(c.Environments, NewEnvironment(environmentName, domain))
	return c.UpdateConfig()
}

func (c Config) RemoveConfigEnvironment(environmentName string) error {
	for i, env := range c.Environments {
		if env.Name == environmentName {
			c.Environments = append(c.Environments[:i], c.Environments[i+1:]...)
			return c.UpdateConfig()
		}
	}

	return nil
}

func (c Config) UpdateConfig() error {
	err := c.writeToJSON()
	if err != nil {
		return err
	}

	return nil
}

func (c Config) ChangeUser(username string) error {
	c.User = username
	return c.UpdateConfig()
}

func (c Config) GetEnvironment(environmentName string) (Environment, error) {
	for _, env := range c.Environments {
		if env.Name == environmentName {
			return env, nil
		}
	}

	return Environment{}, errors.New("unknown environment name")
}

func (c Config) SwitchEnvironment(environmentName string) error {
	env, err := c.GetEnvironment(environmentName)
	if err != nil {
		return err
	}

	c.Default = env.Name
	return c.UpdateConfig()
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
