// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_user_order.proto

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

// UserOrderServiceClient is the client API for UserOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserOrderServiceClient interface {
	// 创建订单
	CreateUserOrder(ctx context.Context, in *CreateUserOrderRequest, opts ...grpc.CallOption) (*CreateUserOrderResponse, error)
	// 查看订单
	FindEnabledUserOrder(ctx context.Context, in *FindEnabledUserOrderRequest, opts ...grpc.CallOption) (*FindEnabledUserOrderResponse, error)
	// 取消订单
	CancelUserOrder(ctx context.Context, in *CancelUserOrderRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 完成订单
	FinishUserOrder(ctx context.Context, in *FinishUserOrderRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算订单数量
	CountEnabledUserOrders(ctx context.Context, in *CountEnabledUserOrdersRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页订单
	ListEnabledUserOrders(ctx context.Context, in *ListEnabledUserOrdersRequest, opts ...grpc.CallOption) (*ListEnabledUserOrdersResponse, error)
}

type userOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserOrderServiceClient(cc grpc.ClientConnInterface) UserOrderServiceClient {
	return &userOrderServiceClient{cc}
}

func (c *userOrderServiceClient) CreateUserOrder(ctx context.Context, in *CreateUserOrderRequest, opts ...grpc.CallOption) (*CreateUserOrderResponse, error) {
	out := new(CreateUserOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/createUserOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderServiceClient) FindEnabledUserOrder(ctx context.Context, in *FindEnabledUserOrderRequest, opts ...grpc.CallOption) (*FindEnabledUserOrderResponse, error) {
	out := new(FindEnabledUserOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/findEnabledUserOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderServiceClient) CancelUserOrder(ctx context.Context, in *CancelUserOrderRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/cancelUserOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderServiceClient) FinishUserOrder(ctx context.Context, in *FinishUserOrderRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/finishUserOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderServiceClient) CountEnabledUserOrders(ctx context.Context, in *CountEnabledUserOrdersRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/countEnabledUserOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderServiceClient) ListEnabledUserOrders(ctx context.Context, in *ListEnabledUserOrdersRequest, opts ...grpc.CallOption) (*ListEnabledUserOrdersResponse, error) {
	out := new(ListEnabledUserOrdersResponse)
	err := c.cc.Invoke(ctx, "/pb.UserOrderService/listEnabledUserOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserOrderServiceServer is the server API for UserOrderService service.
// All implementations must embed UnimplementedUserOrderServiceServer
// for forward compatibility
type UserOrderServiceServer interface {
	// 创建订单
	CreateUserOrder(context.Context, *CreateUserOrderRequest) (*CreateUserOrderResponse, error)
	// 查看订单
	FindEnabledUserOrder(context.Context, *FindEnabledUserOrderRequest) (*FindEnabledUserOrderResponse, error)
	// 取消订单
	CancelUserOrder(context.Context, *CancelUserOrderRequest) (*RPCSuccess, error)
	// 完成订单
	FinishUserOrder(context.Context, *FinishUserOrderRequest) (*RPCSuccess, error)
	// 计算订单数量
	CountEnabledUserOrders(context.Context, *CountEnabledUserOrdersRequest) (*RPCCountResponse, error)
	// 列出单页订单
	ListEnabledUserOrders(context.Context, *ListEnabledUserOrdersRequest) (*ListEnabledUserOrdersResponse, error)
}

// UnimplementedUserOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserOrderServiceServer struct {
}

func (UnimplementedUserOrderServiceServer) CreateUserOrder(context.Context, *CreateUserOrderRequest) (*CreateUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserOrder not implemented")
}
func (UnimplementedUserOrderServiceServer) FindEnabledUserOrder(context.Context, *FindEnabledUserOrderRequest) (*FindEnabledUserOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledUserOrder not implemented")
}
func (UnimplementedUserOrderServiceServer) CancelUserOrder(context.Context, *CancelUserOrderRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUserOrder not implemented")
}
func (UnimplementedUserOrderServiceServer) FinishUserOrder(context.Context, *FinishUserOrderRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishUserOrder not implemented")
}
func (UnimplementedUserOrderServiceServer) CountEnabledUserOrders(context.Context, *CountEnabledUserOrdersRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountEnabledUserOrders not implemented")
}
func (UnimplementedUserOrderServiceServer) ListEnabledUserOrders(context.Context, *ListEnabledUserOrdersRequest) (*ListEnabledUserOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledUserOrders not implemented")
}
func RegisterUserOrderServiceServer(s grpc.ServiceRegistrar, srv UserOrderServiceServer) {
	s.RegisterService(&UserOrderService_ServiceDesc, srv)
}

func _UserOrderService_CreateUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).CreateUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/createUserOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).CreateUserOrder(ctx, req.(*CreateUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrderService_FindEnabledUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).FindEnabledUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/findEnabledUserOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).FindEnabledUserOrder(ctx, req.(*FindEnabledUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrderService_CancelUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).CancelUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/cancelUserOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).CancelUserOrder(ctx, req.(*CancelUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrderService_FinishUserOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishUserOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).FinishUserOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/finishUserOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).FinishUserOrder(ctx, req.(*FinishUserOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrderService_CountEnabledUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountEnabledUserOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).CountEnabledUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/countEnabledUserOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).CountEnabledUserOrders(ctx, req.(*CountEnabledUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOrderService_ListEnabledUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledUserOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOrderServiceServer).ListEnabledUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserOrderService/listEnabledUserOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOrderServiceServer).ListEnabledUserOrders(ctx, req.(*ListEnabledUserOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserOrderService_ServiceDesc is the grpc.ServiceDesc for UserOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserOrderService",
	HandlerType: (*UserOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUserOrder",
			Handler:    _UserOrderService_CreateUserOrder_Handler,
		},
		{
			MethodName: "findEnabledUserOrder",
			Handler:    _UserOrderService_FindEnabledUserOrder_Handler,
		},
		{
			MethodName: "cancelUserOrder",
			Handler:    _UserOrderService_CancelUserOrder_Handler,
		},
		{
			MethodName: "finishUserOrder",
			Handler:    _UserOrderService_FinishUserOrder_Handler,
		},
		{
			MethodName: "countEnabledUserOrders",
			Handler:    _UserOrderService_CountEnabledUserOrders_Handler,
		},
		{
			MethodName: "listEnabledUserOrders",
			Handler:    _UserOrderService_ListEnabledUserOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_order.proto",
}