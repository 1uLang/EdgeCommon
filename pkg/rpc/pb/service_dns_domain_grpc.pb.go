// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_dns_domain.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DNSDomainServiceClient is the client API for DNSDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSDomainServiceClient interface {
	// 创建域名
	CreateDNSDomain(ctx context.Context, in *CreateDNSDomainRequest, opts ...grpc.CallOption) (*CreateDNSDomainResponse, error)
	// 修改域名
	UpdateDNSDomain(ctx context.Context, in *UpdateDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除域名
	DeleteDNSDomain(ctx context.Context, in *DeleteDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 恢复删除的域名
	RecoverDNSDomain(ctx context.Context, in *RecoverDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询单个域名完整信息
	FindDNSDomain(ctx context.Context, in *FindDNSDomainRequest, opts ...grpc.CallOption) (*FindDNSDomainResponse, error)
	// 查询单个域名基础信息
	FindBasicDNSDomain(ctx context.Context, in *FindBasicDNSDomainRequest, opts ...grpc.CallOption) (*FindBasicDNSDomainResponse, error)
	// 计算服务商下的域名数量
	CountAllDNSDomainsWithDNSProviderId(ctx context.Context, in *CountAllDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出服务商下的所有域名
	FindAllDNSDomainsWithDNSProviderId(ctx context.Context, in *FindAllDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*FindAllDNSDomainsWithDNSProviderIdResponse, error)
	// 列出服务商下的所有域名基本信息
	FindAllBasicDNSDomainsWithDNSProviderId(ctx context.Context, in *FindAllBasicDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*FindAllBasicDNSDomainsWithDNSProviderIdResponse, error)
	// 列出服务商下的单页域名信息
	ListBasicDNSDomainsWithDNSProviderId(ctx context.Context, in *ListBasicDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*ListDNSDomainsWithDNSProviderIdResponse, error)
	// 同步域名解析
	SyncDNSDomainData(ctx context.Context, in *SyncDNSDomainDataRequest, opts ...grpc.CallOption) (*SyncDNSDomainDataResponse, error)
	// 查看支持的线路
	FindAllDNSDomainRoutes(ctx context.Context, in *FindAllDNSDomainRoutesRequest, opts ...grpc.CallOption) (*FindAllDNSDomainRoutesResponse, error)
	// 判断是否有域名可选
	ExistAvailableDomains(ctx context.Context, in *ExistAvailableDomainsRequest, opts ...grpc.CallOption) (*ExistAvailableDomainsResponse, error)
	// 检查域名是否在记录中
	ExistDNSDomainRecord(ctx context.Context, in *ExistDNSDomainRecordRequest, opts ...grpc.CallOption) (*ExistDNSDomainRecordResponse, error)
	// 从服务商同步域名
	SyncDNSDomainsFromProvider(ctx context.Context, in *SyncDNSDomainsFromProviderRequest, opts ...grpc.CallOption) (*SyncDNSDomainsFromProviderResponse, error)
}

type dNSDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSDomainServiceClient(cc grpc.ClientConnInterface) DNSDomainServiceClient {
	return &dNSDomainServiceClient{cc}
}

func (c *dNSDomainServiceClient) CreateDNSDomain(ctx context.Context, in *CreateDNSDomainRequest, opts ...grpc.CallOption) (*CreateDNSDomainResponse, error) {
	out := new(CreateDNSDomainResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/createDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) UpdateDNSDomain(ctx context.Context, in *UpdateDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/updateDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) DeleteDNSDomain(ctx context.Context, in *DeleteDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/deleteDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) RecoverDNSDomain(ctx context.Context, in *RecoverDNSDomainRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/recoverDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) FindDNSDomain(ctx context.Context, in *FindDNSDomainRequest, opts ...grpc.CallOption) (*FindDNSDomainResponse, error) {
	out := new(FindDNSDomainResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/findDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) FindBasicDNSDomain(ctx context.Context, in *FindBasicDNSDomainRequest, opts ...grpc.CallOption) (*FindBasicDNSDomainResponse, error) {
	out := new(FindBasicDNSDomainResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/findBasicDNSDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) CountAllDNSDomainsWithDNSProviderId(ctx context.Context, in *CountAllDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/countAllDNSDomainsWithDNSProviderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) FindAllDNSDomainsWithDNSProviderId(ctx context.Context, in *FindAllDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*FindAllDNSDomainsWithDNSProviderIdResponse, error) {
	out := new(FindAllDNSDomainsWithDNSProviderIdResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/findAllDNSDomainsWithDNSProviderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) FindAllBasicDNSDomainsWithDNSProviderId(ctx context.Context, in *FindAllBasicDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*FindAllBasicDNSDomainsWithDNSProviderIdResponse, error) {
	out := new(FindAllBasicDNSDomainsWithDNSProviderIdResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/findAllBasicDNSDomainsWithDNSProviderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) ListBasicDNSDomainsWithDNSProviderId(ctx context.Context, in *ListBasicDNSDomainsWithDNSProviderIdRequest, opts ...grpc.CallOption) (*ListDNSDomainsWithDNSProviderIdResponse, error) {
	out := new(ListDNSDomainsWithDNSProviderIdResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/listBasicDNSDomainsWithDNSProviderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) SyncDNSDomainData(ctx context.Context, in *SyncDNSDomainDataRequest, opts ...grpc.CallOption) (*SyncDNSDomainDataResponse, error) {
	out := new(SyncDNSDomainDataResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/syncDNSDomainData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) FindAllDNSDomainRoutes(ctx context.Context, in *FindAllDNSDomainRoutesRequest, opts ...grpc.CallOption) (*FindAllDNSDomainRoutesResponse, error) {
	out := new(FindAllDNSDomainRoutesResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/findAllDNSDomainRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) ExistAvailableDomains(ctx context.Context, in *ExistAvailableDomainsRequest, opts ...grpc.CallOption) (*ExistAvailableDomainsResponse, error) {
	out := new(ExistAvailableDomainsResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/existAvailableDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) ExistDNSDomainRecord(ctx context.Context, in *ExistDNSDomainRecordRequest, opts ...grpc.CallOption) (*ExistDNSDomainRecordResponse, error) {
	out := new(ExistDNSDomainRecordResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/existDNSDomainRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSDomainServiceClient) SyncDNSDomainsFromProvider(ctx context.Context, in *SyncDNSDomainsFromProviderRequest, opts ...grpc.CallOption) (*SyncDNSDomainsFromProviderResponse, error) {
	out := new(SyncDNSDomainsFromProviderResponse)
	err := c.cc.Invoke(ctx, "/pb.DNSDomainService/syncDNSDomainsFromProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSDomainServiceServer is the server API for DNSDomainService service.
// All implementations must embed UnimplementedDNSDomainServiceServer
// for forward compatibility
type DNSDomainServiceServer interface {
	// 创建域名
	CreateDNSDomain(context.Context, *CreateDNSDomainRequest) (*CreateDNSDomainResponse, error)
	// 修改域名
	UpdateDNSDomain(context.Context, *UpdateDNSDomainRequest) (*RPCSuccess, error)
	// 删除域名
	DeleteDNSDomain(context.Context, *DeleteDNSDomainRequest) (*RPCSuccess, error)
	// 恢复删除的域名
	RecoverDNSDomain(context.Context, *RecoverDNSDomainRequest) (*RPCSuccess, error)
	// 查询单个域名完整信息
	FindDNSDomain(context.Context, *FindDNSDomainRequest) (*FindDNSDomainResponse, error)
	// 查询单个域名基础信息
	FindBasicDNSDomain(context.Context, *FindBasicDNSDomainRequest) (*FindBasicDNSDomainResponse, error)
	// 计算服务商下的域名数量
	CountAllDNSDomainsWithDNSProviderId(context.Context, *CountAllDNSDomainsWithDNSProviderIdRequest) (*RPCCountResponse, error)
	// 列出服务商下的所有域名
	FindAllDNSDomainsWithDNSProviderId(context.Context, *FindAllDNSDomainsWithDNSProviderIdRequest) (*FindAllDNSDomainsWithDNSProviderIdResponse, error)
	// 列出服务商下的所有域名基本信息
	FindAllBasicDNSDomainsWithDNSProviderId(context.Context, *FindAllBasicDNSDomainsWithDNSProviderIdRequest) (*FindAllBasicDNSDomainsWithDNSProviderIdResponse, error)
	// 列出服务商下的单页域名信息
	ListBasicDNSDomainsWithDNSProviderId(context.Context, *ListBasicDNSDomainsWithDNSProviderIdRequest) (*ListDNSDomainsWithDNSProviderIdResponse, error)
	// 同步域名解析
	SyncDNSDomainData(context.Context, *SyncDNSDomainDataRequest) (*SyncDNSDomainDataResponse, error)
	// 查看支持的线路
	FindAllDNSDomainRoutes(context.Context, *FindAllDNSDomainRoutesRequest) (*FindAllDNSDomainRoutesResponse, error)
	// 判断是否有域名可选
	ExistAvailableDomains(context.Context, *ExistAvailableDomainsRequest) (*ExistAvailableDomainsResponse, error)
	// 检查域名是否在记录中
	ExistDNSDomainRecord(context.Context, *ExistDNSDomainRecordRequest) (*ExistDNSDomainRecordResponse, error)
	// 从服务商同步域名
	SyncDNSDomainsFromProvider(context.Context, *SyncDNSDomainsFromProviderRequest) (*SyncDNSDomainsFromProviderResponse, error)
}

// UnimplementedDNSDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDNSDomainServiceServer struct {
}

func (UnimplementedDNSDomainServiceServer) CreateDNSDomain(context.Context, *CreateDNSDomainRequest) (*CreateDNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) UpdateDNSDomain(context.Context, *UpdateDNSDomainRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) DeleteDNSDomain(context.Context, *DeleteDNSDomainRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) RecoverDNSDomain(context.Context, *RecoverDNSDomainRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) FindDNSDomain(context.Context, *FindDNSDomainRequest) (*FindDNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) FindBasicDNSDomain(context.Context, *FindBasicDNSDomainRequest) (*FindBasicDNSDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBasicDNSDomain not implemented")
}
func (UnimplementedDNSDomainServiceServer) CountAllDNSDomainsWithDNSProviderId(context.Context, *CountAllDNSDomainsWithDNSProviderIdRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllDNSDomainsWithDNSProviderId not implemented")
}
func (UnimplementedDNSDomainServiceServer) FindAllDNSDomainsWithDNSProviderId(context.Context, *FindAllDNSDomainsWithDNSProviderIdRequest) (*FindAllDNSDomainsWithDNSProviderIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDNSDomainsWithDNSProviderId not implemented")
}
func (UnimplementedDNSDomainServiceServer) FindAllBasicDNSDomainsWithDNSProviderId(context.Context, *FindAllBasicDNSDomainsWithDNSProviderIdRequest) (*FindAllBasicDNSDomainsWithDNSProviderIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllBasicDNSDomainsWithDNSProviderId not implemented")
}
func (UnimplementedDNSDomainServiceServer) ListBasicDNSDomainsWithDNSProviderId(context.Context, *ListBasicDNSDomainsWithDNSProviderIdRequest) (*ListDNSDomainsWithDNSProviderIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBasicDNSDomainsWithDNSProviderId not implemented")
}
func (UnimplementedDNSDomainServiceServer) SyncDNSDomainData(context.Context, *SyncDNSDomainDataRequest) (*SyncDNSDomainDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncDNSDomainData not implemented")
}
func (UnimplementedDNSDomainServiceServer) FindAllDNSDomainRoutes(context.Context, *FindAllDNSDomainRoutesRequest) (*FindAllDNSDomainRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDNSDomainRoutes not implemented")
}
func (UnimplementedDNSDomainServiceServer) ExistAvailableDomains(context.Context, *ExistAvailableDomainsRequest) (*ExistAvailableDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistAvailableDomains not implemented")
}
func (UnimplementedDNSDomainServiceServer) ExistDNSDomainRecord(context.Context, *ExistDNSDomainRecordRequest) (*ExistDNSDomainRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistDNSDomainRecord not implemented")
}
func (UnimplementedDNSDomainServiceServer) SyncDNSDomainsFromProvider(context.Context, *SyncDNSDomainsFromProviderRequest) (*SyncDNSDomainsFromProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncDNSDomainsFromProvider not implemented")
}
func RegisterDNSDomainServiceServer(s grpc.ServiceRegistrar, srv DNSDomainServiceServer) {
	s.RegisterService(&DNSDomainService_ServiceDesc, srv)
}

func _DNSDomainService_CreateDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).CreateDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/createDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).CreateDNSDomain(ctx, req.(*CreateDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_UpdateDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).UpdateDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/updateDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).UpdateDNSDomain(ctx, req.(*UpdateDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_DeleteDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).DeleteDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/deleteDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).DeleteDNSDomain(ctx, req.(*DeleteDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_RecoverDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).RecoverDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/recoverDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).RecoverDNSDomain(ctx, req.(*RecoverDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_FindDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).FindDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/findDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).FindDNSDomain(ctx, req.(*FindDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_FindBasicDNSDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBasicDNSDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).FindBasicDNSDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/findBasicDNSDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).FindBasicDNSDomain(ctx, req.(*FindBasicDNSDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_CountAllDNSDomainsWithDNSProviderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllDNSDomainsWithDNSProviderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).CountAllDNSDomainsWithDNSProviderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/countAllDNSDomainsWithDNSProviderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).CountAllDNSDomainsWithDNSProviderId(ctx, req.(*CountAllDNSDomainsWithDNSProviderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_FindAllDNSDomainsWithDNSProviderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDNSDomainsWithDNSProviderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).FindAllDNSDomainsWithDNSProviderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/findAllDNSDomainsWithDNSProviderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).FindAllDNSDomainsWithDNSProviderId(ctx, req.(*FindAllDNSDomainsWithDNSProviderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_FindAllBasicDNSDomainsWithDNSProviderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllBasicDNSDomainsWithDNSProviderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).FindAllBasicDNSDomainsWithDNSProviderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/findAllBasicDNSDomainsWithDNSProviderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).FindAllBasicDNSDomainsWithDNSProviderId(ctx, req.(*FindAllBasicDNSDomainsWithDNSProviderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_ListBasicDNSDomainsWithDNSProviderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBasicDNSDomainsWithDNSProviderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).ListBasicDNSDomainsWithDNSProviderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/listBasicDNSDomainsWithDNSProviderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).ListBasicDNSDomainsWithDNSProviderId(ctx, req.(*ListBasicDNSDomainsWithDNSProviderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_SyncDNSDomainData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncDNSDomainDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).SyncDNSDomainData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/syncDNSDomainData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).SyncDNSDomainData(ctx, req.(*SyncDNSDomainDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_FindAllDNSDomainRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDNSDomainRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).FindAllDNSDomainRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/findAllDNSDomainRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).FindAllDNSDomainRoutes(ctx, req.(*FindAllDNSDomainRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_ExistAvailableDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistAvailableDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).ExistAvailableDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/existAvailableDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).ExistAvailableDomains(ctx, req.(*ExistAvailableDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_ExistDNSDomainRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistDNSDomainRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).ExistDNSDomainRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/existDNSDomainRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).ExistDNSDomainRecord(ctx, req.(*ExistDNSDomainRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSDomainService_SyncDNSDomainsFromProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncDNSDomainsFromProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSDomainServiceServer).SyncDNSDomainsFromProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DNSDomainService/syncDNSDomainsFromProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSDomainServiceServer).SyncDNSDomainsFromProvider(ctx, req.(*SyncDNSDomainsFromProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSDomainService_ServiceDesc is the grpc.ServiceDesc for DNSDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNSDomainService",
	HandlerType: (*DNSDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createDNSDomain",
			Handler:    _DNSDomainService_CreateDNSDomain_Handler,
		},
		{
			MethodName: "updateDNSDomain",
			Handler:    _DNSDomainService_UpdateDNSDomain_Handler,
		},
		{
			MethodName: "deleteDNSDomain",
			Handler:    _DNSDomainService_DeleteDNSDomain_Handler,
		},
		{
			MethodName: "recoverDNSDomain",
			Handler:    _DNSDomainService_RecoverDNSDomain_Handler,
		},
		{
			MethodName: "findDNSDomain",
			Handler:    _DNSDomainService_FindDNSDomain_Handler,
		},
		{
			MethodName: "findBasicDNSDomain",
			Handler:    _DNSDomainService_FindBasicDNSDomain_Handler,
		},
		{
			MethodName: "countAllDNSDomainsWithDNSProviderId",
			Handler:    _DNSDomainService_CountAllDNSDomainsWithDNSProviderId_Handler,
		},
		{
			MethodName: "findAllDNSDomainsWithDNSProviderId",
			Handler:    _DNSDomainService_FindAllDNSDomainsWithDNSProviderId_Handler,
		},
		{
			MethodName: "findAllBasicDNSDomainsWithDNSProviderId",
			Handler:    _DNSDomainService_FindAllBasicDNSDomainsWithDNSProviderId_Handler,
		},
		{
			MethodName: "listBasicDNSDomainsWithDNSProviderId",
			Handler:    _DNSDomainService_ListBasicDNSDomainsWithDNSProviderId_Handler,
		},
		{
			MethodName: "syncDNSDomainData",
			Handler:    _DNSDomainService_SyncDNSDomainData_Handler,
		},
		{
			MethodName: "findAllDNSDomainRoutes",
			Handler:    _DNSDomainService_FindAllDNSDomainRoutes_Handler,
		},
		{
			MethodName: "existAvailableDomains",
			Handler:    _DNSDomainService_ExistAvailableDomains_Handler,
		},
		{
			MethodName: "existDNSDomainRecord",
			Handler:    _DNSDomainService_ExistDNSDomainRecord_Handler,
		},
		{
			MethodName: "syncDNSDomainsFromProvider",
			Handler:    _DNSDomainService_SyncDNSDomainsFromProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dns_domain.proto",
}