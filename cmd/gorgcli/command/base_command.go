package command

import (
	"encoding/json"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/pkg"
)

const (
	errMsgArg string = "missing args"
	errMsgFlg string = "missing flags"
	filename  string = "gorg.json"

	notifMsgAdd    string = "successfully added %s\n"
	notifMsgRm     string = "successfully removed %s\n"
	notifMsgSwitch string = "successfully switched to %s\n"
	notifInit      string = "successfully created config file: ~/gorg.json"
)

type Command interface {
	Load() error
	Validate(c *cli.Context) error
	Run() error
	String()
}

type baseCommand struct {
	Cfg        *pkg.Config
	csl        *pkg.Consul
	loadConfig bool
	loadConsul bool
}

func (bc *baseCommand) Load() error {
	if bc.loadConfig {
		cfg := pkg.Config{}
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		file, err := ioutil.ReadFile(home + string(os.PathSeparator) + filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(file, &cfg)
		if err != nil {
			return err
		}
		bc.Cfg = &cfg
	}
	if bc.loadConsul {
		consul, err := pkg.NewConsul(bc.Cfg.Default)
		if err != nil {
			return err
		}
		bc.csl = consul
	}
	return nil
}
