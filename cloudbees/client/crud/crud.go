package crud

import (
	pb "cloudbees/proto"
	"context"
	"fmt"
)

func ClientMenu() string {
	menu := "##########################\n1. CreatePost\n2. GetPost\n3. UpdatePost\n4. DeletePost\n5. Exit\nplease enter operation number -> "
	return menu
}

func CreatePost(ctx context.Context, client pb.BlogServiceClient, postReq *pb.CreatePostRequest) error {
	resp, err := client.CreatePost(ctx, postReq)
	if err != nil {
		return fmt.Errorf("error from create post request -> %s", err)
	}
	fmt.Println("response from create post request ->", resp)
	return nil
}

func GetPost(ctx context.Context, client pb.BlogServiceClient, getPostReq *pb.GetPostRequest) error {
	resp, err := client.GetPost(ctx, getPostReq)
	if err != nil {
		return fmt.Errorf("error from get post request -> %s", err)
	}
	fmt.Println("response from get post request ->", resp)
	return nil
}

func UpdatePost(ctx context.Context, client pb.BlogServiceClient, updateRequest *pb.UpdatePostRequest) error {
	resp, err := client.UpdatePost(ctx, updateRequest)
	if err != nil {
		return fmt.Errorf("error from update post request -> %v", err)
	}
	fmt.Println("response from update post request ->", resp)
	return nil
}

func DeletePost(ctx context.Context, client pb.BlogServiceClient, deletePostReq *pb.DeletePostRequest) error {
	resp, err := client.DeletePost(ctx, deletePostReq)
	if err != nil {
		return fmt.Errorf("error from delete post request -> %v", err)
	}
	fmt.Println("response from delete post request ->", resp)
	return nil
}
