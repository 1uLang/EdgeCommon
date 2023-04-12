// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_plan.proto

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

// PlanServiceClient is the client API for PlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlanServiceClient interface {
	// 创建套餐
	CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error)
	// 修改套餐
	UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除套餐
	DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找单个套餐
	FindEnabledPlan(ctx context.Context, in *FindEnabledPlanRequest, opts ...grpc.CallOption) (*FindEnabledPlanResponse, error)
	// 计算套餐数量
	CountAllEnabledPlans(ctx context.Context, in *CountAllEnabledPlansRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页套餐
	ListEnabledPlans(ctx context.Context, in *ListEnabledPlansRequest, opts ...grpc.CallOption) (*ListEnabledPlansResponse, error)
	// 对套餐进行排序
	SortPlans(ctx context.Context, in *SortPlansRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type planServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlanServiceClient(cc grpc.ClientConnInterface) PlanServiceClient {
	return &planServiceClient{cc}
}

func (c *planServiceClient) CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error) {
	out := new(CreatePlanResponse)
	err := c.cc.Invoke(ctx, "/pb.PlanService/createPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.PlanService/updatePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) DeletePlan(ctx context.Context, in *DeletePlanRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.PlanService/deletePlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) FindEnabledPlan(ctx context.Context, in *FindEnabledPlanRequest, opts ...grpc.CallOption) (*FindEnabledPlanResponse, error) {
	out := new(FindEnabledPlanResponse)
	err := c.cc.Invoke(ctx, "/pb.PlanService/findEnabledPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) CountAllEnabledPlans(ctx context.Context, in *CountAllEnabledPlansRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.PlanService/countAllEnabledPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) ListEnabledPlans(ctx context.Context, in *ListEnabledPlansRequest, opts ...grpc.CallOption) (*ListEnabledPlansResponse, error) {
	out := new(ListEnabledPlansResponse)
	err := c.cc.Invoke(ctx, "/pb.PlanService/listEnabledPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planServiceClient) SortPlans(ctx context.Context, in *SortPlansRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.PlanService/sortPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanServiceServer is the server API for PlanService service.
// All implementations must embed UnimplementedPlanServiceServer
// for forward compatibility
type PlanServiceServer interface {
	// 创建套餐
	CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error)
	// 修改套餐
	UpdatePlan(context.Context, *UpdatePlanRequest) (*RPCSuccess, error)
	// 删除套餐
	DeletePlan(context.Context, *DeletePlanRequest) (*RPCSuccess, error)
	// 查找单个套餐
	FindEnabledPlan(context.Context, *FindEnabledPlanRequest) (*FindEnabledPlanResponse, error)
	// 计算套餐数量
	CountAllEnabledPlans(context.Context, *CountAllEnabledPlansRequest) (*RPCCountResponse, error)
	// 列出单页套餐
	ListEnabledPlans(context.Context, *ListEnabledPlansRequest) (*ListEnabledPlansResponse, error)
	// 对套餐进行排序
	SortPlans(context.Context, *SortPlansRequest) (*RPCSuccess, error)
}

// UnimplementedPlanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlanServiceServer struct {
}

func (UnimplementedPlanServiceServer) CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedPlanServiceServer) UpdatePlan(context.Context, *UpdatePlanRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlan not implemented")
}
func (UnimplementedPlanServiceServer) DeletePlan(context.Context, *DeletePlanRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlan not implemented")
}
func (UnimplementedPlanServiceServer) FindEnabledPlan(context.Context, *FindEnabledPlanRequest) (*FindEnabledPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledPlan not implemented")
}
func (UnimplementedPlanServiceServer) CountAllEnabledPlans(context.Context, *CountAllEnabledPlansRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledPlans not implemented")
}
func (UnimplementedPlanServiceServer) ListEnabledPlans(context.Context, *ListEnabledPlansRequest) (*ListEnabledPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledPlans not implemented")
}
func (UnimplementedPlanServiceServer) SortPlans(context.Context, *SortPlansRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortPlans not implemented")
}
func RegisterPlanServiceServer(s grpc.ServiceRegistrar, srv PlanServiceServer) {
	s.RegisterService(&PlanService_ServiceDesc, srv)
}

func _PlanService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/createPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).CreatePlan(ctx, req.(*CreatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_UpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).UpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/updatePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).UpdatePlan(ctx, req.(*UpdatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_DeletePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).DeletePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/deletePlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).DeletePlan(ctx, req.(*DeletePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_FindEnabledPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).FindEnabledPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/findEnabledPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).FindEnabledPlan(ctx, req.(*FindEnabledPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_CountAllEnabledPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).CountAllEnabledPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/countAllEnabledPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).CountAllEnabledPlans(ctx, req.(*CountAllEnabledPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_ListEnabledPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).ListEnabledPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/listEnabledPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).ListEnabledPlans(ctx, req.(*ListEnabledPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanService_SortPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanServiceServer).SortPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlanService/sortPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanServiceServer).SortPlans(ctx, req.(*SortPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlanService_ServiceDesc is the grpc.ServiceDesc for PlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PlanService",
	HandlerType: (*PlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createPlan",
			Handler:    _PlanService_CreatePlan_Handler,
		},
		{
			MethodName: "updatePlan",
			Handler:    _PlanService_UpdatePlan_Handler,
		},
		{
			MethodName: "deletePlan",
			Handler:    _PlanService_DeletePlan_Handler,
		},
		{
			MethodName: "findEnabledPlan",
			Handler:    _PlanService_FindEnabledPlan_Handler,
		},
		{
			MethodName: "countAllEnabledPlans",
			Handler:    _PlanService_CountAllEnabledPlans_Handler,
		},
		{
			MethodName: "listEnabledPlans",
			Handler:    _PlanService_ListEnabledPlans_Handler,
		},
		{
			MethodName: "sortPlans",
			Handler:    _PlanService_SortPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_plan.proto",
}
