package micros

import (
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/service/grpc"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"

	"lcb123/pkg/config"
	"lcb123/pkg/log"
	"lcb123/pkg/trace"
)

var service micro.Service

//get service
func GetService() micro.Service {
	return service
}

func init() {
	/************************************/
	/********统一日志到服务的日志   ********/
	/************************************/
	log.Info("grpc 初始化。。。" + config.C.Service.Name)
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
	/********** 服务发现  etcd   ********/
	/************************************/
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			config.C.Etcd,
		}
		op.Timeout = 5 * time.Second
	})
	/************************************/
	/********** New GRPC Service   ********/
	/************************************/
	service = grpc.NewService(
		micro.Name(config.C.Service.Name),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*15),      //重新注册时间
		micro.RegisterInterval(time.Second*10), //注册过期时间
		micro.Version(config.C.Service.Version),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(wrapperTrace.NewClientWrapper(opentracing.GlobalTracer())),
	)
	// Initialise service-'?*()
	service.Init()
	log.Info("grpc 初始化完成。。。" + config.C.Service.Name)
}
