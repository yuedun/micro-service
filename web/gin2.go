package main

import (
	"context"
	"log"

	user "micro-service/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

// 该模块是作为微服务一部分来提供服务，和gin.go模块对比，该模块没有绑定端口，而是由框架来指定，会自动注册到服务发现中心，由micro api工具来发现并提供默认的8080端口
type Say2 struct{}

var (
	cl2 user.UserService
)

func (s *Say2) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Say2) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response, err := cl2.QueryUserByName(context.TODO(), &user.Request{
		UserName: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()

	// setup Greeter Server Client
	cl2 = user.NewUserService("go.micro.srv.user", client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say2)
	router := gin.Default()
	router.GET("/user", say.Anything)
	router.GET("/user/:name", say.Hello)

	// Register Handler
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
