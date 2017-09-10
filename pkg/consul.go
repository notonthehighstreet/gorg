package pkg

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/api"
)

const managementTag = "management"

type ConsulCatalog interface {
	Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error)
	Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error)
}

type consulClient struct {
	catalog ConsulCatalog
}

type Consul struct {
	API         consulClient
	Address     string
	Port        int
	Managements []string
	opts        *api.QueryOptions
}

func NewConsul(address string) (*Consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = strings.Split(address, ":")[0]

	c := &Consul{}
	client, err := api.NewClient(cfg)
	if err != nil {
		return c, err
	}

	c.API = consulClient{catalog: client.Catalog()}
	c.opts = &api.QueryOptions{}
	return c, nil
}

func (c *Consul) Services() (map[string][]string, error) {
	ctl, _, err := c.API.catalog.Services(c.opts)
	if err != nil {
		return make(map[string][]string), err
	}
	return ctl, nil
}

func (c *Consul) Discover(name string) ([]*api.CatalogService, error) {
	srv, _, err := c.API.catalog.Service(name, "", c.opts)
	if err != nil {
		return []*api.CatalogService{}, err
	}
	return srv, nil
}

func (c *Consul) ServiceURL(name, tag string) (string, error) {
	srv, _, err := c.API.catalog.Service(name, tag, c.opts)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s:%d", srv[0].Address, srv[0].ServicePort)
	return url, nil
}

func (c *Consul) getManagementServices() error {
	ctl, err := c.Services()
	if err != nil {
		return err
	}

	for key, tags := range ctl {
		for _, tag := range tags {
			if tag == managementTag {
				c.Managements = append(c.Managements, key)
			}
		}
	}
	return nil
}
