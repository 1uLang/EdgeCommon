// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_ip_library.proto

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

type IPLibrary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	File      *File  `protobuf:"bytes,30,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *IPLibrary) Reset() {
	*x = IPLibrary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ip_library_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPLibrary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPLibrary) ProtoMessage() {}

func (x *IPLibrary) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ip_library_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPLibrary.ProtoReflect.Descriptor instead.
func (*IPLibrary) Descriptor() ([]byte, []int) {
	return file_models_model_ip_library_proto_rawDescGZIP(), []int{0}
}

func (x *IPLibrary) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IPLibrary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IPLibrary) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *IPLibrary) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_models_model_ip_library_proto protoreflect.FileDescriptor

var file_models_model_ip_library_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69,
	0x70, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x09,
	0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ip_library_proto_rawDescOnce sync.Once
	file_models_model_ip_library_proto_rawDescData = file_models_model_ip_library_proto_rawDesc
)

func file_models_model_ip_library_proto_rawDescGZIP() []byte {
	file_models_model_ip_library_proto_rawDescOnce.Do(func() {
		file_models_model_ip_library_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ip_library_proto_rawDescData)
	})
	return file_models_model_ip_library_proto_rawDescData
}

var file_models_model_ip_library_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ip_library_proto_goTypes = []interface{}{
	(*IPLibrary)(nil), // 0: pb.IPLibrary
	(*File)(nil),      // 1: pb.File
}
var file_models_model_ip_library_proto_depIdxs = []int32{
	1, // 0: pb.IPLibrary.file:type_name -> pb.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_ip_library_proto_init() }
func file_models_model_ip_library_proto_init() {
	if File_models_model_ip_library_proto != nil {
		return
	}
	file_models_model_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ip_library_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPLibrary); i {
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
			RawDescriptor: file_models_model_ip_library_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ip_library_proto_goTypes,
		DependencyIndexes: file_models_model_ip_library_proto_depIdxs,
		MessageInfos:      file_models_model_ip_library_proto_msgTypes,
	}.Build()
	File_models_model_ip_library_proto = out.File
	file_models_model_ip_library_proto_rawDesc = nil
	file_models_model_ip_library_proto_goTypes = nil
	file_models_model_ip_library_proto_depIdxs = nil
}
