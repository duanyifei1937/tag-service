package main

import (
	"log"
	"net"
	"tag-service/server"

	"google.golang.org/grpc/reflection"

	pb "tag-service/proto"

	grpc "google.golang.org/grpc"
)

func main() {
	port := "8888"
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
