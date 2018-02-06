package dns

import (
	"errors"
	"github.com/ennetech/consul-dns/pkg/zone"
	"strings"
)

func generateChain(txt string) []string {
	var chain []string
	chain = append(chain, txt)

	sub := strings.Split(strings.Trim(txt, "."), ".")

	build := ""
	for i, s := range sub {
		if i < len(sub)-2 {
			build = build + s + "."
			partial := strings.TrimPrefix(txt, build)
			if len(partial) > 0 {
				chain = append(chain, partial)
			}
		}
	}
	return chain
}

func checkZone(z *zone.Zone, s string) error {
	var chain = generateChain(s)
	for _, ring := range chain {

		// We have to load from the repository
		// TODO: don't like this comparison to know if the zone was already set
		if z.Origin() == "" {
			zo, e := repository.Get(ring)
			if e == nil {
				*z = zo
				return nil
			}
		} else {
			// Check if zone is managed
			if z.Origin() == ring {
				return nil
			}
		}

	}

	return errors.New("zone is not managed")
}
