// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: models/model_node_grant.proto

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

type NodeGrant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Method      string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Su          bool   `protobuf:"varint,6,opt,name=su,proto3" json:"su,omitempty"`
	PrivateKey  string `protobuf:"bytes,7,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	Passphrase  string `protobuf:"bytes,10,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	NodeId      int64  `protobuf:"varint,9,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *NodeGrant) Reset() {
	*x = NodeGrant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_node_grant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGrant) ProtoMessage() {}

func (x *NodeGrant) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_node_grant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGrant.ProtoReflect.Descriptor instead.
func (*NodeGrant) Descriptor() ([]byte, []int) {
	return file_models_model_node_grant_proto_rawDescGZIP(), []int{0}
}

func (x *NodeGrant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeGrant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeGrant) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *NodeGrant) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NodeGrant) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NodeGrant) GetSu() bool {
	if x != nil {
		return x.Su
	}
	return false
}

func (x *NodeGrant) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *NodeGrant) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

func (x *NodeGrant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NodeGrant) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

var File_models_model_node_grant_proto protoreflect.FileDescriptor

var file_models_model_node_grant_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x89, 0x02, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x73, 0x75, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_node_grant_proto_rawDescOnce sync.Once
	file_models_model_node_grant_proto_rawDescData = file_models_model_node_grant_proto_rawDesc
)

func file_models_model_node_grant_proto_rawDescGZIP() []byte {
	file_models_model_node_grant_proto_rawDescOnce.Do(func() {
		file_models_model_node_grant_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_node_grant_proto_rawDescData)
	})
	return file_models_model_node_grant_proto_rawDescData
}

var file_models_model_node_grant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_node_grant_proto_goTypes = []interface{}{
	(*NodeGrant)(nil), // 0: pb.NodeGrant
}
var file_models_model_node_grant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_node_grant_proto_init() }
func file_models_model_node_grant_proto_init() {
	if File_models_model_node_grant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_node_grant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGrant); i {
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
			RawDescriptor: file_models_model_node_grant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_node_grant_proto_goTypes,
		DependencyIndexes: file_models_model_node_grant_proto_depIdxs,
		MessageInfos:      file_models_model_node_grant_proto_msgTypes,
	}.Build()
	File_models_model_node_grant_proto = out.File
	file_models_model_node_grant_proto_rawDesc = nil
	file_models_model_node_grant_proto_goTypes = nil
	file_models_model_node_grant_proto_depIdxs = nil
}
