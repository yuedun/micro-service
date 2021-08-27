package handler

import (
	"context"
	"encoding/json"
	post "micro-service/proto/post"

	"github.com/asim/go-micro/v3/logger"

	api "github.com/asim/go-micro/v3/api/proto"

	"strconv"
	"strings"

	"github.com/asim/go-micro/v3/errors"
)

type Article struct {
	Client post.PostService
}

// http://localhost:8080/user/article/getArticle?id=1
func (s *Article) GetArticle(ctx context.Context, req *api.Request, rsp *api.Response) error {
	logger.Info("Received Article.GetArticle API request")

	ID, ok := req.Get["id"]
	if !ok || len(ID.Values) == 0 {
		return errors.BadRequest("go.micro.api.article", "id cannot be blank")
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
