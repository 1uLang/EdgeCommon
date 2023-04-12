// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_metric_stat.proto

package pb

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

// MetricStatServiceClient is the client API for MetricStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricStatServiceClient interface {
	// 上传统计数据
	UploadMetricStats(ctx context.Context, in *UploadMetricStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算指标数据数量
	CountMetricStats(ctx context.Context, in *CountMetricStatsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 读取单页指标数据
	ListMetricStats(ctx context.Context, in *ListMetricStatsRequest, opts ...grpc.CallOption) (*ListMetricStatsResponse, error)
}

type metricStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricStatServiceClient(cc grpc.ClientConnInterface) MetricStatServiceClient {
	return &metricStatServiceClient{cc}
}

func (c *metricStatServiceClient) UploadMetricStats(ctx context.Context, in *UploadMetricStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/uploadMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricStatServiceClient) CountMetricStats(ctx context.Context, in *CountMetricStatsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/countMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricStatServiceClient) ListMetricStats(ctx context.Context, in *ListMetricStatsRequest, opts ...grpc.CallOption) (*ListMetricStatsResponse, error) {
	out := new(ListMetricStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricStatService/listMetricStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricStatServiceServer is the server API for MetricStatService service.
// All implementations must embed UnimplementedMetricStatServiceServer
// for forward compatibility
type MetricStatServiceServer interface {
	// 上传统计数据
	UploadMetricStats(context.Context, *UploadMetricStatsRequest) (*RPCSuccess, error)
	// 计算指标数据数量
	CountMetricStats(context.Context, *CountMetricStatsRequest) (*RPCCountResponse, error)
	// 读取单页指标数据
	ListMetricStats(context.Context, *ListMetricStatsRequest) (*ListMetricStatsResponse, error)
}

// UnimplementedMetricStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricStatServiceServer struct {
}

func (UnimplementedMetricStatServiceServer) UploadMetricStats(context.Context, *UploadMetricStatsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadMetricStats not implemented")
}
func (UnimplementedMetricStatServiceServer) CountMetricStats(context.Context, *CountMetricStatsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountMetricStats not implemented")
}
func (UnimplementedMetricStatServiceServer) ListMetricStats(context.Context, *ListMetricStatsRequest) (*ListMetricStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetricStats not implemented")
}
func RegisterMetricStatServiceServer(s grpc.ServiceRegistrar, srv MetricStatServiceServer) {
	s.RegisterService(&MetricStatService_ServiceDesc, srv)
}

func _MetricStatService_UploadMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).UploadMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/uploadMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).UploadMetricStats(ctx, req.(*UploadMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricStatService_CountMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).CountMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/countMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).CountMetricStats(ctx, req.(*CountMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricStatService_ListMetricStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricStatServiceServer).ListMetricStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricStatService/listMetricStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricStatServiceServer).ListMetricStats(ctx, req.(*ListMetricStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricStatService_ServiceDesc is the grpc.ServiceDesc for MetricStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MetricStatService",
	HandlerType: (*MetricStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "uploadMetricStats",
			Handler:    _MetricStatService_UploadMetricStats_Handler,
		},
		{
			MethodName: "countMetricStats",
			Handler:    _MetricStatService_CountMetricStats_Handler,
		},
		{
			MethodName: "listMetricStats",
			Handler:    _MetricStatService_ListMetricStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_metric_stat.proto",
}