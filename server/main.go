/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"log"
	"net"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "micro-service/protos"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Md2Html(ctx context.Context, in *pb.MethodRequest) (*pb.MethodReply, error) {
	log.Printf("Received: %v", in.MarkdownStr)
	//先做一个markdown转HTML的练习和go module依赖管理
	input := []byte(in.MarkdownStr)
	unsafe := blackfriday.Run(input)
	htmlByte := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	log.Println("Reply:", string(htmlByte))
	return &pb.MethodReply{Html: string(htmlByte)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMarkdownServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
