// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_sys_locker.proto

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

// SysLockerServiceClient is the client API for SysLockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysLockerServiceClient interface {
	// 获得锁
	SysLockerLock(ctx context.Context, in *SysLockerLockRequest, opts ...grpc.CallOption) (*SysLockerLockResponse, error)
	// 释放锁
	SysLockerUnlock(ctx context.Context, in *SysLockerUnlockRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type sysLockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSysLockerServiceClient(cc grpc.ClientConnInterface) SysLockerServiceClient {
	return &sysLockerServiceClient{cc}
}

func (c *sysLockerServiceClient) SysLockerLock(ctx context.Context, in *SysLockerLockRequest, opts ...grpc.CallOption) (*SysLockerLockResponse, error) {
	out := new(SysLockerLockResponse)
	err := c.cc.Invoke(ctx, "/pb.SysLockerService/SysLockerLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysLockerServiceClient) SysLockerUnlock(ctx context.Context, in *SysLockerUnlockRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.SysLockerService/SysLockerUnlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysLockerServiceServer is the server API for SysLockerService service.
// All implementations must embed UnimplementedSysLockerServiceServer
// for forward compatibility
type SysLockerServiceServer interface {
	// 获得锁
	SysLockerLock(context.Context, *SysLockerLockRequest) (*SysLockerLockResponse, error)
	// 释放锁
	SysLockerUnlock(context.Context, *SysLockerUnlockRequest) (*RPCSuccess, error)
}

// UnimplementedSysLockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSysLockerServiceServer struct {
}

func (UnimplementedSysLockerServiceServer) SysLockerLock(context.Context, *SysLockerLockRequest) (*SysLockerLockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLockerLock not implemented")
}
func (UnimplementedSysLockerServiceServer) SysLockerUnlock(context.Context, *SysLockerUnlockRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLockerUnlock not implemented")
}
func RegisterSysLockerServiceServer(s grpc.ServiceRegistrar, srv SysLockerServiceServer) {
	s.RegisterService(&SysLockerService_ServiceDesc, srv)
}

func _SysLockerService_SysLockerLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLockerLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysLockerServiceServer).SysLockerLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SysLockerService/SysLockerLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysLockerServiceServer).SysLockerLock(ctx, req.(*SysLockerLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysLockerService_SysLockerUnlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLockerUnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysLockerServiceServer).SysLockerUnlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SysLockerService/SysLockerUnlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysLockerServiceServer).SysLockerUnlock(ctx, req.(*SysLockerUnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysLockerService_ServiceDesc is the grpc.ServiceDesc for SysLockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysLockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SysLockerService",
	HandlerType: (*SysLockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysLockerLock",
			Handler:    _SysLockerService_SysLockerLock_Handler,
		},
		{
			MethodName: "SysLockerUnlock",
			Handler:    _SysLockerService_SysLockerUnlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_sys_locker.proto",
}