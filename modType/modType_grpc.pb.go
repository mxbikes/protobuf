// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: modType/modType.proto

package ModType

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

// ModTypeServiceClient is the client API for ModTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModTypeServiceClient interface {
	GetModTypeByID(ctx context.Context, in *GetModTypeByIDRequest, opts ...grpc.CallOption) (*GetModTypeByIDResponse, error)
	GetModTypes(ctx context.Context, in *GetModTypesRequest, opts ...grpc.CallOption) (*GetModTypesResponse, error)
}

type modTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModTypeServiceClient(cc grpc.ClientConnInterface) ModTypeServiceClient {
	return &modTypeServiceClient{cc}
}

func (c *modTypeServiceClient) GetModTypeByID(ctx context.Context, in *GetModTypeByIDRequest, opts ...grpc.CallOption) (*GetModTypeByIDResponse, error) {
	out := new(GetModTypeByIDResponse)
	err := c.cc.Invoke(ctx, "/modType_service.ModTypeService/GetModTypeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modTypeServiceClient) GetModTypes(ctx context.Context, in *GetModTypesRequest, opts ...grpc.CallOption) (*GetModTypesResponse, error) {
	out := new(GetModTypesResponse)
	err := c.cc.Invoke(ctx, "/modType_service.ModTypeService/GetModTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModTypeServiceServer is the server API for ModTypeService service.
// All implementations must embed UnimplementedModTypeServiceServer
// for forward compatibility
type ModTypeServiceServer interface {
	GetModTypeByID(context.Context, *GetModTypeByIDRequest) (*GetModTypeByIDResponse, error)
	GetModTypes(context.Context, *GetModTypesRequest) (*GetModTypesResponse, error)
	mustEmbedUnimplementedModTypeServiceServer()
}

// UnimplementedModTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModTypeServiceServer struct {
}

func (UnimplementedModTypeServiceServer) GetModTypeByID(context.Context, *GetModTypeByIDRequest) (*GetModTypeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModTypeByID not implemented")
}
func (UnimplementedModTypeServiceServer) GetModTypes(context.Context, *GetModTypesRequest) (*GetModTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModTypes not implemented")
}
func (UnimplementedModTypeServiceServer) mustEmbedUnimplementedModTypeServiceServer() {}

// UnsafeModTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModTypeServiceServer will
// result in compilation errors.
type UnsafeModTypeServiceServer interface {
	mustEmbedUnimplementedModTypeServiceServer()
}

func RegisterModTypeServiceServer(s grpc.ServiceRegistrar, srv ModTypeServiceServer) {
	s.RegisterService(&ModTypeService_ServiceDesc, srv)
}

func _ModTypeService_GetModTypeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModTypeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModTypeServiceServer).GetModTypeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modType_service.ModTypeService/GetModTypeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModTypeServiceServer).GetModTypeByID(ctx, req.(*GetModTypeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModTypeService_GetModTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModTypeServiceServer).GetModTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/modType_service.ModTypeService/GetModTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModTypeServiceServer).GetModTypes(ctx, req.(*GetModTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModTypeService_ServiceDesc is the grpc.ServiceDesc for ModTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modType_service.ModTypeService",
	HandlerType: (*ModTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModTypeByID",
			Handler:    _ModTypeService_GetModTypeByID_Handler,
		},
		{
			MethodName: "GetModTypes",
			Handler:    _ModTypeService_GetModTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modType/modType.proto",
}
