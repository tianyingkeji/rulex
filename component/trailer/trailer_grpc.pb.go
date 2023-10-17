// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: trailer.proto

package trailer

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

const (
	Trailer_Init_FullMethodName    = "/trailer.trailer/Init"
	Trailer_Start_FullMethodName   = "/trailer.trailer/Start"
	Trailer_Status_FullMethodName  = "/trailer.trailer/Status"
	Trailer_Service_FullMethodName = "/trailer.trailer/Service"
	Trailer_Query_FullMethodName   = "/trailer.trailer/Query"
	Trailer_Schema_FullMethodName  = "/trailer.trailer/Schema"
	Trailer_Stop_FullMethodName    = "/trailer.trailer/Stop"
)

// TrailerClient is the client API for Trailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrailerClient interface {
	// 初始化, 主要是为了传配置进去
	Init(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error)
	// 启动
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 获取状态
	Status(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// 服务调用
	Service(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
	// 数据查询
	Query(ctx context.Context, in *DataRowsRequest, opts ...grpc.CallOption) (*DataRowsResponse, error)
	// 数据模型
	Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResponse, error)
	// 停止
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type trailerClient struct {
	cc grpc.ClientConnInterface
}

func NewTrailerClient(cc grpc.ClientConnInterface) TrailerClient {
	return &trailerClient{cc}
}

func (c *trailerClient) Init(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Trailer_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Trailer_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Status(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Trailer_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Service(ctx context.Context, in *ServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, Trailer_Service_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Query(ctx context.Context, in *DataRowsRequest, opts ...grpc.CallOption) (*DataRowsResponse, error) {
	out := new(DataRowsResponse)
	err := c.cc.Invoke(ctx, Trailer_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Schema(ctx context.Context, in *SchemaRequest, opts ...grpc.CallOption) (*SchemaResponse, error) {
	out := new(SchemaResponse)
	err := c.cc.Invoke(ctx, Trailer_Schema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trailerClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Trailer_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrailerServer is the server API for Trailer service.
// All implementations must embed UnimplementedTrailerServer
// for forward compatibility
type TrailerServer interface {
	// 初始化, 主要是为了传配置进去
	Init(context.Context, *Config) (*Response, error)
	// 启动
	Start(context.Context, *Request) (*Response, error)
	// 获取状态
	Status(context.Context, *Request) (*Response, error)
	// 服务调用
	Service(context.Context, *ServiceRequest) (*ServiceResponse, error)
	// 数据查询
	Query(context.Context, *DataRowsRequest) (*DataRowsResponse, error)
	// 数据模型
	Schema(context.Context, *SchemaRequest) (*SchemaResponse, error)
	// 停止
	Stop(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedTrailerServer()
}

// UnimplementedTrailerServer must be embedded to have forward compatible implementations.
type UnimplementedTrailerServer struct {
}

func (UnimplementedTrailerServer) Init(context.Context, *Config) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedTrailerServer) Start(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedTrailerServer) Status(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTrailerServer) Service(context.Context, *ServiceRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Service not implemented")
}
func (UnimplementedTrailerServer) Query(context.Context, *DataRowsRequest) (*DataRowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedTrailerServer) Schema(context.Context, *SchemaRequest) (*SchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schema not implemented")
}
func (UnimplementedTrailerServer) Stop(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTrailerServer) mustEmbedUnimplementedTrailerServer() {}

// UnsafeTrailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrailerServer will
// result in compilation errors.
type UnsafeTrailerServer interface {
	mustEmbedUnimplementedTrailerServer()
}

func RegisterTrailerServer(s grpc.ServiceRegistrar, srv TrailerServer) {
	s.RegisterService(&Trailer_ServiceDesc, srv)
}

func _Trailer_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Init(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Start(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Status(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Service_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Service(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Service_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Service(ctx, req.(*ServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Query(ctx, req.(*DataRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Schema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Schema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Schema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Schema(ctx, req.(*SchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trailer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrailerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Trailer_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrailerServer).Stop(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Trailer_ServiceDesc is the grpc.ServiceDesc for Trailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Trailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trailer.trailer",
	HandlerType: (*TrailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Trailer_Init_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Trailer_Start_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Trailer_Status_Handler,
		},
		{
			MethodName: "Service",
			Handler:    _Trailer_Service_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Trailer_Query_Handler,
		},
		{
			MethodName: "Schema",
			Handler:    _Trailer_Schema_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Trailer_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trailer.proto",
}
