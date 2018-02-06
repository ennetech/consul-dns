package config

type ConsulConfig struct {
	AuthToken   string
	HttpAddress string
	DnsAddress  string
}

type SystemConfig struct {
	DnsPort string
	HttpPort string
	TsigKey string
}

type ConsulDnsConfig struct {
	ConsulConfig ConsulConfig
	SystemConfig SystemConfig
}
