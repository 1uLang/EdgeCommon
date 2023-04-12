// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_http_cache_policy.proto

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

// HTTPCachePolicyServiceClient is the client API for HTTPCachePolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPCachePolicyServiceClient interface {
	// 获取所有可用策略
	FindAllEnabledHTTPCachePolicies(ctx context.Context, in *FindAllEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*FindAllEnabledHTTPCachePoliciesResponse, error)
	// 创建缓存策略
	CreateHTTPCachePolicy(ctx context.Context, in *CreateHTTPCachePolicyRequest, opts ...grpc.CallOption) (*CreateHTTPCachePolicyResponse, error)
	// 修改缓存策略
	UpdateHTTPCachePolicy(ctx context.Context, in *UpdateHTTPCachePolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除缓存策略
	DeleteHTTPCachePolicy(ctx context.Context, in *DeleteHTTPCachePolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算缓存策略数量
	CountAllEnabledHTTPCachePolicies(ctx context.Context, in *CountAllEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页的缓存策略
	ListEnabledHTTPCachePolicies(ctx context.Context, in *ListEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*ListEnabledHTTPCachePoliciesResponse, error)
	// 查找单个缓存策略配置
	FindEnabledHTTPCachePolicyConfig(ctx context.Context, in *FindEnabledHTTPCachePolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPCachePolicyConfigResponse, error)
	// 查找单个缓存策略信息
	FindEnabledHTTPCachePolicy(ctx context.Context, in *FindEnabledHTTPCachePolicyRequest, opts ...grpc.CallOption) (*FindEnabledHTTPCachePolicyResponse, error)
	// 设置缓存策略的默认条件
	UpdateHTTPCachePolicyRefs(ctx context.Context, in *UpdateHTTPCachePolicyRefsRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPCachePolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPCachePolicyServiceClient(cc grpc.ClientConnInterface) HTTPCachePolicyServiceClient {
	return &hTTPCachePolicyServiceClient{cc}
}

func (c *hTTPCachePolicyServiceClient) FindAllEnabledHTTPCachePolicies(ctx context.Context, in *FindAllEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*FindAllEnabledHTTPCachePoliciesResponse, error) {
	out := new(FindAllEnabledHTTPCachePoliciesResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/findAllEnabledHTTPCachePolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) CreateHTTPCachePolicy(ctx context.Context, in *CreateHTTPCachePolicyRequest, opts ...grpc.CallOption) (*CreateHTTPCachePolicyResponse, error) {
	out := new(CreateHTTPCachePolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/createHTTPCachePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) UpdateHTTPCachePolicy(ctx context.Context, in *UpdateHTTPCachePolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/updateHTTPCachePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) DeleteHTTPCachePolicy(ctx context.Context, in *DeleteHTTPCachePolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/deleteHTTPCachePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) CountAllEnabledHTTPCachePolicies(ctx context.Context, in *CountAllEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/countAllEnabledHTTPCachePolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) ListEnabledHTTPCachePolicies(ctx context.Context, in *ListEnabledHTTPCachePoliciesRequest, opts ...grpc.CallOption) (*ListEnabledHTTPCachePoliciesResponse, error) {
	out := new(ListEnabledHTTPCachePoliciesResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/listEnabledHTTPCachePolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) FindEnabledHTTPCachePolicyConfig(ctx context.Context, in *FindEnabledHTTPCachePolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPCachePolicyConfigResponse, error) {
	out := new(FindEnabledHTTPCachePolicyConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/findEnabledHTTPCachePolicyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) FindEnabledHTTPCachePolicy(ctx context.Context, in *FindEnabledHTTPCachePolicyRequest, opts ...grpc.CallOption) (*FindEnabledHTTPCachePolicyResponse, error) {
	out := new(FindEnabledHTTPCachePolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/findEnabledHTTPCachePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCachePolicyServiceClient) UpdateHTTPCachePolicyRefs(ctx context.Context, in *UpdateHTTPCachePolicyRefsRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPCachePolicyService/updateHTTPCachePolicyRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPCachePolicyServiceServer is the server API for HTTPCachePolicyService service.
// All implementations must embed UnimplementedHTTPCachePolicyServiceServer
// for forward compatibility
type HTTPCachePolicyServiceServer interface {
	// 获取所有可用策略
	FindAllEnabledHTTPCachePolicies(context.Context, *FindAllEnabledHTTPCachePoliciesRequest) (*FindAllEnabledHTTPCachePoliciesResponse, error)
	// 创建缓存策略
	CreateHTTPCachePolicy(context.Context, *CreateHTTPCachePolicyRequest) (*CreateHTTPCachePolicyResponse, error)
	// 修改缓存策略
	UpdateHTTPCachePolicy(context.Context, *UpdateHTTPCachePolicyRequest) (*RPCSuccess, error)
	// 删除缓存策略
	DeleteHTTPCachePolicy(context.Context, *DeleteHTTPCachePolicyRequest) (*RPCSuccess, error)
	// 计算缓存策略数量
	CountAllEnabledHTTPCachePolicies(context.Context, *CountAllEnabledHTTPCachePoliciesRequest) (*RPCCountResponse, error)
	// 列出单页的缓存策略
	ListEnabledHTTPCachePolicies(context.Context, *ListEnabledHTTPCachePoliciesRequest) (*ListEnabledHTTPCachePoliciesResponse, error)
	// 查找单个缓存策略配置
	FindEnabledHTTPCachePolicyConfig(context.Context, *FindEnabledHTTPCachePolicyConfigRequest) (*FindEnabledHTTPCachePolicyConfigResponse, error)
	// 查找单个缓存策略信息
	FindEnabledHTTPCachePolicy(context.Context, *FindEnabledHTTPCachePolicyRequest) (*FindEnabledHTTPCachePolicyResponse, error)
	// 设置缓存策略的默认条件
	UpdateHTTPCachePolicyRefs(context.Context, *UpdateHTTPCachePolicyRefsRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPCachePolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPCachePolicyServiceServer struct {
}

func (UnimplementedHTTPCachePolicyServiceServer) FindAllEnabledHTTPCachePolicies(context.Context, *FindAllEnabledHTTPCachePoliciesRequest) (*FindAllEnabledHTTPCachePoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledHTTPCachePolicies not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) CreateHTTPCachePolicy(context.Context, *CreateHTTPCachePolicyRequest) (*CreateHTTPCachePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPCachePolicy not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) UpdateHTTPCachePolicy(context.Context, *UpdateHTTPCachePolicyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPCachePolicy not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) DeleteHTTPCachePolicy(context.Context, *DeleteHTTPCachePolicyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHTTPCachePolicy not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) CountAllEnabledHTTPCachePolicies(context.Context, *CountAllEnabledHTTPCachePoliciesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledHTTPCachePolicies not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) ListEnabledHTTPCachePolicies(context.Context, *ListEnabledHTTPCachePoliciesRequest) (*ListEnabledHTTPCachePoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEnabledHTTPCachePolicies not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) FindEnabledHTTPCachePolicyConfig(context.Context, *FindEnabledHTTPCachePolicyConfigRequest) (*FindEnabledHTTPCachePolicyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPCachePolicyConfig not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) FindEnabledHTTPCachePolicy(context.Context, *FindEnabledHTTPCachePolicyRequest) (*FindEnabledHTTPCachePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPCachePolicy not implemented")
}
func (UnimplementedHTTPCachePolicyServiceServer) UpdateHTTPCachePolicyRefs(context.Context, *UpdateHTTPCachePolicyRefsRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPCachePolicyRefs not implemented")
}

func RegisterHTTPCachePolicyServiceServer(s grpc.ServiceRegistrar, srv HTTPCachePolicyServiceServer) {
	s.RegisterService(&HTTPCachePolicyService_ServiceDesc, srv)
}

func _HTTPCachePolicyService_FindAllEnabledHTTPCachePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledHTTPCachePoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).FindAllEnabledHTTPCachePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/findAllEnabledHTTPCachePolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).FindAllEnabledHTTPCachePolicies(ctx, req.(*FindAllEnabledHTTPCachePoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_CreateHTTPCachePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPCachePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).CreateHTTPCachePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/createHTTPCachePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).CreateHTTPCachePolicy(ctx, req.(*CreateHTTPCachePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_UpdateHTTPCachePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPCachePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).UpdateHTTPCachePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/updateHTTPCachePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).UpdateHTTPCachePolicy(ctx, req.(*UpdateHTTPCachePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_DeleteHTTPCachePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHTTPCachePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).DeleteHTTPCachePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/deleteHTTPCachePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).DeleteHTTPCachePolicy(ctx, req.(*DeleteHTTPCachePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_CountAllEnabledHTTPCachePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledHTTPCachePoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).CountAllEnabledHTTPCachePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/countAllEnabledHTTPCachePolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).CountAllEnabledHTTPCachePolicies(ctx, req.(*CountAllEnabledHTTPCachePoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_ListEnabledHTTPCachePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnabledHTTPCachePoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).ListEnabledHTTPCachePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/listEnabledHTTPCachePolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).ListEnabledHTTPCachePolicies(ctx, req.(*ListEnabledHTTPCachePoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_FindEnabledHTTPCachePolicyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPCachePolicyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).FindEnabledHTTPCachePolicyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/findEnabledHTTPCachePolicyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).FindEnabledHTTPCachePolicyConfig(ctx, req.(*FindEnabledHTTPCachePolicyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_FindEnabledHTTPCachePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPCachePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).FindEnabledHTTPCachePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/findEnabledHTTPCachePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).FindEnabledHTTPCachePolicy(ctx, req.(*FindEnabledHTTPCachePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCachePolicyService_UpdateHTTPCachePolicyRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPCachePolicyRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCachePolicyServiceServer).UpdateHTTPCachePolicyRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCachePolicyService/updateHTTPCachePolicyRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCachePolicyServiceServer).UpdateHTTPCachePolicyRefs(ctx, req.(*UpdateHTTPCachePolicyRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPCachePolicyService_ServiceDesc is the grpc.ServiceDesc for HTTPCachePolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPCachePolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPCachePolicyService",
	HandlerType: (*HTTPCachePolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledHTTPCachePolicies",
			Handler:    _HTTPCachePolicyService_FindAllEnabledHTTPCachePolicies_Handler,
		},
		{
			MethodName: "createHTTPCachePolicy",
			Handler:    _HTTPCachePolicyService_CreateHTTPCachePolicy_Handler,
		},
		{
			MethodName: "updateHTTPCachePolicy",
			Handler:    _HTTPCachePolicyService_UpdateHTTPCachePolicy_Handler,
		},
		{
			MethodName: "deleteHTTPCachePolicy",
			Handler:    _HTTPCachePolicyService_DeleteHTTPCachePolicy_Handler,
		},
		{
			MethodName: "countAllEnabledHTTPCachePolicies",
			Handler:    _HTTPCachePolicyService_CountAllEnabledHTTPCachePolicies_Handler,
		},
		{
			MethodName: "listEnabledHTTPCachePolicies",
			Handler:    _HTTPCachePolicyService_ListEnabledHTTPCachePolicies_Handler,
		},
		{
			MethodName: "findEnabledHTTPCachePolicyConfig",
			Handler:    _HTTPCachePolicyService_FindEnabledHTTPCachePolicyConfig_Handler,
		},
		{
			MethodName: "findEnabledHTTPCachePolicy",
			Handler:    _HTTPCachePolicyService_FindEnabledHTTPCachePolicy_Handler,
		},
		{
			MethodName: "updateHTTPCachePolicyRefs",
			Handler:    _HTTPCachePolicyService_UpdateHTTPCachePolicyRefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_cache_policy.proto",
}