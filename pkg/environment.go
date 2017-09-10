package pkg

import (
	"fmt"

	"github.com/fatih/color"
)

type Environment struct {
	Name     string
	Domain   string
	Services Services
}

func NewEnvironment(name string, domain string) Environment {
	s := Services{}
	s.ConsulUI = fmt.Sprintf("consul-ui.service.%s.%s:8500", name, domain)
	s.Marathon = fmt.Sprintf("marathon.service.%s.%s", name, domain)
	s.Mesos = fmt.Sprintf("mesos.service.%s.%s", name, domain)
	s.Chronos = fmt.Sprintf("chronos.service.%s.%s", name, domain)
	s.Kibana = fmt.Sprintf("kibana.service.%s.%s", name, domain)
	s.WWW = fmt.Sprintf("www.public.%s.%s", name, domain)

	return Environment{
		Name:     name,
		Domain:   domain,
		Services: s,
	}
}

type Services struct {
	ConsulUI string
	Marathon string
	Mesos    string
	Chronos  string
	Kibana   string
	WWW      string
}

func (s Services) String() string {
	ret := `
    ConsulUI: %s
    Marathon: %s
    Mesos:    %s
    Chronos:  %s
    Kibana:   %s
    WWW:      %s
	`

	return fmt.Sprintf(ret,
		color.YellowString(s.ConsulUI),
		color.YellowString(s.Marathon),
		color.YellowString(s.Mesos),
		color.YellowString(s.Chronos),
		color.YellowString(s.Kibana),
		color.YellowString(s.WWW))
}
