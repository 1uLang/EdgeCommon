// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: service_db.proto

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

// 获取所有表信息
type FindAllDBTablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllDBTablesRequest) Reset() {
	*x = FindAllDBTablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllDBTablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllDBTablesRequest) ProtoMessage() {}

func (x *FindAllDBTablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllDBTablesRequest.ProtoReflect.Descriptor instead.
func (*FindAllDBTablesRequest) Descriptor() ([]byte, []int) {
	return file_service_db_proto_rawDescGZIP(), []int{0}
}

type FindAllDBTablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbTables []*DBTable `protobuf:"bytes,1,rep,name=dbTables,proto3" json:"dbTables,omitempty"`
}

func (x *FindAllDBTablesResponse) Reset() {
	*x = FindAllDBTablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllDBTablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllDBTablesResponse) ProtoMessage() {}

func (x *FindAllDBTablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllDBTablesResponse.ProtoReflect.Descriptor instead.
func (*FindAllDBTablesResponse) Descriptor() ([]byte, []int) {
	return file_service_db_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllDBTablesResponse) GetDbTables() []*DBTable {
	if x != nil {
		return x.DbTables
	}
	return nil
}

// 删除表
type DeleteDBTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbTable string `protobuf:"bytes,1,opt,name=dbTable,proto3" json:"dbTable,omitempty"`
}

func (x *DeleteDBTableRequest) Reset() {
	*x = DeleteDBTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDBTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDBTableRequest) ProtoMessage() {}

func (x *DeleteDBTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDBTableRequest.ProtoReflect.Descriptor instead.
func (*DeleteDBTableRequest) Descriptor() ([]byte, []int) {
	return file_service_db_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDBTableRequest) GetDbTable() string {
	if x != nil {
		return x.DbTable
	}
	return ""
}

// 清空表
type TruncateDBTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbTable string `protobuf:"bytes,1,opt,name=dbTable,proto3" json:"dbTable,omitempty"`
}

func (x *TruncateDBTableRequest) Reset() {
	*x = TruncateDBTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TruncateDBTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TruncateDBTableRequest) ProtoMessage() {}

func (x *TruncateDBTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TruncateDBTableRequest.ProtoReflect.Descriptor instead.
func (*TruncateDBTableRequest) Descriptor() ([]byte, []int) {
	return file_service_db_proto_rawDescGZIP(), []int{3}
}

func (x *TruncateDBTableRequest) GetDbTable() string {
	if x != nil {
		return x.DbTable
	}
	return ""
}

var File_service_db_proto protoreflect.FileDescriptor

var file_service_db_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x64, 0x62, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18,
	0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x64, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x08, 0x64, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x32,
	0x0a, 0x16, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x62, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x32, 0xd1, 0x01, 0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x44, 0x42, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x74, 0x72, 0x75, 0x6e, 0x63,
	0x61, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_db_proto_rawDescOnce sync.Once
	file_service_db_proto_rawDescData = file_service_db_proto_rawDesc
)

func file_service_db_proto_rawDescGZIP() []byte {
	file_service_db_proto_rawDescOnce.Do(func() {
		file_service_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_db_proto_rawDescData)
	})
	return file_service_db_proto_rawDescData
}

var file_service_db_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_db_proto_goTypes = []interface{}{
	(*FindAllDBTablesRequest)(nil),  // 0: pb.FindAllDBTablesRequest
	(*FindAllDBTablesResponse)(nil), // 1: pb.FindAllDBTablesResponse
	(*DeleteDBTableRequest)(nil),    // 2: pb.DeleteDBTableRequest
	(*TruncateDBTableRequest)(nil),  // 3: pb.TruncateDBTableRequest
	(*DBTable)(nil),                 // 4: pb.DBTable
	(*RPCSuccess)(nil),              // 5: pb.RPCSuccess
}
var file_service_db_proto_depIdxs = []int32{
	4, // 0: pb.FindAllDBTablesResponse.dbTables:type_name -> pb.DBTable
	0, // 1: pb.DBService.findAllDBTables:input_type -> pb.FindAllDBTablesRequest
	2, // 2: pb.DBService.deleteDBTable:input_type -> pb.DeleteDBTableRequest
	3, // 3: pb.DBService.truncateDBTable:input_type -> pb.TruncateDBTableRequest
	1, // 4: pb.DBService.findAllDBTables:output_type -> pb.FindAllDBTablesResponse
	5, // 5: pb.DBService.deleteDBTable:output_type -> pb.RPCSuccess
	5, // 6: pb.DBService.truncateDBTable:output_type -> pb.RPCSuccess
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_db_proto_init() }
func file_service_db_proto_init() {
	if File_service_db_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_db_table_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllDBTablesRequest); i {
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
		file_service_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllDBTablesResponse); i {
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
		file_service_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDBTableRequest); i {
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
		file_service_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TruncateDBTableRequest); i {
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
			RawDescriptor: file_service_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_db_proto_goTypes,
		DependencyIndexes: file_service_db_proto_depIdxs,
		MessageInfos:      file_service_db_proto_msgTypes,
	}.Build()
	File_service_db_proto = out.File
	file_service_db_proto_rawDesc = nil
	file_service_db_proto_goTypes = nil
	file_service_db_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBServiceClient interface {
	// 获取所有表信息
	FindAllDBTables(ctx context.Context, in *FindAllDBTablesRequest, opts ...grpc.CallOption) (*FindAllDBTablesResponse, error)
	// 删除表
	DeleteDBTable(ctx context.Context, in *DeleteDBTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 清空表
	TruncateDBTable(ctx context.Context, in *TruncateDBTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type dBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBServiceClient(cc grpc.ClientConnInterface) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) FindAllDBTables(ctx context.Context, in *FindAllDBTablesRequest, opts ...grpc.CallOption) (*FindAllDBTablesResponse, error) {
	out := new(FindAllDBTablesResponse)
	err := c.cc.Invoke(ctx, "/pb.DBService/findAllDBTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) DeleteDBTable(ctx context.Context, in *DeleteDBTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DBService/deleteDBTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) TruncateDBTable(ctx context.Context, in *TruncateDBTableRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DBService/truncateDBTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
type DBServiceServer interface {
	// 获取所有表信息
	FindAllDBTables(context.Context, *FindAllDBTablesRequest) (*FindAllDBTablesResponse, error)
	// 删除表
	DeleteDBTable(context.Context, *DeleteDBTableRequest) (*RPCSuccess, error)
	// 清空表
	TruncateDBTable(context.Context, *TruncateDBTableRequest) (*RPCSuccess, error)
}

// UnimplementedDBServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (*UnimplementedDBServiceServer) FindAllDBTables(context.Context, *FindAllDBTablesRequest) (*FindAllDBTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDBTables not implemented")
}
func (*UnimplementedDBServiceServer) DeleteDBTable(context.Context, *DeleteDBTableRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDBTable not implemented")
}
func (*UnimplementedDBServiceServer) TruncateDBTable(context.Context, *TruncateDBTableRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TruncateDBTable not implemented")
}

func RegisterDBServiceServer(s *grpc.Server, srv DBServiceServer) {
	s.RegisterService(&_DBService_serviceDesc, srv)
}

func _DBService_FindAllDBTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDBTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).FindAllDBTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DBService/FindAllDBTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).FindAllDBTables(ctx, req.(*FindAllDBTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_DeleteDBTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDBTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).DeleteDBTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DBService/DeleteDBTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).DeleteDBTable(ctx, req.(*DeleteDBTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_TruncateDBTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateDBTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).TruncateDBTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DBService/TruncateDBTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).TruncateDBTable(ctx, req.(*TruncateDBTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllDBTables",
			Handler:    _DBService_FindAllDBTables_Handler,
		},
		{
			MethodName: "deleteDBTable",
			Handler:    _DBService_DeleteDBTable_Handler,
		},
		{
			MethodName: "truncateDBTable",
			Handler:    _DBService_TruncateDBTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_db.proto",
}
