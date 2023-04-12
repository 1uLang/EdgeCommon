// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_message_recipient_group.proto

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

// 创建分组
type CreateMessageRecipientGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateMessageRecipientGroupRequest) Reset() {
	*x = CreateMessageRecipientGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRecipientGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRecipientGroupRequest) ProtoMessage() {}

func (x *CreateMessageRecipientGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRecipientGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRecipientGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRecipientGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateMessageRecipientGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroupId int64 `protobuf:"varint,1,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
}

func (x *CreateMessageRecipientGroupResponse) Reset() {
	*x = CreateMessageRecipientGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRecipientGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRecipientGroupResponse) ProtoMessage() {}

func (x *CreateMessageRecipientGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRecipientGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageRecipientGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageRecipientGroupResponse) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

// 修改分组
type UpdateMessageRecipientGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroupId int64  `protobuf:"varint,1,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
	Name                    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn                    bool   `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateMessageRecipientGroupRequest) Reset() {
	*x = UpdateMessageRecipientGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRecipientGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRecipientGroupRequest) ProtoMessage() {}

func (x *UpdateMessageRecipientGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRecipientGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRecipientGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMessageRecipientGroupRequest) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

func (x *UpdateMessageRecipientGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMessageRecipientGroupRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 查找所有可用的分组
type FindAllEnabledMessageRecipientGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllEnabledMessageRecipientGroupsRequest) Reset() {
	*x = FindAllEnabledMessageRecipientGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledMessageRecipientGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledMessageRecipientGroupsRequest) ProtoMessage() {}

func (x *FindAllEnabledMessageRecipientGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledMessageRecipientGroupsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledMessageRecipientGroupsRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{3}
}

type FindAllEnabledMessageRecipientGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroups []*MessageRecipientGroup `protobuf:"bytes,1,rep,name=messageRecipientGroups,proto3" json:"messageRecipientGroups,omitempty"`
}

func (x *FindAllEnabledMessageRecipientGroupsResponse) Reset() {
	*x = FindAllEnabledMessageRecipientGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledMessageRecipientGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledMessageRecipientGroupsResponse) ProtoMessage() {}

func (x *FindAllEnabledMessageRecipientGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledMessageRecipientGroupsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledMessageRecipientGroupsResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllEnabledMessageRecipientGroupsResponse) GetMessageRecipientGroups() []*MessageRecipientGroup {
	if x != nil {
		return x.MessageRecipientGroups
	}
	return nil
}

// 删除分组
type DeleteMessageRecipientGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroupId int64 `protobuf:"varint,1,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
}

func (x *DeleteMessageRecipientGroupRequest) Reset() {
	*x = DeleteMessageRecipientGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRecipientGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRecipientGroupRequest) ProtoMessage() {}

func (x *DeleteMessageRecipientGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRecipientGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRecipientGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMessageRecipientGroupRequest) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

// 查找单个分组信息
type FindEnabledMessageRecipientGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroupId int64 `protobuf:"varint,1,opt,name=messageRecipientGroupId,proto3" json:"messageRecipientGroupId,omitempty"`
}

func (x *FindEnabledMessageRecipientGroupRequest) Reset() {
	*x = FindEnabledMessageRecipientGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageRecipientGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageRecipientGroupRequest) ProtoMessage() {}

func (x *FindEnabledMessageRecipientGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageRecipientGroupRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageRecipientGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledMessageRecipientGroupRequest) GetMessageRecipientGroupId() int64 {
	if x != nil {
		return x.MessageRecipientGroupId
	}
	return 0
}

type FindEnabledMessageRecipientGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageRecipientGroup *MessageRecipientGroup `protobuf:"bytes,1,opt,name=messageRecipientGroup,proto3" json:"messageRecipientGroup,omitempty"`
}

func (x *FindEnabledMessageRecipientGroupResponse) Reset() {
	*x = FindEnabledMessageRecipientGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_recipient_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageRecipientGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageRecipientGroupResponse) ProtoMessage() {}

func (x *FindEnabledMessageRecipientGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_recipient_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageRecipientGroupResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageRecipientGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_message_recipient_group_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledMessageRecipientGroupResponse) GetMessageRecipientGroup() *MessageRecipientGroup {
	if x != nil {
		return x.MessageRecipientGroup
	}
	return nil
}

var File_service_message_recipient_group_proto protoreflect.FileDescriptor

var file_service_message_recipient_group_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x2a, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x38, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x23,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x86, 0x01,
	0x0a, 0x22, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x2d, 0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x2c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x5e, 0x0a, 0x22, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x27, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x7b,
	0x0a, 0x28, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x32, 0xc7, 0x04, 0x0a, 0x1c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x1b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x1b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x24, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2f, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x55, 0x0a, 0x1b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x7d, 0x0a, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2b, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_recipient_group_proto_rawDescOnce sync.Once
	file_service_message_recipient_group_proto_rawDescData = file_service_message_recipient_group_proto_rawDesc
)

func file_service_message_recipient_group_proto_rawDescGZIP() []byte {
	file_service_message_recipient_group_proto_rawDescOnce.Do(func() {
		file_service_message_recipient_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_recipient_group_proto_rawDescData)
	})
	return file_service_message_recipient_group_proto_rawDescData
}

var file_service_message_recipient_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_message_recipient_group_proto_goTypes = []interface{}{
	(*CreateMessageRecipientGroupRequest)(nil),           // 0: pb.CreateMessageRecipientGroupRequest
	(*CreateMessageRecipientGroupResponse)(nil),          // 1: pb.CreateMessageRecipientGroupResponse
	(*UpdateMessageRecipientGroupRequest)(nil),           // 2: pb.UpdateMessageRecipientGroupRequest
	(*FindAllEnabledMessageRecipientGroupsRequest)(nil),  // 3: pb.FindAllEnabledMessageRecipientGroupsRequest
	(*FindAllEnabledMessageRecipientGroupsResponse)(nil), // 4: pb.FindAllEnabledMessageRecipientGroupsResponse
	(*DeleteMessageRecipientGroupRequest)(nil),           // 5: pb.DeleteMessageRecipientGroupRequest
	(*FindEnabledMessageRecipientGroupRequest)(nil),      // 6: pb.FindEnabledMessageRecipientGroupRequest
	(*FindEnabledMessageRecipientGroupResponse)(nil),     // 7: pb.FindEnabledMessageRecipientGroupResponse
	(*MessageRecipientGroup)(nil),                        // 8: pb.MessageRecipientGroup
	(*RPCSuccess)(nil),                                   // 9: pb.RPCSuccess
}
var file_service_message_recipient_group_proto_depIdxs = []int32{
	8, // 0: pb.FindAllEnabledMessageRecipientGroupsResponse.messageRecipientGroups:type_name -> pb.MessageRecipientGroup
	8, // 1: pb.FindEnabledMessageRecipientGroupResponse.messageRecipientGroup:type_name -> pb.MessageRecipientGroup
	0, // 2: pb.MessageRecipientGroupService.createMessageRecipientGroup:input_type -> pb.CreateMessageRecipientGroupRequest
	2, // 3: pb.MessageRecipientGroupService.updateMessageRecipientGroup:input_type -> pb.UpdateMessageRecipientGroupRequest
	3, // 4: pb.MessageRecipientGroupService.findAllEnabledMessageRecipientGroups:input_type -> pb.FindAllEnabledMessageRecipientGroupsRequest
	5, // 5: pb.MessageRecipientGroupService.deleteMessageRecipientGroup:input_type -> pb.DeleteMessageRecipientGroupRequest
	6, // 6: pb.MessageRecipientGroupService.findEnabledMessageRecipientGroup:input_type -> pb.FindEnabledMessageRecipientGroupRequest
	1, // 7: pb.MessageRecipientGroupService.createMessageRecipientGroup:output_type -> pb.CreateMessageRecipientGroupResponse
	9, // 8: pb.MessageRecipientGroupService.updateMessageRecipientGroup:output_type -> pb.RPCSuccess
	4, // 9: pb.MessageRecipientGroupService.findAllEnabledMessageRecipientGroups:output_type -> pb.FindAllEnabledMessageRecipientGroupsResponse
	9, // 10: pb.MessageRecipientGroupService.deleteMessageRecipientGroup:output_type -> pb.RPCSuccess
	7, // 11: pb.MessageRecipientGroupService.findEnabledMessageRecipientGroup:output_type -> pb.FindEnabledMessageRecipientGroupResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_message_recipient_group_proto_init() }
func file_service_message_recipient_group_proto_init() {
	if File_service_message_recipient_group_proto != nil {
		return
	}
	file_models_model_message_recipient_group_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_recipient_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRecipientGroupRequest); i {
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
		file_service_message_recipient_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRecipientGroupResponse); i {
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
		file_service_message_recipient_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRecipientGroupRequest); i {
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
		file_service_message_recipient_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledMessageRecipientGroupsRequest); i {
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
		file_service_message_recipient_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledMessageRecipientGroupsResponse); i {
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
		file_service_message_recipient_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRecipientGroupRequest); i {
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
		file_service_message_recipient_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageRecipientGroupRequest); i {
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
		file_service_message_recipient_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageRecipientGroupResponse); i {
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
			RawDescriptor: file_service_message_recipient_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_recipient_group_proto_goTypes,
		DependencyIndexes: file_service_message_recipient_group_proto_depIdxs,
		MessageInfos:      file_service_message_recipient_group_proto_msgTypes,
	}.Build()
	File_service_message_recipient_group_proto = out.File
	file_service_message_recipient_group_proto_rawDesc = nil
	file_service_message_recipient_group_proto_goTypes = nil
	file_service_message_recipient_group_proto_depIdxs = nil
}
