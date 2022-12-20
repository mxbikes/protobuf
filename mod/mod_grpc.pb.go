// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: mod/mod.proto

package service_Mod

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

// ModServiceClient is the client API for ModService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModServiceClient interface {
	SearchMod(ctx context.Context, in *SearchModRequest, opts ...grpc.CallOption) (*SearchModResponse, error)
	GetModByID(ctx context.Context, in *GetModByIDRequest, opts ...grpc.CallOption) (*GetModByIDResponse, error)
}

type modServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModServiceClient(cc grpc.ClientConnInterface) ModServiceClient {
	return &modServiceClient{cc}
}

func (c *modServiceClient) SearchMod(ctx context.Context, in *SearchModRequest, opts ...grpc.CallOption) (*SearchModResponse, error) {
	out := new(SearchModResponse)
	err := c.cc.Invoke(ctx, "/mod_service.ModService/SearchMod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modServiceClient) GetModByID(ctx context.Context, in *GetModByIDRequest, opts ...grpc.CallOption) (*GetModByIDResponse, error) {
	out := new(GetModByIDResponse)
	err := c.cc.Invoke(ctx, "/mod_service.ModService/GetModByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModServiceServer is the server API for ModService service.
// All implementations must embed UnimplementedModServiceServer
// for forward compatibility
type ModServiceServer interface {
	SearchMod(context.Context, *SearchModRequest) (*SearchModResponse, error)
	GetModByID(context.Context, *GetModByIDRequest) (*GetModByIDResponse, error)
	mustEmbedUnimplementedModServiceServer()
}

// UnimplementedModServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModServiceServer struct {
}

func (UnimplementedModServiceServer) SearchMod(context.Context, *SearchModRequest) (*SearchModResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMod not implemented")
}
func (UnimplementedModServiceServer) GetModByID(context.Context, *GetModByIDRequest) (*GetModByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModByID not implemented")
}
func (UnimplementedModServiceServer) mustEmbedUnimplementedModServiceServer() {}

// UnsafeModServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModServiceServer will
// result in compilation errors.
type UnsafeModServiceServer interface {
	mustEmbedUnimplementedModServiceServer()
}

func RegisterModServiceServer(s grpc.ServiceRegistrar, srv ModServiceServer) {
	s.RegisterService(&ModService_ServiceDesc, srv)
}

func _ModService_SearchMod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchModRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModServiceServer).SearchMod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mod_service.ModService/SearchMod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModServiceServer).SearchMod(ctx, req.(*SearchModRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModService_GetModByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModServiceServer).GetModByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mod_service.ModService/GetModByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModServiceServer).GetModByID(ctx, req.(*GetModByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModService_ServiceDesc is the grpc.ServiceDesc for ModService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mod_service.ModService",
	HandlerType: (*ModServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMod",
			Handler:    _ModService_SearchMod_Handler,
		},
		{
			MethodName: "GetModByID",
			Handler:    _ModService_GetModByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mod/mod.proto",
}
