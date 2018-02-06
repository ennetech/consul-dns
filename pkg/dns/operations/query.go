package operations

import (
	"github.com/ennetech/consul-dns/pkg/zone"
	"github.com/miekg/dns"
)

// Only one query must be handled by this file, the convenient solution is to

func HandleQuery(q dns.Question, z zone.Zone) []dns.RR {
	return z.QueryRR(q)
}
