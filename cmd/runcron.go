package cmd

import (
	"time"

	"github.com/zx06/ddnss/config"
)

func RunCron() {
	config.Init()
	clients := config.RegisterAll()
	for _, client := range clients {
		client.Update()
	}
	ticker := time.NewTicker(time.Hour * 1)
	for range ticker.C {
		for _, client := range clients {
			client.Update()
		}
	}
}
