// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_server_region_province_monthly_stat.proto

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

// 查找前N个城市
type FindTopServerRegionProvinceMonthlyStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month     string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"` // YYYYMM
	ServerId  int64  `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	CountryId int64  `protobuf:"varint,3,opt,name=countryId,proto3" json:"countryId,omitempty"`
	Offset    int64  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size      int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) Reset() {
	*x = FindTopServerRegionProvinceMonthlyStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionProvinceMonthlyStatsRequest) ProtoMessage() {}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionProvinceMonthlyStatsRequest.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionProvinceMonthlyStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_region_province_monthly_stat_proto_rawDescGZIP(), []int{0}
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindTopServerRegionProvinceMonthlyStatsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindTopServerRegionProvinceMonthlyStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*FindTopServerRegionProvinceMonthlyStatsResponse_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse) Reset() {
	*x = FindTopServerRegionProvinceMonthlyStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionProvinceMonthlyStatsResponse) ProtoMessage() {}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionProvinceMonthlyStatsResponse.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionProvinceMonthlyStatsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_region_province_monthly_stat_proto_rawDescGZIP(), []int{1}
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse) GetStats() []*FindTopServerRegionProvinceMonthlyStatsResponse_Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type FindTopServerRegionProvinceMonthlyStatsResponse_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCountry  *RegionCountry  `protobuf:"bytes,1,opt,name=regionCountry,proto3" json:"regionCountry,omitempty"`
	RegionProvince *RegionProvince `protobuf:"bytes,2,opt,name=regionProvince,proto3" json:"regionProvince,omitempty"`
	Count          int64           `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) Reset() {
	*x = FindTopServerRegionProvinceMonthlyStatsResponse_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionProvinceMonthlyStatsResponse_Stat) ProtoMessage() {}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_province_monthly_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionProvinceMonthlyStatsResponse_Stat.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionProvinceMonthlyStatsResponse_Stat) Descriptor() ([]byte, []int) {
	return file_service_server_region_province_monthly_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) GetRegionCountry() *RegionCountry {
	if x != nil {
		return x.RegionCountry
	}
	return nil
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) GetRegionProvince() *RegionProvince {
	if x != nil {
		return x.RegionProvince
	}
	return nil
}

func (x *FindTopServerRegionProvinceMonthlyStatsResponse_Stat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_server_region_province_monthly_stat_proto protoreflect.FileDescriptor

var file_service_server_region_province_monthly_stat_proto_rawDesc = []byte{
	0x0a, 0x31, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x01, 0x0a, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x95, 0x02,
	0x0a, 0x2f, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x1a, 0x91, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbd, 0x01, 0x0a, 0x26, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x92, 0x01, 0x0a, 0x27, 0x66, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_region_province_monthly_stat_proto_rawDescOnce sync.Once
	file_service_server_region_province_monthly_stat_proto_rawDescData = file_service_server_region_province_monthly_stat_proto_rawDesc
)

func file_service_server_region_province_monthly_stat_proto_rawDescGZIP() []byte {
	file_service_server_region_province_monthly_stat_proto_rawDescOnce.Do(func() {
		file_service_server_region_province_monthly_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_region_province_monthly_stat_proto_rawDescData)
	})
	return file_service_server_region_province_monthly_stat_proto_rawDescData
}

var file_service_server_region_province_monthly_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_server_region_province_monthly_stat_proto_goTypes = []interface{}{
	(*FindTopServerRegionProvinceMonthlyStatsRequest)(nil),       // 0: pb.FindTopServerRegionProvinceMonthlyStatsRequest
	(*FindTopServerRegionProvinceMonthlyStatsResponse)(nil),      // 1: pb.FindTopServerRegionProvinceMonthlyStatsResponse
	(*FindTopServerRegionProvinceMonthlyStatsResponse_Stat)(nil), // 2: pb.FindTopServerRegionProvinceMonthlyStatsResponse.Stat
	(*RegionCountry)(nil),  // 3: pb.RegionCountry
	(*RegionProvince)(nil), // 4: pb.RegionProvince
}
var file_service_server_region_province_monthly_stat_proto_depIdxs = []int32{
	2, // 0: pb.FindTopServerRegionProvinceMonthlyStatsResponse.stats:type_name -> pb.FindTopServerRegionProvinceMonthlyStatsResponse.Stat
	3, // 1: pb.FindTopServerRegionProvinceMonthlyStatsResponse.Stat.regionCountry:type_name -> pb.RegionCountry
	4, // 2: pb.FindTopServerRegionProvinceMonthlyStatsResponse.Stat.regionProvince:type_name -> pb.RegionProvince
	0, // 3: pb.ServerRegionProvinceMonthlyStatService.findTopServerRegionProvinceMonthlyStats:input_type -> pb.FindTopServerRegionProvinceMonthlyStatsRequest
	1, // 4: pb.ServerRegionProvinceMonthlyStatService.findTopServerRegionProvinceMonthlyStats:output_type -> pb.FindTopServerRegionProvinceMonthlyStatsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_server_region_province_monthly_stat_proto_init() }
func file_service_server_region_province_monthly_stat_proto_init() {
	if File_service_server_region_province_monthly_stat_proto != nil {
		return
	}
	file_models_model_region_country_proto_init()
	file_models_model_region_province_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_region_province_monthly_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionProvinceMonthlyStatsRequest); i {
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
		file_service_server_region_province_monthly_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionProvinceMonthlyStatsResponse); i {
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
		file_service_server_region_province_monthly_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionProvinceMonthlyStatsResponse_Stat); i {
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
			RawDescriptor: file_service_server_region_province_monthly_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_region_province_monthly_stat_proto_goTypes,
		DependencyIndexes: file_service_server_region_province_monthly_stat_proto_depIdxs,
		MessageInfos:      file_service_server_region_province_monthly_stat_proto_msgTypes,
	}.Build()
	File_service_server_region_province_monthly_stat_proto = out.File
	file_service_server_region_province_monthly_stat_proto_rawDesc = nil
	file_service_server_region_province_monthly_stat_proto_goTypes = nil
	file_service_server_region_province_monthly_stat_proto_depIdxs = nil
}
