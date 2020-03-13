package main

import (
	"encoding/json"
	"log"
	"strings"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	user "micro-service/proto/user"

	"context"
)

type Say struct {
	Client user.UserService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name cannot be blank")
	}

	response, err := s.Client.QueryUserByName(ctx, &user.Request{
		UserName: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.User.Name,
	})
	rsp.Body = string(b)

	return nil
}

// 将grpc服务转为restful接口
func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			// 此处我要说一下，这个坑爹的指针类型没传对，搞了3天就是跑不通
			&Say{Client: user.NewUserService("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
