// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_user_node.proto

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

// UserNodeServiceClient is the client API for UserNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserNodeServiceClient interface {
	// 创建用户节点
	CreateUserNode(ctx context.Context, in *CreateUserNodeRequest, opts ...grpc.CallOption) (*CreateUserNodeResponse, error)
	// 修改用户节点
	UpdateUserNode(ctx context.Context, in *UpdateUserNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除用户节点
	DeleteUserNode(ctx context.Context, in *DeleteUserNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 列出所有可用用户节点
	FindAllEnabledUserNodes(ctx context.Context, in *FindAllEnabledUserNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledUserNodesResponse, error)
	// 计算用户节点数量
	CountAllEnabledUserNodes(ctx context.Context, in *CountAllEnabledUserNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的用户节点
	ListEnabledUserNodes(ctx context.Context, in *ListEnabledUserNodesRequest, opts ...grpc.CallOption) (*ListEnabledUserNodesResponse, error)
	// 根据ID查找节点
	FindEnabledUserNode(ctx context.Context, in *FindEnabledUserNodeRequest, opts ...grpc.CallOption) (*FindEnabledUserNodeResponse, error)
	// 获取当前用户节点信息
	FindCurrentUserNode(ctx context.Context, in *FindCurrentUserNodeRequest, opts ...grpc.CallOption) (*FindCurrentUserNodeResponse, error)
	// 更新节点状态
	UpdateUserNodeStatus(ctx context.Context, in *UpdateUserNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算使用某个SSL证书的用户节点数量
	CountAllEnabledUserNodesWithSSLCertId(ctx context.Context, in *CountAllEnabledUserNodesWithSSLCertIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
}

type userNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserNodeServiceClient(cc grpc.ClientConnInterface) UserNodeServiceClient {
	return &userNodeServiceClient{cc}
}

func (c *userNodeServiceClient) CreateUserNode(ctx context.Context, in *CreateUserNodeRequest, opts ...grpc.CallOption) (*CreateUserNodeResponse, error) {
	out := new(CreateUserNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/createUserNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) UpdateUserNode(ctx context.Context, in *UpdateUserNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/updateUserNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) DeleteUserNode(ctx context.Context, in *DeleteUserNodeRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/deleteUserNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) FindAllEnabledUserNodes(ctx context.Context, in *FindAllEnabledUserNodesRequest, opts ...grpc.CallOption) (*FindAllEnabledUserNodesResponse, error) {
	out := new(FindAllEnabledUserNodesResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/findAllEnabledUserNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) CountAllEnabledUserNodes(ctx context.Context, in *CountAllEnabledUserNodesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/countAllEnabledUserNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) ListEnabledUserNodes(ctx context.Context, in *ListEnabledUserNodesRequest, opts ...grpc.CallOption) (*ListEnabledUserNodesResponse, error) {
	out := new(ListEnabledUserNodesResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/listEnabledUserNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) FindEnabledUserNode(ctx context.Context, in *FindEnabledUserNodeRequest, opts ...grpc.CallOption) (*FindEnabledUserNodeResponse, error) {
	out := new(FindEnabledUserNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/findEnabledUserNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) FindCurrentUserNode(ctx context.Context, in *FindCurrentUserNodeRequest, opts ...grpc.CallOption) (*FindCurrentUserNodeResponse, error) {
	out := new(FindCurrentUserNodeResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/findCurrentUserNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) UpdateUserNodeStatus(ctx context.Context, in *UpdateUserNodeStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/updateUserNodeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userNodeServiceClient) CountAllEnabledUserNodesWithSSLCertId(ctx context.Context, in *CountAllEnabledUserNodesWithSSLCertIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserNodeService/countAllEnabledUserNodesWithSSLCertId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserNodeServiceServer is the server API for UserNodeService service.
// All implementations must embed UnimplementedUserNodeServiceServer
// for forward compatibility
type UserNodeServiceServer interface {
	// 创建用户节点
	CreateUserNode(context.Context, *CreateUserNodeRequest) (*CreateUserNodeResponse, error)
	// 修改用户节点
	UpdateUserNode(context.Context, *UpdateUserNodeRequest) (*RPCSuccess, error)
	// 删除用户节点
	DeleteUserNode(context.Context, *DeleteUserNodeRequest) (*RPCSuccess, error)
	// 列出所有可用用户节点
	FindAllEnabledUserNodes(context.Context, *FindAllEnabledUserNodesRequest) (*FindAllEnabledUserNodesResponse, error)
	// 计算用户节点数量
	CountAllEnabledUserNodes(context.Context, *CountAllEnabledUserNodesRequest) (*RPCCountResponse, error)
	// 列出单页的用户节点
	ListEnabledUserNodes(context.Context, *ListEnabledUserNodesRequest) (*ListEnabledUserNodesResponse, error)
	// 根据ID查找节点
	FindEnabledUserNode(context.Context, *FindEnabledUserNodeRequest) (*FindEnabledUserNodeResponse, error)
	// 获取当前用户节点信息
	FindCurrentUserNode(context.Context, *FindCurrentUserNodeRequest) (*FindCurrentUserNodeResponse, error)
	// 更新节点状态
	UpdateUserNodeStatus(context.Context, *UpdateUserNodeStatusRequest) (*RPCSuccess, error)
	// 计算使用某个SSL证书的用户节点数量
	CountAllEnabledUserNodesWithSSLCertId(context.Context, *CountAllEnabledUserNodesWithSSLCertIdRequest) (*RPCCountResponse, error)
}

// UnimplementedUserNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserNodeServiceServer struct {
}

func (UnimplementedUserNodeServiceServer) CreateUserNode(context.Context, *CreateUserNodeRequest) (*CreateUserNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserNode not implemented")
}
func (UnimplementedUserNodeServiceServer) UpdateUserNode(context.Context, *UpdateUserNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserNode not implemented")
}
func (UnimplementedUserNodeServiceServer) DeleteUserNode(context.Context, *DeleteUserNodeRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserNode not implemented")
}
func (UnimplementedUserNodeServiceServer) FindAllEnabledUserNodes(context.Context, *FindAllEnabledUserNodesRequest) (*FindAllEnabledUserNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledUserNodes not implemented")
}
func (UnimplementedUserNodeServiceServer) CountAllEnabledUserNodes(context.Context, *CountAllEnabledUserNodesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledUserNodes not implemented")
}
func (UnimplementedUserNodeServiceServer) ListEnabledUserNodes(context.Context, *ListEnabledUserNodesRequest) (*ListEnabledUserNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledUserNodes not implemented")
}
func (UnimplementedUserNodeServiceServer) FindEnabledUserNode(context.Context, *FindEnabledUserNodeRequest) (*FindEnabledUserNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledUserNode not implemented")
}
func (UnimplementedUserNodeServiceServer) FindCurrentUserNode(context.Context, *FindCurrentUserNodeRequest) (*FindCurrentUserNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCurrentUserNode not implemented")
}
func (UnimplementedUserNodeServiceServer) UpdateUserNodeStatus(context.Context, *UpdateUserNodeStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserNodeStatus not implemented")
}
func (UnimplementedUserNodeServiceServer) CountAllEnabledUserNodesWithSSLCertId(context.Context, *CountAllEnabledUserNodesWithSSLCertIdRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledUserNodesWithSSLCertId not implemented")
}
func RegisterUserNodeServiceServer(s grpc.ServiceRegistrar, srv UserNodeServiceServer) {
	s.RegisterService(&UserNodeService_ServiceDesc, srv)
}

func _UserNodeService_CreateUserNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).CreateUserNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/createUserNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).CreateUserNode(ctx, req.(*CreateUserNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_UpdateUserNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).UpdateUserNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/updateUserNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).UpdateUserNode(ctx, req.(*UpdateUserNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_DeleteUserNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).DeleteUserNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/deleteUserNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).DeleteUserNode(ctx, req.(*DeleteUserNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_FindAllEnabledUserNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledUserNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).FindAllEnabledUserNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/findAllEnabledUserNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).FindAllEnabledUserNodes(ctx, req.(*FindAllEnabledUserNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_CountAllEnabledUserNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledUserNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).CountAllEnabledUserNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/countAllEnabledUserNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).CountAllEnabledUserNodes(ctx, req.(*CountAllEnabledUserNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_ListEnabledUserNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledUserNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).ListEnabledUserNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/listEnabledUserNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).ListEnabledUserNodes(ctx, req.(*ListEnabledUserNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_FindEnabledUserNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledUserNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).FindEnabledUserNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/findEnabledUserNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).FindEnabledUserNode(ctx, req.(*FindEnabledUserNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_FindCurrentUserNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCurrentUserNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).FindCurrentUserNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/findCurrentUserNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).FindCurrentUserNode(ctx, req.(*FindCurrentUserNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_UpdateUserNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).UpdateUserNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/updateUserNodeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).UpdateUserNodeStatus(ctx, req.(*UpdateUserNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserNodeService_CountAllEnabledUserNodesWithSSLCertId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledUserNodesWithSSLCertIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserNodeServiceServer).CountAllEnabledUserNodesWithSSLCertId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserNodeService/countAllEnabledUserNodesWithSSLCertId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserNodeServiceServer).CountAllEnabledUserNodesWithSSLCertId(ctx, req.(*CountAllEnabledUserNodesWithSSLCertIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserNodeService_ServiceDesc is the grpc.ServiceDesc for UserNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserNodeService",
	HandlerType: (*UserNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUserNode",
			Handler:    _UserNodeService_CreateUserNode_Handler,
		},
		{
			MethodName: "updateUserNode",
			Handler:    _UserNodeService_UpdateUserNode_Handler,
		},
		{
			MethodName: "deleteUserNode",
			Handler:    _UserNodeService_DeleteUserNode_Handler,
		},
		{
			MethodName: "findAllEnabledUserNodes",
			Handler:    _UserNodeService_FindAllEnabledUserNodes_Handler,
		},
		{
			MethodName: "countAllEnabledUserNodes",
			Handler:    _UserNodeService_CountAllEnabledUserNodes_Handler,
		},
		{
			MethodName: "listEnabledUserNodes",
			Handler:    _UserNodeService_ListEnabledUserNodes_Handler,
		},
		{
			MethodName: "findEnabledUserNode",
			Handler:    _UserNodeService_FindEnabledUserNode_Handler,
		},
		{
			MethodName: "findCurrentUserNode",
			Handler:    _UserNodeService_FindCurrentUserNode_Handler,
		},
		{
			MethodName: "updateUserNodeStatus",
			Handler:    _UserNodeService_UpdateUserNodeStatus_Handler,
		},
		{
			MethodName: "countAllEnabledUserNodesWithSSLCertId",
			Handler:    _UserNodeService_CountAllEnabledUserNodesWithSSLCertId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_node.proto",
}
