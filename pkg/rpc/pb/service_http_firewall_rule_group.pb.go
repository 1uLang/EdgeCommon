// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_http_firewall_rule_group.proto

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

// 设置是否启用分组
type UpdateHTTPFirewallRuleGroupIsOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId int64 `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
	IsOn                bool  `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateHTTPFirewallRuleGroupIsOnRequest) Reset() {
	*x = UpdateHTTPFirewallRuleGroupIsOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFirewallRuleGroupIsOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFirewallRuleGroupIsOnRequest) ProtoMessage() {}

func (x *UpdateHTTPFirewallRuleGroupIsOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFirewallRuleGroupIsOnRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFirewallRuleGroupIsOnRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateHTTPFirewallRuleGroupIsOnRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *UpdateHTTPFirewallRuleGroupIsOnRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 创建分组
type CreateHTTPFirewallRuleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOn        bool   `protobuf:"varint,1,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateHTTPFirewallRuleGroupRequest) Reset() {
	*x = CreateHTTPFirewallRuleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPFirewallRuleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPFirewallRuleGroupRequest) ProtoMessage() {}

func (x *CreateHTTPFirewallRuleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPFirewallRuleGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateHTTPFirewallRuleGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHTTPFirewallRuleGroupRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *CreateHTTPFirewallRuleGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHTTPFirewallRuleGroupRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateHTTPFirewallRuleGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateHTTPFirewallRuleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId int64 `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
}

func (x *CreateHTTPFirewallRuleGroupResponse) Reset() {
	*x = CreateHTTPFirewallRuleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHTTPFirewallRuleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHTTPFirewallRuleGroupResponse) ProtoMessage() {}

func (x *CreateHTTPFirewallRuleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHTTPFirewallRuleGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateHTTPFirewallRuleGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHTTPFirewallRuleGroupResponse) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

// 修改分组
type UpdateHTTPFirewallRuleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId int64  `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
	IsOn                bool   `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name                string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description         string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Code                string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UpdateHTTPFirewallRuleGroupRequest) Reset() {
	*x = UpdateHTTPFirewallRuleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFirewallRuleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFirewallRuleGroupRequest) ProtoMessage() {}

func (x *UpdateHTTPFirewallRuleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFirewallRuleGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFirewallRuleGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateHTTPFirewallRuleGroupRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *UpdateHTTPFirewallRuleGroupRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UpdateHTTPFirewallRuleGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHTTPFirewallRuleGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateHTTPFirewallRuleGroupRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// 获取分组配置
type FindEnabledHTTPFirewallRuleGroupConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId int64 `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigRequest) Reset() {
	*x = FindEnabledHTTPFirewallRuleGroupConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleGroupConfigRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleGroupConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleGroupConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleGroupConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

type FindEnabledHTTPFirewallRuleGroupConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupJSON []byte `protobuf:"bytes,1,opt,name=firewallRuleGroupJSON,proto3" json:"firewallRuleGroupJSON,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigResponse) Reset() {
	*x = FindEnabledHTTPFirewallRuleGroupConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleGroupConfigResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleGroupConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleGroupConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleGroupConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledHTTPFirewallRuleGroupConfigResponse) GetFirewallRuleGroupJSON() []byte {
	if x != nil {
		return x.FirewallRuleGroupJSON
	}
	return nil
}

// 获取分组信息
type FindEnabledHTTPFirewallRuleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId int64 `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleGroupRequest) Reset() {
	*x = FindEnabledHTTPFirewallRuleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleGroupRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleGroupRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledHTTPFirewallRuleGroupRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

type FindEnabledHTTPFirewallRuleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroup *HTTPFirewallRuleGroup `protobuf:"bytes,1,opt,name=firewallRuleGroup,proto3" json:"firewallRuleGroup,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleGroupResponse) Reset() {
	*x = FindEnabledHTTPFirewallRuleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleGroupResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleGroupResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledHTTPFirewallRuleGroupResponse) GetFirewallRuleGroup() *HTTPFirewallRuleGroup {
	if x != nil {
		return x.FirewallRuleGroup
	}
	return nil
}

// 修改分组的规则集
type UpdateHTTPFirewallRuleGroupSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId  int64  `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
	FirewallRuleSetsJSON []byte `protobuf:"bytes,2,opt,name=firewallRuleSetsJSON,proto3" json:"firewallRuleSetsJSON,omitempty"`
}

func (x *UpdateHTTPFirewallRuleGroupSetsRequest) Reset() {
	*x = UpdateHTTPFirewallRuleGroupSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFirewallRuleGroupSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFirewallRuleGroupSetsRequest) ProtoMessage() {}

func (x *UpdateHTTPFirewallRuleGroupSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFirewallRuleGroupSetsRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFirewallRuleGroupSetsRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateHTTPFirewallRuleGroupSetsRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *UpdateHTTPFirewallRuleGroupSetsRequest) GetFirewallRuleSetsJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetsJSON
	}
	return nil
}

// 添加规则集
type AddHTTPFirewallRuleGroupSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleGroupId       int64  `protobuf:"varint,1,opt,name=firewallRuleGroupId,proto3" json:"firewallRuleGroupId,omitempty"`
	FirewallRuleSetConfigJSON []byte `protobuf:"bytes,2,opt,name=firewallRuleSetConfigJSON,proto3" json:"firewallRuleSetConfigJSON,omitempty"`
}

func (x *AddHTTPFirewallRuleGroupSetRequest) Reset() {
	*x = AddHTTPFirewallRuleGroupSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHTTPFirewallRuleGroupSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHTTPFirewallRuleGroupSetRequest) ProtoMessage() {}

func (x *AddHTTPFirewallRuleGroupSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHTTPFirewallRuleGroupSetRequest.ProtoReflect.Descriptor instead.
func (*AddHTTPFirewallRuleGroupSetRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_group_proto_rawDescGZIP(), []int{9}
}

func (x *AddHTTPFirewallRuleGroupSetRequest) GetFirewallRuleGroupId() int64 {
	if x != nil {
		return x.FirewallRuleGroupId
	}
	return 0
}

func (x *AddHTTPFirewallRuleGroupSetRequest) GetFirewallRuleSetConfigJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetConfigJSON
	}
	return nil
}

var File_service_http_firewall_rule_group_proto protoreflect.FileDescriptor

var file_service_http_firewall_rule_group_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x23, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x22, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x61, 0x0a, 0x2d, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x15, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x5b, 0x0a, 0x27, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x73, 0x0a, 0x28, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x8e, 0x01, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13,
	0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x14, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x94, 0x01, 0x0a, 0x22, 0x41, 0x64, 0x64, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x3c, 0x0a, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f, 0x4e, 0x32, 0x8b,
	0x06, 0x0a, 0x1c, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5d, 0x0a, 0x1f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x73,
	0x4f, 0x6e, 0x12, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x6e,
	0x0a, 0x1b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x26, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x20, 0x66, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2b, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x1f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x55, 0x0a, 0x1b, 0x61, 0x64, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_firewall_rule_group_proto_rawDescOnce sync.Once
	file_service_http_firewall_rule_group_proto_rawDescData = file_service_http_firewall_rule_group_proto_rawDesc
)

func file_service_http_firewall_rule_group_proto_rawDescGZIP() []byte {
	file_service_http_firewall_rule_group_proto_rawDescOnce.Do(func() {
		file_service_http_firewall_rule_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_firewall_rule_group_proto_rawDescData)
	})
	return file_service_http_firewall_rule_group_proto_rawDescData
}

var file_service_http_firewall_rule_group_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_http_firewall_rule_group_proto_goTypes = []interface{}{
	(*UpdateHTTPFirewallRuleGroupIsOnRequest)(nil),         // 0: pb.UpdateHTTPFirewallRuleGroupIsOnRequest
	(*CreateHTTPFirewallRuleGroupRequest)(nil),             // 1: pb.CreateHTTPFirewallRuleGroupRequest
	(*CreateHTTPFirewallRuleGroupResponse)(nil),            // 2: pb.CreateHTTPFirewallRuleGroupResponse
	(*UpdateHTTPFirewallRuleGroupRequest)(nil),             // 3: pb.UpdateHTTPFirewallRuleGroupRequest
	(*FindEnabledHTTPFirewallRuleGroupConfigRequest)(nil),  // 4: pb.FindEnabledHTTPFirewallRuleGroupConfigRequest
	(*FindEnabledHTTPFirewallRuleGroupConfigResponse)(nil), // 5: pb.FindEnabledHTTPFirewallRuleGroupConfigResponse
	(*FindEnabledHTTPFirewallRuleGroupRequest)(nil),        // 6: pb.FindEnabledHTTPFirewallRuleGroupRequest
	(*FindEnabledHTTPFirewallRuleGroupResponse)(nil),       // 7: pb.FindEnabledHTTPFirewallRuleGroupResponse
	(*UpdateHTTPFirewallRuleGroupSetsRequest)(nil),         // 8: pb.UpdateHTTPFirewallRuleGroupSetsRequest
	(*AddHTTPFirewallRuleGroupSetRequest)(nil),             // 9: pb.AddHTTPFirewallRuleGroupSetRequest
	(*HTTPFirewallRuleGroup)(nil),                          // 10: pb.HTTPFirewallRuleGroup
	(*RPCSuccess)(nil),                                     // 11: pb.RPCSuccess
}
var file_service_http_firewall_rule_group_proto_depIdxs = []int32{
	10, // 0: pb.FindEnabledHTTPFirewallRuleGroupResponse.firewallRuleGroup:type_name -> pb.HTTPFirewallRuleGroup
	0,  // 1: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroupIsOn:input_type -> pb.UpdateHTTPFirewallRuleGroupIsOnRequest
	1,  // 2: pb.HTTPFirewallRuleGroupService.createHTTPFirewallRuleGroup:input_type -> pb.CreateHTTPFirewallRuleGroupRequest
	3,  // 3: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroup:input_type -> pb.UpdateHTTPFirewallRuleGroupRequest
	4,  // 4: pb.HTTPFirewallRuleGroupService.findEnabledHTTPFirewallRuleGroupConfig:input_type -> pb.FindEnabledHTTPFirewallRuleGroupConfigRequest
	6,  // 5: pb.HTTPFirewallRuleGroupService.findEnabledHTTPFirewallRuleGroup:input_type -> pb.FindEnabledHTTPFirewallRuleGroupRequest
	8,  // 6: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroupSets:input_type -> pb.UpdateHTTPFirewallRuleGroupSetsRequest
	9,  // 7: pb.HTTPFirewallRuleGroupService.addHTTPFirewallRuleGroupSet:input_type -> pb.AddHTTPFirewallRuleGroupSetRequest
	11, // 8: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroupIsOn:output_type -> pb.RPCSuccess
	2,  // 9: pb.HTTPFirewallRuleGroupService.createHTTPFirewallRuleGroup:output_type -> pb.CreateHTTPFirewallRuleGroupResponse
	11, // 10: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroup:output_type -> pb.RPCSuccess
	5,  // 11: pb.HTTPFirewallRuleGroupService.findEnabledHTTPFirewallRuleGroupConfig:output_type -> pb.FindEnabledHTTPFirewallRuleGroupConfigResponse
	7,  // 12: pb.HTTPFirewallRuleGroupService.findEnabledHTTPFirewallRuleGroup:output_type -> pb.FindEnabledHTTPFirewallRuleGroupResponse
	11, // 13: pb.HTTPFirewallRuleGroupService.updateHTTPFirewallRuleGroupSets:output_type -> pb.RPCSuccess
	11, // 14: pb.HTTPFirewallRuleGroupService.addHTTPFirewallRuleGroupSet:output_type -> pb.RPCSuccess
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_http_firewall_rule_group_proto_init() }
func file_service_http_firewall_rule_group_proto_init() {
	if File_service_http_firewall_rule_group_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_http_firewall_rule_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_firewall_rule_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFirewallRuleGroupIsOnRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPFirewallRuleGroupRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHTTPFirewallRuleGroupResponse); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFirewallRuleGroupRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleGroupConfigRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleGroupConfigResponse); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleGroupRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleGroupResponse); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFirewallRuleGroupSetsRequest); i {
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
		file_service_http_firewall_rule_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHTTPFirewallRuleGroupSetRequest); i {
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
			RawDescriptor: file_service_http_firewall_rule_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_firewall_rule_group_proto_goTypes,
		DependencyIndexes: file_service_http_firewall_rule_group_proto_depIdxs,
		MessageInfos:      file_service_http_firewall_rule_group_proto_msgTypes,
	}.Build()
	File_service_http_firewall_rule_group_proto = out.File
	file_service_http_firewall_rule_group_proto_rawDesc = nil
	file_service_http_firewall_rule_group_proto_goTypes = nil
	file_service_http_firewall_rule_group_proto_depIdxs = nil
}
