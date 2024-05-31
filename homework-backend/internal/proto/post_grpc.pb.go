// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: post.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	PostCreate(ctx context.Context, in *PostCreateRequest, opts ...grpc.CallOption) (*PostCreateResponse, error)
	PostUpdate(ctx context.Context, in *PostUpdateRequest, opts ...grpc.CallOption) (*PostMainResponse, error)
	PostDelete(ctx context.Context, in *PostMainRequest, opts ...grpc.CallOption) (*PostMainResponse, error)
	PostGet(ctx context.Context, in *PostMainRequest, opts ...grpc.CallOption) (*PostGetResponse, error)
	PostFeed(ctx context.Context, in *PostFeedRequest, opts ...grpc.CallOption) (*PostFeedResponse, error)
	FriendSet(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error)
	FriendDelete(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) PostCreate(ctx context.Context, in *PostCreateRequest, opts ...grpc.CallOption) (*PostCreateResponse, error) {
	out := new(PostCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/PostCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostUpdate(ctx context.Context, in *PostUpdateRequest, opts ...grpc.CallOption) (*PostMainResponse, error) {
	out := new(PostMainResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/PostUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostDelete(ctx context.Context, in *PostMainRequest, opts ...grpc.CallOption) (*PostMainResponse, error) {
	out := new(PostMainResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/PostDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostGet(ctx context.Context, in *PostMainRequest, opts ...grpc.CallOption) (*PostGetResponse, error) {
	out := new(PostGetResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/PostGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) PostFeed(ctx context.Context, in *PostFeedRequest, opts ...grpc.CallOption) (*PostFeedResponse, error) {
	out := new(PostFeedResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/PostFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) FriendSet(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error) {
	out := new(FriendResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/FriendSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) FriendDelete(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error) {
	out := new(FriendResponse)
	err := c.cc.Invoke(ctx, "/proto.PostService/FriendDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	PostCreate(context.Context, *PostCreateRequest) (*PostCreateResponse, error)
	PostUpdate(context.Context, *PostUpdateRequest) (*PostMainResponse, error)
	PostDelete(context.Context, *PostMainRequest) (*PostMainResponse, error)
	PostGet(context.Context, *PostMainRequest) (*PostGetResponse, error)
	PostFeed(context.Context, *PostFeedRequest) (*PostFeedResponse, error)
	FriendSet(context.Context, *FriendRequest) (*FriendResponse, error)
	FriendDelete(context.Context, *FriendRequest) (*FriendResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) PostCreate(context.Context, *PostCreateRequest) (*PostCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCreate not implemented")
}
func (UnimplementedPostServiceServer) PostUpdate(context.Context, *PostUpdateRequest) (*PostMainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUpdate not implemented")
}
func (UnimplementedPostServiceServer) PostDelete(context.Context, *PostMainRequest) (*PostMainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDelete not implemented")
}
func (UnimplementedPostServiceServer) PostGet(context.Context, *PostMainRequest) (*PostGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGet not implemented")
}
func (UnimplementedPostServiceServer) PostFeed(context.Context, *PostFeedRequest) (*PostFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostFeed not implemented")
}
func (UnimplementedPostServiceServer) FriendSet(context.Context, *FriendRequest) (*FriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendSet not implemented")
}
func (UnimplementedPostServiceServer) FriendDelete(context.Context, *FriendRequest) (*FriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendDelete not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_PostCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/PostCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostCreate(ctx, req.(*PostCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/PostUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostUpdate(ctx, req.(*PostUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/PostDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostDelete(ctx, req.(*PostMainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/PostGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostGet(ctx, req.(*PostMainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_PostFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).PostFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/PostFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).PostFeed(ctx, req.(*PostFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_FriendSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).FriendSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/FriendSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).FriendSet(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_FriendDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).FriendDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PostService/FriendDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).FriendDelete(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCreate",
			Handler:    _PostService_PostCreate_Handler,
		},
		{
			MethodName: "PostUpdate",
			Handler:    _PostService_PostUpdate_Handler,
		},
		{
			MethodName: "PostDelete",
			Handler:    _PostService_PostDelete_Handler,
		},
		{
			MethodName: "PostGet",
			Handler:    _PostService_PostGet_Handler,
		},
		{
			MethodName: "PostFeed",
			Handler:    _PostService_PostFeed_Handler,
		},
		{
			MethodName: "FriendSet",
			Handler:    _PostService_FriendSet_Handler,
		},
		{
			MethodName: "FriendDelete",
			Handler:    _PostService_FriendDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}