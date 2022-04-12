// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: models/model_acme_provider.proto

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

type ACMEProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code           string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description    string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ApiURL         string `protobuf:"bytes,5,opt,name=apiURL,proto3" json:"apiURL,omitempty"`
	RequireEAB     bool   `protobuf:"varint,6,opt,name=requireEAB,proto3" json:"requireEAB,omitempty"`
	EabDescription string `protobuf:"bytes,7,opt,name=eabDescription,proto3" json:"eabDescription,omitempty"`
}

func (x *ACMEProvider) Reset() {
	*x = ACMEProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_acme_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACMEProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACMEProvider) ProtoMessage() {}

func (x *ACMEProvider) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_acme_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACMEProvider.ProtoReflect.Descriptor instead.
func (*ACMEProvider) Descriptor() ([]byte, []int) {
	return file_models_model_acme_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ACMEProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ACMEProvider) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ACMEProvider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ACMEProvider) GetApiURL() string {
	if x != nil {
		return x.ApiURL
	}
	return ""
}

func (x *ACMEProvider) GetRequireEAB() bool {
	if x != nil {
		return x.RequireEAB
	}
	return false
}

func (x *ACMEProvider) GetEabDescription() string {
	if x != nil {
		return x.EabDescription
	}
	return ""
}

var File_models_model_acme_provider_proto protoreflect.FileDescriptor

var file_models_model_acme_provider_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x63, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb8, 0x01, 0x0a, 0x0c, 0x41, 0x43, 0x4d, 0x45, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x45, 0x41, 0x42, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x45, 0x41, 0x42, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x61, 0x62,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x65, 0x61, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_models_model_acme_provider_proto_rawDescOnce sync.Once
	file_models_model_acme_provider_proto_rawDescData = file_models_model_acme_provider_proto_rawDesc
)

func file_models_model_acme_provider_proto_rawDescGZIP() []byte {
	file_models_model_acme_provider_proto_rawDescOnce.Do(func() {
		file_models_model_acme_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_acme_provider_proto_rawDescData)
	})
	return file_models_model_acme_provider_proto_rawDescData
}

var file_models_model_acme_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_acme_provider_proto_goTypes = []interface{}{
	(*ACMEProvider)(nil), // 0: pb.ACMEProvider
}
var file_models_model_acme_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_acme_provider_proto_init() }
func file_models_model_acme_provider_proto_init() {
	if File_models_model_acme_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_acme_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACMEProvider); i {
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
			RawDescriptor: file_models_model_acme_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_acme_provider_proto_goTypes,
		DependencyIndexes: file_models_model_acme_provider_proto_depIdxs,
		MessageInfos:      file_models_model_acme_provider_proto_msgTypes,
	}.Build()
	File_models_model_acme_provider_proto = out.File
	file_models_model_acme_provider_proto_rawDesc = nil
	file_models_model_acme_provider_proto_goTypes = nil
	file_models_model_acme_provider_proto_depIdxs = nil
}
