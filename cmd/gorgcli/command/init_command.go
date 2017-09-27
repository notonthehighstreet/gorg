package command

import (
	"errors"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/pkg"
)

type InitCommand struct {
	baseCommand
	domain string
	ename  string
}

func NewInitCmd() InitCommand {
	return InitCommand{
		baseCommand: baseCommand{
			loadConfig: false,
			loadConsul: false,
		},
	}
}

func (ic *InitCommand) Validate(c *cli.Context) error {
	ic.domain = c.String("domain")
	if ic.domain == "" {
		return errors.New(errMsgFlg)
	}
	return nil
}

func (ic *InitCommand) Run() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	cfg := pkg.NewConfig(home+string(os.PathSeparator)+filename, ic.domain)
	if err := cfg.Update(); err != nil {
		return err
	}
	ic.Cfg = cfg
	return nil
}

func (ic *InitCommand) String() {
	fmt.Println(notifInit)
}
