package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro/v2"
	userProto "micro-service/proto/user"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("user.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	user := userProto.NewUserService("user", service.Client())

	// Call the user
	rsp, err := user.QueryUserByName(context.TODO(), &userProto.Request{UserName: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.GetUser())
}
