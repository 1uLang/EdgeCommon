// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: models/model_dns_provider.proto

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

type DNSProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TypeName      string `protobuf:"bytes,4,opt,name=typeName,proto3" json:"typeName,omitempty"`
	ApiParamsJSON []byte `protobuf:"bytes,5,opt,name=apiParamsJSON,proto3" json:"apiParamsJSON,omitempty"`
	DataUpdatedAt int64  `protobuf:"varint,6,opt,name=dataUpdatedAt,proto3" json:"dataUpdatedAt,omitempty"`
}

func (x *DNSProvider) Reset() {
	*x = DNSProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_dns_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSProvider) ProtoMessage() {}

func (x *DNSProvider) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_dns_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSProvider.ProtoReflect.Descriptor instead.
func (*DNSProvider) Descriptor() ([]byte, []int) {
	return file_models_model_dns_provider_proto_rawDescGZIP(), []int{0}
}

func (x *DNSProvider) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DNSProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSProvider) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DNSProvider) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *DNSProvider) GetApiParamsJSON() []byte {
	if x != nil {
		return x.ApiParamsJSON
	}
	return nil
}

func (x *DNSProvider) GetDataUpdatedAt() int64 {
	if x != nil {
		return x.DataUpdatedAt
	}
	return 0
}

var File_models_model_dns_provider_proto protoreflect.FileDescriptor

var file_models_model_dns_provider_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64,
	0x6e, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xad, 0x01, 0x0a, 0x0b, 0x44, 0x4e, 0x53, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x69,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x61, 0x70, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_dns_provider_proto_rawDescOnce sync.Once
	file_models_model_dns_provider_proto_rawDescData = file_models_model_dns_provider_proto_rawDesc
)

func file_models_model_dns_provider_proto_rawDescGZIP() []byte {
	file_models_model_dns_provider_proto_rawDescOnce.Do(func() {
		file_models_model_dns_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_dns_provider_proto_rawDescData)
	})
	return file_models_model_dns_provider_proto_rawDescData
}

var file_models_model_dns_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_dns_provider_proto_goTypes = []interface{}{
	(*DNSProvider)(nil), // 0: pb.DNSProvider
}
var file_models_model_dns_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_dns_provider_proto_init() }
func file_models_model_dns_provider_proto_init() {
	if File_models_model_dns_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_dns_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSProvider); i {
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
			RawDescriptor: file_models_model_dns_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_dns_provider_proto_goTypes,
		DependencyIndexes: file_models_model_dns_provider_proto_depIdxs,
		MessageInfos:      file_models_model_dns_provider_proto_msgTypes,
	}.Build()
	File_models_model_dns_provider_proto = out.File
	file_models_model_dns_provider_proto_rawDesc = nil
	file_models_model_dns_provider_proto_goTypes = nil
	file_models_model_dns_provider_proto_depIdxs = nil
}
