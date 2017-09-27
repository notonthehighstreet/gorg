package command

import (
	"errors"

	"fmt"

	"github.com/urfave/cli"
)

type ConfigRemoveCommand struct {
	baseCommand
	name string
}

func NewConfigRemoveCmd() ConfigRemoveCommand {
	return ConfigRemoveCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: false,
		},
	}
}

func (crc *ConfigRemoveCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	crc.name = c.Args()[0]
	return nil
}

func (crc *ConfigRemoveCommand) Run() error {
	err := crc.Cfg.RemoveEnvironment(crc.name)
	if err != nil {
		return err
	}
	return crc.Cfg.Update()
}

func (crc *ConfigRemoveCommand) String() {
	fmt.Printf(notifMsgRm, crc.name)
}
