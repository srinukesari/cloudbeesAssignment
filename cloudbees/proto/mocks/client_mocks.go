// Code generated by MockGen. DO NOT EDIT.
// Source: ./blog_grpc.pb.go

// Package proto is a generated GoMock package.
package proto

import (
	proto "cloudbees/proto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBlogServiceClient is a mock of BlogServiceClient interface.
type MockBlogServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlogServiceClientMockRecorder
}

// MockBlogServiceClientMockRecorder is the mock recorder for MockBlogServiceClient.
type MockBlogServiceClientMockRecorder struct {
	mock *MockBlogServiceClient
}

// NewMockBlogServiceClient creates a new mock instance.
func NewMockBlogServiceClient(ctrl *gomock.Controller) *MockBlogServiceClient {
	mock := &MockBlogServiceClient{ctrl: ctrl}
	mock.recorder = &MockBlogServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogServiceClient) EXPECT() *MockBlogServiceClientMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockBlogServiceClient) CreatePost(ctx context.Context, in *proto.CreatePostRequest, opts ...grpc.CallOption) (*proto.CreatePostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePost", varargs...)
	ret0, _ := ret[0].(*proto.CreatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockBlogServiceClientMockRecorder) CreatePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockBlogServiceClient)(nil).CreatePost), varargs...)
}

// DeletePost mocks base method.
func (m *MockBlogServiceClient) DeletePost(ctx context.Context, in *proto.DeletePostRequest, opts ...grpc.CallOption) (*proto.DeletePostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePost", varargs...)
	ret0, _ := ret[0].(*proto.DeletePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockBlogServiceClientMockRecorder) DeletePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockBlogServiceClient)(nil).DeletePost), varargs...)
}

// GetPost mocks base method.
func (m *MockBlogServiceClient) GetPost(ctx context.Context, in *proto.GetPostRequest, opts ...grpc.CallOption) (*proto.GetPostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPost", varargs...)
	ret0, _ := ret[0].(*proto.GetPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockBlogServiceClientMockRecorder) GetPost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockBlogServiceClient)(nil).GetPost), varargs...)
}

// UpdatePost mocks base method.
func (m *MockBlogServiceClient) UpdatePost(ctx context.Context, in *proto.UpdatePostRequest, opts ...grpc.CallOption) (*proto.UpdatePostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePost", varargs...)
	ret0, _ := ret[0].(*proto.UpdatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockBlogServiceClientMockRecorder) UpdatePost(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockBlogServiceClient)(nil).UpdatePost), varargs...)
}

// MockBlogServiceServer is a mock of BlogServiceServer interface.
type MockBlogServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBlogServiceServerMockRecorder
}

// MockBlogServiceServerMockRecorder is the mock recorder for MockBlogServiceServer.
type MockBlogServiceServerMockRecorder struct {
	mock *MockBlogServiceServer
}

// NewMockBlogServiceServer creates a new mock instance.
func NewMockBlogServiceServer(ctrl *gomock.Controller) *MockBlogServiceServer {
	mock := &MockBlogServiceServer{ctrl: ctrl}
	mock.recorder = &MockBlogServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogServiceServer) EXPECT() *MockBlogServiceServerMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockBlogServiceServer) CreatePost(arg0 context.Context, arg1 *proto.CreatePostRequest) (*proto.CreatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockBlogServiceServerMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockBlogServiceServer)(nil).CreatePost), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockBlogServiceServer) DeletePost(arg0 context.Context, arg1 *proto.DeletePostRequest) (*proto.DeletePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeletePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockBlogServiceServerMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockBlogServiceServer)(nil).DeletePost), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockBlogServiceServer) GetPost(arg0 context.Context, arg1 *proto.GetPostRequest) (*proto.GetPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockBlogServiceServerMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockBlogServiceServer)(nil).GetPost), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockBlogServiceServer) UpdatePost(arg0 context.Context, arg1 *proto.UpdatePostRequest) (*proto.UpdatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(*proto.UpdatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockBlogServiceServerMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockBlogServiceServer)(nil).UpdatePost), arg0, arg1)
}

// mustEmbedUnimplementedBlogServiceServer mocks base method.
func (m *MockBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBlogServiceServer")
}

// mustEmbedUnimplementedBlogServiceServer indicates an expected call of mustEmbedUnimplementedBlogServiceServer.
func (mr *MockBlogServiceServerMockRecorder) mustEmbedUnimplementedBlogServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBlogServiceServer", reflect.TypeOf((*MockBlogServiceServer)(nil).mustEmbedUnimplementedBlogServiceServer))
}

// MockUnsafeBlogServiceServer is a mock of UnsafeBlogServiceServer interface.
type MockUnsafeBlogServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeBlogServiceServerMockRecorder
}

// MockUnsafeBlogServiceServerMockRecorder is the mock recorder for MockUnsafeBlogServiceServer.
type MockUnsafeBlogServiceServerMockRecorder struct {
	mock *MockUnsafeBlogServiceServer
}

// NewMockUnsafeBlogServiceServer creates a new mock instance.
func NewMockUnsafeBlogServiceServer(ctrl *gomock.Controller) *MockUnsafeBlogServiceServer {
	mock := &MockUnsafeBlogServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeBlogServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeBlogServiceServer) EXPECT() *MockUnsafeBlogServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedBlogServiceServer mocks base method.
func (m *MockUnsafeBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedBlogServiceServer")
}

// mustEmbedUnimplementedBlogServiceServer indicates an expected call of mustEmbedUnimplementedBlogServiceServer.
func (mr *MockUnsafeBlogServiceServerMockRecorder) mustEmbedUnimplementedBlogServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedBlogServiceServer", reflect.TypeOf((*MockUnsafeBlogServiceServer)(nil).mustEmbedUnimplementedBlogServiceServer))
}