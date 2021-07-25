package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/wenyunji/common"
	"github.com/wenyunji/payment/domain/repository"
	service2 "github.com/wenyunji/payment/domain/service"
	"github.com/wenyunji/payment/handler"
	"github.com/wenyunji/payment/proto/payment"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		common.Error(err)
	}
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.0:8500",
		}
	})

	//jaeger 链路追踪
	t, io, err := common.NewTracer("go.micro.service.payment", "localhost:6831")
	if err != err {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//mysql 设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	//初始化数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Password+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		common.Error(err)
	}
	defer db.Close()
	//禁止复数表
	db.SingularTable(true)

	//创建表
	tableInit := repository.NewPaymentRepository(db)
	tableInit.InitTable()

	//监控
	common.PrometheusBoot(9089)

	// Create service
	srv := micro.NewService(
		micro.Name("payment"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),
		//注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//加载限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		//加载监控 NewHandlerWrapper
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	srv.Init()
	paymentService := service2.NewPaymentService(repository.NewPaymentRepository(db))
	payment.RegisterPaymentHandler(srv.Server(), &handler.Payment{Payment: paymentService})

	// Run service
	if err := srv.Run(); err != nil {
		common.Error(err)
	}
}
