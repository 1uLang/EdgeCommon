// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_user_access_key.proto

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

// 创建AccessKey
type CreateUserAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AdminId     int64  `protobuf:"varint,3,opt,name=adminId,proto3" json:"adminId,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateUserAccessKeyRequest) Reset() {
	*x = CreateUserAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserAccessKeyRequest) ProtoMessage() {}

func (x *CreateUserAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateUserAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserAccessKeyRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserAccessKeyRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CreateUserAccessKeyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateUserAccessKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccessKeyId int64 `protobuf:"varint,1,opt,name=userAccessKeyId,proto3" json:"userAccessKeyId,omitempty"`
}

func (x *CreateUserAccessKeyResponse) Reset() {
	*x = CreateUserAccessKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserAccessKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserAccessKeyResponse) ProtoMessage() {}

func (x *CreateUserAccessKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserAccessKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateUserAccessKeyResponse) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserAccessKeyResponse) GetUserAccessKeyId() int64 {
	if x != nil {
		return x.UserAccessKeyId
	}
	return 0
}

// 查找所有的AccessKey
type FindAllEnabledUserAccessKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AdminId int64 `protobuf:"varint,2,opt,name=adminId,proto3" json:"adminId,omitempty"`
}

func (x *FindAllEnabledUserAccessKeysRequest) Reset() {
	*x = FindAllEnabledUserAccessKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledUserAccessKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledUserAccessKeysRequest) ProtoMessage() {}

func (x *FindAllEnabledUserAccessKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledUserAccessKeysRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledUserAccessKeysRequest) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllEnabledUserAccessKeysRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindAllEnabledUserAccessKeysRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type FindAllEnabledUserAccessKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccessKeys []*UserAccessKey `protobuf:"bytes,1,rep,name=userAccessKeys,proto3" json:"userAccessKeys,omitempty"`
}

func (x *FindAllEnabledUserAccessKeysResponse) Reset() {
	*x = FindAllEnabledUserAccessKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledUserAccessKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledUserAccessKeysResponse) ProtoMessage() {}

func (x *FindAllEnabledUserAccessKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledUserAccessKeysResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledUserAccessKeysResponse) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllEnabledUserAccessKeysResponse) GetUserAccessKeys() []*UserAccessKey {
	if x != nil {
		return x.UserAccessKeys
	}
	return nil
}

// 删除AccessKey
type DeleteUserAccessKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccessKeyId int64 `protobuf:"varint,1,opt,name=userAccessKeyId,proto3" json:"userAccessKeyId,omitempty"`
}

func (x *DeleteUserAccessKeyRequest) Reset() {
	*x = DeleteUserAccessKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAccessKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAccessKeyRequest) ProtoMessage() {}

func (x *DeleteUserAccessKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAccessKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserAccessKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserAccessKeyRequest) GetUserAccessKeyId() int64 {
	if x != nil {
		return x.UserAccessKeyId
	}
	return 0
}

// 设置是否启用AccessKey
type UpdateUserAccessKeyIsOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccessKeyId int64 `protobuf:"varint,1,opt,name=userAccessKeyId,proto3" json:"userAccessKeyId,omitempty"`
	IsOn            bool  `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateUserAccessKeyIsOnRequest) Reset() {
	*x = UpdateUserAccessKeyIsOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserAccessKeyIsOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserAccessKeyIsOnRequest) ProtoMessage() {}

func (x *UpdateUserAccessKeyIsOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserAccessKeyIsOnRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserAccessKeyIsOnRequest) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserAccessKeyIsOnRequest) GetUserAccessKeyId() int64 {
	if x != nil {
		return x.UserAccessKeyId
	}
	return 0
}

func (x *UpdateUserAccessKeyIsOnRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 计算AccessKey数量
type CountAllEnabledUserAccessKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId int64 `protobuf:"varint,1,opt,name=adminId,proto3" json:"adminId,omitempty"`
	UserId  int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CountAllEnabledUserAccessKeysRequest) Reset() {
	*x = CountAllEnabledUserAccessKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_access_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledUserAccessKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledUserAccessKeysRequest) ProtoMessage() {}

func (x *CountAllEnabledUserAccessKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_access_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledUserAccessKeysRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledUserAccessKeysRequest) Descriptor() ([]byte, []int) {
	return file_service_user_access_key_proto_rawDescGZIP(), []int{6}
}

func (x *CountAllEnabledUserAccessKeysRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CountAllEnabledUserAccessKeysRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_service_user_access_key_proto protoreflect.FileDescriptor

var file_service_user_access_key_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x70, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x57, 0x0a,
	0x23, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x46, 0x0a, 0x1a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49,
	0x64, 0x22, 0x5e, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x22, 0x58, 0x0a, 0x24, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xd8, 0x03, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x1c,
	0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x27, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x73, 0x4f,
	0x6e, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x1d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_access_key_proto_rawDescOnce sync.Once
	file_service_user_access_key_proto_rawDescData = file_service_user_access_key_proto_rawDesc
)

func file_service_user_access_key_proto_rawDescGZIP() []byte {
	file_service_user_access_key_proto_rawDescOnce.Do(func() {
		file_service_user_access_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_access_key_proto_rawDescData)
	})
	return file_service_user_access_key_proto_rawDescData
}

var file_service_user_access_key_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_user_access_key_proto_goTypes = []interface{}{
	(*CreateUserAccessKeyRequest)(nil),           // 0: pb.CreateUserAccessKeyRequest
	(*CreateUserAccessKeyResponse)(nil),          // 1: pb.CreateUserAccessKeyResponse
	(*FindAllEnabledUserAccessKeysRequest)(nil),  // 2: pb.FindAllEnabledUserAccessKeysRequest
	(*FindAllEnabledUserAccessKeysResponse)(nil), // 3: pb.FindAllEnabledUserAccessKeysResponse
	(*DeleteUserAccessKeyRequest)(nil),           // 4: pb.DeleteUserAccessKeyRequest
	(*UpdateUserAccessKeyIsOnRequest)(nil),       // 5: pb.UpdateUserAccessKeyIsOnRequest
	(*CountAllEnabledUserAccessKeysRequest)(nil), // 6: pb.CountAllEnabledUserAccessKeysRequest
	(*UserAccessKey)(nil),                        // 7: pb.UserAccessKey
	(*RPCSuccess)(nil),                           // 8: pb.RPCSuccess
	(*RPCCountResponse)(nil),                     // 9: pb.RPCCountResponse
}
var file_service_user_access_key_proto_depIdxs = []int32{
	7, // 0: pb.FindAllEnabledUserAccessKeysResponse.userAccessKeys:type_name -> pb.UserAccessKey
	0, // 1: pb.UserAccessKeyService.createUserAccessKey:input_type -> pb.CreateUserAccessKeyRequest
	2, // 2: pb.UserAccessKeyService.findAllEnabledUserAccessKeys:input_type -> pb.FindAllEnabledUserAccessKeysRequest
	4, // 3: pb.UserAccessKeyService.deleteUserAccessKey:input_type -> pb.DeleteUserAccessKeyRequest
	5, // 4: pb.UserAccessKeyService.updateUserAccessKeyIsOn:input_type -> pb.UpdateUserAccessKeyIsOnRequest
	6, // 5: pb.UserAccessKeyService.countAllEnabledUserAccessKeys:input_type -> pb.CountAllEnabledUserAccessKeysRequest
	1, // 6: pb.UserAccessKeyService.createUserAccessKey:output_type -> pb.CreateUserAccessKeyResponse
	3, // 7: pb.UserAccessKeyService.findAllEnabledUserAccessKeys:output_type -> pb.FindAllEnabledUserAccessKeysResponse
	8, // 8: pb.UserAccessKeyService.deleteUserAccessKey:output_type -> pb.RPCSuccess
	8, // 9: pb.UserAccessKeyService.updateUserAccessKeyIsOn:output_type -> pb.RPCSuccess
	9, // 10: pb.UserAccessKeyService.countAllEnabledUserAccessKeys:output_type -> pb.RPCCountResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_access_key_proto_init() }
func file_service_user_access_key_proto_init() {
	if File_service_user_access_key_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_user_access_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_access_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserAccessKeyRequest); i {
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
		file_service_user_access_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserAccessKeyResponse); i {
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
		file_service_user_access_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledUserAccessKeysRequest); i {
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
		file_service_user_access_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledUserAccessKeysResponse); i {
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
		file_service_user_access_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAccessKeyRequest); i {
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
		file_service_user_access_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserAccessKeyIsOnRequest); i {
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
		file_service_user_access_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledUserAccessKeysRequest); i {
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
			RawDescriptor: file_service_user_access_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_access_key_proto_goTypes,
		DependencyIndexes: file_service_user_access_key_proto_depIdxs,
		MessageInfos:      file_service_user_access_key_proto_msgTypes,
	}.Build()
	File_service_user_access_key_proto = out.File
	file_service_user_access_key_proto_rawDesc = nil
	file_service_user_access_key_proto_goTypes = nil
	file_service_user_access_key_proto_depIdxs = nil
}
