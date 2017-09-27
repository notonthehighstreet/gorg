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

func NewSwitchEnvironmentCmd() SwitchEnvironmentCommand {
	return SwitchEnvironmentCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: false,
		},
	}
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
	return sec.Cfg.Update()
}

func (sec *SwitchEnvironmentCommand) String() {
	fmt.Printf(notifMsgSwitch, sec.name)
}
