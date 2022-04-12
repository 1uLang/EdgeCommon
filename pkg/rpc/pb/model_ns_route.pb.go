// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: models/model_ns_route.proto

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

// 线路
type NSRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn       bool       `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name       string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RangesJSON []byte     `protobuf:"bytes,4,opt,name=rangesJSON,proto3" json:"rangesJSON,omitempty"`
	IsDeleted  bool       `protobuf:"varint,5,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Order      int64      `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
	Version    int64      `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
	Code       string     `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"`
	NsCluster  *NSCluster `protobuf:"bytes,30,opt,name=nsCluster,proto3" json:"nsCluster,omitempty"`
	NsDomain   *NSDomain  `protobuf:"bytes,31,opt,name=nsDomain,proto3" json:"nsDomain,omitempty"`
}

func (x *NSRoute) Reset() {
	*x = NSRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSRoute) ProtoMessage() {}

func (x *NSRoute) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSRoute.ProtoReflect.Descriptor instead.
func (*NSRoute) Descriptor() ([]byte, []int) {
	return file_models_model_ns_route_proto_rawDescGZIP(), []int{0}
}

func (x *NSRoute) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSRoute) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSRoute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSRoute) GetRangesJSON() []byte {
	if x != nil {
		return x.RangesJSON
	}
	return nil
}

func (x *NSRoute) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *NSRoute) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *NSRoute) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NSRoute) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NSRoute) GetNsCluster() *NSCluster {
	if x != nil {
		return x.NsCluster
	}
	return nil
}

func (x *NSRoute) GetNsDomain() *NSDomain {
	if x != nil {
		return x.NsDomain
	}
	return nil
}

var File_models_model_ns_route_proto protoreflect.FileDescriptor

var file_models_model_ns_route_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a,
	0x02, 0x0a, 0x07, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x4a, 0x53,
	0x4f, 0x4e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x09, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_route_proto_rawDescOnce sync.Once
	file_models_model_ns_route_proto_rawDescData = file_models_model_ns_route_proto_rawDesc
)

func file_models_model_ns_route_proto_rawDescGZIP() []byte {
	file_models_model_ns_route_proto_rawDescOnce.Do(func() {
		file_models_model_ns_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_route_proto_rawDescData)
	})
	return file_models_model_ns_route_proto_rawDescData
}

var file_models_model_ns_route_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_route_proto_goTypes = []interface{}{
	(*NSRoute)(nil),   // 0: pb.NSRoute
	(*NSCluster)(nil), // 1: pb.NSCluster
	(*NSDomain)(nil),  // 2: pb.NSDomain
}
var file_models_model_ns_route_proto_depIdxs = []int32{
	1, // 0: pb.NSRoute.nsCluster:type_name -> pb.NSCluster
	2, // 1: pb.NSRoute.nsDomain:type_name -> pb.NSDomain
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_ns_route_proto_init() }
func file_models_model_ns_route_proto_init() {
	if File_models_model_ns_route_proto != nil {
		return
	}
	file_models_model_ns_cluster_proto_init()
	file_models_model_ns_domain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSRoute); i {
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
			RawDescriptor: file_models_model_ns_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_route_proto_goTypes,
		DependencyIndexes: file_models_model_ns_route_proto_depIdxs,
		MessageInfos:      file_models_model_ns_route_proto_msgTypes,
	}.Build()
	File_models_model_ns_route_proto = out.File
	file_models_model_ns_route_proto_rawDesc = nil
	file_models_model_ns_route_proto_goTypes = nil
	file_models_model_ns_route_proto_depIdxs = nil
}
