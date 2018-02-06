package dns

import "github.com/miekg/dns"

func sendMessage(w dns.ResponseWriter, r *dns.Msg, recordsAnswer []dns.RR, rcode int) dns.Msg {
	a := new(dns.Msg)
	a.SetReply(r)
	a.Compress = true
	a.Authoritative = true
	a.Answer = recordsAnswer
	a.Rcode = rcode
	w.WriteMsg(a)
	return *a
}

func sendSuccess(w dns.ResponseWriter, r *dns.Msg, recordsAnswer []dns.RR) dns.Msg {
	return sendMessage(w, r, recordsAnswer, dns.RcodeSuccess)
}

func sendNxDomain(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	return sendMessage(w, r, []dns.RR{}, dns.RcodeNameError)
}
func sendNotImplemented(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	return sendMessage(w, r, []dns.RR{}, dns.RcodeNotImplemented)
}

func sendRefused(w dns.ResponseWriter, r *dns.Msg) dns.Msg {
	return sendMessage(w, r, []dns.RR{}, dns.RcodeRefused)
}
