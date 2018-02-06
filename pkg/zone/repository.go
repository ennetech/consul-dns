package zone

type Repository interface {
	Put(name, zone string) bool
	Get(zone string) (Zone, error)
}
