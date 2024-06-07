package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Playlist struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreateAt   time.Time          `bson:"create_at"`
	UpdateAt   time.Time          `bson:"update_at"`
	DeleteAt   time.Time          `json:"-" bson:"delete_at,omitempty"`
	Title      string             `bson:"title"`
	AlbumCover string             `bson:"album_cover,omitempty"`
	TrackIDs   []string           `bson:"track_ids,omitempty"`
}

func (r *PlaylistRepository) Create(ctx context.Context, playlist *Playlist) (*Playlist, error) {
	playlist.CreateAt = time.Now()
	playlist.UpdateAt = playlist.CreateAt
	playlist.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, playlist)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (r *PlaylistRepository) FindByID(ctx context.Context, playlistID primitive.ObjectID) (*Playlist, error) {
	var playlist Playlist
	log.Printf("playlistID: %s", playlistID)
	filter := bson.M{
		"_id": playlistID,
		"delete_at": bson.M{
			"$exists": false,
		},
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&playlist)
	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

func (r *PlaylistRepository) Update(ctx context.Context, playlist *Playlist) error {
	filter := bson.M{"_id": playlist.ID}
	update := bson.M{"$set": bson.M{
		"title":       playlist.Title,
		"album_cover": playlist.AlbumCover,
		"update_at":   time.Now(),
		"track_ids":   playlist.TrackIDs,
	}}
	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}
	return nil
}

// Soft delete playlist record
func (r *PlaylistRepository) Delete(ctx context.Context, playlistID primitive.ObjectID) error {
	filter := bson.M{
		"_id": playlistID,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "delete_at",
					Value: time.Now(),
				},
			},
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *PlaylistRepository) FindMany(ctx context.Context) ([]*Playlist, error) {
	var playlists []*Playlist

	filter := bson.M{"delete_at": bson.M{"$eq": nil}}
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var playlist *Playlist
		if err := cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return playlists, nil
}

func (r *PlaylistRepository) Search(ctx context.Context, searchKey string) ([]*Playlist, error) {
	if searchKey == "" {
		return nil, errors.New("search key cannot be empty")
	}

	filter := bson.M{}
	filter["title"] = bson.M{"$regex": searchKey, "$options": "i"}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var playlists []*Playlist
	for cursor.Next(ctx) {
		var playlist Playlist
		if err := cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, &playlist)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return playlists, nil
}
