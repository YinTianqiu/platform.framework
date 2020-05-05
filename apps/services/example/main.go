package main

import (
	"platform.framework/apps/services/example/handler"
	"platform.framework/apps/services/example/service"
	basicservice "platform.framework/library/basic.service"
	"platform.framework/library/logger"
)

func main() {

	handle := handler.NewExampleSrv()
	engine := basicservice.NewBasicService()
	if err := service.RegisterExampleSrvHandler(engine.Server(), handle); err != nil {
		logger.ErrorF("RegisterExampleSrvHandler error: %v", err)
	}
	if err := engine.Run(); err != nil {
		logger.ErrorF("Run error: %v", err)
		return
	}

}
