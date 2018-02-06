package main

import (
	"flag"
	"github.com/ennetech/consul-dns/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/ennetech/consul-dns/pkg/config"
	"github.com/ennetech/consul-dns/pkg/dns"
	"github.com/ennetech/consul-dns/pkg/http"
	"github.com/ennetech/consul-dns/pkg/repositories"
)

var (
	version    = "0.0.0"
	configPath = flag.String("c", "", "Path to configuration file")
)

func main() {
	flag.Parse()
	logger.Info("consul-dns v" + version)
	logger.Info("starting...")

	// Load the configuration
	consulDnsConfig, err := config.LoadConfiguration(*configPath)
	if err != nil {
		logger.Warn("Failed to load the configuration: " + err.Error())
	}

	// Initialize the storage
	repositories.SetConsulClient(consulDnsConfig.ConsulConfig)

	// Start the DNS interface
	dns.Init(consulDnsConfig, repositories.ConsulRepository{})
	logger.Info("Started dns interface on port: " + consulDnsConfig.SystemConfig.DnsPort)

	// Start the REST interface
	http.Init(consulDnsConfig)
	logger.Info("Started rest interface on port: " + consulDnsConfig.SystemConfig.HttpPort)

	logger.Info("...running...")
	// Gracefull shutdown routine
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	logger.Info("...shutting down (" + s.String() + ")...")
	config.SaveConfiguration(consulDnsConfig, *configPath)
	logger.Info("...stopped")
}
