package router

import (
	"github.com/gin-gonic/gin"
	"gofly/api"
)

func InitHostRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostAPI := api.NewHostAPI()
		rgAuthHost := rgAuth.Group("/host")
		{

			rgAuthHost.POST("/shutdown", hostAPI.ShutDown)
		}

	})
}
