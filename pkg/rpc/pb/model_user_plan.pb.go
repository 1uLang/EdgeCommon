// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_user_plan.proto

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

type UserPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PlanId int64   `protobuf:"varint,3,opt,name=planId,proto3" json:"planId,omitempty"`
	IsOn   bool    `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`
	DayTo  string  `protobuf:"bytes,5,opt,name=dayTo,proto3" json:"dayTo,omitempty"`
	Name   string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	User   *User   `protobuf:"bytes,30,opt,name=user,proto3" json:"user,omitempty"`
	Plan   *Plan   `protobuf:"bytes,31,opt,name=plan,proto3" json:"plan,omitempty"`
	Server *Server `protobuf:"bytes,32,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *UserPlan) Reset() {
	*x = UserPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPlan) ProtoMessage() {}

func (x *UserPlan) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPlan.ProtoReflect.Descriptor instead.
func (*UserPlan) Descriptor() ([]byte, []int) {
	return file_models_model_user_plan_proto_rawDescGZIP(), []int{0}
}

func (x *UserPlan) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPlan) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPlan) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

func (x *UserPlan) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *UserPlan) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *UserPlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPlan) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserPlan) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *UserPlan) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

var File_models_model_user_plan_proto protoreflect.FileDescriptor

var file_models_model_user_plan_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_plan_proto_rawDescOnce sync.Once
	file_models_model_user_plan_proto_rawDescData = file_models_model_user_plan_proto_rawDesc
)

func file_models_model_user_plan_proto_rawDescGZIP() []byte {
	file_models_model_user_plan_proto_rawDescOnce.Do(func() {
		file_models_model_user_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_plan_proto_rawDescData)
	})
	return file_models_model_user_plan_proto_rawDescData
}

var file_models_model_user_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_plan_proto_goTypes = []interface{}{
	(*UserPlan)(nil), // 0: pb.UserPlan
	(*User)(nil),     // 1: pb.User
	(*Plan)(nil),     // 2: pb.Plan
	(*Server)(nil),   // 3: pb.Server
}
var file_models_model_user_plan_proto_depIdxs = []int32{
	1, // 0: pb.UserPlan.user:type_name -> pb.User
	2, // 1: pb.UserPlan.plan:type_name -> pb.Plan
	3, // 2: pb.UserPlan.server:type_name -> pb.Server
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_model_user_plan_proto_init() }
func file_models_model_user_plan_proto_init() {
	if File_models_model_user_plan_proto != nil {
		return
	}
	file_models_model_user_proto_init()
	file_models_model_plan_proto_init()
	file_models_model_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPlan); i {
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
			RawDescriptor: file_models_model_user_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_plan_proto_goTypes,
		DependencyIndexes: file_models_model_user_plan_proto_depIdxs,
		MessageInfos:      file_models_model_user_plan_proto_msgTypes,
	}.Build()
	File_models_model_user_plan_proto = out.File
	file_models_model_user_plan_proto_rawDesc = nil
	file_models_model_user_plan_proto_goTypes = nil
	file_models_model_user_plan_proto_depIdxs = nil
}
