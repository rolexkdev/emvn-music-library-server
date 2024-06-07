package database

import (
	"context"
	"time"

	"github.com/rolexkdev/emvn-music-library-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup(c *config.Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Database.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(c.Database.Name)
	return db, nil
}
