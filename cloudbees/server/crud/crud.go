package crud

import (
	constant "cloudbees"
	pb "cloudbees/proto"
	"context"
	"fmt"
	"strings"
)

type BlogServer struct {
	pb.UnimplementedBlogServiceServer
}

var blogPost = make(map[int32]*pb.BlogPost)

func (s *BlogServer) CreatePost(ctx context.Context, in *pb.CreatePostRequest,
) (*pb.CreatePostResponse, error) {

	err := ValidateCreatePost(in)
	if err != nil {
		return nil, err
	}

	if _, isExists := blogPost[in.CreatePostRequest.PostId]; !isExists {
		blogPost[in.CreatePostRequest.PostId] = in.CreatePostRequest

		return &pb.CreatePostResponse{
			CreatePostResponse: blogPost[in.CreatePostRequest.PostId],
			Message:            "post created successfully",
		}, nil
	}
	return nil, fmt.Errorf("duplicate post ID found")
}

func (s *BlogServer) GetPost(ctx context.Context, in *pb.GetPostRequest,
) (*pb.GetPostResponse, error) {
	err := ValidateGetPost(in)
	if err != nil {
		return nil, err
	}

	if val, isExists := blogPost[in.PostId]; isExists {
		return &pb.GetPostResponse{
			GetPostResponse: val,
		}, nil
	}

	return &pb.GetPostResponse{
		Error: "post not found",
	}, nil
}

func (s *BlogServer) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest,
) (*pb.UpdatePostResponse, error) {

	err := ValidateUpdatePost(in)
	if err != nil {
		return nil, err
	}

	postID := in.UpdatePostRequest.GetPostId()
	masks := strings.Split(in.Mask, ",")

	if post, isExists := blogPost[postID]; isExists {
		for _, val := range masks {
			switch strings.ToLower(val) {
			case constant.Author:
				post.Author = in.UpdatePostRequest.Author
			case constant.Title:
				post.Title = in.UpdatePostRequest.Title
			case constant.Content:
				post.Content = in.UpdatePostRequest.Content
			case constant.Tags:
				post.Tags = in.UpdatePostRequest.Tags
			default:
				continue
			}
		}
		blogPost[postID] = post
		return &pb.UpdatePostResponse{
			UpdatePostResponse: blogPost[postID],
			Message:            "updated successfully",
		}, nil
	}
	return &pb.UpdatePostResponse{
		UpdatePostResponse: blogPost[postID],
		Error:              "post not found to update",
	}, nil
}

func (s *BlogServer) DeletePost(ctx context.Context, in *pb.DeletePostRequest,
) (*pb.DeletePostResponse, error) {
	err := ValidateDeletePost(in)
	if err != nil {
		return nil, err
	}

	if _, isExists := blogPost[in.PostId]; isExists {
		delete(blogPost, in.PostId)
		return &pb.DeletePostResponse{
			Message: "delete successfully",
		}, nil
	}

	return &pb.DeletePostResponse{
		Error: "post doesn't exists",
	}, nil
}
