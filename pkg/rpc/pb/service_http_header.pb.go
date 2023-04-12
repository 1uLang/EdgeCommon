// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_http_header.proto

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

// 创建Header
type CreateHTTPHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value             string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Status            []int32  `protobuf:"varint,3,rep,packed,name=status,proto3" json:"status,omitempty"`
	DisableRedirect   bool     `protobuf:"varint,4,opt,name=disableRedirect,proto3" json:"disableRedirect,omitempty"`
	ShouldAppend      bool     `protobuf:"varint,5,opt,name=shouldAppend,proto3" json:"shouldAppend,omitempty"`
	ShouldReplace     bool     `protobuf:"varint,6,opt,name=shouldReplace,proto3" json:"shouldReplace,omitempty"`
	ReplaceValuesJSON []byte   `protobuf:"bytes,7,opt,name=replaceValuesJSON,proto3" json:"replaceValuesJSON,omitempty"`
	Methods           []string `protobuf:"bytes,8,rep,name=methods,proto3" json:"methods,omitempty"`
	Domains           []string `protobuf:"bytes,9,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *CreateHTTPHeaderRequest) Reset() {
	*x = CreateHTTPHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPHeaderRequest) ProtoMessage() {}

func (x *CreateHTTPHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPHeaderRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPHeaderRequest) Descriptor() ([]byte, []int) {
	return file_service_http_header_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHTTPHeaderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHTTPHeaderRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CreateHTTPHeaderRequest) GetStatus() []int32 {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreateHTTPHeaderRequest) GetDisableRedirect() bool {
	if x != nil {
		return x.DisableRedirect
	}
	return false
}

func (x *CreateHTTPHeaderRequest) GetShouldAppend() bool {
	if x != nil {
		return x.ShouldAppend
	}
	return false
}

func (x *CreateHTTPHeaderRequest) GetShouldReplace() bool {
	if x != nil {
		return x.ShouldReplace
	}
	return false
}

func (x *CreateHTTPHeaderRequest) GetReplaceValuesJSON() []byte {
	if x != nil {
		return x.ReplaceValuesJSON
	}
	return nil
}

func (x *CreateHTTPHeaderRequest) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *CreateHTTPHeaderRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

type CreateHTTPHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderId int64 `protobuf:"varint,1,opt,name=headerId,proto3" json:"headerId,omitempty"`
}

func (x *CreateHTTPHeaderResponse) Reset() {
	*x = CreateHTTPHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_header_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPHeaderResponse) ProtoMessage() {}

func (x *CreateHTTPHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_header_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPHeaderResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPHeaderResponse) Descriptor() ([]byte, []int) {
	return file_service_http_header_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPHeaderResponse) GetHeaderId() int64 {
	if x != nil {
		return x.HeaderId
	}
	return 0
}

// 修改Header
type UpdateHTTPHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderId          int64    `protobuf:"varint,1,opt,name=headerId,proto3" json:"headerId,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value             string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Status            []int32  `protobuf:"varint,4,rep,packed,name=status,proto3" json:"status,omitempty"`
	DisableRedirect   bool     `protobuf:"varint,5,opt,name=disableRedirect,proto3" json:"disableRedirect,omitempty"`
	ShouldAppend      bool     `protobuf:"varint,6,opt,name=shouldAppend,proto3" json:"shouldAppend,omitempty"`
	ShouldReplace     bool     `protobuf:"varint,7,opt,name=shouldReplace,proto3" json:"shouldReplace,omitempty"`
	ReplaceValuesJSON []byte   `protobuf:"bytes,8,opt,name=replaceValuesJSON,proto3" json:"replaceValuesJSON,omitempty"`
	Methods           []string `protobuf:"bytes,9,rep,name=methods,proto3" json:"methods,omitempty"`
	Domains           []string `protobuf:"bytes,10,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *UpdateHTTPHeaderRequest) Reset() {
	*x = UpdateHTTPHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_header_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPHeaderRequest) ProtoMessage() {}

func (x *UpdateHTTPHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_header_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPHeaderRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPHeaderRequest) Descriptor() ([]byte, []int) {
	return file_service_http_header_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPHeaderRequest) GetHeaderId() int64 {
	if x != nil {
		return x.HeaderId
	}
	return 0
}

func (x *UpdateHTTPHeaderRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHTTPHeaderRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UpdateHTTPHeaderRequest) GetStatus() []int32 {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *UpdateHTTPHeaderRequest) GetDisableRedirect() bool {
	if x != nil {
		return x.DisableRedirect
	}
	return false
}

func (x *UpdateHTTPHeaderRequest) GetShouldAppend() bool {
	if x != nil {
		return x.ShouldAppend
	}
	return false
}

func (x *UpdateHTTPHeaderRequest) GetShouldReplace() bool {
	if x != nil {
		return x.ShouldReplace
	}
	return false
}

func (x *UpdateHTTPHeaderRequest) GetReplaceValuesJSON() []byte {
	if x != nil {
		return x.ReplaceValuesJSON
	}
	return nil
}

func (x *UpdateHTTPHeaderRequest) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *UpdateHTTPHeaderRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

// 查找配置
type FindEnabledHTTPHeaderConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderId int64 `protobuf:"varint,1,opt,name=headerId,proto3" json:"headerId,omitempty"`
}

func (x *FindEnabledHTTPHeaderConfigRequest) Reset() {
	*x = FindEnabledHTTPHeaderConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_header_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPHeaderConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPHeaderConfigRequest) ProtoMessage() {}

func (x *FindEnabledHTTPHeaderConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_header_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPHeaderConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPHeaderConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_header_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledHTTPHeaderConfigRequest) GetHeaderId() int64 {
	if x != nil {
		return x.HeaderId
	}
	return 0
}

type FindEnabledHTTPHeaderConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderJSON []byte `protobuf:"bytes,1,opt,name=headerJSON,proto3" json:"headerJSON,omitempty"`
}

func (x *FindEnabledHTTPHeaderConfigResponse) Reset() {
	*x = FindEnabledHTTPHeaderConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_header_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPHeaderConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPHeaderConfigResponse) ProtoMessage() {}

func (x *FindEnabledHTTPHeaderConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_header_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPHeaderConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPHeaderConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_header_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPHeaderConfigResponse) GetHeaderJSON() []byte {
	if x != nil {
		return x.HeaderJSON
	}
	return nil
}

var File_service_http_header_proto protoreflect.FileDescriptor

var file_service_http_header_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x36,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xcd, 0x02, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x11, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53,
	0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4a, 0x53, 0x4f, 0x4e, 0x32,
	0x93, 0x02, 0x0a, 0x11, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6e, 0x0a, 0x1b, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54,
	0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_header_proto_rawDescOnce sync.Once
	file_service_http_header_proto_rawDescData = file_service_http_header_proto_rawDesc
)

func file_service_http_header_proto_rawDescGZIP() []byte {
	file_service_http_header_proto_rawDescOnce.Do(func() {
		file_service_http_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_header_proto_rawDescData)
	})
	return file_service_http_header_proto_rawDescData
}

var file_service_http_header_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_http_header_proto_goTypes = []interface{}{
	(*CreateHTTPHeaderRequest)(nil),             // 0: pb.CreateHTTPHeaderRequest
	(*CreateHTTPHeaderResponse)(nil),            // 1: pb.CreateHTTPHeaderResponse
	(*UpdateHTTPHeaderRequest)(nil),             // 2: pb.UpdateHTTPHeaderRequest
	(*FindEnabledHTTPHeaderConfigRequest)(nil),  // 3: pb.FindEnabledHTTPHeaderConfigRequest
	(*FindEnabledHTTPHeaderConfigResponse)(nil), // 4: pb.FindEnabledHTTPHeaderConfigResponse
	(*RPCSuccess)(nil),                          // 5: pb.RPCSuccess
}
var file_service_http_header_proto_depIdxs = []int32{
	0, // 0: pb.HTTPHeaderService.createHTTPHeader:input_type -> pb.CreateHTTPHeaderRequest
	2, // 1: pb.HTTPHeaderService.updateHTTPHeader:input_type -> pb.UpdateHTTPHeaderRequest
	3, // 2: pb.HTTPHeaderService.findEnabledHTTPHeaderConfig:input_type -> pb.FindEnabledHTTPHeaderConfigRequest
	1, // 3: pb.HTTPHeaderService.createHTTPHeader:output_type -> pb.CreateHTTPHeaderResponse
	5, // 4: pb.HTTPHeaderService.updateHTTPHeader:output_type -> pb.RPCSuccess
	4, // 5: pb.HTTPHeaderService.findEnabledHTTPHeaderConfig:output_type -> pb.FindEnabledHTTPHeaderConfigResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_http_header_proto_init() }
func file_service_http_header_proto_init() {
	if File_service_http_header_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPHeaderRequest); i {
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
		file_service_http_header_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPHeaderResponse); i {
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
		file_service_http_header_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPHeaderRequest); i {
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
		file_service_http_header_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPHeaderConfigRequest); i {
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
		file_service_http_header_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPHeaderConfigResponse); i {
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
			RawDescriptor: file_service_http_header_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_header_proto_goTypes,
		DependencyIndexes: file_service_http_header_proto_depIdxs,
		MessageInfos:      file_service_http_header_proto_msgTypes,
	}.Build()
	File_service_http_header_proto = out.File
	file_service_http_header_proto_rawDesc = nil
	file_service_http_header_proto_goTypes = nil
	file_service_http_header_proto_depIdxs = nil
}
