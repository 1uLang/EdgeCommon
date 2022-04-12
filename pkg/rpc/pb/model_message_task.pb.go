// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: models/model_message_task.proto

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

type MessageTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MessageRecipient     *MessageRecipient     `protobuf:"bytes,2,opt,name=messageRecipient,proto3" json:"messageRecipient,omitempty"`
	User                 string                `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Subject              string                `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                 string                `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            int64                 `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Status               int32                 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	SentAt               int64                 `protobuf:"varint,8,opt,name=sentAt,proto3" json:"sentAt,omitempty"`
	Result               *MessageTaskResult    `protobuf:"bytes,9,opt,name=result,proto3" json:"result,omitempty"`
	MessageMediaInstance *MessageMediaInstance `protobuf:"bytes,10,opt,name=messageMediaInstance,proto3" json:"messageMediaInstance,omitempty"`
}

func (x *MessageTask) Reset() {
	*x = MessageTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_message_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTask) ProtoMessage() {}

func (x *MessageTask) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_message_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTask.ProtoReflect.Descriptor instead.
func (*MessageTask) Descriptor() ([]byte, []int) {
	return file_models_model_message_task_proto_rawDescGZIP(), []int{0}
}

func (x *MessageTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageTask) GetMessageRecipient() *MessageRecipient {
	if x != nil {
		return x.MessageRecipient
	}
	return nil
}

func (x *MessageTask) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MessageTask) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *MessageTask) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *MessageTask) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MessageTask) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MessageTask) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

func (x *MessageTask) GetResult() *MessageTaskResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *MessageTask) GetMessageMediaInstance() *MessageMediaInstance {
	if x != nil {
		return x.MessageMediaInstance
	}
	return nil
}

type MessageTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk     bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Error    string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MessageTaskResult) Reset() {
	*x = MessageTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_message_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTaskResult) ProtoMessage() {}

func (x *MessageTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_message_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTaskResult.ProtoReflect.Descriptor instead.
func (*MessageTaskResult) Descriptor() ([]byte, []int) {
	return file_models_model_message_task_proto_rawDescGZIP(), []int{1}
}

func (x *MessageTaskResult) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *MessageTaskResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MessageTaskResult) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_models_model_message_task_proto protoreflect.FileDescriptor

var file_models_model_message_task_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_message_task_proto_rawDescOnce sync.Once
	file_models_model_message_task_proto_rawDescData = file_models_model_message_task_proto_rawDesc
)

func file_models_model_message_task_proto_rawDescGZIP() []byte {
	file_models_model_message_task_proto_rawDescOnce.Do(func() {
		file_models_model_message_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_message_task_proto_rawDescData)
	})
	return file_models_model_message_task_proto_rawDescData
}

var file_models_model_message_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_model_message_task_proto_goTypes = []interface{}{
	(*MessageTask)(nil),          // 0: pb.MessageTask
	(*MessageTaskResult)(nil),    // 1: pb.MessageTaskResult
	(*MessageRecipient)(nil),     // 2: pb.MessageRecipient
	(*MessageMediaInstance)(nil), // 3: pb.MessageMediaInstance
}
var file_models_model_message_task_proto_depIdxs = []int32{
	2, // 0: pb.MessageTask.messageRecipient:type_name -> pb.MessageRecipient
	1, // 1: pb.MessageTask.result:type_name -> pb.MessageTaskResult
	3, // 2: pb.MessageTask.messageMediaInstance:type_name -> pb.MessageMediaInstance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_message_task_proto_init() }
func file_models_model_message_task_proto_init() {
	if File_models_model_message_task_proto != nil {
		return
	}
	file_models_model_message_recipient_proto_init()
	file_models_model_message_media_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_message_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageTask); i {
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
		file_models_model_message_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageTaskResult); i {
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
			RawDescriptor: file_models_model_message_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_message_task_proto_goTypes,
		DependencyIndexes: file_models_model_message_task_proto_depIdxs,
		MessageInfos:      file_models_model_message_task_proto_msgTypes,
	}.Build()
	File_models_model_message_task_proto = out.File
	file_models_model_message_task_proto_rawDesc = nil
	file_models_model_message_task_proto_goTypes = nil
	file_models_model_message_task_proto_depIdxs = nil
}
