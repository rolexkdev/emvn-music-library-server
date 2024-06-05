package main

import (
	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/server"
)

func main() {
	// Read configs base on environment variable
	// config := utils.Read("")
	config := &config.Config{
		Server: config.ServerConfig{
			AppVersion:  "v1",
			Port:        "8089",
			Host:        "127.0.0.1",
			Environment: config.DEVELOPMENT,
		},
	}
	server.InitServer(config)
}
