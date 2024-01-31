package main

import (
	constant "cloudbees"
	pb "cloudbees/proto"
	"cloudbees/server/crud"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", constant.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	blogServer := &crud.BlogServer{}

	pb.RegisterBlogServiceServer(grpcServer, blogServer)

	log.Println("Blog Server started listening on ", constant.Port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
