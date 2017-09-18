package command

import (
	"encoding/json"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"strings"

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

type baseCommand struct {
	Cfg *pkg.Config
	cat *pkg.Consul
}

func (bc *baseCommand) loadConfig() error {
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
	return nil
}

func (bc *baseCommand) loadConsul() error {
	env, err := bc.Cfg.LoadEnvironment(bc.Cfg.Default)
	if err != nil {
		return err
	}
	address := strings.Split(env.Services.ConsulUI, string(os.PathListSeparator))[0]
	consul, err := pkg.NewConsul(address)
	if err != nil {
		return err
	}
	bc.cat = consul
	return nil
}
