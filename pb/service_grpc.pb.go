// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/service.proto

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

// ServiceHelloClient is the client API for ServiceHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceHelloClient interface {
	ValidateCPF(ctx context.Context, in *CPFRequest, opts ...grpc.CallOption) (*CPFResponse, error)
	ValidateCNPJ(ctx context.Context, in *CNPJRequest, opts ...grpc.CallOption) (*CNPJResponse, error)
}

type serviceHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceHelloClient(cc grpc.ClientConnInterface) ServiceHelloClient {
	return &serviceHelloClient{cc}
}

func (c *serviceHelloClient) ValidateCPF(ctx context.Context, in *CPFRequest, opts ...grpc.CallOption) (*CPFResponse, error) {
	out := new(CPFResponse)
	err := c.cc.Invoke(ctx, "/serviceHello/ValidateCPF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceHelloClient) ValidateCNPJ(ctx context.Context, in *CNPJRequest, opts ...grpc.CallOption) (*CNPJResponse, error) {
	out := new(CNPJResponse)
	err := c.cc.Invoke(ctx, "/serviceHello/ValidateCNPJ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceHelloServer is the server API for ServiceHello service.
// All implementations must embed UnimplementedServiceHelloServer
// for forward compatibility
type ServiceHelloServer interface {
	ValidateCPF(context.Context, *CPFRequest) (*CPFResponse, error)
	ValidateCNPJ(context.Context, *CNPJRequest) (*CNPJResponse, error)
	mustEmbedUnimplementedServiceHelloServer()
}

// UnimplementedServiceHelloServer must be embedded to have forward compatible implementations.
type UnimplementedServiceHelloServer struct {
}

func (UnimplementedServiceHelloServer) ValidateCPF(context.Context, *CPFRequest) (*CPFResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCPF not implemented")
}
func (UnimplementedServiceHelloServer) ValidateCNPJ(context.Context, *CNPJRequest) (*CNPJResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCNPJ not implemented")
}
func (UnimplementedServiceHelloServer) mustEmbedUnimplementedServiceHelloServer() {}

// UnsafeServiceHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceHelloServer will
// result in compilation errors.
type UnsafeServiceHelloServer interface {
	mustEmbedUnimplementedServiceHelloServer()
}

func RegisterServiceHelloServer(s grpc.ServiceRegistrar, srv ServiceHelloServer) {
	s.RegisterService(&ServiceHello_ServiceDesc, srv)
}

func _ServiceHello_ValidateCPF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CPFRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceHelloServer).ValidateCPF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceHello/ValidateCPF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceHelloServer).ValidateCPF(ctx, req.(*CPFRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceHello_ValidateCNPJ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CNPJRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceHelloServer).ValidateCNPJ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceHello/ValidateCNPJ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceHelloServer).ValidateCNPJ(ctx, req.(*CNPJRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceHello_ServiceDesc is the grpc.ServiceDesc for ServiceHello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceHello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serviceHello",
	HandlerType: (*ServiceHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCPF",
			Handler:    _ServiceHello_ValidateCPF_Handler,
		},
		{
			MethodName: "ValidateCNPJ",
			Handler:    _ServiceHello_ValidateCNPJ_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
