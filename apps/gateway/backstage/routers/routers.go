package routers

import (
	"github.com/gin-gonic/gin"
	"platform.framework/apps/gateway/backstage/handler"
)

// NewRouter 新的路由.
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("hello", handler.Hello)
	return router
}
