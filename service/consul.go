package service

// import (
// 	"errors"
//
// 	"github.com/hashicorp/consul/api"
// )
//
// type ConsulCatalog interface {
// 	Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error)
// }
//
// type ConsulService struct {
// 	catalog ConsulCatalog
// }
//
// // Consul defines a consul service
// type Consul struct {
// 	Address string
// 	Port    int
// }
//
// func NewConsul(consulAddress string) *Consul {
// 	return &Consul{catalog: createConsulAPI(consulAddress)}
// }
//
// func createConsulAPI(consulAddress string) ConsulCatalog {
// 	config := api.DefaultConfig()
// 	config.Address = consulAddress
//
// 	client, err := api.NewClient(config)
// 	if err != nil {
// 		return nil
// 	}
//
// 	return client.Catalog()
// }
//
// // DiscoverService returns the service address and port from a consul
// func (c *Consul) DiscoverService(serviceKey string) (Service, error) {
// 	services, _, err := c.catalog.Service(serviceKey, "", nil)
// 	if err != nil {
// 		return Service{}, err
// 	}
//
// 	if len(services) == 0 {
// 		return Service{}, errors.New("consul: no services for key")
// 	}
//
// 	svc := Service{Address: services[0].Address, Port: services[0].ServicePort}
// 	return svc, nil
// }
