package handler

import (
	"context"
	"encoding/json"
	api "github.com/micro/go-micro/api/proto"
	post "micro-service/proto/post"

	"github.com/micro/go-micro/errors"
	"log"
	"strconv"
	"strings"
)

type Article struct {
	Client post.PostService
}

func (s *Article) GetArticle(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	ID, ok := req.Get["id"]
	if !ok || len(ID.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name cannot be blank")
	}

	// 在restful api中调用rpc服务
	response, err := s.Client.QueryUserPosts(ctx, &post.Request{
		PostID: strings.Join(ID.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"id":      strconv.FormatInt(response.Post.Id, 10),
		"title":   response.Post.Title,
		"content": response.Post.Content,
	})
	rsp.Body = string(b)

	return nil
}
