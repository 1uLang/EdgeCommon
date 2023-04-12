// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: service_region_city.proto

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

// RegionCityServiceClient is the client API for RegionCityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionCityServiceClient interface {
	// Deprecated: Do not use.
	// 查找所有城市
	FindAllEnabledRegionCities(ctx context.Context, in *FindAllEnabledRegionCitiesRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionCitiesResponse, error)
	// Deprecated: Do not use.
	// 查找单个城市信息
	FindEnabledRegionCity(ctx context.Context, in *FindEnabledRegionCityRequest, opts ...grpc.CallOption) (*FindEnabledRegionCityResponse, error)
	// 查找所有城市
	FindAllRegionCities(ctx context.Context, in *FindAllRegionCitiesRequest, opts ...grpc.CallOption) (*FindAllRegionCitiesResponse, error)
	// 查找某个省份的所有城市
	FindAllRegionCitiesWithRegionProvinceId(ctx context.Context, in *FindAllRegionCitiesWithRegionProvinceIdRequest, opts ...grpc.CallOption) (*FindAllRegionCitiesWithRegionProvinceIdResponse, error)
	// 查找单个城市信息
	FindRegionCity(ctx context.Context, in *FindRegionCityRequest, opts ...grpc.CallOption) (*FindRegionCityResponse, error)
	// 修改城市定制信息
	UpdateRegionCityCustom(ctx context.Context, in *UpdateRegionCityCustomRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type regionCityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionCityServiceClient(cc grpc.ClientConnInterface) RegionCityServiceClient {
	return &regionCityServiceClient{cc}
}

// Deprecated: Do not use.
func (c *regionCityServiceClient) FindAllEnabledRegionCities(ctx context.Context, in *FindAllEnabledRegionCitiesRequest, opts ...grpc.CallOption) (*FindAllEnabledRegionCitiesResponse, error) {
	out := new(FindAllEnabledRegionCitiesResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/findAllEnabledRegionCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *regionCityServiceClient) FindEnabledRegionCity(ctx context.Context, in *FindEnabledRegionCityRequest, opts ...grpc.CallOption) (*FindEnabledRegionCityResponse, error) {
	out := new(FindEnabledRegionCityResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/findEnabledRegionCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionCityServiceClient) FindAllRegionCities(ctx context.Context, in *FindAllRegionCitiesRequest, opts ...grpc.CallOption) (*FindAllRegionCitiesResponse, error) {
	out := new(FindAllRegionCitiesResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/findAllRegionCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionCityServiceClient) FindAllRegionCitiesWithRegionProvinceId(ctx context.Context, in *FindAllRegionCitiesWithRegionProvinceIdRequest, opts ...grpc.CallOption) (*FindAllRegionCitiesWithRegionProvinceIdResponse, error) {
	out := new(FindAllRegionCitiesWithRegionProvinceIdResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/findAllRegionCitiesWithRegionProvinceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionCityServiceClient) FindRegionCity(ctx context.Context, in *FindRegionCityRequest, opts ...grpc.CallOption) (*FindRegionCityResponse, error) {
	out := new(FindRegionCityResponse)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/findRegionCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionCityServiceClient) UpdateRegionCityCustom(ctx context.Context, in *UpdateRegionCityCustomRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.RegionCityService/updateRegionCityCustom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionCityServiceServer is the server API for RegionCityService service.
// All implementations must embed UnimplementedRegionCityServiceServer
// for forward compatibility
type RegionCityServiceServer interface {
	// Deprecated: Do not use.
	// 查找所有城市
	FindAllEnabledRegionCities(context.Context, *FindAllEnabledRegionCitiesRequest) (*FindAllEnabledRegionCitiesResponse, error)
	// Deprecated: Do not use.
	// 查找单个城市信息
	FindEnabledRegionCity(context.Context, *FindEnabledRegionCityRequest) (*FindEnabledRegionCityResponse, error)
	// 查找所有城市
	FindAllRegionCities(context.Context, *FindAllRegionCitiesRequest) (*FindAllRegionCitiesResponse, error)
	// 查找某个省份的所有城市
	FindAllRegionCitiesWithRegionProvinceId(context.Context, *FindAllRegionCitiesWithRegionProvinceIdRequest) (*FindAllRegionCitiesWithRegionProvinceIdResponse, error)
	// 查找单个城市信息
	FindRegionCity(context.Context, *FindRegionCityRequest) (*FindRegionCityResponse, error)
	// 修改城市定制信息
	UpdateRegionCityCustom(context.Context, *UpdateRegionCityCustomRequest) (*RPCSuccess, error)
}

// UnimplementedRegionCityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegionCityServiceServer struct {
}

func (UnimplementedRegionCityServiceServer) FindAllEnabledRegionCities(context.Context, *FindAllEnabledRegionCitiesRequest) (*FindAllEnabledRegionCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledRegionCities not implemented")
}
func (UnimplementedRegionCityServiceServer) FindEnabledRegionCity(context.Context, *FindEnabledRegionCityRequest) (*FindEnabledRegionCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledRegionCity not implemented")
}
func (UnimplementedRegionCityServiceServer) FindAllRegionCities(context.Context, *FindAllRegionCitiesRequest) (*FindAllRegionCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllRegionCities not implemented")
}
func (UnimplementedRegionCityServiceServer) FindAllRegionCitiesWithRegionProvinceId(context.Context, *FindAllRegionCitiesWithRegionProvinceIdRequest) (*FindAllRegionCitiesWithRegionProvinceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllRegionCitiesWithRegionProvinceId not implemented")
}
func (UnimplementedRegionCityServiceServer) FindRegionCity(context.Context, *FindRegionCityRequest) (*FindRegionCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRegionCity not implemented")
}
func (UnimplementedRegionCityServiceServer) UpdateRegionCityCustom(context.Context, *UpdateRegionCityCustomRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegionCityCustom not implemented")
}
func RegisterRegionCityServiceServer(s grpc.ServiceRegistrar, srv RegionCityServiceServer) {
	s.RegisterService(&RegionCityService_ServiceDesc, srv)
}

func _RegionCityService_FindAllEnabledRegionCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledRegionCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).FindAllEnabledRegionCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/findAllEnabledRegionCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).FindAllEnabledRegionCities(ctx, req.(*FindAllEnabledRegionCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionCityService_FindEnabledRegionCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledRegionCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).FindEnabledRegionCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/findEnabledRegionCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).FindEnabledRegionCity(ctx, req.(*FindEnabledRegionCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionCityService_FindAllRegionCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRegionCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).FindAllRegionCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/findAllRegionCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).FindAllRegionCities(ctx, req.(*FindAllRegionCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionCityService_FindAllRegionCitiesWithRegionProvinceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRegionCitiesWithRegionProvinceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).FindAllRegionCitiesWithRegionProvinceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/findAllRegionCitiesWithRegionProvinceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).FindAllRegionCitiesWithRegionProvinceId(ctx, req.(*FindAllRegionCitiesWithRegionProvinceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionCityService_FindRegionCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRegionCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).FindRegionCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/findRegionCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).FindRegionCity(ctx, req.(*FindRegionCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionCityService_UpdateRegionCityCustom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionCityCustomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionCityServiceServer).UpdateRegionCityCustom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RegionCityService/updateRegionCityCustom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionCityServiceServer).UpdateRegionCityCustom(ctx, req.(*UpdateRegionCityCustomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionCityService_ServiceDesc is the grpc.ServiceDesc for RegionCityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionCityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RegionCityService",
	HandlerType: (*RegionCityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledRegionCities",
			Handler:    _RegionCityService_FindAllEnabledRegionCities_Handler,
		},
		{
			MethodName: "findEnabledRegionCity",
			Handler:    _RegionCityService_FindEnabledRegionCity_Handler,
		},
		{
			MethodName: "findAllRegionCities",
			Handler:    _RegionCityService_FindAllRegionCities_Handler,
		},
		{
			MethodName: "findAllRegionCitiesWithRegionProvinceId",
			Handler:    _RegionCityService_FindAllRegionCitiesWithRegionProvinceId_Handler,
		},
		{
			MethodName: "findRegionCity",
			Handler:    _RegionCityService_FindRegionCity_Handler,
		},
		{
			MethodName: "updateRegionCityCustom",
			Handler:    _RegionCityService_UpdateRegionCityCustom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_region_city.proto",
}
