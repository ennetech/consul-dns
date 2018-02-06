package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/ennetech/consul-dns/pkg/common"
	"github.com/ennetech/consul-dns/pkg/logger"
)

func LoadConfiguration(configPath string) (conf ConsulDnsConfig, err error) {
	// Set defaults
	conf.ConsulConfig.AuthToken = "anonymous"
	conf.ConsulConfig.HttpAddress = "http://127.0.0.1:8500"
	conf.ConsulConfig.DnsAddress = "127.0.0.1:8600"
	conf.SystemConfig.DnsPort = "53"
	conf.SystemConfig.HttpPort = "4367"

	// Load from file
	if configPath != "" {
		if !path.IsAbs(configPath) {
			wd, _ := os.Getwd()
			configPath = path.Clean(wd + string(os.PathSeparator) + configPath)
		}
		logger.Info("config specified: " + configPath)

		err := json.Unmarshal(common.ReadFile(configPath), &conf)
		if err != nil {
			logger.Error(err.Error())
		}
	} else {
		logger.Warn("no configuration specified")
	}

	// Detect env
	envOrDefault("CONSUL_HTTP_TOKEN", &conf.ConsulConfig.HttpAddress)
	envOrDefault("CONSUL_HTTP_ADDR", &conf.ConsulConfig.HttpAddress)
	envOrDefault("CONSUL_DNS_ADDR", &conf.ConsulConfig.DnsAddress)

	envOrDefault("CONSULDNS_DNS_PORT", &conf.SystemConfig.DnsPort)
	envOrDefault("CONSULDNS_HTTP_PORT", &conf.SystemConfig.HttpPort)
	envOrDefault("CONSULDNS_TSIG_KEY", &conf.SystemConfig.TsigKey)
	return conf, nil
}

func envOrDefault(key string, target *string) {
	v := os.Getenv(key)
	if v != "" {
		*target = v
		logger.Info("Detected env variable " + key)
	}
}

func SaveConfiguration(conf ConsulDnsConfig, configPath string) (err error) {
	j, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	common.WriteFile(configPath, []byte(j))
	return nil
}
