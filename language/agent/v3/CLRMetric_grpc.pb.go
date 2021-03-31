// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking/network/common/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CLRMetricReportServiceClient is the client API for CLRMetricReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CLRMetricReportServiceClient interface {
	Collect(ctx context.Context, in *CLRMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error)
}

type cLRMetricReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCLRMetricReportServiceClient(cc grpc.ClientConnInterface) CLRMetricReportServiceClient {
	return &cLRMetricReportServiceClient{cc}
}

func (c *cLRMetricReportServiceClient) Collect(ctx context.Context, in *CLRMetricCollection, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v3.CLRMetricReportService/collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CLRMetricReportServiceServer is the server API for CLRMetricReportService service.
// All implementations must embed UnimplementedCLRMetricReportServiceServer
// for forward compatibility
type CLRMetricReportServiceServer interface {
	Collect(context.Context, *CLRMetricCollection) (*v3.Commands, error)
	mustEmbedUnimplementedCLRMetricReportServiceServer()
}

// UnimplementedCLRMetricReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCLRMetricReportServiceServer struct {
}

func (UnimplementedCLRMetricReportServiceServer) Collect(context.Context, *CLRMetricCollection) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedCLRMetricReportServiceServer) mustEmbedUnimplementedCLRMetricReportServiceServer() {
}

// UnsafeCLRMetricReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CLRMetricReportServiceServer will
// result in compilation errors.
type UnsafeCLRMetricReportServiceServer interface {
	mustEmbedUnimplementedCLRMetricReportServiceServer()
}

func RegisterCLRMetricReportServiceServer(s grpc.ServiceRegistrar, srv CLRMetricReportServiceServer) {
	s.RegisterService(&CLRMetricReportService_ServiceDesc, srv)
}

func _CLRMetricReportService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CLRMetricCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CLRMetricReportServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.CLRMetricReportService/collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CLRMetricReportServiceServer).Collect(ctx, req.(*CLRMetricCollection))
	}
	return interceptor(ctx, in, info, handler)
}

// CLRMetricReportService_ServiceDesc is the grpc.ServiceDesc for CLRMetricReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CLRMetricReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.CLRMetricReportService",
	HandlerType: (*CLRMetricReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "collect",
			Handler:    _CLRMetricReportService_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/CLRMetric.proto",
}
