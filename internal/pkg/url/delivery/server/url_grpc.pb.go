// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.0
// source: proto/url.proto

package server

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

// DecreaseUrlClient is the client API for DecreaseUrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecreaseUrlClient interface {
	Create(ctx context.Context, in *Url, opts ...grpc.CallOption) (*TinyUrl, error)
	Get(ctx context.Context, in *TinyUrl, opts ...grpc.CallOption) (*Url, error)
}

type decreaseUrlClient struct {
	cc grpc.ClientConnInterface
}

func NewDecreaseUrlClient(cc grpc.ClientConnInterface) DecreaseUrlClient {
	return &decreaseUrlClient{cc}
}

func (c *decreaseUrlClient) Create(ctx context.Context, in *Url, opts ...grpc.CallOption) (*TinyUrl, error) {
	out := new(TinyUrl)
	err := c.cc.Invoke(ctx, "/url.DecreaseUrl/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decreaseUrlClient) Get(ctx context.Context, in *TinyUrl, opts ...grpc.CallOption) (*Url, error) {
	out := new(Url)
	err := c.cc.Invoke(ctx, "/url.DecreaseUrl/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecreaseUrlServer is the server API for DecreaseUrl service.
// All implementations must embed UnimplementedDecreaseUrlServer
// for forward compatibility
type DecreaseUrlServer interface {
	Create(context.Context, *Url) (*TinyUrl, error)
	Get(context.Context, *TinyUrl) (*Url, error)
	mustEmbedUnimplementedDecreaseUrlServer()
}

// UnimplementedDecreaseUrlServer must be embedded to have forward compatible implementations.
type UnimplementedDecreaseUrlServer struct {
}

func (UnimplementedDecreaseUrlServer) Create(context.Context, *Url) (*TinyUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDecreaseUrlServer) Get(context.Context, *TinyUrl) (*Url, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDecreaseUrlServer) mustEmbedUnimplementedDecreaseUrlServer() {}

// UnsafeDecreaseUrlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecreaseUrlServer will
// result in compilation errors.
type UnsafeDecreaseUrlServer interface {
	mustEmbedUnimplementedDecreaseUrlServer()
}

func RegisterDecreaseUrlServer(s grpc.ServiceRegistrar, srv DecreaseUrlServer) {
	s.RegisterService(&DecreaseUrl_ServiceDesc, srv)
}

func _DecreaseUrl_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Url)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecreaseUrlServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/url.DecreaseUrl/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecreaseUrlServer).Create(ctx, req.(*Url))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecreaseUrl_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TinyUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecreaseUrlServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/url.DecreaseUrl/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecreaseUrlServer).Get(ctx, req.(*TinyUrl))
	}
	return interceptor(ctx, in, info, handler)
}

// DecreaseUrl_ServiceDesc is the grpc.ServiceDesc for DecreaseUrl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecreaseUrl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "url.DecreaseUrl",
	HandlerType: (*DecreaseUrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DecreaseUrl_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DecreaseUrl_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/url.proto",
}
