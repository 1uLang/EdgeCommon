// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: models/model_plan.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn             bool    `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name             string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ClusterId        int64   `protobuf:"varint,4,opt,name=clusterId,proto3" json:"clusterId,omitempty"`
	TrafficLimitJSON []byte  `protobuf:"bytes,5,opt,name=trafficLimitJSON,proto3" json:"trafficLimitJSON,omitempty"`
	FeaturesJSON     []byte  `protobuf:"bytes,6,opt,name=featuresJSON,proto3" json:"featuresJSON,omitempty"`
	PriceType        string  `protobuf:"bytes,7,opt,name=priceType,proto3" json:"priceType,omitempty"`
	TrafficPriceJSON []byte  `protobuf:"bytes,8,opt,name=trafficPriceJSON,proto3" json:"trafficPriceJSON,omitempty"`
	MonthlyPrice     float32 `protobuf:"fixed32,9,opt,name=monthlyPrice,proto3" json:"monthlyPrice,omitempty"`
	SeasonallyPrice  float32 `protobuf:"fixed32,10,opt,name=seasonallyPrice,proto3" json:"seasonallyPrice,omitempty"`
	YearlyPrice      float32 `protobuf:"fixed32,11,opt,name=yearlyPrice,proto3" json:"yearlyPrice,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_models_model_plan_proto_rawDescGZIP(), []int{0}
}

func (x *Plan) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Plan) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetClusterId() int64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *Plan) GetTrafficLimitJSON() []byte {
	if x != nil {
		return x.TrafficLimitJSON
	}
	return nil
}

func (x *Plan) GetFeaturesJSON() []byte {
	if x != nil {
		return x.FeaturesJSON
	}
	return nil
}

func (x *Plan) GetPriceType() string {
	if x != nil {
		return x.PriceType
	}
	return ""
}

func (x *Plan) GetTrafficPriceJSON() []byte {
	if x != nil {
		return x.TrafficPriceJSON
	}
	return nil
}

func (x *Plan) GetMonthlyPrice() float32 {
	if x != nil {
		return x.MonthlyPrice
	}
	return 0
}

func (x *Plan) GetSeasonallyPrice() float32 {
	if x != nil {
		return x.SeasonallyPrice
	}
	return 0
}

func (x *Plan) GetYearlyPrice() float32 {
	if x != nil {
		return x.YearlyPrice
	}
	return 0
}

var File_models_model_plan_proto protoreflect.FileDescriptor

var file_models_model_plan_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xe6, 0x02,
	0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x79, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x6c,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_plan_proto_rawDescOnce sync.Once
	file_models_model_plan_proto_rawDescData = file_models_model_plan_proto_rawDesc
)

func file_models_model_plan_proto_rawDescGZIP() []byte {
	file_models_model_plan_proto_rawDescOnce.Do(func() {
		file_models_model_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_plan_proto_rawDescData)
	})
	return file_models_model_plan_proto_rawDescData
}

var file_models_model_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_plan_proto_goTypes = []interface{}{
	(*Plan)(nil), // 0: pb.Plan
}
var file_models_model_plan_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_plan_proto_init() }
func file_models_model_plan_proto_init() {
	if File_models_model_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
			RawDescriptor: file_models_model_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_plan_proto_goTypes,
		DependencyIndexes: file_models_model_plan_proto_depIdxs,
		MessageInfos:      file_models_model_plan_proto_msgTypes,
	}.Build()
	File_models_model_plan_proto = out.File
	file_models_model_plan_proto_rawDesc = nil
	file_models_model_plan_proto_goTypes = nil
	file_models_model_plan_proto_depIdxs = nil
}
