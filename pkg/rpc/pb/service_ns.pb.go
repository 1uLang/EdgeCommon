// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: service_ns.proto

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

// 组合看板数据
type ComposeNSBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ComposeNSBoardRequest) Reset() {
	*x = ComposeNSBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardRequest) ProtoMessage() {}

func (x *ComposeNSBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardRequest.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{0}
}

type ComposeNSBoardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountNSDomains      int64                                       `protobuf:"varint,1,opt,name=countNSDomains,proto3" json:"countNSDomains,omitempty"`
	CountNSRecords      int64                                       `protobuf:"varint,2,opt,name=countNSRecords,proto3" json:"countNSRecords,omitempty"`
	CountNSClusters     int64                                       `protobuf:"varint,3,opt,name=countNSClusters,proto3" json:"countNSClusters,omitempty"`
	CountNSNodes        int64                                       `protobuf:"varint,4,opt,name=countNSNodes,proto3" json:"countNSNodes,omitempty"`
	CountOfflineNSNodes int64                                       `protobuf:"varint,5,opt,name=countOfflineNSNodes,proto3" json:"countOfflineNSNodes,omitempty"`
	DailyTrafficStats   []*ComposeNSBoardResponse_DailyTrafficStat  `protobuf:"bytes,30,rep,name=dailyTrafficStats,proto3" json:"dailyTrafficStats,omitempty"`
	HourlyTrafficStats  []*ComposeNSBoardResponse_HourlyTrafficStat `protobuf:"bytes,31,rep,name=hourlyTrafficStats,proto3" json:"hourlyTrafficStats,omitempty"`
	TopNSNodeStats      []*ComposeNSBoardResponse_NodeStat          `protobuf:"bytes,32,rep,name=topNSNodeStats,proto3" json:"topNSNodeStats,omitempty"`
	TopNSDomainStats    []*ComposeNSBoardResponse_DomainStat        `protobuf:"bytes,33,rep,name=topNSDomainStats,proto3" json:"topNSDomainStats,omitempty"`
	CpuNodeValues       []*NodeValue                                `protobuf:"bytes,34,rep,name=cpuNodeValues,proto3" json:"cpuNodeValues,omitempty"`
	MemoryNodeValues    []*NodeValue                                `protobuf:"bytes,35,rep,name=memoryNodeValues,proto3" json:"memoryNodeValues,omitempty"`
	LoadNodeValues      []*NodeValue                                `protobuf:"bytes,36,rep,name=loadNodeValues,proto3" json:"loadNodeValues,omitempty"`
}

func (x *ComposeNSBoardResponse) Reset() {
	*x = ComposeNSBoardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardResponse) ProtoMessage() {}

func (x *ComposeNSBoardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardResponse.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeNSBoardResponse) GetCountNSDomains() int64 {
	if x != nil {
		return x.CountNSDomains
	}
	return 0
}

func (x *ComposeNSBoardResponse) GetCountNSRecords() int64 {
	if x != nil {
		return x.CountNSRecords
	}
	return 0
}

func (x *ComposeNSBoardResponse) GetCountNSClusters() int64 {
	if x != nil {
		return x.CountNSClusters
	}
	return 0
}

func (x *ComposeNSBoardResponse) GetCountNSNodes() int64 {
	if x != nil {
		return x.CountNSNodes
	}
	return 0
}

func (x *ComposeNSBoardResponse) GetCountOfflineNSNodes() int64 {
	if x != nil {
		return x.CountOfflineNSNodes
	}
	return 0
}

func (x *ComposeNSBoardResponse) GetDailyTrafficStats() []*ComposeNSBoardResponse_DailyTrafficStat {
	if x != nil {
		return x.DailyTrafficStats
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetHourlyTrafficStats() []*ComposeNSBoardResponse_HourlyTrafficStat {
	if x != nil {
		return x.HourlyTrafficStats
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetTopNSNodeStats() []*ComposeNSBoardResponse_NodeStat {
	if x != nil {
		return x.TopNSNodeStats
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetTopNSDomainStats() []*ComposeNSBoardResponse_DomainStat {
	if x != nil {
		return x.TopNSDomainStats
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetCpuNodeValues() []*NodeValue {
	if x != nil {
		return x.CpuNodeValues
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetMemoryNodeValues() []*NodeValue {
	if x != nil {
		return x.MemoryNodeValues
	}
	return nil
}

func (x *ComposeNSBoardResponse) GetLoadNodeValues() []*NodeValue {
	if x != nil {
		return x.LoadNodeValues
	}
	return nil
}

type ComposeNSBoardResponse_DailyTrafficStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day           string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
	Bytes         int64  `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	CountRequests int64  `protobuf:"varint,3,opt,name=countRequests,proto3" json:"countRequests,omitempty"`
}

func (x *ComposeNSBoardResponse_DailyTrafficStat) Reset() {
	*x = ComposeNSBoardResponse_DailyTrafficStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardResponse_DailyTrafficStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardResponse_DailyTrafficStat) ProtoMessage() {}

func (x *ComposeNSBoardResponse_DailyTrafficStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardResponse_DailyTrafficStat.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardResponse_DailyTrafficStat) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ComposeNSBoardResponse_DailyTrafficStat) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ComposeNSBoardResponse_DailyTrafficStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *ComposeNSBoardResponse_DailyTrafficStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

type ComposeNSBoardResponse_HourlyTrafficStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hour          string `protobuf:"bytes,1,opt,name=hour,proto3" json:"hour,omitempty"`
	Bytes         int64  `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	CountRequests int64  `protobuf:"varint,3,opt,name=countRequests,proto3" json:"countRequests,omitempty"`
}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) Reset() {
	*x = ComposeNSBoardResponse_HourlyTrafficStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardResponse_HourlyTrafficStat) ProtoMessage() {}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardResponse_HourlyTrafficStat.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardResponse_HourlyTrafficStat) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) GetHour() string {
	if x != nil {
		return x.Hour
	}
	return ""
}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *ComposeNSBoardResponse_HourlyTrafficStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

type ComposeNSBoardResponse_NodeStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsClusterId   int64  `protobuf:"varint,1,opt,name=nsClusterId,proto3" json:"nsClusterId,omitempty"`
	NsNodeId      int64  `protobuf:"varint,2,opt,name=nsNodeId,proto3" json:"nsNodeId,omitempty"`
	NsNodeName    string `protobuf:"bytes,3,opt,name=nsNodeName,proto3" json:"nsNodeName,omitempty"`
	CountRequests int64  `protobuf:"varint,4,opt,name=countRequests,proto3" json:"countRequests,omitempty"`
	Bytes         int64  `protobuf:"varint,5,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *ComposeNSBoardResponse_NodeStat) Reset() {
	*x = ComposeNSBoardResponse_NodeStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardResponse_NodeStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardResponse_NodeStat) ProtoMessage() {}

func (x *ComposeNSBoardResponse_NodeStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardResponse_NodeStat.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardResponse_NodeStat) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{1, 2}
}

func (x *ComposeNSBoardResponse_NodeStat) GetNsClusterId() int64 {
	if x != nil {
		return x.NsClusterId
	}
	return 0
}

func (x *ComposeNSBoardResponse_NodeStat) GetNsNodeId() int64 {
	if x != nil {
		return x.NsNodeId
	}
	return 0
}

func (x *ComposeNSBoardResponse_NodeStat) GetNsNodeName() string {
	if x != nil {
		return x.NsNodeName
	}
	return ""
}

func (x *ComposeNSBoardResponse_NodeStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

func (x *ComposeNSBoardResponse_NodeStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

type ComposeNSBoardResponse_DomainStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsDomainId    int64  `protobuf:"varint,1,opt,name=nsDomainId,proto3" json:"nsDomainId,omitempty"`
	NsDomainName  string `protobuf:"bytes,2,opt,name=nsDomainName,proto3" json:"nsDomainName,omitempty"`
	CountRequests int64  `protobuf:"varint,3,opt,name=countRequests,proto3" json:"countRequests,omitempty"`
	Bytes         int64  `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *ComposeNSBoardResponse_DomainStat) Reset() {
	*x = ComposeNSBoardResponse_DomainStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeNSBoardResponse_DomainStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeNSBoardResponse_DomainStat) ProtoMessage() {}

func (x *ComposeNSBoardResponse_DomainStat) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeNSBoardResponse_DomainStat.ProtoReflect.Descriptor instead.
func (*ComposeNSBoardResponse_DomainStat) Descriptor() ([]byte, []int) {
	return file_service_ns_proto_rawDescGZIP(), []int{1, 3}
}

func (x *ComposeNSBoardResponse_DomainStat) GetNsDomainId() int64 {
	if x != nil {
		return x.NsDomainId
	}
	return 0
}

func (x *ComposeNSBoardResponse_DomainStat) GetNsDomainName() string {
	if x != nil {
		return x.NsDomainName
	}
	return ""
}

func (x *ComposeNSBoardResponse_DomainStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

func (x *ComposeNSBoardResponse_DomainStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

var File_service_ns_proto protoreflect.FileDescriptor

var file_service_ns_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe5,
	0x09, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x53, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x4e, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x11, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x1e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x11, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x5c, 0x0a, 0x12, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x6f, 0x75,
	0x72, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x12,
	0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x74, 0x6f, 0x70, 0x4e, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x0e, 0x74, 0x6f, 0x70, 0x4e, 0x53, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x51, 0x0a, 0x10, 0x74, 0x6f, 0x70, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x21, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x10, 0x74, 0x6f, 0x70, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x63, 0x70, 0x75, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x22, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x63, 0x70, 0x75, 0x4e, 0x6f, 0x64,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x35, 0x0a, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6c, 0x6f, 0x61, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x60, 0x0a, 0x10, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x63, 0x0a, 0x11, 0x48,
	0x6f, 0x75, 0x72, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x75, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x1a, 0xa4, 0x01, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x73, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x8c, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x73, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x73,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x32, 0x54, 0x0a, 0x09, 0x4e, 0x53, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x4e, 0x53, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x4e, 0x53, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ns_proto_rawDescOnce sync.Once
	file_service_ns_proto_rawDescData = file_service_ns_proto_rawDesc
)

func file_service_ns_proto_rawDescGZIP() []byte {
	file_service_ns_proto_rawDescOnce.Do(func() {
		file_service_ns_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ns_proto_rawDescData)
	})
	return file_service_ns_proto_rawDescData
}

var file_service_ns_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_ns_proto_goTypes = []interface{}{
	(*ComposeNSBoardRequest)(nil),                    // 0: pb.ComposeNSBoardRequest
	(*ComposeNSBoardResponse)(nil),                   // 1: pb.ComposeNSBoardResponse
	(*ComposeNSBoardResponse_DailyTrafficStat)(nil),  // 2: pb.ComposeNSBoardResponse.DailyTrafficStat
	(*ComposeNSBoardResponse_HourlyTrafficStat)(nil), // 3: pb.ComposeNSBoardResponse.HourlyTrafficStat
	(*ComposeNSBoardResponse_NodeStat)(nil),          // 4: pb.ComposeNSBoardResponse.NodeStat
	(*ComposeNSBoardResponse_DomainStat)(nil),        // 5: pb.ComposeNSBoardResponse.DomainStat
	(*NodeValue)(nil),                                // 6: pb.NodeValue
}
var file_service_ns_proto_depIdxs = []int32{
	2, // 0: pb.ComposeNSBoardResponse.dailyTrafficStats:type_name -> pb.ComposeNSBoardResponse.DailyTrafficStat
	3, // 1: pb.ComposeNSBoardResponse.hourlyTrafficStats:type_name -> pb.ComposeNSBoardResponse.HourlyTrafficStat
	4, // 2: pb.ComposeNSBoardResponse.topNSNodeStats:type_name -> pb.ComposeNSBoardResponse.NodeStat
	5, // 3: pb.ComposeNSBoardResponse.topNSDomainStats:type_name -> pb.ComposeNSBoardResponse.DomainStat
	6, // 4: pb.ComposeNSBoardResponse.cpuNodeValues:type_name -> pb.NodeValue
	6, // 5: pb.ComposeNSBoardResponse.memoryNodeValues:type_name -> pb.NodeValue
	6, // 6: pb.ComposeNSBoardResponse.loadNodeValues:type_name -> pb.NodeValue
	0, // 7: pb.NSService.composeNSBoard:input_type -> pb.ComposeNSBoardRequest
	1, // 8: pb.NSService.composeNSBoard:output_type -> pb.ComposeNSBoardResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_service_ns_proto_init() }
func file_service_ns_proto_init() {
	if File_service_ns_proto != nil {
		return
	}
	file_models_model_node_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardRequest); i {
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
		file_service_ns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardResponse); i {
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
		file_service_ns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardResponse_DailyTrafficStat); i {
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
		file_service_ns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardResponse_HourlyTrafficStat); i {
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
		file_service_ns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardResponse_NodeStat); i {
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
		file_service_ns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeNSBoardResponse_DomainStat); i {
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
			RawDescriptor: file_service_ns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ns_proto_goTypes,
		DependencyIndexes: file_service_ns_proto_depIdxs,
		MessageInfos:      file_service_ns_proto_msgTypes,
	}.Build()
	File_service_ns_proto = out.File
	file_service_ns_proto_rawDesc = nil
	file_service_ns_proto_goTypes = nil
	file_service_ns_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NSServiceClient is the client API for NSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NSServiceClient interface {
	// 组合看板数据
	ComposeNSBoard(ctx context.Context, in *ComposeNSBoardRequest, opts ...grpc.CallOption) (*ComposeNSBoardResponse, error)
}

type nSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSServiceClient(cc grpc.ClientConnInterface) NSServiceClient {
	return &nSServiceClient{cc}
}

func (c *nSServiceClient) ComposeNSBoard(ctx context.Context, in *ComposeNSBoardRequest, opts ...grpc.CallOption) (*ComposeNSBoardResponse, error) {
	out := new(ComposeNSBoardResponse)
	err := c.cc.Invoke(ctx, "/pb.NSService/composeNSBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSServiceServer is the server API for NSService service.
type NSServiceServer interface {
	// 组合看板数据
	ComposeNSBoard(context.Context, *ComposeNSBoardRequest) (*ComposeNSBoardResponse, error)
}

// UnimplementedNSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNSServiceServer struct {
}

func (*UnimplementedNSServiceServer) ComposeNSBoard(context.Context, *ComposeNSBoardRequest) (*ComposeNSBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComposeNSBoard not implemented")
}

func RegisterNSServiceServer(s *grpc.Server, srv NSServiceServer) {
	s.RegisterService(&_NSService_serviceDesc, srv)
}

func _NSService_ComposeNSBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComposeNSBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSServiceServer).ComposeNSBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NSService/ComposeNSBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSServiceServer).ComposeNSBoard(ctx, req.(*ComposeNSBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NSService",
	HandlerType: (*NSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "composeNSBoard",
			Handler:    _NSService_ComposeNSBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ns.proto",
}
