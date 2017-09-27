package command

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type KVListCommand struct {
	baseCommand
	srvName string
	kvs     map[string]string
}

func NewKVListCmd() KVListCommand {
	return KVListCommand{
		baseCommand: baseCommand{
			loadConfig: true,
			loadConsul: true,
		},
	}
}

func (cmd *KVListCommand) Validate(c *cli.Context) error {
	if c.NArg() < 1 {
		return errors.New(errMsgArg)
	}
	cmd.srvName = c.Args()[0]
	return nil
}

func (cmd *KVListCommand) Run() error {
	kvs, err := cmd.csl.KVList(cmd.srvName)
	if err != nil {
		return err
	}
	cmd.kvs = kvs
	return nil
}

func (cmd *KVListCommand) String() {
	keys := []string{}
	for k := range cmd.kvs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 50, 10, 4, ' ', 0)
	for _, key := range keys {
		line := fmt.Sprintf("%s\t%s", color.YellowString(key), cmd.kvs[key])
		fmt.Fprintln(w, line)
	}

	fmt.Fprintln(w)
	w.Flush()
}
