// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service_server_region_country_monthly_stat.proto

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

// 查找前N个城市
type FindTopServerRegionCountryMonthlyStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month    string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"` // YYYYMM
	ServerId int64  `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Offset   int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) Reset() {
	*x = FindTopServerRegionCountryMonthlyStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCountryMonthlyStatsRequest) ProtoMessage() {}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCountryMonthlyStatsRequest.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCountryMonthlyStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_region_country_monthly_stat_proto_rawDescGZIP(), []int{0}
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindTopServerRegionCountryMonthlyStatsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindTopServerRegionCountryMonthlyStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*FindTopServerRegionCountryMonthlyStatsResponse_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse) Reset() {
	*x = FindTopServerRegionCountryMonthlyStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCountryMonthlyStatsResponse) ProtoMessage() {}

func (x *FindTopServerRegionCountryMonthlyStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCountryMonthlyStatsResponse.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCountryMonthlyStatsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_region_country_monthly_stat_proto_rawDescGZIP(), []int{1}
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse) GetStats() []*FindTopServerRegionCountryMonthlyStatsResponse_Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type FindTopServerRegionCountryMonthlyStatsResponse_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCountry *RegionCountry `protobuf:"bytes,1,opt,name=regionCountry,proto3" json:"regionCountry,omitempty"`
	Count         int64          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse_Stat) Reset() {
	*x = FindTopServerRegionCountryMonthlyStatsResponse_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCountryMonthlyStatsResponse_Stat) ProtoMessage() {}

func (x *FindTopServerRegionCountryMonthlyStatsResponse_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_country_monthly_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCountryMonthlyStatsResponse_Stat.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCountryMonthlyStatsResponse_Stat) Descriptor() ([]byte, []int) {
	return file_service_server_region_country_monthly_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse_Stat) GetRegionCountry() *RegionCountry {
	if x != nil {
		return x.RegionCountry
	}
	return nil
}

func (x *FindTopServerRegionCountryMonthlyStatsResponse_Stat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_server_region_country_monthly_stat_proto protoreflect.FileDescriptor

var file_service_server_region_country_monthly_stat_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x2d, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x55, 0x0a, 0x04, 0x53,
	0x74, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xb9, 0x01, 0x0a, 0x25, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a,
	0x26, 0x66, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_region_country_monthly_stat_proto_rawDescOnce sync.Once
	file_service_server_region_country_monthly_stat_proto_rawDescData = file_service_server_region_country_monthly_stat_proto_rawDesc
)

func file_service_server_region_country_monthly_stat_proto_rawDescGZIP() []byte {
	file_service_server_region_country_monthly_stat_proto_rawDescOnce.Do(func() {
		file_service_server_region_country_monthly_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_region_country_monthly_stat_proto_rawDescData)
	})
	return file_service_server_region_country_monthly_stat_proto_rawDescData
}

var file_service_server_region_country_monthly_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_server_region_country_monthly_stat_proto_goTypes = []interface{}{
	(*FindTopServerRegionCountryMonthlyStatsRequest)(nil),       // 0: pb.FindTopServerRegionCountryMonthlyStatsRequest
	(*FindTopServerRegionCountryMonthlyStatsResponse)(nil),      // 1: pb.FindTopServerRegionCountryMonthlyStatsResponse
	(*FindTopServerRegionCountryMonthlyStatsResponse_Stat)(nil), // 2: pb.FindTopServerRegionCountryMonthlyStatsResponse.Stat
	(*RegionCountry)(nil), // 3: pb.RegionCountry
}
var file_service_server_region_country_monthly_stat_proto_depIdxs = []int32{
	2, // 0: pb.FindTopServerRegionCountryMonthlyStatsResponse.stats:type_name -> pb.FindTopServerRegionCountryMonthlyStatsResponse.Stat
	3, // 1: pb.FindTopServerRegionCountryMonthlyStatsResponse.Stat.regionCountry:type_name -> pb.RegionCountry
	0, // 2: pb.ServerRegionCountryMonthlyStatService.findTopServerRegionCountryMonthlyStats:input_type -> pb.FindTopServerRegionCountryMonthlyStatsRequest
	1, // 3: pb.ServerRegionCountryMonthlyStatService.findTopServerRegionCountryMonthlyStats:output_type -> pb.FindTopServerRegionCountryMonthlyStatsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_server_region_country_monthly_stat_proto_init() }
func file_service_server_region_country_monthly_stat_proto_init() {
	if File_service_server_region_country_monthly_stat_proto != nil {
		return
	}
	file_models_model_region_country_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_region_country_monthly_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCountryMonthlyStatsRequest); i {
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
		file_service_server_region_country_monthly_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCountryMonthlyStatsResponse); i {
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
		file_service_server_region_country_monthly_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCountryMonthlyStatsResponse_Stat); i {
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
			RawDescriptor: file_service_server_region_country_monthly_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_region_country_monthly_stat_proto_goTypes,
		DependencyIndexes: file_service_server_region_country_monthly_stat_proto_depIdxs,
		MessageInfos:      file_service_server_region_country_monthly_stat_proto_msgTypes,
	}.Build()
	File_service_server_region_country_monthly_stat_proto = out.File
	file_service_server_region_country_monthly_stat_proto_rawDesc = nil
	file_service_server_region_country_monthly_stat_proto_goTypes = nil
	file_service_server_region_country_monthly_stat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerRegionCountryMonthlyStatServiceClient is the client API for ServerRegionCountryMonthlyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerRegionCountryMonthlyStatServiceClient interface {
	// 查找前N个地区
	FindTopServerRegionCountryMonthlyStats(ctx context.Context, in *FindTopServerRegionCountryMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCountryMonthlyStatsResponse, error)
}

type serverRegionCountryMonthlyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerRegionCountryMonthlyStatServiceClient(cc grpc.ClientConnInterface) ServerRegionCountryMonthlyStatServiceClient {
	return &serverRegionCountryMonthlyStatServiceClient{cc}
}

func (c *serverRegionCountryMonthlyStatServiceClient) FindTopServerRegionCountryMonthlyStats(ctx context.Context, in *FindTopServerRegionCountryMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCountryMonthlyStatsResponse, error) {
	out := new(FindTopServerRegionCountryMonthlyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerRegionCountryMonthlyStatService/findTopServerRegionCountryMonthlyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerRegionCountryMonthlyStatServiceServer is the server API for ServerRegionCountryMonthlyStatService service.
type ServerRegionCountryMonthlyStatServiceServer interface {
	// 查找前N个地区
	FindTopServerRegionCountryMonthlyStats(context.Context, *FindTopServerRegionCountryMonthlyStatsRequest) (*FindTopServerRegionCountryMonthlyStatsResponse, error)
}

// UnimplementedServerRegionCountryMonthlyStatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerRegionCountryMonthlyStatServiceServer struct {
}

func (*UnimplementedServerRegionCountryMonthlyStatServiceServer) FindTopServerRegionCountryMonthlyStats(context.Context, *FindTopServerRegionCountryMonthlyStatsRequest) (*FindTopServerRegionCountryMonthlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTopServerRegionCountryMonthlyStats not implemented")
}

func RegisterServerRegionCountryMonthlyStatServiceServer(s *grpc.Server, srv ServerRegionCountryMonthlyStatServiceServer) {
	s.RegisterService(&_ServerRegionCountryMonthlyStatService_serviceDesc, srv)
}

func _ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTopServerRegionCountryMonthlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRegionCountryMonthlyStatServiceServer).FindTopServerRegionCountryMonthlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerRegionCountryMonthlyStatService/FindTopServerRegionCountryMonthlyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRegionCountryMonthlyStatServiceServer).FindTopServerRegionCountryMonthlyStats(ctx, req.(*FindTopServerRegionCountryMonthlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerRegionCountryMonthlyStatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerRegionCountryMonthlyStatService",
	HandlerType: (*ServerRegionCountryMonthlyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findTopServerRegionCountryMonthlyStats",
			Handler:    _ServerRegionCountryMonthlyStatService_FindTopServerRegionCountryMonthlyStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_region_country_monthly_stat.proto",
}
