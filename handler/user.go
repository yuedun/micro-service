package handler

import (
	"context"
	"github.com/micro/go-log"

	user "micro-service/proto/user"
)

type User struct{}
// 实现了user.pb.micro.go中的UserHandler接口
func (e *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received Example.Call request")
	rsp.User.Name = "Hello " + req.UserName
	return nil
}

