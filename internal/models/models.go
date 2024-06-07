package models

import (
	"errors"
	"log"

	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB         *mongo.Database
	Repository *AppRepository
)

var (
	ErrNotFound = errors.New("record not found")
)

func Setup(c *config.Config) {
	var err error

	DB, err = database.Setup(c)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	Repository = &AppRepository{
		Track:    &TrackRepository{DB.Collection("track")},
		Playlist: &PlaylistRepository{DB.Collection("playlist")},
	}
}
