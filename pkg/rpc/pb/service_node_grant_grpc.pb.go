// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_node_grant.proto

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

// NodeGrantServiceClient is the client API for NodeGrantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeGrantServiceClient interface {
	// 创建认证
	CreateNodeGrant(ctx context.Context, in *CreateNodeGrantRequest, opts ...grpc.CallOption) (*CreateNodeGrantResponse, error)
	// 修改认证
	UpdateNodeGrant(ctx context.Context, in *UpdateNodeGrantRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 禁用认证
	DisableNodeGrant(ctx context.Context, in *DisableNodeGrantRequest, opts ...grpc.CallOption) (*DisableNodeGrantResponse, error)
	// 计算认证的数量
	CountAllEnabledNodeGrants(ctx context.Context, in *CountAllEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页认证
	ListEnabledNodeGrants(ctx context.Context, in *ListEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*ListEnabledNodeGrantsResponse, error)
	// 列出所有认证
	FindAllEnabledNodeGrants(ctx context.Context, in *FindAllEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGrantsResponse, error)
	// 获取单个认证信息
	FindEnabledNodeGrant(ctx context.Context, in *FindEnabledNodeGrantRequest, opts ...grpc.CallOption) (*FindEnabledNodeGrantResponse, error)
	// 测试连接
	TestNodeGrant(ctx context.Context, in *TestNodeGrantRequest, opts ...grpc.CallOption) (*TestNodeGrantResponse, error)
	// 查找集群推荐的认证
	FindSuggestNodeGrants(ctx context.Context, in *FindSuggestNodeGrantsRequest, opts ...grpc.CallOption) (*FindSuggestNodeGrantsResponse, error)
}

type nodeGrantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeGrantServiceClient(cc grpc.ClientConnInterface) NodeGrantServiceClient {
	return &nodeGrantServiceClient{cc}
}

func (c *nodeGrantServiceClient) CreateNodeGrant(ctx context.Context, in *CreateNodeGrantRequest, opts ...grpc.CallOption) (*CreateNodeGrantResponse, error) {
	out := new(CreateNodeGrantResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/createNodeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) UpdateNodeGrant(ctx context.Context, in *UpdateNodeGrantRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/updateNodeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) DisableNodeGrant(ctx context.Context, in *DisableNodeGrantRequest, opts ...grpc.CallOption) (*DisableNodeGrantResponse, error) {
	out := new(DisableNodeGrantResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/disableNodeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) CountAllEnabledNodeGrants(ctx context.Context, in *CountAllEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/countAllEnabledNodeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) ListEnabledNodeGrants(ctx context.Context, in *ListEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*ListEnabledNodeGrantsResponse, error) {
	out := new(ListEnabledNodeGrantsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/listEnabledNodeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) FindAllEnabledNodeGrants(ctx context.Context, in *FindAllEnabledNodeGrantsRequest, opts ...grpc.CallOption) (*FindAllEnabledNodeGrantsResponse, error) {
	out := new(FindAllEnabledNodeGrantsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/findAllEnabledNodeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) FindEnabledNodeGrant(ctx context.Context, in *FindEnabledNodeGrantRequest, opts ...grpc.CallOption) (*FindEnabledNodeGrantResponse, error) {
	out := new(FindEnabledNodeGrantResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/findEnabledNodeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) TestNodeGrant(ctx context.Context, in *TestNodeGrantRequest, opts ...grpc.CallOption) (*TestNodeGrantResponse, error) {
	out := new(TestNodeGrantResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/testNodeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeGrantServiceClient) FindSuggestNodeGrants(ctx context.Context, in *FindSuggestNodeGrantsRequest, opts ...grpc.CallOption) (*FindSuggestNodeGrantsResponse, error) {
	out := new(FindSuggestNodeGrantsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeGrantService/findSuggestNodeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeGrantServiceServer is the server API for NodeGrantService service.
// All implementations must embed UnimplementedNodeGrantServiceServer
// for forward compatibility
type NodeGrantServiceServer interface {
	// 创建认证
	CreateNodeGrant(context.Context, *CreateNodeGrantRequest) (*CreateNodeGrantResponse, error)
	// 修改认证
	UpdateNodeGrant(context.Context, *UpdateNodeGrantRequest) (*RPCSuccess, error)
	// 禁用认证
	DisableNodeGrant(context.Context, *DisableNodeGrantRequest) (*DisableNodeGrantResponse, error)
	// 计算认证的数量
	CountAllEnabledNodeGrants(context.Context, *CountAllEnabledNodeGrantsRequest) (*RPCCountResponse, error)
	// 列出单页认证
	ListEnabledNodeGrants(context.Context, *ListEnabledNodeGrantsRequest) (*ListEnabledNodeGrantsResponse, error)
	// 列出所有认证
	FindAllEnabledNodeGrants(context.Context, *FindAllEnabledNodeGrantsRequest) (*FindAllEnabledNodeGrantsResponse, error)
	// 获取单个认证信息
	FindEnabledNodeGrant(context.Context, *FindEnabledNodeGrantRequest) (*FindEnabledNodeGrantResponse, error)
	// 测试连接
	TestNodeGrant(context.Context, *TestNodeGrantRequest) (*TestNodeGrantResponse, error)
	// 查找集群推荐的认证
	FindSuggestNodeGrants(context.Context, *FindSuggestNodeGrantsRequest) (*FindSuggestNodeGrantsResponse, error)
}

// UnimplementedNodeGrantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeGrantServiceServer struct {
}

func (UnimplementedNodeGrantServiceServer) CreateNodeGrant(context.Context, *CreateNodeGrantRequest) (*CreateNodeGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeGrant not implemented")
}
func (UnimplementedNodeGrantServiceServer) UpdateNodeGrant(context.Context, *UpdateNodeGrantRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeGrant not implemented")
}
func (UnimplementedNodeGrantServiceServer) DisableNodeGrant(context.Context, *DisableNodeGrantRequest) (*DisableNodeGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableNodeGrant not implemented")
}
func (UnimplementedNodeGrantServiceServer) CountAllEnabledNodeGrants(context.Context, *CountAllEnabledNodeGrantsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledNodeGrants not implemented")
}
func (UnimplementedNodeGrantServiceServer) ListEnabledNodeGrants(context.Context, *ListEnabledNodeGrantsRequest) (*ListEnabledNodeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledNodeGrants not implemented")
}
func (UnimplementedNodeGrantServiceServer) FindAllEnabledNodeGrants(context.Context, *FindAllEnabledNodeGrantsRequest) (*FindAllEnabledNodeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledNodeGrants not implemented")
}
func (UnimplementedNodeGrantServiceServer) FindEnabledNodeGrant(context.Context, *FindEnabledNodeGrantRequest) (*FindEnabledNodeGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledNodeGrant not implemented")
}
func (UnimplementedNodeGrantServiceServer) TestNodeGrant(context.Context, *TestNodeGrantRequest) (*TestNodeGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestNodeGrant not implemented")
}
func (UnimplementedNodeGrantServiceServer) FindSuggestNodeGrants(context.Context, *FindSuggestNodeGrantsRequest) (*FindSuggestNodeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuggestNodeGrants not implemented")
}
func RegisterNodeGrantServiceServer(s grpc.ServiceRegistrar, srv NodeGrantServiceServer) {
	s.RegisterService(&NodeGrantService_ServiceDesc, srv)
}

func _NodeGrantService_CreateNodeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).CreateNodeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/createNodeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).CreateNodeGrant(ctx, req.(*CreateNodeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_UpdateNodeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).UpdateNodeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/updateNodeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).UpdateNodeGrant(ctx, req.(*UpdateNodeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_DisableNodeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableNodeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).DisableNodeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/disableNodeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).DisableNodeGrant(ctx, req.(*DisableNodeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_CountAllEnabledNodeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledNodeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).CountAllEnabledNodeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/countAllEnabledNodeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).CountAllEnabledNodeGrants(ctx, req.(*CountAllEnabledNodeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_ListEnabledNodeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledNodeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).ListEnabledNodeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/listEnabledNodeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).ListEnabledNodeGrants(ctx, req.(*ListEnabledNodeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_FindAllEnabledNodeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledNodeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).FindAllEnabledNodeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/findAllEnabledNodeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).FindAllEnabledNodeGrants(ctx, req.(*FindAllEnabledNodeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_FindEnabledNodeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledNodeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).FindEnabledNodeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/findEnabledNodeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).FindEnabledNodeGrant(ctx, req.(*FindEnabledNodeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_TestNodeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestNodeGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).TestNodeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/testNodeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).TestNodeGrant(ctx, req.(*TestNodeGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeGrantService_FindSuggestNodeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSuggestNodeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeGrantServiceServer).FindSuggestNodeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeGrantService/findSuggestNodeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeGrantServiceServer).FindSuggestNodeGrants(ctx, req.(*FindSuggestNodeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeGrantService_ServiceDesc is the grpc.ServiceDesc for NodeGrantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeGrantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeGrantService",
	HandlerType: (*NodeGrantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeGrant",
			Handler:    _NodeGrantService_CreateNodeGrant_Handler,
		},
		{
			MethodName: "updateNodeGrant",
			Handler:    _NodeGrantService_UpdateNodeGrant_Handler,
		},
		{
			MethodName: "disableNodeGrant",
			Handler:    _NodeGrantService_DisableNodeGrant_Handler,
		},
		{
			MethodName: "countAllEnabledNodeGrants",
			Handler:    _NodeGrantService_CountAllEnabledNodeGrants_Handler,
		},
		{
			MethodName: "listEnabledNodeGrants",
			Handler:    _NodeGrantService_ListEnabledNodeGrants_Handler,
		},
		{
			MethodName: "findAllEnabledNodeGrants",
			Handler:    _NodeGrantService_FindAllEnabledNodeGrants_Handler,
		},
		{
			MethodName: "findEnabledNodeGrant",
			Handler:    _NodeGrantService_FindEnabledNodeGrant_Handler,
		},
		{
			MethodName: "testNodeGrant",
			Handler:    _NodeGrantService_TestNodeGrant_Handler,
		},
		{
			MethodName: "findSuggestNodeGrants",
			Handler:    _NodeGrantService_FindSuggestNodeGrants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_grant.proto",
}
