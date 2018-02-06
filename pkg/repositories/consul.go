package repositories

import (
	"github.com/ennetech/consul-dns/pkg/zone"
	"github.com/hashicorp/consul/api"
	"github.com/ennetech/consul-dns/pkg/config"
	"errors"
)

type ConsulRepository struct {
	zone.Repository
}

func (r ConsulRepository) Put(name, zone string) bool {
	data, _, _ := consulClient.KV().Get("dns/"+name, &api.QueryOptions{})
	if data != nil {
		data.Value = []byte(zone)
		_, err := consulClient.KV().Put(data, &api.WriteOptions{})
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

var consulClient *api.Client

func SetConsulClient(config config.ConsulConfig){
	consulClient, _ = api.NewClient(&api.Config{
		Address: config.HttpAddress,
		Token:   config.AuthToken,
	})
}

func (r ConsulRepository) Get(name string) (zone.Zone, error) {
	data, _, _ := consulClient.KV().Get("dns/"+name, &api.QueryOptions{})
	if data != nil {
		z := zone.NewZone(string(data.Value), name)
		return *z, nil
	} else {
		return zone.Zone{}, errors.New("empty data on consul")
	}
}
