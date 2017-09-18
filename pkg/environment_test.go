package pkg

import (
	"testing"
)

var (
	name   = "integration"
	domain = "qa.domain.com"
)

func TestNewEnvironmentReturnsBuiltServices(t *testing.T) {
	e := NewEnvironment(name, domain)

	if name != e.Name {
		t.Error("expected %s, got %s", name, e.Name)
	}
	if domain != e.Domain {
		t.Error("expected %s, got %s", domain, e.Domain)
	}
}

func TestStringReturnsFormattedString(t *testing.T) {
	e := NewEnvironment(name, domain)
	s := e.Services
	expected :=
		"\n    ConsulUI: \x1b[33mconsul-ui.service.integration.qa.domain.com:8500\x1b[0m" +
			"\n    Marathon: \x1b[33mmarathon.service.integration.qa.domain.com\x1b[0m" +
			"\n    Mesos:    \x1b[33mmesos.service.integration.qa.domain.com\x1b[0m" +
			"\n    Chronos:  \x1b[33mchronos.service.integration.qa.domain.com\x1b[0m" +
			"\n    Kibana:   \x1b[33mkibana.service.integration.qa.domain.com\x1b[0m" +
			"\n    WWW:      \x1b[33mwww.public.integration.qa.domain.com\x1b[0m\n\t"

	if s.String() != expected {
		t.Errorf("expected:\n%s\n got:\n%s\n", expected, s.String())
	}
}
