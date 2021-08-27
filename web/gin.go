package main

import (
	"context"

	"github.com/asim/go-micro/v3/logger"

	user "micro-service/proto/user"

	micro "github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
)

// 该模块演示的是在常规的web服务中调用微服务，比如在gin，beego，echo搭建的服务中调用微服务
type Say struct{}

var (
	cl user.UserService
)

func (s *Say) Anything(c *gin.Context) {
	logger.Info("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Say) Hello(c *gin.Context) {
	logger.Info("Received Say.Hello API request")

	name := c.Param("name")

	response, err := cl.QueryUserByName(context.TODO(), &user.Request{
		UserName: name,
	})

	if err != nil {
		logger.Error(err)
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main1() {
	// Create a new service
	service := micro.NewService(micro.Name("web.gin"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	cl = user.NewUserService("go.micro.srv.user", service.Client())

	// Create RESTful handler (using Gin)
	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)
	router.Run(":8081")
}
