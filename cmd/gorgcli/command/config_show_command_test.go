package command

import (
	"testing"

	"github.com/notonthehighstreet/gorg/pkg"
)

func TestConfigShowCommand_Run(t *testing.T) {
	env := pkg.NewEnvironment("bogus", "domain.com")
	scc := ConfigShowCommand{}
	scc.baseCommand = baseCommand{
		Cfg: &pkg.Config{
			Default:      env.Name,
			Environments: []pkg.Environment{env},
		},
	}

	err := scc.Run()
	if err != nil {
		t.Error("unexpected error:", err)
	}
}

func TestConfigShowCommand_String(t *testing.T) {
	scc := ConfigShowCommand{
		defaultEnv: "bogus",
		services: pkg.Services{
			ConsulUI: "bogus.service.consul-ui.domain",
			Marathon: "bogus.service.marathon.domain",
			Mesos:    "bogus.service.mesos.domain",
			Chronos:  "bogus.service.chronos.domain",
			Kibana:   "bogus.service.kibana.domain",
			WWW:      "bogus.service.www.domain",
		},
		envs: []pkg.Environment{
			{Name: "jughead"},
			{Name: "bogus"},
		},
	}
	scc.String()
}
