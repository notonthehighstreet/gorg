package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name   = "integration"
	domain = "qa.domain.com"
)

func TestNewEnvironmentReturnsBuiltServices(t *testing.T) {
	e := NewEnvironment(name, domain)

	assert.Equal(t, name, e.Name)
	assert.Equal(t, domain, e.Domain)
	assert.NotNil(t, e.Services)
}

func TestStringReturnsFormattedString(t *testing.T) {
	e := NewEnvironment(name, domain)
	s := e.Services
	expected :=
		"\n    ConsulUI: consul-ui.service.integration.qa.domain.com:8500" +
			"\n    Marathon: marathon.service.integration.qa.domain.com" +
			"\n    Mesos:    mesos.service.integration.qa.domain.com" +
			"\n    Chronos:  chronos.service.integration.qa.domain.com" +
			"\n    Kibana:   kibana.service.integration.qa.domain.com" +
			"\n    WWW:      www.public.integration.qa.domain.com\n\t"

	assert.NotNil(t, s.String())
	assert.Equal(t, expected, s.String())
}
