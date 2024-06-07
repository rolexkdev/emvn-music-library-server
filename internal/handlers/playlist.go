package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/common/app"
	"github.com/rolexkdev/emvn-music-library-server/common/e"
	"github.com/rolexkdev/emvn-music-library-server/common/utils"
	"github.com/rolexkdev/emvn-music-library-server/dto"
	"github.com/rolexkdev/emvn-music-library-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreatePlaylist godoc
//
//	@Summary		Create a playlist
//	@Description	create a playlist
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//	@Param			input		    body		dto.CreatePlaylistRequest	true	"Create Playlist Request input"
//
//	@Success		201				{object}	app.Response
//	@Failure		400				{object}	app.Response
//	@Failure		403				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists [post]
func CreatePlaylist(c *gin.Context) {
	appG := app.Gin{C: c}

	var request dto.CreatePlaylistRequest
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

	playlist := &models.Playlist{
		Title:      request.Title,
		AlbumCover: request.AlbumCover,
		TrackIDs:   request.TrackIDs,
	}

	playlistCreated, err := models.Repository.Playlist.Create(context.Background(), playlist)
	if err != nil {
		appG.Response500(e.ERROR, "create playlist failed with error: "+err.Error())
		return
	}

	appG.Response201(playlistCreated)
}

// GetPlaylist godoc
//
//	@Summary		Get a playlist
//	@Description	Get a playlist
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	true	"playlist id"
//
//	@Success		200				{object}	app.Response
//	@Failure		404				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/{id} [get]
func GetPlaylist(c *gin.Context) {
	appG := app.Gin{C: c}
	playlistID := c.Param("id")

	// convert playlist_id string to objectID
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	playlist, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Playlist not exist")
			return
		}
		appG.Response500(e.ERROR, "Get playlist by id failed with err: "+err.Error())
		return
	}

	appG.Response200(playlist)
}

// GetPlaylists godoc
//
//	@Summary		Get list playlists
//	@Description	Get list playlists
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//
//	@Success		200				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/ [get]
func GetPlaylists(c *gin.Context) {
	appG := app.Gin{C: c}

	playlists, err := models.Repository.Playlist.FindMany(context.Background())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		appG.Response500(e.ERROR, "Get playlist by id failed with err: "+err.Error())
		return
	}

	appG.Response200(playlists)
}

// DeletePlaylist godoc
//
//	@Summary		Delete a playlist by id
//	@Description	Delete a playlist by id
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	true	"playlist id"
//
//	@Success		200				{object}	app.Response
//	@Success		400				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/{id} [delete]
func DeletePlaylist(c *gin.Context) {
	appG := app.Gin{C: c}
	playlistID := c.Param("id")

	// convert playlist_id string to objectID
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	_, err = models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Playlist not exist")
			return
		}
		appG.Response500(e.ERROR, "Get playlist by id failed with err: "+err.Error())
		return
	}

	err = models.Repository.Playlist.Delete(context.Background(), objID)
	if err != nil {
		appG.Response500(e.ERROR, "Delete a playlist by id failed with err: "+err.Error())
		return
	}

	appG.Response200("success")
}

// UpdatePlaylist godoc
//
//	@Summary		Update a playlist by id
//	@Description	Update a playlist by id
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	                     true	"playlist id"
//	@Param			input		body		dto.UpdatePlaylistRequest	true	"Update Playlist Request input"
//
//	@Success		200				{object}	app.Response
//	@Failure		404				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/{id} [PUT]
func UpdatePlaylist(c *gin.Context) {
	appG := app.Gin{C: c}
	playlistID := c.Param("id")

	var request dto.UpdatePlaylistRequest
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

	// convert playlist_id string to objectID
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	playlist, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Playlist not exist")
			return
		}
		appG.Response500(e.ERROR, "Get playlist by id failed with err: "+err.Error())
		return
	}

	if request.AlbumCover != nil {
		playlist.AlbumCover = *request.AlbumCover
	}

	if request.Title != nil {
		playlist.Title = *request.Title
	}

	err = models.Repository.Playlist.Update(context.Background(), playlist)
	if err != nil {
		appG.Response500(e.ERROR, "Delete a playlist by id failed with err: "+err.Error())
		return
	}

	playlistUpdated, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		appG.Response500(e.ERROR, "Get playlist updated failed with err: "+err.Error())
		return
	}

	appG.Response200(playlistUpdated)
}

// UpdatePlaylistTrack godoc
//
//	@Summary		Add or remove a track for playlist
//	@Description	Add or remove a track for playlist
//	@Tags			playlist
//	@Accept			json
//	@Produce		json
//
//	@Param			id		    path		string	                        true	"playlist id"
//	@Param			input		body		dto.UpdatePlaylistTrackRequest	true	"Update Playlist Track Request input"
//
//	@Success		200				{object}	app.Response
//	@Failure		404				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/{id}/tracks [POST]
func UpdatePlaylistTrack(c *gin.Context) {
	appG := app.Gin{C: c}
	playlistID := c.Param("id")

	var request dto.UpdatePlaylistTrackRequest
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

	// convert playlist_id string to objectID
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	playlist, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Playlist not exist")
			return
		}
		appG.Response500(e.ERROR, "Get playlist by id failed with err: "+err.Error())
		return
	}

	if request.IsDelete != nil {
		if *request.IsDelete {
			utils.RemoveElementFromArray(&playlist.TrackIDs, request.TrackID)
		} else {
			playlist.TrackIDs = append(playlist.TrackIDs, request.TrackID)
		}
	}

	err = models.Repository.Playlist.Update(context.Background(), playlist)
	if err != nil {
		appG.Response500(e.ERROR, "Delete a playlist by id failed with err: "+err.Error())
		return
	}

	playlistUpdated, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		appG.Response500(e.ERROR, "Get playlist updated failed with err: "+err.Error())
		return
	}

	appG.Response200(playlistUpdated)
}

// GenerateM3UPlaylist godoc
//
//	@Summary		Generate M3U playlist
//	@Description	Generate M3U playlist
//	@Tags			playlist
//	@Produce		audio/x-mpegurl
//
//	@Param			id	path	string	true	"Playlist ID"
//
//	@Success		200	            {string}	string	"M3U playlist"
//	@Failure		400				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/playlists/{id}/m3u [get]
func GenerateM3UPlaylist(c *gin.Context) {
	appG := app.Gin{C: c}
	playlistID := c.Param("id")

	// convert playlist_id string to objectID
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
		return
	}
	fmt.Println("objID: ", objID)

	playlist, err := models.Repository.Playlist.FindByID(context.Background(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			appG.Response404(e.NOTFOUND, "Playlist not exist")
			return
		}
		appG.Response500(e.ERROR, "Get playlist updated failed with err: "+err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "audio/x-mpegurl")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.m3u\"", playlistID))

	var m3uContent string

	for _, trackID := range playlist.TrackIDs {
		// convert track_id string to objectID
		objID, err := primitive.ObjectIDFromHex(trackID)
		if err != nil {
			appG.Response500(e.ERROR, "convert id string to objectID failed"+err.Error())
			return
		}
		track, err := models.Repository.Track.FindByID(context.Background(), objID)
		if err != nil {
			log.Printf("Get track by id: %s failed with error: %v", trackID, err)
		} else {
			m3uContent += track.FileURL + "\n"
		}

	}

	_, err = c.Writer.Write([]byte(m3uContent))
	if err != nil {
		appG.Response500(e.ERROR, "Write to  m3u failed: "+err.Error())
		return
	}
}
