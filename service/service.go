package service

import (
	"context"
	"fmt"
	"go-distributed-system/registry"
	"log"
	"net/http"
)

//公共函数
//用来集中启动服务，任何服务只要提供了内容，服务名称，主机，端口号以及处理用的handler，都可以通过这个公共函数启动

func Start(ctx context.Context, host, port string, reg registry.Registration,
	handlerFunc func()) (context.Context, error) {
	//先运行处理请求用的handler
	handlerFunc()
	//启动服务
	ctx = startService(ctx, reg.ServiceName, host, port)

	//注册
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

//启动服务函数

func startService(ctx context.Context, serviceName registry.ServiceName,
	host, port string) context.Context {

	//调用方法返回ctx的一个拷贝和一个取消函数
	ctx, cancel := context.WithCancel(ctx)

	//定义一个服务器
	var server http.Server
	//初始化服务器的端口号
	server.Addr = ":" + port //如 ":8080"

	//使用服务器监听并处理web请求
	//在这里使用
	go func() {
		log.Println(server.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. press any key to stop. \n", serviceName)
		var s string
		//用户可以通过输入任何值终止服务
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
		//关闭服务
		server.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
