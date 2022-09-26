// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: models/model_metric_chart.proto

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

// 指标图表
type MetricChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type            string      `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	WidthDiv        int32       `protobuf:"varint,4,opt,name=widthDiv,proto3" json:"widthDiv,omitempty"`
	ParamsJSON      []byte      `protobuf:"bytes,5,opt,name=paramsJSON,proto3" json:"paramsJSON,omitempty"`
	IsOn            bool        `protobuf:"varint,6,opt,name=isOn,proto3" json:"isOn,omitempty"`
	MaxItems        int32       `protobuf:"varint,7,opt,name=maxItems,proto3" json:"maxItems,omitempty"`
	IgnoreEmptyKeys bool        `protobuf:"varint,8,opt,name=ignoreEmptyKeys,proto3" json:"ignoreEmptyKeys,omitempty"`
	IgnoredKeys     []string    `protobuf:"bytes,9,rep,name=ignoredKeys,proto3" json:"ignoredKeys,omitempty"`
	MetricItem      *MetricItem `protobuf:"bytes,30,opt,name=metricItem,proto3" json:"metricItem,omitempty"`
}

func (x *MetricChart) Reset() {
	*x = MetricChart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_metric_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricChart) ProtoMessage() {}

func (x *MetricChart) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_metric_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricChart.ProtoReflect.Descriptor instead.
func (*MetricChart) Descriptor() ([]byte, []int) {
	return file_models_model_metric_chart_proto_rawDescGZIP(), []int{0}
}

func (x *MetricChart) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetricChart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricChart) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MetricChart) GetWidthDiv() int32 {
	if x != nil {
		return x.WidthDiv
	}
	return 0
}

func (x *MetricChart) GetParamsJSON() []byte {
	if x != nil {
		return x.ParamsJSON
	}
	return nil
}

func (x *MetricChart) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *MetricChart) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *MetricChart) GetIgnoreEmptyKeys() bool {
	if x != nil {
		return x.IgnoreEmptyKeys
	}
	return false
}

func (x *MetricChart) GetIgnoredKeys() []string {
	if x != nil {
		return x.IgnoredKeys
	}
	return nil
}

func (x *MetricChart) GetMetricItem() *MetricItem {
	if x != nil {
		return x.MetricItem
	}
	return nil
}

var File_models_model_metric_chart_proto protoreflect.FileDescriptor

var file_models_model_metric_chart_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x77, 0x69, 0x64, 0x74, 0x68, 0x44, 0x69, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x77, 0x69, 0x64, 0x74, 0x68, 0x44, 0x69, 0x76, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x4b, 0x65,
	0x79, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_metric_chart_proto_rawDescOnce sync.Once
	file_models_model_metric_chart_proto_rawDescData = file_models_model_metric_chart_proto_rawDesc
)

func file_models_model_metric_chart_proto_rawDescGZIP() []byte {
	file_models_model_metric_chart_proto_rawDescOnce.Do(func() {
		file_models_model_metric_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_metric_chart_proto_rawDescData)
	})
	return file_models_model_metric_chart_proto_rawDescData
}

var file_models_model_metric_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_metric_chart_proto_goTypes = []interface{}{
	(*MetricChart)(nil), // 0: pb.MetricChart
	(*MetricItem)(nil),  // 1: pb.MetricItem
}
var file_models_model_metric_chart_proto_depIdxs = []int32{
	1, // 0: pb.MetricChart.metricItem:type_name -> pb.MetricItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_metric_chart_proto_init() }
func file_models_model_metric_chart_proto_init() {
	if File_models_model_metric_chart_proto != nil {
		return
	}
	file_models_model_metric_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_metric_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricChart); i {
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
			RawDescriptor: file_models_model_metric_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_metric_chart_proto_goTypes,
		DependencyIndexes: file_models_model_metric_chart_proto_depIdxs,
		MessageInfos:      file_models_model_metric_chart_proto_msgTypes,
	}.Build()
	File_models_model_metric_chart_proto = out.File
	file_models_model_metric_chart_proto_rawDesc = nil
	file_models_model_metric_chart_proto_goTypes = nil
	file_models_model_metric_chart_proto_depIdxs = nil
}
