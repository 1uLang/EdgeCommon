// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_server_daily_stat.proto

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

// ServerDailyStatServiceClient is the client API for ServerDailyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerDailyStatServiceClient interface {
	// 上传统计
	UploadServerDailyStats(ctx context.Context, in *UploadServerDailyStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 按小时读取统计数据
	FindLatestServerHourlyStats(ctx context.Context, in *FindLatestServerHourlyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerHourlyStatsResponse, error)
	// 按分钟读取统计数据，并返回秒级平均流量
	FindLatestServerMinutelyStats(ctx context.Context, in *FindLatestServerMinutelyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerMinutelyStatsResponse, error)
	// 读取某天的5分钟间隔流量
	FindServer5MinutelyStatsWithDay(ctx context.Context, in *FindServer5MinutelyStatsWithDayRequest, opts ...grpc.CallOption) (*FindServer5MinutelyStatsWithDayResponse, error)
	// 按日读取统计数据
	FindLatestServerDailyStats(ctx context.Context, in *FindLatestServerDailyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerDailyStatsResponse, error)
	// 查找单个服务当前时刻（N分钟内）统计数据
	SumCurrentServerDailyStats(ctx context.Context, in *SumCurrentServerDailyStatsRequest, opts ...grpc.CallOption) (*SumCurrentServerDailyStatsResponse, error)
	// 计算单个服务的日统计
	SumServerDailyStats(ctx context.Context, in *SumServerDailyStatsRequest, opts ...grpc.CallOption) (*SumServerDailyStatsResponse, error)
	// 计算单个服务的月统计
	SumServerMonthlyStats(ctx context.Context, in *SumServerMonthlyStatsRequest, opts ...grpc.CallOption) (*SumServerMonthlyStatsResponse, error)
}

type serverDailyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerDailyStatServiceClient(cc grpc.ClientConnInterface) ServerDailyStatServiceClient {
	return &serverDailyStatServiceClient{cc}
}

func (c *serverDailyStatServiceClient) UploadServerDailyStats(ctx context.Context, in *UploadServerDailyStatsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/uploadServerDailyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) FindLatestServerHourlyStats(ctx context.Context, in *FindLatestServerHourlyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerHourlyStatsResponse, error) {
	out := new(FindLatestServerHourlyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/findLatestServerHourlyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) FindLatestServerMinutelyStats(ctx context.Context, in *FindLatestServerMinutelyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerMinutelyStatsResponse, error) {
	out := new(FindLatestServerMinutelyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/findLatestServerMinutelyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) FindServer5MinutelyStatsWithDay(ctx context.Context, in *FindServer5MinutelyStatsWithDayRequest, opts ...grpc.CallOption) (*FindServer5MinutelyStatsWithDayResponse, error) {
	out := new(FindServer5MinutelyStatsWithDayResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/findServer5MinutelyStatsWithDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) FindLatestServerDailyStats(ctx context.Context, in *FindLatestServerDailyStatsRequest, opts ...grpc.CallOption) (*FindLatestServerDailyStatsResponse, error) {
	out := new(FindLatestServerDailyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/findLatestServerDailyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) SumCurrentServerDailyStats(ctx context.Context, in *SumCurrentServerDailyStatsRequest, opts ...grpc.CallOption) (*SumCurrentServerDailyStatsResponse, error) {
	out := new(SumCurrentServerDailyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/sumCurrentServerDailyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) SumServerDailyStats(ctx context.Context, in *SumServerDailyStatsRequest, opts ...grpc.CallOption) (*SumServerDailyStatsResponse, error) {
	out := new(SumServerDailyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/sumServerDailyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverDailyStatServiceClient) SumServerMonthlyStats(ctx context.Context, in *SumServerMonthlyStatsRequest, opts ...grpc.CallOption) (*SumServerMonthlyStatsResponse, error) {
	out := new(SumServerMonthlyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerDailyStatService/sumServerMonthlyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerDailyStatServiceServer is the server API for ServerDailyStatService service.
// All implementations must embed UnimplementedServerDailyStatServiceServer
// for forward compatibility
type ServerDailyStatServiceServer interface {
	// 上传统计
	UploadServerDailyStats(context.Context, *UploadServerDailyStatsRequest) (*RPCSuccess, error)
	// 按小时读取统计数据
	FindLatestServerHourlyStats(context.Context, *FindLatestServerHourlyStatsRequest) (*FindLatestServerHourlyStatsResponse, error)
	// 按分钟读取统计数据，并返回秒级平均流量
	FindLatestServerMinutelyStats(context.Context, *FindLatestServerMinutelyStatsRequest) (*FindLatestServerMinutelyStatsResponse, error)
	// 读取某天的5分钟间隔流量
	FindServer5MinutelyStatsWithDay(context.Context, *FindServer5MinutelyStatsWithDayRequest) (*FindServer5MinutelyStatsWithDayResponse, error)
	// 按日读取统计数据
	FindLatestServerDailyStats(context.Context, *FindLatestServerDailyStatsRequest) (*FindLatestServerDailyStatsResponse, error)
	// 查找单个服务当前时刻（N分钟内）统计数据
	SumCurrentServerDailyStats(context.Context, *SumCurrentServerDailyStatsRequest) (*SumCurrentServerDailyStatsResponse, error)
	// 计算单个服务的日统计
	SumServerDailyStats(context.Context, *SumServerDailyStatsRequest) (*SumServerDailyStatsResponse, error)
	// 计算单个服务的月统计
	SumServerMonthlyStats(context.Context, *SumServerMonthlyStatsRequest) (*SumServerMonthlyStatsResponse, error)
}

// UnimplementedServerDailyStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerDailyStatServiceServer struct {
}

func (UnimplementedServerDailyStatServiceServer) UploadServerDailyStats(context.Context, *UploadServerDailyStatsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadServerDailyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) FindLatestServerHourlyStats(context.Context, *FindLatestServerHourlyStatsRequest) (*FindLatestServerHourlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestServerHourlyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) FindLatestServerMinutelyStats(context.Context, *FindLatestServerMinutelyStatsRequest) (*FindLatestServerMinutelyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestServerMinutelyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) FindServer5MinutelyStatsWithDay(context.Context, *FindServer5MinutelyStatsWithDayRequest) (*FindServer5MinutelyStatsWithDayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindServer5MinutelyStatsWithDay not implemented")
}
func (UnimplementedServerDailyStatServiceServer) FindLatestServerDailyStats(context.Context, *FindLatestServerDailyStatsRequest) (*FindLatestServerDailyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestServerDailyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) SumCurrentServerDailyStats(context.Context, *SumCurrentServerDailyStatsRequest) (*SumCurrentServerDailyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumCurrentServerDailyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) SumServerDailyStats(context.Context, *SumServerDailyStatsRequest) (*SumServerDailyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumServerDailyStats not implemented")
}
func (UnimplementedServerDailyStatServiceServer) SumServerMonthlyStats(context.Context, *SumServerMonthlyStatsRequest) (*SumServerMonthlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumServerMonthlyStats not implemented")
}

func RegisterServerDailyStatServiceServer(s grpc.ServiceRegistrar, srv ServerDailyStatServiceServer) {
	s.RegisterService(&ServerDailyStatService_ServiceDesc, srv)
}

func _ServerDailyStatService_UploadServerDailyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadServerDailyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).UploadServerDailyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/uploadServerDailyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).UploadServerDailyStats(ctx, req.(*UploadServerDailyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_FindLatestServerHourlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestServerHourlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).FindLatestServerHourlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/findLatestServerHourlyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).FindLatestServerHourlyStats(ctx, req.(*FindLatestServerHourlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_FindLatestServerMinutelyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestServerMinutelyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).FindLatestServerMinutelyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/findLatestServerMinutelyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).FindLatestServerMinutelyStats(ctx, req.(*FindLatestServerMinutelyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_FindServer5MinutelyStatsWithDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindServer5MinutelyStatsWithDayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).FindServer5MinutelyStatsWithDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/findServer5MinutelyStatsWithDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).FindServer5MinutelyStatsWithDay(ctx, req.(*FindServer5MinutelyStatsWithDayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_FindLatestServerDailyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestServerDailyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).FindLatestServerDailyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/findLatestServerDailyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).FindLatestServerDailyStats(ctx, req.(*FindLatestServerDailyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_SumCurrentServerDailyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumCurrentServerDailyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).SumCurrentServerDailyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/sumCurrentServerDailyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).SumCurrentServerDailyStats(ctx, req.(*SumCurrentServerDailyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_SumServerDailyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumServerDailyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).SumServerDailyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/sumServerDailyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).SumServerDailyStats(ctx, req.(*SumServerDailyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerDailyStatService_SumServerMonthlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumServerMonthlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerDailyStatServiceServer).SumServerMonthlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerDailyStatService/sumServerMonthlyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerDailyStatServiceServer).SumServerMonthlyStats(ctx, req.(*SumServerMonthlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerDailyStatService_ServiceDesc is the grpc.ServiceDesc for ServerDailyStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerDailyStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerDailyStatService",
	HandlerType: (*ServerDailyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "uploadServerDailyStats",
			Handler:    _ServerDailyStatService_UploadServerDailyStats_Handler,
		},
		{
			MethodName: "findLatestServerHourlyStats",
			Handler:    _ServerDailyStatService_FindLatestServerHourlyStats_Handler,
		},
		{
			MethodName: "findLatestServerMinutelyStats",
			Handler:    _ServerDailyStatService_FindLatestServerMinutelyStats_Handler,
		},
		{
			MethodName: "findServer5MinutelyStatsWithDay",
			Handler:    _ServerDailyStatService_FindServer5MinutelyStatsWithDay_Handler,
		},
		{
			MethodName: "findLatestServerDailyStats",
			Handler:    _ServerDailyStatService_FindLatestServerDailyStats_Handler,
		},
		{
			MethodName: "sumCurrentServerDailyStats",
			Handler:    _ServerDailyStatService_SumCurrentServerDailyStats_Handler,
		},
		{
			MethodName: "sumServerDailyStats",
			Handler:    _ServerDailyStatService_SumServerDailyStats_Handler,
		},
		{
			MethodName: "sumServerMonthlyStats",
			Handler:    _ServerDailyStatService_SumServerMonthlyStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_daily_stat.proto",
}
