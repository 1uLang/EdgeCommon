// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_acme_provider_account.proto

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

type ACMEProviderAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn         bool          `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
	ProviderCode string        `protobuf:"bytes,4,opt,name=providerCode,proto3" json:"providerCode,omitempty"`
	EabKid       string        `protobuf:"bytes,5,opt,name=eabKid,proto3" json:"eabKid,omitempty"`
	EabKey       string        `protobuf:"bytes,6,opt,name=eabKey,proto3" json:"eabKey,omitempty"`
	Error        string        `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	AcmeProvider *ACMEProvider `protobuf:"bytes,30,opt,name=acmeProvider,proto3" json:"acmeProvider,omitempty"`
}

func (x *ACMEProviderAccount) Reset() {
	*x = ACMEProviderAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_acme_provider_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACMEProviderAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACMEProviderAccount) ProtoMessage() {}

func (x *ACMEProviderAccount) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_acme_provider_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACMEProviderAccount.ProtoReflect.Descriptor instead.
func (*ACMEProviderAccount) Descriptor() ([]byte, []int) {
	return file_models_model_acme_provider_account_proto_rawDescGZIP(), []int{0}
}

func (x *ACMEProviderAccount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ACMEProviderAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ACMEProviderAccount) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *ACMEProviderAccount) GetProviderCode() string {
	if x != nil {
		return x.ProviderCode
	}
	return ""
}

func (x *ACMEProviderAccount) GetEabKid() string {
	if x != nil {
		return x.EabKid
	}
	return ""
}

func (x *ACMEProviderAccount) GetEabKey() string {
	if x != nil {
		return x.EabKey
	}
	return ""
}

func (x *ACMEProviderAccount) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ACMEProviderAccount) GetAcmeProvider() *ACMEProvider {
	if x != nil {
		return x.AcmeProvider
	}
	return nil
}

var File_models_model_acme_provider_account_proto protoreflect.FileDescriptor

var file_models_model_acme_provider_account_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61,
	0x63, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x20,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x63, 0x6d,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xed, 0x01, 0x0a, 0x13, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x61, 0x62, 0x4b, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x61, 0x62, 0x4b, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x61, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x61,
	0x62, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x0c, 0x61, 0x63,
	0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0c, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_acme_provider_account_proto_rawDescOnce sync.Once
	file_models_model_acme_provider_account_proto_rawDescData = file_models_model_acme_provider_account_proto_rawDesc
)

func file_models_model_acme_provider_account_proto_rawDescGZIP() []byte {
	file_models_model_acme_provider_account_proto_rawDescOnce.Do(func() {
		file_models_model_acme_provider_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_acme_provider_account_proto_rawDescData)
	})
	return file_models_model_acme_provider_account_proto_rawDescData
}

var file_models_model_acme_provider_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_acme_provider_account_proto_goTypes = []interface{}{
	(*ACMEProviderAccount)(nil), // 0: pb.ACMEProviderAccount
	(*ACMEProvider)(nil),        // 1: pb.ACMEProvider
}
var file_models_model_acme_provider_account_proto_depIdxs = []int32{
	1, // 0: pb.ACMEProviderAccount.acmeProvider:type_name -> pb.ACMEProvider
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_acme_provider_account_proto_init() }
func file_models_model_acme_provider_account_proto_init() {
	if File_models_model_acme_provider_account_proto != nil {
		return
	}
	file_models_model_acme_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_acme_provider_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACMEProviderAccount); i {
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
			RawDescriptor: file_models_model_acme_provider_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_acme_provider_account_proto_goTypes,
		DependencyIndexes: file_models_model_acme_provider_account_proto_depIdxs,
		MessageInfos:      file_models_model_acme_provider_account_proto_msgTypes,
	}.Build()
	File_models_model_acme_provider_account_proto = out.File
	file_models_model_acme_provider_account_proto_rawDesc = nil
	file_models_model_acme_provider_account_proto_goTypes = nil
	file_models_model_acme_provider_account_proto_depIdxs = nil
}
