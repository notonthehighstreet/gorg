package pkg

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/hashicorp/consul/api"
)

type ConsulCatalog interface {
	Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error)
	Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error)
}

type ConsulKV interface {
	List(string, *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error)
	Get(string, *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)
}

type consulClient struct {
	Catalog ConsulCatalog
	KV      ConsulKV
}

type Consul struct {
	Client consulClient
	opts   *api.QueryOptions
}

func NewConsul(address string) (*Consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = strings.Split(address, string(os.PathListSeparator))[0]
	client, err := api.NewClient(cfg)
	if err != nil {
		return &Consul{}, err
	}

	cc := consulClient{
		Catalog: client.Catalog(),
		KV:      client.KV(),
	}
	return &Consul{Client: cc, opts: &api.QueryOptions{}}, nil
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

func (c *Consul) ServiceAddress(name string) (string, error) {
	srv, _, err := c.Client.Catalog.Service(name, "", c.opts)
	if err != nil {
		return "", err
	}
	if reflect.DeepEqual(srv, []*api.CatalogService{}) {
		return "", errors.New("service not found")
	}
	return srv[0].Address, nil
}

func (c *Consul) KVList(service string) (map[string]string, error) {
	kvs, _, err := c.Client.KV.List(service, c.opts)
	if err != nil {
		return nil, err
	}
	list := map[string]string{}
	for _, kv := range kvs {
		list[kv.Key] = string(kv.Value)
	}
	return list, err
}

func (c *Consul) KVGet(name string) (map[string]string, error) {
	kv, _, err := c.Client.KV.Get(name, c.opts)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, errors.New("key not found")
	}
	res := make(map[string]string)
	res[kv.Key] = string(kv.Value)
	return res, err
}
