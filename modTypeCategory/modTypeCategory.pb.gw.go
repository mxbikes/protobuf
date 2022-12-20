// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: modTypeCategory/modTypeCategory.proto

/*
Package service_ModTypeCategory is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service_ModTypeCategory

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_ModTypeCategoryService_GetModTypeCategoryByID_0(ctx context.Context, marshaler runtime.Marshaler, client ModTypeCategoryServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoryByIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["ID"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "ID")
	}

	protoReq.ID, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "ID", err)
	}

	msg, err := client.GetModTypeCategoryByID(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ModTypeCategoryService_GetModTypeCategoryByID_0(ctx context.Context, marshaler runtime.Marshaler, server ModTypeCategoryServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoryByIDRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["ID"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "ID")
	}

	protoReq.ID, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "ID", err)
	}

	msg, err := server.GetModTypeCategoryByID(ctx, &protoReq)
	return msg, metadata, err

}

func request_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(ctx context.Context, marshaler runtime.Marshaler, client ModTypeCategoryServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoriesByModTypeIDRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetModTypeCategoriesByModTypeID(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(ctx context.Context, marshaler runtime.Marshaler, server ModTypeCategoryServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoriesByModTypeIDRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetModTypeCategoriesByModTypeID(ctx, &protoReq)
	return msg, metadata, err

}

func request_ModTypeCategoryService_GetModTypeCategories_0(ctx context.Context, marshaler runtime.Marshaler, client ModTypeCategoryServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoriesRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetModTypeCategories(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ModTypeCategoryService_GetModTypeCategories_0(ctx context.Context, marshaler runtime.Marshaler, server ModTypeCategoryServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetModTypeCategoriesRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetModTypeCategories(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterModTypeCategoryServiceHandlerServer registers the http handlers for service ModTypeCategoryService to "mux".
// UnaryRPC     :call ModTypeCategoryServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterModTypeCategoryServiceHandlerFromEndpoint instead.
func RegisterModTypeCategoryServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ModTypeCategoryServiceServer) error {

	mux.Handle("GET", pattern_ModTypeCategoryService_GetModTypeCategoryByID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategoryByID", runtime.WithHTTPPathPattern("/v1/modTypeCategory/{ID}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ModTypeCategoryService_GetModTypeCategoryByID_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategoryByID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategoriesByModTypeID", runtime.WithHTTPPathPattern("/v1/modTypeCategory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ModTypeCategoryService_GetModTypeCategories_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategories", runtime.WithHTTPPathPattern("/v1/modTypeCategory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ModTypeCategoryService_GetModTypeCategories_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategories_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterModTypeCategoryServiceHandlerFromEndpoint is same as RegisterModTypeCategoryServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterModTypeCategoryServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterModTypeCategoryServiceHandler(ctx, mux, conn)
}

// RegisterModTypeCategoryServiceHandler registers the http handlers for service ModTypeCategoryService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterModTypeCategoryServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterModTypeCategoryServiceHandlerClient(ctx, mux, NewModTypeCategoryServiceClient(conn))
}

// RegisterModTypeCategoryServiceHandlerClient registers the http handlers for service ModTypeCategoryService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ModTypeCategoryServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ModTypeCategoryServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ModTypeCategoryServiceClient" to call the correct interceptors.
func RegisterModTypeCategoryServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ModTypeCategoryServiceClient) error {

	mux.Handle("GET", pattern_ModTypeCategoryService_GetModTypeCategoryByID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategoryByID", runtime.WithHTTPPathPattern("/v1/modTypeCategory/{ID}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ModTypeCategoryService_GetModTypeCategoryByID_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategoryByID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategoriesByModTypeID", runtime.WithHTTPPathPattern("/v1/modTypeCategory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ModTypeCategoryService_GetModTypeCategories_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/modTypeCategory_service.ModTypeCategoryService/GetModTypeCategories", runtime.WithHTTPPathPattern("/v1/modTypeCategory"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ModTypeCategoryService_GetModTypeCategories_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ModTypeCategoryService_GetModTypeCategories_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ModTypeCategoryService_GetModTypeCategoryByID_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"v1", "modTypeCategory", "ID"}, ""))

	pattern_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "modTypeCategory"}, ""))

	pattern_ModTypeCategoryService_GetModTypeCategories_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "modTypeCategory"}, ""))
)

var (
	forward_ModTypeCategoryService_GetModTypeCategoryByID_0 = runtime.ForwardResponseMessage

	forward_ModTypeCategoryService_GetModTypeCategoriesByModTypeID_0 = runtime.ForwardResponseMessage

	forward_ModTypeCategoryService_GetModTypeCategories_0 = runtime.ForwardResponseMessage
)
