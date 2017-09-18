package command

import (
	"errors"

	"fmt"

	"github.com/urfave/cli"
)

type SwitchEnvironmentCommand struct {
	baseCommand
	name string
}

func (sec *SwitchEnvironmentCommand) Load() error {
	err := sec.loadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (sec *SwitchEnvironmentCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	sec.name = c.Args()[0]
	return nil
}

func (sec *SwitchEnvironmentCommand) Run() error {
	err := sec.Cfg.SwitchEnvironment(sec.name)
	if err != nil {
		return err
	}
	err = sec.Cfg.Update()
	if err != nil {
		return err
	}
	return nil
}

func (sec *SwitchEnvironmentCommand) String() {
	fmt.Printf(notifMsgSwitch, sec.name)
}
