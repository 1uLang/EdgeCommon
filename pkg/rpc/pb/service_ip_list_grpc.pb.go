// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_ip_list.proto

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

// IPListServiceClient is the client API for IPListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPListServiceClient interface {
	// 创建IP列表
	CreateIPList(ctx context.Context, in *CreateIPListRequest, opts ...grpc.CallOption) (*CreateIPListResponse, error)
	// 修改IP列表
	UpdateIPList(ctx context.Context, in *UpdateIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找IP列表信息
	FindEnabledIPList(ctx context.Context, in *FindEnabledIPListRequest, opts ...grpc.CallOption) (*FindEnabledIPListResponse, error)
	// 计算名单数量
	CountAllEnabledIPLists(ctx context.Context, in *CountAllEnabledIPListsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页名单
	ListEnabledIPLists(ctx context.Context, in *ListEnabledIPListsRequest, opts ...grpc.CallOption) (*ListEnabledIPListsResponse, error)
	// 删除IP名单
	DeleteIPList(ctx context.Context, in *DeleteIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 检查IPList是否存在
	ExistsEnabledIPList(ctx context.Context, in *ExistsEnabledIPListRequest, opts ...grpc.CallOption) (*ExistsEnabledIPListResponse, error)
	// 根据IP来搜索IP名单
	FindEnabledIPListContainsIP(ctx context.Context, in *FindEnabledIPListContainsIPRequest, opts ...grpc.CallOption) (*FindEnabledIPListContainsIPResponse, error)
}

type iPListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPListServiceClient(cc grpc.ClientConnInterface) IPListServiceClient {
	return &iPListServiceClient{cc}
}

func (c *iPListServiceClient) CreateIPList(ctx context.Context, in *CreateIPListRequest, opts ...grpc.CallOption) (*CreateIPListResponse, error) {
	out := new(CreateIPListResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/createIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) UpdateIPList(ctx context.Context, in *UpdateIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.IPListService/updateIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) FindEnabledIPList(ctx context.Context, in *FindEnabledIPListRequest, opts ...grpc.CallOption) (*FindEnabledIPListResponse, error) {
	out := new(FindEnabledIPListResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/findEnabledIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) CountAllEnabledIPLists(ctx context.Context, in *CountAllEnabledIPListsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/countAllEnabledIPLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) ListEnabledIPLists(ctx context.Context, in *ListEnabledIPListsRequest, opts ...grpc.CallOption) (*ListEnabledIPListsResponse, error) {
	out := new(ListEnabledIPListsResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/listEnabledIPLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) DeleteIPList(ctx context.Context, in *DeleteIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.IPListService/deleteIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) ExistsEnabledIPList(ctx context.Context, in *ExistsEnabledIPListRequest, opts ...grpc.CallOption) (*ExistsEnabledIPListResponse, error) {
	out := new(ExistsEnabledIPListResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/existsEnabledIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) FindEnabledIPListContainsIP(ctx context.Context, in *FindEnabledIPListContainsIPRequest, opts ...grpc.CallOption) (*FindEnabledIPListContainsIPResponse, error) {
	out := new(FindEnabledIPListContainsIPResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/findEnabledIPListContainsIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPListServiceServer is the server API for IPListService service.
// All implementations must embed UnimplementedIPListServiceServer
// for forward compatibility
type IPListServiceServer interface {
	// 创建IP列表
	CreateIPList(context.Context, *CreateIPListRequest) (*CreateIPListResponse, error)
	// 修改IP列表
	UpdateIPList(context.Context, *UpdateIPListRequest) (*RPCSuccess, error)
	// 查找IP列表信息
	FindEnabledIPList(context.Context, *FindEnabledIPListRequest) (*FindEnabledIPListResponse, error)
	// 计算名单数量
	CountAllEnabledIPLists(context.Context, *CountAllEnabledIPListsRequest) (*RPCCountResponse, error)
	// 列出单页名单
	ListEnabledIPLists(context.Context, *ListEnabledIPListsRequest) (*ListEnabledIPListsResponse, error)
	// 删除IP名单
	DeleteIPList(context.Context, *DeleteIPListRequest) (*RPCSuccess, error)
	// 检查IPList是否存在
	ExistsEnabledIPList(context.Context, *ExistsEnabledIPListRequest) (*ExistsEnabledIPListResponse, error)
	// 根据IP来搜索IP名单
	FindEnabledIPListContainsIP(context.Context, *FindEnabledIPListContainsIPRequest) (*FindEnabledIPListContainsIPResponse, error)
}

// UnimplementedIPListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIPListServiceServer struct {
}

func (UnimplementedIPListServiceServer) CreateIPList(context.Context, *CreateIPListRequest) (*CreateIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPList not implemented")
}
func (UnimplementedIPListServiceServer) UpdateIPList(context.Context, *UpdateIPListRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPList not implemented")
}
func (UnimplementedIPListServiceServer) FindEnabledIPList(context.Context, *FindEnabledIPListRequest) (*FindEnabledIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledIPList not implemented")
}
func (UnimplementedIPListServiceServer) CountAllEnabledIPLists(context.Context, *CountAllEnabledIPListsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledIPLists not implemented")
}
func (UnimplementedIPListServiceServer) ListEnabledIPLists(context.Context, *ListEnabledIPListsRequest) (*ListEnabledIPListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledIPLists not implemented")
}
func (UnimplementedIPListServiceServer) DeleteIPList(context.Context, *DeleteIPListRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPList not implemented")
}
func (UnimplementedIPListServiceServer) ExistsEnabledIPList(context.Context, *ExistsEnabledIPListRequest) (*ExistsEnabledIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsEnabledIPList not implemented")
}
func (UnimplementedIPListServiceServer) FindEnabledIPListContainsIP(context.Context, *FindEnabledIPListContainsIPRequest) (*FindEnabledIPListContainsIPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledIPListContainsIP not implemented")
}
func RegisterIPListServiceServer(s grpc.ServiceRegistrar, srv IPListServiceServer) {
	s.RegisterService(&IPListService_ServiceDesc, srv)
}

func _IPListService_CreateIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).CreateIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/createIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).CreateIPList(ctx, req.(*CreateIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_UpdateIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).UpdateIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/updateIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).UpdateIPList(ctx, req.(*UpdateIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_FindEnabledIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).FindEnabledIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/findEnabledIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).FindEnabledIPList(ctx, req.(*FindEnabledIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_CountAllEnabledIPLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledIPListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).CountAllEnabledIPLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/countAllEnabledIPLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).CountAllEnabledIPLists(ctx, req.(*CountAllEnabledIPListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_ListEnabledIPLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledIPListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).ListEnabledIPLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/listEnabledIPLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).ListEnabledIPLists(ctx, req.(*ListEnabledIPListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_DeleteIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).DeleteIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/deleteIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).DeleteIPList(ctx, req.(*DeleteIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_ExistsEnabledIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsEnabledIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).ExistsEnabledIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/existsEnabledIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).ExistsEnabledIPList(ctx, req.(*ExistsEnabledIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_FindEnabledIPListContainsIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledIPListContainsIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).FindEnabledIPListContainsIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/findEnabledIPListContainsIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).FindEnabledIPListContainsIP(ctx, req.(*FindEnabledIPListContainsIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPListService_ServiceDesc is the grpc.ServiceDesc for IPListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPListService",
	HandlerType: (*IPListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createIPList",
			Handler:    _IPListService_CreateIPList_Handler,
		},
		{
			MethodName: "updateIPList",
			Handler:    _IPListService_UpdateIPList_Handler,
		},
		{
			MethodName: "findEnabledIPList",
			Handler:    _IPListService_FindEnabledIPList_Handler,
		},
		{
			MethodName: "countAllEnabledIPLists",
			Handler:    _IPListService_CountAllEnabledIPLists_Handler,
		},
		{
			MethodName: "listEnabledIPLists",
			Handler:    _IPListService_ListEnabledIPLists_Handler,
		},
		{
			MethodName: "deleteIPList",
			Handler:    _IPListService_DeleteIPList_Handler,
		},
		{
			MethodName: "existsEnabledIPList",
			Handler:    _IPListService_ExistsEnabledIPList_Handler,
		},
		{
			MethodName: "findEnabledIPListContainsIP",
			Handler:    _IPListService_FindEnabledIPListContainsIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ip_list.proto",
}
