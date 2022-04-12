// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: service_node_value.proto

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

// 记录数据
type CreateNodeValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item      string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	ValueJSON []byte `protobuf:"bytes,2,opt,name=valueJSON,proto3" json:"valueJSON,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreateNodeValueRequest) Reset() {
	*x = CreateNodeValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeValueRequest) ProtoMessage() {}

func (x *CreateNodeValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeValueRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeValueRequest) Descriptor() ([]byte, []int) {
	return file_service_node_value_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeValueRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *CreateNodeValueRequest) GetValueJSON() []byte {
	if x != nil {
		return x.ValueJSON
	}
	return nil
}

func (x *CreateNodeValueRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// 读取数据
type ListNodeValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role   string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	NodeId int64  `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Item   string `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	Range  string `protobuf:"bytes,10,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *ListNodeValuesRequest) Reset() {
	*x = ListNodeValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeValuesRequest) ProtoMessage() {}

func (x *ListNodeValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeValuesRequest.ProtoReflect.Descriptor instead.
func (*ListNodeValuesRequest) Descriptor() ([]byte, []int) {
	return file_service_node_value_proto_rawDescGZIP(), []int{1}
}

func (x *ListNodeValuesRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ListNodeValuesRequest) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *ListNodeValuesRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *ListNodeValuesRequest) GetRange() string {
	if x != nil {
		return x.Range
	}
	return ""
}

type ListNodeValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeValues []*NodeValue `protobuf:"bytes,1,rep,name=nodeValues,proto3" json:"nodeValues,omitempty"`
}

func (x *ListNodeValuesResponse) Reset() {
	*x = ListNodeValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeValuesResponse) ProtoMessage() {}

func (x *ListNodeValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeValuesResponse.ProtoReflect.Descriptor instead.
func (*ListNodeValuesResponse) Descriptor() ([]byte, []int) {
	return file_service_node_value_proto_rawDescGZIP(), []int{2}
}

func (x *ListNodeValuesResponse) GetNodeValues() []*NodeValue {
	if x != nil {
		return x.NodeValues
	}
	return nil
}

var File_service_node_value_proto protoreflect.FileDescriptor

var file_service_node_value_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x6d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x22, 0x47, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x6e,
	0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x6e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x10, 0x4e,
	0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x47,
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_value_proto_rawDescOnce sync.Once
	file_service_node_value_proto_rawDescData = file_service_node_value_proto_rawDesc
)

func file_service_node_value_proto_rawDescGZIP() []byte {
	file_service_node_value_proto_rawDescOnce.Do(func() {
		file_service_node_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_value_proto_rawDescData)
	})
	return file_service_node_value_proto_rawDescData
}

var file_service_node_value_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_node_value_proto_goTypes = []interface{}{
	(*CreateNodeValueRequest)(nil), // 0: pb.CreateNodeValueRequest
	(*ListNodeValuesRequest)(nil),  // 1: pb.ListNodeValuesRequest
	(*ListNodeValuesResponse)(nil), // 2: pb.ListNodeValuesResponse
	(*NodeValue)(nil),              // 3: pb.NodeValue
	(*RPCSuccess)(nil),             // 4: pb.RPCSuccess
}
var file_service_node_value_proto_depIdxs = []int32{
	3, // 0: pb.ListNodeValuesResponse.nodeValues:type_name -> pb.NodeValue
	0, // 1: pb.NodeValueService.createNodeValue:input_type -> pb.CreateNodeValueRequest
	1, // 2: pb.NodeValueService.listNodeValues:input_type -> pb.ListNodeValuesRequest
	4, // 3: pb.NodeValueService.createNodeValue:output_type -> pb.RPCSuccess
	2, // 4: pb.NodeValueService.listNodeValues:output_type -> pb.ListNodeValuesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_node_value_proto_init() }
func file_service_node_value_proto_init() {
	if File_service_node_value_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_node_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeValueRequest); i {
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
		file_service_node_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeValuesRequest); i {
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
		file_service_node_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeValuesResponse); i {
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
			RawDescriptor: file_service_node_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_value_proto_goTypes,
		DependencyIndexes: file_service_node_value_proto_depIdxs,
		MessageInfos:      file_service_node_value_proto_msgTypes,
	}.Build()
	File_service_node_value_proto = out.File
	file_service_node_value_proto_rawDesc = nil
	file_service_node_value_proto_goTypes = nil
	file_service_node_value_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeValueServiceClient is the client API for NodeValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeValueServiceClient interface {
	// 记录数据
	CreateNodeValue(ctx context.Context, in *CreateNodeValueRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取数据
	ListNodeValues(ctx context.Context, in *ListNodeValuesRequest, opts ...grpc.CallOption) (*ListNodeValuesResponse, error)
}

type nodeValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeValueServiceClient(cc grpc.ClientConnInterface) NodeValueServiceClient {
	return &nodeValueServiceClient{cc}
}

func (c *nodeValueServiceClient) CreateNodeValue(ctx context.Context, in *CreateNodeValueRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeValueService/createNodeValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeValueServiceClient) ListNodeValues(ctx context.Context, in *ListNodeValuesRequest, opts ...grpc.CallOption) (*ListNodeValuesResponse, error) {
	out := new(ListNodeValuesResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeValueService/listNodeValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeValueServiceServer is the server API for NodeValueService service.
type NodeValueServiceServer interface {
	// 记录数据
	CreateNodeValue(context.Context, *CreateNodeValueRequest) (*RPCSuccess, error)
	// 读取数据
	ListNodeValues(context.Context, *ListNodeValuesRequest) (*ListNodeValuesResponse, error)
}

// UnimplementedNodeValueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeValueServiceServer struct {
}

func (*UnimplementedNodeValueServiceServer) CreateNodeValue(context.Context, *CreateNodeValueRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeValue not implemented")
}
func (*UnimplementedNodeValueServiceServer) ListNodeValues(context.Context, *ListNodeValuesRequest) (*ListNodeValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodeValues not implemented")
}

func RegisterNodeValueServiceServer(s *grpc.Server, srv NodeValueServiceServer) {
	s.RegisterService(&_NodeValueService_serviceDesc, srv)
}

func _NodeValueService_CreateNodeValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeValueServiceServer).CreateNodeValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeValueService/CreateNodeValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeValueServiceServer).CreateNodeValue(ctx, req.(*CreateNodeValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeValueService_ListNodeValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeValueServiceServer).ListNodeValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeValueService/ListNodeValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeValueServiceServer).ListNodeValues(ctx, req.(*ListNodeValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeValueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeValueService",
	HandlerType: (*NodeValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createNodeValue",
			Handler:    _NodeValueService_CreateNodeValue_Handler,
		},
		{
			MethodName: "listNodeValues",
			Handler:    _NodeValueService_ListNodeValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_value.proto",
}
