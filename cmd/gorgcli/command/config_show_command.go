package command

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/pkg"
)

type ConfigShowCommand struct {
	baseCommand
	defaultEnv string
	services   pkg.Services
	envs       []pkg.Environment
}

func (csc *ConfigShowCommand) Load() error {
	err := csc.loadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (csc *ConfigShowCommand) Validate(c *cli.Context) error {
	return nil
}

func (csc *ConfigShowCommand) Run() error {
	def, err := csc.Cfg.LoadEnvironment(csc.Cfg.Default)
	if err != nil {
		return err
	}

	csc.defaultEnv = def.Name
	csc.services = def.Services
	csc.envs = csc.Cfg.Environments
	return nil
}

func (csc *ConfigShowCommand) String() {
	fmt.Printf("Using %s environment by default: \n", csc.defaultEnv)
	fmt.Printf("%s", csc.services)
	fmt.Print("\nEnvironments available: \n")
	for _, env := range csc.envs {
		fmt.Printf("  - %s \n", env.Name)
	}
}
