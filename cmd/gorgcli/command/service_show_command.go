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

func (ssc *ServiceShowCommand) Load() error {
	err := ssc.loadConfig()
	if err != nil {
		return err
	}
	err = ssc.loadConsul()
	if err != nil {
		return err
	}
	return nil
}

func (ssc *ServiceShowCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	ssc.name = c.Args()[0]
	return nil
}

func (ssc *ServiceShowCommand) Run() error {
	services, err := ssc.cat.Service(ssc.name)
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
