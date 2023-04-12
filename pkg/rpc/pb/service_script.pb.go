// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_script.proto

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

// 添加脚本
type CreateScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreateScriptRequest) Reset() {
	*x = CreateScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScriptRequest) ProtoMessage() {}

func (x *CreateScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScriptRequest.ProtoReflect.Descriptor instead.
func (*CreateScriptRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScriptRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateScriptRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateScriptRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CreateScriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptId int64 `protobuf:"varint,1,opt,name=scriptId,proto3" json:"scriptId,omitempty"`
}

func (x *CreateScriptResponse) Reset() {
	*x = CreateScriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScriptResponse) ProtoMessage() {}

func (x *CreateScriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScriptResponse.ProtoReflect.Descriptor instead.
func (*CreateScriptResponse) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScriptResponse) GetScriptId() int64 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

// 删除脚本
type DeleteScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptId int64 `protobuf:"varint,1,opt,name=scriptId,proto3" json:"scriptId,omitempty"`
}

func (x *DeleteScriptRequest) Reset() {
	*x = DeleteScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScriptRequest) ProtoMessage() {}

func (x *DeleteScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScriptRequest.ProtoReflect.Descriptor instead.
func (*DeleteScriptRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteScriptRequest) GetScriptId() int64 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

// 计算脚本数量
type CountAllEnabledScriptsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CountAllEnabledScriptsRequest) Reset() {
	*x = CountAllEnabledScriptsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountAllEnabledScriptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountAllEnabledScriptsRequest) ProtoMessage() {}

func (x *CountAllEnabledScriptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountAllEnabledScriptsRequest.ProtoReflect.Descriptor instead.
func (*CountAllEnabledScriptsRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{3}
}

func (x *CountAllEnabledScriptsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 列出单页脚本
type ListEnabledScriptsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size   int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListEnabledScriptsRequest) Reset() {
	*x = ListEnabledScriptsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledScriptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledScriptsRequest) ProtoMessage() {}

func (x *ListEnabledScriptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledScriptsRequest.ProtoReflect.Descriptor instead.
func (*ListEnabledScriptsRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{4}
}

func (x *ListEnabledScriptsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListEnabledScriptsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEnabledScriptsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListEnabledScriptsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scripts []*Script `protobuf:"bytes,1,rep,name=scripts,proto3" json:"scripts,omitempty"`
}

func (x *ListEnabledScriptsResponse) Reset() {
	*x = ListEnabledScriptsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnabledScriptsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnabledScriptsResponse) ProtoMessage() {}

func (x *ListEnabledScriptsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnabledScriptsResponse.ProtoReflect.Descriptor instead.
func (*ListEnabledScriptsResponse) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{5}
}

func (x *ListEnabledScriptsResponse) GetScripts() []*Script {
	if x != nil {
		return x.Scripts
	}
	return nil
}

// 发布脚本
type PublishScriptsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *PublishScriptsRequest) Reset() {
	*x = PublishScriptsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishScriptsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishScriptsRequest) ProtoMessage() {}

func (x *PublishScriptsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishScriptsRequest.ProtoReflect.Descriptor instead.
func (*PublishScriptsRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{6}
}

func (x *PublishScriptsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 检查脚本是否需要有更新
type CheckScriptUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CheckScriptUpdatesRequest) Reset() {
	*x = CheckScriptUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckScriptUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckScriptUpdatesRequest) ProtoMessage() {}

func (x *CheckScriptUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckScriptUpdatesRequest.ProtoReflect.Descriptor instead.
func (*CheckScriptUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{7}
}

func (x *CheckScriptUpdatesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckScriptUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasUpdates bool  `protobuf:"varint,1,opt,name=hasUpdates,proto3" json:"hasUpdates,omitempty"`
	Version    int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CheckScriptUpdatesResponse) Reset() {
	*x = CheckScriptUpdatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckScriptUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckScriptUpdatesResponse) ProtoMessage() {}

func (x *CheckScriptUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckScriptUpdatesResponse.ProtoReflect.Descriptor instead.
func (*CheckScriptUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{8}
}

func (x *CheckScriptUpdatesResponse) GetHasUpdates() bool {
	if x != nil {
		return x.HasUpdates
	}
	return false
}

func (x *CheckScriptUpdatesResponse) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

// 查找单个脚本
type FindEnabledScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptId int64 `protobuf:"varint,1,opt,name=scriptId,proto3" json:"scriptId,omitempty"`
}

func (x *FindEnabledScriptRequest) Reset() {
	*x = FindEnabledScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledScriptRequest) ProtoMessage() {}

func (x *FindEnabledScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledScriptRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledScriptRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{9}
}

func (x *FindEnabledScriptRequest) GetScriptId() int64 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

type FindEnabledScriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script *Script `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *FindEnabledScriptResponse) Reset() {
	*x = FindEnabledScriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledScriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledScriptResponse) ProtoMessage() {}

func (x *FindEnabledScriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledScriptResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledScriptResponse) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{10}
}

func (x *FindEnabledScriptResponse) GetScript() *Script {
	if x != nil {
		return x.Script
	}
	return nil
}

// 修改脚本
type UpdateScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptId int64  `protobuf:"varint,1,opt,name=scriptId,proto3" json:"scriptId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Code     string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	IsOn     bool   `protobuf:"varint,5,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateScriptRequest) Reset() {
	*x = UpdateScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScriptRequest) ProtoMessage() {}

func (x *UpdateScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScriptRequest.ProtoReflect.Descriptor instead.
func (*UpdateScriptRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateScriptRequest) GetScriptId() int64 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

func (x *UpdateScriptRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateScriptRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UpdateScriptRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateScriptRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 组合脚本配置
type ComposeScriptConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeScriptConfigsRequest) Reset() {
	*x = ComposeScriptConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeScriptConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeScriptConfigsRequest) ProtoMessage() {}

func (x *ComposeScriptConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeScriptConfigsRequest.ProtoReflect.Descriptor instead.
func (*ComposeScriptConfigsRequest) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{12}
}

type ComposeScriptConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptConfigsJSON []byte `protobuf:"bytes,1,opt,name=scriptConfigsJSON,proto3" json:"scriptConfigsJSON,omitempty"`
}

func (x *ComposeScriptConfigsResponse) Reset() {
	*x = ComposeScriptConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_script_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeScriptConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeScriptConfigsResponse) ProtoMessage() {}

func (x *ComposeScriptConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_script_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeScriptConfigsResponse.ProtoReflect.Descriptor instead.
func (*ComposeScriptConfigsResponse) Descriptor() ([]byte, []int) {
	return file_service_script_proto_rawDescGZIP(), []int{13}
}

func (x *ComposeScriptConfigsResponse) GetScriptConfigsJSON() []byte {
	if x != nil {
		return x.ScriptConfigsJSON
	}
	return nil
}

var File_service_script_proto protoreflect.FileDescriptor

var file_service_script_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x59, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x22,
	0x31, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x49, 0x64, 0x22, 0x37, 0x0a, 0x1d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x19, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x42, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73,
	0x22, 0x2f, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x33, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x36,
	0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52,
	0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x4c, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x32, 0xab, 0x05, 0x0a, 0x0d, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x51,
	0x0a, 0x16, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x53, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x59, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_script_proto_rawDescOnce sync.Once
	file_service_script_proto_rawDescData = file_service_script_proto_rawDesc
)

func file_service_script_proto_rawDescGZIP() []byte {
	file_service_script_proto_rawDescOnce.Do(func() {
		file_service_script_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_script_proto_rawDescData)
	})
	return file_service_script_proto_rawDescData
}

var file_service_script_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_service_script_proto_goTypes = []interface{}{
	(*CreateScriptRequest)(nil),           // 0: pb.CreateScriptRequest
	(*CreateScriptResponse)(nil),          // 1: pb.CreateScriptResponse
	(*DeleteScriptRequest)(nil),           // 2: pb.DeleteScriptRequest
	(*CountAllEnabledScriptsRequest)(nil), // 3: pb.CountAllEnabledScriptsRequest
	(*ListEnabledScriptsRequest)(nil),     // 4: pb.ListEnabledScriptsRequest
	(*ListEnabledScriptsResponse)(nil),    // 5: pb.ListEnabledScriptsResponse
	(*PublishScriptsRequest)(nil),         // 6: pb.PublishScriptsRequest
	(*CheckScriptUpdatesRequest)(nil),     // 7: pb.CheckScriptUpdatesRequest
	(*CheckScriptUpdatesResponse)(nil),    // 8: pb.CheckScriptUpdatesResponse
	(*FindEnabledScriptRequest)(nil),      // 9: pb.FindEnabledScriptRequest
	(*FindEnabledScriptResponse)(nil),     // 10: pb.FindEnabledScriptResponse
	(*UpdateScriptRequest)(nil),           // 11: pb.UpdateScriptRequest
	(*ComposeScriptConfigsRequest)(nil),   // 12: pb.ComposeScriptConfigsRequest
	(*ComposeScriptConfigsResponse)(nil),  // 13: pb.ComposeScriptConfigsResponse
	(*Script)(nil),                        // 14: pb.Script
	(*RPCSuccess)(nil),                    // 15: pb.RPCSuccess
	(*RPCCountResponse)(nil),              // 16: pb.RPCCountResponse
}
var file_service_script_proto_depIdxs = []int32{
	14, // 0: pb.ListEnabledScriptsResponse.scripts:type_name -> pb.Script
	14, // 1: pb.FindEnabledScriptResponse.script:type_name -> pb.Script
	0,  // 2: pb.ScriptService.createScript:input_type -> pb.CreateScriptRequest
	2,  // 3: pb.ScriptService.deleteScript:input_type -> pb.DeleteScriptRequest
	3,  // 4: pb.ScriptService.countAllEnabledScripts:input_type -> pb.CountAllEnabledScriptsRequest
	4,  // 5: pb.ScriptService.listEnabledScripts:input_type -> pb.ListEnabledScriptsRequest
	6,  // 6: pb.ScriptService.publishScripts:input_type -> pb.PublishScriptsRequest
	7,  // 7: pb.ScriptService.checkScriptUpdates:input_type -> pb.CheckScriptUpdatesRequest
	9,  // 8: pb.ScriptService.findEnabledScript:input_type -> pb.FindEnabledScriptRequest
	11, // 9: pb.ScriptService.updateScript:input_type -> pb.UpdateScriptRequest
	12, // 10: pb.ScriptService.composeScriptConfigs:input_type -> pb.ComposeScriptConfigsRequest
	1,  // 11: pb.ScriptService.createScript:output_type -> pb.CreateScriptResponse
	15, // 12: pb.ScriptService.deleteScript:output_type -> pb.RPCSuccess
	16, // 13: pb.ScriptService.countAllEnabledScripts:output_type -> pb.RPCCountResponse
	5,  // 14: pb.ScriptService.listEnabledScripts:output_type -> pb.ListEnabledScriptsResponse
	15, // 15: pb.ScriptService.publishScripts:output_type -> pb.RPCSuccess
	8,  // 16: pb.ScriptService.checkScriptUpdates:output_type -> pb.CheckScriptUpdatesResponse
	10, // 17: pb.ScriptService.findEnabledScript:output_type -> pb.FindEnabledScriptResponse
	15, // 18: pb.ScriptService.updateScript:output_type -> pb.RPCSuccess
	13, // 19: pb.ScriptService.composeScriptConfigs:output_type -> pb.ComposeScriptConfigsResponse
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_script_proto_init() }
func file_service_script_proto_init() {
	if File_service_script_proto != nil {
		return
	}
	file_models_model_script_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_script_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScriptRequest); i {
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
		file_service_script_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScriptResponse); i {
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
		file_service_script_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScriptRequest); i {
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
		file_service_script_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountAllEnabledScriptsRequest); i {
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
		file_service_script_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledScriptsRequest); i {
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
		file_service_script_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnabledScriptsResponse); i {
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
		file_service_script_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishScriptsRequest); i {
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
		file_service_script_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckScriptUpdatesRequest); i {
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
		file_service_script_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckScriptUpdatesResponse); i {
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
		file_service_script_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledScriptRequest); i {
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
		file_service_script_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledScriptResponse); i {
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
		file_service_script_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScriptRequest); i {
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
		file_service_script_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeScriptConfigsRequest); i {
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
		file_service_script_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeScriptConfigsResponse); i {
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
			RawDescriptor: file_service_script_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_script_proto_goTypes,
		DependencyIndexes: file_service_script_proto_depIdxs,
		MessageInfos:      file_service_script_proto_msgTypes,
	}.Build()
	File_service_script_proto = out.File
	file_service_script_proto_rawDesc = nil
	file_service_script_proto_goTypes = nil
	file_service_script_proto_depIdxs = nil
}
