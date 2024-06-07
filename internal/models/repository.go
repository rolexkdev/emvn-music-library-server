package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppRepository struct {
	Track    TrackRepositoryInterface
	Playlist PlaylistRepositoryInterface
}

type TrackRepository struct {
	Collection *mongo.Collection
}
type PlaylistRepository struct {
	Collection *mongo.Collection
}

type TrackRepositoryInterface interface {
	Create(ctx context.Context, track *Track) (*Track, error)
	FindByID(ctx context.Context, trackID primitive.ObjectID) (*Track, error)
	Update(ctx context.Context, track *Track) error
	Delete(ctx context.Context, trackID primitive.ObjectID) error
	FindMany(ctx context.Context) ([]*Track, error)
}

type PlaylistRepositoryInterface interface {
	Create(ctx context.Context, playlist *Playlist) (*Playlist, error)
	FindByID(ctx context.Context, playlistID primitive.ObjectID) (*Playlist, error)
	Update(ctx context.Context, playlist *Playlist) error
	Delete(ctx context.Context, playlistID primitive.ObjectID) error
	FindMany(ctx context.Context) ([]*Playlist, error)
}
