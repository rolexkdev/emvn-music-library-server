package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Track struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreateAt    time.Time          `bson:"create_at"`
	UpdateAt    time.Time          `bson:"update_at"`
	DeleteAt    time.Time          `json:"-" bson:"delete_at,omitempty"`
	Name        string             `bson:"name"`
	Title       string             `bson:"title"`
	ArtistID    string             `bson:"artist_id"`
	Album       string             `bson:"album"`
	Genre       string             `bson:"genre,omitempty"`
	ReleaseDate int64              `bson:"release_date"`
	Duration    int64              `bson:"duration"`
	FileURL     string             `bson:"file_url"`
}

func (r *TrackRepository) Create(ctx context.Context, track *Track) (*Track, error) {
	track.CreateAt = time.Now()
	track.UpdateAt = track.CreateAt
	track.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(ctx, track)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func (r *TrackRepository) FindByID(ctx context.Context, trackID primitive.ObjectID) (*Track, error) {
	var track Track
	log.Printf("trackID: %s", trackID)
	filter := bson.M{
		"_id": trackID,
		"delete_at": bson.M{
			"$exists": false,
		},
	}

	err := r.Collection.FindOne(ctx, filter).Decode(&track)
	if err != nil {
		return nil, err
	}

	return &track, nil
}

func (r *TrackRepository) Update(ctx context.Context, track *Track) error {
	filter := bson.M{"_id": track.ID}
	update := bson.M{"$set": bson.M{
		"title":     track.Title,
		"name":      track.Name,
		"album":     track.Album,
		"update_at": time.Now(),
		"artist_id": track.ArtistID,
		"genre":     track.Genre,
		"file_url":  track.FileURL,
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

// Soft delete track record
func (r *TrackRepository) Delete(ctx context.Context, trackID primitive.ObjectID) error {
	filter := bson.M{
		"_id": trackID,
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

func (r *TrackRepository) FindMany(ctx context.Context) ([]*Track, error) {
	var tracks []*Track

	filter := bson.M{"delete_at": bson.M{"$eq": nil}}
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var track *Track
		if err := cursor.Decode(&track); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *TrackRepository) Search(ctx context.Context, searchKey string) ([]*Track, error) {
	if searchKey == "" {
		return nil, errors.New("search key cannot be empty")
	}

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"title": bson.M{"$regex": searchKey, "$options": "i"}},
			bson.M{"name": bson.M{"$regex": searchKey, "$options": "i"}},
			bson.M{"album": bson.M{"$regex": searchKey, "$options": "i"}},
			bson.M{"genre": bson.M{"$regex": searchKey, "$options": "i"}},
		},
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tracks []*Track
	for cursor.Next(ctx) {
		var track Track
		if err := cursor.Decode(&track); err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}
