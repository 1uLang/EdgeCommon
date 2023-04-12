// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_message_recipient.proto

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

// 创建接收人
type CreateMessageRecipientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId                  int64   `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	MessageMediaInstanceId   int64   `protobuf:"varint,2,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
	MessageRecipientGroupIds []int64 `protobuf:"varint,3,rep,packed,name=messageRecipientGroupIds,proto3" json:"messageRecipientGroupIds,omitempty"`
	Description              string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	User                     string  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	TimeFrom                 string  `protobuf:"bytes,6,opt,name=timeFrom,proto3" json:"timeFrom,omitempty"`
	TimeTo                   string  `protobuf:"bytes,7,opt,name=timeTo,proto3" json:"timeTo,omitempty"`
}

func (x *CreateMessageRecipientRequest) Reset() {
	*x = CreateMessageRecipientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRecipientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRecipientRequest) ProtoMessage() {}

func (x *CreateMessageRecipientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRecipientRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRecipientRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRecipientRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CreateMessageRecipientRequest) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

func (x *CreateMessageRecipientRequest) GetMessageRecipientGroupIds() []int64 {
	if x != nil {
		return x.MessageRecipientGroupIds
	}
	return nil
}

func (x *CreateMessageRecipientRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMessageRecipientRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateMessageRecipientRequest) GetTimeFrom() string {
	if x != nil {
		return x.TimeFrom
	}
	return ""
}

func (x *CreateMessageRecipientRequest) GetTimeTo() string {
	if x != nil {
		return x.TimeTo
	}
	return ""
}

type CreateMessageRecipientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientId int64 `protobuf:"varint,1,opt,name=messageRecipientId,proto3" json:"messageRecipientId,omitempty"`
}

func (x *CreateMessageRecipientResponse) Reset() {
	*x = CreateMessageRecipientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRecipientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRecipientResponse) ProtoMessage() {}

func (x *CreateMessageRecipientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRecipientResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageRecipientResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageRecipientResponse) GetMessageRecipientId() int64 {
	if x != nil {
		return x.MessageRecipientId
	}
	return 0
}

// 修改接收人
type UpdateMessageRecipientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientId       int64   `protobuf:"varint,1,opt,name=messageRecipientId,proto3" json:"messageRecipientId,omitempty"`
	AdminId                  int64   `protobuf:"varint,2,opt,name=adminId,proto3" json:"adminId,omitempty"`
	MessageMediaInstanceId   int64   `protobuf:"varint,3,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
	MessageRecipientGroupIds []int64 `protobuf:"varint,4,rep,packed,name=messageRecipientGroupIds,proto3" json:"messageRecipientGroupIds,omitempty"`
	Description              string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	IsOn                     bool    `protobuf:"varint,6,opt,name=isOn,proto3" json:"isOn,omitempty"`
	User                     string  `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
	TimeFrom                 string  `protobuf:"bytes,8,opt,name=timeFrom,proto3" json:"timeFrom,omitempty"`
	TimeTo                   string  `protobuf:"bytes,9,opt,name=timeTo,proto3" json:"timeTo,omitempty"`
}

func (x *UpdateMessageRecipientRequest) Reset() {
	*x = UpdateMessageRecipientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRecipientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRecipientRequest) ProtoMessage() {}

func (x *UpdateMessageRecipientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRecipientRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRecipientRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMessageRecipientRequest) GetMessageRecipientId() int64 {
	if x != nil {
		return x.MessageRecipientId
	}
	return 0
}

func (x *UpdateMessageRecipientRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *UpdateMessageRecipientRequest) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

func (x *UpdateMessageRecipientRequest) GetMessageRecipientGroupIds() []int64 {
	if x != nil {
		return x.MessageRecipientGroupIds
	}
	return nil
}

func (x *UpdateMessageRecipientRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMessageRecipientRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UpdateMessageRecipientRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UpdateMessageRecipientRequest) GetTimeFrom() string {
	if x != nil {
		return x.TimeFrom
	}
	return ""
}

func (x *UpdateMessageRecipientRequest) GetTimeTo() string {
	if x != nil {
		return x.TimeTo
	}
	return ""
}

// 删除接收人
type DeleteMessageRecipientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientId int64 `protobuf:"varint,1,opt,name=messageRecipientId,proto3" json:"messageRecipientId,omitempty"`
}

func (x *DeleteMessageRecipientRequest) Reset() {
	*x = DeleteMessageRecipientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRecipientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRecipientRequest) ProtoMessage() {}

func (x *DeleteMessageRecipientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRecipientRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRecipientRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMessageRecipientRequest) GetMessageRecipientId() int64 {
	if x != nil {
		return x.MessageRecipientId
	}
	return 0
}

// 计算接收人数量
type CountAllEnabledMessageRecipientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId                 int64  `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	MediaType               string `protobuf:"bytes,2,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	MessageRecipientGroupId int64  `protobuf:"varint,3,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
	Keyword                 string `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *CountAllEnabledMessageRecipientsRequest) Reset() {
	*x = CountAllEnabledMessageRecipientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledMessageRecipientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledMessageRecipientsRequest) ProtoMessage() {}

func (x *CountAllEnabledMessageRecipientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledMessageRecipientsRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledMessageRecipientsRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{4}
}

func (x *CountAllEnabledMessageRecipientsRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CountAllEnabledMessageRecipientsRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *CountAllEnabledMessageRecipientsRequest) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

func (x *CountAllEnabledMessageRecipientsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// 列出单页接收人
type ListEnabledMessageRecipientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId                 int64  `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	MediaType               string `protobuf:"bytes,2,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	MessageRecipientGroupId int64  `protobuf:"varint,3,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
	Keyword                 string `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Offset                  int64  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Size                    int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListEnabledMessageRecipientsRequest) Reset() {
	*x = ListEnabledMessageRecipientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledMessageRecipientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledMessageRecipientsRequest) ProtoMessage() {}

func (x *ListEnabledMessageRecipientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledMessageRecipientsRequest.ProtoReflect.Descriptor instead.
func (*ListEnabledMessageRecipientsRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{5}
}

func (x *ListEnabledMessageRecipientsRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *ListEnabledMessageRecipientsRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ListEnabledMessageRecipientsRequest) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

func (x *ListEnabledMessageRecipientsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListEnabledMessageRecipientsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEnabledMessageRecipientsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListEnabledMessageRecipientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipients []*MessageRecipient `protobuf:"bytes,1,rep,name=messageRecipients,proto3" json:"messageRecipients,omitempty"`
}

func (x *ListEnabledMessageRecipientsResponse) Reset() {
	*x = ListEnabledMessageRecipientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledMessageRecipientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledMessageRecipientsResponse) ProtoMessage() {}

func (x *ListEnabledMessageRecipientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledMessageRecipientsResponse.ProtoReflect.Descriptor instead.
func (*ListEnabledMessageRecipientsResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{6}
}

func (x *ListEnabledMessageRecipientsResponse) GetMessageRecipients() []*MessageRecipient {
	if x != nil {
		return x.MessageRecipients
	}
	return nil
}

// 查找单个接收人信息
type FindEnabledMessageRecipientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientId int64 `protobuf:"varint,1,opt,name=messageRecipientId,proto3" json:"messageRecipientId,omitempty"`
}

func (x *FindEnabledMessageRecipientRequest) Reset() {
	*x = FindEnabledMessageRecipientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageRecipientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageRecipientRequest) ProtoMessage() {}

func (x *FindEnabledMessageRecipientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageRecipientRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageRecipientRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledMessageRecipientRequest) GetMessageRecipientId() int64 {
	if x != nil {
		return x.MessageRecipientId
	}
	return 0
}

type FindEnabledMessageRecipientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipient *MessageRecipient `protobuf:"bytes,1,opt,name=messageRecipient,proto3" json:"messageRecipient,omitempty"`
}

func (x *FindEnabledMessageRecipientResponse) Reset() {
	*x = FindEnabledMessageRecipientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageRecipientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageRecipientResponse) ProtoMessage() {}

func (x *FindEnabledMessageRecipientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageRecipientResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageRecipientResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_proto_rawDescGZIP(), []int{8}
}

func (x *FindEnabledMessageRecipientResponse) GetMessageRecipient() *MessageRecipient {
	if x != nil {
		return x.MessageRecipient
	}
	return nil
}

var File_service_message_recipient_proto protoreflect.FileDescriptor

var file_service_message_recipient_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x18, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x18, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x22, 0x50, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0xdb, 0x02, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x22, 0x4f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0xb5, 0x01, 0x0a, 0x27, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x23, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x6a, 0x0a, 0x24, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x23, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x32, 0xde, 0x04, 0x0a, 0x17, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5f, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4b,
	0x0a, 0x16, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x65, 0x0a, 0x20, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x71, 0x0a, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x1b, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_recipient_proto_rawDescOnce sync.Once
	file_service_message_recipient_proto_rawDescData = file_service_message_recipient_proto_rawDesc
)

func file_service_message_recipient_proto_rawDescGZIP() []byte {
	file_service_message_recipient_proto_rawDescOnce.Do(func() {
		file_service_message_recipient_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_recipient_proto_rawDescData)
	})
	return file_service_message_recipient_proto_rawDescData
}

var file_service_message_recipient_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_message_recipient_proto_goTypes = []interface{}{
	(*CreateMessageRecipientRequest)(nil),           // 0: pb.CreateMessageRecipientRequest
	(*CreateMessageRecipientResponse)(nil),          // 1: pb.CreateMessageRecipientResponse
	(*UpdateMessageRecipientRequest)(nil),           // 2: pb.UpdateMessageRecipientRequest
	(*DeleteMessageRecipientRequest)(nil),           // 3: pb.DeleteMessageRecipientRequest
	(*CountAllEnabledMessageRecipientsRequest)(nil), // 4: pb.CountAllEnabledMessageRecipientsRequest
	(*ListEnabledMessageRecipientsRequest)(nil),     // 5: pb.ListEnabledMessageRecipientsRequest
	(*ListEnabledMessageRecipientsResponse)(nil),    // 6: pb.ListEnabledMessageRecipientsResponse
	(*FindEnabledMessageRecipientRequest)(nil),      // 7: pb.FindEnabledMessageRecipientRequest
	(*FindEnabledMessageRecipientResponse)(nil),     // 8: pb.FindEnabledMessageRecipientResponse
	(*MessageRecipient)(nil),                        // 9: pb.MessageRecipient
	(*RPCSuccess)(nil),                              // 10: pb.RPCSuccess
	(*RPCCountResponse)(nil),                        // 11: pb.RPCCountResponse
}
var file_service_message_recipient_proto_depIdxs = []int32{
	9,  // 0: pb.ListEnabledMessageRecipientsResponse.messageRecipients:type_name -> pb.MessageRecipient
	9,  // 1: pb.FindEnabledMessageRecipientResponse.messageRecipient:type_name -> pb.MessageRecipient
	0,  // 2: pb.MessageRecipientService.createMessageRecipient:input_type -> pb.CreateMessageRecipientRequest
	2,  // 3: pb.MessageRecipientService.updateMessageRecipient:input_type -> pb.UpdateMessageRecipientRequest
	3,  // 4: pb.MessageRecipientService.deleteMessageRecipient:input_type -> pb.DeleteMessageRecipientRequest
	4,  // 5: pb.MessageRecipientService.countAllEnabledMessageRecipients:input_type -> pb.CountAllEnabledMessageRecipientsRequest
	5,  // 6: pb.MessageRecipientService.listEnabledMessageRecipients:input_type -> pb.ListEnabledMessageRecipientsRequest
	7,  // 7: pb.MessageRecipientService.findEnabledMessageRecipient:input_type -> pb.FindEnabledMessageRecipientRequest
	1,  // 8: pb.MessageRecipientService.createMessageRecipient:output_type -> pb.CreateMessageRecipientResponse
	10, // 9: pb.MessageRecipientService.updateMessageRecipient:output_type -> pb.RPCSuccess
	10, // 10: pb.MessageRecipientService.deleteMessageRecipient:output_type -> pb.RPCSuccess
	11, // 11: pb.MessageRecipientService.countAllEnabledMessageRecipients:output_type -> pb.RPCCountResponse
	6,  // 12: pb.MessageRecipientService.listEnabledMessageRecipients:output_type -> pb.ListEnabledMessageRecipientsResponse
	8,  // 13: pb.MessageRecipientService.findEnabledMessageRecipient:output_type -> pb.FindEnabledMessageRecipientResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_message_recipient_proto_init() }
func file_service_message_recipient_proto_init() {
	if File_service_message_recipient_proto != nil {
		return
	}
	file_models_model_message_recipient_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_recipient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRecipientRequest); i {
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
		file_service_message_recipient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRecipientResponse); i {
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
		file_service_message_recipient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRecipientRequest); i {
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
		file_service_message_recipient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRecipientRequest); i {
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
		file_service_message_recipient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledMessageRecipientsRequest); i {
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
		file_service_message_recipient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledMessageRecipientsRequest); i {
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
		file_service_message_recipient_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledMessageRecipientsResponse); i {
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
		file_service_message_recipient_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageRecipientRequest); i {
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
		file_service_message_recipient_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageRecipientResponse); i {
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
			RawDescriptor: file_service_message_recipient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_recipient_proto_goTypes,
		DependencyIndexes: file_service_message_recipient_proto_depIdxs,
		MessageInfos:      file_service_message_recipient_proto_msgTypes,
	}.Build()
	File_service_message_recipient_proto = out.File
	file_service_message_recipient_proto_rawDesc = nil
	file_service_message_recipient_proto_goTypes = nil
	file_service_message_recipient_proto_depIdxs = nil
}
