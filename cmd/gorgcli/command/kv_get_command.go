package command

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type KVGetCommand struct {
	baseCommand
	key string
	kv  map[string]string
}

func NewKVGetCmd() KVGetCommand {
	return KVGetCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: true,
		},
	}
}

func (cmd *KVGetCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	cmd.key = c.Args()[0]
	return nil
}

func (cmd *KVGetCommand) Run() error {
	kv, err := cmd.csl.KVGet(cmd.key)
	if err != nil {
		return err
	}
	cmd.kv = kv
	return nil
}

func (cmd *KVGetCommand) String() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 50, 10, 4, ' ', 0)
	line := fmt.Sprintf("%s\t%s", color.YellowString(cmd.key), cmd.kv[cmd.key])
	fmt.Fprintln(w, line)
	w.Flush()
}
