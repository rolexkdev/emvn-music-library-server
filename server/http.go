package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/docs"
	v1 "github.com/rolexkdev/emvn-music-library-server/internal/handlers"
)

func InitHttpRoutes(r *gin.RouterGroup, conf *config.Config) {
	router := r.Group("/")

	docs.InitSwaggerRoute(router.Group("/swagger"), conf)

	// Upload
	upload := router.Group("/uploads")
	upload.POST("", v1.UploadFile)
	upload.GET("/:filename", v1.RetrieveFile)

	//tracks
	tracks := router.Group("/tracks")
	tracks.POST("", v1.CreateTrack)
	tracks.GET("", v1.GetTracks)
	tracks.GET("/:id", v1.GetTrack)
	tracks.DELETE("/:id", v1.DeleteTrack)
	tracks.PUT("/:id", v1.UpdateTrack)

	//playlist
	playlists := router.Group("/playlists")
	playlists.POST("", v1.CreatePlaylist)
	playlists.GET("", v1.GetPlaylists)
	playlists.GET("/:id", v1.GetPlaylist)
	playlists.DELETE("/:id", v1.DeletePlaylist)
	playlists.PUT("/:id", v1.UpdatePlaylist)
	playlists.POST("/:id/tracks", v1.UpdatePlaylistTrack)
	playlists.GET("/:id/m3u", v1.GenerateM3UPlaylist)

	//search
	search := router.Group("/search")
	search.GET("", v1.Search)

}
