// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_authority_node.proto

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

// AuthorityNodeServiceClient is the client API for AuthorityNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorityNodeServiceClient interface {
	// 创建认证节点
	CreateAuthorityNode(ctx context.Context, in *CreateAuthorityNodeRequest, opts ...grpc.CallOption) (*CreateAuthorityNodeResponse, error)
	// 修改认证节点
	UpdateAuthorityNode(ctx context.Context, in *UpdateAuthorityNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除认证节点
	DeleteAuthorityNode(ctx context.Context, in *DeleteAuthorityNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 列出所有可用认证节点
	FindAllEnabledAuthorityNodes(ctx context.Context, in *FindAllEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledAuthorityNodesResponse, error)
	// 计算认证节点数量
	CountAllEnabledAuthorityNodes(ctx context.Context, in *CountAllEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的认证节点
	ListEnabledAuthorityNodes(ctx context.Context, in *ListEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*ListEnabledAuthorityNodesResponse, error)
	// 根据ID查找节点
	FindEnabledAuthorityNode(ctx context.Context, in *FindEnabledAuthorityNodeRequest, opts ...grpc.CallOption) (*FindEnabledAuthorityNodeResponse, error)
	// 获取当前认证节点信息
	FindCurrentAuthorityNode(ctx context.Context, in *FindCurrentAuthorityNodeRequest, opts ...grpc.CallOption) (*FindCurrentAuthorityNodeResponse, error)
	// 更新节点状态
	UpdateAuthorityNodeStatus(ctx context.Context, in *UpdateAuthorityNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type authorityNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityNodeServiceClient(cc grpc.ClientConnInterface) AuthorityNodeServiceClient {
	return &authorityNodeServiceClient{cc}
}

func (c *authorityNodeServiceClient) CreateAuthorityNode(ctx context.Context, in *CreateAuthorityNodeRequest, opts ...grpc.CallOption) (*CreateAuthorityNodeResponse, error) {
	out := new(CreateAuthorityNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/createAuthorityNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) UpdateAuthorityNode(ctx context.Context, in *UpdateAuthorityNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/updateAuthorityNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) DeleteAuthorityNode(ctx context.Context, in *DeleteAuthorityNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/deleteAuthorityNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) FindAllEnabledAuthorityNodes(ctx context.Context, in *FindAllEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledAuthorityNodesResponse, error) {
	out := new(FindAllEnabledAuthorityNodesResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/findAllEnabledAuthorityNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) CountAllEnabledAuthorityNodes(ctx context.Context, in *CountAllEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/countAllEnabledAuthorityNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) ListEnabledAuthorityNodes(ctx context.Context, in *ListEnabledAuthorityNodesRequest, opts ...grpc.CallOption) (*ListEnabledAuthorityNodesResponse, error) {
	out := new(ListEnabledAuthorityNodesResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/listEnabledAuthorityNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) FindEnabledAuthorityNode(ctx context.Context, in *FindEnabledAuthorityNodeRequest, opts ...grpc.CallOption) (*FindEnabledAuthorityNodeResponse, error) {
	out := new(FindEnabledAuthorityNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/findEnabledAuthorityNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) FindCurrentAuthorityNode(ctx context.Context, in *FindCurrentAuthorityNodeRequest, opts ...grpc.CallOption) (*FindCurrentAuthorityNodeResponse, error) {
	out := new(FindCurrentAuthorityNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/findCurrentAuthorityNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityNodeServiceClient) UpdateAuthorityNodeStatus(ctx context.Context, in *UpdateAuthorityNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.AuthorityNodeService/updateAuthorityNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityNodeServiceServer is the server API for AuthorityNodeService service.
// All implementations must embed UnimplementedAuthorityNodeServiceServer
// for forward compatibility
type AuthorityNodeServiceServer interface {
	// 创建认证节点
	CreateAuthorityNode(context.Context, *CreateAuthorityNodeRequest) (*CreateAuthorityNodeResponse, error)
	// 修改认证节点
	UpdateAuthorityNode(context.Context, *UpdateAuthorityNodeRequest) (*RPCSuccess, error)
	// 删除认证节点
	DeleteAuthorityNode(context.Context, *DeleteAuthorityNodeRequest) (*RPCSuccess, error)
	// 列出所有可用认证节点
	FindAllEnabledAuthorityNodes(context.Context, *FindAllEnabledAuthorityNodesRequest) (*FindAllEnabledAuthorityNodesResponse, error)
	// 计算认证节点数量
	CountAllEnabledAuthorityNodes(context.Context, *CountAllEnabledAuthorityNodesRequest) (*RPCCountResponse, error)
	// 列出单页的认证节点
	ListEnabledAuthorityNodes(context.Context, *ListEnabledAuthorityNodesRequest) (*ListEnabledAuthorityNodesResponse, error)
	// 根据ID查找节点
	FindEnabledAuthorityNode(context.Context, *FindEnabledAuthorityNodeRequest) (*FindEnabledAuthorityNodeResponse, error)
	// 获取当前认证节点信息
	FindCurrentAuthorityNode(context.Context, *FindCurrentAuthorityNodeRequest) (*FindCurrentAuthorityNodeResponse, error)
	// 更新节点状态
	UpdateAuthorityNodeStatus(context.Context, *UpdateAuthorityNodeStatusRequest) (*RPCSuccess, error)
}

// UnimplementedAuthorityNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorityNodeServiceServer struct {
}

func (UnimplementedAuthorityNodeServiceServer) CreateAuthorityNode(context.Context, *CreateAuthorityNodeRequest) (*CreateAuthorityNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthorityNode not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) UpdateAuthorityNode(context.Context, *UpdateAuthorityNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthorityNode not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) DeleteAuthorityNode(context.Context, *DeleteAuthorityNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthorityNode not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) FindAllEnabledAuthorityNodes(context.Context, *FindAllEnabledAuthorityNodesRequest) (*FindAllEnabledAuthorityNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledAuthorityNodes not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) CountAllEnabledAuthorityNodes(context.Context, *CountAllEnabledAuthorityNodesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledAuthorityNodes not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) ListEnabledAuthorityNodes(context.Context, *ListEnabledAuthorityNodesRequest) (*ListEnabledAuthorityNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledAuthorityNodes not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) FindEnabledAuthorityNode(context.Context, *FindEnabledAuthorityNodeRequest) (*FindEnabledAuthorityNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledAuthorityNode not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) FindCurrentAuthorityNode(context.Context, *FindCurrentAuthorityNodeRequest) (*FindCurrentAuthorityNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCurrentAuthorityNode not implemented")
}
func (UnimplementedAuthorityNodeServiceServer) UpdateAuthorityNodeStatus(context.Context, *UpdateAuthorityNodeStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthorityNodeStatus not implemented")
}
func RegisterAuthorityNodeServiceServer(s grpc.ServiceRegistrar, srv AuthorityNodeServiceServer) {
	s.RegisterService(&AuthorityNodeService_ServiceDesc, srv)
}

func _AuthorityNodeService_CreateAuthorityNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorityNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).CreateAuthorityNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/createAuthorityNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).CreateAuthorityNode(ctx, req.(*CreateAuthorityNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_UpdateAuthorityNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorityNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).UpdateAuthorityNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/updateAuthorityNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).UpdateAuthorityNode(ctx, req.(*UpdateAuthorityNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_DeleteAuthorityNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthorityNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).DeleteAuthorityNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/deleteAuthorityNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).DeleteAuthorityNode(ctx, req.(*DeleteAuthorityNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_FindAllEnabledAuthorityNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledAuthorityNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).FindAllEnabledAuthorityNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/findAllEnabledAuthorityNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).FindAllEnabledAuthorityNodes(ctx, req.(*FindAllEnabledAuthorityNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_CountAllEnabledAuthorityNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledAuthorityNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).CountAllEnabledAuthorityNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/countAllEnabledAuthorityNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).CountAllEnabledAuthorityNodes(ctx, req.(*CountAllEnabledAuthorityNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_ListEnabledAuthorityNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledAuthorityNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).ListEnabledAuthorityNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/listEnabledAuthorityNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).ListEnabledAuthorityNodes(ctx, req.(*ListEnabledAuthorityNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_FindEnabledAuthorityNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledAuthorityNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).FindEnabledAuthorityNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/findEnabledAuthorityNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).FindEnabledAuthorityNode(ctx, req.(*FindEnabledAuthorityNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_FindCurrentAuthorityNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCurrentAuthorityNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).FindCurrentAuthorityNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/findCurrentAuthorityNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).FindCurrentAuthorityNode(ctx, req.(*FindCurrentAuthorityNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityNodeService_UpdateAuthorityNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorityNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityNodeServiceServer).UpdateAuthorityNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityNodeService/updateAuthorityNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityNodeServiceServer).UpdateAuthorityNodeStatus(ctx, req.(*UpdateAuthorityNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorityNodeService_ServiceDesc is the grpc.ServiceDesc for AuthorityNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorityNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthorityNodeService",
	HandlerType: (*AuthorityNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAuthorityNode",
			Handler:    _AuthorityNodeService_CreateAuthorityNode_Handler,
		},
		{
			MethodName: "updateAuthorityNode",
			Handler:    _AuthorityNodeService_UpdateAuthorityNode_Handler,
		},
		{
			MethodName: "deleteAuthorityNode",
			Handler:    _AuthorityNodeService_DeleteAuthorityNode_Handler,
		},
		{
			MethodName: "findAllEnabledAuthorityNodes",
			Handler:    _AuthorityNodeService_FindAllEnabledAuthorityNodes_Handler,
		},
		{
			MethodName: "countAllEnabledAuthorityNodes",
			Handler:    _AuthorityNodeService_CountAllEnabledAuthorityNodes_Handler,
		},
		{
			MethodName: "listEnabledAuthorityNodes",
			Handler:    _AuthorityNodeService_ListEnabledAuthorityNodes_Handler,
		},
		{
			MethodName: "findEnabledAuthorityNode",
			Handler:    _AuthorityNodeService_FindEnabledAuthorityNode_Handler,
		},
		{
			MethodName: "findCurrentAuthorityNode",
			Handler:    _AuthorityNodeService_FindCurrentAuthorityNode_Handler,
		},
		{
			MethodName: "updateAuthorityNodeStatus",
			Handler:    _AuthorityNodeService_UpdateAuthorityNodeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_authority_node.proto",
}
