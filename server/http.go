package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/docs"
)

func InitHttpRoutes(r *gin.RouterGroup, conf *config.Config) {
	router := r.Group("/")

	docs.InitSwaggerRoute(router.Group("/swagger"), conf)
}
