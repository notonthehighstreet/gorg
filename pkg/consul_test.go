package pkg

import (
	"errors"
	"testing"

	"github.com/hashicorp/consul/api"
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

func TestServices_WithNoError(t *testing.T) {
	d := DoubleConsulCatalog{services: mockedServices()}
	s := consulClient{catalog: d}
	c := Consul{API: s}
	exp, err := c.Services()

	if err != nil {
		t.Error("unexpected error:", err)
	}
	if len(mockedServices()) != len(exp) {
		t.Errorf("expecting %d services listed, got: %d", len(mockedServices()), len(exp))
	}
}

func TestServices_WithError(t *testing.T) {
	d := DoubleConsulCatalog{err: errors.New("error message")}
	s := consulClient{catalog: d}
	c := Consul{API: s}
	exp, err := c.Services()

	if err == nil {
		t.Error("expected error got nil")
	}
	if len(exp) != 0 {
		t.Errorf("unexpected content: %v", exp)
	}
}

func TestDiscover_WitNoError(t *testing.T) {
	key := "cdn"
	d := DoubleConsulCatalog{serviceDetail: mockedService(key)}
	s := consulClient{catalog: d}
	c := Consul{API: s}
	exp, err := c.Discover(key)

	if err != nil {
		t.Error("unexpected error:", err)
	}
	if len(mockedService(key)) != len(exp) {
		t.Errorf("expecting %d, got: %d", len(mockedService(key)), len(exp))
	}
}

func TestDiscover_WithError(t *testing.T) {
	d := DoubleConsulCatalog{err: errors.New("error message")}
	s := consulClient{catalog: d}
	c := Consul{API: s}
	_, err := c.Discover("cdn")

	if err == nil {
		t.Error("expecting an error got nil")
	}
}
