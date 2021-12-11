// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service_node_login.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 读取建议的端口
type FindNodeLoginSuggestPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *FindNodeLoginSuggestPortsRequest) Reset() {
	*x = FindNodeLoginSuggestPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeLoginSuggestPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeLoginSuggestPortsRequest) ProtoMessage() {}

func (x *FindNodeLoginSuggestPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeLoginSuggestPortsRequest.ProtoReflect.Descriptor instead.
func (*FindNodeLoginSuggestPortsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_login_proto_rawDescGZIP(), []int{0}
}

func (x *FindNodeLoginSuggestPortsRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type FindNodeLoginSuggestPortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports          []int32 `protobuf:"varint,1,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	AvailablePorts []int32 `protobuf:"varint,2,rep,packed,name=availablePorts,proto3" json:"availablePorts,omitempty"`
}

func (x *FindNodeLoginSuggestPortsResponse) Reset() {
	*x = FindNodeLoginSuggestPortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeLoginSuggestPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeLoginSuggestPortsResponse) ProtoMessage() {}

func (x *FindNodeLoginSuggestPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeLoginSuggestPortsResponse.ProtoReflect.Descriptor instead.
func (*FindNodeLoginSuggestPortsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_login_proto_rawDescGZIP(), []int{1}
}

func (x *FindNodeLoginSuggestPortsResponse) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *FindNodeLoginSuggestPortsResponse) GetAvailablePorts() []int32 {
	if x != nil {
		return x.AvailablePorts
	}
	return nil
}

var File_service_node_login_proto protoreflect.FileDescriptor

var file_service_node_login_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x36,
	0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x7c, 0x0a, 0x10, 0x4e, 0x6f, 0x64,
	0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a,
	0x19, 0x66, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_login_proto_rawDescOnce sync.Once
	file_service_node_login_proto_rawDescData = file_service_node_login_proto_rawDesc
)

func file_service_node_login_proto_rawDescGZIP() []byte {
	file_service_node_login_proto_rawDescOnce.Do(func() {
		file_service_node_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_login_proto_rawDescData)
	})
	return file_service_node_login_proto_rawDescData
}

var file_service_node_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_node_login_proto_goTypes = []interface{}{
	(*FindNodeLoginSuggestPortsRequest)(nil),  // 0: pb.FindNodeLoginSuggestPortsRequest
	(*FindNodeLoginSuggestPortsResponse)(nil), // 1: pb.FindNodeLoginSuggestPortsResponse
}
var file_service_node_login_proto_depIdxs = []int32{
	0, // 0: pb.NodeLoginService.findNodeLoginSuggestPorts:input_type -> pb.FindNodeLoginSuggestPortsRequest
	1, // 1: pb.NodeLoginService.findNodeLoginSuggestPorts:output_type -> pb.FindNodeLoginSuggestPortsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_node_login_proto_init() }
func file_service_node_login_proto_init() {
	if File_service_node_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_node_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeLoginSuggestPortsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_node_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeLoginSuggestPortsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_node_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_login_proto_goTypes,
		DependencyIndexes: file_service_node_login_proto_depIdxs,
		MessageInfos:      file_service_node_login_proto_msgTypes,
	}.Build()
	File_service_node_login_proto = out.File
	file_service_node_login_proto_rawDesc = nil
	file_service_node_login_proto_goTypes = nil
	file_service_node_login_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeLoginServiceClient is the client API for NodeLoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeLoginServiceClient interface {
	// 读取建议的端口
	FindNodeLoginSuggestPorts(ctx context.Context, in *FindNodeLoginSuggestPortsRequest, opts ...grpc.CallOption) (*FindNodeLoginSuggestPortsResponse, error)
}

type nodeLoginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeLoginServiceClient(cc grpc.ClientConnInterface) NodeLoginServiceClient {
	return &nodeLoginServiceClient{cc}
}

func (c *nodeLoginServiceClient) FindNodeLoginSuggestPorts(ctx context.Context, in *FindNodeLoginSuggestPortsRequest, opts ...grpc.CallOption) (*FindNodeLoginSuggestPortsResponse, error) {
	out := new(FindNodeLoginSuggestPortsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeLoginService/findNodeLoginSuggestPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeLoginServiceServer is the server API for NodeLoginService service.
type NodeLoginServiceServer interface {
	// 读取建议的端口
	FindNodeLoginSuggestPorts(context.Context, *FindNodeLoginSuggestPortsRequest) (*FindNodeLoginSuggestPortsResponse, error)
}

// UnimplementedNodeLoginServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeLoginServiceServer struct {
}

func (*UnimplementedNodeLoginServiceServer) FindNodeLoginSuggestPorts(context.Context, *FindNodeLoginSuggestPortsRequest) (*FindNodeLoginSuggestPortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNodeLoginSuggestPorts not implemented")
}

func RegisterNodeLoginServiceServer(s *grpc.Server, srv NodeLoginServiceServer) {
	s.RegisterService(&_NodeLoginService_serviceDesc, srv)
}

func _NodeLoginService_FindNodeLoginSuggestPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodeLoginSuggestPortsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeLoginServiceServer).FindNodeLoginSuggestPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeLoginService/FindNodeLoginSuggestPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeLoginServiceServer).FindNodeLoginSuggestPorts(ctx, req.(*FindNodeLoginSuggestPortsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeLoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeLoginService",
	HandlerType: (*NodeLoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findNodeLoginSuggestPorts",
			Handler:    _NodeLoginService_FindNodeLoginSuggestPorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_login.proto",
}
