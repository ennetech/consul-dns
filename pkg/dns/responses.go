package dns

import "github.com/miekg/dns"

func sendSuccess(w dns.ResponseWriter, r *dns.Msg, recordsAnswer []dns.RR) string {
	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	a.Answer = recordsAnswer
	a.Rcode = dns.RcodeSuccess
	w.WriteMsg(a)
	return a.String()
}

func sendNxDomain(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	a.Rcode = dns.RcodeNameError
	w.WriteMsg(a)
	return *a
}
func sendNotImplemented(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	a.Rcode = dns.RcodeNotImplemented
	w.WriteMsg(a)
	return *a
}

func sendRefused(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	a.Rcode = dns.RcodeRefused
	w.WriteMsg(a)
	return *a
}
