package docs

import (
	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRoute(r *gin.RouterGroup, conf *config.Config) {
	serverConfig := conf.Server
	SwaggerInfo.Version = serverConfig.AppVersion
	SwaggerInfo.BasePath = "/api/v1"
	SwaggerInfo.Schemes = []string{"http", "https"}
	SwaggerInfo.Title = "EMVN Web Server APIs"
	SwaggerInfo.Description = "API documents for emvn music library server"

	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
