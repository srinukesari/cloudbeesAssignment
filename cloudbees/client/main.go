package main

import (
	constant "cloudbees"
	"context"
	"fmt"
	"log"

	"cloudbees/client/crud"
	pb "cloudbees/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost"+constant.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("connected to server, enter number for below operation")
	var op int
clientLoop:
	for {
		fmt.Print(crud.ClientMenu())
		fmt.Scan(&op)
		switch op {
		case 1:
			crud.CreatePostInput(ctx, client)
		case 2:
			crud.GetPostInput(ctx, client)
		case 3:
			crud.UpdatePostInput(ctx, client)
		case 4:
			crud.DeletePostInput(ctx, client)
		default:
			break clientLoop
		}
	}
}
