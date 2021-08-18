package client

import "log"

type DDNSProvider interface {
	Update() bool
}

type DDNSClient struct {
	Name     string
	Provider DDNSProvider
}

func NewDDNSClient(name string, provider DDNSProvider) *DDNSClient {
	return &DDNSClient{
		Name:     name,
		Provider: provider,
	}
}

func (client *DDNSClient) Update() bool {
	result := client.Provider.Update()
	if result {
		log.Printf("%s: Update Succeed\n", client.Name)
	}else {
		log.Printf("%s: Update Failed\n", client.Name)
	}
	return result
}
