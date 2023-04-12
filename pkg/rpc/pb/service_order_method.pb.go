// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: service_order_method.proto

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

// 创建支付方式
type CreateOrderMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateOrderMethodRequest) Reset() {
	*x = CreateOrderMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderMethodRequest) ProtoMessage() {}

func (x *CreateOrderMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderMethodRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderMethodRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderMethodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOrderMethodRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateOrderMethodRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateOrderMethodRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateOrderMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethodId int64 `protobuf:"varint,1,opt,name=orderMethodId,proto3" json:"orderMethodId,omitempty"`
}

func (x *CreateOrderMethodResponse) Reset() {
	*x = CreateOrderMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderMethodResponse) ProtoMessage() {}

func (x *CreateOrderMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderMethodResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderMethodResponse) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderMethodResponse) GetOrderMethodId() int64 {
	if x != nil {
		return x.OrderMethodId
	}
	return 0
}

// 修改支付方式
type UpdateOrderMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethodId int64  `protobuf:"varint,1,opt,name=orderMethodId,proto3" json:"orderMethodId,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code          string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url           string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	IsOn          bool   `protobuf:"varint,6,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateOrderMethodRequest) Reset() {
	*x = UpdateOrderMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderMethodRequest) ProtoMessage() {}

func (x *UpdateOrderMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderMethodRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderMethodRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrderMethodRequest) GetOrderMethodId() int64 {
	if x != nil {
		return x.OrderMethodId
	}
	return 0
}

func (x *UpdateOrderMethodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateOrderMethodRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateOrderMethodRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateOrderMethodRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateOrderMethodRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 删除支付方式
type DeleteOrderMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethodId int64 `protobuf:"varint,1,opt,name=orderMethodId,proto3" json:"orderMethodId,omitempty"`
}

func (x *DeleteOrderMethodRequest) Reset() {
	*x = DeleteOrderMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderMethodRequest) ProtoMessage() {}

func (x *DeleteOrderMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderMethodRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderMethodRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteOrderMethodRequest) GetOrderMethodId() int64 {
	if x != nil {
		return x.OrderMethodId
	}
	return 0
}

// 查找单个支付方式
type FindEnabledOrderMethodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethodId int64 `protobuf:"varint,1,opt,name=orderMethodId,proto3" json:"orderMethodId,omitempty"`
}

func (x *FindEnabledOrderMethodRequest) Reset() {
	*x = FindEnabledOrderMethodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOrderMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOrderMethodRequest) ProtoMessage() {}

func (x *FindEnabledOrderMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOrderMethodRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledOrderMethodRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledOrderMethodRequest) GetOrderMethodId() int64 {
	if x != nil {
		return x.OrderMethodId
	}
	return 0
}

type FindEnabledOrderMethodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethod *OrderMethod `protobuf:"bytes,1,opt,name=orderMethod,proto3" json:"orderMethod,omitempty"`
}

func (x *FindEnabledOrderMethodResponse) Reset() {
	*x = FindEnabledOrderMethodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOrderMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOrderMethodResponse) ProtoMessage() {}

func (x *FindEnabledOrderMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOrderMethodResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledOrderMethodResponse) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledOrderMethodResponse) GetOrderMethod() *OrderMethod {
	if x != nil {
		return x.OrderMethod
	}
	return nil
}

// 根据代号查找支付方式
type FindEnabledOrderMethodWithCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FindEnabledOrderMethodWithCodeRequest) Reset() {
	*x = FindEnabledOrderMethodWithCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOrderMethodWithCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOrderMethodWithCodeRequest) ProtoMessage() {}

func (x *FindEnabledOrderMethodWithCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOrderMethodWithCodeRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledOrderMethodWithCodeRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledOrderMethodWithCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type FindEnabledOrderMethodWithCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethod *OrderMethod `protobuf:"bytes,1,opt,name=orderMethod,proto3" json:"orderMethod,omitempty"`
}

func (x *FindEnabledOrderMethodWithCodeResponse) Reset() {
	*x = FindEnabledOrderMethodWithCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledOrderMethodWithCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledOrderMethodWithCodeResponse) ProtoMessage() {}

func (x *FindEnabledOrderMethodWithCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledOrderMethodWithCodeResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledOrderMethodWithCodeResponse) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledOrderMethodWithCodeResponse) GetOrderMethod() *OrderMethod {
	if x != nil {
		return x.OrderMethod
	}
	return nil
}

// 查找所有支付方式
type FindAllEnabledOrderMethodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllEnabledOrderMethodsRequest) Reset() {
	*x = FindAllEnabledOrderMethodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledOrderMethodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledOrderMethodsRequest) ProtoMessage() {}

func (x *FindAllEnabledOrderMethodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledOrderMethodsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledOrderMethodsRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{8}
}

type FindAllEnabledOrderMethodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethods []*OrderMethod `protobuf:"bytes,1,rep,name=orderMethods,proto3" json:"orderMethods,omitempty"`
}

func (x *FindAllEnabledOrderMethodsResponse) Reset() {
	*x = FindAllEnabledOrderMethodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledOrderMethodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledOrderMethodsResponse) ProtoMessage() {}

func (x *FindAllEnabledOrderMethodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledOrderMethodsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledOrderMethodsResponse) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{9}
}

func (x *FindAllEnabledOrderMethodsResponse) GetOrderMethods() []*OrderMethod {
	if x != nil {
		return x.OrderMethods
	}
	return nil
}

// 查找所有已启用的支付方式
type FindAllAvailableOrderMethodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllAvailableOrderMethodsRequest) Reset() {
	*x = FindAllAvailableOrderMethodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableOrderMethodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableOrderMethodsRequest) ProtoMessage() {}

func (x *FindAllAvailableOrderMethodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableOrderMethodsRequest.ProtoReflect.Descriptor instead.
func (*FindAllAvailableOrderMethodsRequest) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{10}
}

type FindAllAvailableOrderMethodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMethods []*OrderMethod `protobuf:"bytes,1,rep,name=orderMethods,proto3" json:"orderMethods,omitempty"`
}

func (x *FindAllAvailableOrderMethodsResponse) Reset() {
	*x = FindAllAvailableOrderMethodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_method_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllAvailableOrderMethodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllAvailableOrderMethodsResponse) ProtoMessage() {}

func (x *FindAllAvailableOrderMethodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_method_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllAvailableOrderMethodsResponse.ProtoReflect.Descriptor instead.
func (*FindAllAvailableOrderMethodsResponse) Descriptor() ([]byte, []int) {
	return file_service_order_method_proto_rawDescGZIP(), []int{11}
}

func (x *FindAllAvailableOrderMethodsResponse) GetOrderMethods() []*OrderMethod {
	if x != nil {
		return x.OrderMethods
	}
	return nil
}

var File_service_order_method_proto protoreflect.FileDescriptor

var file_service_order_method_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x41, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x1d,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x3b, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5b, 0x0a, 0x26, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x22, 0x23, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x22, 0x25, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x24, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x32, 0xa6, 0x05, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x1e, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6b, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x25, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x1c,
	0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x27, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_order_method_proto_rawDescOnce sync.Once
	file_service_order_method_proto_rawDescData = file_service_order_method_proto_rawDesc
)

func file_service_order_method_proto_rawDescGZIP() []byte {
	file_service_order_method_proto_rawDescOnce.Do(func() {
		file_service_order_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_order_method_proto_rawDescData)
	})
	return file_service_order_method_proto_rawDescData
}

var file_service_order_method_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_service_order_method_proto_goTypes = []interface{}{
	(*CreateOrderMethodRequest)(nil),               // 0: pb.CreateOrderMethodRequest
	(*CreateOrderMethodResponse)(nil),              // 1: pb.CreateOrderMethodResponse
	(*UpdateOrderMethodRequest)(nil),               // 2: pb.UpdateOrderMethodRequest
	(*DeleteOrderMethodRequest)(nil),               // 3: pb.DeleteOrderMethodRequest
	(*FindEnabledOrderMethodRequest)(nil),          // 4: pb.FindEnabledOrderMethodRequest
	(*FindEnabledOrderMethodResponse)(nil),         // 5: pb.FindEnabledOrderMethodResponse
	(*FindEnabledOrderMethodWithCodeRequest)(nil),  // 6: pb.FindEnabledOrderMethodWithCodeRequest
	(*FindEnabledOrderMethodWithCodeResponse)(nil), // 7: pb.FindEnabledOrderMethodWithCodeResponse
	(*FindAllEnabledOrderMethodsRequest)(nil),      // 8: pb.FindAllEnabledOrderMethodsRequest
	(*FindAllEnabledOrderMethodsResponse)(nil),     // 9: pb.FindAllEnabledOrderMethodsResponse
	(*FindAllAvailableOrderMethodsRequest)(nil),    // 10: pb.FindAllAvailableOrderMethodsRequest
	(*FindAllAvailableOrderMethodsResponse)(nil),   // 11: pb.FindAllAvailableOrderMethodsResponse
	(*OrderMethod)(nil),                            // 12: pb.OrderMethod
	(*RPCSuccess)(nil),                             // 13: pb.RPCSuccess
}
var file_service_order_method_proto_depIdxs = []int32{
	12, // 0: pb.FindEnabledOrderMethodResponse.orderMethod:type_name -> pb.OrderMethod
	12, // 1: pb.FindEnabledOrderMethodWithCodeResponse.orderMethod:type_name -> pb.OrderMethod
	12, // 2: pb.FindAllEnabledOrderMethodsResponse.orderMethods:type_name -> pb.OrderMethod
	12, // 3: pb.FindAllAvailableOrderMethodsResponse.orderMethods:type_name -> pb.OrderMethod
	0,  // 4: pb.OrderMethodService.createOrderMethod:input_type -> pb.CreateOrderMethodRequest
	2,  // 5: pb.OrderMethodService.updateOrderMethod:input_type -> pb.UpdateOrderMethodRequest
	3,  // 6: pb.OrderMethodService.deleteOrderMethod:input_type -> pb.DeleteOrderMethodRequest
	4,  // 7: pb.OrderMethodService.findEnabledOrderMethod:input_type -> pb.FindEnabledOrderMethodRequest
	6,  // 8: pb.OrderMethodService.findEnabledOrderMethodWithCode:input_type -> pb.FindEnabledOrderMethodWithCodeRequest
	8,  // 9: pb.OrderMethodService.findAllEnabledOrderMethods:input_type -> pb.FindAllEnabledOrderMethodsRequest
	10, // 10: pb.OrderMethodService.findAllAvailableOrderMethods:input_type -> pb.FindAllAvailableOrderMethodsRequest
	1,  // 11: pb.OrderMethodService.createOrderMethod:output_type -> pb.CreateOrderMethodResponse
	13, // 12: pb.OrderMethodService.updateOrderMethod:output_type -> pb.RPCSuccess
	13, // 13: pb.OrderMethodService.deleteOrderMethod:output_type -> pb.RPCSuccess
	5,  // 14: pb.OrderMethodService.findEnabledOrderMethod:output_type -> pb.FindEnabledOrderMethodResponse
	7,  // 15: pb.OrderMethodService.findEnabledOrderMethodWithCode:output_type -> pb.FindEnabledOrderMethodWithCodeResponse
	9,  // 16: pb.OrderMethodService.findAllEnabledOrderMethods:output_type -> pb.FindAllEnabledOrderMethodsResponse
	11, // 17: pb.OrderMethodService.findAllAvailableOrderMethods:output_type -> pb.FindAllAvailableOrderMethodsResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_service_order_method_proto_init() }
func file_service_order_method_proto_init() {
	if File_service_order_method_proto != nil {
		return
	}
	file_models_model_order_method_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_order_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderMethodRequest); i {
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
		file_service_order_method_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderMethodResponse); i {
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
		file_service_order_method_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderMethodRequest); i {
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
		file_service_order_method_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderMethodRequest); i {
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
		file_service_order_method_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOrderMethodRequest); i {
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
		file_service_order_method_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOrderMethodResponse); i {
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
		file_service_order_method_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOrderMethodWithCodeRequest); i {
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
		file_service_order_method_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledOrderMethodWithCodeResponse); i {
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
		file_service_order_method_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledOrderMethodsRequest); i {
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
		file_service_order_method_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledOrderMethodsResponse); i {
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
		file_service_order_method_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableOrderMethodsRequest); i {
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
		file_service_order_method_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllAvailableOrderMethodsResponse); i {
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
			RawDescriptor: file_service_order_method_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_order_method_proto_goTypes,
		DependencyIndexes: file_service_order_method_proto_depIdxs,
		MessageInfos:      file_service_order_method_proto_msgTypes,
	}.Build()
	File_service_order_method_proto = out.File
	file_service_order_method_proto_rawDesc = nil
	file_service_order_method_proto_goTypes = nil
	file_service_order_method_proto_depIdxs = nil
}