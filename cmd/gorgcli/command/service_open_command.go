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

func (soc *ServiceOpenCommand) Load() error {
	err := soc.loadConfig()
	if err != nil {
		return err
	}
	err = soc.loadConsul()
	if err != nil {
		return err
	}
	return nil
}

func (soc *ServiceOpenCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	soc.name = c.Args()[0]
	return nil
}

func (soc *ServiceOpenCommand) Run() error {
	url, err := soc.cat.ServiceURL(soc.name, "http")
	if err != nil {
		return err
	}
	return browser.OpenURL(url)
}

func (soc *ServiceOpenCommand) String() {
}
