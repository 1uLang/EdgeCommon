// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_ip_library_artifact.proto

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

// IPLibraryArtifactServiceClient is the client API for IPLibraryArtifactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPLibraryArtifactServiceClient interface {
	// 创建制品
	CreateIPLibraryArtifact(ctx context.Context, in *CreateIPLibraryArtifactRequest, opts ...grpc.CallOption) (*CreateIPLibraryArtifactResponse, error)
	// 使用/取消使用制品
	UpdateIPLibraryArtifactIsPublic(ctx context.Context, in *UpdateIPLibraryArtifactIsPublicRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询所有制品
	FindAllIPLibraryArtifacts(ctx context.Context, in *FindAllIPLibraryArtifactsRequest, opts ...grpc.CallOption) (*FindAllIPLibraryArtifactsResponse, error)
	// 查找单个制品信息
	FindIPLibraryArtifact(ctx context.Context, in *FindIPLibraryArtifactRequest, opts ...grpc.CallOption) (*FindIPLibraryArtifactResponse, error)
	// 查找当前正在使用的制品
	FindPublicIPLibraryArtifact(ctx context.Context, in *FindPublicIPLibraryArtifactRequest, opts ...grpc.CallOption) (*FindPublicIPLibraryArtifactResponse, error)
	// 删除制品
	DeleteIPLibraryArtifact(ctx context.Context, in *DeleteIPLibraryArtifactRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type iPLibraryArtifactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPLibraryArtifactServiceClient(cc grpc.ClientConnInterface) IPLibraryArtifactServiceClient {
	return &iPLibraryArtifactServiceClient{cc}
}

func (c *iPLibraryArtifactServiceClient) CreateIPLibraryArtifact(ctx context.Context, in *CreateIPLibraryArtifactRequest, opts ...grpc.CallOption) (*CreateIPLibraryArtifactResponse, error) {
	out := new(CreateIPLibraryArtifactResponse)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/createIPLibraryArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryArtifactServiceClient) UpdateIPLibraryArtifactIsPublic(ctx context.Context, in *UpdateIPLibraryArtifactIsPublicRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/updateIPLibraryArtifactIsPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryArtifactServiceClient) FindAllIPLibraryArtifacts(ctx context.Context, in *FindAllIPLibraryArtifactsRequest, opts ...grpc.CallOption) (*FindAllIPLibraryArtifactsResponse, error) {
	out := new(FindAllIPLibraryArtifactsResponse)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/findAllIPLibraryArtifacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryArtifactServiceClient) FindIPLibraryArtifact(ctx context.Context, in *FindIPLibraryArtifactRequest, opts ...grpc.CallOption) (*FindIPLibraryArtifactResponse, error) {
	out := new(FindIPLibraryArtifactResponse)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/findIPLibraryArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryArtifactServiceClient) FindPublicIPLibraryArtifact(ctx context.Context, in *FindPublicIPLibraryArtifactRequest, opts ...grpc.CallOption) (*FindPublicIPLibraryArtifactResponse, error) {
	out := new(FindPublicIPLibraryArtifactResponse)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/findPublicIPLibraryArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryArtifactServiceClient) DeleteIPLibraryArtifact(ctx context.Context, in *DeleteIPLibraryArtifactRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.IPLibraryArtifactService/deleteIPLibraryArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPLibraryArtifactServiceServer is the server API for IPLibraryArtifactService service.
// All implementations must embed UnimplementedIPLibraryArtifactServiceServer
// for forward compatibility
type IPLibraryArtifactServiceServer interface {
	// 创建制品
	CreateIPLibraryArtifact(context.Context, *CreateIPLibraryArtifactRequest) (*CreateIPLibraryArtifactResponse, error)
	// 使用/取消使用制品
	UpdateIPLibraryArtifactIsPublic(context.Context, *UpdateIPLibraryArtifactIsPublicRequest) (*RPCSuccess, error)
	// 查询所有制品
	FindAllIPLibraryArtifacts(context.Context, *FindAllIPLibraryArtifactsRequest) (*FindAllIPLibraryArtifactsResponse, error)
	// 查找单个制品信息
	FindIPLibraryArtifact(context.Context, *FindIPLibraryArtifactRequest) (*FindIPLibraryArtifactResponse, error)
	// 查找当前正在使用的制品
	FindPublicIPLibraryArtifact(context.Context, *FindPublicIPLibraryArtifactRequest) (*FindPublicIPLibraryArtifactResponse, error)
	// 删除制品
	DeleteIPLibraryArtifact(context.Context, *DeleteIPLibraryArtifactRequest) (*RPCSuccess, error)
}

// UnimplementedIPLibraryArtifactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIPLibraryArtifactServiceServer struct {
}

func (UnimplementedIPLibraryArtifactServiceServer) CreateIPLibraryArtifact(context.Context, *CreateIPLibraryArtifactRequest) (*CreateIPLibraryArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPLibraryArtifact not implemented")
}
func (UnimplementedIPLibraryArtifactServiceServer) UpdateIPLibraryArtifactIsPublic(context.Context, *UpdateIPLibraryArtifactIsPublicRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPLibraryArtifactIsPublic not implemented")
}
func (UnimplementedIPLibraryArtifactServiceServer) FindAllIPLibraryArtifacts(context.Context, *FindAllIPLibraryArtifactsRequest) (*FindAllIPLibraryArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllIPLibraryArtifacts not implemented")
}
func (UnimplementedIPLibraryArtifactServiceServer) FindIPLibraryArtifact(context.Context, *FindIPLibraryArtifactRequest) (*FindIPLibraryArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindIPLibraryArtifact not implemented")
}
func (UnimplementedIPLibraryArtifactServiceServer) FindPublicIPLibraryArtifact(context.Context, *FindPublicIPLibraryArtifactRequest) (*FindPublicIPLibraryArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPublicIPLibraryArtifact not implemented")
}
func (UnimplementedIPLibraryArtifactServiceServer) DeleteIPLibraryArtifact(context.Context, *DeleteIPLibraryArtifactRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPLibraryArtifact not implemented")
}
func RegisterIPLibraryArtifactServiceServer(s grpc.ServiceRegistrar, srv IPLibraryArtifactServiceServer) {
	s.RegisterService(&IPLibraryArtifactService_ServiceDesc, srv)
}

func _IPLibraryArtifactService_CreateIPLibraryArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPLibraryArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).CreateIPLibraryArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/createIPLibraryArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).CreateIPLibraryArtifact(ctx, req.(*CreateIPLibraryArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryArtifactService_UpdateIPLibraryArtifactIsPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPLibraryArtifactIsPublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).UpdateIPLibraryArtifactIsPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/updateIPLibraryArtifactIsPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).UpdateIPLibraryArtifactIsPublic(ctx, req.(*UpdateIPLibraryArtifactIsPublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryArtifactService_FindAllIPLibraryArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllIPLibraryArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).FindAllIPLibraryArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/findAllIPLibraryArtifacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).FindAllIPLibraryArtifacts(ctx, req.(*FindAllIPLibraryArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryArtifactService_FindIPLibraryArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindIPLibraryArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).FindIPLibraryArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/findIPLibraryArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).FindIPLibraryArtifact(ctx, req.(*FindIPLibraryArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryArtifactService_FindPublicIPLibraryArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPublicIPLibraryArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).FindPublicIPLibraryArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/findPublicIPLibraryArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).FindPublicIPLibraryArtifact(ctx, req.(*FindPublicIPLibraryArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryArtifactService_DeleteIPLibraryArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIPLibraryArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryArtifactServiceServer).DeleteIPLibraryArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPLibraryArtifactService/deleteIPLibraryArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryArtifactServiceServer).DeleteIPLibraryArtifact(ctx, req.(*DeleteIPLibraryArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPLibraryArtifactService_ServiceDesc is the grpc.ServiceDesc for IPLibraryArtifactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPLibraryArtifactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPLibraryArtifactService",
	HandlerType: (*IPLibraryArtifactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createIPLibraryArtifact",
			Handler:    _IPLibraryArtifactService_CreateIPLibraryArtifact_Handler,
		},
		{
			MethodName: "updateIPLibraryArtifactIsPublic",
			Handler:    _IPLibraryArtifactService_UpdateIPLibraryArtifactIsPublic_Handler,
		},
		{
			MethodName: "findAllIPLibraryArtifacts",
			Handler:    _IPLibraryArtifactService_FindAllIPLibraryArtifacts_Handler,
		},
		{
			MethodName: "findIPLibraryArtifact",
			Handler:    _IPLibraryArtifactService_FindIPLibraryArtifact_Handler,
		},
		{
			MethodName: "findPublicIPLibraryArtifact",
			Handler:    _IPLibraryArtifactService_FindPublicIPLibraryArtifact_Handler,
		},
		{
			MethodName: "deleteIPLibraryArtifact",
			Handler:    _IPLibraryArtifactService_DeleteIPLibraryArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ip_library_artifact.proto",
}
