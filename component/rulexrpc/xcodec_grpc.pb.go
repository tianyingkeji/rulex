// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: xcodec.proto

package rulexrpc

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
	Codec_Decode_FullMethodName = "/rulexrpc.Codec/Decode"
	Codec_Encode_FullMethodName = "/rulexrpc.Codec/Encode"
)

// CodecClient is the client API for Codec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodecClient interface {
	// 编码
	Decode(ctx context.Context, in *CodecRequest, opts ...grpc.CallOption) (*CodecResponse, error)
	// 解码
	Encode(ctx context.Context, in *CodecRequest, opts ...grpc.CallOption) (*CodecResponse, error)
}

type codecClient struct {
	cc grpc.ClientConnInterface
}

func NewCodecClient(cc grpc.ClientConnInterface) CodecClient {
	return &codecClient{cc}
}

func (c *codecClient) Decode(ctx context.Context, in *CodecRequest, opts ...grpc.CallOption) (*CodecResponse, error) {
	out := new(CodecResponse)
	err := c.cc.Invoke(ctx, Codec_Decode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codecClient) Encode(ctx context.Context, in *CodecRequest, opts ...grpc.CallOption) (*CodecResponse, error) {
	out := new(CodecResponse)
	err := c.cc.Invoke(ctx, Codec_Encode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodecServer is the server API for Codec service.
// All implementations must embed UnimplementedCodecServer
// for forward compatibility
type CodecServer interface {
	// 编码
	Decode(context.Context, *CodecRequest) (*CodecResponse, error)
	// 解码
	Encode(context.Context, *CodecRequest) (*CodecResponse, error)
	mustEmbedUnimplementedCodecServer()
}

// UnimplementedCodecServer must be embedded to have forward compatible implementations.
type UnimplementedCodecServer struct {
}

func (UnimplementedCodecServer) Decode(context.Context, *CodecRequest) (*CodecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (UnimplementedCodecServer) Encode(context.Context, *CodecRequest) (*CodecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encode not implemented")
}
func (UnimplementedCodecServer) mustEmbedUnimplementedCodecServer() {}

// UnsafeCodecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodecServer will
// result in compilation errors.
type UnsafeCodecServer interface {
	mustEmbedUnimplementedCodecServer()
}

func RegisterCodecServer(s grpc.ServiceRegistrar, srv CodecServer) {
	s.RegisterService(&Codec_ServiceDesc, srv)
}

func _Codec_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodecServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Codec_Decode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodecServer).Decode(ctx, req.(*CodecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Codec_Encode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodecServer).Encode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Codec_Encode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodecServer).Encode(ctx, req.(*CodecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Codec_ServiceDesc is the grpc.ServiceDesc for Codec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Codec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rulexrpc.Codec",
	HandlerType: (*CodecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decode",
			Handler:    _Codec_Decode_Handler,
		},
		{
			MethodName: "Encode",
			Handler:    _Codec_Encode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xcodec.proto",
}
