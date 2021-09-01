package initialize

import (
	"github.com/gin-gonic/gin"
	"go-demo/routers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	routers.UserRoute(router)

	return router
}
