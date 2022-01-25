package grpcChat

import (
	"log"
	"net"

	"github.com/Joe-Degs/golang-journey/systems_golang/grpc-chat/chat"
	"google.golang.org/grpc"
)

func Main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("starting grpc server on %s", lis.Addr().String())
	s := &chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
