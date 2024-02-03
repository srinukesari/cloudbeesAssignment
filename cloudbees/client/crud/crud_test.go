package crud

import (
	"context"
	"fmt"
	"testing"

	pb "cloudbees/proto"
	mock "cloudbees/proto/mocks"

	"github.com/golang/mock/gomock"
)

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serverMock := mock.NewMockBlogServiceClient(ctrl)

	tests := []struct {
		name          string
		createPostReq *pb.CreatePostRequest
		mocks         func()
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
			mocks: func() {
				serverMock.
					EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("port id can't be empty"))
			},
			wantErr: true,
		},
		{
			name: "create post successfully",
			createPostReq: &pb.CreatePostRequest{
				CreatePostRequest: &pb.BlogPost{
					PostId: 333,
					Title:  "test title",
					Author: "ssr",
				},
			},
			mocks: func() {
				serverMock.
					EXPECT().
					CreatePost(gomock.Any(), gomock.Any()).
					Return(&pb.CreatePostResponse{Message: "created succesfully"}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()
			err := CreatePost(context.Background(),
				serverMock,
				tt.createPostReq,
			)

			if (err != nil) != tt.wantErr {
				t.Fatal("CreatePost() error -> ", err)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serverMock := mock.NewMockBlogServiceClient(ctrl)

	tests := []struct {
		name       string
		getPostReq *pb.GetPostRequest
		mocks      func()
		wantErr    bool
	}{
		{
			name:       "post id can't be null",
			getPostReq: &pb.GetPostRequest{},
			mocks: func() {
				serverMock.
					EXPECT().
					GetPost(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("port id can't be null"))
			},
			wantErr: true,
		},
		{
			name: "get post successfully",
			getPostReq: &pb.GetPostRequest{
				PostId: 123,
			},
			mocks: func() {
				serverMock.
					EXPECT().
					GetPost(gomock.Any(), gomock.Any()).
					Return(&pb.GetPostResponse{
						GetPostResponse: &pb.BlogPost{
							PostId: 123,
							Title:  "test title",
							Author: "ssr",
						},
					}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()
			err := GetPost(context.Background(),
				serverMock,
				tt.getPostReq,
			)

			if (err != nil) != tt.wantErr {
				t.Fatal("GetPost() error -> ", err)
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serverMock := mock.NewMockBlogServiceClient(ctrl)

	tests := []struct {
		name          string
		updatePostReq *pb.UpdatePostRequest
		mocks         func()
		wantErr       bool
	}{
		{
			name:          "post id can't be null",
			updatePostReq: &pb.UpdatePostRequest{},
			mocks: func() {
				serverMock.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("port id can't be null"))
			},
			wantErr: true,
		},
		{
			name: "mask can't be empty in updatePost",
			updatePostReq: &pb.UpdatePostRequest{
				UpdatePostRequest: &pb.BlogPost{
					PostId: 333,
					Author: "srinu",
				},
			},
			mocks: func() {
				serverMock.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("mask is not empty"))
			},
			wantErr: true,
		},
		{
			name: "get post successfully",
			updatePostReq: &pb.UpdatePostRequest{
				UpdatePostRequest: &pb.BlogPost{
					PostId: 333,
					Author: "srinu",
				},
				Mask: "author",
			},
			mocks: func() {
				serverMock.
					EXPECT().
					UpdatePost(gomock.Any(), gomock.Any()).
					Return(&pb.UpdatePostResponse{
						Message: "updated successfully",
					}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()
			err := UpdatePost(context.Background(),
				serverMock,
				tt.updatePostReq,
			)

			if (err != nil) != tt.wantErr {
				t.Fatal("DeletePost() error -> ", err)
			}
		})
	}
}

func TestGetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serverMock := mock.NewMockBlogServiceClient(ctrl)

	tests := []struct {
		name          string
		deletePostReq *pb.DeletePostRequest
		mocks         func()
		wantErr       bool
	}{
		{
			name:          "post id can't be null",
			deletePostReq: &pb.DeletePostRequest{},
			mocks: func() {
				serverMock.
					EXPECT().
					DeletePost(gomock.Any(), gomock.Any()).
					Return(nil, fmt.Errorf("port id can't be null"))
			},
			wantErr: true,
		},
		{
			name: "get post successfully",
			deletePostReq: &pb.DeletePostRequest{
				PostId: 123,
			},
			mocks: func() {
				serverMock.
					EXPECT().
					DeletePost(gomock.Any(), gomock.Any()).
					Return(&pb.DeletePostResponse{
						Message: "deleted successfully",
					}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocks()
			err := DeletePost(context.Background(),
				serverMock,
				tt.deletePostReq,
			)

			if (err != nil) != tt.wantErr {
				t.Fatal("DeletePost() error -> ", err)
			}
		})
	}
}
