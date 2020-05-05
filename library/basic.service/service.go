package basicservice

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-plugins/registry/consul/v2"

	configure "platform.framework/library/config"
)

// BasicService 基础服务封装.
type BasicService struct {
	service micro.Service
	configs config.Config
}

// NewBasicService 新的基础服务.
func NewBasicService(opts ...micro.Option) *BasicService {
	cfg, err := configure.LoadConfigFile("./config.json")
	if err != nil {
		panic(err)
	}

	msc, err := configure.WithMicroSrvConf(cfg)
	if err != nil {
		panic(err)
	}

	srv := &BasicService{
		service: micro.NewService(
			micro.Name(msc.Basic.Name),
			micro.Version(msc.Basic.Version),
			micro.Address(msc.Basic.Address),
			micro.Registry(consul.NewRegistry(registry.Addrs(msc.Basic.Registry))),
			micro.RegisterTTL(time.Second*time.Duration(msc.Basic.RegisterTTL)),
			micro.RegisterInterval(time.Second*time.Duration(msc.Basic.RegisterInterval)),
		),
		configs: cfg,
	}
	srv.service.Init(opts...)

	return srv
}

// Run 运行.
func (bs *BasicService) Run() error {
	if err := bs.service.Run(); err != nil {
		return err
	}
	return nil
}

// Server 获取MicroServer.
func (bs *BasicService) Server() server.Server {
	return bs.service.Server()
}

// Client 获取MicroClient.
func (bs *BasicService) Client() client.Client {
	return bs.service.Client()
}

// Config 获取配置文件.
func (bs *BasicService) Config() config.Config {
	return bs.configs
}
