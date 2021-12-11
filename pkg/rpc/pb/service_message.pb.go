// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service_message.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 计算未读消息数
type CountUnreadMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountUnreadMessagesRequest) Reset() {
	*x = CountUnreadMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUnreadMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUnreadMessagesRequest) ProtoMessage() {}

func (x *CountUnreadMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUnreadMessagesRequest.ProtoReflect.Descriptor instead.
func (*CountUnreadMessagesRequest) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{0}
}

// 列出单页未读消息
type ListUnreadMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUnreadMessagesRequest) Reset() {
	*x = ListUnreadMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUnreadMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnreadMessagesRequest) ProtoMessage() {}

func (x *ListUnreadMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnreadMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListUnreadMessagesRequest) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{1}
}

func (x *ListUnreadMessagesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUnreadMessagesRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUnreadMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ListUnreadMessagesResponse) Reset() {
	*x = ListUnreadMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUnreadMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUnreadMessagesResponse) ProtoMessage() {}

func (x *ListUnreadMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUnreadMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListUnreadMessagesResponse) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{2}
}

func (x *ListUnreadMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// 设置消息已读状态
type UpdateMessageReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId int64 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	IsRead    bool  `protobuf:"varint,2,opt,name=isRead,proto3" json:"isRead,omitempty"`
}

func (x *UpdateMessageReadRequest) Reset() {
	*x = UpdateMessageReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageReadRequest) ProtoMessage() {}

func (x *UpdateMessageReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageReadRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageReadRequest) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateMessageReadRequest) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *UpdateMessageReadRequest) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

// 设置一组消息已读状态
type UpdateMessagesReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageIds []int64 `protobuf:"varint,1,rep,packed,name=messageIds,proto3" json:"messageIds,omitempty"`
	IsRead     bool    `protobuf:"varint,2,opt,name=isRead,proto3" json:"isRead,omitempty"`
}

func (x *UpdateMessagesReadRequest) Reset() {
	*x = UpdateMessagesReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessagesReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessagesReadRequest) ProtoMessage() {}

func (x *UpdateMessagesReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessagesReadRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessagesReadRequest) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMessagesReadRequest) GetMessageIds() []int64 {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

func (x *UpdateMessagesReadRequest) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

// 设置所有消息为已读
type UpdateAllMessagesReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAllMessagesReadRequest) Reset() {
	*x = UpdateAllMessagesReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAllMessagesReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAllMessagesReadRequest) ProtoMessage() {}

func (x *UpdateAllMessagesReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAllMessagesReadRequest.ProtoReflect.Descriptor instead.
func (*UpdateAllMessagesReadRequest) Descriptor() ([]byte, []int) {
	return file_service_message_proto_rawDescGZIP(), []int{5}
}

var File_service_message_proto protoreflect.FileDescriptor

var file_service_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x47, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x45, 0x0a, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x50, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x22, 0x53, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x85, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x13, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x6e, 0x72, 0x65,
	0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x55,
	0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x43, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x20, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_proto_rawDescOnce sync.Once
	file_service_message_proto_rawDescData = file_service_message_proto_rawDesc
)

func file_service_message_proto_rawDescGZIP() []byte {
	file_service_message_proto_rawDescOnce.Do(func() {
		file_service_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_proto_rawDescData)
	})
	return file_service_message_proto_rawDescData
}

var file_service_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_message_proto_goTypes = []interface{}{
	(*CountUnreadMessagesRequest)(nil),   // 0: pb.CountUnreadMessagesRequest
	(*ListUnreadMessagesRequest)(nil),    // 1: pb.ListUnreadMessagesRequest
	(*ListUnreadMessagesResponse)(nil),   // 2: pb.ListUnreadMessagesResponse
	(*UpdateMessageReadRequest)(nil),     // 3: pb.UpdateMessageReadRequest
	(*UpdateMessagesReadRequest)(nil),    // 4: pb.UpdateMessagesReadRequest
	(*UpdateAllMessagesReadRequest)(nil), // 5: pb.UpdateAllMessagesReadRequest
	(*Message)(nil),                      // 6: pb.Message
	(*RPCCountResponse)(nil),             // 7: pb.RPCCountResponse
	(*RPCSuccess)(nil),                   // 8: pb.RPCSuccess
}
var file_service_message_proto_depIdxs = []int32{
	6, // 0: pb.ListUnreadMessagesResponse.messages:type_name -> pb.Message
	0, // 1: pb.MessageService.countUnreadMessages:input_type -> pb.CountUnreadMessagesRequest
	1, // 2: pb.MessageService.listUnreadMessages:input_type -> pb.ListUnreadMessagesRequest
	3, // 3: pb.MessageService.updateMessageRead:input_type -> pb.UpdateMessageReadRequest
	4, // 4: pb.MessageService.updateMessagesRead:input_type -> pb.UpdateMessagesReadRequest
	5, // 5: pb.MessageService.updateAllMessagesRead:input_type -> pb.UpdateAllMessagesReadRequest
	7, // 6: pb.MessageService.countUnreadMessages:output_type -> pb.RPCCountResponse
	2, // 7: pb.MessageService.listUnreadMessages:output_type -> pb.ListUnreadMessagesResponse
	8, // 8: pb.MessageService.updateMessageRead:output_type -> pb.RPCSuccess
	8, // 9: pb.MessageService.updateMessagesRead:output_type -> pb.RPCSuccess
	8, // 10: pb.MessageService.updateAllMessagesRead:output_type -> pb.RPCSuccess
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_message_proto_init() }
func file_service_message_proto_init() {
	if File_service_message_proto != nil {
		return
	}
	file_models_model_message_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUnreadMessagesRequest); i {
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
		file_service_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUnreadMessagesRequest); i {
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
		file_service_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUnreadMessagesResponse); i {
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
		file_service_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageReadRequest); i {
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
		file_service_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessagesReadRequest); i {
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
		file_service_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAllMessagesReadRequest); i {
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
			RawDescriptor: file_service_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_proto_goTypes,
		DependencyIndexes: file_service_message_proto_depIdxs,
		MessageInfos:      file_service_message_proto_msgTypes,
	}.Build()
	File_service_message_proto = out.File
	file_service_message_proto_rawDesc = nil
	file_service_message_proto_goTypes = nil
	file_service_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	// 计算未读消息数
	CountUnreadMessages(ctx context.Context, in *CountUnreadMessagesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页未读消息
	ListUnreadMessages(ctx context.Context, in *ListUnreadMessagesRequest, opts ...grpc.CallOption) (*ListUnreadMessagesResponse, error)
	// 设置消息已读状态
	UpdateMessageRead(ctx context.Context, in *UpdateMessageReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 设置一组消息已读状态
	UpdateMessagesRead(ctx context.Context, in *UpdateMessagesReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 设置所有消息为已读
	UpdateAllMessagesRead(ctx context.Context, in *UpdateAllMessagesReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CountUnreadMessages(ctx context.Context, in *CountUnreadMessagesRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageService/countUnreadMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) ListUnreadMessages(ctx context.Context, in *ListUnreadMessagesRequest, opts ...grpc.CallOption) (*ListUnreadMessagesResponse, error) {
	out := new(ListUnreadMessagesResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageService/listUnreadMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UpdateMessageRead(ctx context.Context, in *UpdateMessageReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageService/updateMessageRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UpdateMessagesRead(ctx context.Context, in *UpdateMessagesReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageService/updateMessagesRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UpdateAllMessagesRead(ctx context.Context, in *UpdateAllMessagesReadRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageService/updateAllMessagesRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	// 计算未读消息数
	CountUnreadMessages(context.Context, *CountUnreadMessagesRequest) (*RPCCountResponse, error)
	// 列出单页未读消息
	ListUnreadMessages(context.Context, *ListUnreadMessagesRequest) (*ListUnreadMessagesResponse, error)
	// 设置消息已读状态
	UpdateMessageRead(context.Context, *UpdateMessageReadRequest) (*RPCSuccess, error)
	// 设置一组消息已读状态
	UpdateMessagesRead(context.Context, *UpdateMessagesReadRequest) (*RPCSuccess, error)
	// 设置所有消息为已读
	UpdateAllMessagesRead(context.Context, *UpdateAllMessagesReadRequest) (*RPCSuccess, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) CountUnreadMessages(context.Context, *CountUnreadMessagesRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUnreadMessages not implemented")
}
func (*UnimplementedMessageServiceServer) ListUnreadMessages(context.Context, *ListUnreadMessagesRequest) (*ListUnreadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUnreadMessages not implemented")
}
func (*UnimplementedMessageServiceServer) UpdateMessageRead(context.Context, *UpdateMessageReadRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageRead not implemented")
}
func (*UnimplementedMessageServiceServer) UpdateMessagesRead(context.Context, *UpdateMessagesReadRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessagesRead not implemented")
}
func (*UnimplementedMessageServiceServer) UpdateAllMessagesRead(context.Context, *UpdateAllMessagesReadRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAllMessagesRead not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_CountUnreadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUnreadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CountUnreadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/CountUnreadMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CountUnreadMessages(ctx, req.(*CountUnreadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_ListUnreadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUnreadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).ListUnreadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/ListUnreadMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).ListUnreadMessages(ctx, req.(*ListUnreadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UpdateMessageRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UpdateMessageRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/UpdateMessageRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UpdateMessageRead(ctx, req.(*UpdateMessageReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UpdateMessagesRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessagesReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UpdateMessagesRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/UpdateMessagesRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UpdateMessagesRead(ctx, req.(*UpdateMessagesReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UpdateAllMessagesRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAllMessagesReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UpdateAllMessagesRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/UpdateAllMessagesRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UpdateAllMessagesRead(ctx, req.(*UpdateAllMessagesReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "countUnreadMessages",
			Handler:    _MessageService_CountUnreadMessages_Handler,
		},
		{
			MethodName: "listUnreadMessages",
			Handler:    _MessageService_ListUnreadMessages_Handler,
		},
		{
			MethodName: "updateMessageRead",
			Handler:    _MessageService_UpdateMessageRead_Handler,
		},
		{
			MethodName: "updateMessagesRead",
			Handler:    _MessageService_UpdateMessagesRead_Handler,
		},
		{
			MethodName: "updateAllMessagesRead",
			Handler:    _MessageService_UpdateAllMessagesRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_message.proto",
}
