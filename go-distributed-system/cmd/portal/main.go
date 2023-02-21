package main

import (
	"context"
	"fmt"
	"go-distributed-system/log"
	"go-distributed-system/portal"
	"go-distributed-system/registry"
	"go-distributed-system/service"
	stlog "log"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "127.0.0.1", "5001"
	serviceAddr := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  serviceAddr,
		RequiredServices: []registry.ServiceName{
			registry.LogService,
			registry.GradingService,
		},
		ServiceUpdateURL: serviceAddr + "/services",
		HeartBeatURL:     serviceAddr + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		portal.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	//start 中使用了goroutine，会发送响应的信号
	<-ctx.Done()
	fmt.Println("Shutting down portal.")
}
