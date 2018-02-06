package operations

import (
	"github.com/miekg/dns"
	"strings"
)

func HandleMasquerade(questionName string, questionType uint16, tld string, server string) ([]dns.RR, error) {

	// We have to masq. current zone and then send to consul
	partial := strings.TrimSuffix(questionName, tld)
	partial = partial + "consul."

	m := new(dns.Msg)
	m.SetQuestion(partial, questionType)
	c := new(dns.Client)
	mes, _, _ := c.Exchange(m, server)
	// TODO: error handling
	for _, x := range mes.Answer {
		x.Header().Name = questionName
	}

	return mes.Answer, nil
}
