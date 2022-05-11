// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/reverser.proto

package api

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

// ReverserClient is the client API for Reverser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReverserClient interface {
	Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error)
}

type reverserClient struct {
	cc grpc.ClientConnInterface
}

func NewReverserClient(cc grpc.ClientConnInterface) ReverserClient {
	return &reverserClient{cc}
}

func (c *reverserClient) Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error) {
	out := new(ReverseResponse)
	err := c.cc.Invoke(ctx, "/proto.Reverser/Reverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReverserServer is the server API for Reverser service.
// All implementations must embed UnimplementedReverserServer
// for forward compatibility
type ReverserServer interface {
	Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error)
	//mustEmbedUnimplementedReverserServer()
}

// UnimplementedReverserServer must be embedded to have forward compatible implementations.
type UnimplementedReverserServer struct {
}

func (UnimplementedReverserServer) Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reverse not implemented")
}
//func (UnimplementedReverserServer) mustEmbedUnimplementedReverserServer() {}

// UnsafeReverserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReverserServer will
// result in compilation errors.
type UnsafeReverserServer interface {
	//mustEmbedUnimplementedReverserServer()
}

func RegisterReverserServer(s grpc.ServiceRegistrar, srv ReverserServer) {
	s.RegisterService(&Reverser_ServiceDesc, srv)
}

func _Reverser_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReverserServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Reverser/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReverserServer).Reverse(ctx, req.(*ReverseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reverser_ServiceDesc is the grpc.ServiceDesc for Reverser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reverser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Reverser",
	HandlerType: (*ReverserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reverse",
			Handler:    _Reverser_Reverse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reverser.proto",
}
