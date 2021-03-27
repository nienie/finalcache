package testimp

import (
	"context"
	"fmt"

	finalcachepb "github.com/nienie/finalcache/pb"
)

//FinalCacheImpl ...
type FinalCacheImpl struct {}

//Get ...
func (o *FinalCacheImpl) Get(ctx context.Context, in *finalcachepb.GetRequest, out *finalcachepb.GetResponse) error {
	out.Result = &finalcachepb.Result{
		Code: 1,
		Msg: "FinalCacheImpl.Get",
	}
	fmt.Printf("msg=FinalCacheImpl.Get||in=%+v||out=%+v\n", in, out)
	return nil
}

//Set ...
func (o *FinalCacheImpl) Set(ctx context.Context, in *finalcachepb.SetRequest, out *finalcachepb.SetResponse) error {
	out.Result = &finalcachepb.Result{
		Code: 1,
		Msg: "FinalCacheImpl.Set",
	}
	fmt.Printf("msg=FinalCacheImpl.Set||in=%+v||out=%+v\n", in, out)
	return nil
}

//Del ...
func (o *FinalCacheImpl) Del(ctx context.Context, in *finalcachepb.DelRequest, out *finalcachepb.DelResponse) error {
	out.Result = &finalcachepb.Result{
		Code: 1,
		Msg: "FinalCacheImpl.Del",
	}
	fmt.Printf("msg=FinalCacheImpl.Del||in=%+v||out=%+v\n", in, out)
	return nil
}
