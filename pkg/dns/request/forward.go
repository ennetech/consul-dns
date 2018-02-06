package request

import "github.com/miekg/dns"

func Forward(server, zone string, t uint16) ([]dns.RR, error) {
	m := new(dns.Msg)
	m.SetQuestion(zone, t)
	c := new(dns.Client)
	mes, _, err := c.Exchange(m, server)
	return mes.Answer, err
}
