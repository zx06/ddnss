package client

import (
	"log"
	"time"
)

type DDNSProvider interface {
	Update() bool
}

type DDNSClient struct {
	Name     string
	Interval int
	Provider DDNSProvider
}

func NewDDNSClient(name string, interval int, provider DDNSProvider) *DDNSClient {
	return &DDNSClient{
		Name:     name,
		Interval: interval,
		Provider: provider,
	}
}

func (client *DDNSClient) Update() bool {
	result := client.Provider.Update()
	if result {
		log.Printf("%s: Update Succeed\n", client.Name)
	} else {
		log.Printf("%s: Update Failed\n", client.Name)
	}
	return result
}

func (client *DDNSClient) Schedule(runImmediately bool) {
	go func() {
		if runImmediately {
			client.Update()
		}
		d := time.Hour * time.Duration(client.Interval)
		ticker := time.NewTicker(d)
		log.Printf("%s: Scheduled to run every %d hours\n", client.Name, client.Interval)
		// next run at
		log.Printf("%s: Scheduled next run at %s\n", client.Name, time.Now().Add(d).Format("15:04:05"))
		for range ticker.C {
			client.Update()
			log.Printf("%s: Scheduled next run at %s\n", client.Name, time.Now().Add(d).Format("15:04:05"))
		}
	}()
}
