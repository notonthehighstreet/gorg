package command

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type ServiceListCommand struct {
	baseCommand
	services map[string][]string
}

func (slc *ServiceListCommand) Load() error {
	err := slc.loadConfig()
	if err != nil {
		return err
	}
	err = slc.loadConsul()
	if err != nil {
		return err
	}
	return nil
}

func (slc *ServiceListCommand) Validate(c *cli.Context) error {
	return nil
}

func (slc *ServiceListCommand) Run() error {
	services, err := slc.cat.Services()
	if err != nil {
		return err
	}
	slc.services = services
	return nil
}

func (slc *ServiceListCommand) String() {
	keys := []string{}
	for k := range slc.services {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 50, 10, 4, ' ', 0)
	for _, key := range keys {
		line := fmt.Sprintf("%s\t%s", color.YellowString(key), strings.Join(slc.services[key], ", "))
		fmt.Fprintln(w, line)
	}

	fmt.Fprintln(w)
	w.Flush()
}
