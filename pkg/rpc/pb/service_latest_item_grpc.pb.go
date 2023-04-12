// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_latest_item.proto

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

// LatestItemServiceClient is the client API for LatestItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LatestItemServiceClient interface {
	// 记录最近使用的条目
	IncreaseLatestItem(ctx context.Context, in *IncreaseLatestItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type latestItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLatestItemServiceClient(cc grpc.ClientConnInterface) LatestItemServiceClient {
	return &latestItemServiceClient{cc}
}

func (c *latestItemServiceClient) IncreaseLatestItem(ctx context.Context, in *IncreaseLatestItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.LatestItemService/increaseLatestItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LatestItemServiceServer is the server API for LatestItemService service.
// All implementations must embed UnimplementedLatestItemServiceServer
// for forward compatibility
type LatestItemServiceServer interface {
	// 记录最近使用的条目
	IncreaseLatestItem(context.Context, *IncreaseLatestItemRequest) (*RPCSuccess, error)
}

// UnimplementedLatestItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLatestItemServiceServer struct {
}

func (UnimplementedLatestItemServiceServer) IncreaseLatestItem(context.Context, *IncreaseLatestItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseLatestItem not implemented")
}

func RegisterLatestItemServiceServer(s grpc.ServiceRegistrar, srv LatestItemServiceServer) {
	s.RegisterService(&LatestItemService_ServiceDesc, srv)
}

func _LatestItemService_IncreaseLatestItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseLatestItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LatestItemServiceServer).IncreaseLatestItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LatestItemService/increaseLatestItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LatestItemServiceServer).IncreaseLatestItem(ctx, req.(*IncreaseLatestItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LatestItemService_ServiceDesc is the grpc.ServiceDesc for LatestItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LatestItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LatestItemService",
	HandlerType: (*LatestItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "increaseLatestItem",
			Handler:    _LatestItemService_IncreaseLatestItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_latest_item.proto",
}