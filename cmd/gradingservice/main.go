package main

import (
	"context"
	"fmt"
	"go-distributed-system/grades"
	"go-distributed-system/log"
	"go-distributed-system/registry"
	"go-distributed-system/service"
	stlog "log"
)

func main() {
	host, port := "127.0.0.1", "6000"
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddr,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddr + "/services",
		HeartBeatURL:     serviceAddr + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(), host, port, r, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	//start 中使用了goroutine，会发送响应的信号
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
