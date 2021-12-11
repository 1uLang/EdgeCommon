// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service_http_gzip.proto

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

// 创建Gzip配置
type CreateHTTPGzipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level     int32         `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	MinLength *SizeCapacity `protobuf:"bytes,2,opt,name=minLength,proto3" json:"minLength,omitempty"`
	MaxLength *SizeCapacity `protobuf:"bytes,3,opt,name=maxLength,proto3" json:"maxLength,omitempty"`
	CondsJSON []byte        `protobuf:"bytes,4,opt,name=condsJSON,proto3" json:"condsJSON,omitempty"`
}

func (x *CreateHTTPGzipRequest) Reset() {
	*x = CreateHTTPGzipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_gzip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPGzipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPGzipRequest) ProtoMessage() {}

func (x *CreateHTTPGzipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_gzip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPGzipRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPGzipRequest) Descriptor() ([]byte, []int) {
	return file_service_http_gzip_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPGzipRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CreateHTTPGzipRequest) GetMinLength() *SizeCapacity {
	if x != nil {
		return x.MinLength
	}
	return nil
}

func (x *CreateHTTPGzipRequest) GetMaxLength() *SizeCapacity {
	if x != nil {
		return x.MaxLength
	}
	return nil
}

func (x *CreateHTTPGzipRequest) GetCondsJSON() []byte {
	if x != nil {
		return x.CondsJSON
	}
	return nil
}

type CreateHTTPGzipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpGzipId int64 `protobuf:"varint,1,opt,name=httpGzipId,proto3" json:"httpGzipId,omitempty"`
}

func (x *CreateHTTPGzipResponse) Reset() {
	*x = CreateHTTPGzipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_gzip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPGzipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPGzipResponse) ProtoMessage() {}

func (x *CreateHTTPGzipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_gzip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPGzipResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPGzipResponse) Descriptor() ([]byte, []int) {
	return file_service_http_gzip_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPGzipResponse) GetHttpGzipId() int64 {
	if x != nil {
		return x.HttpGzipId
	}
	return 0
}

// 查找Gzip配置
type FindEnabledGzipConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpGzipId int64 `protobuf:"varint,1,opt,name=httpGzipId,proto3" json:"httpGzipId,omitempty"`
}

func (x *FindEnabledGzipConfigRequest) Reset() {
	*x = FindEnabledGzipConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_gzip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledGzipConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledGzipConfigRequest) ProtoMessage() {}

func (x *FindEnabledGzipConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_gzip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledGzipConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledGzipConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_gzip_proto_rawDescGZIP(), []int{2}
}

func (x *FindEnabledGzipConfigRequest) GetHttpGzipId() int64 {
	if x != nil {
		return x.HttpGzipId
	}
	return 0
}

type FindEnabledGzipConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpGzipJSON []byte `protobuf:"bytes,1,opt,name=httpGzipJSON,proto3" json:"httpGzipJSON,omitempty"`
}

func (x *FindEnabledGzipConfigResponse) Reset() {
	*x = FindEnabledGzipConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_gzip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledGzipConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledGzipConfigResponse) ProtoMessage() {}

func (x *FindEnabledGzipConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_gzip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledGzipConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledGzipConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_gzip_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledGzipConfigResponse) GetHttpGzipJSON() []byte {
	if x != nil {
		return x.HttpGzipJSON
	}
	return nil
}

// 修改Gzip配置
type UpdateHTTPGzipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpGzipId int64         `protobuf:"varint,1,opt,name=httpGzipId,proto3" json:"httpGzipId,omitempty"`
	Level      int32         `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MinLength  *SizeCapacity `protobuf:"bytes,3,opt,name=minLength,proto3" json:"minLength,omitempty"`
	MaxLength  *SizeCapacity `protobuf:"bytes,4,opt,name=maxLength,proto3" json:"maxLength,omitempty"`
	CondsJSON  []byte        `protobuf:"bytes,5,opt,name=condsJSON,proto3" json:"condsJSON,omitempty"`
}

func (x *UpdateHTTPGzipRequest) Reset() {
	*x = UpdateHTTPGzipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_gzip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPGzipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPGzipRequest) ProtoMessage() {}

func (x *UpdateHTTPGzipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_gzip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPGzipRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPGzipRequest) Descriptor() ([]byte, []int) {
	return file_service_http_gzip_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHTTPGzipRequest) GetHttpGzipId() int64 {
	if x != nil {
		return x.HttpGzipId
	}
	return 0
}

func (x *UpdateHTTPGzipRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *UpdateHTTPGzipRequest) GetMinLength() *SizeCapacity {
	if x != nil {
		return x.MinLength
	}
	return nil
}

func (x *UpdateHTTPGzipRequest) GetMaxLength() *SizeCapacity {
	if x != nil {
		return x.MaxLength
	}
	return nil
}

func (x *UpdateHTTPGzipRequest) GetCondsJSON() []byte {
	if x != nil {
		return x.CondsJSON
	}
	return nil
}

var File_service_http_gzip_proto protoreflect.FileDescriptor

var file_service_http_gzip_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67,
	0x7a, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x20, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x69,
	0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52,
	0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52,
	0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x38, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70,
	0x49, 0x64, 0x22, 0x3e, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x47, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70,
	0x49, 0x64, 0x22, 0x43, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x47, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x47,
	0x7a, 0x69, 0x70, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0xcb, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x47, 0x7a, 0x69, 0x70, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x09, 0x6d, 0x69,
	0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x7a, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0xf9, 0x01, 0x0a, 0x0f, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a,
	0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x19, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x47, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x47, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x47, 0x7a, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_http_gzip_proto_rawDescOnce sync.Once
	file_service_http_gzip_proto_rawDescData = file_service_http_gzip_proto_rawDesc
)

func file_service_http_gzip_proto_rawDescGZIP() []byte {
	file_service_http_gzip_proto_rawDescOnce.Do(func() {
		file_service_http_gzip_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_gzip_proto_rawDescData)
	})
	return file_service_http_gzip_proto_rawDescData
}

var file_service_http_gzip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_http_gzip_proto_goTypes = []interface{}{
	(*CreateHTTPGzipRequest)(nil),         // 0: pb.CreateHTTPGzipRequest
	(*CreateHTTPGzipResponse)(nil),        // 1: pb.CreateHTTPGzipResponse
	(*FindEnabledGzipConfigRequest)(nil),  // 2: pb.FindEnabledGzipConfigRequest
	(*FindEnabledGzipConfigResponse)(nil), // 3: pb.FindEnabledGzipConfigResponse
	(*UpdateHTTPGzipRequest)(nil),         // 4: pb.UpdateHTTPGzipRequest
	(*SizeCapacity)(nil),                  // 5: pb.SizeCapacity
	(*RPCSuccess)(nil),                    // 6: pb.RPCSuccess
}
var file_service_http_gzip_proto_depIdxs = []int32{
	5, // 0: pb.CreateHTTPGzipRequest.minLength:type_name -> pb.SizeCapacity
	5, // 1: pb.CreateHTTPGzipRequest.maxLength:type_name -> pb.SizeCapacity
	5, // 2: pb.UpdateHTTPGzipRequest.minLength:type_name -> pb.SizeCapacity
	5, // 3: pb.UpdateHTTPGzipRequest.maxLength:type_name -> pb.SizeCapacity
	0, // 4: pb.HTTPGzipService.createHTTPGzip:input_type -> pb.CreateHTTPGzipRequest
	2, // 5: pb.HTTPGzipService.findEnabledHTTPGzipConfig:input_type -> pb.FindEnabledGzipConfigRequest
	4, // 6: pb.HTTPGzipService.updateHTTPGzip:input_type -> pb.UpdateHTTPGzipRequest
	1, // 7: pb.HTTPGzipService.createHTTPGzip:output_type -> pb.CreateHTTPGzipResponse
	3, // 8: pb.HTTPGzipService.findEnabledHTTPGzipConfig:output_type -> pb.FindEnabledGzipConfigResponse
	6, // 9: pb.HTTPGzipService.updateHTTPGzip:output_type -> pb.RPCSuccess
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_http_gzip_proto_init() }
func file_service_http_gzip_proto_init() {
	if File_service_http_gzip_proto != nil {
		return
	}
	file_models_model_size_capacity_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_gzip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPGzipRequest); i {
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
		file_service_http_gzip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPGzipResponse); i {
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
		file_service_http_gzip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledGzipConfigRequest); i {
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
		file_service_http_gzip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledGzipConfigResponse); i {
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
		file_service_http_gzip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPGzipRequest); i {
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
			RawDescriptor: file_service_http_gzip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_gzip_proto_goTypes,
		DependencyIndexes: file_service_http_gzip_proto_depIdxs,
		MessageInfos:      file_service_http_gzip_proto_msgTypes,
	}.Build()
	File_service_http_gzip_proto = out.File
	file_service_http_gzip_proto_rawDesc = nil
	file_service_http_gzip_proto_goTypes = nil
	file_service_http_gzip_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPGzipServiceClient is the client API for HTTPGzipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPGzipServiceClient interface {
	// 创建Gzip配置
	CreateHTTPGzip(ctx context.Context, in *CreateHTTPGzipRequest, opts ...grpc.CallOption) (*CreateHTTPGzipResponse, error)
	// 查找Gzip配置
	FindEnabledHTTPGzipConfig(ctx context.Context, in *FindEnabledGzipConfigRequest, opts ...grpc.CallOption) (*FindEnabledGzipConfigResponse, error)
	// 修改Gzip配置
	UpdateHTTPGzip(ctx context.Context, in *UpdateHTTPGzipRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPGzipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPGzipServiceClient(cc grpc.ClientConnInterface) HTTPGzipServiceClient {
	return &hTTPGzipServiceClient{cc}
}

func (c *hTTPGzipServiceClient) CreateHTTPGzip(ctx context.Context, in *CreateHTTPGzipRequest, opts ...grpc.CallOption) (*CreateHTTPGzipResponse, error) {
	out := new(CreateHTTPGzipResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPGzipService/createHTTPGzip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPGzipServiceClient) FindEnabledHTTPGzipConfig(ctx context.Context, in *FindEnabledGzipConfigRequest, opts ...grpc.CallOption) (*FindEnabledGzipConfigResponse, error) {
	out := new(FindEnabledGzipConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPGzipService/findEnabledHTTPGzipConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPGzipServiceClient) UpdateHTTPGzip(ctx context.Context, in *UpdateHTTPGzipRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPGzipService/updateHTTPGzip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPGzipServiceServer is the server API for HTTPGzipService service.
type HTTPGzipServiceServer interface {
	// 创建Gzip配置
	CreateHTTPGzip(context.Context, *CreateHTTPGzipRequest) (*CreateHTTPGzipResponse, error)
	// 查找Gzip配置
	FindEnabledHTTPGzipConfig(context.Context, *FindEnabledGzipConfigRequest) (*FindEnabledGzipConfigResponse, error)
	// 修改Gzip配置
	UpdateHTTPGzip(context.Context, *UpdateHTTPGzipRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPGzipServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPGzipServiceServer struct {
}

func (*UnimplementedHTTPGzipServiceServer) CreateHTTPGzip(context.Context, *CreateHTTPGzipRequest) (*CreateHTTPGzipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPGzip not implemented")
}
func (*UnimplementedHTTPGzipServiceServer) FindEnabledHTTPGzipConfig(context.Context, *FindEnabledGzipConfigRequest) (*FindEnabledGzipConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPGzipConfig not implemented")
}
func (*UnimplementedHTTPGzipServiceServer) UpdateHTTPGzip(context.Context, *UpdateHTTPGzipRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPGzip not implemented")
}

func RegisterHTTPGzipServiceServer(s *grpc.Server, srv HTTPGzipServiceServer) {
	s.RegisterService(&_HTTPGzipService_serviceDesc, srv)
}

func _HTTPGzipService_CreateHTTPGzip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPGzipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPGzipServiceServer).CreateHTTPGzip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPGzipService/CreateHTTPGzip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPGzipServiceServer).CreateHTTPGzip(ctx, req.(*CreateHTTPGzipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPGzipService_FindEnabledHTTPGzipConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledGzipConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPGzipServiceServer).FindEnabledHTTPGzipConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPGzipService/FindEnabledHTTPGzipConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPGzipServiceServer).FindEnabledHTTPGzipConfig(ctx, req.(*FindEnabledGzipConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPGzipService_UpdateHTTPGzip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPGzipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPGzipServiceServer).UpdateHTTPGzip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPGzipService/UpdateHTTPGzip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPGzipServiceServer).UpdateHTTPGzip(ctx, req.(*UpdateHTTPGzipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPGzipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPGzipService",
	HandlerType: (*HTTPGzipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPGzip",
			Handler:    _HTTPGzipService_CreateHTTPGzip_Handler,
		},
		{
			MethodName: "findEnabledHTTPGzipConfig",
			Handler:    _HTTPGzipService_FindEnabledHTTPGzipConfig_Handler,
		},
		{
			MethodName: "updateHTTPGzip",
			Handler:    _HTTPGzipService_UpdateHTTPGzip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_gzip.proto",
}
