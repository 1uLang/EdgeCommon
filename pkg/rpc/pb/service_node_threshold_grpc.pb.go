// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_node_threshold.proto

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

// NodeThresholdServiceClient is the client API for NodeThresholdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeThresholdServiceClient interface {
	// 创建阈值
	CreateNodeThreshold(ctx context.Context, in *CreateNodeThresholdRequest, opts ...grpc.CallOption) (*CreateNodeThresholdResponse, error)
	// 修改阈值
	UpdateNodeThreshold(ctx context.Context, in *UpdateNodeThresholdRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除阈值
	DeleteNodeThreshold(ctx context.Context, in *DeleteNodeThresholdRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询阈值
	FindAllEnabledNodeThresholds(ctx context.Context, in *FindAllEnabledNodeThresholdsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeThresholdsResponse, error)
	// 计算阈值数量
	CountAllEnabledNodeThresholds(ctx context.Context, in *CountAllEnabledNodeThresholdsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 查询单个阈值详情
	FindEnabledNodeThreshold(ctx context.Context, in *FindEnabledNodeThresholdRequest, opts ...grpc.CallOption) (*FindEnabledNodeThresholdResponse, error)
}

type nodeThresholdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeThresholdServiceClient(cc grpc.ClientConnInterface) NodeThresholdServiceClient {
	return &nodeThresholdServiceClient{cc}
}

func (c *nodeThresholdServiceClient) CreateNodeThreshold(ctx context.Context, in *CreateNodeThresholdRequest, opts ...grpc.CallOption) (*CreateNodeThresholdResponse, error) {
	out := new(CreateNodeThresholdResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/createNodeThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeThresholdServiceClient) UpdateNodeThreshold(ctx context.Context, in *UpdateNodeThresholdRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/updateNodeThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeThresholdServiceClient) DeleteNodeThreshold(ctx context.Context, in *DeleteNodeThresholdRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/deleteNodeThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeThresholdServiceClient) FindAllEnabledNodeThresholds(ctx context.Context, in *FindAllEnabledNodeThresholdsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeThresholdsResponse, error) {
	out := new(FindAllEnabledNodeThresholdsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/findAllEnabledNodeThresholds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeThresholdServiceClient) CountAllEnabledNodeThresholds(ctx context.Context, in *CountAllEnabledNodeThresholdsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/countAllEnabledNodeThresholds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeThresholdServiceClient) FindEnabledNodeThreshold(ctx context.Context, in *FindEnabledNodeThresholdRequest, opts ...grpc.CallOption) (*FindEnabledNodeThresholdResponse, error) {
	out := new(FindEnabledNodeThresholdResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeThresholdService/findEnabledNodeThreshold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeThresholdServiceServer is the server API for NodeThresholdService service.
// All implementations must embed UnimplementedNodeThresholdServiceServer
// for forward compatibility
type NodeThresholdServiceServer interface {
	// 创建阈值
	CreateNodeThreshold(context.Context, *CreateNodeThresholdRequest) (*CreateNodeThresholdResponse, error)
	// 修改阈值
	UpdateNodeThreshold(context.Context, *UpdateNodeThresholdRequest) (*RPCSuccess, error)
	// 删除阈值
	DeleteNodeThreshold(context.Context, *DeleteNodeThresholdRequest) (*RPCSuccess, error)
	// 查询阈值
	FindAllEnabledNodeThresholds(context.Context, *FindAllEnabledNodeThresholdsRequest) (*FindAllEnabledNodeThresholdsResponse, error)
	// 计算阈值数量
	CountAllEnabledNodeThresholds(context.Context, *CountAllEnabledNodeThresholdsRequest) (*RPCCountResponse, error)
	// 查询单个阈值详情
	FindEnabledNodeThreshold(context.Context, *FindEnabledNodeThresholdRequest) (*FindEnabledNodeThresholdResponse, error)
}

// UnimplementedNodeThresholdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeThresholdServiceServer struct {
}

func (UnimplementedNodeThresholdServiceServer) CreateNodeThreshold(context.Context, *CreateNodeThresholdRequest) (*CreateNodeThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeThreshold not implemented")
}
func (UnimplementedNodeThresholdServiceServer) UpdateNodeThreshold(context.Context, *UpdateNodeThresholdRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeThreshold not implemented")
}
func (UnimplementedNodeThresholdServiceServer) DeleteNodeThreshold(context.Context, *DeleteNodeThresholdRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeThreshold not implemented")
}
func (UnimplementedNodeThresholdServiceServer) FindAllEnabledNodeThresholds(context.Context, *FindAllEnabledNodeThresholdsRequest) (*FindAllEnabledNodeThresholdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledNodeThresholds not implemented")
}
func (UnimplementedNodeThresholdServiceServer) CountAllEnabledNodeThresholds(context.Context, *CountAllEnabledNodeThresholdsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledNodeThresholds not implemented")
}
func (UnimplementedNodeThresholdServiceServer) FindEnabledNodeThreshold(context.Context, *FindEnabledNodeThresholdRequest) (*FindEnabledNodeThresholdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledNodeThreshold not implemented")
}
func RegisterNodeThresholdServiceServer(s grpc.ServiceRegistrar, srv NodeThresholdServiceServer) {
	s.RegisterService(&NodeThresholdService_ServiceDesc, srv)
}

func _NodeThresholdService_CreateNodeThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).CreateNodeThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/createNodeThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).CreateNodeThreshold(ctx, req.(*CreateNodeThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeThresholdService_UpdateNodeThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).UpdateNodeThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/updateNodeThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).UpdateNodeThreshold(ctx, req.(*UpdateNodeThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeThresholdService_DeleteNodeThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).DeleteNodeThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/deleteNodeThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).DeleteNodeThreshold(ctx, req.(*DeleteNodeThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeThresholdService_FindAllEnabledNodeThresholds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledNodeThresholdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).FindAllEnabledNodeThresholds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/findAllEnabledNodeThresholds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).FindAllEnabledNodeThresholds(ctx, req.(*FindAllEnabledNodeThresholdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeThresholdService_CountAllEnabledNodeThresholds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledNodeThresholdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).CountAllEnabledNodeThresholds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/countAllEnabledNodeThresholds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).CountAllEnabledNodeThresholds(ctx, req.(*CountAllEnabledNodeThresholdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeThresholdService_FindEnabledNodeThreshold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledNodeThresholdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeThresholdServiceServer).FindEnabledNodeThreshold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeThresholdService/findEnabledNodeThreshold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeThresholdServiceServer).FindEnabledNodeThreshold(ctx, req.(*FindEnabledNodeThresholdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeThresholdService_ServiceDesc is the grpc.ServiceDesc for NodeThresholdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeThresholdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeThresholdService",
	HandlerType: (*NodeThresholdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeThreshold",
			Handler:    _NodeThresholdService_CreateNodeThreshold_Handler,
		},
		{
			MethodName: "updateNodeThreshold",
			Handler:    _NodeThresholdService_UpdateNodeThreshold_Handler,
		},
		{
			MethodName: "deleteNodeThreshold",
			Handler:    _NodeThresholdService_DeleteNodeThreshold_Handler,
		},
		{
			MethodName: "findAllEnabledNodeThresholds",
			Handler:    _NodeThresholdService_FindAllEnabledNodeThresholds_Handler,
		},
		{
			MethodName: "countAllEnabledNodeThresholds",
			Handler:    _NodeThresholdService_CountAllEnabledNodeThresholds_Handler,
		},
		{
			MethodName: "findEnabledNodeThreshold",
			Handler:    _NodeThresholdService_FindEnabledNodeThreshold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_threshold.proto",
}
