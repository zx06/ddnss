package cmd

import "github.com/zx06/ddnss/config"

func RunOnce() {
	config.Init()
	clients := config.RegisterAll()
	for _, client := range clients {
		client.Update()
	}
}
