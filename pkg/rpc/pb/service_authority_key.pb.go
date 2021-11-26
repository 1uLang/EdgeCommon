// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_authority_key.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 设置Key
type UpdateAuthorityKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	DayFrom      string   `protobuf:"bytes,2,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`
	DayTo        string   `protobuf:"bytes,3,opt,name=dayTo,proto3" json:"dayTo,omitempty"`
	Hostname     string   `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	MacAddresses []string `protobuf:"bytes,5,rep,name=macAddresses,proto3" json:"macAddresses,omitempty"`
	Company      string   `protobuf:"bytes,6,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *UpdateAuthorityKeyRequest) Reset() {
	*x = UpdateAuthorityKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthorityKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthorityKeyRequest) ProtoMessage() {}

func (x *UpdateAuthorityKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthorityKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthorityKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAuthorityKeyRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UpdateAuthorityKeyRequest) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *UpdateAuthorityKeyRequest) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *UpdateAuthorityKeyRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *UpdateAuthorityKeyRequest) GetMacAddresses() []string {
	if x != nil {
		return x.MacAddresses
	}
	return nil
}

func (x *UpdateAuthorityKeyRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

// 读取Key
type ReadAuthorityKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadAuthorityKeyRequest) Reset() {
	*x = ReadAuthorityKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAuthorityKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAuthorityKeyRequest) ProtoMessage() {}

func (x *ReadAuthorityKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAuthorityKeyRequest.ProtoReflect.Descriptor instead.
func (*ReadAuthorityKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{1}
}

type ReadAuthorityKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorityKey *AuthorityKey `protobuf:"bytes,1,opt,name=authorityKey,proto3" json:"authorityKey,omitempty"`
}

func (x *ReadAuthorityKeyResponse) Reset() {
	*x = ReadAuthorityKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAuthorityKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAuthorityKeyResponse) ProtoMessage() {}

func (x *ReadAuthorityKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAuthorityKeyResponse.ProtoReflect.Descriptor instead.
func (*ReadAuthorityKeyResponse) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{2}
}

func (x *ReadAuthorityKeyResponse) GetAuthorityKey() *AuthorityKey {
	if x != nil {
		return x.AuthorityKey
	}
	return nil
}

// 重置Key
type ResetAuthorityKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetAuthorityKeyRequest) Reset() {
	*x = ResetAuthorityKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetAuthorityKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetAuthorityKeyRequest) ProtoMessage() {}

func (x *ResetAuthorityKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetAuthorityKeyRequest.ProtoReflect.Descriptor instead.
func (*ResetAuthorityKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{3}
}

// 校验Key
type ValidateAuthorityKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ValidateAuthorityKeyRequest) Reset() {
	*x = ValidateAuthorityKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAuthorityKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAuthorityKeyRequest) ProtoMessage() {}

func (x *ValidateAuthorityKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAuthorityKeyRequest.ProtoReflect.Descriptor instead.
func (*ValidateAuthorityKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateAuthorityKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ValidateAuthorityKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOk  bool   `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ValidateAuthorityKeyResponse) Reset() {
	*x = ValidateAuthorityKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_authority_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateAuthorityKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAuthorityKeyResponse) ProtoMessage() {}

func (x *ValidateAuthorityKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_authority_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAuthorityKeyResponse.ProtoReflect.Descriptor instead.
func (*ValidateAuthorityKeyResponse) Descriptor() ([]byte, []int) {
	return file_service_authority_key_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateAuthorityKeyResponse) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ValidateAuthorityKeyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_service_authority_key_proto protoreflect.FileDescriptor

var file_service_authority_key_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79,
	0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x19, 0x0a, 0x17,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x1b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x48, 0x0a, 0x1c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0xc7, 0x02, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a,
	0x10, 0x72, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x59, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_authority_key_proto_rawDescOnce sync.Once
	file_service_authority_key_proto_rawDescData = file_service_authority_key_proto_rawDesc
)

func file_service_authority_key_proto_rawDescGZIP() []byte {
	file_service_authority_key_proto_rawDescOnce.Do(func() {
		file_service_authority_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_authority_key_proto_rawDescData)
	})
	return file_service_authority_key_proto_rawDescData
}

var file_service_authority_key_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_authority_key_proto_goTypes = []interface{}{
	(*UpdateAuthorityKeyRequest)(nil),    // 0: pb.UpdateAuthorityKeyRequest
	(*ReadAuthorityKeyRequest)(nil),      // 1: pb.ReadAuthorityKeyRequest
	(*ReadAuthorityKeyResponse)(nil),     // 2: pb.ReadAuthorityKeyResponse
	(*ResetAuthorityKeyRequest)(nil),     // 3: pb.ResetAuthorityKeyRequest
	(*ValidateAuthorityKeyRequest)(nil),  // 4: pb.ValidateAuthorityKeyRequest
	(*ValidateAuthorityKeyResponse)(nil), // 5: pb.ValidateAuthorityKeyResponse
	(*AuthorityKey)(nil),                 // 6: pb.AuthorityKey
	(*RPCSuccess)(nil),                   // 7: pb.RPCSuccess
}
var file_service_authority_key_proto_depIdxs = []int32{
	6, // 0: pb.ReadAuthorityKeyResponse.authorityKey:type_name -> pb.AuthorityKey
	0, // 1: pb.AuthorityKeyService.updateAuthorityKey:input_type -> pb.UpdateAuthorityKeyRequest
	1, // 2: pb.AuthorityKeyService.readAuthorityKey:input_type -> pb.ReadAuthorityKeyRequest
	3, // 3: pb.AuthorityKeyService.resetAuthorityKey:input_type -> pb.ResetAuthorityKeyRequest
	4, // 4: pb.AuthorityKeyService.validateAuthorityKey:input_type -> pb.ValidateAuthorityKeyRequest
	7, // 5: pb.AuthorityKeyService.updateAuthorityKey:output_type -> pb.RPCSuccess
	2, // 6: pb.AuthorityKeyService.readAuthorityKey:output_type -> pb.ReadAuthorityKeyResponse
	7, // 7: pb.AuthorityKeyService.resetAuthorityKey:output_type -> pb.RPCSuccess
	5, // 8: pb.AuthorityKeyService.validateAuthorityKey:output_type -> pb.ValidateAuthorityKeyResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_authority_key_proto_init() }
func file_service_authority_key_proto_init() {
	if File_service_authority_key_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_authority_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_authority_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthorityKeyRequest); i {
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
		file_service_authority_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAuthorityKeyRequest); i {
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
		file_service_authority_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAuthorityKeyResponse); i {
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
		file_service_authority_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetAuthorityKeyRequest); i {
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
		file_service_authority_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAuthorityKeyRequest); i {
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
		file_service_authority_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateAuthorityKeyResponse); i {
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
			RawDescriptor: file_service_authority_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_authority_key_proto_goTypes,
		DependencyIndexes: file_service_authority_key_proto_depIdxs,
		MessageInfos:      file_service_authority_key_proto_msgTypes,
	}.Build()
	File_service_authority_key_proto = out.File
	file_service_authority_key_proto_rawDesc = nil
	file_service_authority_key_proto_goTypes = nil
	file_service_authority_key_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorityKeyServiceClient is the client API for AuthorityKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorityKeyServiceClient interface {
	// 设置Key
	UpdateAuthorityKey(ctx context.Context, in *UpdateAuthorityKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取Key
	ReadAuthorityKey(ctx context.Context, in *ReadAuthorityKeyRequest, opts ...grpc.CallOption) (*ReadAuthorityKeyResponse, error)
	// 重置Key
	ResetAuthorityKey(ctx context.Context, in *ResetAuthorityKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 校验Key
	ValidateAuthorityKey(ctx context.Context, in *ValidateAuthorityKeyRequest, opts ...grpc.CallOption) (*ValidateAuthorityKeyResponse, error)
}

type authorityKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityKeyServiceClient(cc grpc.ClientConnInterface) AuthorityKeyServiceClient {
	return &authorityKeyServiceClient{cc}
}

func (c *authorityKeyServiceClient) UpdateAuthorityKey(ctx context.Context, in *UpdateAuthorityKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.AuthorityKeyService/updateAuthorityKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityKeyServiceClient) ReadAuthorityKey(ctx context.Context, in *ReadAuthorityKeyRequest, opts ...grpc.CallOption) (*ReadAuthorityKeyResponse, error) {
	out := new(ReadAuthorityKeyResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityKeyService/readAuthorityKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityKeyServiceClient) ResetAuthorityKey(ctx context.Context, in *ResetAuthorityKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.AuthorityKeyService/resetAuthorityKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorityKeyServiceClient) ValidateAuthorityKey(ctx context.Context, in *ValidateAuthorityKeyRequest, opts ...grpc.CallOption) (*ValidateAuthorityKeyResponse, error) {
	out := new(ValidateAuthorityKeyResponse)
	err := c.cc.Invoke(ctx, "/pb.AuthorityKeyService/validateAuthorityKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityKeyServiceServer is the server API for AuthorityKeyService service.
type AuthorityKeyServiceServer interface {
	// 设置Key
	UpdateAuthorityKey(context.Context, *UpdateAuthorityKeyRequest) (*RPCSuccess, error)
	// 读取Key
	ReadAuthorityKey(context.Context, *ReadAuthorityKeyRequest) (*ReadAuthorityKeyResponse, error)
	// 重置Key
	ResetAuthorityKey(context.Context, *ResetAuthorityKeyRequest) (*RPCSuccess, error)
	// 校验Key
	ValidateAuthorityKey(context.Context, *ValidateAuthorityKeyRequest) (*ValidateAuthorityKeyResponse, error)
}

// UnimplementedAuthorityKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorityKeyServiceServer struct {
}

func (*UnimplementedAuthorityKeyServiceServer) UpdateAuthorityKey(context.Context, *UpdateAuthorityKeyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthorityKey not implemented")
}
func (*UnimplementedAuthorityKeyServiceServer) ReadAuthorityKey(context.Context, *ReadAuthorityKeyRequest) (*ReadAuthorityKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAuthorityKey not implemented")
}
func (*UnimplementedAuthorityKeyServiceServer) ResetAuthorityKey(context.Context, *ResetAuthorityKeyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAuthorityKey not implemented")
}
func (*UnimplementedAuthorityKeyServiceServer) ValidateAuthorityKey(context.Context, *ValidateAuthorityKeyRequest) (*ValidateAuthorityKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAuthorityKey not implemented")
}

func RegisterAuthorityKeyServiceServer(s *grpc.Server, srv AuthorityKeyServiceServer) {
	s.RegisterService(&_AuthorityKeyService_serviceDesc, srv)
}

func _AuthorityKeyService_UpdateAuthorityKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorityKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityKeyServiceServer).UpdateAuthorityKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityKeyService/UpdateAuthorityKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityKeyServiceServer).UpdateAuthorityKey(ctx, req.(*UpdateAuthorityKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityKeyService_ReadAuthorityKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAuthorityKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityKeyServiceServer).ReadAuthorityKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityKeyService/ReadAuthorityKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityKeyServiceServer).ReadAuthorityKey(ctx, req.(*ReadAuthorityKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityKeyService_ResetAuthorityKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAuthorityKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityKeyServiceServer).ResetAuthorityKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityKeyService/ResetAuthorityKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityKeyServiceServer).ResetAuthorityKey(ctx, req.(*ResetAuthorityKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorityKeyService_ValidateAuthorityKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAuthorityKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityKeyServiceServer).ValidateAuthorityKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthorityKeyService/ValidateAuthorityKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityKeyServiceServer).ValidateAuthorityKey(ctx, req.(*ValidateAuthorityKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthorityKeyService",
	HandlerType: (*AuthorityKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "updateAuthorityKey",
			Handler:    _AuthorityKeyService_UpdateAuthorityKey_Handler,
		},
		{
			MethodName: "readAuthorityKey",
			Handler:    _AuthorityKeyService_ReadAuthorityKey_Handler,
		},
		{
			MethodName: "resetAuthorityKey",
			Handler:    _AuthorityKeyService_ResetAuthorityKey_Handler,
		},
		{
			MethodName: "validateAuthorityKey",
			Handler:    _AuthorityKeyService_ValidateAuthorityKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_authority_key.proto",
}
