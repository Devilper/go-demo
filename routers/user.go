package routers

import (
	"github.com/gin-gonic/gin"
	"go-demo/api"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user/", api.SaveUser)
	router.GET("/user/", api.GetUsers)
	router.GET("/user/:id", api.GetOneUser)
	router.PUT("/user/", api.UpdateUser)
	router.DELETE("/user/:id", api.DeleteUser)
}
