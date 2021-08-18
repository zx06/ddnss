package config

import (
	"github.com/zx06/ddnss/client"
	"github.com/zx06/ddnss/providers"
	"log"
)

type CfgItem struct {
	// ddns服务商类型
	Type providers.ProviderType
	// 需要绑定的域名
	Domain string
	// ddns服务商的api key
	ApiKey string
	// 更新间隔(小时,默认1小时)
	Interval int
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
				log.Println("ddnss config error:", err)
				continue
			}
			log.Printf("ddnss config: [%s] type=%s domain=%s interval=%d", name, v.Type, v.Domain, v.Interval)
			clients = append(clients, client.NewDDNSClient(name, v.Interval, p))
		}
	}
	return clients
}
