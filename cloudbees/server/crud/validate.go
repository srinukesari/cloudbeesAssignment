package crud

import (
	pb "cloudbees/proto"
	"fmt"
)

func ValidateCreatePost(request *pb.CreatePostRequest) error {
	if request == nil {
		return fmt.Errorf("request can't be empty")
	}

	if request.GetCreatePostRequest() == nil {
		return fmt.Errorf("empty create post request")
	}

	if request.CreatePostRequest.GetPostId() == 0 {
		return fmt.Errorf("post ID can't be null")
	}

	if request.CreatePostRequest.GetAuthor() == "" {
		return fmt.Errorf("author can't be Empty")
	}

	return nil
}

func ValidateGetPost(request *pb.GetPostRequest) error {
	if request == nil {
		return fmt.Errorf("request can't be empty")
	}

	if request.GetPostId() == 0 {
		return fmt.Errorf("post ID can't be zero")
	}

	return nil
}

func ValidateUpdatePost(request *pb.UpdatePostRequest) error {
	if request == nil {
		return fmt.Errorf("request can't be empty")
	}

	if request.GetUpdatePostRequest() == nil || request.GetMask() == "" {
		return fmt.Errorf("UpdatePostRequest/Mask is empty")
	}

	if request.UpdatePostRequest.GetPostId() == 0 {
		return fmt.Errorf("post ID can't be zero")
	}

	return nil
}

func ValidateDeletePost(request *pb.DeletePostRequest) error {
	if request == nil {
		return fmt.Errorf("request can't be empty")
	}

	if request.GetPostId() == 0 {
		return fmt.Errorf("post ID can't be zero")
	}

	return nil
}
