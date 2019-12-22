package main

import (
	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/trace"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	microLog "github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	opentracing "github.com/opentracing/opentracing-go"

	"lcb123/api-gateway/router"
)

func main() {
	//统一日志到服务的日志
	microLog.SetLogger(log.NewMicroLogger())

	/************************************/
	/********** 服务发现  etcd   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			config.C.Etcd,
		}
		op.Timeout = 5 * time.Second
	})
	/************************************/
	/********** 链路追踪  trace   ********/
	/************************************/
	trace.SetSamplingFrequency(50)
	t, io, err := trace.NewTracer(config.C.Service.Name, config.C.Jaeger)
	if err != nil {
		microLog.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	/************************************/

	/************************************/
	/********** gin  路由框架     ********/
	/************************************/
	//注册 gin  routers
	ginHandler := router.Init()

	// create new web service
	service := web.NewService(
		web.Name(config.C.Service.Name),
		web.Registry(reg),
		web.RegisterTTL(time.Second*15),      //重新注册时间
		web.RegisterInterval(time.Second*10), //注册过期时间
		web.Version(config.C.Service.Version),
		web.Address(config.C.Service.Port),
		web.Handler(ginHandler),
	)

	// initialise service
	if err := service.Init(); err != nil {
		microLog.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		microLog.Fatal(err)
	}
}
