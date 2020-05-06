package main

import (
	"platform.framework/apps/gateway/backstage/routers"
	basicgateway "platform.framework/library/basic.gateway"
	"platform.framework/library/logger"
)

func main() {
	gateway := basicgateway.NewGateway(routers.NewRouter())

	if err := gateway.Run(); err != nil {
		logger.ErrorF("Run error: %v", err)
		return
	}
}
