package handler

import (
	"context"
	"encoding/json"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"log"
	user "micro-service/proto/user"
	"strconv"
	"strings"
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

	// 在restful api中调用rpc服务
	response, err := s.Client.QueryUserByName(ctx, &user.Request{
		UserID:   "123",
		UserName: strings.Join(name.Values, " "),
		UserPwd:  "dslhgfoif40u9b9",
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"id":       strconv.FormatInt(response.User.Id, 10),
		"username": response.User.Name,
		"password": response.User.Pwd,
	})
	rsp.Body = string(b)

	return nil
}
