package cluster

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"github.com/davecgh/go-spew/spew"

	finalcachepb "github.com/nienie/finalcache/pb"
	"github.com/nienie/finalcache/pb/testimp"
)

func TestService(t *testing.T) {
	name := "test"

	service1 := NewService(name)
	service2 := NewService(name)
	service3 := NewService(name)

	finalcachepb.RegisterFinalCacheHandler(service1.Server(), &testimp.FinalCacheImpl{})
	finalcachepb.RegisterFinalCacheHandler(service2.Server(), &testimp.FinalCacheImpl{})
	finalcachepb.RegisterFinalCacheHandler(service3.Server(), &testimp.FinalCacheImpl{})

	//启动service1
	go func() {
		err := service1.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()

	//启动service2
	go func() {
		err := service2.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()

	//启动service3
	go func() {
		err := service3.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()

	time.Sleep(1 * time.Second)

	services, err := service1.Options().Registry.GetService(GetServiceName(name))
	require.NoError(t, err)
	t.Logf(spew.Sprintf("services=%v", services))
	require.Len(t, services, 1)
	service := services[0]
	t.Logf(spew.Sprintf("service=%v", service))
	require.Len(t, service.Nodes, 3)

	client := finalcachepb.NewFinalCacheService(GetServiceName(name), service1.Client())
	getResp, err := client.Get(context.Background(), &finalcachepb.GetRequest{})
	require.NoError(t, err)
	t.Logf("GetResponse=%+v", getResp)

	setResp, err := client.Set(context.Background(), &finalcachepb.SetRequest{})
	require.NoError(t, err)
	t.Logf("SetResponse=%+v", setResp)

	delResp, err := client.Del(context.Background(), &finalcachepb.DelRequest{})
	require.NoError(t, err)
	t.Logf("DelResponse=%+v", delResp)
}