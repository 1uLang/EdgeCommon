// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: service_user_account.proto

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

// 计算账户数量
type CountUserAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *CountUserAccountsRequest) Reset() {
	*x = CountUserAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserAccountsRequest) ProtoMessage() {}

func (x *CountUserAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserAccountsRequest.ProtoReflect.Descriptor instead.
func (*CountUserAccountsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *CountUserAccountsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// 列出单页账户
type ListUserAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Offset  int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUserAccountsRequest) Reset() {
	*x = ListUserAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserAccountsRequest) ProtoMessage() {}

func (x *ListUserAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListUserAccountsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserAccountsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListUserAccountsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserAccountsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUserAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccounts []*UserAccount `protobuf:"bytes,1,rep,name=userAccounts,proto3" json:"userAccounts,omitempty"`
}

func (x *ListUserAccountsResponse) Reset() {
	*x = ListUserAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserAccountsResponse) ProtoMessage() {}

func (x *ListUserAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListUserAccountsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserAccountsResponse) GetUserAccounts() []*UserAccount {
	if x != nil {
		return x.UserAccounts
	}
	return nil
}

// 根据用户ID查找单个账户
type FindEnabledUserAccountWithUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindEnabledUserAccountWithUserIdRequest) Reset() {
	*x = FindEnabledUserAccountWithUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledUserAccountWithUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledUserAccountWithUserIdRequest) ProtoMessage() {}

func (x *FindEnabledUserAccountWithUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledUserAccountWithUserIdRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledUserAccountWithUserIdRequest) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledUserAccountWithUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindEnabledUserAccountWithUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccount *UserAccount `protobuf:"bytes,1,opt,name=userAccount,proto3" json:"userAccount,omitempty"`
}

func (x *FindEnabledUserAccountWithUserIdResponse) Reset() {
	*x = FindEnabledUserAccountWithUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledUserAccountWithUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledUserAccountWithUserIdResponse) ProtoMessage() {}

func (x *FindEnabledUserAccountWithUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledUserAccountWithUserIdResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledUserAccountWithUserIdResponse) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledUserAccountWithUserIdResponse) GetUserAccount() *UserAccount {
	if x != nil {
		return x.UserAccount
	}
	return nil
}

// 查找单个账户
type FindEnabledUserAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccountId int64 `protobuf:"varint,1,opt,name=userAccountId,proto3" json:"userAccountId,omitempty"`
}

func (x *FindEnabledUserAccountRequest) Reset() {
	*x = FindEnabledUserAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledUserAccountRequest) ProtoMessage() {}

func (x *FindEnabledUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledUserAccountRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledUserAccountRequest) GetUserAccountId() int64 {
	if x != nil {
		return x.UserAccountId
	}
	return 0
}

type FindEnabledUserAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccount *UserAccount `protobuf:"bytes,1,opt,name=userAccount,proto3" json:"userAccount,omitempty"`
}

func (x *FindEnabledUserAccountResponse) Reset() {
	*x = FindEnabledUserAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledUserAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledUserAccountResponse) ProtoMessage() {}

func (x *FindEnabledUserAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledUserAccountResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledUserAccountResponse) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledUserAccountResponse) GetUserAccount() *UserAccount {
	if x != nil {
		return x.UserAccount
	}
	return nil
}

// 修改用户账户
type UpdateUserAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccountId int64   `protobuf:"varint,1,opt,name=userAccountId,proto3" json:"userAccountId,omitempty"`
	Delta         float32 `protobuf:"fixed32,2,opt,name=delta,proto3" json:"delta,omitempty"`
	EventType     string  `protobuf:"bytes,3,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Description   string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ParamsJSON    []byte  `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
}

func (x *UpdateUserAccountRequest) Reset() {
	*x = UpdateUserAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserAccountRequest) ProtoMessage() {}

func (x *UpdateUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_service_user_account_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserAccountRequest) GetUserAccountId() int64 {
	if x != nil {
		return x.UserAccountId
	}
	return 0
}

func (x *UpdateUserAccountRequest) GetDelta() float32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

func (x *UpdateUserAccountRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *UpdateUserAccountRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateUserAccountRequest) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

var File_service_user_account_proto protoreflect.FileDescriptor

var file_service_user_account_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x18,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x5f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x28, 0x46, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x53, 0x0a,
	0x1e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0xcf, 0x03, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x6c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x20, 0x66, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_account_proto_rawDescOnce sync.Once
	file_service_user_account_proto_rawDescData = file_service_user_account_proto_rawDesc
)

func file_service_user_account_proto_rawDescGZIP() []byte {
	file_service_user_account_proto_rawDescOnce.Do(func() {
		file_service_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_account_proto_rawDescData)
	})
	return file_service_user_account_proto_rawDescData
}

var file_service_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_user_account_proto_goTypes = []interface{}{
	(*CountUserAccountsRequest)(nil),                 // 0: pb.CountUserAccountsRequest
	(*ListUserAccountsRequest)(nil),                  // 1: pb.ListUserAccountsRequest
	(*ListUserAccountsResponse)(nil),                 // 2: pb.ListUserAccountsResponse
	(*FindEnabledUserAccountWithUserIdRequest)(nil),  // 3: pb.FindEnabledUserAccountWithUserIdRequest
	(*FindEnabledUserAccountWithUserIdResponse)(nil), // 4: pb.FindEnabledUserAccountWithUserIdResponse
	(*FindEnabledUserAccountRequest)(nil),            // 5: pb.FindEnabledUserAccountRequest
	(*FindEnabledUserAccountResponse)(nil),           // 6: pb.FindEnabledUserAccountResponse
	(*UpdateUserAccountRequest)(nil),                 // 7: pb.UpdateUserAccountRequest
	(*UserAccount)(nil),                              // 8: pb.UserAccount
	(*RPCCountResponse)(nil),                         // 9: pb.RPCCountResponse
	(*RPCSuccess)(nil),                               // 10: pb.RPCSuccess
}
var file_service_user_account_proto_depIdxs = []int32{
	8,  // 0: pb.ListUserAccountsResponse.userAccounts:type_name -> pb.UserAccount
	8,  // 1: pb.FindEnabledUserAccountWithUserIdResponse.userAccount:type_name -> pb.UserAccount
	8,  // 2: pb.FindEnabledUserAccountResponse.userAccount:type_name -> pb.UserAccount
	0,  // 3: pb.UserAccountService.countUserAccounts:input_type -> pb.CountUserAccountsRequest
	1,  // 4: pb.UserAccountService.listUserAccounts:input_type -> pb.ListUserAccountsRequest
	3,  // 5: pb.UserAccountService.findEnabledUserAccountWithUserId:input_type -> pb.FindEnabledUserAccountWithUserIdRequest
	5,  // 6: pb.UserAccountService.findEnabledUserAccount:input_type -> pb.FindEnabledUserAccountRequest
	7,  // 7: pb.UserAccountService.updateUserAccount:input_type -> pb.UpdateUserAccountRequest
	9,  // 8: pb.UserAccountService.countUserAccounts:output_type -> pb.RPCCountResponse
	2,  // 9: pb.UserAccountService.listUserAccounts:output_type -> pb.ListUserAccountsResponse
	4,  // 10: pb.UserAccountService.findEnabledUserAccountWithUserId:output_type -> pb.FindEnabledUserAccountWithUserIdResponse
	6,  // 11: pb.UserAccountService.findEnabledUserAccount:output_type -> pb.FindEnabledUserAccountResponse
	10, // 12: pb.UserAccountService.updateUserAccount:output_type -> pb.RPCSuccess
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_user_account_proto_init() }
func file_service_user_account_proto_init() {
	if File_service_user_account_proto != nil {
		return
	}
	file_models_model_user_account_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserAccountsRequest); i {
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
		file_service_user_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserAccountsRequest); i {
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
		file_service_user_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserAccountsResponse); i {
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
		file_service_user_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledUserAccountWithUserIdRequest); i {
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
		file_service_user_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledUserAccountWithUserIdResponse); i {
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
		file_service_user_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledUserAccountRequest); i {
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
		file_service_user_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledUserAccountResponse); i {
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
		file_service_user_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserAccountRequest); i {
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
			RawDescriptor: file_service_user_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_account_proto_goTypes,
		DependencyIndexes: file_service_user_account_proto_depIdxs,
		MessageInfos:      file_service_user_account_proto_msgTypes,
	}.Build()
	File_service_user_account_proto = out.File
	file_service_user_account_proto_rawDesc = nil
	file_service_user_account_proto_goTypes = nil
	file_service_user_account_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserAccountServiceClient is the client API for UserAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAccountServiceClient interface {
	// 计算账户数量
	CountUserAccounts(ctx context.Context, in *CountUserAccountsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页账户
	ListUserAccounts(ctx context.Context, in *ListUserAccountsRequest, opts ...grpc.CallOption) (*ListUserAccountsResponse, error)
	// 根据用户ID查找单个账户
	FindEnabledUserAccountWithUserId(ctx context.Context, in *FindEnabledUserAccountWithUserIdRequest, opts ...grpc.CallOption) (*FindEnabledUserAccountWithUserIdResponse, error)
	// 查找单个账户
	FindEnabledUserAccount(ctx context.Context, in *FindEnabledUserAccountRequest, opts ...grpc.CallOption) (*FindEnabledUserAccountResponse, error)
	// 修改用户账户
	UpdateUserAccount(ctx context.Context, in *UpdateUserAccountRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type userAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountServiceClient(cc grpc.ClientConnInterface) UserAccountServiceClient {
	return &userAccountServiceClient{cc}
}

func (c *userAccountServiceClient) CountUserAccounts(ctx context.Context, in *CountUserAccountsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserAccountService/countUserAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceClient) ListUserAccounts(ctx context.Context, in *ListUserAccountsRequest, opts ...grpc.CallOption) (*ListUserAccountsResponse, error) {
	out := new(ListUserAccountsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserAccountService/listUserAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceClient) FindEnabledUserAccountWithUserId(ctx context.Context, in *FindEnabledUserAccountWithUserIdRequest, opts ...grpc.CallOption) (*FindEnabledUserAccountWithUserIdResponse, error) {
	out := new(FindEnabledUserAccountWithUserIdResponse)
	err := c.cc.Invoke(ctx, "/pb.UserAccountService/findEnabledUserAccountWithUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceClient) FindEnabledUserAccount(ctx context.Context, in *FindEnabledUserAccountRequest, opts ...grpc.CallOption) (*FindEnabledUserAccountResponse, error) {
	out := new(FindEnabledUserAccountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserAccountService/findEnabledUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceClient) UpdateUserAccount(ctx context.Context, in *UpdateUserAccountRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserAccountService/updateUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountServiceServer is the server API for UserAccountService service.
type UserAccountServiceServer interface {
	// 计算账户数量
	CountUserAccounts(context.Context, *CountUserAccountsRequest) (*RPCCountResponse, error)
	// 列出单页账户
	ListUserAccounts(context.Context, *ListUserAccountsRequest) (*ListUserAccountsResponse, error)
	// 根据用户ID查找单个账户
	FindEnabledUserAccountWithUserId(context.Context, *FindEnabledUserAccountWithUserIdRequest) (*FindEnabledUserAccountWithUserIdResponse, error)
	// 查找单个账户
	FindEnabledUserAccount(context.Context, *FindEnabledUserAccountRequest) (*FindEnabledUserAccountResponse, error)
	// 修改用户账户
	UpdateUserAccount(context.Context, *UpdateUserAccountRequest) (*RPCSuccess, error)
}

// UnimplementedUserAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserAccountServiceServer struct {
}

func (*UnimplementedUserAccountServiceServer) CountUserAccounts(context.Context, *CountUserAccountsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUserAccounts not implemented")
}
func (*UnimplementedUserAccountServiceServer) ListUserAccounts(context.Context, *ListUserAccountsRequest) (*ListUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserAccounts not implemented")
}
func (*UnimplementedUserAccountServiceServer) FindEnabledUserAccountWithUserId(context.Context, *FindEnabledUserAccountWithUserIdRequest) (*FindEnabledUserAccountWithUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledUserAccountWithUserId not implemented")
}
func (*UnimplementedUserAccountServiceServer) FindEnabledUserAccount(context.Context, *FindEnabledUserAccountRequest) (*FindEnabledUserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledUserAccount not implemented")
}
func (*UnimplementedUserAccountServiceServer) UpdateUserAccount(context.Context, *UpdateUserAccountRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAccount not implemented")
}

func RegisterUserAccountServiceServer(s *grpc.Server, srv UserAccountServiceServer) {
	s.RegisterService(&_UserAccountService_serviceDesc, srv)
}

func _UserAccountService_CountUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).CountUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserAccountService/CountUserAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).CountUserAccounts(ctx, req.(*CountUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountService_ListUserAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).ListUserAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserAccountService/ListUserAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).ListUserAccounts(ctx, req.(*ListUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountService_FindEnabledUserAccountWithUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledUserAccountWithUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).FindEnabledUserAccountWithUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserAccountService/FindEnabledUserAccountWithUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).FindEnabledUserAccountWithUserId(ctx, req.(*FindEnabledUserAccountWithUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountService_FindEnabledUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).FindEnabledUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserAccountService/FindEnabledUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).FindEnabledUserAccount(ctx, req.(*FindEnabledUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountService_UpdateUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).UpdateUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserAccountService/UpdateUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).UpdateUserAccount(ctx, req.(*UpdateUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserAccountService",
	HandlerType: (*UserAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "countUserAccounts",
			Handler:    _UserAccountService_CountUserAccounts_Handler,
		},
		{
			MethodName: "listUserAccounts",
			Handler:    _UserAccountService_ListUserAccounts_Handler,
		},
		{
			MethodName: "findEnabledUserAccountWithUserId",
			Handler:    _UserAccountService_FindEnabledUserAccountWithUserId_Handler,
		},
		{
			MethodName: "findEnabledUserAccount",
			Handler:    _UserAccountService_FindEnabledUserAccount_Handler,
		},
		{
			MethodName: "updateUserAccount",
			Handler:    _UserAccountService_UpdateUserAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_account.proto",
}
