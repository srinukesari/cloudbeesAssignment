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
		var err error
		switch int(op) {
		case 1:
			postReq := crud.CreatePostUserInput()
			err = crud.CreatePost(ctx, client, postReq)
		case 2:
			getReq := crud.GetPostUserInput()
			err = crud.GetPost(ctx, client, getReq)
		case 3:
			updateReq := crud.UpdatePostUserInput()
			err = crud.UpdatePost(ctx, client, updateReq)
		case 4:
			deleteReq := crud.DeletePostUserInput()
			err = crud.DeletePost(ctx, client, deleteReq)
		default:
			break clientLoop
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
