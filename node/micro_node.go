package node

import (
	"context"

	"github.com/micro/go-micro/v2/service"
	"github.com/nienie/finalcache/consistenthash"
	"github.com/nienie/finalcache/pb"
)

//MicroNode ...
type MicroNode struct {
	service service.Service
	ring    *consistenthash.Map
}

//WhoAmI ...
func (o *MicroNode) WhoAmI() string {
	return o.service.Options().Server.Options().Advertise
}

//Lookup ...
func (o *MicroNode) Lookup(key string) string {
	return o.ring.Get(key)
}

//Run ...
func (o *MicroNode) Run() error {
	go func() {
		err := o.service.Run()
		if err != nil {
			//TODO: Add log
		}
	}()
	return nil
}

//Get ...
func (o *MicroNode) Get(context.Context, *finalcachepb.GetRequest, *finalcachepb.GetResponse) error {

	return nil
}

//Set ...
func (o *MicroNode) Set(context.Context, *finalcachepb.SetRequest, *finalcachepb.SetResponse) error {

	return nil
}

//Del ...
func (o *MicroNode) Del(context.Context, *finalcachepb.DelRequest, *finalcachepb.DelResponse) error {

	return nil
}
