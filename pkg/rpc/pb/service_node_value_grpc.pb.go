// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_node_value.proto

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

// NodeValueServiceClient is the client API for NodeValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeValueServiceClient interface {
	// 记录数据
	CreateNodeValue(ctx context.Context, in *CreateNodeValueRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取数据
	ListNodeValues(ctx context.Context, in *ListNodeValuesRequest, opts ...grpc.CallOption) (*ListNodeValuesResponse, error)
	// 读取所有节点的最新数据
	SumAllNodeValueStats(ctx context.Context, in *SumAllNodeValueStatsRequest, opts ...grpc.CallOption) (*SumAllNodeValueStatsResponse, error)
}

type nodeValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeValueServiceClient(cc grpc.ClientConnInterface) NodeValueServiceClient {
	return &nodeValueServiceClient{cc}
}

func (c *nodeValueServiceClient) CreateNodeValue(ctx context.Context, in *CreateNodeValueRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeValueService/createNodeValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeValueServiceClient) ListNodeValues(ctx context.Context, in *ListNodeValuesRequest, opts ...grpc.CallOption) (*ListNodeValuesResponse, error) {
	out := new(ListNodeValuesResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeValueService/listNodeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeValueServiceClient) SumAllNodeValueStats(ctx context.Context, in *SumAllNodeValueStatsRequest, opts ...grpc.CallOption) (*SumAllNodeValueStatsResponse, error) {
	out := new(SumAllNodeValueStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeValueService/sumAllNodeValueStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeValueServiceServer is the server API for NodeValueService service.
// All implementations must embed UnimplementedNodeValueServiceServer
// for forward compatibility
type NodeValueServiceServer interface {
	// 记录数据
	CreateNodeValue(context.Context, *CreateNodeValueRequest) (*RPCSuccess, error)
	// 读取数据
	ListNodeValues(context.Context, *ListNodeValuesRequest) (*ListNodeValuesResponse, error)
	// 读取所有节点的最新数据
	SumAllNodeValueStats(context.Context, *SumAllNodeValueStatsRequest) (*SumAllNodeValueStatsResponse, error)
}

// UnimplementedNodeValueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeValueServiceServer struct {
}

func (UnimplementedNodeValueServiceServer) CreateNodeValue(context.Context, *CreateNodeValueRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeValue not implemented")
}
func (UnimplementedNodeValueServiceServer) ListNodeValues(context.Context, *ListNodeValuesRequest) (*ListNodeValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodeValues not implemented")
}
func (UnimplementedNodeValueServiceServer) SumAllNodeValueStats(context.Context, *SumAllNodeValueStatsRequest) (*SumAllNodeValueStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumAllNodeValueStats not implemented")
}
func RegisterNodeValueServiceServer(s grpc.ServiceRegistrar, srv NodeValueServiceServer) {
	s.RegisterService(&NodeValueService_ServiceDesc, srv)
}

func _NodeValueService_CreateNodeValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeValueServiceServer).CreateNodeValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeValueService/createNodeValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeValueServiceServer).CreateNodeValue(ctx, req.(*CreateNodeValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeValueService_ListNodeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeValueServiceServer).ListNodeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeValueService/listNodeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeValueServiceServer).ListNodeValues(ctx, req.(*ListNodeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeValueService_SumAllNodeValueStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumAllNodeValueStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeValueServiceServer).SumAllNodeValueStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeValueService/sumAllNodeValueStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeValueServiceServer).SumAllNodeValueStats(ctx, req.(*SumAllNodeValueStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeValueService_ServiceDesc is the grpc.ServiceDesc for NodeValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeValueService",
	HandlerType: (*NodeValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeValue",
			Handler:    _NodeValueService_CreateNodeValue_Handler,
		},
		{
			MethodName: "listNodeValues",
			Handler:    _NodeValueService_ListNodeValues_Handler,
		},
		{
			MethodName: "sumAllNodeValueStats",
			Handler:    _NodeValueService_SumAllNodeValueStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_value.proto",
}