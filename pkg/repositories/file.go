package repositories

import (
	"github.com/ennetech/consul-dns/pkg/zone"
	"io/ioutil"
	"os"
)

type FileRepository struct {
	zone.Repository
}

func fileName(name string) string {
	return "zones/" + name
}

func (r FileRepository) Put(name, zone string) bool {
	err := ioutil.WriteFile(fileName(name), []byte(zone), os.ModePerm)
	return err == nil
}

func (r FileRepository) Get(name string) (zone.Zone, error) {
	b, err := ioutil.ReadFile(fileName(name))

	if err == nil {
		z := zone.NewZone(string(b), name)
		return *z, nil
	} else {
		return zone.Zone{}, err
	}
}
