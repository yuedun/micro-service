package handler

import (
	"context"
	"strconv"

	"github.com/micro/go-log"

	post "micro-service/proto/post"
)

type Post struct{}

// 实现了user.pb.micro.go中的UserHandler接口
func (e *Post) QueryUserPosts(ctx context.Context, req *post.Request, rsp *post.Response) error {
	log.Log("Received QueryUserPosts request:", req.GetPostID())
	ID64, _ := strconv.ParseInt(req.PostID, 10, 64)
	rsp.Post = &post.Post{
		Id:    ID64,
		Title: req.Title,
	}
	rsp.Success = true
	rsp.Error = nil
	return nil
}
