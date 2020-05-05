package handler

import (
	"context"

	"platform.framework/apps/services/example/service"
	"platform.framework/common/errors"
	"platform.framework/library/logger"
)

// ExampleSrv 对example service protobuf 的接口实现.
type ExampleSrv struct{}

// NewExampleSrv 新的实例服务业务处理实现类.
func NewExampleSrv() *ExampleSrv {
	return &ExampleSrv{}
}

// Hello 实现Hello功能.
func (es *ExampleSrv) Hello(ctx context.Context, req *service.HelloRequest, rsp *service.HelloResponse) error {
	if req.Name == "" {
		logger.Error("ExampleSrv Hello request param name is empty string")
		return errors.NewError(errors.RequestParamIsNull)
	}
	rsp.Greeting = "hi " + req.Name
	return nil
}
