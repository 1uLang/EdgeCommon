// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_message_media_instance.proto

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
type CreateMessageMediaInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MediaType   string  `protobuf:"bytes,2,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	ParamsJSON  []byte  `protobuf:"bytes,3,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	GroupIds    []int64 `protobuf:"varint,4,rep,packed,name=groupIds,proto3" json:"groupIds,omitempty"`
	Description string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	RateJSON    []byte  `protobuf:"bytes,6,opt,name=rateJSON,proto3" json:"rateJSON,omitempty"`
	HashLife    int32   `protobuf:"varint,7,opt,name=hashLife,proto3" json:"hashLife,omitempty"`
}

func (x *CreateMessageMediaInstanceRequest) Reset() {
	*x = CreateMessageMediaInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageMediaInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageMediaInstanceRequest) ProtoMessage() {}

func (x *CreateMessageMediaInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageMediaInstanceRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageMediaInstanceRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageMediaInstanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMessageMediaInstanceRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *CreateMessageMediaInstanceRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *CreateMessageMediaInstanceRequest) GetGroupIds() []int64 {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

func (x *CreateMessageMediaInstanceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMessageMediaInstanceRequest) GetRateJSON() []byte {
	if x != nil {
		return x.RateJSON
	}
	return nil
}

func (x *CreateMessageMediaInstanceRequest) GetHashLife() int32 {
	if x != nil {
		return x.HashLife
	}
	return 0
}

type CreateMessageMediaInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstanceId int64 `protobuf:"varint,1,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
}

func (x *CreateMessageMediaInstanceResponse) Reset() {
	*x = CreateMessageMediaInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageMediaInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageMediaInstanceResponse) ProtoMessage() {}

func (x *CreateMessageMediaInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageMediaInstanceResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageMediaInstanceResponse) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageMediaInstanceResponse) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

// 修改接收人
type UpdateMessageMediaInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstanceId int64  `protobuf:"varint,1,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
	Name                   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MediaType              string `protobuf:"bytes,3,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	ParamsJSON             []byte `protobuf:"bytes,4,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	Description            string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	RateJSON               []byte `protobuf:"bytes,7,opt,name=rateJSON,proto3" json:"rateJSON,omitempty"`
	HashLife               int32  `protobuf:"varint,8,opt,name=hashLife,proto3" json:"hashLife,omitempty"`
	IsOn                   bool   `protobuf:"varint,6,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateMessageMediaInstanceRequest) Reset() {
	*x = UpdateMessageMediaInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageMediaInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageMediaInstanceRequest) ProtoMessage() {}

func (x *UpdateMessageMediaInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageMediaInstanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageMediaInstanceRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMessageMediaInstanceRequest) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

func (x *UpdateMessageMediaInstanceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMessageMediaInstanceRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *UpdateMessageMediaInstanceRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *UpdateMessageMediaInstanceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMessageMediaInstanceRequest) GetRateJSON() []byte {
	if x != nil {
		return x.RateJSON
	}
	return nil
}

func (x *UpdateMessageMediaInstanceRequest) GetHashLife() int32 {
	if x != nil {
		return x.HashLife
	}
	return 0
}

func (x *UpdateMessageMediaInstanceRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 删除接收人
type DeleteMessageMediaInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstanceId int64 `protobuf:"varint,1,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
}

func (x *DeleteMessageMediaInstanceRequest) Reset() {
	*x = DeleteMessageMediaInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageMediaInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageMediaInstanceRequest) ProtoMessage() {}

func (x *DeleteMessageMediaInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageMediaInstanceRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageMediaInstanceRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMessageMediaInstanceRequest) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

// 计算接收人数量
type CountAllEnabledMessageMediaInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType string `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Keyword   string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *CountAllEnabledMessageMediaInstancesRequest) Reset() {
	*x = CountAllEnabledMessageMediaInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledMessageMediaInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledMessageMediaInstancesRequest) ProtoMessage() {}

func (x *CountAllEnabledMessageMediaInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledMessageMediaInstancesRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledMessageMediaInstancesRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{4}
}

func (x *CountAllEnabledMessageMediaInstancesRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *CountAllEnabledMessageMediaInstancesRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// 列出单页接收人
type ListEnabledMessageMediaInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType string `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Keyword   string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Offset    int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size      int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListEnabledMessageMediaInstancesRequest) Reset() {
	*x = ListEnabledMessageMediaInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledMessageMediaInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledMessageMediaInstancesRequest) ProtoMessage() {}

func (x *ListEnabledMessageMediaInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledMessageMediaInstancesRequest.ProtoReflect.Descriptor instead.
func (*ListEnabledMessageMediaInstancesRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{5}
}

func (x *ListEnabledMessageMediaInstancesRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ListEnabledMessageMediaInstancesRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListEnabledMessageMediaInstancesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEnabledMessageMediaInstancesRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListEnabledMessageMediaInstancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstances []*MessageMediaInstance `protobuf:"bytes,1,rep,name=messageMediaInstances,proto3" json:"messageMediaInstances,omitempty"`
}

func (x *ListEnabledMessageMediaInstancesResponse) Reset() {
	*x = ListEnabledMessageMediaInstancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledMessageMediaInstancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledMessageMediaInstancesResponse) ProtoMessage() {}

func (x *ListEnabledMessageMediaInstancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledMessageMediaInstancesResponse.ProtoReflect.Descriptor instead.
func (*ListEnabledMessageMediaInstancesResponse) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{6}
}

func (x *ListEnabledMessageMediaInstancesResponse) GetMessageMediaInstances() []*MessageMediaInstance {
	if x != nil {
		return x.MessageMediaInstances
	}
	return nil
}

// 查找单个接收人信息
type FindEnabledMessageMediaInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstanceId int64 `protobuf:"varint,1,opt,name=messageMediaInstanceId,proto3" json:"messageMediaInstanceId,omitempty"`
}

func (x *FindEnabledMessageMediaInstanceRequest) Reset() {
	*x = FindEnabledMessageMediaInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageMediaInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageMediaInstanceRequest) ProtoMessage() {}

func (x *FindEnabledMessageMediaInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageMediaInstanceRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageMediaInstanceRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledMessageMediaInstanceRequest) GetMessageMediaInstanceId() int64 {
	if x != nil {
		return x.MessageMediaInstanceId
	}
	return 0
}

type FindEnabledMessageMediaInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMediaInstance *MessageMediaInstance `protobuf:"bytes,1,opt,name=messageMediaInstance,proto3" json:"messageMediaInstance,omitempty"`
}

func (x *FindEnabledMessageMediaInstanceResponse) Reset() {
	*x = FindEnabledMessageMediaInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_instance_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageMediaInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageMediaInstanceResponse) ProtoMessage() {}

func (x *FindEnabledMessageMediaInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_instance_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageMediaInstanceResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageMediaInstanceResponse) Descriptor() ([]byte, []int) {
	return file_service_message_media_instance_proto_rawDescGZIP(), []int{8}
}

func (x *FindEnabledMessageMediaInstanceResponse) GetMessageMediaInstance() *MessageMediaInstance {
	if x != nil {
		return x.MessageMediaInstance
	}
	return nil
}

var File_service_message_media_instance_proto protoreflect.FileDescriptor

var file_service_message_media_instance_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x29, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xeb, 0x01, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x53,
	0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x66, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x4c, 0x69, 0x66, 0x65, 0x22, 0x5c,
	0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x9b, 0x02, 0x0a,
	0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61,
	0x73, 0x68, 0x4c, 0x69, 0x66, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x61,
	0x73, 0x68, 0x4c, 0x69, 0x66, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x5b, 0x0a, 0x21, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x2b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8d,
	0x01, 0x0a, 0x27, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x7a,
	0x0a, 0x28, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x26, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x27,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x9e, 0x05, 0x0a, 0x1b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x1a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x1a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x53, 0x0a, 0x1a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6d, 0x0a, 0x24,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x20, 0x6c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x1f, 0x66, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_media_instance_proto_rawDescOnce sync.Once
	file_service_message_media_instance_proto_rawDescData = file_service_message_media_instance_proto_rawDesc
)

func file_service_message_media_instance_proto_rawDescGZIP() []byte {
	file_service_message_media_instance_proto_rawDescOnce.Do(func() {
		file_service_message_media_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_media_instance_proto_rawDescData)
	})
	return file_service_message_media_instance_proto_rawDescData
}

var file_service_message_media_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_message_media_instance_proto_goTypes = []interface{}{
	(*CreateMessageMediaInstanceRequest)(nil),           // 0: pb.CreateMessageMediaInstanceRequest
	(*CreateMessageMediaInstanceResponse)(nil),          // 1: pb.CreateMessageMediaInstanceResponse
	(*UpdateMessageMediaInstanceRequest)(nil),           // 2: pb.UpdateMessageMediaInstanceRequest
	(*DeleteMessageMediaInstanceRequest)(nil),           // 3: pb.DeleteMessageMediaInstanceRequest
	(*CountAllEnabledMessageMediaInstancesRequest)(nil), // 4: pb.CountAllEnabledMessageMediaInstancesRequest
	(*ListEnabledMessageMediaInstancesRequest)(nil),     // 5: pb.ListEnabledMessageMediaInstancesRequest
	(*ListEnabledMessageMediaInstancesResponse)(nil),    // 6: pb.ListEnabledMessageMediaInstancesResponse
	(*FindEnabledMessageMediaInstanceRequest)(nil),      // 7: pb.FindEnabledMessageMediaInstanceRequest
	(*FindEnabledMessageMediaInstanceResponse)(nil),     // 8: pb.FindEnabledMessageMediaInstanceResponse
	(*MessageMediaInstance)(nil),                        // 9: pb.MessageMediaInstance
	(*RPCSuccess)(nil),                                  // 10: pb.RPCSuccess
	(*RPCCountResponse)(nil),                            // 11: pb.RPCCountResponse
}
var file_service_message_media_instance_proto_depIdxs = []int32{
	9,  // 0: pb.ListEnabledMessageMediaInstancesResponse.messageMediaInstances:type_name -> pb.MessageMediaInstance
	9,  // 1: pb.FindEnabledMessageMediaInstanceResponse.messageMediaInstance:type_name -> pb.MessageMediaInstance
	0,  // 2: pb.MessageMediaInstanceService.createMessageMediaInstance:input_type -> pb.CreateMessageMediaInstanceRequest
	2,  // 3: pb.MessageMediaInstanceService.updateMessageMediaInstance:input_type -> pb.UpdateMessageMediaInstanceRequest
	3,  // 4: pb.MessageMediaInstanceService.deleteMessageMediaInstance:input_type -> pb.DeleteMessageMediaInstanceRequest
	4,  // 5: pb.MessageMediaInstanceService.countAllEnabledMessageMediaInstances:input_type -> pb.CountAllEnabledMessageMediaInstancesRequest
	5,  // 6: pb.MessageMediaInstanceService.listEnabledMessageMediaInstances:input_type -> pb.ListEnabledMessageMediaInstancesRequest
	7,  // 7: pb.MessageMediaInstanceService.findEnabledMessageMediaInstance:input_type -> pb.FindEnabledMessageMediaInstanceRequest
	1,  // 8: pb.MessageMediaInstanceService.createMessageMediaInstance:output_type -> pb.CreateMessageMediaInstanceResponse
	10, // 9: pb.MessageMediaInstanceService.updateMessageMediaInstance:output_type -> pb.RPCSuccess
	10, // 10: pb.MessageMediaInstanceService.deleteMessageMediaInstance:output_type -> pb.RPCSuccess
	11, // 11: pb.MessageMediaInstanceService.countAllEnabledMessageMediaInstances:output_type -> pb.RPCCountResponse
	6,  // 12: pb.MessageMediaInstanceService.listEnabledMessageMediaInstances:output_type -> pb.ListEnabledMessageMediaInstancesResponse
	8,  // 13: pb.MessageMediaInstanceService.findEnabledMessageMediaInstance:output_type -> pb.FindEnabledMessageMediaInstanceResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_message_media_instance_proto_init() }
func file_service_message_media_instance_proto_init() {
	if File_service_message_media_instance_proto != nil {
		return
	}
	file_models_model_message_media_instance_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_media_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageMediaInstanceRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageMediaInstanceResponse); i {
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
		file_service_message_media_instance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageMediaInstanceRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageMediaInstanceRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledMessageMediaInstancesRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledMessageMediaInstancesRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledMessageMediaInstancesResponse); i {
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
		file_service_message_media_instance_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageMediaInstanceRequest); i {
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
		file_service_message_media_instance_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageMediaInstanceResponse); i {
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
			RawDescriptor: file_service_message_media_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_media_instance_proto_goTypes,
		DependencyIndexes: file_service_message_media_instance_proto_depIdxs,
		MessageInfos:      file_service_message_media_instance_proto_msgTypes,
	}.Build()
	File_service_message_media_instance_proto = out.File
	file_service_message_media_instance_proto_rawDesc = nil
	file_service_message_media_instance_proto_goTypes = nil
	file_service_message_media_instance_proto_depIdxs = nil
}
