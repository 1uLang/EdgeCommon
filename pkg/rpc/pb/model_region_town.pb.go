// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_region_town.proto

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

type RegionTown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Codes        []string    `protobuf:"bytes,3,rep,name=codes,proto3" json:"codes,omitempty"`
	RegionCityId int64       `protobuf:"varint,4,opt,name=regionCityId,proto3" json:"regionCityId,omitempty"`
	CustomName   string      `protobuf:"bytes,5,opt,name=customName,proto3" json:"customName,omitempty"`
	CustomCodes  []string    `protobuf:"bytes,6,rep,name=customCodes,proto3" json:"customCodes,omitempty"`
	DisplayName  string      `protobuf:"bytes,7,opt,name=displayName,proto3" json:"displayName,omitempty"`
	RegionCity   *RegionCity `protobuf:"bytes,30,opt,name=regionCity,proto3" json:"regionCity,omitempty"`
}

func (x *RegionTown) Reset() {
	*x = RegionTown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_region_town_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionTown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionTown) ProtoMessage() {}

func (x *RegionTown) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_region_town_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionTown.ProtoReflect.Descriptor instead.
func (*RegionTown) Descriptor() ([]byte, []int) {
	return file_models_model_region_town_proto_rawDescGZIP(), []int{0}
}

func (x *RegionTown) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegionTown) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionTown) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *RegionTown) GetRegionCityId() int64 {
	if x != nil {
		return x.RegionCityId
	}
	return 0
}

func (x *RegionTown) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

func (x *RegionTown) GetCustomCodes() []string {
	if x != nil {
		return x.CustomCodes
	}
	return nil
}

func (x *RegionTown) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *RegionTown) GetRegionCity() *RegionCity {
	if x != nil {
		return x.RegionCity
	}
	return nil
}

var File_models_model_region_town_proto protoreflect.FileDescriptor

var file_models_model_region_town_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x77, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x69, 0x74, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x69, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_region_town_proto_rawDescOnce sync.Once
	file_models_model_region_town_proto_rawDescData = file_models_model_region_town_proto_rawDesc
)

func file_models_model_region_town_proto_rawDescGZIP() []byte {
	file_models_model_region_town_proto_rawDescOnce.Do(func() {
		file_models_model_region_town_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_region_town_proto_rawDescData)
	})
	return file_models_model_region_town_proto_rawDescData
}

var file_models_model_region_town_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_region_town_proto_goTypes = []interface{}{
	(*RegionTown)(nil), // 0: pb.RegionTown
	(*RegionCity)(nil), // 1: pb.RegionCity
}
var file_models_model_region_town_proto_depIdxs = []int32{
	1, // 0: pb.RegionTown.regionCity:type_name -> pb.RegionCity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_region_town_proto_init() }
func file_models_model_region_town_proto_init() {
	if File_models_model_region_town_proto != nil {
		return
	}
	file_models_model_region_city_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_region_town_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionTown); i {
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
			RawDescriptor: file_models_model_region_town_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_region_town_proto_goTypes,
		DependencyIndexes: file_models_model_region_town_proto_depIdxs,
		MessageInfos:      file_models_model_region_town_proto_msgTypes,
	}.Build()
	File_models_model_region_town_proto = out.File
	file_models_model_region_town_proto_rawDesc = nil
	file_models_model_region_town_proto_goTypes = nil
	file_models_model_region_town_proto_depIdxs = nil
}
