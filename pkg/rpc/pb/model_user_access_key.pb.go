// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_user_access_key.proto

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

// 用户AccessKey
type UserAccessKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	SubUserId   int64  `protobuf:"varint,3,opt,name=subUserId,proto3" json:"subUserId,omitempty"`
	IsOn        bool   `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`
	UniqueId    string `protobuf:"bytes,5,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret      string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	AccessedAt  int64  `protobuf:"varint,8,opt,name=accessedAt,proto3" json:"accessedAt,omitempty"`
}

func (x *UserAccessKey) Reset() {
	*x = UserAccessKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_access_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccessKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccessKey) ProtoMessage() {}

func (x *UserAccessKey) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_access_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccessKey.ProtoReflect.Descriptor instead.
func (*UserAccessKey) Descriptor() ([]byte, []int) {
	return file_models_model_user_access_key_proto_rawDescGZIP(), []int{0}
}

func (x *UserAccessKey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserAccessKey) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAccessKey) GetSubUserId() int64 {
	if x != nil {
		return x.SubUserId
	}
	return 0
}

func (x *UserAccessKey) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UserAccessKey) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *UserAccessKey) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UserAccessKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserAccessKey) GetAccessedAt() int64 {
	if x != nil {
		return x.AccessedAt
	}
	return 0
}

var File_models_model_user_access_key_proto protoreflect.FileDescriptor

var file_models_model_user_access_key_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xdf, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_access_key_proto_rawDescOnce sync.Once
	file_models_model_user_access_key_proto_rawDescData = file_models_model_user_access_key_proto_rawDesc
)

func file_models_model_user_access_key_proto_rawDescGZIP() []byte {
	file_models_model_user_access_key_proto_rawDescOnce.Do(func() {
		file_models_model_user_access_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_access_key_proto_rawDescData)
	})
	return file_models_model_user_access_key_proto_rawDescData
}

var file_models_model_user_access_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_access_key_proto_goTypes = []interface{}{
	(*UserAccessKey)(nil), // 0: pb.UserAccessKey
}
var file_models_model_user_access_key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_user_access_key_proto_init() }
func file_models_model_user_access_key_proto_init() {
	if File_models_model_user_access_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_access_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccessKey); i {
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
			RawDescriptor: file_models_model_user_access_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_access_key_proto_goTypes,
		DependencyIndexes: file_models_model_user_access_key_proto_depIdxs,
		MessageInfos:      file_models_model_user_access_key_proto_msgTypes,
	}.Build()
	File_models_model_user_access_key_proto = out.File
	file_models_model_user_access_key_proto_rawDesc = nil
	file_models_model_user_access_key_proto_goTypes = nil
	file_models_model_user_access_key_proto_depIdxs = nil
}
