// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_ns_key.proto

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

// NS密钥
type NSKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn       bool      `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name       string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Algo       string    `protobuf:"bytes,4,opt,name=algo,proto3" json:"algo,omitempty"`
	Secret     string    `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	SecretType string    `protobuf:"bytes,6,opt,name=secretType,proto3" json:"secretType,omitempty"`
	IsDeleted  bool      `protobuf:"varint,7,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Version    int64     `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	NsDomain   *NSDomain `protobuf:"bytes,30,opt,name=nsDomain,proto3" json:"nsDomain,omitempty"`
	NsZone     *NSZone   `protobuf:"bytes,31,opt,name=nsZone,proto3" json:"nsZone,omitempty"`
}

func (x *NSKey) Reset() {
	*x = NSKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSKey) ProtoMessage() {}

func (x *NSKey) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSKey.ProtoReflect.Descriptor instead.
func (*NSKey) Descriptor() ([]byte, []int) {
	return file_models_model_ns_key_proto_rawDescGZIP(), []int{0}
}

func (x *NSKey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSKey) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSKey) GetAlgo() string {
	if x != nil {
		return x.Algo
	}
	return ""
}

func (x *NSKey) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *NSKey) GetSecretType() string {
	if x != nil {
		return x.SecretType
	}
	return ""
}

func (x *NSKey) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *NSKey) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NSKey) GetNsDomain() *NSDomain {
	if x != nil {
		return x.NsDomain
	}
	return nil
}

func (x *NSKey) GetNsZone() *NSZone {
	if x != nil {
		return x.NsZone
	}
	return nil
}

var File_models_model_ns_key_proto protoreflect.FileDescriptor

var file_models_model_ns_key_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x73,
	0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x73, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x05, 0x4e, 0x53,
	0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x6c, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x73, 0x5a,
	0x6f, 0x6e, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x53, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x06, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_key_proto_rawDescOnce sync.Once
	file_models_model_ns_key_proto_rawDescData = file_models_model_ns_key_proto_rawDesc
)

func file_models_model_ns_key_proto_rawDescGZIP() []byte {
	file_models_model_ns_key_proto_rawDescOnce.Do(func() {
		file_models_model_ns_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_key_proto_rawDescData)
	})
	return file_models_model_ns_key_proto_rawDescData
}

var file_models_model_ns_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_key_proto_goTypes = []interface{}{
	(*NSKey)(nil),    // 0: pb.NSKey
	(*NSDomain)(nil), // 1: pb.NSDomain
	(*NSZone)(nil),   // 2: pb.NSZone
}
var file_models_model_ns_key_proto_depIdxs = []int32{
	1, // 0: pb.NSKey.nsDomain:type_name -> pb.NSDomain
	2, // 1: pb.NSKey.nsZone:type_name -> pb.NSZone
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_ns_key_proto_init() }
func file_models_model_ns_key_proto_init() {
	if File_models_model_ns_key_proto != nil {
		return
	}
	file_models_model_ns_domain_proto_init()
	file_models_model_ns_zone_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSKey); i {
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
			RawDescriptor: file_models_model_ns_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_key_proto_goTypes,
		DependencyIndexes: file_models_model_ns_key_proto_depIdxs,
		MessageInfos:      file_models_model_ns_key_proto_msgTypes,
	}.Build()
	File_models_model_ns_key_proto = out.File
	file_models_model_ns_key_proto_rawDesc = nil
	file_models_model_ns_key_proto_goTypes = nil
	file_models_model_ns_key_proto_depIdxs = nil
}
