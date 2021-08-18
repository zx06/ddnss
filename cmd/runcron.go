package cmd

import (
	"github.com/zx06/ddnss/config"
)

func RunCron() {
	config.Init()
	clients := config.RegisterAll()
	for _, client := range clients {
		client.Schedule(true)
	}
	select {}
}
