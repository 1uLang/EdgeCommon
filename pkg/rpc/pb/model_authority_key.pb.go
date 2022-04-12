// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: models/model_authority_key.proto

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

// 版本认证
type AuthorityKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DayFrom      string   `protobuf:"bytes,2,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`
	DayTo        string   `protobuf:"bytes,3,opt,name=dayTo,proto3" json:"dayTo,omitempty"`
	Hostname     string   `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	MacAddresses []string `protobuf:"bytes,5,rep,name=macAddresses,proto3" json:"macAddresses,omitempty"`
	UpdatedAt    int64    `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Company      string   `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	Nodes        int32    `protobuf:"varint,8,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *AuthorityKey) Reset() {
	*x = AuthorityKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_authority_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorityKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorityKey) ProtoMessage() {}

func (x *AuthorityKey) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_authority_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorityKey.ProtoReflect.Descriptor instead.
func (*AuthorityKey) Descriptor() ([]byte, []int) {
	return file_models_model_authority_key_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorityKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *AuthorityKey) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *AuthorityKey) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *AuthorityKey) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AuthorityKey) GetMacAddresses() []string {
	if x != nil {
		return x.MacAddresses
	}
	return nil
}

func (x *AuthorityKey) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AuthorityKey) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *AuthorityKey) GetNodes() int32 {
	if x != nil {
		return x.Nodes
	}
	return 0
}

var File_models_model_authority_key_proto protoreflect.FileDescriptor

var file_models_model_authority_key_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe2, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_authority_key_proto_rawDescOnce sync.Once
	file_models_model_authority_key_proto_rawDescData = file_models_model_authority_key_proto_rawDesc
)

func file_models_model_authority_key_proto_rawDescGZIP() []byte {
	file_models_model_authority_key_proto_rawDescOnce.Do(func() {
		file_models_model_authority_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_authority_key_proto_rawDescData)
	})
	return file_models_model_authority_key_proto_rawDescData
}

var file_models_model_authority_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_authority_key_proto_goTypes = []interface{}{
	(*AuthorityKey)(nil), // 0: pb.AuthorityKey
}
var file_models_model_authority_key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_authority_key_proto_init() }
func file_models_model_authority_key_proto_init() {
	if File_models_model_authority_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_authority_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorityKey); i {
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
			RawDescriptor: file_models_model_authority_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_authority_key_proto_goTypes,
		DependencyIndexes: file_models_model_authority_key_proto_depIdxs,
		MessageInfos:      file_models_model_authority_key_proto_msgTypes,
	}.Build()
	File_models_model_authority_key_proto = out.File
	file_models_model_authority_key_proto_rawDesc = nil
	file_models_model_authority_key_proto_goTypes = nil
	file_models_model_authority_key_proto_depIdxs = nil
}
