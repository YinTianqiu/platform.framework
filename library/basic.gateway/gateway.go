package basicgateway

import (
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"

	"platform.framework/library/config"
)

// NewGateway 新的网关.
func NewGateway(handler http.Handler) web.Service {
	cfg, err := config.LoadConfigFile("./config.json")
	if err != nil {
		panic(err)
	}

	msc, err := config.WithMicroSrvConf(cfg)
	if err != nil {
		panic(err)
	}

	srv := web.NewService(
		web.Name(msc.Basic.Name),
		web.Version(msc.Basic.Version),
		web.Address(msc.Basic.Address),
		web.Registry(consul.NewRegistry(registry.Addrs(msc.Basic.Registry))),
		web.RegisterTTL(time.Second*time.Duration(msc.Basic.RegisterTTL)),
		web.RegisterInterval(time.Second*time.Duration(msc.Basic.RegisterInterval)),
	)
	srv.Init()
	srv.Handle("/", handler)

	return srv
}
