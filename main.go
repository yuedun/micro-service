package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	// "github.com/micro/go-micro/v2/registry"
	// "github.com/micro/go-micro/v2/registry/etcd"
	"micro-service/handler"
	post "micro-service/proto/post"
	user "micro-service/proto/user"
)

func main() {
	// reg := etcd.NewRegistry(func(options *registry.Options) {
	// 	options.Addrs = []string{
	// 		"xx.xx.xx.xx:2379",
	// 	}
	// })
	// New Service
	service := micro.NewService(
		// micro.Registry(reg),
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))
	post.RegisterPostHandler(service.Server(), new(handler.Post))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
