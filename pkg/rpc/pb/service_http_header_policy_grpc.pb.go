// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_http_header_policy.proto

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

// HTTPHeaderPolicyServiceClient is the client API for HTTPHeaderPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPHeaderPolicyServiceClient interface {
	// 查找策略配置
	FindEnabledHTTPHeaderPolicyConfig(ctx context.Context, in *FindEnabledHTTPHeaderPolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPHeaderPolicyConfigResponse, error)
	// 创建策略
	CreateHTTPHeaderPolicy(ctx context.Context, in *CreateHTTPHeaderPolicyRequest, opts ...grpc.CallOption) (*CreateHTTPHeaderPolicyResponse, error)
	// 修改AddHeaders
	UpdateHTTPHeaderPolicyAddingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyAddingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改SetHeaders
	UpdateHTTPHeaderPolicySettingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicySettingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改AddTrailers
	UpdateHTTPHeaderPolicyAddingTrailers(ctx context.Context, in *UpdateHTTPHeaderPolicyAddingTrailersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改ReplaceHeaders
	UpdateHTTPHeaderPolicyReplacingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyReplacingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 修改删除的Headers
	UpdateHTTPHeaderPolicyDeletingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyDeletingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPHeaderPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPHeaderPolicyServiceClient(cc grpc.ClientConnInterface) HTTPHeaderPolicyServiceClient {
	return &hTTPHeaderPolicyServiceClient{cc}
}

func (c *hTTPHeaderPolicyServiceClient) FindEnabledHTTPHeaderPolicyConfig(ctx context.Context, in *FindEnabledHTTPHeaderPolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPHeaderPolicyConfigResponse, error) {
	out := new(FindEnabledHTTPHeaderPolicyConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/findEnabledHTTPHeaderPolicyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) CreateHTTPHeaderPolicy(ctx context.Context, in *CreateHTTPHeaderPolicyRequest, opts ...grpc.CallOption) (*CreateHTTPHeaderPolicyResponse, error) {
	out := new(CreateHTTPHeaderPolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/createHTTPHeaderPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) UpdateHTTPHeaderPolicyAddingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyAddingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyAddingHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) UpdateHTTPHeaderPolicySettingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicySettingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicySettingHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) UpdateHTTPHeaderPolicyAddingTrailers(ctx context.Context, in *UpdateHTTPHeaderPolicyAddingTrailersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyAddingTrailers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) UpdateHTTPHeaderPolicyReplacingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyReplacingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyReplacingHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPHeaderPolicyServiceClient) UpdateHTTPHeaderPolicyDeletingHeaders(ctx context.Context, in *UpdateHTTPHeaderPolicyDeletingHeadersRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyDeletingHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPHeaderPolicyServiceServer is the server API for HTTPHeaderPolicyService service.
// All implementations must embed UnimplementedHTTPHeaderPolicyServiceServer
// for forward compatibility
type HTTPHeaderPolicyServiceServer interface {
	// 查找策略配置
	FindEnabledHTTPHeaderPolicyConfig(context.Context, *FindEnabledHTTPHeaderPolicyConfigRequest) (*FindEnabledHTTPHeaderPolicyConfigResponse, error)
	// 创建策略
	CreateHTTPHeaderPolicy(context.Context, *CreateHTTPHeaderPolicyRequest) (*CreateHTTPHeaderPolicyResponse, error)
	// 修改AddHeaders
	UpdateHTTPHeaderPolicyAddingHeaders(context.Context, *UpdateHTTPHeaderPolicyAddingHeadersRequest) (*RPCSuccess, error)
	// 修改SetHeaders
	UpdateHTTPHeaderPolicySettingHeaders(context.Context, *UpdateHTTPHeaderPolicySettingHeadersRequest) (*RPCSuccess, error)
	// 修改AddTrailers
	UpdateHTTPHeaderPolicyAddingTrailers(context.Context, *UpdateHTTPHeaderPolicyAddingTrailersRequest) (*RPCSuccess, error)
	// 修改ReplaceHeaders
	UpdateHTTPHeaderPolicyReplacingHeaders(context.Context, *UpdateHTTPHeaderPolicyReplacingHeadersRequest) (*RPCSuccess, error)
	// 修改删除的Headers
	UpdateHTTPHeaderPolicyDeletingHeaders(context.Context, *UpdateHTTPHeaderPolicyDeletingHeadersRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPHeaderPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPHeaderPolicyServiceServer struct {
}

func (UnimplementedHTTPHeaderPolicyServiceServer) FindEnabledHTTPHeaderPolicyConfig(context.Context, *FindEnabledHTTPHeaderPolicyConfigRequest) (*FindEnabledHTTPHeaderPolicyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPHeaderPolicyConfig not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) CreateHTTPHeaderPolicy(context.Context, *CreateHTTPHeaderPolicyRequest) (*CreateHTTPHeaderPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPHeaderPolicy not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) UpdateHTTPHeaderPolicyAddingHeaders(context.Context, *UpdateHTTPHeaderPolicyAddingHeadersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPHeaderPolicyAddingHeaders not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) UpdateHTTPHeaderPolicySettingHeaders(context.Context, *UpdateHTTPHeaderPolicySettingHeadersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPHeaderPolicySettingHeaders not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) UpdateHTTPHeaderPolicyAddingTrailers(context.Context, *UpdateHTTPHeaderPolicyAddingTrailersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPHeaderPolicyAddingTrailers not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) UpdateHTTPHeaderPolicyReplacingHeaders(context.Context, *UpdateHTTPHeaderPolicyReplacingHeadersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPHeaderPolicyReplacingHeaders not implemented")
}
func (UnimplementedHTTPHeaderPolicyServiceServer) UpdateHTTPHeaderPolicyDeletingHeaders(context.Context, *UpdateHTTPHeaderPolicyDeletingHeadersRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPHeaderPolicyDeletingHeaders not implemented")
}
func RegisterHTTPHeaderPolicyServiceServer(s grpc.ServiceRegistrar, srv HTTPHeaderPolicyServiceServer) {
	s.RegisterService(&HTTPHeaderPolicyService_ServiceDesc, srv)
}

func _HTTPHeaderPolicyService_FindEnabledHTTPHeaderPolicyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPHeaderPolicyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).FindEnabledHTTPHeaderPolicyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/findEnabledHTTPHeaderPolicyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).FindEnabledHTTPHeaderPolicyConfig(ctx, req.(*FindEnabledHTTPHeaderPolicyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_CreateHTTPHeaderPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPHeaderPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).CreateHTTPHeaderPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/createHTTPHeaderPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).CreateHTTPHeaderPolicy(ctx, req.(*CreateHTTPHeaderPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyAddingHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPHeaderPolicyAddingHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyAddingHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyAddingHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyAddingHeaders(ctx, req.(*UpdateHTTPHeaderPolicyAddingHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicySettingHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPHeaderPolicySettingHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicySettingHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicySettingHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicySettingHeaders(ctx, req.(*UpdateHTTPHeaderPolicySettingHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyAddingTrailers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPHeaderPolicyAddingTrailersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyAddingTrailers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyAddingTrailers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyAddingTrailers(ctx, req.(*UpdateHTTPHeaderPolicyAddingTrailersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyReplacingHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPHeaderPolicyReplacingHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyReplacingHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyReplacingHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyReplacingHeaders(ctx, req.(*UpdateHTTPHeaderPolicyReplacingHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyDeletingHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPHeaderPolicyDeletingHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyDeletingHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPHeaderPolicyService/updateHTTPHeaderPolicyDeletingHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPHeaderPolicyServiceServer).UpdateHTTPHeaderPolicyDeletingHeaders(ctx, req.(*UpdateHTTPHeaderPolicyDeletingHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPHeaderPolicyService_ServiceDesc is the grpc.ServiceDesc for HTTPHeaderPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPHeaderPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPHeaderPolicyService",
	HandlerType: (*HTTPHeaderPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findEnabledHTTPHeaderPolicyConfig",
			Handler:    _HTTPHeaderPolicyService_FindEnabledHTTPHeaderPolicyConfig_Handler,
		},
		{
			MethodName: "createHTTPHeaderPolicy",
			Handler:    _HTTPHeaderPolicyService_CreateHTTPHeaderPolicy_Handler,
		},
		{
			MethodName: "updateHTTPHeaderPolicyAddingHeaders",
			Handler:    _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyAddingHeaders_Handler,
		},
		{
			MethodName: "updateHTTPHeaderPolicySettingHeaders",
			Handler:    _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicySettingHeaders_Handler,
		},
		{
			MethodName: "updateHTTPHeaderPolicyAddingTrailers",
			Handler:    _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyAddingTrailers_Handler,
		},
		{
			MethodName: "updateHTTPHeaderPolicyReplacingHeaders",
			Handler:    _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyReplacingHeaders_Handler,
		},
		{
			MethodName: "updateHTTPHeaderPolicyDeletingHeaders",
			Handler:    _HTTPHeaderPolicyService_UpdateHTTPHeaderPolicyDeletingHeaders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_header_policy.proto",
}
