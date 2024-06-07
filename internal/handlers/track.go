package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/common/app"
	"github.com/rolexkdev/emvn-music-library-server/common/e"
	"github.com/rolexkdev/emvn-music-library-server/common/utils"
	"github.com/rolexkdev/emvn-music-library-server/dto"
	"github.com/rolexkdev/emvn-music-library-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTrack godoc
//
//	@Summary		Create a track
//	@Description	create a track
//	@Tags			track
//	@Accept			json
//	@Produce		json
//
//	@Param			input		    body		dto.CreateTrackRequest	true	"Create Track Request input"
//
//	@Success		201				{object}	app.Response
//	@Failure		400				{object}	app.Response
//	@Failure		403				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/tracks [post]
func CreateTrack(c *gin.Context) {
	appG := app.Gin{C: c}

	var request dto.CreateTrackRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		appG.Response400(e.INVALID_PARAMS, "Parse JSON body failed: "+err.Error())
		return
	}
	fmt.Println("Request input: ", request)

	// Validate request input
	if err := utils.Validator.Struct(request); err != nil {
		appG.Response400(e.INVALID_PARAMS, "Validate JSON body failed: "+err.Error())
		return
	}

	track := &models.Track{
		Name:        request.Name,
		Title:       request.Title,
		ArtistID:    request.ArtistID,
		Album:       request.Album,
		Genre:       request.Genre,
		ReleaseDate: request.ReleaseDate,
		Duration:    request.Duration,
		FileURL:     request.FileURL,
	}

	trackCreated, err := models.Repository.Track.Create(context.Background(), track)
	if err != nil {
		appG.Response500(e.ERROR, "create track failed with error: "+err.Error())
		return
	}

	appG.Response201(trackCreated)
}

// GetTrack godoc
//
//	@Summary		Get a track
//	@Description	Get a track
//	@Tags			track
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	true	"track id"
//
//	@Success		200				{object}	app.Response
//	@Failure		404				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/tracks/{id} [get]
func GetTrack(c *gin.Context) {
	appG := app.Gin{C: c}
	trackID := c.Param("id")

	// convert track_id string to objectID
	objID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	track, err := models.Repository.Track.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Track not exist")
			return
		}
		appG.Response500(e.ERROR, "Get track by id failed with err: "+err.Error())
		return
	}

	appG.Response200(track)
}

// GetTracks godoc
//
//	@Summary		Get list tracks
//	@Description	Get list tracks
//	@Tags			track
//	@Accept			json
//	@Produce		json
//
//
//	@Success		200				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/tracks/ [get]
func GetTracks(c *gin.Context) {
	appG := app.Gin{C: c}

	tracks, err := models.Repository.Track.FindMany(context.Background())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		appG.Response500(e.ERROR, "Get track by id failed with err: "+err.Error())
		return
	}

	appG.Response200(tracks)
}

// DeleteTrack godoc
//
//	@Summary		Delete a track by id
//	@Description	Delete a track by id
//	@Tags			track
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	true	"track id"
//
//	@Success		200				{object}	app.Response
//	@Success		404				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/tracks/{id} [delete]
func DeleteTrack(c *gin.Context) {
	appG := app.Gin{C: c}
	trackID := c.Param("id")

	// convert track_id string to objectID
	objID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	_, err = models.Repository.Track.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Track not exist")
			return
		}
		appG.Response500(e.ERROR, "Get track by id failed with err: "+err.Error())
		return
	}

	err = models.Repository.Track.Delete(context.Background(), objID)
	if err != nil {
		appG.Response500(e.ERROR, "Delete a track by id failed with err: "+err.Error())
		return
	}

	appG.Response200("success")
}

// UpdateTrack godoc
//
//	@Summary		Update a track by id
//	@Description	Update a track by id
//	@Tags			track
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	true	"track id"
//	@Param			input		    body		dto.UpdateTrackRequest	true	"Update Track Request input"
//
//	@Success		200				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/tracks/{id} [PUT]
func UpdateTrack(c *gin.Context) {
	appG := app.Gin{C: c}
	trackID := c.Param("id")

	var request dto.UpdateTrackRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		appG.Response400(e.INVALID_PARAMS, "Parse JSON body failed: "+err.Error())
		return
	}
	fmt.Println("Request input: ", request)

	// Validate request input
	if err := utils.Validator.Struct(request); err != nil {
		appG.Response400(e.INVALID_PARAMS, "Validate JSON body failed: "+err.Error())
		return
	}

	// convert track_id string to objectID
	objID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	track, err := models.Repository.Track.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response400(e.INVALID_PARAMS, "Track not exist")
			return
		}
		appG.Response500(e.ERROR, "Get track by id failed with err: "+err.Error())
		return
	}

	if request.Name != nil {
		track.Name = *request.Name
	}

	if request.Album != nil {
		track.Album = *request.Album
	}

	if request.ArtistID != nil {
		track.ArtistID = *request.ArtistID
	}

	if request.Title != nil {
		track.Title = *request.Title
	}

	if request.FileURL != nil {
		track.FileURL = *request.FileURL
	}

	if request.Genre != nil {
		track.Genre = *request.Genre
	}

	err = models.Repository.Track.Update(context.Background(), track)
	if err != nil {
		appG.Response500(e.ERROR, "Delete a track by id failed with err: "+err.Error())
		return
	}

	trackUpdated, err := models.Repository.Track.FindByID(context.Background(), objID)
	if err != nil {
		appG.Response500(e.ERROR, "Get track updated failed with err: "+err.Error())
		return
	}

	appG.Response200(trackUpdated)
}
