// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: models/model_message_task_log.proto

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

// 消息任务日志
type MessageTaskLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt   int64        `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	IsOk        bool         `protobuf:"varint,3,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Error       string       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Response    string       `protobuf:"bytes,5,opt,name=response,proto3" json:"response,omitempty"`
	MessageTask *MessageTask `protobuf:"bytes,6,opt,name=messageTask,proto3" json:"messageTask,omitempty"`
}

func (x *MessageTaskLog) Reset() {
	*x = MessageTaskLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_message_task_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTaskLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTaskLog) ProtoMessage() {}

func (x *MessageTaskLog) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_message_task_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTaskLog.ProtoReflect.Descriptor instead.
func (*MessageTaskLog) Descriptor() ([]byte, []int) {
	return file_models_model_message_task_log_proto_rawDescGZIP(), []int{0}
}

func (x *MessageTaskLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageTaskLog) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MessageTaskLog) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *MessageTaskLog) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MessageTaskLog) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MessageTaskLog) GetMessageTask() *MessageTask {
	if x != nil {
		return x.MessageTask
	}
	return nil
}

var File_models_model_message_task_log_proto protoreflect.FileDescriptor

var file_models_model_message_task_log_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_message_task_log_proto_rawDescOnce sync.Once
	file_models_model_message_task_log_proto_rawDescData = file_models_model_message_task_log_proto_rawDesc
)

func file_models_model_message_task_log_proto_rawDescGZIP() []byte {
	file_models_model_message_task_log_proto_rawDescOnce.Do(func() {
		file_models_model_message_task_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_message_task_log_proto_rawDescData)
	})
	return file_models_model_message_task_log_proto_rawDescData
}

var file_models_model_message_task_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_message_task_log_proto_goTypes = []interface{}{
	(*MessageTaskLog)(nil), // 0: pb.MessageTaskLog
	(*MessageTask)(nil),    // 1: pb.MessageTask
}
var file_models_model_message_task_log_proto_depIdxs = []int32{
	1, // 0: pb.MessageTaskLog.messageTask:type_name -> pb.MessageTask
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_message_task_log_proto_init() }
func file_models_model_message_task_log_proto_init() {
	if File_models_model_message_task_log_proto != nil {
		return
	}
	file_models_model_message_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_message_task_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageTaskLog); i {
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
			RawDescriptor: file_models_model_message_task_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_message_task_log_proto_goTypes,
		DependencyIndexes: file_models_model_message_task_log_proto_depIdxs,
		MessageInfos:      file_models_model_message_task_log_proto_msgTypes,
	}.Build()
	File_models_model_message_task_log_proto = out.File
	file_models_model_message_task_log_proto_rawDesc = nil
	file_models_model_message_task_log_proto_goTypes = nil
	file_models_model_message_task_log_proto_depIdxs = nil
}
