// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_server_http_firewall_daily_stat.proto

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

// ServerHTTPFirewallDailyStatServiceClient is the client API for ServerHTTPFirewallDailyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerHTTPFirewallDailyStatServiceClient interface {
	// 组合服务的Dashboard
	ComposeServerHTTPFirewallDashboard(ctx context.Context, in *ComposeServerHTTPFirewallDashboardRequest, opts ...grpc.CallOption) (*ComposeServerHTTPFirewallDashboardResponse, error)
}

type serverHTTPFirewallDailyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerHTTPFirewallDailyStatServiceClient(cc grpc.ClientConnInterface) ServerHTTPFirewallDailyStatServiceClient {
	return &serverHTTPFirewallDailyStatServiceClient{cc}
}

func (c *serverHTTPFirewallDailyStatServiceClient) ComposeServerHTTPFirewallDashboard(ctx context.Context, in *ComposeServerHTTPFirewallDashboardRequest, opts ...grpc.CallOption) (*ComposeServerHTTPFirewallDashboardResponse, error) {
	out := new(ComposeServerHTTPFirewallDashboardResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerHTTPFirewallDailyStatService/composeServerHTTPFirewallDashboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerHTTPFirewallDailyStatServiceServer is the server API for ServerHTTPFirewallDailyStatService service.
// All implementations must embed UnimplementedServerHTTPFirewallDailyStatServiceServer
// for forward compatibility
type ServerHTTPFirewallDailyStatServiceServer interface {
	// 组合服务的Dashboard
	ComposeServerHTTPFirewallDashboard(context.Context, *ComposeServerHTTPFirewallDashboardRequest) (*ComposeServerHTTPFirewallDashboardResponse, error)
}

// UnimplementedServerHTTPFirewallDailyStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerHTTPFirewallDailyStatServiceServer struct {
}

func (UnimplementedServerHTTPFirewallDailyStatServiceServer) ComposeServerHTTPFirewallDashboard(context.Context, *ComposeServerHTTPFirewallDashboardRequest) (*ComposeServerHTTPFirewallDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeServerHTTPFirewallDashboard not implemented")
}
func RegisterServerHTTPFirewallDailyStatServiceServer(s grpc.ServiceRegistrar, srv ServerHTTPFirewallDailyStatServiceServer) {
	s.RegisterService(&ServerHTTPFirewallDailyStatService_ServiceDesc, srv)
}

func _ServerHTTPFirewallDailyStatService_ComposeServerHTTPFirewallDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeServerHTTPFirewallDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerHTTPFirewallDailyStatServiceServer).ComposeServerHTTPFirewallDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerHTTPFirewallDailyStatService/composeServerHTTPFirewallDashboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerHTTPFirewallDailyStatServiceServer).ComposeServerHTTPFirewallDashboard(ctx, req.(*ComposeServerHTTPFirewallDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerHTTPFirewallDailyStatService_ServiceDesc is the grpc.ServiceDesc for ServerHTTPFirewallDailyStatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerHTTPFirewallDailyStatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerHTTPFirewallDailyStatService",
	HandlerType: (*ServerHTTPFirewallDailyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "composeServerHTTPFirewallDashboard",
			Handler:    _ServerHTTPFirewallDailyStatService_ComposeServerHTTPFirewallDashboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_http_firewall_daily_stat.proto",
}
