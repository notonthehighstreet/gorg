package command

import (
	"errors"

	"github.com/urfave/cli"

	"fmt"

	"github.com/notonthehighstreet/gorg/pkg"
)

type ConfigAddCommand struct {
	baseCommand
	name string
}

func (cac *ConfigAddCommand) Load() error {
	err := cac.loadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (cac *ConfigAddCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	cac.name = c.Args()[0]
	return nil
}

func (cac *ConfigAddCommand) Run() error {
	env := pkg.NewEnvironment(cac.name, cac.Cfg.Domain)
	err := cac.Cfg.AddEnvironment(env)
	if err != nil {
		return err
	}
	err = cac.Cfg.Update()
	if err != nil {
		return err
	}
	return nil
}

func (cac *ConfigAddCommand) String() {
	fmt.Printf(notifMsgAdd, cac.name)
}
