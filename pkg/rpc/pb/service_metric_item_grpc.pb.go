// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_metric_item.proto

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

// MetricItemServiceClient is the client API for MetricItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricItemServiceClient interface {
	// 创建指标
	CreateMetricItem(ctx context.Context, in *CreateMetricItemRequest, opts ...grpc.CallOption) (*CreateMetricItemResponse, error)
	// 修改指标
	UpdateMetricItem(ctx context.Context, in *UpdateMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个指标信息
	FindEnabledMetricItem(ctx context.Context, in *FindEnabledMetricItemRequest, opts ...grpc.CallOption) (*FindEnabledMetricItemResponse, error)
	// 计算指标数量
	CountAllEnabledMetricItems(ctx context.Context, in *CountAllEnabledMetricItemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页指标
	ListEnabledMetricItems(ctx context.Context, in *ListEnabledMetricItemsRequest, opts ...grpc.CallOption) (*ListEnabledMetricItemsResponse, error)
	// 删除指标
	DeleteMetricItem(ctx context.Context, in *DeleteMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type metricItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricItemServiceClient(cc grpc.ClientConnInterface) MetricItemServiceClient {
	return &metricItemServiceClient{cc}
}

func (c *metricItemServiceClient) CreateMetricItem(ctx context.Context, in *CreateMetricItemRequest, opts ...grpc.CallOption) (*CreateMetricItemResponse, error) {
	out := new(CreateMetricItemResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/createMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricItemServiceClient) UpdateMetricItem(ctx context.Context, in *UpdateMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/updateMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricItemServiceClient) FindEnabledMetricItem(ctx context.Context, in *FindEnabledMetricItemRequest, opts ...grpc.CallOption) (*FindEnabledMetricItemResponse, error) {
	out := new(FindEnabledMetricItemResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/findEnabledMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricItemServiceClient) CountAllEnabledMetricItems(ctx context.Context, in *CountAllEnabledMetricItemsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/countAllEnabledMetricItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricItemServiceClient) ListEnabledMetricItems(ctx context.Context, in *ListEnabledMetricItemsRequest, opts ...grpc.CallOption) (*ListEnabledMetricItemsResponse, error) {
	out := new(ListEnabledMetricItemsResponse)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/listEnabledMetricItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricItemServiceClient) DeleteMetricItem(ctx context.Context, in *DeleteMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MetricItemService/deleteMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricItemServiceServer is the server API for MetricItemService service.
// All implementations must embed UnimplementedMetricItemServiceServer
// for forward compatibility
type MetricItemServiceServer interface {
	// 创建指标
	CreateMetricItem(context.Context, *CreateMetricItemRequest) (*CreateMetricItemResponse, error)
	// 修改指标
	UpdateMetricItem(context.Context, *UpdateMetricItemRequest) (*RPCSuccess, error)
	// 查找单个指标信息
	FindEnabledMetricItem(context.Context, *FindEnabledMetricItemRequest) (*FindEnabledMetricItemResponse, error)
	// 计算指标数量
	CountAllEnabledMetricItems(context.Context, *CountAllEnabledMetricItemsRequest) (*RPCCountResponse, error)
	// 列出单页指标
	ListEnabledMetricItems(context.Context, *ListEnabledMetricItemsRequest) (*ListEnabledMetricItemsResponse, error)
	// 删除指标
	DeleteMetricItem(context.Context, *DeleteMetricItemRequest) (*RPCSuccess, error)
}

// UnimplementedMetricItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricItemServiceServer struct {
}

func (UnimplementedMetricItemServiceServer) CreateMetricItem(context.Context, *CreateMetricItemRequest) (*CreateMetricItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMetricItem not implemented")
}
func (UnimplementedMetricItemServiceServer) UpdateMetricItem(context.Context, *UpdateMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetricItem not implemented")
}
func (UnimplementedMetricItemServiceServer) FindEnabledMetricItem(context.Context, *FindEnabledMetricItemRequest) (*FindEnabledMetricItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledMetricItem not implemented")
}
func (UnimplementedMetricItemServiceServer) CountAllEnabledMetricItems(context.Context, *CountAllEnabledMetricItemsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledMetricItems not implemented")
}
func (UnimplementedMetricItemServiceServer) ListEnabledMetricItems(context.Context, *ListEnabledMetricItemsRequest) (*ListEnabledMetricItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledMetricItems not implemented")
}
func (UnimplementedMetricItemServiceServer) DeleteMetricItem(context.Context, *DeleteMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMetricItem not implemented")
}
func RegisterMetricItemServiceServer(s grpc.ServiceRegistrar, srv MetricItemServiceServer) {
	s.RegisterService(&MetricItemService_ServiceDesc, srv)
}

func _MetricItemService_CreateMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).CreateMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/createMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).CreateMetricItem(ctx, req.(*CreateMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricItemService_UpdateMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).UpdateMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/updateMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).UpdateMetricItem(ctx, req.(*UpdateMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricItemService_FindEnabledMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).FindEnabledMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/findEnabledMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).FindEnabledMetricItem(ctx, req.(*FindEnabledMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricItemService_CountAllEnabledMetricItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledMetricItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).CountAllEnabledMetricItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/countAllEnabledMetricItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).CountAllEnabledMetricItems(ctx, req.(*CountAllEnabledMetricItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricItemService_ListEnabledMetricItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledMetricItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).ListEnabledMetricItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/listEnabledMetricItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).ListEnabledMetricItems(ctx, req.(*ListEnabledMetricItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricItemService_DeleteMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricItemServiceServer).DeleteMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MetricItemService/deleteMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricItemServiceServer).DeleteMetricItem(ctx, req.(*DeleteMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricItemService_ServiceDesc is the grpc.ServiceDesc for MetricItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MetricItemService",
	HandlerType: (*MetricItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMetricItem",
			Handler:    _MetricItemService_CreateMetricItem_Handler,
		},
		{
			MethodName: "updateMetricItem",
			Handler:    _MetricItemService_UpdateMetricItem_Handler,
		},
		{
			MethodName: "findEnabledMetricItem",
			Handler:    _MetricItemService_FindEnabledMetricItem_Handler,
		},
		{
			MethodName: "countAllEnabledMetricItems",
			Handler:    _MetricItemService_CountAllEnabledMetricItems_Handler,
		},
		{
			MethodName: "listEnabledMetricItems",
			Handler:    _MetricItemService_ListEnabledMetricItems_Handler,
		},
		{
			MethodName: "deleteMetricItem",
			Handler:    _MetricItemService_DeleteMetricItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_metric_item.proto",
}
