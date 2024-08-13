// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/v1/administration_usage_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

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

func request_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, client AdministrationUsageServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetCurrentSecuredUnitsUsage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, server AdministrationUsageServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetCurrentSecuredUnitsUsage(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_AdministrationUsageService_GetMaxSecuredUnitsUsage_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, client AdministrationUsageServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TimeRange
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_AdministrationUsageService_GetMaxSecuredUnitsUsage_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetMaxSecuredUnitsUsage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(ctx context.Context, marshaler runtime.Marshaler, server AdministrationUsageServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TimeRange
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_AdministrationUsageService_GetMaxSecuredUnitsUsage_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetMaxSecuredUnitsUsage(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterAdministrationUsageServiceHandlerServer registers the http handlers for service AdministrationUsageService to "mux".
// UnaryRPC     :call AdministrationUsageServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterAdministrationUsageServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterAdministrationUsageServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server AdministrationUsageServiceServer) error {

	mux.Handle("GET", pattern_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/v1.AdministrationUsageService/GetCurrentSecuredUnitsUsage", runtime.WithHTTPPathPattern("/v1/administration/usage/secured-units/current"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_AdministrationUsageService_GetMaxSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/v1.AdministrationUsageService/GetMaxSecuredUnitsUsage", runtime.WithHTTPPathPattern("/v1/administration/usage/secured-units/max"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterAdministrationUsageServiceHandlerFromEndpoint is same as RegisterAdministrationUsageServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAdministrationUsageServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAdministrationUsageServiceHandler(ctx, mux, conn)
}

// RegisterAdministrationUsageServiceHandler registers the http handlers for service AdministrationUsageService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAdministrationUsageServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAdministrationUsageServiceHandlerClient(ctx, mux, NewAdministrationUsageServiceClient(conn))
}

// RegisterAdministrationUsageServiceHandlerClient registers the http handlers for service AdministrationUsageService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AdministrationUsageServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AdministrationUsageServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AdministrationUsageServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterAdministrationUsageServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AdministrationUsageServiceClient) error {

	mux.Handle("GET", pattern_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/v1.AdministrationUsageService/GetCurrentSecuredUnitsUsage", runtime.WithHTTPPathPattern("/v1/administration/usage/secured-units/current"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_AdministrationUsageService_GetMaxSecuredUnitsUsage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/v1.AdministrationUsageService/GetMaxSecuredUnitsUsage", runtime.WithHTTPPathPattern("/v1/administration/usage/secured-units/max"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AdministrationUsageService_GetMaxSecuredUnitsUsage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"v1", "administration", "usage", "secured-units", "current"}, ""))

	pattern_AdministrationUsageService_GetMaxSecuredUnitsUsage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"v1", "administration", "usage", "secured-units", "max"}, ""))
)

var (
	forward_AdministrationUsageService_GetCurrentSecuredUnitsUsage_0 = runtime.ForwardResponseMessage

	forward_AdministrationUsageService_GetMaxSecuredUnitsUsage_0 = runtime.ForwardResponseMessage
)
