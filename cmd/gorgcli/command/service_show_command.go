package command

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

type ServiceShowCommand struct {
	baseCommand
	services []string
	name     string
}

func NewServiceShowCmd() ServiceShowCommand {
	return ServiceShowCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: true,
		},
	}
}

func (ssc *ServiceShowCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	ssc.name = c.Args()[0]
	return nil
}

func (ssc *ServiceShowCommand) Run() error {
	services, err := ssc.csl.Service(ssc.name)
	if err != nil {
		return err
	}
	for _, srv := range services {
		output, err := json.MarshalIndent(srv, "", "\t")
		if err != nil {
			return err
		}
		ssc.services = append(ssc.services, string(output)+string('\n'))
	}
	return nil
}

func (ssc *ServiceShowCommand) String() {
	for _, srv := range ssc.services {
		fmt.Println(srv)
	}
}
