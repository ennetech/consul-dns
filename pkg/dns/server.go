package dns

import (
	"github.com/ennetech/consul-dns/pkg/logger"
	"github.com/miekg/dns"
)

func serveInterface(net string, port string) {
	server := &dns.Server{Addr: ":" + port, Net: net, TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Failed to setup the " + net + " server: " + err.Error())
	}
}

func startServers(port string, handler func(dns.ResponseWriter, *dns.Msg)) error {
	dns.HandleFunc(".", handler)
	go serveInterface("tcp", port)
	logger.Info("...attached tcp:" + port)
	go serveInterface("udp", port)
	logger.Info("...attached udp:" + port)
	return nil
}
