package config

import (
	"os"
	"strconv"
	"strings"
)

const (
	ENV_PREFIX   = "DDNSS_"
	TYPE_KEY     = "TYPE"
	DOMAIN_KEY   = "DOMAIN"
	APIKEY_KEY   = "APIKEY"
	INTERVAL_KEY = "INTERVAL"
)

type EnvType [2]string

func getAllenv() [][2]string {
	var result [][2]string
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, ENV_PREFIX) {
			var tmp = [2]string{}
			copy(tmp[:], strings.Split(v, "=")[:2])
			result = append(result, [2]string{strings.Replace(tmp[0], ENV_PREFIX, "", 1), tmp[1]})
		}

	}
	return result
}

func GetEnvCfg() map[string]*CfgItem {
	var result = make(map[string]*CfgItem)
	for _, v := range getAllenv() {

		envData := strings.Split(v[0], "_")
		if len(envData) != 2 {
			continue
		}
		// 该配置为空时才初始化
		if _, ok := result[envData[0]]; !ok {
			result[envData[0]] = &CfgItem{
				Interval: 1,
			}
		}
		// 根据类型获取配置项
		switch envData[1] {
		case TYPE_KEY:
			result[envData[0]].Type = v[1]
		case DOMAIN_KEY:
			result[envData[0]].Domain = v[1]
		case APIKEY_KEY:
			result[envData[0]].ApiKey = v[1]
		case INTERVAL_KEY:
			interval, err := strconv.Atoi(v[1])
			if err == nil {
				result[envData[0]].Interval = interval
			}

		}

	}
	return result
}
