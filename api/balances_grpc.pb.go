// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/balances.proto

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

// BalancesClient is the client API for Balances service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancesClient interface {
	// Получение остатков
	Get(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error)
}

type balancesClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancesClient(cc grpc.ClientConnInterface) BalancesClient {
	return &balancesClient{cc}
}

func (c *balancesClient) Get(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error) {
	out := new(BalancesResponse)
	err := c.cc.Invoke(ctx, "/balances.Balances/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancesServer is the server API for Balances service.
// All implementations must embed UnimplementedBalancesServer
// for forward compatibility
type BalancesServer interface {
	// Получение остатков
	Get(context.Context, *BalancesRequest) (*BalancesResponse, error)
	mustEmbedUnimplementedBalancesServer()
}

// UnimplementedBalancesServer must be embedded to have forward compatible implementations.
type UnimplementedBalancesServer struct {
}

func (UnimplementedBalancesServer) Get(context.Context, *BalancesRequest) (*BalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBalancesServer) mustEmbedUnimplementedBalancesServer() {}

// UnsafeBalancesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancesServer will
// result in compilation errors.
type UnsafeBalancesServer interface {
	mustEmbedUnimplementedBalancesServer()
}

func RegisterBalancesServer(s grpc.ServiceRegistrar, srv BalancesServer) {
	s.RegisterService(&Balances_ServiceDesc, srv)
}

func _Balances_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balances.Balances/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancesServer).Get(ctx, req.(*BalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Balances_ServiceDesc is the grpc.ServiceDesc for Balances service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Balances_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "balances.Balances",
	HandlerType: (*BalancesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Balances_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/balances.proto",
}
