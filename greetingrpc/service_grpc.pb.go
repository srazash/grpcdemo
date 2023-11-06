// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: service.proto

package greetingrpc

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

// GreetingRpcClient is the client API for GreetingRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingRpcClient interface {
	GetGreeting(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Greeting, error)
}

type greetingRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingRpcClient(cc grpc.ClientConnInterface) GreetingRpcClient {
	return &greetingRpcClient{cc}
}

func (c *greetingRpcClient) GetGreeting(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := c.cc.Invoke(ctx, "/GreetingRpc/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingRpcServer is the server API for GreetingRpc service.
// All implementations must embed UnimplementedGreetingRpcServer
// for forward compatibility
type GreetingRpcServer interface {
	GetGreeting(context.Context, *Name) (*Greeting, error)
	mustEmbedUnimplementedGreetingRpcServer()
}

// UnimplementedGreetingRpcServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingRpcServer struct {
}

func (UnimplementedGreetingRpcServer) GetGreeting(context.Context, *Name) (*Greeting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeting not implemented")
}
func (UnimplementedGreetingRpcServer) mustEmbedUnimplementedGreetingRpcServer() {}

// UnsafeGreetingRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingRpcServer will
// result in compilation errors.
type UnsafeGreetingRpcServer interface {
	mustEmbedUnimplementedGreetingRpcServer()
}

func RegisterGreetingRpcServer(s grpc.ServiceRegistrar, srv GreetingRpcServer) {
	s.RegisterService(&GreetingRpc_ServiceDesc, srv)
}

func _GreetingRpc_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingRpcServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreetingRpc/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingRpcServer).GetGreeting(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingRpc_ServiceDesc is the grpc.ServiceDesc for GreetingRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GreetingRpc",
	HandlerType: (*GreetingRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _GreetingRpc_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}