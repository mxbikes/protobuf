// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: subscription/subscription.proto

package Subscription

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

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	GetSubscriptionsByUserID(ctx context.Context, in *GetSubscriptionsByIDRequest, opts ...grpc.CallOption) (*GetSubscriptionsByIDResponse, error)
	AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*AddSubscriptionResponse, error)
	RemoveSubscription(ctx context.Context, in *RemoveSubscriptionRequest, opts ...grpc.CallOption) (*RemoveSubscriptionResponse, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) GetSubscriptionsByUserID(ctx context.Context, in *GetSubscriptionsByIDRequest, opts ...grpc.CallOption) (*GetSubscriptionsByIDResponse, error) {
	out := new(GetSubscriptionsByIDResponse)
	err := c.cc.Invoke(ctx, "/subscription_service.SubscriptionService/GetSubscriptionsByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*AddSubscriptionResponse, error) {
	out := new(AddSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/subscription_service.SubscriptionService/AddSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) RemoveSubscription(ctx context.Context, in *RemoveSubscriptionRequest, opts ...grpc.CallOption) (*RemoveSubscriptionResponse, error) {
	out := new(RemoveSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/subscription_service.SubscriptionService/RemoveSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	GetSubscriptionsByUserID(context.Context, *GetSubscriptionsByIDRequest) (*GetSubscriptionsByIDResponse, error)
	AddSubscription(context.Context, *AddSubscriptionRequest) (*AddSubscriptionResponse, error)
	RemoveSubscription(context.Context, *RemoveSubscriptionRequest) (*RemoveSubscriptionResponse, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) GetSubscriptionsByUserID(context.Context, *GetSubscriptionsByIDRequest) (*GetSubscriptionsByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionsByUserID not implemented")
}
func (UnimplementedSubscriptionServiceServer) AddSubscription(context.Context, *AddSubscriptionRequest) (*AddSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) RemoveSubscription(context.Context, *RemoveSubscriptionRequest) (*RemoveSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_GetSubscriptionsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionsByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscriptionsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription_service.SubscriptionService/GetSubscriptionsByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscriptionsByUserID(ctx, req.(*GetSubscriptionsByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription_service.SubscriptionService/AddSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).AddSubscription(ctx, req.(*AddSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_RemoveSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).RemoveSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subscription_service.SubscriptionService/RemoveSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).RemoveSubscription(ctx, req.(*RemoveSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription_service.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriptionsByUserID",
			Handler:    _SubscriptionService_GetSubscriptionsByUserID_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _SubscriptionService_AddSubscription_Handler,
		},
		{
			MethodName: "RemoveSubscription",
			Handler:    _SubscriptionService_RemoveSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription/subscription.proto",
}
