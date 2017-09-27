package command

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/pkg"
)

type ConfigAddCommand struct {
	baseCommand
	name string
}

func NewConfigAddCmd() ConfigAddCommand {
	return ConfigAddCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: false,
		},
	}
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
	return cac.Cfg.Update()
}

func (cac *ConfigAddCommand) String() {
	fmt.Printf(notifMsgAdd, cac.name)
}
