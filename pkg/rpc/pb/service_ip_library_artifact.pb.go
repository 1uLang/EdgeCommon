// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_ip_library_artifact.proto

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

// 创建制品
type CreateIPLibraryArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId   int64  `protobuf:"varint,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
	MetaJSON []byte `protobuf:"bytes,2,opt,name=metaJSON,proto3" json:"metaJSON,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateIPLibraryArtifactRequest) Reset() {
	*x = CreateIPLibraryArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIPLibraryArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIPLibraryArtifactRequest) ProtoMessage() {}

func (x *CreateIPLibraryArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIPLibraryArtifactRequest.ProtoReflect.Descriptor instead.
func (*CreateIPLibraryArtifactRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIPLibraryArtifactRequest) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *CreateIPLibraryArtifactRequest) GetMetaJSON() []byte {
	if x != nil {
		return x.MetaJSON
	}
	return nil
}

func (x *CreateIPLibraryArtifactRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateIPLibraryArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifactId int64 `protobuf:"varint,1,opt,name=ipLibraryArtifactId,proto3" json:"ipLibraryArtifactId,omitempty"`
}

func (x *CreateIPLibraryArtifactResponse) Reset() {
	*x = CreateIPLibraryArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIPLibraryArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIPLibraryArtifactResponse) ProtoMessage() {}

func (x *CreateIPLibraryArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIPLibraryArtifactResponse.ProtoReflect.Descriptor instead.
func (*CreateIPLibraryArtifactResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIPLibraryArtifactResponse) GetIpLibraryArtifactId() int64 {
	if x != nil {
		return x.IpLibraryArtifactId
	}
	return 0
}

// 使用/取消使用制品
type UpdateIPLibraryArtifactIsPublicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifactId int64 `protobuf:"varint,1,opt,name=ipLibraryArtifactId,proto3" json:"ipLibraryArtifactId,omitempty"`
	IsPublic            bool  `protobuf:"varint,2,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
}

func (x *UpdateIPLibraryArtifactIsPublicRequest) Reset() {
	*x = UpdateIPLibraryArtifactIsPublicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIPLibraryArtifactIsPublicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIPLibraryArtifactIsPublicRequest) ProtoMessage() {}

func (x *UpdateIPLibraryArtifactIsPublicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIPLibraryArtifactIsPublicRequest.ProtoReflect.Descriptor instead.
func (*UpdateIPLibraryArtifactIsPublicRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateIPLibraryArtifactIsPublicRequest) GetIpLibraryArtifactId() int64 {
	if x != nil {
		return x.IpLibraryArtifactId
	}
	return 0
}

func (x *UpdateIPLibraryArtifactIsPublicRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

// 查询所有制品
type FindAllIPLibraryArtifactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllIPLibraryArtifactsRequest) Reset() {
	*x = FindAllIPLibraryArtifactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllIPLibraryArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllIPLibraryArtifactsRequest) ProtoMessage() {}

func (x *FindAllIPLibraryArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllIPLibraryArtifactsRequest.ProtoReflect.Descriptor instead.
func (*FindAllIPLibraryArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{3}
}

type FindAllIPLibraryArtifactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifacts []*IPLibraryArtifact `protobuf:"bytes,1,rep,name=ipLibraryArtifacts,proto3" json:"ipLibraryArtifacts,omitempty"`
}

func (x *FindAllIPLibraryArtifactsResponse) Reset() {
	*x = FindAllIPLibraryArtifactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllIPLibraryArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllIPLibraryArtifactsResponse) ProtoMessage() {}

func (x *FindAllIPLibraryArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllIPLibraryArtifactsResponse.ProtoReflect.Descriptor instead.
func (*FindAllIPLibraryArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllIPLibraryArtifactsResponse) GetIpLibraryArtifacts() []*IPLibraryArtifact {
	if x != nil {
		return x.IpLibraryArtifacts
	}
	return nil
}

// 查找单个制品信息
type FindIPLibraryArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifactId int64 `protobuf:"varint,1,opt,name=ipLibraryArtifactId,proto3" json:"ipLibraryArtifactId,omitempty"`
}

func (x *FindIPLibraryArtifactRequest) Reset() {
	*x = FindIPLibraryArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindIPLibraryArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindIPLibraryArtifactRequest) ProtoMessage() {}

func (x *FindIPLibraryArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindIPLibraryArtifactRequest.ProtoReflect.Descriptor instead.
func (*FindIPLibraryArtifactRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{5}
}

func (x *FindIPLibraryArtifactRequest) GetIpLibraryArtifactId() int64 {
	if x != nil {
		return x.IpLibraryArtifactId
	}
	return 0
}

type FindIPLibraryArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifact *IPLibraryArtifact `protobuf:"bytes,1,opt,name=ipLibraryArtifact,proto3" json:"ipLibraryArtifact,omitempty"`
}

func (x *FindIPLibraryArtifactResponse) Reset() {
	*x = FindIPLibraryArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindIPLibraryArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindIPLibraryArtifactResponse) ProtoMessage() {}

func (x *FindIPLibraryArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindIPLibraryArtifactResponse.ProtoReflect.Descriptor instead.
func (*FindIPLibraryArtifactResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{6}
}

func (x *FindIPLibraryArtifactResponse) GetIpLibraryArtifact() *IPLibraryArtifact {
	if x != nil {
		return x.IpLibraryArtifact
	}
	return nil
}

// 查找当前正在使用的制品
type FindPublicIPLibraryArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindPublicIPLibraryArtifactRequest) Reset() {
	*x = FindPublicIPLibraryArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPublicIPLibraryArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPublicIPLibraryArtifactRequest) ProtoMessage() {}

func (x *FindPublicIPLibraryArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPublicIPLibraryArtifactRequest.ProtoReflect.Descriptor instead.
func (*FindPublicIPLibraryArtifactRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{7}
}

type FindPublicIPLibraryArtifactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifact *IPLibraryArtifact `protobuf:"bytes,1,opt,name=ipLibraryArtifact,proto3" json:"ipLibraryArtifact,omitempty"`
}

func (x *FindPublicIPLibraryArtifactResponse) Reset() {
	*x = FindPublicIPLibraryArtifactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindPublicIPLibraryArtifactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindPublicIPLibraryArtifactResponse) ProtoMessage() {}

func (x *FindPublicIPLibraryArtifactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindPublicIPLibraryArtifactResponse.ProtoReflect.Descriptor instead.
func (*FindPublicIPLibraryArtifactResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{8}
}

func (x *FindPublicIPLibraryArtifactResponse) GetIpLibraryArtifact() *IPLibraryArtifact {
	if x != nil {
		return x.IpLibraryArtifact
	}
	return nil
}

// 删除制品
type DeleteIPLibraryArtifactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpLibraryArtifactId int64 `protobuf:"varint,1,opt,name=ipLibraryArtifactId,proto3" json:"ipLibraryArtifactId,omitempty"`
}

func (x *DeleteIPLibraryArtifactRequest) Reset() {
	*x = DeleteIPLibraryArtifactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_library_artifact_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIPLibraryArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIPLibraryArtifactRequest) ProtoMessage() {}

func (x *DeleteIPLibraryArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_library_artifact_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIPLibraryArtifactRequest.ProtoReflect.Descriptor instead.
func (*DeleteIPLibraryArtifactRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_library_artifact_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteIPLibraryArtifactRequest) GetIpLibraryArtifactId() int64 {
	if x != nil {
		return x.IpLibraryArtifactId
	}
	return 0
}

var File_service_ip_library_artifact_proto protoreflect.FileDescriptor

var file_service_ip_library_artifact_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x1e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x70, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x26, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x49, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x22, 0x22, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x50, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x69, 0x70,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x50, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x12, 0x69,
	0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x22, 0x50, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13,
	0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x50, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x11, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x11, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x22, 0x24, 0x0a, 0x22, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6a, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x11, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x11, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x22, 0x52, 0x0a, 0x1e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x13, 0x69, 0x70, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x69, 0x70, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x32,
	0xe4, 0x04, 0x0a, 0x18, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x17,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5d, 0x0a, 0x1f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x12, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x49, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x68, 0x0a, 0x19, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x49,
	0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x66, 0x69, 0x6e,
	0x64, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x50, 0x4c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x50,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x1b, 0x66, 0x69, 0x6e, 0x64, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x50,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x50, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ip_library_artifact_proto_rawDescOnce sync.Once
	file_service_ip_library_artifact_proto_rawDescData = file_service_ip_library_artifact_proto_rawDesc
)

func file_service_ip_library_artifact_proto_rawDescGZIP() []byte {
	file_service_ip_library_artifact_proto_rawDescOnce.Do(func() {
		file_service_ip_library_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ip_library_artifact_proto_rawDescData)
	})
	return file_service_ip_library_artifact_proto_rawDescData
}

var file_service_ip_library_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_ip_library_artifact_proto_goTypes = []interface{}{
	(*CreateIPLibraryArtifactRequest)(nil),         // 0: pb.CreateIPLibraryArtifactRequest
	(*CreateIPLibraryArtifactResponse)(nil),        // 1: pb.CreateIPLibraryArtifactResponse
	(*UpdateIPLibraryArtifactIsPublicRequest)(nil), // 2: pb.UpdateIPLibraryArtifactIsPublicRequest
	(*FindAllIPLibraryArtifactsRequest)(nil),       // 3: pb.FindAllIPLibraryArtifactsRequest
	(*FindAllIPLibraryArtifactsResponse)(nil),      // 4: pb.FindAllIPLibraryArtifactsResponse
	(*FindIPLibraryArtifactRequest)(nil),           // 5: pb.FindIPLibraryArtifactRequest
	(*FindIPLibraryArtifactResponse)(nil),          // 6: pb.FindIPLibraryArtifactResponse
	(*FindPublicIPLibraryArtifactRequest)(nil),     // 7: pb.FindPublicIPLibraryArtifactRequest
	(*FindPublicIPLibraryArtifactResponse)(nil),    // 8: pb.FindPublicIPLibraryArtifactResponse
	(*DeleteIPLibraryArtifactRequest)(nil),         // 9: pb.DeleteIPLibraryArtifactRequest
	(*IPLibraryArtifact)(nil),                      // 10: pb.IPLibraryArtifact
	(*RPCSuccess)(nil),                             // 11: pb.RPCSuccess
}
var file_service_ip_library_artifact_proto_depIdxs = []int32{
	10, // 0: pb.FindAllIPLibraryArtifactsResponse.ipLibraryArtifacts:type_name -> pb.IPLibraryArtifact
	10, // 1: pb.FindIPLibraryArtifactResponse.ipLibraryArtifact:type_name -> pb.IPLibraryArtifact
	10, // 2: pb.FindPublicIPLibraryArtifactResponse.ipLibraryArtifact:type_name -> pb.IPLibraryArtifact
	0,  // 3: pb.IPLibraryArtifactService.createIPLibraryArtifact:input_type -> pb.CreateIPLibraryArtifactRequest
	2,  // 4: pb.IPLibraryArtifactService.updateIPLibraryArtifactIsPublic:input_type -> pb.UpdateIPLibraryArtifactIsPublicRequest
	3,  // 5: pb.IPLibraryArtifactService.findAllIPLibraryArtifacts:input_type -> pb.FindAllIPLibraryArtifactsRequest
	5,  // 6: pb.IPLibraryArtifactService.findIPLibraryArtifact:input_type -> pb.FindIPLibraryArtifactRequest
	7,  // 7: pb.IPLibraryArtifactService.findPublicIPLibraryArtifact:input_type -> pb.FindPublicIPLibraryArtifactRequest
	9,  // 8: pb.IPLibraryArtifactService.deleteIPLibraryArtifact:input_type -> pb.DeleteIPLibraryArtifactRequest
	1,  // 9: pb.IPLibraryArtifactService.createIPLibraryArtifact:output_type -> pb.CreateIPLibraryArtifactResponse
	11, // 10: pb.IPLibraryArtifactService.updateIPLibraryArtifactIsPublic:output_type -> pb.RPCSuccess
	4,  // 11: pb.IPLibraryArtifactService.findAllIPLibraryArtifacts:output_type -> pb.FindAllIPLibraryArtifactsResponse
	6,  // 12: pb.IPLibraryArtifactService.findIPLibraryArtifact:output_type -> pb.FindIPLibraryArtifactResponse
	8,  // 13: pb.IPLibraryArtifactService.findPublicIPLibraryArtifact:output_type -> pb.FindPublicIPLibraryArtifactResponse
	11, // 14: pb.IPLibraryArtifactService.deleteIPLibraryArtifact:output_type -> pb.RPCSuccess
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_ip_library_artifact_proto_init() }
func file_service_ip_library_artifact_proto_init() {
	if File_service_ip_library_artifact_proto != nil {
		return
	}
	file_models_model_ip_library_artifact_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ip_library_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIPLibraryArtifactRequest); i {
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
		file_service_ip_library_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIPLibraryArtifactResponse); i {
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
		file_service_ip_library_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIPLibraryArtifactIsPublicRequest); i {
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
		file_service_ip_library_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllIPLibraryArtifactsRequest); i {
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
		file_service_ip_library_artifact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllIPLibraryArtifactsResponse); i {
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
		file_service_ip_library_artifact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindIPLibraryArtifactRequest); i {
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
		file_service_ip_library_artifact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindIPLibraryArtifactResponse); i {
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
		file_service_ip_library_artifact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPublicIPLibraryArtifactRequest); i {
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
		file_service_ip_library_artifact_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindPublicIPLibraryArtifactResponse); i {
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
		file_service_ip_library_artifact_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIPLibraryArtifactRequest); i {
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
			RawDescriptor: file_service_ip_library_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ip_library_artifact_proto_goTypes,
		DependencyIndexes: file_service_ip_library_artifact_proto_depIdxs,
		MessageInfos:      file_service_ip_library_artifact_proto_msgTypes,
	}.Build()
	File_service_ip_library_artifact_proto = out.File
	file_service_ip_library_artifact_proto_rawDesc = nil
	file_service_ip_library_artifact_proto_goTypes = nil
	file_service_ip_library_artifact_proto_depIdxs = nil
}