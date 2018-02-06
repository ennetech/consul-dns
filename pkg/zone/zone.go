package zone

import (
	"github.com/ennetech/consul-dns/pkg/common"
	"github.com/ennetech/consul-dns/pkg/logger"
	"github.com/miekg/dns"
	"sort"
	"strings"
)

type Zone struct {
	text    string
	origin  string
	records []dns.RR
}

func NewZone(text, origin string) *Zone {
	z := new(Zone)
	z.text = text
	z.origin = origin
	z.updateRecords()
	return z
}

func (zone *Zone) String() string {
	return zone.text
}
func (zone *Zone) Origin() string {
	return zone.origin
}

func (zone *Zone) updateRecords() {
	for x := range dns.ParseZone(strings.NewReader(zone.text), dns.Fqdn(zone.origin), "") {
		if x.Error != nil {
			// Ignore parsing error
			// panic(x.Error)
		} else {
			zone.records = append(zone.records, x.RR)
		}
	}
}

func (zone *Zone) updateText() {
	build := common.NewChainText()
	for _, rr := range zone.records {
		build.Append(rr.String())
	}
	sp := strings.Split(build.String(), "\n")
	sort.Strings(sp)
	zone.text = strings.Join(reverseStringsSlice(sp), "\n")
}

func (zone *Zone) QueryRR(question dns.Question) []dns.RR {
	var answerRecords []dns.RR

	for _, rr := range zone.records {
		namematch := rr.Header().Name == question.Name
		typematch := rr.Header().Rrtype == question.Qtype
		// Perfect match
		if namematch && typematch {
			answerRecords = append(answerRecords, rr)
			// Wildcard match
		} else if !namematch && strings.HasPrefix(rr.Header().Name, "*") && (len(answerRecords) == 0) {
			// Remove first level on question
			s := strings.Split(question.Name, ".")
			copy(s, s[1:])
			s = s[:len(s)-1]
			builded := "*." + strings.Join(s, ".")

			if (builded == rr.Header().Name) && typematch {
				// We have to correct the name (so we don't expose we have a wildcard)
				rr.Header().Name = question.Name
				answerRecords = append(answerRecords, rr)
			}
			// Cname
		} else if question.Qtype == dns.TypeA && rr.Header().Rrtype == dns.TypeCNAME {
			if cname, ok := rr.(*dns.CNAME); ok {
				m := new(dns.Msg)
				m.SetQuestion(cname.Target, dns.TypeA)
				c := new(dns.Client)
				mes, _, _ := c.Exchange(m, "8.8.8.8:53")
				answerRecords = append(answerRecords, mes.Answer...)
			}
		}
	}
	return answerRecords
}

func (zone *Zone) AddRR(input dns.RR) {
	newRecord := true
	for _, rr := range zone.records {
		if input.String() == rr.String() {
			newRecord = false
		}
	}
	if newRecord {
		zone.records = append(zone.records, input)
	}
	// Update the text version of the zone
	zone.updateText()

	logger.Debug(zone.String(), "Updated zone after AddRR")
}

func (zone *Zone) DeleteRR(input dns.RR) {
	// We maybe need to delete more than one item
	var deleteIndices []int

	// Scan all the zone records
	for i, rr := range zone.records {
		// Uniform TTL
		// TODO: i think this can be deleted
		//input.Header().Ttl = 300
		//rr.Header().Ttl = 300

		// Additional check, goes false when a specific case happen
		additional := true

		// A record value must match
		t1, o1 := input.(*dns.A)
		t2, o2 := rr.(*dns.A)
		if o1 && o2 && input.Header().Class != dns.ClassANY {
			additional = t1.A.String() == t2.A.String()
		}

		// Type + Name + Additional check matching
		if input.Header().Rrtype == rr.Header().Rrtype && input.Header().Name == rr.Header().Name && additional {
			deleteIndices = append(deleteIndices, i)
		}
	}

	// Delete in reverse order
	deleteIndices = reverseIntegersSlice(deleteIndices)
	for _, d := range deleteIndices {
		zone.records = append(zone.records[:d], zone.records[d+1:]...)
	}

	// Update the text version of the zone
	zone.updateText()

	logger.Debug(zone.String(), "Updated zone after RemoveRR")
}

// Helper functions

func reverseIntegersSlice(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func reverseStringsSlice(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
