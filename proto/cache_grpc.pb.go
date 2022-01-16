// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// CacherClient is the client API for Cacher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacherClient interface {
	GetKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	SetKey(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*KeyPair, error)
}

type cacherClient struct {
	cc grpc.ClientConnInterface
}

func NewCacherClient(cc grpc.ClientConnInterface) CacherClient {
	return &cacherClient{cc}
}

func (c *cacherClient) GetKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/proto.Cacher/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) SetKey(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*KeyPair, error) {
	out := new(KeyPair)
	err := c.cc.Invoke(ctx, "/proto.Cacher/SetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacherServer is the server API for Cacher service.
// All implementations must embed UnimplementedCacherServer
// for forward compatibility
type CacherServer interface {
	GetKey(context.Context, *Key) (*Value, error)
	SetKey(context.Context, *KeyPair) (*KeyPair, error)
	mustEmbedUnimplementedCacherServer()
}

// UnimplementedCacherServer must be embedded to have forward compatible implementations.
type UnimplementedCacherServer struct {
}

func (UnimplementedCacherServer) GetKey(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedCacherServer) SetKey(context.Context, *KeyPair) (*KeyPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKey not implemented")
}
func (UnimplementedCacherServer) mustEmbedUnimplementedCacherServer() {}

// UnsafeCacherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacherServer will
// result in compilation errors.
type UnsafeCacherServer interface {
	mustEmbedUnimplementedCacherServer()
}

func RegisterCacherServer(s grpc.ServiceRegistrar, srv CacherServer) {
	s.RegisterService(&Cacher_ServiceDesc, srv)
}

func _Cacher_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cacher/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).GetKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_SetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).SetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cacher/SetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).SetKey(ctx, req.(*KeyPair))
	}
	return interceptor(ctx, in, info, handler)
}

// Cacher_ServiceDesc is the grpc.ServiceDesc for Cacher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cacher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cacher",
	HandlerType: (*CacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKey",
			Handler:    _Cacher_GetKey_Handler,
		},
		{
			MethodName: "SetKey",
			Handler:    _Cacher_SetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}