package membership

import (
	"context"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	finalcachepb "github.com/nienie/finalcache/pb"
	"github.com/nienie/finalcache/pb/testimp"
)

func TestService(t *testing.T) {
	//初始化服务
	name := "test"
	service1 := NewService(name)
	service2 := NewService(name)
	service3 := NewService(name)
	finalcachepb.RegisterFinalCacheHandler(service1.Server(), &testimp.FinalCacheImpl{})
	finalcachepb.RegisterFinalCacheHandler(service2.Server(), &testimp.FinalCacheImpl{})
	finalcachepb.RegisterFinalCacheHandler(service3.Server(), &testimp.FinalCacheImpl{})

	//启动服务
	go func() {
		err := service1.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()
	go func() {
		err := service2.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()
	go func() {
		err := service3.Run()
		t.Logf("err=%v", err)
		t.FailNow()
	}()
	time.Sleep(1 * time.Second)

	//验证服务是否在服务发现上正常
	services, err := service1.Options().Registry.GetService(GetServiceName(name))
	require.NoError(t, err)
	t.Logf(spew.Sprintf("services=%v", services))
	require.Len(t, services, 1)
	service := services[0]
	t.Logf(spew.Sprintf("service=%v", service))
	require.Len(t, service.Nodes, 3) //验证节点数是否是3

	//grpc调用接口，判断服务示范正常
	client := finalcachepb.NewFinalCacheService(GetServiceName(name), service1.Client())
	//调用Get接口
	getResp, err := client.Get(context.Background(), &finalcachepb.GetRequest{})
	require.NoError(t, err)
	t.Logf("GetResponse=%+v", getResp)
	//调用Set接口
	setResp, err := client.Set(context.Background(), &finalcachepb.SetRequest{})
	require.NoError(t, err)
	t.Logf("SetResponse=%+v", setResp)
	//调用Del接口
	delResp, err := client.Del(context.Background(), &finalcachepb.DelRequest{})
	require.NoError(t, err)
	t.Logf("DelResponse=%+v", delResp)
}