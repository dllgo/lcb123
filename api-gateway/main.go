package main

import (
	"lcb123/pkg/config"
	"net/http"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"

	"lcb123/api-gateway/handler"
)

func main() {
	/************************************/
	/********** 服务发现  cousul   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			config.C.Etcd,
		}
	})
	// create new web service
	service := web.NewService(
		web.Name(config.C.Service.Name),
		web.Registry(reg),
		web.RegisterTTL(time.Second*15),      //重新注册时间
		web.RegisterInterval(time.Second*10), //注册过期时间
		web.Version(config.C.Service.Version),
		web.Address(config.C.Service.Port),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/api/call", handler.ApiCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
