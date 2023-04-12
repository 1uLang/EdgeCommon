// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_server_stat_board.proto

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

// ServerStatBoardServiceClient is the client API for ServerStatBoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStatBoardServiceClient interface {
	// 读取所有看板
	FindAllEnabledServerStatBoards(ctx context.Context, in *FindAllEnabledServerStatBoardsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerStatBoardsResponse, error)
	// 组合集群看板数据
	ComposeServerStatNodeClusterBoard(ctx context.Context, in *ComposeServerStatNodeClusterBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatNodeClusterBoardResponse, error)
	// 组合节点看板数据
	ComposeServerStatNodeBoard(ctx context.Context, in *ComposeServerStatNodeBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatNodeBoardResponse, error)
	// 组合服务看板数据
	ComposeServerStatBoard(ctx context.Context, in *ComposeServerStatBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatBoardResponse, error)
}

type serverStatBoardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStatBoardServiceClient(cc grpc.ClientConnInterface) ServerStatBoardServiceClient {
	return &serverStatBoardServiceClient{cc}
}

func (c *serverStatBoardServiceClient) FindAllEnabledServerStatBoards(ctx context.Context, in *FindAllEnabledServerStatBoardsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerStatBoardsResponse, error) {
	out := new(FindAllEnabledServerStatBoardsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerStatBoardService/findAllEnabledServerStatBoards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatBoardServiceClient) ComposeServerStatNodeClusterBoard(ctx context.Context, in *ComposeServerStatNodeClusterBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatNodeClusterBoardResponse, error) {
	out := new(ComposeServerStatNodeClusterBoardResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerStatBoardService/composeServerStatNodeClusterBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatBoardServiceClient) ComposeServerStatNodeBoard(ctx context.Context, in *ComposeServerStatNodeBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatNodeBoardResponse, error) {
	out := new(ComposeServerStatNodeBoardResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerStatBoardService/composeServerStatNodeBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatBoardServiceClient) ComposeServerStatBoard(ctx context.Context, in *ComposeServerStatBoardRequest, opts ...grpc.CallOption) (*ComposeServerStatBoardResponse, error) {
	out := new(ComposeServerStatBoardResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerStatBoardService/composeServerStatBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerStatBoardServiceServer is the server API for ServerStatBoardService service.
// All implementations must embed UnimplementedServerStatBoardServiceServer
// for forward compatibility
type ServerStatBoardServiceServer interface {
	// 读取所有看板
	FindAllEnabledServerStatBoards(context.Context, *FindAllEnabledServerStatBoardsRequest) (*FindAllEnabledServerStatBoardsResponse, error)
	// 组合集群看板数据
	ComposeServerStatNodeClusterBoard(context.Context, *ComposeServerStatNodeClusterBoardRequest) (*ComposeServerStatNodeClusterBoardResponse, error)
	// 组合节点看板数据
	ComposeServerStatNodeBoard(context.Context, *ComposeServerStatNodeBoardRequest) (*ComposeServerStatNodeBoardResponse, error)
	// 组合服务看板数据
	ComposeServerStatBoard(context.Context, *ComposeServerStatBoardRequest) (*ComposeServerStatBoardResponse, error)
}

// UnimplementedServerStatBoardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerStatBoardServiceServer struct {
}

func (UnimplementedServerStatBoardServiceServer) FindAllEnabledServerStatBoards(context.Context, *FindAllEnabledServerStatBoardsRequest) (*FindAllEnabledServerStatBoardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledServerStatBoards not implemented")
}
func (UnimplementedServerStatBoardServiceServer) ComposeServerStatNodeClusterBoard(context.Context, *ComposeServerStatNodeClusterBoardRequest) (*ComposeServerStatNodeClusterBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeServerStatNodeClusterBoard not implemented")
}
func (UnimplementedServerStatBoardServiceServer) ComposeServerStatNodeBoard(context.Context, *ComposeServerStatNodeBoardRequest) (*ComposeServerStatNodeBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeServerStatNodeBoard not implemented")
}
func (UnimplementedServerStatBoardServiceServer) ComposeServerStatBoard(context.Context, *ComposeServerStatBoardRequest) (*ComposeServerStatBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeServerStatBoard not implemented")
}

func RegisterServerStatBoardServiceServer(s grpc.ServiceRegistrar, srv ServerStatBoardServiceServer) {
	s.RegisterService(&ServerStatBoardService_ServiceDesc, srv)
}

func _ServerStatBoardService_FindAllEnabledServerStatBoards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledServerStatBoardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardServiceServer).FindAllEnabledServerStatBoards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerStatBoardService/findAllEnabledServerStatBoards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardServiceServer).FindAllEnabledServerStatBoards(ctx, req.(*FindAllEnabledServerStatBoardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatBoardService_ComposeServerStatNodeClusterBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServerStatNodeClusterBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatNodeClusterBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerStatBoardService/composeServerStatNodeClusterBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatNodeClusterBoard(ctx, req.(*ComposeServerStatNodeClusterBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatBoardService_ComposeServerStatNodeBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServerStatNodeBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatNodeBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerStatBoardService/composeServerStatNodeBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatNodeBoard(ctx, req.(*ComposeServerStatNodeBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatBoardService_ComposeServerStatBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServerStatBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerStatBoardService/composeServerStatBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardServiceServer).ComposeServerStatBoard(ctx, req.(*ComposeServerStatBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerStatBoardService_ServiceDesc is the grpc.ServiceDesc for ServerStatBoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStatBoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerStatBoardService",
	HandlerType: (*ServerStatBoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledServerStatBoards",
			Handler:    _ServerStatBoardService_FindAllEnabledServerStatBoards_Handler,
		},
		{
			MethodName: "composeServerStatNodeClusterBoard",
			Handler:    _ServerStatBoardService_ComposeServerStatNodeClusterBoard_Handler,
		},
		{
			MethodName: "composeServerStatNodeBoard",
			Handler:    _ServerStatBoardService_ComposeServerStatNodeBoard_Handler,
		},
		{
			MethodName: "composeServerStatBoard",
			Handler:    _ServerStatBoardService_ComposeServerStatBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_stat_board.proto",
}