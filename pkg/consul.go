package pkg

import (
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

type ConsulCatalog interface {
	Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error)
	Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error)
}

type ConsulKV interface {
}

type consulClient struct {
	Catalog ConsulCatalog
	KV      ConsulKV
}

type Consul struct {
	Client  consulClient
	opts    *api.QueryOptions
	address string
}

func NewConsul(address string) (*Consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = address
	client, err := api.NewClient(cfg)
	if err != nil {
		return &Consul{}, err
	}

	cc := consulClient{
		Catalog: client.Catalog(),
		KV:      client.KV(),
	}
	return &Consul{Client: cc, address: address, opts: &api.QueryOptions{}}, nil
}

func (c *Consul) Service(name string) ([]*api.CatalogService, error) {
	srv, _, err := c.Client.Catalog.Service(name, "", c.opts)
	if err != nil {
		return []*api.CatalogService{}, err
	}
	return srv, err
}

func (c *Consul) Services() (map[string][]string, error) {
	srv, _, err := c.Client.Catalog.Services(c.opts)
	if err != nil {
		return nil, err
	}
	return srv, err
}

func (c *Consul) ServiceURL(name, tag string) (string, error) {
	srv, _, err := c.Client.Catalog.Service(name, tag, c.opts)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("http://%s%s%d", srv[0].Address, string(os.PathListSeparator), srv[0].ServicePort)
	return url, nil
}
