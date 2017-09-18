package command

import (
	"errors"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"

	"fmt"

	"github.com/notonthehighstreet/gorg/pkg"
)

type InitCommand struct {
	baseCommand
	domain string
	ename  string
}

func (ic *InitCommand) Load() error {
	return nil
}

func (ic *InitCommand) Validate(c *cli.Context) error {
	ic.domain = c.String("domain")
	if ic.domain == "" {
		return errors.New(errMsgFlg)
	}
	ic.ename = c.String("environment-name")
	if ic.ename == "" {
		return errors.New(errMsgFlg)
	}
	return nil
}

func (ic *InitCommand) Run() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	env := pkg.NewEnvironment(ic.ename, ic.domain)
	cfg := pkg.NewConfig(env, home+string(os.PathSeparator)+filename, ic.domain)
	if err := cfg.Update(); err != nil {
		return err
	}

	ic.Cfg = cfg
	return nil
}

func (ic *InitCommand) String() {
	fmt.Println(notifInit)
}
