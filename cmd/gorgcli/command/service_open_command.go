package command

import (
	"errors"

	"github.com/pkg/browser"
	"github.com/urfave/cli"
)

type ServiceOpenCommand struct {
	baseCommand
	name string
}

func NewServiceOpenCmd() ServiceOpenCommand {
	return ServiceOpenCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: true,
		},
	}
}

func (soc *ServiceOpenCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	soc.name = c.Args()[0]
	return nil
}

func (soc *ServiceOpenCommand) Run() error {
	url, err := soc.csl.ServiceURL(soc.name, "http")
	if err != nil {
		return err
	}
	return browser.OpenURL(url)
}

func (soc *ServiceOpenCommand) String() {
}
