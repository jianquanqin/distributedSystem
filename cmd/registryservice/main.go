package main

import (
	"context"
	"fmt"
	"go-distributed-system/registry"
	"log"
	"net/http"
)

func main() {

	//健康检查
	registry.SetupRegistryService()

	//如下的函数比http.HandleFunc()对handler进行了更高级的封装
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var server http.Server
	server.Addr = registry.ServerPort

	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Println("Registry service started, Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
