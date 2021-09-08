package initialize

import (
	"github.com/gin-gonic/gin"
	"go-demo/api"
	"go-demo/routers"

	"go-demo/middleware/jwt"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login/", api.LoginUser)
	router.POST("/user/", api.SaveUser)
	router.Use(jwt.Auth())
	routers.UserRoute(router)

	return router
}
