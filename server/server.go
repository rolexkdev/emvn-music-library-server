package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/config"
	"github.com/rolexkdev/emvn-music-library-server/middleware"
)

func InitServer(conf *config.Config) {
	serverConfig := conf.Server
	var appEngine *gin.Engine

	// App
	if serverConfig.Environment == config.DEVELOPMENT {
		gin.SetMode(gin.DebugMode)
		appEngine = gin.Default()
		appEngine.SetTrustedProxies([]string{"127.0.0.1"})
	} else {
		gin.SetMode(gin.ReleaseMode)
		appEngine = gin.Default()
		appEngine.SetTrustedProxies(nil)
	}

	// Middlewares
	appEngine.Use(middleware.CORS)

	router := appEngine.Group("/" + serverConfig.AppVersion)

	// Routes
	InitHttpRoutes(router, conf)

	appEngine.Run(serverConfig.Host + ":" + serverConfig.Port)
}
