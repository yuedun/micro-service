package main

import (
	"context"

	"github.com/asim/go-micro/v3/logger"

	postProto "micro-service/proto/post"
	userProto "micro-service/proto/user"

	micro "github.com/asim/go-micro/v3"
)

// 使用方式：将下面代码插入到web，api项目中来调用user.QueryUserByName服务。可以参考web/gin.go。
func main() {
	// Create a new service
	service := micro.NewService(micro.Name("user.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	user := userProto.NewUserService("go.micro.srv.user", service.Client())
	post := postProto.NewPostService("go.micro.srv.user", service.Client())

	// Call the user
	rsp, err := user.QueryUserByName(context.TODO(), &userProto.Request{UserName: "John"})
	if err != nil {
		logger.Fatal(err)
	}

	// Print response
	logger.Info(rsp.GetUser())

	rsp2, err2 := post.QueryUserPosts(context.TODO(), &postProto.Request{UserID: 1})
	if err2 != nil {
		logger.Fatal(err2)
	}
	// Print response
	logger.Info(rsp2.GetPost().Title)
}
