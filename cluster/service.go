package cluster

import (
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/registry/mdns"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/service/grpc"
)

//NewService ...
func NewService(name string) service.Service {
	opt := []service.Option{
		service.Name(GetServiceName(name)),
		service.Registry(mdns.NewRegistry()),
		service.RegisterInterval(10 * time.Second),
		service.RegisterTTL(30 * time.Second),
		setServerID(uuid.New().String()),
	}
	return grpc.NewService(opt...)
}

func setServerID(id string) service.Option {
	return func(o *service.Options) {
		o.Server.Init(func(opt *server.Options) {
			opt.Id = id
		})
	}
}

//GetServiceName ...
func GetServiceName(name string) string  {
	return "finalcache." + name
}