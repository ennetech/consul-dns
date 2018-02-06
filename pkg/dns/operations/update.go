package operations

import (
	"github.com/ennetech/consul-dns/pkg/zone"
	"github.com/miekg/dns"
)

func HandleUpdate(ns dns.RR, z *zone.Zone) error {
	if ns.Header().Class == dns.ClassANY || ns.Header().Class == dns.ClassNONE { // Deletion
		z.DeleteRR(ns)
	} else { // Addition
		z.AddRR(ns)
	}
	return nil
}
