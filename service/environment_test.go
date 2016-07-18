package service

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
		"\n    ConsulUI: http://consul-ui.service.integration.qa.domain.com:8500" +
			"\n    Marathon: http://marathon.service.integration.qa.domain.com" +
			"\n    Mesos:    http://mesos.service.integration.qa.domain.com" +
			"\n    Chronos:  http://chronos.service.integration.qa.domain.com" +
			"\n    Kibana:   http://kibana.service.integration.qa.domain.com" +
			"\n    WWW:      http://www.public.integration.qa.domain.com\n\t"

	assert.NotNil(t, s.String())
	assert.Equal(t, expected, s.String())
}
