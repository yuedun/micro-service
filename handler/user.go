package handler

import (
	"context"
	"github.com/micro/go-log"

	user "micro-service/proto/user"
)

type User struct{}

// 实现了user.pb.micro.go中的UserHandler接口
func (e *User) QueryUserByName(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Log("Received QueryUserByName request:", req.GetUserName())
	//rsp.User.Name = "Hello " + req.UserName//rsp.User是零值（nil），不能直接对其属性赋值，所以需要创建新对象赋值
	rsp.User = &user.User{Name: req.UserName}
	rsp.Success = true
	return nil
}
