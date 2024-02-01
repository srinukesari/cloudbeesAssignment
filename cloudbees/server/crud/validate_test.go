package crud

import (
	"testing"

	pb "cloudbees/proto"
)

func TestValidateCreatePost(t *testing.T) {
	tests := []struct {
		name          string
		createPostReq *pb.CreatePostRequest
		wantErr       bool
	}{
		{
			name: "post id not found error",
			createPostReq: &pb.CreatePostRequest{
				CreatePostRequest: &pb.BlogPost{
					Title:  "test title",
					Author: "ssr",
				},
			},
			wantErr: true,
		},
		{
			name: "valid create post request",
			createPostReq: &pb.CreatePostRequest{
				CreatePostRequest: &pb.BlogPost{
					PostId: 333,
					Title:  "test title",
					Author: "ssr",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreatePost(tt.createPostReq)

			if (err != nil) != tt.wantErr {
				t.Fatal("ValidateCreatePost() error -> ", err)
			}
		})
	}
}

func TestValidateGetPost(t *testing.T) {
	tests := []struct {
		name       string
		getPostReq *pb.GetPostRequest
		wantErr    bool
	}{
		{
			name:       "post id not found error",
			getPostReq: &pb.GetPostRequest{},
			wantErr:    true,
		},
		{
			name: "valid Get post request",
			getPostReq: &pb.GetPostRequest{
				PostId: 332,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetPost(tt.getPostReq)

			if (err != nil) != tt.wantErr {
				t.Fatal("ValidateGetPost() error -> ", err)
			}
		})
	}
}

func TestValidateUpdatePost(t *testing.T) {
	tests := []struct {
		name          string
		updatePostReq *pb.UpdatePostRequest
		wantErr       bool
	}{
		{
			name:          "Invalid Update posy req",
			updatePostReq: &pb.UpdatePostRequest{},
			wantErr:       true,
		},
		{
			name: "valid update post request",
			updatePostReq: &pb.UpdatePostRequest{
				UpdatePostRequest: &pb.BlogPost{
					PostId: 211,
					Author: "ssr",
				},
				Mask: "author",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdatePost(tt.updatePostReq)

			if (err != nil) != tt.wantErr {
				t.Fatal("ValidateUpdatePost() error -> ", err)
			}
		})
	}
}

func TestValidateDeletePost(t *testing.T) {
	tests := []struct {
		name          string
		deletePostReq *pb.DeletePostRequest
		wantErr       bool
	}{
		{
			name:          "Invalid delete post req",
			deletePostReq: &pb.DeletePostRequest{},
			wantErr:       true,
		},
		{
			name: "valid delete post request",
			deletePostReq: &pb.DeletePostRequest{
				PostId: 323,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDeletePost(tt.deletePostReq)

			if (err != nil) != tt.wantErr {
				t.Fatal("ValidateDeletePost() error -> ", err)
			}
		})
	}
}
