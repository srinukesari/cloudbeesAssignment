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

func CreatePostUserInput() *pb.CreatePostRequest {
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
	fmt.Print("enter publicationdate in YYYY-MM-DD format -> ")
	date, _ := reader.ReadString('\n')

	date = strings.TrimSpace(date)
	timestamp, _ := time.Parse("2006-01-02", date)
	timeInProtoFormat := timestamppb.New(timestamp)

	postReq := &pb.CreatePostRequest{
		CreatePostRequest: &pb.BlogPost{
			PostId:          postid,
			Author:          strings.TrimSpace(author),
			Title:           strings.TrimSpace(title),
			Content:         strings.TrimSpace(content),
			PublicationDate: timeInProtoFormat,
		},
	}
	return postReq
}

func GetPostUserInput() *pb.GetPostRequest {
	var postID int32
	fmt.Print("Enter post id to get -> ")
	fmt.Scan(&postID)
	return &pb.GetPostRequest{
		PostId: postID,
	}
}

func UpdatePostUserInput(client pb.BlogServiceClient) *pb.UpdatePostRequest {
	var postID int32
	var mask string
	fmt.Print("Enter post id to Update -> ")
	fmt.Scan(&postID)
	fmt.Println("feilds that can updated are author,title,content and publication_date")
	fmt.Print("Enter mask(updated fields) with comma seperator -> ")
	fmt.Scan(&mask)
	updateRequest := &pb.UpdatePostRequest{
		UpdatePostRequest: &pb.BlogPost{
			PostId: postID,
		},
		Mask: mask,
	}

	resp, err := client.GetPost(context.Background(), &pb.GetPostRequest{PostId: postID})
	if err != nil {
		fmt.Errorf("post not found to update")
		return nil
	}

	reader := bufio.NewReader(os.Stdin)

	updatedMask := strings.Split(mask, ",")
	for _, i := range updatedMask {
		switch strings.ToLower(i) {
		case constant.Title:
			fmt.Print("present Title = ", resp.GetPostResponse.Title, "\nplease enter title to update -> ")
			title, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Title = strings.TrimSpace(title)

		case constant.Content:
			fmt.Print("present Content = ", resp.GetPostResponse.Content, "\nplease enter Content to update -> ")
			content, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Content = strings.TrimSpace(content)

		case constant.Author:
			fmt.Print("present Author = ", resp.GetPostResponse.Author, "\nplease enter author to update -> ")
			author, _ := reader.ReadString('\n')
			updateRequest.UpdatePostRequest.Author = strings.TrimSpace(author)

		case constant.Tags:
			fmt.Print("present Tags = ", resp.GetPostResponse.Tags, "\nplease enter updated tags as comma seperated values -> ")
			tags, _ := reader.ReadString('\n')
			tagsList := strings.Split(strings.TrimSpace(tags), ",")
			updateRequest.UpdatePostRequest.Tags = tagsList

		case constant.PublicationDate:
			fmt.Print("present PublicationDate = ", resp.GetPostResponse.PublicationDate.AsTime(), "\nplease enter Publication Date to update in YYYY-MM-DD  -> ")
			date, _ := reader.ReadString('\n')
			date = strings.TrimSpace(date)
			timestamp, _ := time.Parse("2006-01-02", date)
			timeInProtoFormat := timestamppb.New(timestamp)
			updateRequest.UpdatePostRequest.PublicationDate = timeInProtoFormat
		}
	}
	return updateRequest
}

func DeletePostUserInput() *pb.DeletePostRequest {
	var postID int32
	fmt.Print("Enter post id to delete -> ")
	fmt.Scan(&postID)
	return &pb.DeletePostRequest{
		PostId: postID,
	}
}
