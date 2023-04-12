// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_node_ip_address_log.proto

package pb

import (
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

// 计算日志数量
type CountAllNodeIPAddressLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId int64 `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
}

func (x *CountAllNodeIPAddressLogsRequest) Reset() {
	*x = CountAllNodeIPAddressLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllNodeIPAddressLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllNodeIPAddressLogsRequest) ProtoMessage() {}

func (x *CountAllNodeIPAddressLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllNodeIPAddressLogsRequest.ProtoReflect.Descriptor instead.
func (*CountAllNodeIPAddressLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_log_proto_rawDescGZIP(), []int{0}
}

func (x *CountAllNodeIPAddressLogsRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

// 列出单页日志
type ListNodeIPAddressLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressId int64 `protobuf:"varint,1,opt,name=nodeIPAddressId,proto3" json:"nodeIPAddressId,omitempty"`
	Offset          int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size            int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListNodeIPAddressLogsRequest) Reset() {
	*x = ListNodeIPAddressLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeIPAddressLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeIPAddressLogsRequest) ProtoMessage() {}

func (x *ListNodeIPAddressLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeIPAddressLogsRequest.ProtoReflect.Descriptor instead.
func (*ListNodeIPAddressLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_log_proto_rawDescGZIP(), []int{1}
}

func (x *ListNodeIPAddressLogsRequest) GetNodeIPAddressId() int64 {
	if x != nil {
		return x.NodeIPAddressId
	}
	return 0
}

func (x *ListNodeIPAddressLogsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListNodeIPAddressLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListNodeIPAddressLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIPAddressLogs []*NodeIPAddressLog `protobuf:"bytes,1,rep,name=nodeIPAddressLogs,proto3" json:"nodeIPAddressLogs,omitempty"`
}

func (x *ListNodeIPAddressLogsResponse) Reset() {
	*x = ListNodeIPAddressLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_ip_address_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeIPAddressLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeIPAddressLogsResponse) ProtoMessage() {}

func (x *ListNodeIPAddressLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_ip_address_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeIPAddressLogsResponse.ProtoReflect.Descriptor instead.
func (*ListNodeIPAddressLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_ip_address_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListNodeIPAddressLogsResponse) GetNodeIPAddressLogs() []*NodeIPAddressLog {
	if x != nil {
		return x.NodeIPAddressLogs
	}
	return nil
}

var File_service_node_ip_address_log_proto protoreflect.FileDescriptor

var file_service_node_ip_address_log_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x20, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x63,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x52, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c,
	0x6f, 0x67, 0x73, 0x32, 0xd0, 0x01, 0x0a, 0x17, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x57, 0x0a, 0x19, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x24, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_ip_address_log_proto_rawDescOnce sync.Once
	file_service_node_ip_address_log_proto_rawDescData = file_service_node_ip_address_log_proto_rawDesc
)

func file_service_node_ip_address_log_proto_rawDescGZIP() []byte {
	file_service_node_ip_address_log_proto_rawDescOnce.Do(func() {
		file_service_node_ip_address_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_ip_address_log_proto_rawDescData)
	})
	return file_service_node_ip_address_log_proto_rawDescData
}

var file_service_node_ip_address_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_node_ip_address_log_proto_goTypes = []interface{}{
	(*CountAllNodeIPAddressLogsRequest)(nil), // 0: pb.CountAllNodeIPAddressLogsRequest
	(*ListNodeIPAddressLogsRequest)(nil),     // 1: pb.ListNodeIPAddressLogsRequest
	(*ListNodeIPAddressLogsResponse)(nil),    // 2: pb.ListNodeIPAddressLogsResponse
	(*NodeIPAddressLog)(nil),                 // 3: pb.NodeIPAddressLog
	(*RPCCountResponse)(nil),                 // 4: pb.RPCCountResponse
}
var file_service_node_ip_address_log_proto_depIdxs = []int32{
	3, // 0: pb.ListNodeIPAddressLogsResponse.nodeIPAddressLogs:type_name -> pb.NodeIPAddressLog
	0, // 1: pb.NodeIPAddressLogService.countAllNodeIPAddressLogs:input_type -> pb.CountAllNodeIPAddressLogsRequest
	1, // 2: pb.NodeIPAddressLogService.listNodeIPAddressLogs:input_type -> pb.ListNodeIPAddressLogsRequest
	4, // 3: pb.NodeIPAddressLogService.countAllNodeIPAddressLogs:output_type -> pb.RPCCountResponse
	2, // 4: pb.NodeIPAddressLogService.listNodeIPAddressLogs:output_type -> pb.ListNodeIPAddressLogsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_node_ip_address_log_proto_init() }
func file_service_node_ip_address_log_proto_init() {
	if File_service_node_ip_address_log_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_node_ip_address_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_ip_address_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllNodeIPAddressLogsRequest); i {
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
		file_service_node_ip_address_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeIPAddressLogsRequest); i {
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
		file_service_node_ip_address_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeIPAddressLogsResponse); i {
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
			RawDescriptor: file_service_node_ip_address_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_ip_address_log_proto_goTypes,
		DependencyIndexes: file_service_node_ip_address_log_proto_depIdxs,
		MessageInfos:      file_service_node_ip_address_log_proto_msgTypes,
	}.Build()
	File_service_node_ip_address_log_proto = out.File
	file_service_node_ip_address_log_proto_rawDesc = nil
	file_service_node_ip_address_log_proto_goTypes = nil
	file_service_node_ip_address_log_proto_depIdxs = nil
}
