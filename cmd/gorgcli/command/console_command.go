package command

import (
	"errors"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

type ConsoleCommand struct {
	baseCommand
	name string
}

func NewConsoleCmd() ConsoleCommand {
	return ConsoleCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: true,
		},
	}
}

func (cmd *ConsoleCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	cmd.name = c.Args()[0]
	return nil
}

func (cmd *ConsoleCommand) Run() error {
	address, err := cmd.csl.ServiceAddress(cmd.name)
	if err != nil {
		return err
	}

	c := exec.Command("ssh", address)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	return c.Run()
}

func (cmd *ConsoleCommand) String() {
}
