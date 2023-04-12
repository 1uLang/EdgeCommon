// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_node_task.proto

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

// NodeTaskServiceClient is the client API for NodeTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeTaskServiceClient interface {
	// 获取单节点同步任务
	FindNodeTasks(ctx context.Context, in *FindNodeTasksRequest, opts ...grpc.CallOption) (*FindNodeTasksResponse, error)
	// 报告同步任务结果
	ReportNodeTaskDone(ctx context.Context, in *ReportNodeTaskDoneRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 获取所有正在同步的集群信息
	FindNodeClusterTasks(ctx context.Context, in *FindNodeClusterTasksRequest, opts ...grpc.CallOption) (*FindNodeClusterTasksResponse, error)
	// 检查是否有正在执行的任务
	ExistsNodeTasks(ctx context.Context, in *ExistsNodeTasksRequest, opts ...grpc.CallOption) (*ExistsNodeTasksResponse, error)
	// 删除任务
	DeleteNodeTask(ctx context.Context, in *DeleteNodeTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 批量删除任务
	DeleteNodeTasks(ctx context.Context, in *DeleteNodeTasksRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算正在执行的任务数量
	CountDoingNodeTasks(ctx context.Context, in *CountDoingNodeTasksRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 查找需要通知的任务
	FindNotifyingNodeTasks(ctx context.Context, in *FindNotifyingNodeTasksRequest, opts ...grpc.CallOption) (*FindNotifyingNodeTasksResponse, error)
	// 设置任务已通知
	UpdateNodeTasksNotified(ctx context.Context, in *UpdateNodeTasksNotifiedRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type nodeTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeTaskServiceClient(cc grpc.ClientConnInterface) NodeTaskServiceClient {
	return &nodeTaskServiceClient{cc}
}

func (c *nodeTaskServiceClient) FindNodeTasks(ctx context.Context, in *FindNodeTasksRequest, opts ...grpc.CallOption) (*FindNodeTasksResponse, error) {
	out := new(FindNodeTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/findNodeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) ReportNodeTaskDone(ctx context.Context, in *ReportNodeTaskDoneRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/reportNodeTaskDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) FindNodeClusterTasks(ctx context.Context, in *FindNodeClusterTasksRequest, opts ...grpc.CallOption) (*FindNodeClusterTasksResponse, error) {
	out := new(FindNodeClusterTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/findNodeClusterTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) ExistsNodeTasks(ctx context.Context, in *ExistsNodeTasksRequest, opts ...grpc.CallOption) (*ExistsNodeTasksResponse, error) {
	out := new(ExistsNodeTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/existsNodeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) DeleteNodeTask(ctx context.Context, in *DeleteNodeTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/deleteNodeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) DeleteNodeTasks(ctx context.Context, in *DeleteNodeTasksRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/deleteNodeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) CountDoingNodeTasks(ctx context.Context, in *CountDoingNodeTasksRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/countDoingNodeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) FindNotifyingNodeTasks(ctx context.Context, in *FindNotifyingNodeTasksRequest, opts ...grpc.CallOption) (*FindNotifyingNodeTasksResponse, error) {
	out := new(FindNotifyingNodeTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/findNotifyingNodeTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeTaskServiceClient) UpdateNodeTasksNotified(ctx context.Context, in *UpdateNodeTasksNotifiedRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeTaskService/updateNodeTasksNotified", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeTaskServiceServer is the server API for NodeTaskService service.
// All implementations must embed UnimplementedNodeTaskServiceServer
// for forward compatibility
type NodeTaskServiceServer interface {
	// 获取单节点同步任务
	FindNodeTasks(context.Context, *FindNodeTasksRequest) (*FindNodeTasksResponse, error)
	// 报告同步任务结果
	ReportNodeTaskDone(context.Context, *ReportNodeTaskDoneRequest) (*RPCSuccess, error)
	// 获取所有正在同步的集群信息
	FindNodeClusterTasks(context.Context, *FindNodeClusterTasksRequest) (*FindNodeClusterTasksResponse, error)
	// 检查是否有正在执行的任务
	ExistsNodeTasks(context.Context, *ExistsNodeTasksRequest) (*ExistsNodeTasksResponse, error)
	// 删除任务
	DeleteNodeTask(context.Context, *DeleteNodeTaskRequest) (*RPCSuccess, error)
	// 批量删除任务
	DeleteNodeTasks(context.Context, *DeleteNodeTasksRequest) (*RPCSuccess, error)
	// 计算正在执行的任务数量
	CountDoingNodeTasks(context.Context, *CountDoingNodeTasksRequest) (*RPCCountResponse, error)
	// 查找需要通知的任务
	FindNotifyingNodeTasks(context.Context, *FindNotifyingNodeTasksRequest) (*FindNotifyingNodeTasksResponse, error)
	// 设置任务已通知
	UpdateNodeTasksNotified(context.Context, *UpdateNodeTasksNotifiedRequest) (*RPCSuccess, error)
}

// UnimplementedNodeTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeTaskServiceServer struct {
}

func (UnimplementedNodeTaskServiceServer) FindNodeTasks(context.Context, *FindNodeTasksRequest) (*FindNodeTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNodeTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) ReportNodeTaskDone(context.Context, *ReportNodeTaskDoneRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportNodeTaskDone not implemented")
}
func (UnimplementedNodeTaskServiceServer) FindNodeClusterTasks(context.Context, *FindNodeClusterTasksRequest) (*FindNodeClusterTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNodeClusterTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) ExistsNodeTasks(context.Context, *ExistsNodeTasksRequest) (*ExistsNodeTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsNodeTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) DeleteNodeTask(context.Context, *DeleteNodeTaskRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeTask not implemented")
}
func (UnimplementedNodeTaskServiceServer) DeleteNodeTasks(context.Context, *DeleteNodeTasksRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNodeTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) CountDoingNodeTasks(context.Context, *CountDoingNodeTasksRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDoingNodeTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) FindNotifyingNodeTasks(context.Context, *FindNotifyingNodeTasksRequest) (*FindNotifyingNodeTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNotifyingNodeTasks not implemented")
}
func (UnimplementedNodeTaskServiceServer) UpdateNodeTasksNotified(context.Context, *UpdateNodeTasksNotifiedRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeTasksNotified not implemented")
}
func RegisterNodeTaskServiceServer(s grpc.ServiceRegistrar, srv NodeTaskServiceServer) {
	s.RegisterService(&NodeTaskService_ServiceDesc, srv)
}

func _NodeTaskService_FindNodeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).FindNodeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/findNodeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).FindNodeTasks(ctx, req.(*FindNodeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_ReportNodeTaskDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportNodeTaskDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).ReportNodeTaskDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/reportNodeTaskDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).ReportNodeTaskDone(ctx, req.(*ReportNodeTaskDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_FindNodeClusterTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodeClusterTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).FindNodeClusterTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/findNodeClusterTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).FindNodeClusterTasks(ctx, req.(*FindNodeClusterTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_ExistsNodeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsNodeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).ExistsNodeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/existsNodeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).ExistsNodeTasks(ctx, req.(*ExistsNodeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_DeleteNodeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).DeleteNodeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/deleteNodeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).DeleteNodeTask(ctx, req.(*DeleteNodeTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_DeleteNodeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).DeleteNodeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/deleteNodeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).DeleteNodeTasks(ctx, req.(*DeleteNodeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_CountDoingNodeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountDoingNodeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).CountDoingNodeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/countDoingNodeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).CountDoingNodeTasks(ctx, req.(*CountDoingNodeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_FindNotifyingNodeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNotifyingNodeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).FindNotifyingNodeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/findNotifyingNodeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).FindNotifyingNodeTasks(ctx, req.(*FindNotifyingNodeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeTaskService_UpdateNodeTasksNotified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeTasksNotifiedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeTaskServiceServer).UpdateNodeTasksNotified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeTaskService/updateNodeTasksNotified",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeTaskServiceServer).UpdateNodeTasksNotified(ctx, req.(*UpdateNodeTasksNotifiedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeTaskService_ServiceDesc is the grpc.ServiceDesc for NodeTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeTaskService",
	HandlerType: (*NodeTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findNodeTasks",
			Handler:    _NodeTaskService_FindNodeTasks_Handler,
		},
		{
			MethodName: "reportNodeTaskDone",
			Handler:    _NodeTaskService_ReportNodeTaskDone_Handler,
		},
		{
			MethodName: "findNodeClusterTasks",
			Handler:    _NodeTaskService_FindNodeClusterTasks_Handler,
		},
		{
			MethodName: "existsNodeTasks",
			Handler:    _NodeTaskService_ExistsNodeTasks_Handler,
		},
		{
			MethodName: "deleteNodeTask",
			Handler:    _NodeTaskService_DeleteNodeTask_Handler,
		},
		{
			MethodName: "deleteNodeTasks",
			Handler:    _NodeTaskService_DeleteNodeTasks_Handler,
		},
		{
			MethodName: "countDoingNodeTasks",
			Handler:    _NodeTaskService_CountDoingNodeTasks_Handler,
		},
		{
			MethodName: "findNotifyingNodeTasks",
			Handler:    _NodeTaskService_FindNotifyingNodeTasks_Handler,
		},
		{
			MethodName: "updateNodeTasksNotified",
			Handler:    _NodeTaskService_UpdateNodeTasksNotified_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_task.proto",
}