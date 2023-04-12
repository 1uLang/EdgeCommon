// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_message_media.proto

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

// MessageMediaServiceClient is the client API for MessageMediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageMediaServiceClient interface {
	// 获取所有支持的媒介
	FindAllMessageMedias(ctx context.Context, in *FindAllMessageMediasRequest, opts ...grpc.CallOption) (*FindAllMessageMediasResponse, error)
	// 设置所有支持的媒介
	UpdateMessageMedias(ctx context.Context, in *UpdateMessageMediasRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type messageMediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageMediaServiceClient(cc grpc.ClientConnInterface) MessageMediaServiceClient {
	return &messageMediaServiceClient{cc}
}

func (c *messageMediaServiceClient) FindAllMessageMedias(ctx context.Context, in *FindAllMessageMediasRequest, opts ...grpc.CallOption) (*FindAllMessageMediasResponse, error) {
	out := new(FindAllMessageMediasResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageMediaService/findAllMessageMedias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageMediaServiceClient) UpdateMessageMedias(ctx context.Context, in *UpdateMessageMediasRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageMediaService/updateMessageMedias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageMediaServiceServer is the server API for MessageMediaService service.
// All implementations must embed UnimplementedMessageMediaServiceServer
// for forward compatibility
type MessageMediaServiceServer interface {
	// 获取所有支持的媒介
	FindAllMessageMedias(context.Context, *FindAllMessageMediasRequest) (*FindAllMessageMediasResponse, error)
	// 设置所有支持的媒介
	UpdateMessageMedias(context.Context, *UpdateMessageMediasRequest) (*RPCSuccess, error)
}

// UnimplementedMessageMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageMediaServiceServer struct {
}

func (UnimplementedMessageMediaServiceServer) FindAllMessageMedias(context.Context, *FindAllMessageMediasRequest) (*FindAllMessageMediasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllMessageMedias not implemented")
}
func (UnimplementedMessageMediaServiceServer) UpdateMessageMedias(context.Context, *UpdateMessageMediasRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageMedias not implemented")
}
func RegisterMessageMediaServiceServer(s grpc.ServiceRegistrar, srv MessageMediaServiceServer) {
	s.RegisterService(&MessageMediaService_ServiceDesc, srv)
}

func _MessageMediaService_FindAllMessageMedias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllMessageMediasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaServiceServer).FindAllMessageMedias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageMediaService/findAllMessageMedias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaServiceServer).FindAllMessageMedias(ctx, req.(*FindAllMessageMediasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageMediaService_UpdateMessageMedias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageMediasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageMediaServiceServer).UpdateMessageMedias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageMediaService/updateMessageMedias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageMediaServiceServer).UpdateMessageMedias(ctx, req.(*UpdateMessageMediasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageMediaService_ServiceDesc is the grpc.ServiceDesc for MessageMediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageMediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageMediaService",
	HandlerType: (*MessageMediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllMessageMedias",
			Handler:    _MessageMediaService_FindAllMessageMedias_Handler,
		},
		{
			MethodName: "updateMessageMedias",
			Handler:    _MessageMediaService_UpdateMessageMedias_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_message_media.proto",
}
