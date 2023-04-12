// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_node_threshold.proto

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

type NodeThreshold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId      int64  `protobuf:"varint,2,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	Node           *Node  `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	Item           string `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	Param          string `protobuf:"bytes,5,opt,name=param,proto3" json:"param,omitempty"`
	Operator       string `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	ValueJSON      []byte `protobuf:"bytes,7,opt,name=valueJSON,proto3" json:"valueJSON,omitempty"`
	Message        string `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	Duration       int32  `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
	DurationUnit   string `protobuf:"bytes,10,opt,name=durationUnit,proto3" json:"durationUnit,omitempty"`
	SumMethod      string `protobuf:"bytes,11,opt,name=sumMethod,proto3" json:"sumMethod,omitempty"`
	IsOn           bool   `protobuf:"varint,12,opt,name=isOn,proto3" json:"isOn,omitempty"`
	NotifyDuration int32  `protobuf:"varint,13,opt,name=notifyDuration,proto3" json:"notifyDuration,omitempty"`
}

func (x *NodeThreshold) Reset() {
	*x = NodeThreshold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_node_threshold_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeThreshold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeThreshold) ProtoMessage() {}

func (x *NodeThreshold) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_node_threshold_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeThreshold.ProtoReflect.Descriptor instead.
func (*NodeThreshold) Descriptor() ([]byte, []int) {
	return file_models_model_node_threshold_proto_rawDescGZIP(), []int{0}
}

func (x *NodeThreshold) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeThreshold) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *NodeThreshold) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *NodeThreshold) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *NodeThreshold) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *NodeThreshold) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *NodeThreshold) GetValueJSON() []byte {
	if x != nil {
		return x.ValueJSON
	}
	return nil
}

func (x *NodeThreshold) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NodeThreshold) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *NodeThreshold) GetDurationUnit() string {
	if x != nil {
		return x.DurationUnit
	}
	return ""
}

func (x *NodeThreshold) GetSumMethod() string {
	if x != nil {
		return x.SumMethod
	}
	return ""
}

func (x *NodeThreshold) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NodeThreshold) GetNotifyDuration() int32 {
	if x != nil {
		return x.NotifyDuration
	}
	return 0
}

var File_models_model_node_threshold_proto protoreflect.FileDescriptor

var file_models_model_node_threshold_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf3, 0x02, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4a, 0x53,
	0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x75, 0x6d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x75, 0x6d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_node_threshold_proto_rawDescOnce sync.Once
	file_models_model_node_threshold_proto_rawDescData = file_models_model_node_threshold_proto_rawDesc
)

func file_models_model_node_threshold_proto_rawDescGZIP() []byte {
	file_models_model_node_threshold_proto_rawDescOnce.Do(func() {
		file_models_model_node_threshold_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_node_threshold_proto_rawDescData)
	})
	return file_models_model_node_threshold_proto_rawDescData
}

var file_models_model_node_threshold_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_node_threshold_proto_goTypes = []interface{}{
	(*NodeThreshold)(nil), // 0: pb.NodeThreshold
	(*Node)(nil),          // 1: pb.Node
}
var file_models_model_node_threshold_proto_depIdxs = []int32{
	1, // 0: pb.NodeThreshold.node:type_name -> pb.Node
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_node_threshold_proto_init() }
func file_models_model_node_threshold_proto_init() {
	if File_models_model_node_threshold_proto != nil {
		return
	}
	file_models_model_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_node_threshold_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeThreshold); i {
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
			RawDescriptor: file_models_model_node_threshold_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_node_threshold_proto_goTypes,
		DependencyIndexes: file_models_model_node_threshold_proto_depIdxs,
		MessageInfos:      file_models_model_node_threshold_proto_msgTypes,
	}.Build()
	File_models_model_node_threshold_proto = out.File
	file_models_model_node_threshold_proto_rawDesc = nil
	file_models_model_node_threshold_proto_goTypes = nil
	file_models_model_node_threshold_proto_depIdxs = nil
}
