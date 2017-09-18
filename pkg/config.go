package pkg

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Default      string
	Domain       string
	Username     string
	Filepath     string
	Environments []Environment
}

func NewConfig(env Environment, path, domain string) *Config {
	return &Config{
		Default:      env.Name,
		Filepath:     path,
		Domain:       domain,
		Environments: []Environment{env},
	}
}

func (c *Config) AddEnvironment(env Environment) error {
	for _, e := range c.Environments {
		if e == env {
			return errors.New("environment has already exists")
		}
	}
	c.Environments = append(c.Environments, env)
	return nil
}

func (c *Config) RemoveEnvironment(name string) error {
	if c.Default == name {
		return errors.New("can not remove default environment")
	}
	for i, env := range c.Environments {
		if env.Name == name {
			c.Environments = append(c.Environments[:i], c.Environments[i+1:]...)
			return nil
		}
	}
	return nil
}

func (c *Config) LoadEnvironment(name string) (Environment, error) {
	for _, env := range c.Environments {
		if env.Name == name {
			return env, nil
		}
	}
	return Environment{}, errors.New("unknown environment name")
}

func (c *Config) SwitchEnvironment(name string) error {
	env, err := c.LoadEnvironment(name)
	if err != nil {
		return err
	}
	c.Default = env.Name
	return nil
}

func (c *Config) ChangeUser(name string) {
	c.Username = name
}

func (c *Config) Update() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	file, err := os.Create(c.Filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	w.Write(data)
	w.Flush()
	return nil
}
