// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: models/model_message_recipient.proto

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

type MessageRecipient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Admin                  *Admin                   `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	MessageMediaInstance   *MessageMediaInstance    `protobuf:"bytes,3,opt,name=messageMediaInstance,proto3" json:"messageMediaInstance,omitempty"`
	IsOn                   bool                     `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`
	MessageRecipientGroups []*MessageRecipientGroup `protobuf:"bytes,5,rep,name=messageRecipientGroups,proto3" json:"messageRecipientGroups,omitempty"`
	Description            string                   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	User                   string                   `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	TimeFrom               string                   `protobuf:"bytes,8,opt,name=timeFrom,proto3" json:"timeFrom,omitempty"`
	TimeTo                 string                   `protobuf:"bytes,9,opt,name=timeTo,proto3" json:"timeTo,omitempty"`
}

func (x *MessageRecipient) Reset() {
	*x = MessageRecipient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_message_recipient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRecipient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRecipient) ProtoMessage() {}

func (x *MessageRecipient) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_message_recipient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRecipient.ProtoReflect.Descriptor instead.
func (*MessageRecipient) Descriptor() ([]byte, []int) {
	return file_models_model_message_recipient_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRecipient) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageRecipient) GetAdmin() *Admin {
	if x != nil {
		return x.Admin
	}
	return nil
}

func (x *MessageRecipient) GetMessageMediaInstance() *MessageMediaInstance {
	if x != nil {
		return x.MessageMediaInstance
	}
	return nil
}

func (x *MessageRecipient) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *MessageRecipient) GetMessageRecipientGroups() []*MessageRecipientGroup {
	if x != nil {
		return x.MessageRecipientGroups
	}
	return nil
}

func (x *MessageRecipient) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MessageRecipient) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MessageRecipient) GetTimeFrom() string {
	if x != nil {
		return x.TimeFrom
	}
	return ""
}

func (x *MessageRecipient) GetTimeTo() string {
	if x != nil {
		return x.TimeTo
	}
	return ""
}

var File_models_model_message_recipient_proto protoreflect.FileDescriptor

var file_models_model_message_recipient_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x18, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x10,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x4c, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x12, 0x51, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x16,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_message_recipient_proto_rawDescOnce sync.Once
	file_models_model_message_recipient_proto_rawDescData = file_models_model_message_recipient_proto_rawDesc
)

func file_models_model_message_recipient_proto_rawDescGZIP() []byte {
	file_models_model_message_recipient_proto_rawDescOnce.Do(func() {
		file_models_model_message_recipient_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_message_recipient_proto_rawDescData)
	})
	return file_models_model_message_recipient_proto_rawDescData
}

var file_models_model_message_recipient_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_message_recipient_proto_goTypes = []interface{}{
	(*MessageRecipient)(nil),      // 0: pb.MessageRecipient
	(*Admin)(nil),                 // 1: pb.Admin
	(*MessageMediaInstance)(nil),  // 2: pb.MessageMediaInstance
	(*MessageRecipientGroup)(nil), // 3: pb.MessageRecipientGroup
}
var file_models_model_message_recipient_proto_depIdxs = []int32{
	1, // 0: pb.MessageRecipient.admin:type_name -> pb.Admin
	2, // 1: pb.MessageRecipient.messageMediaInstance:type_name -> pb.MessageMediaInstance
	3, // 2: pb.MessageRecipient.messageRecipientGroups:type_name -> pb.MessageRecipientGroup
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_message_recipient_proto_init() }
func file_models_model_message_recipient_proto_init() {
	if File_models_model_message_recipient_proto != nil {
		return
	}
	file_models_model_admin_proto_init()
	file_models_model_message_recipient_group_proto_init()
	file_models_model_message_media_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_message_recipient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRecipient); i {
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
			RawDescriptor: file_models_model_message_recipient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_message_recipient_proto_goTypes,
		DependencyIndexes: file_models_model_message_recipient_proto_depIdxs,
		MessageInfos:      file_models_model_message_recipient_proto_msgTypes,
	}.Build()
	File_models_model_message_recipient_proto = out.File
	file_models_model_message_recipient_proto_rawDesc = nil
	file_models_model_message_recipient_proto_goTypes = nil
	file_models_model_message_recipient_proto_depIdxs = nil
}
