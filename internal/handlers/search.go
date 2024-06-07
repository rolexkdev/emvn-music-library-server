package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/common/app"
	"github.com/rolexkdev/emvn-music-library-server/common/e"
	"github.com/rolexkdev/emvn-music-library-server/dto"
	"github.com/rolexkdev/emvn-music-library-server/internal/models"
)

// Search godoc
//
//	@Summary		Search tracks and playlist
//	@Description	Search tracks and playlist
//	@Tags			search
//	@Accept			json
//	@Produce		json
//
//	@Param			query		    query		string	true	"search string"
//
//	@Success		200				{object}	app.Response
//	@Failure		400				{object}	app.Response
//	@Failure		500				{object}	app.Response
//	@Router			/search [get]
func Search(c *gin.Context) {
	appG := app.Gin{C: c}
	var req dto.SearchRequest
	if err := c.BindQuery(&req); err != nil {
		appG.Response400(e.INVALID_PARAMS, "Bind query failed"+err.Error())
		return
	}

	tracks, err := models.Repository.Track.Search(context.Background(), req.Query)
	if err != nil {
		appG.Response500(e.ERROR, "Search tracks failed"+err.Error())
		return
	}

	playlists, err := models.Repository.Playlist.Search(context.Background(), req.Query)
	if err != nil {
		appG.Response500(e.ERROR, "Search playlist failed"+err.Error())
		return
	}

	appG.Response200(dto.SearchResponse{
		Tracks:    tracks,
		Playlists: playlists,
	})
}
