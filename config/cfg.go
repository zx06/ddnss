package config

import (
	"log"
	"github.com/zx06/ddnss/client"
	"github.com/zx06/ddnss/providers"
)

type CfgItem struct {
	Type   providers.ProviderType
	Domain string
	ApiKey string
}

type Config struct {
	Items map[string]*CfgItem
}

var Cfg Config = Config{
	Items: make(map[string]*CfgItem),
}

func Init() {
	// 从环境变量读取配置
	envCfg := GetEnvCfg()
	// 合并配置
	for k, v := range envCfg {
		Cfg.Items[k] = v
	}
}

func RegisterAll() []*client.DDNSClient {
	clients := make([]*client.DDNSClient, 0)
	for name, v := range Cfg.Items {
		switch v.Type {
		case providers.ProviderType_DYNU:
			p, err := providers.NewDynuProvider(v.Domain, v.ApiKey, "")
			if err != nil {
				log.Panicln("ddnss config error:", err)
			}
			clients = append(clients, client.NewDDNSClient(name, p))
		}
	}
	return clients
}
