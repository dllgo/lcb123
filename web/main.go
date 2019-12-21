package main

import (
	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/queue"
	"lcb123/pkg/trace"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	microLog "github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/opentracing/opentracing-go"

	"lcb123/web/routers"
)

// @title Swagger Example API12222
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1

func main() {
	//统一日志到服务的日志
	microLog.SetLogger(log.NewMicroLogger())
	/************************************/
	/********** 服务发现  cousul   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			config.C.Etcd,
		}
	})
	/************************************/
	/********** gin  路由框架     ********/
	/************************************/
	//注册 gin  routers
	ginHandler := routers.Init()

	// create new api service
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
	if err := service.Init(
		web.Action(func(ctx *cli.Context) {
			/************************************/
			/********** 链路追踪  trace   ********/
			/************************************/
			trace.SetSamplingFrequency(50)
			t, io, err := trace.NewTracer(config.C.Service.Name, config.C.Jaeger)
			if err != nil {
				log.Fatal(err)
			}
			defer io.Close()
			opentracing.SetGlobalTracer(t)

			/************************************/
			/********** 消息队列  queue   ********/
			/************************************/
			queue.Init(config.C.Nsq.Address, config.C.Nsq.Lookup, config.C.Nsq.MaxInFlight)

		}),
	); err != nil {
		log.Error(err.Error())
	}

	// run service
	if err := service.Run(); err != nil {
		log.Error(err.Error())
	}
}
