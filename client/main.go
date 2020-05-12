package main

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	userProto "micro-service/proto/user"

	micro "github.com/micro/go-micro/v2"
)

// 使用方式：将下面代码插入到web，api项目中来调用user.QueryUserByName服务。可以参考web/gin.go。
func main() {
	// Create a new service
	service := micro.NewService(micro.Name("user.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	user := userProto.NewUserService("go.micro.srv.user", service.Client())

	// Call the user
	rsp, err := user.QueryUserByName(context.TODO(), &userProto.Request{UserName: "John"})
	if err != nil {
		logger.Fatal(err)
	}

	// Print response
	logger.Info(rsp.GetUser())
}
