// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: models/model_ssl_cert.proto

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

type SSLCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn        bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TimeBeginAt int64  `protobuf:"varint,4,opt,name=timeBeginAt,proto3" json:"timeBeginAt,omitempty"`
	TimeEndAt   int64  `protobuf:"varint,5,opt,name=timeEndAt,proto3" json:"timeEndAt,omitempty"`
}

func (x *SSLCert) Reset() {
	*x = SSLCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ssl_cert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSLCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSLCert) ProtoMessage() {}

func (x *SSLCert) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ssl_cert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSLCert.ProtoReflect.Descriptor instead.
func (*SSLCert) Descriptor() ([]byte, []int) {
	return file_models_model_ssl_cert_proto_rawDescGZIP(), []int{0}
}

func (x *SSLCert) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SSLCert) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *SSLCert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSLCert) GetTimeBeginAt() int64 {
	if x != nil {
		return x.TimeBeginAt
	}
	return 0
}

func (x *SSLCert) GetTimeEndAt() int64 {
	if x != nil {
		return x.TimeEndAt
	}
	return 0
}

var File_models_model_ssl_cert_proto protoreflect.FileDescriptor

var file_models_model_ssl_cert_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x73, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x53, 0x53, 0x4c, 0x43, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x45,
	0x6e, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x45, 0x6e, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ssl_cert_proto_rawDescOnce sync.Once
	file_models_model_ssl_cert_proto_rawDescData = file_models_model_ssl_cert_proto_rawDesc
)

func file_models_model_ssl_cert_proto_rawDescGZIP() []byte {
	file_models_model_ssl_cert_proto_rawDescOnce.Do(func() {
		file_models_model_ssl_cert_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ssl_cert_proto_rawDescData)
	})
	return file_models_model_ssl_cert_proto_rawDescData
}

var file_models_model_ssl_cert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ssl_cert_proto_goTypes = []interface{}{
	(*SSLCert)(nil), // 0: pb.SSLCert
}
var file_models_model_ssl_cert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_ssl_cert_proto_init() }
func file_models_model_ssl_cert_proto_init() {
	if File_models_model_ssl_cert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_ssl_cert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSLCert); i {
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
			RawDescriptor: file_models_model_ssl_cert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ssl_cert_proto_goTypes,
		DependencyIndexes: file_models_model_ssl_cert_proto_depIdxs,
		MessageInfos:      file_models_model_ssl_cert_proto_msgTypes,
	}.Build()
	File_models_model_ssl_cert_proto = out.File
	file_models_model_ssl_cert_proto_rawDesc = nil
	file_models_model_ssl_cert_proto_goTypes = nil
	file_models_model_ssl_cert_proto_depIdxs = nil
}
