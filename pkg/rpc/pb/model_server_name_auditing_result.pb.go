// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_server_name_auditing_result.proto

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

type ServerNameAuditingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk      bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Reason    string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *ServerNameAuditingResult) Reset() {
	*x = ServerNameAuditingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_server_name_auditing_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerNameAuditingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameAuditingResult) ProtoMessage() {}

func (x *ServerNameAuditingResult) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_server_name_auditing_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameAuditingResult.ProtoReflect.Descriptor instead.
func (*ServerNameAuditingResult) Descriptor() ([]byte, []int) {
	return file_models_model_server_name_auditing_result_proto_rawDescGZIP(), []int{0}
}

func (x *ServerNameAuditingResult) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ServerNameAuditingResult) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ServerNameAuditingResult) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_models_model_server_name_auditing_result_proto protoreflect.FileDescriptor

var file_models_model_server_name_auditing_result_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x64, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x4f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_server_name_auditing_result_proto_rawDescOnce sync.Once
	file_models_model_server_name_auditing_result_proto_rawDescData = file_models_model_server_name_auditing_result_proto_rawDesc
)

func file_models_model_server_name_auditing_result_proto_rawDescGZIP() []byte {
	file_models_model_server_name_auditing_result_proto_rawDescOnce.Do(func() {
		file_models_model_server_name_auditing_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_server_name_auditing_result_proto_rawDescData)
	})
	return file_models_model_server_name_auditing_result_proto_rawDescData
}

var file_models_model_server_name_auditing_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_server_name_auditing_result_proto_goTypes = []interface{}{
	(*ServerNameAuditingResult)(nil), // 0: pb.ServerNameAuditingResult
}
var file_models_model_server_name_auditing_result_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_server_name_auditing_result_proto_init() }
func file_models_model_server_name_auditing_result_proto_init() {
	if File_models_model_server_name_auditing_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_server_name_auditing_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerNameAuditingResult); i {
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
			RawDescriptor: file_models_model_server_name_auditing_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_server_name_auditing_result_proto_goTypes,
		DependencyIndexes: file_models_model_server_name_auditing_result_proto_depIdxs,
		MessageInfos:      file_models_model_server_name_auditing_result_proto_msgTypes,
	}.Build()
	File_models_model_server_name_auditing_result_proto = out.File
	file_models_model_server_name_auditing_result_proto_rawDesc = nil
	file_models_model_server_name_auditing_result_proto_goTypes = nil
	file_models_model_server_name_auditing_result_proto_depIdxs = nil
}
