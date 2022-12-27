// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: modImage/modImage.proto

package modImage

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

// ModImageServiceClient is the client API for ModImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModImageServiceClient interface {
	GetModImagesByModID(ctx context.Context, in *GetModImagesByModIDRequest, opts ...grpc.CallOption) (*GetModImagesByModIDResponse, error)
}

type modImageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModImageServiceClient(cc grpc.ClientConnInterface) ModImageServiceClient {
	return &modImageServiceClient{cc}
}

func (c *modImageServiceClient) GetModImagesByModID(ctx context.Context, in *GetModImagesByModIDRequest, opts ...grpc.CallOption) (*GetModImagesByModIDResponse, error) {
	out := new(GetModImagesByModIDResponse)
	err := c.cc.Invoke(ctx, "/modImage_service.ModImageService/GetModImagesByModID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModImageServiceServer is the server API for ModImageService service.
// All implementations must embed UnimplementedModImageServiceServer
// for forward compatibility
type ModImageServiceServer interface {
	GetModImagesByModID(context.Context, *GetModImagesByModIDRequest) (*GetModImagesByModIDResponse, error)
	mustEmbedUnimplementedModImageServiceServer()
}

// UnimplementedModImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModImageServiceServer struct {
}

func (UnimplementedModImageServiceServer) GetModImagesByModID(context.Context, *GetModImagesByModIDRequest) (*GetModImagesByModIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModImagesByModID not implemented")
}
func (UnimplementedModImageServiceServer) mustEmbedUnimplementedModImageServiceServer() {}

// UnsafeModImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModImageServiceServer will
// result in compilation errors.
type UnsafeModImageServiceServer interface {
	mustEmbedUnimplementedModImageServiceServer()
}

func RegisterModImageServiceServer(s grpc.ServiceRegistrar, srv ModImageServiceServer) {
	s.RegisterService(&ModImageService_ServiceDesc, srv)
}

func _ModImageService_GetModImagesByModID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModImagesByModIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModImageServiceServer).GetModImagesByModID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modImage_service.ModImageService/GetModImagesByModID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModImageServiceServer).GetModImagesByModID(ctx, req.(*GetModImagesByModIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModImageService_ServiceDesc is the grpc.ServiceDesc for ModImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modImage_service.ModImageService",
	HandlerType: (*ModImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModImagesByModID",
			Handler:    _ModImageService_GetModImagesByModID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modImage/modImage.proto",
}
