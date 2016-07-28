package service

import (
	"fmt"

	"github.com/fatih/color"
)

type Environment struct {
	Name     string
	Domain   string
	Services Services
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

func NewEnvironment(name string, domain string) Environment {
	return Environment{Name: name, Domain: domain, Services: buildServicesMap(fmt.Sprintf("%s.%s", name, domain))}
}

func buildServicesMap(envDomain string) Services {
	s := Services{}
	s.ConsulUI = fmt.Sprintf("consul-ui.service.%s:8500", envDomain)
	s.Marathon = fmt.Sprintf("marathon.service.%s", envDomain)
	s.Mesos = fmt.Sprintf("mesos.service.%s", envDomain)
	s.Chronos = fmt.Sprintf("chronos.service.%s", envDomain)
	s.Kibana = fmt.Sprintf("kibana.service.%s", envDomain)
	s.WWW = fmt.Sprintf("www.public.%s", envDomain)
	return s
}
