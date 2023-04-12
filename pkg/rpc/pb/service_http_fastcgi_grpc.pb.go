// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_http_fastcgi.proto

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

// HTTPFastcgiServiceClient is the client API for HTTPFastcgiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPFastcgiServiceClient interface {
	// 创建Fastcgi
	CreateHTTPFastcgi(ctx context.Context, in *CreateHTTPFastcgiRequest, opts ...grpc.CallOption) (*CreateHTTPFastcgiResponse, error)
	// 修改Fastcgi
	UpdateHTTPFastcgi(ctx context.Context, in *UpdateHTTPFastcgiRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 获取Fastcgi详情
	FindEnabledHTTPFastcgi(ctx context.Context, in *FindEnabledHTTPFastcgiRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiResponse, error)
	// 获取Fastcgi配置
	FindEnabledHTTPFastcgiConfig(ctx context.Context, in *FindEnabledHTTPFastcgiConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiConfigResponse, error)
}

type hTTPFastcgiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPFastcgiServiceClient(cc grpc.ClientConnInterface) HTTPFastcgiServiceClient {
	return &hTTPFastcgiServiceClient{cc}
}

func (c *hTTPFastcgiServiceClient) CreateHTTPFastcgi(ctx context.Context, in *CreateHTTPFastcgiRequest, opts ...grpc.CallOption) (*CreateHTTPFastcgiResponse, error) {
	out := new(CreateHTTPFastcgiResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/createHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) UpdateHTTPFastcgi(ctx context.Context, in *UpdateHTTPFastcgiRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/updateHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) FindEnabledHTTPFastcgi(ctx context.Context, in *FindEnabledHTTPFastcgiRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiResponse, error) {
	out := new(FindEnabledHTTPFastcgiResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/findEnabledHTTPFastcgi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFastcgiServiceClient) FindEnabledHTTPFastcgiConfig(ctx context.Context, in *FindEnabledHTTPFastcgiConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFastcgiConfigResponse, error) {
	out := new(FindEnabledHTTPFastcgiConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFastcgiService/findEnabledHTTPFastcgiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPFastcgiServiceServer is the server API for HTTPFastcgiService service.
// All implementations must embed UnimplementedHTTPFastcgiServiceServer
// for forward compatibility
type HTTPFastcgiServiceServer interface {
	// 创建Fastcgi
	CreateHTTPFastcgi(context.Context, *CreateHTTPFastcgiRequest) (*CreateHTTPFastcgiResponse, error)
	// 修改Fastcgi
	UpdateHTTPFastcgi(context.Context, *UpdateHTTPFastcgiRequest) (*RPCSuccess, error)
	// 获取Fastcgi详情
	FindEnabledHTTPFastcgi(context.Context, *FindEnabledHTTPFastcgiRequest) (*FindEnabledHTTPFastcgiResponse, error)
	// 获取Fastcgi配置
	FindEnabledHTTPFastcgiConfig(context.Context, *FindEnabledHTTPFastcgiConfigRequest) (*FindEnabledHTTPFastcgiConfigResponse, error)
}

// UnimplementedHTTPFastcgiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPFastcgiServiceServer struct {
}

func (UnimplementedHTTPFastcgiServiceServer) CreateHTTPFastcgi(context.Context, *CreateHTTPFastcgiRequest) (*CreateHTTPFastcgiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPFastcgi not implemented")
}
func (UnimplementedHTTPFastcgiServiceServer) UpdateHTTPFastcgi(context.Context, *UpdateHTTPFastcgiRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPFastcgi not implemented")
}
func (UnimplementedHTTPFastcgiServiceServer) FindEnabledHTTPFastcgi(context.Context, *FindEnabledHTTPFastcgiRequest) (*FindEnabledHTTPFastcgiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFastcgi not implemented")
}
func (UnimplementedHTTPFastcgiServiceServer) FindEnabledHTTPFastcgiConfig(context.Context, *FindEnabledHTTPFastcgiConfigRequest) (*FindEnabledHTTPFastcgiConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFastcgiConfig not implemented")
}
func RegisterHTTPFastcgiServiceServer(s grpc.ServiceRegistrar, srv HTTPFastcgiServiceServer) {
	s.RegisterService(&HTTPFastcgiService_ServiceDesc, srv)
}

func _HTTPFastcgiService_CreateHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).CreateHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/createHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).CreateHTTPFastcgi(ctx, req.(*CreateHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_UpdateHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).UpdateHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/updateHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).UpdateHTTPFastcgi(ctx, req.(*UpdateHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_FindEnabledHTTPFastcgi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFastcgiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/findEnabledHTTPFastcgi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgi(ctx, req.(*FindEnabledHTTPFastcgiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFastcgiService_FindEnabledHTTPFastcgiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFastcgiConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFastcgiService/findEnabledHTTPFastcgiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFastcgiServiceServer).FindEnabledHTTPFastcgiConfig(ctx, req.(*FindEnabledHTTPFastcgiConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPFastcgiService_ServiceDesc is the grpc.ServiceDesc for HTTPFastcgiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPFastcgiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPFastcgiService",
	HandlerType: (*HTTPFastcgiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPFastcgi",
			Handler:    _HTTPFastcgiService_CreateHTTPFastcgi_Handler,
		},
		{
			MethodName: "updateHTTPFastcgi",
			Handler:    _HTTPFastcgiService_UpdateHTTPFastcgi_Handler,
		},
		{
			MethodName: "findEnabledHTTPFastcgi",
			Handler:    _HTTPFastcgiService_FindEnabledHTTPFastcgi_Handler,
		},
		{
			MethodName: "findEnabledHTTPFastcgiConfig",
			Handler:    _HTTPFastcgiService_FindEnabledHTTPFastcgiConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_fastcgi.proto",
}