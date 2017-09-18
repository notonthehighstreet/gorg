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

func (crc *ConfigRemoveCommand) Load() error {
	err := crc.loadConfig()
	if err != nil {
		return err
	}
	return nil
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
	err = crc.Cfg.Update()
	if err != nil {
		return err
	}
	return nil
}

func (crc *ConfigRemoveCommand) String() {
	fmt.Printf(notifMsgRm, crc.name)
}
