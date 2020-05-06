package main

import (
	"platform.framework/apps/services/example/handler"
	"platform.framework/apps/services/example/service"
	basicservice "platform.framework/library/basic.service"
	"platform.framework/library/logger"
)

func main() {
	engine := basicservice.NewBasicService()

	if err := service.RegisterExampleSrvHandler(engine.Server(), handler.NewExampleSrv()); err != nil {
		logger.ErrorF("RegisterExampleSrvHandler error: %v", err)
		return
	}

	if err := engine.Run(); err != nil {
		logger.ErrorF("Run error: %v", err)
		return
	}
}
