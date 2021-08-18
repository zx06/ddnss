package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var ipv4Providers = []func() (string, error){
	func() (string, error) {

		const apiUrl = "https://httpbin.org/ip"
		var result = struct {
			Origin string `json:"origin"`
		}{}
		content, err := getHttpContent(apiUrl)
		if err != nil {
			return "", err
		}
		err = json.Unmarshal(*content, &result)
		if err != nil {
			return "", err
		}
		return result.Origin, nil
	},
}

func getHttpContent(url string) (*[]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// GetIPV4 获取本机公网ipv4地址
func GetIPV4() (string, error) {
	var (
		ip  string
		err error
	)
	// 遍历所有Provider
	for _, p := range ipv4Providers {
		ip, err = p()
		// 成功获取返回ip
		if err == nil {
			break
		}
	}
	log.Printf("Now wan ipv4 is: %s", ip)
	return ip, nil
}
