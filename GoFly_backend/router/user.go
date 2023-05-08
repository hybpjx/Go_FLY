package router

import (
	"github.com/gin-gonic/gin"
	"gofly/api"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userAPI := api.NewUserAPI()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userAPI.Login)
		}
		rgAuthUser := rgAuth.Group("user")

		{
			rgAuthUser.POST("", userAPI.AddUser)
			rgAuthUser.POST("list", userAPI.GetUserList)
			rgAuthUser.GET("/:id", userAPI.GetUserByID)
			rgAuthUser.PUT("/:id", userAPI.UpdateUser)
			rgAuthUser.DELETE("/:id", userAPI.DeleteUserByID)
		}

	})
}
