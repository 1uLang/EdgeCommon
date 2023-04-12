// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_region_city.proto

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

// 查找所有城市
type FindAllEnabledRegionCitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeRegionProvince bool `protobuf:"varint,1,opt,name=includeRegionProvince,proto3" json:"includeRegionProvince,omitempty"`
}

func (x *FindAllEnabledRegionCitiesRequest) Reset() {
	*x = FindAllEnabledRegionCitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionCitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionCitiesRequest) ProtoMessage() {}

func (x *FindAllEnabledRegionCitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionCitiesRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionCitiesRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllEnabledRegionCitiesRequest) GetIncludeRegionProvince() bool {
	if x != nil {
		return x.IncludeRegionProvince
	}
	return false
}

type FindAllEnabledRegionCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCities []*RegionCity `protobuf:"bytes,1,rep,name=regionCities,proto3" json:"regionCities,omitempty"`
}

func (x *FindAllEnabledRegionCitiesResponse) Reset() {
	*x = FindAllEnabledRegionCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionCitiesResponse) ProtoMessage() {}

func (x *FindAllEnabledRegionCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionCitiesResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionCitiesResponse) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllEnabledRegionCitiesResponse) GetRegionCities() []*RegionCity {
	if x != nil {
		return x.RegionCities
	}
	return nil
}

// 查找单个城市信息
type FindEnabledRegionCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCityId int64 `protobuf:"varint,1,opt,name=regionCityId,proto3" json:"regionCityId,omitempty"`
}

func (x *FindEnabledRegionCityRequest) Reset() {
	*x = FindEnabledRegionCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionCityRequest) ProtoMessage() {}

func (x *FindEnabledRegionCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionCityRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionCityRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{2}
}

func (x *FindEnabledRegionCityRequest) GetRegionCityId() int64 {
	if x != nil {
		return x.RegionCityId
	}
	return 0
}

type FindEnabledRegionCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCity *RegionCity `protobuf:"bytes,1,opt,name=regionCity,proto3" json:"regionCity,omitempty"`
}

func (x *FindEnabledRegionCityResponse) Reset() {
	*x = FindEnabledRegionCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionCityResponse) ProtoMessage() {}

func (x *FindEnabledRegionCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionCityResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionCityResponse) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledRegionCityResponse) GetRegionCity() *RegionCity {
	if x != nil {
		return x.RegionCity
	}
	return nil
}

// 查找所有城市
type FindAllRegionCitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeRegionProvince bool `protobuf:"varint,1,opt,name=includeRegionProvince,proto3" json:"includeRegionProvince,omitempty"`
}

func (x *FindAllRegionCitiesRequest) Reset() {
	*x = FindAllRegionCitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionCitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionCitiesRequest) ProtoMessage() {}

func (x *FindAllRegionCitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionCitiesRequest.ProtoReflect.Descriptor instead.
func (*FindAllRegionCitiesRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllRegionCitiesRequest) GetIncludeRegionProvince() bool {
	if x != nil {
		return x.IncludeRegionProvince
	}
	return false
}

type FindAllRegionCitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCities []*RegionCity `protobuf:"bytes,1,rep,name=regionCities,proto3" json:"regionCities,omitempty"`
}

func (x *FindAllRegionCitiesResponse) Reset() {
	*x = FindAllRegionCitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionCitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionCitiesResponse) ProtoMessage() {}

func (x *FindAllRegionCitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionCitiesResponse.ProtoReflect.Descriptor instead.
func (*FindAllRegionCitiesResponse) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllRegionCitiesResponse) GetRegionCities() []*RegionCity {
	if x != nil {
		return x.RegionCities
	}
	return nil
}

// 查找某个省份的所有城市
type FindAllRegionCitiesWithRegionProvinceIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinceId int64 `protobuf:"varint,1,opt,name=regionProvinceId,proto3" json:"regionProvinceId,omitempty"`
}

func (x *FindAllRegionCitiesWithRegionProvinceIdRequest) Reset() {
	*x = FindAllRegionCitiesWithRegionProvinceIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionCitiesWithRegionProvinceIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionCitiesWithRegionProvinceIdRequest) ProtoMessage() {}

func (x *FindAllRegionCitiesWithRegionProvinceIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionCitiesWithRegionProvinceIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllRegionCitiesWithRegionProvinceIdRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{6}
}

func (x *FindAllRegionCitiesWithRegionProvinceIdRequest) GetRegionProvinceId() int64 {
	if x != nil {
		return x.RegionProvinceId
	}
	return 0
}

type FindAllRegionCitiesWithRegionProvinceIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCities []*RegionCity `protobuf:"bytes,1,rep,name=regionCities,proto3" json:"regionCities,omitempty"`
}

func (x *FindAllRegionCitiesWithRegionProvinceIdResponse) Reset() {
	*x = FindAllRegionCitiesWithRegionProvinceIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionCitiesWithRegionProvinceIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionCitiesWithRegionProvinceIdResponse) ProtoMessage() {}

func (x *FindAllRegionCitiesWithRegionProvinceIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionCitiesWithRegionProvinceIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllRegionCitiesWithRegionProvinceIdResponse) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{7}
}

func (x *FindAllRegionCitiesWithRegionProvinceIdResponse) GetRegionCities() []*RegionCity {
	if x != nil {
		return x.RegionCities
	}
	return nil
}

// 查找单个城市信息
type FindRegionCityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCityId int64 `protobuf:"varint,1,opt,name=regionCityId,proto3" json:"regionCityId,omitempty"`
}

func (x *FindRegionCityRequest) Reset() {
	*x = FindRegionCityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionCityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionCityRequest) ProtoMessage() {}

func (x *FindRegionCityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionCityRequest.ProtoReflect.Descriptor instead.
func (*FindRegionCityRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{8}
}

func (x *FindRegionCityRequest) GetRegionCityId() int64 {
	if x != nil {
		return x.RegionCityId
	}
	return 0
}

type FindRegionCityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCity *RegionCity `protobuf:"bytes,1,opt,name=regionCity,proto3" json:"regionCity,omitempty"`
}

func (x *FindRegionCityResponse) Reset() {
	*x = FindRegionCityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionCityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionCityResponse) ProtoMessage() {}

func (x *FindRegionCityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionCityResponse.ProtoReflect.Descriptor instead.
func (*FindRegionCityResponse) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{9}
}

func (x *FindRegionCityResponse) GetRegionCity() *RegionCity {
	if x != nil {
		return x.RegionCity
	}
	return nil
}

// 修改城市定制信息
type UpdateRegionCityCustomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCityId int64    `protobuf:"varint,1,opt,name=regionCityId,proto3" json:"regionCityId,omitempty"`
	CustomName   string   `protobuf:"bytes,2,opt,name=customName,proto3" json:"customName,omitempty"`
	CustomCodes  []string `protobuf:"bytes,3,rep,name=customCodes,proto3" json:"customCodes,omitempty"`
}

func (x *UpdateRegionCityCustomRequest) Reset() {
	*x = UpdateRegionCityCustomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_city_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegionCityCustomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionCityCustomRequest) ProtoMessage() {}

func (x *UpdateRegionCityCustomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_city_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionCityCustomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionCityCustomRequest) Descriptor() ([]byte, []int) {
	return file_service_region_city_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRegionCityCustomRequest) GetRegionCityId() int64 {
	if x != nil {
		return x.RegionCityId
	}
	return 0
}

func (x *UpdateRegionCityCustomRequest) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

func (x *UpdateRegionCityCustomRequest) GetCustomCodes() []string {
	if x != nil {
		return x.CustomCodes
	}
	return nil
}

var File_service_region_city_proto protoreflect.FileDescriptor

var file_service_region_city_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x21, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x42, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x69, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x51, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x2f, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69,
	0x74, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x22, 0x3b, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x16, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x32,
	0xeb, 0x04, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x61, 0x0a, 0x15, 0x66, 0x69, 0x6e, 0x64, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x56, 0x0a, 0x13, 0x66, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x27, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x69, 0x74, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_region_city_proto_rawDescOnce sync.Once
	file_service_region_city_proto_rawDescData = file_service_region_city_proto_rawDesc
)

func file_service_region_city_proto_rawDescGZIP() []byte {
	file_service_region_city_proto_rawDescOnce.Do(func() {
		file_service_region_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_region_city_proto_rawDescData)
	})
	return file_service_region_city_proto_rawDescData
}

var file_service_region_city_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_region_city_proto_goTypes = []interface{}{
	(*FindAllEnabledRegionCitiesRequest)(nil),               // 0: pb.FindAllEnabledRegionCitiesRequest
	(*FindAllEnabledRegionCitiesResponse)(nil),              // 1: pb.FindAllEnabledRegionCitiesResponse
	(*FindEnabledRegionCityRequest)(nil),                    // 2: pb.FindEnabledRegionCityRequest
	(*FindEnabledRegionCityResponse)(nil),                   // 3: pb.FindEnabledRegionCityResponse
	(*FindAllRegionCitiesRequest)(nil),                      // 4: pb.FindAllRegionCitiesRequest
	(*FindAllRegionCitiesResponse)(nil),                     // 5: pb.FindAllRegionCitiesResponse
	(*FindAllRegionCitiesWithRegionProvinceIdRequest)(nil),  // 6: pb.FindAllRegionCitiesWithRegionProvinceIdRequest
	(*FindAllRegionCitiesWithRegionProvinceIdResponse)(nil), // 7: pb.FindAllRegionCitiesWithRegionProvinceIdResponse
	(*FindRegionCityRequest)(nil),                           // 8: pb.FindRegionCityRequest
	(*FindRegionCityResponse)(nil),                          // 9: pb.FindRegionCityResponse
	(*UpdateRegionCityCustomRequest)(nil),                   // 10: pb.UpdateRegionCityCustomRequest
	(*RegionCity)(nil),                                      // 11: pb.RegionCity
	(*RPCSuccess)(nil),                                      // 12: pb.RPCSuccess
}
var file_service_region_city_proto_depIdxs = []int32{
	11, // 0: pb.FindAllEnabledRegionCitiesResponse.regionCities:type_name -> pb.RegionCity
	11, // 1: pb.FindEnabledRegionCityResponse.regionCity:type_name -> pb.RegionCity
	11, // 2: pb.FindAllRegionCitiesResponse.regionCities:type_name -> pb.RegionCity
	11, // 3: pb.FindAllRegionCitiesWithRegionProvinceIdResponse.regionCities:type_name -> pb.RegionCity
	11, // 4: pb.FindRegionCityResponse.regionCity:type_name -> pb.RegionCity
	0,  // 5: pb.RegionCityService.findAllEnabledRegionCities:input_type -> pb.FindAllEnabledRegionCitiesRequest
	2,  // 6: pb.RegionCityService.findEnabledRegionCity:input_type -> pb.FindEnabledRegionCityRequest
	4,  // 7: pb.RegionCityService.findAllRegionCities:input_type -> pb.FindAllRegionCitiesRequest
	6,  // 8: pb.RegionCityService.findAllRegionCitiesWithRegionProvinceId:input_type -> pb.FindAllRegionCitiesWithRegionProvinceIdRequest
	8,  // 9: pb.RegionCityService.findRegionCity:input_type -> pb.FindRegionCityRequest
	10, // 10: pb.RegionCityService.updateRegionCityCustom:input_type -> pb.UpdateRegionCityCustomRequest
	1,  // 11: pb.RegionCityService.findAllEnabledRegionCities:output_type -> pb.FindAllEnabledRegionCitiesResponse
	3,  // 12: pb.RegionCityService.findEnabledRegionCity:output_type -> pb.FindEnabledRegionCityResponse
	5,  // 13: pb.RegionCityService.findAllRegionCities:output_type -> pb.FindAllRegionCitiesResponse
	7,  // 14: pb.RegionCityService.findAllRegionCitiesWithRegionProvinceId:output_type -> pb.FindAllRegionCitiesWithRegionProvinceIdResponse
	9,  // 15: pb.RegionCityService.findRegionCity:output_type -> pb.FindRegionCityResponse
	12, // 16: pb.RegionCityService.updateRegionCityCustom:output_type -> pb.RPCSuccess
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_service_region_city_proto_init() }
func file_service_region_city_proto_init() {
	if File_service_region_city_proto != nil {
		return
	}
	file_models_model_region_city_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_region_city_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionCitiesRequest); i {
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
		file_service_region_city_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionCitiesResponse); i {
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
		file_service_region_city_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionCityRequest); i {
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
		file_service_region_city_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionCityResponse); i {
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
		file_service_region_city_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionCitiesRequest); i {
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
		file_service_region_city_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionCitiesResponse); i {
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
		file_service_region_city_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionCitiesWithRegionProvinceIdRequest); i {
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
		file_service_region_city_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionCitiesWithRegionProvinceIdResponse); i {
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
		file_service_region_city_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionCityRequest); i {
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
		file_service_region_city_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionCityResponse); i {
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
		file_service_region_city_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegionCityCustomRequest); i {
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
			RawDescriptor: file_service_region_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_region_city_proto_goTypes,
		DependencyIndexes: file_service_region_city_proto_depIdxs,
		MessageInfos:      file_service_region_city_proto_msgTypes,
	}.Build()
	File_service_region_city_proto = out.File
	file_service_region_city_proto_rawDesc = nil
	file_service_region_city_proto_goTypes = nil
	file_service_region_city_proto_depIdxs = nil
}
