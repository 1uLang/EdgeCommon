// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service_region_province.proto

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

// 查找所有省份
type FindAllEnabledRegionProvincesWithCountryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryId int64 `protobuf:"varint,1,opt,name=countryId,proto3" json:"countryId,omitempty"`
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) Reset() {
	*x = FindAllEnabledRegionProvincesWithCountryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvincesWithCountryIdRequest) ProtoMessage() {}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvincesWithCountryIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvincesWithCountryIdRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

type FindAllEnabledRegionProvincesWithCountryIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provinces []*RegionProvince `protobuf:"bytes,1,rep,name=provinces,proto3" json:"provinces,omitempty"`
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) Reset() {
	*x = FindAllEnabledRegionProvincesWithCountryIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvincesWithCountryIdResponse) ProtoMessage() {}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvincesWithCountryIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvincesWithCountryIdResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) GetProvinces() []*RegionProvince {
	if x != nil {
		return x.Provinces
	}
	return nil
}

// 查找单个省份信息
type FindEnabledRegionProvinceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceId int64 `protobuf:"varint,1,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
}

func (x *FindEnabledRegionProvinceRequest) Reset() {
	*x = FindEnabledRegionProvinceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProvinceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProvinceRequest) ProtoMessage() {}

func (x *FindEnabledRegionProvinceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProvinceRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProvinceRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{2}
}

func (x *FindEnabledRegionProvinceRequest) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

type FindEnabledRegionProvinceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Province *RegionProvince `protobuf:"bytes,1,opt,name=province,proto3" json:"province,omitempty"`
}

func (x *FindEnabledRegionProvinceResponse) Reset() {
	*x = FindEnabledRegionProvinceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProvinceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProvinceResponse) ProtoMessage() {}

func (x *FindEnabledRegionProvinceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProvinceResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProvinceResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledRegionProvinceResponse) GetProvince() *RegionProvince {
	if x != nil {
		return x.Province
	}
	return nil
}

var File_service_region_province_proto protoreflect.FileDescriptor

var file_service_region_province_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x31, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x32, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x42, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x32, 0x9f, 0x02, 0x0a, 0x15,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x2a, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x35, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x19, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_region_province_proto_rawDescOnce sync.Once
	file_service_region_province_proto_rawDescData = file_service_region_province_proto_rawDesc
)

func file_service_region_province_proto_rawDescGZIP() []byte {
	file_service_region_province_proto_rawDescOnce.Do(func() {
		file_service_region_province_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_region_province_proto_rawDescData)
	})
	return file_service_region_province_proto_rawDescData
}

var file_service_region_province_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_region_province_proto_goTypes = []interface{}{
	(*FindAllEnabledRegionProvincesWithCountryIdRequest)(nil),  // 0: pb.FindAllEnabledRegionProvincesWithCountryIdRequest
	(*FindAllEnabledRegionProvincesWithCountryIdResponse)(nil), // 1: pb.FindAllEnabledRegionProvincesWithCountryIdResponse
	(*FindEnabledRegionProvinceRequest)(nil),                   // 2: pb.FindEnabledRegionProvinceRequest
	(*FindEnabledRegionProvinceResponse)(nil),                  // 3: pb.FindEnabledRegionProvinceResponse
	(*RegionProvince)(nil),                                     // 4: pb.RegionProvince
}
var file_service_region_province_proto_depIdxs = []int32{
	4, // 0: pb.FindAllEnabledRegionProvincesWithCountryIdResponse.provinces:type_name -> pb.RegionProvince
	4, // 1: pb.FindEnabledRegionProvinceResponse.province:type_name -> pb.RegionProvince
	0, // 2: pb.RegionProvinceService.findAllEnabledRegionProvincesWithCountryId:input_type -> pb.FindAllEnabledRegionProvincesWithCountryIdRequest
	2, // 3: pb.RegionProvinceService.findEnabledRegionProvince:input_type -> pb.FindEnabledRegionProvinceRequest
	1, // 4: pb.RegionProvinceService.findAllEnabledRegionProvincesWithCountryId:output_type -> pb.FindAllEnabledRegionProvincesWithCountryIdResponse
	3, // 5: pb.RegionProvinceService.findEnabledRegionProvince:output_type -> pb.FindEnabledRegionProvinceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_region_province_proto_init() }
func file_service_region_province_proto_init() {
	if File_service_region_province_proto != nil {
		return
	}
	file_models_model_region_province_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_region_province_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvincesWithCountryIdRequest); i {
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
		file_service_region_province_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvincesWithCountryIdResponse); i {
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
		file_service_region_province_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProvinceRequest); i {
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
		file_service_region_province_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProvinceResponse); i {
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
			RawDescriptor: file_service_region_province_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_region_province_proto_goTypes,
		DependencyIndexes: file_service_region_province_proto_depIdxs,
		MessageInfos:      file_service_region_province_proto_msgTypes,
	}.Build()
	File_service_region_province_proto = out.File
	file_service_region_province_proto_rawDesc = nil
	file_service_region_province_proto_goTypes = nil
	file_service_region_province_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegionProvinceServiceClient is the client API for RegionProvinceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegionProvinceServiceClient interface {
	// 查找所有省份
	FindAllEnabledRegionProvincesWithCountryId(ctx context.Context, in *FindAllEnabledRegionProvincesWithCountryIdRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error)
	// 查找单个省份信息
	FindEnabledRegionProvince(ctx context.Context, in *FindEnabledRegionProvinceRequest, opts ...grpc.CallOption) (*FindEnabledRegionProvinceResponse, error)
}

type regionProvinceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionProvinceServiceClient(cc grpc.ClientConnInterface) RegionProvinceServiceClient {
	return &regionProvinceServiceClient{cc}
}

func (c *regionProvinceServiceClient) FindAllEnabledRegionProvincesWithCountryId(ctx context.Context, in *FindAllEnabledRegionProvincesWithCountryIdRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error) {
	out := new(FindAllEnabledRegionProvincesWithCountryIdResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionProvinceService/findAllEnabledRegionProvincesWithCountryId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionProvinceServiceClient) FindEnabledRegionProvince(ctx context.Context, in *FindEnabledRegionProvinceRequest, opts ...grpc.CallOption) (*FindEnabledRegionProvinceResponse, error) {
	out := new(FindEnabledRegionProvinceResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionProvinceService/findEnabledRegionProvince", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionProvinceServiceServer is the server API for RegionProvinceService service.
type RegionProvinceServiceServer interface {
	// 查找所有省份
	FindAllEnabledRegionProvincesWithCountryId(context.Context, *FindAllEnabledRegionProvincesWithCountryIdRequest) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error)
	// 查找单个省份信息
	FindEnabledRegionProvince(context.Context, *FindEnabledRegionProvinceRequest) (*FindEnabledRegionProvinceResponse, error)
}

// UnimplementedRegionProvinceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegionProvinceServiceServer struct {
}

func (*UnimplementedRegionProvinceServiceServer) FindAllEnabledRegionProvincesWithCountryId(context.Context, *FindAllEnabledRegionProvincesWithCountryIdRequest) (*FindAllEnabledRegionProvincesWithCountryIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledRegionProvincesWithCountryId not implemented")
}
func (*UnimplementedRegionProvinceServiceServer) FindEnabledRegionProvince(context.Context, *FindEnabledRegionProvinceRequest) (*FindEnabledRegionProvinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledRegionProvince not implemented")
}

func RegisterRegionProvinceServiceServer(s *grpc.Server, srv RegionProvinceServiceServer) {
	s.RegisterService(&_RegionProvinceService_serviceDesc, srv)
}

func _RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledRegionProvincesWithCountryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindAllEnabledRegionProvincesWithCountryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionProvinceService/FindAllEnabledRegionProvincesWithCountryId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindAllEnabledRegionProvincesWithCountryId(ctx, req.(*FindAllEnabledRegionProvincesWithCountryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionProvinceService_FindEnabledRegionProvince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledRegionProvinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionProvinceServiceServer).FindEnabledRegionProvince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionProvinceService/FindEnabledRegionProvince",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionProvinceServiceServer).FindEnabledRegionProvince(ctx, req.(*FindEnabledRegionProvinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegionProvinceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegionProvinceService",
	HandlerType: (*RegionProvinceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledRegionProvincesWithCountryId",
			Handler:    _RegionProvinceService_FindAllEnabledRegionProvincesWithCountryId_Handler,
		},
		{
			MethodName: "findEnabledRegionProvince",
			Handler:    _RegionProvinceService_FindEnabledRegionProvince_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_region_province.proto",
}
