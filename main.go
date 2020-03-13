package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"micro-service/handler"
	post "micro-service/proto/post"
	user "micro-service/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
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
