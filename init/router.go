package init

import (
	"github.com/gin-gonic/gin"
	"go-demo/routers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	routers.UserRouter(router)
	return router
}
