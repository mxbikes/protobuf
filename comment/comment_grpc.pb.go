// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: comment/comment.proto

package comment

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

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetCommentByModID(ctx context.Context, in *GetCommentByModIDRequest, opts ...grpc.CallOption) (*GetCommentByModIDResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetCommentByModID(ctx context.Context, in *GetCommentByModIDRequest, opts ...grpc.CallOption) (*GetCommentByModIDResponse, error) {
	out := new(GetCommentByModIDResponse)
	err := c.cc.Invoke(ctx, "/comment_service.CommentService/GetCommentByModID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment_service.CommentService/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/comment_service.CommentService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment_service.CommentService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	GetCommentByModID(context.Context, *GetCommentByModIDRequest) (*GetCommentByModIDResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) GetCommentByModID(context.Context, *GetCommentByModIDRequest) (*GetCommentByModIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentByModID not implemented")
}
func (UnimplementedCommentServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_GetCommentByModID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentByModIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetCommentByModID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_service.CommentService/GetCommentByModID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetCommentByModID(ctx, req.(*GetCommentByModIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_service.CommentService/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_service.CommentService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment_service.CommentService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment_service.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommentByModID",
			Handler:    _CommentService_GetCommentByModID_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _CommentService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentService_DeleteComment_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment/comment.proto",
}
