package handler

import (
	"context"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/micro/go-log"

	example "micro-service/proto/example"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Log("Received Example.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
	log.Logf("Received Example.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&example.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

// SayHello implements helloworld.GreeterServer
func (e *Example) Md2Html(ctx context.Context, req *example.MethodRequest, res *example.MethodReply) error {
	log.Logf("Received: %v", req.MarkdownStr)
	//先做一个markdown转HTML的练习和go module依赖管理
	input := []byte(req.MarkdownStr)
	unsafe := blackfriday.Run(input)
	htmlByte := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	log.Logf("Reply:", string(htmlByte))
	res.Html=string(htmlByte)//获取到的时Unicode码，蛋疼
	return nil
}