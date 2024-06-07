package main

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/rolexkdev/emvn-music-library-server/common/utils"
	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/internal/models"
	"github.com/rolexkdev/emvn-music-library-server/server"
)

func main() {
	// Read configs base on environment variable
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	models.Setup(cfg)
	utils.Validator = validator.New()
	server.InitServer(cfg)
}
