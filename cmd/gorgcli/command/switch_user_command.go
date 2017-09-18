package command

import (
	"errors"

	"fmt"

	"github.com/urfave/cli"
)

type SwitchUserCommand struct {
	baseCommand
	name string
}

func (suc *SwitchUserCommand) Load() error {
	err := suc.loadConfig()
	if err != nil {
		return err
	}
	return nil
}

func (suc *SwitchUserCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	suc.name = c.Args()[0]
	return nil
}

func (suc *SwitchUserCommand) Run() error {
	suc.Cfg.ChangeUser(suc.name)
	err := suc.Cfg.Update()
	if err != nil {
		return err
	}
	return nil
}

func (suc *SwitchUserCommand) String() {
	fmt.Printf(notifMsgSwitch, suc.name)
}
