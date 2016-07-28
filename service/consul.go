package service

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/hashicorp/consul/api"
)

type ConsulCatalog interface {
	Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error)
	Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error)
}

type ConsulService struct {
	Catalog ConsulCatalog
}

type Consul struct {
	Service ConsulService
	Address string
	Port    int
}

func NewConsul(consulAddress string) *Consul {
	return &Consul{Service: ConsulService{Catalog: createConsulAPI(consulAddress)}}
}

func createConsulAPI(consulAddress string) ConsulCatalog {
	// setup consul config
	config := api.DefaultConfig()
	config.Address = strings.Split(consulAddress, ":")[0]

	// setup consul client
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	return client.Catalog()
}

func (c Consul) ListServices() {
	// set tab writer with field size
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 50, 10, 4, ' ', 0)

	// iterate through services
	services := c.Services()
	for _, k := range c.servicesKeys() {
		line := fmt.Sprintf("%s\t%s", color.YellowString(k), strings.Join(services[k], ", "))
		fmt.Fprintln(w, line)
	}

	// print out
	fmt.Fprintln(w)
	w.Flush()
}

func (c Consul) ShowService(serviceName string) {
	services := c.DiscoverService(serviceName)
	for _, service := range services {
		output, err := json.MarshalIndent(service, "", "\t")
		if err != nil {
			panic(err)
		}
		formatted := append(output, '\n')
		os.Stdout.Write(formatted)
	}
}

func (c Consul) OpenService(serviceName string) string {
	services := c.DiscoverService(serviceName)
	return serviceURL(services[0])
}

func (c Consul) DiscoverService(serviceName string) []*api.CatalogService {
	service, _, err := c.Service.Catalog.Service(serviceName, "", &api.QueryOptions{})
	if err != nil || len(service) == 0 {
		return []*api.CatalogService{}
	}
	return service
}

func (c Consul) Services() map[string][]string {
	catalog, _, err := c.Service.Catalog.Services(&api.QueryOptions{})
	if err != nil {
		panic(err)
	}
	return catalog
}

func serviceURL(service *api.CatalogService) string {
	if httpService(service.ServiceTags) != true {
		panic("This service is not tagged as HTTP service ")
	}

	return fmt.Sprintf("http://%s:%d", service.Address, service.ServicePort)
}

func httpService(tags []string) bool {
	if len(tags) > 0 {
		for _, tag := range tags {
			if tag == "http" {
				return true
			}
		}
	}
	return false
}

func (c Consul) servicesKeys() []string {
	var keys []string
	for k := range c.Services() {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
