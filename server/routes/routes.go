package routes

import (
	"github.com/reinaldodribeiro/api-movon/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}

		main.POST("login", controllers.Login)
	}

	return router
}