package routes

import (
	controller "github.com/lordscoba/golang-auth/pkg/controllers"
	"github.com/lordscoba/golang-auth/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/admin", controller.GetUsers())
	incomingRoutes.GET("/admin/:user_id", controller.GetUser())
	incomingRoutes.POST("/admin/create", controller.PostUser())
	incomingRoutes.PATCH("/admin/update/:user_id", controller.UpdateUser())
}
