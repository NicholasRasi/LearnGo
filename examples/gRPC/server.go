package main

import (
	"LearnGo/examples/gRPC/chat"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
)

type Server struct {
    chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Received: %v", in.Body)
	return &chat.Message{Body: "Hello!"}, nil
}

func main()  {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	chat.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}