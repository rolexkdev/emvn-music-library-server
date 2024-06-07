package dto

import "github.com/rolexkdev/emvn-music-library-server/internal/models"

type SearchRequest struct {
	Query string `form:"query"`
}

type SearchResponse struct {
	Tracks    []*models.Track    `json:"tracks"`
	Playlists []*models.Playlist `json:"playlists"`
}
