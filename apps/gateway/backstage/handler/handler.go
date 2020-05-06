package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"

	"platform.framework/apps/services/example/service"
	"platform.framework/library/logger"
)

// Hello hello api.
func Hello(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	srv := service.NewExampleSrvService("platform.services.example", client.DefaultClient)
	rsp, err := srv.Hello(ctx, &service.HelloRequest{Name: name})
	if err != nil {
		logger.ErrorF("Hello call platform.services.example error: %v", err)
		ctx.JSON(500, nil)
		return
	}

	ctx.JSON(200, rsp)
}
