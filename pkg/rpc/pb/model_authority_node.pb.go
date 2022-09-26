// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_authority_node.proto

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

type AuthorityNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn        bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	UniqueId    string `protobuf:"bytes,3,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret      string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Name        string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	StatusJSON  []byte `protobuf:"bytes,7,opt,name=statusJSON,proto3" json:"statusJSON,omitempty"`
}

func (x *AuthorityNode) Reset() {
	*x = AuthorityNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_authority_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorityNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorityNode) ProtoMessage() {}

func (x *AuthorityNode) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_authority_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorityNode.ProtoReflect.Descriptor instead.
func (*AuthorityNode) Descriptor() ([]byte, []int) {
	return file_models_model_authority_node_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorityNode) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorityNode) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *AuthorityNode) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *AuthorityNode) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *AuthorityNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorityNode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthorityNode) GetStatusJSON() []byte {
	if x != nil {
		return x.StatusJSON
	}
	return nil
}

var File_models_model_authority_node_proto protoreflect.FileDescriptor

var file_models_model_authority_node_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_authority_node_proto_rawDescOnce sync.Once
	file_models_model_authority_node_proto_rawDescData = file_models_model_authority_node_proto_rawDesc
)

func file_models_model_authority_node_proto_rawDescGZIP() []byte {
	file_models_model_authority_node_proto_rawDescOnce.Do(func() {
		file_models_model_authority_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_authority_node_proto_rawDescData)
	})
	return file_models_model_authority_node_proto_rawDescData
}

var file_models_model_authority_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_authority_node_proto_goTypes = []interface{}{
	(*AuthorityNode)(nil), // 0: pb.AuthorityNode
}
var file_models_model_authority_node_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_authority_node_proto_init() }
func file_models_model_authority_node_proto_init() {
	if File_models_model_authority_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_authority_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorityNode); i {
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
			RawDescriptor: file_models_model_authority_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_authority_node_proto_goTypes,
		DependencyIndexes: file_models_model_authority_node_proto_depIdxs,
		MessageInfos:      file_models_model_authority_node_proto_msgTypes,
	}.Build()
	File_models_model_authority_node_proto = out.File
	file_models_model_authority_node_proto_rawDesc = nil
	file_models_model_authority_node_proto_goTypes = nil
	file_models_model_authority_node_proto_depIdxs = nil
}
