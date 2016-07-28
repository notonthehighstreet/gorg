package service

import (
	"testing"

	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

func mockedServices() map[string][]string {
	services := make(map[string][]string, 2)
	services["cdn"] = []string{"public", "http"}
	services["consul"] = []string{}
	services["chronos"] = []string{"http", "management"}

	return services
}

func mockedService(key string) []*api.CatalogService {
	if key == "cdn" {
		cdn := api.CatalogService{
			Node:                     "10.12.36.2",
			Address:                  "10.12.36.2",
			ServiceID:                "10.12.36.2:mesos-1e918be3-5179-40bb-a397-5e4d8fc837d6-S3.a2ab426e-5bb5-4f17-864e-6309dda3e69e:80",
			ServiceName:              "cdn",
			ServiceAddress:           "10.12.36.2",
			ServiceTags:              []string{"public", "http"},
			ServicePort:              31853,
			ServiceEnableTagOverride: false,
		}

		return []*api.CatalogService{&cdn}
	}
	return []*api.CatalogService{}
}

type DoubleConsulCatalog struct {
	services      map[string][]string
	serviceDetail []*api.CatalogService
	err           error
}

func (d DoubleConsulCatalog) Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error) {
	return d.services, nil, d.err
}

func (d DoubleConsulCatalog) Service(serviceName, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	return d.serviceDetail, nil, d.err
}

var serviceKey = "elasticsearch-logstash"

func TestServicesReturnsArrayOfService(t *testing.T) {
	d := DoubleConsulCatalog{services: mockedServices()}
	s := ConsulService{Catalog: d}
	c := Consul{Service: s}

	assert.Equal(t, len(mockedServices()), len(c.Services()))
}

func TestDiscoverServiceReturnsServiceDetails(t *testing.T) {
	key := "cdn"
	d := DoubleConsulCatalog{serviceDetail: mockedService(key)}
	s := ConsulService{Catalog: d}
	c := Consul{Service: s}

	assert.Equal(t, mockedService(key), c.DiscoverService(key))
}
