package handler

import (
	"context"
	"strconv"

	"github.com/asim/go-micro/v3/logger"

	post "micro-service/proto/post"
)

type Post struct{}

// 实现了user.pb.micro.go中的UserHandler接口
func (e *Post) QueryUserPosts(ctx context.Context, req *post.Request, rsp *post.Response) error {
	logger.Info("Received QueryUserPosts request:", req.GetPostID())
	ID64, _ := strconv.ParseInt(req.PostID, 10, 64)
	rsp.Post = &post.Post{
		Id:    ID64,
		Title: "这是通过id获取到的文章标题", //req.Title,
	}
	rsp.Success = true
	rsp.Error = nil
	return nil
}
