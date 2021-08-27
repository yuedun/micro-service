package handler

import (
	"context"
	"encoding/json"

	"github.com/asim/go-micro/v3/logger"

	user "micro-service/proto/user"
	"strconv"
	"strings"

	api "github.com/asim/go-micro/v3/api/proto"
	"github.com/asim/go-micro/v3/errors"
)

type Say struct {
	Client user.UserService
}

//http://localhost:8080/user/say/hello?name=sdgfh
func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	logger.Info("Received Say.Hello API request")

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
