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

func NewSwitchUserCmd() SwitchUserCommand {
	return SwitchUserCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: false,
		},
	}
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
	return suc.Cfg.Update()
}

func (suc *SwitchUserCommand) String() {
	fmt.Printf(notifMsgSwitch, suc.name)
}
