package main

import (
	"context"
	"fmt"
	"go-distributed-system/log"
	"go-distributed-system/registry"
	"go-distributed-system/service"
	stlog "log"
)

func main() {

	//先运行log服务
	log.Run("distributed.log")

	//接着启动服务
	//实际业务中，如下信息应该从环境变量或者配置文件中获取
	host, port := "127.0.0.1", "4000"

	//使用公共函数启用
	//没有ctx，就使用context.Background,它是一个非nil的空值
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddr,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddr + "/services",
		HeartBeatURL:     serviceAddr + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandler)
	if err != nil {
		stlog.Fatalln(err)
	}
	//start 中使用了goroutine，会发送响应的信号
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
