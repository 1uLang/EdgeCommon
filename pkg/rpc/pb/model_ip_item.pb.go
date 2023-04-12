// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_ip_item.proto

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

type IPItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IpFrom                        string                 `protobuf:"bytes,2,opt,name=ipFrom,proto3" json:"ipFrom,omitempty"`
	IpTo                          string                 `protobuf:"bytes,3,opt,name=ipTo,proto3" json:"ipTo,omitempty"`
	Version                       int64                  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExpiredAt                     int64                  `protobuf:"varint,5,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Reason                        string                 `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason,omitempty"`
	ListId                        int64                  `protobuf:"varint,7,opt,name=listId,proto3" json:"listId,omitempty"`
	IsDeleted                     bool                   `protobuf:"varint,8,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Type                          string                 `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	EventLevel                    string                 `protobuf:"bytes,10,opt,name=eventLevel,proto3" json:"eventLevel,omitempty"` // 级别
	ListType                      string                 `protobuf:"bytes,11,opt,name=listType,proto3" json:"listType,omitempty"`     // 所在名单类型，来自名单
	IsGlobal                      bool                   `protobuf:"varint,20,opt,name=isGlobal,proto3" json:"isGlobal,omitempty"`    // 是否全局，来自名单
	CreatedAt                     int64                  `protobuf:"varint,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	NodeId                        int64                  `protobuf:"varint,13,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	ServerId                      int64                  `protobuf:"varint,14,opt,name=serverId,proto3" json:"serverId,omitempty"`
	SourceNodeId                  int64                  `protobuf:"varint,15,opt,name=sourceNodeId,proto3" json:"sourceNodeId,omitempty"`
	SourceServerId                int64                  `protobuf:"varint,16,opt,name=sourceServerId,proto3" json:"sourceServerId,omitempty"`
	SourceHTTPFirewallPolicyId    int64                  `protobuf:"varint,17,opt,name=sourceHTTPFirewallPolicyId,proto3" json:"sourceHTTPFirewallPolicyId,omitempty"`
	SourceHTTPFirewallRuleGroupId int64                  `protobuf:"varint,18,opt,name=sourceHTTPFirewallRuleGroupId,proto3" json:"sourceHTTPFirewallRuleGroupId,omitempty"`
	SourceHTTPFirewallRuleSetId   int64                  `protobuf:"varint,19,opt,name=sourceHTTPFirewallRuleSetId,proto3" json:"sourceHTTPFirewallRuleSetId,omitempty"`
	IsRead                        bool                   `protobuf:"varint,21,opt,name=isRead,proto3" json:"isRead,omitempty"`
	SourceServer                  *Server                `protobuf:"bytes,30,opt,name=sourceServer,proto3" json:"sourceServer,omitempty"`
	Server                        *Server                `protobuf:"bytes,34,opt,name=server,proto3" json:"server,omitempty"`
	SourceHTTPFirewallPolicy      *HTTPFirewallPolicy    `protobuf:"bytes,31,opt,name=sourceHTTPFirewallPolicy,proto3" json:"sourceHTTPFirewallPolicy,omitempty"`
	SourceHTTPFirewallRuleGroup   *HTTPFirewallRuleGroup `protobuf:"bytes,32,opt,name=sourceHTTPFirewallRuleGroup,proto3" json:"sourceHTTPFirewallRuleGroup,omitempty"`
	SourceHTTPFirewallRuleSet     *HTTPFirewallRuleSet   `protobuf:"bytes,33,opt,name=sourceHTTPFirewallRuleSet,proto3" json:"sourceHTTPFirewallRuleSet,omitempty"`
	SourceNode                    *Node                  `protobuf:"bytes,35,opt,name=sourceNode,proto3" json:"sourceNode,omitempty"`
}

func (x *IPItem) Reset() {
	*x = IPItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ip_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPItem) ProtoMessage() {}

func (x *IPItem) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ip_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPItem.ProtoReflect.Descriptor instead.
func (*IPItem) Descriptor() ([]byte, []int) {
	return file_models_model_ip_item_proto_rawDescGZIP(), []int{0}
}

func (x *IPItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IPItem) GetIpFrom() string {
	if x != nil {
		return x.IpFrom
	}
	return ""
}

func (x *IPItem) GetIpTo() string {
	if x != nil {
		return x.IpTo
	}
	return ""
}

func (x *IPItem) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *IPItem) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *IPItem) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *IPItem) GetListId() int64 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *IPItem) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *IPItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IPItem) GetEventLevel() string {
	if x != nil {
		return x.EventLevel
	}
	return ""
}

func (x *IPItem) GetListType() string {
	if x != nil {
		return x.ListType
	}
	return ""
}

func (x *IPItem) GetIsGlobal() bool {
	if x != nil {
		return x.IsGlobal
	}
	return false
}

func (x *IPItem) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *IPItem) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *IPItem) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *IPItem) GetSourceNodeId() int64 {
	if x != nil {
		return x.SourceNodeId
	}
	return 0
}

func (x *IPItem) GetSourceServerId() int64 {
	if x != nil {
		return x.SourceServerId
	}
	return 0
}

func (x *IPItem) GetSourceHTTPFirewallPolicyId() int64 {
	if x != nil {
		return x.SourceHTTPFirewallPolicyId
	}
	return 0
}

func (x *IPItem) GetSourceHTTPFirewallRuleGroupId() int64 {
	if x != nil {
		return x.SourceHTTPFirewallRuleGroupId
	}
	return 0
}

func (x *IPItem) GetSourceHTTPFirewallRuleSetId() int64 {
	if x != nil {
		return x.SourceHTTPFirewallRuleSetId
	}
	return 0
}

func (x *IPItem) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *IPItem) GetSourceServer() *Server {
	if x != nil {
		return x.SourceServer
	}
	return nil
}

func (x *IPItem) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *IPItem) GetSourceHTTPFirewallPolicy() *HTTPFirewallPolicy {
	if x != nil {
		return x.SourceHTTPFirewallPolicy
	}
	return nil
}

func (x *IPItem) GetSourceHTTPFirewallRuleGroup() *HTTPFirewallRuleGroup {
	if x != nil {
		return x.SourceHTTPFirewallRuleGroup
	}
	return nil
}

func (x *IPItem) GetSourceHTTPFirewallRuleSet() *HTTPFirewallRuleSet {
	if x != nil {
		return x.SourceHTTPFirewallRuleSet
	}
	return nil
}

func (x *IPItem) GetSourceNode() *Node {
	if x != nil {
		return x.SourceNode
	}
	return nil
}

var File_models_model_ip_item_proto protoreflect.FileDescriptor

var file_models_model_ip_item_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69,
	0x70, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x08, 0x0a, 0x06, 0x49, 0x50, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x54, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x1a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x1d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x1d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x40, 0x0a, 0x1b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x22, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x52, 0x0a, 0x18, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x18, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x5b, 0x0a, 0x1b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x1b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x55, 0x0a, 0x19, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x19, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_models_model_ip_item_proto_rawDescOnce sync.Once
	file_models_model_ip_item_proto_rawDescData = file_models_model_ip_item_proto_rawDesc
)

func file_models_model_ip_item_proto_rawDescGZIP() []byte {
	file_models_model_ip_item_proto_rawDescOnce.Do(func() {
		file_models_model_ip_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ip_item_proto_rawDescData)
	})
	return file_models_model_ip_item_proto_rawDescData
}

var file_models_model_ip_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ip_item_proto_goTypes = []interface{}{
	(*IPItem)(nil),                // 0: pb.IPItem
	(*Server)(nil),                // 1: pb.Server
	(*HTTPFirewallPolicy)(nil),    // 2: pb.HTTPFirewallPolicy
	(*HTTPFirewallRuleGroup)(nil), // 3: pb.HTTPFirewallRuleGroup
	(*HTTPFirewallRuleSet)(nil),   // 4: pb.HTTPFirewallRuleSet
	(*Node)(nil),                  // 5: pb.Node
}
var file_models_model_ip_item_proto_depIdxs = []int32{
	1, // 0: pb.IPItem.sourceServer:type_name -> pb.Server
	1, // 1: pb.IPItem.server:type_name -> pb.Server
	2, // 2: pb.IPItem.sourceHTTPFirewallPolicy:type_name -> pb.HTTPFirewallPolicy
	3, // 3: pb.IPItem.sourceHTTPFirewallRuleGroup:type_name -> pb.HTTPFirewallRuleGroup
	4, // 4: pb.IPItem.sourceHTTPFirewallRuleSet:type_name -> pb.HTTPFirewallRuleSet
	5, // 5: pb.IPItem.sourceNode:type_name -> pb.Node
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_models_model_ip_item_proto_init() }
func file_models_model_ip_item_proto_init() {
	if File_models_model_ip_item_proto != nil {
		return
	}
	file_models_model_http_firewall_policy_proto_init()
	file_models_model_http_firewall_rule_group_proto_init()
	file_models_model_http_firewall_rule_set_proto_init()
	file_models_model_server_proto_init()
	file_models_model_node_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ip_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPItem); i {
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
			RawDescriptor: file_models_model_ip_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ip_item_proto_goTypes,
		DependencyIndexes: file_models_model_ip_item_proto_depIdxs,
		MessageInfos:      file_models_model_ip_item_proto_msgTypes,
	}.Build()
	File_models_model_ip_item_proto = out.File
	file_models_model_ip_item_proto_rawDesc = nil
	file_models_model_ip_item_proto_goTypes = nil
	file_models_model_ip_item_proto_depIdxs = nil
}
