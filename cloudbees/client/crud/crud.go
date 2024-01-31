package crud

import (
	"bufio"
	constant "cloudbees"
	pb "cloudbees/proto"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ClientMenu() string {
	menu := "##########################\n1. CreatePost\n2. GetPost\n3. UpdatePost\n4. DeletePost\n5. Exit\nplease enter operation number -> "
	return menu
}

func CreatePostInput(ctx context.Context, client pb.BlogServiceClient) {
	var postid int32
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter post id -> ")
	fmt.Scan(&postid)
	fmt.Print("enter title -> ")
	title, _ := reader.ReadString('\n')
	fmt.Print("enter content -> ")
	content, _ := reader.ReadString('\n')
	fmt.Print("enter author -> ")
	author, _ := reader.ReadString('\n')
	fmt.Print("enter publicationdate in DD-MM-YYYY format")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	timestamp, _ := time.Parse("03-11-1998", date)

	_ = timestamppb.New(timestamp)

	postReq := &pb.CreatePostRequest{
		CreatePostRequest: &pb.BlogPost{
			PostId:  postid,
			Author:  strings.TrimSpace(author),
			Title:   strings.TrimSpace(title),
			Content: strings.TrimSpace(content),
			// Timestamp: pbs,
		},
	}

	resp, err := client.CreatePost(ctx, postReq)
	if err != nil {
		fmt.Println("error from create post request -> ", err)
		return
	}
	fmt.Println("response from create post request ->", resp)
}

func GetPostInput(ctx context.Context, client pb.BlogServiceClient) {
	var postID int32
	fmt.Print("Enter post id to get -> ")
	fmt.Scan(&postID)
	resp, err := client.GetPost(ctx, &pb.GetPostRequest{PostId: postID})
	if err != nil {
		fmt.Println("error from get post request -> ", err)
		return
	}
	fmt.Println("response from get post request ->", resp)
}

func UpdatePostInput(ctx context.Context, client pb.BlogServiceClient) {
	var postID int32
	var mask string
	fmt.Print("Enter post id to Update -> ")
	fmt.Scan(&postID)
	fmt.Print("Enter mask(updated fields) with comma seperator -> ")
	fmt.Scan(&mask)
	updateRequest := &pb.UpdatePostRequest{
		UpdatePostRequest: &pb.BlogPost{
			PostId: postID,
		},
		Mask: mask,
	}

	reader := bufio.NewReader(os.Stdin)

	updatedMask := strings.Split(mask, ",")
	for _, i := range updatedMask {
		switch strings.ToLower(i) {
		case constant.Title:
			fmt.Print("please enter updated title -> ")
			title, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Title = strings.TrimSpace(title)
		case constant.Content:
			fmt.Print("please enter updated Content -> ")
			content, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Content = strings.TrimSpace(content)
		case constant.Author:
			fmt.Print("please enter updated author -> ")
			author, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Author = strings.TrimSpace(author)
		case constant.Tags:
			var tags string
			fmt.Print("please enter updated tags as commaseperated values -> ")
			fmt.Scan(&tags)
			tagsList := strings.Split(tags, ",")
			updateRequest.UpdatePostRequest.Tags = tagsList
		}
	}
	resp, err := client.UpdatePost(ctx, updateRequest)
	if err != nil {
		fmt.Println("error from update post request -> ", err)
		return
	}
	fmt.Println("response from update post request ->", resp)
}

func DeletePostInput(ctx context.Context, client pb.BlogServiceClient) {
	var postID int32
	fmt.Print("Enter post id to delete -> ")
	fmt.Scan(&postID)
	resp, err := client.DeletePost(ctx, &pb.DeletePostRequest{PostId: postID})
	if err != nil {
		fmt.Println("error from delete post request -> ", err)
		return
	}
	fmt.Println("response from delete post request ->", resp)
}
