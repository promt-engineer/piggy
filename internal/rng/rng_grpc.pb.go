// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: internal/rng/rng.proto

package rng

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RNG_Rand_FullMethodName        = "/RNG/Rand"
	RNG_RandFloat_FullMethodName   = "/RNG/RandFloat"
	RNG_HealthCheck_FullMethodName = "/RNG/HealthCheck"
)

// RNGClient is the client API for RNG service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RNGClient interface {
	Rand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error)
	RandFloat(ctx context.Context, in *RandRequestFloat, opts ...grpc.CallOption) (*RandResponseFloat, error)
	HealthCheck(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Status, Status], error)
}

type rNGClient struct {
	cc grpc.ClientConnInterface
}

func NewRNGClient(cc grpc.ClientConnInterface) RNGClient {
	return &rNGClient{cc}
}

func (c *rNGClient) Rand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandResponse)
	err := c.cc.Invoke(ctx, RNG_Rand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rNGClient) RandFloat(ctx context.Context, in *RandRequestFloat, opts ...grpc.CallOption) (*RandResponseFloat, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandResponseFloat)
	err := c.cc.Invoke(ctx, RNG_RandFloat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rNGClient) HealthCheck(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Status, Status], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RNG_ServiceDesc.Streams[0], RNG_HealthCheck_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Status, Status]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RNG_HealthCheckClient = grpc.BidiStreamingClient[Status, Status]

// RNGServer is the server API for RNG service.
// All implementations must embed UnimplementedRNGServer
// for forward compatibility.
type RNGServer interface {
	Rand(context.Context, *RandRequest) (*RandResponse, error)
	RandFloat(context.Context, *RandRequestFloat) (*RandResponseFloat, error)
	HealthCheck(grpc.BidiStreamingServer[Status, Status]) error
	mustEmbedUnimplementedRNGServer()
}

// UnimplementedRNGServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRNGServer struct{}

func (UnimplementedRNGServer) Rand(context.Context, *RandRequest) (*RandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rand not implemented")
}
func (UnimplementedRNGServer) RandFloat(context.Context, *RandRequestFloat) (*RandResponseFloat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandFloat not implemented")
}
func (UnimplementedRNGServer) HealthCheck(grpc.BidiStreamingServer[Status, Status]) error {
	return status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedRNGServer) mustEmbedUnimplementedRNGServer() {}
func (UnimplementedRNGServer) testEmbeddedByValue()             {}

// UnsafeRNGServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RNGServer will
// result in compilation errors.
type UnsafeRNGServer interface {
	mustEmbedUnimplementedRNGServer()
}

func RegisterRNGServer(s grpc.ServiceRegistrar, srv RNGServer) {
	// If the following call pancis, it indicates UnimplementedRNGServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RNG_ServiceDesc, srv)
}

func _RNG_Rand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RNGServer).Rand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RNG_Rand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RNGServer).Rand(ctx, req.(*RandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RNG_RandFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandRequestFloat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RNGServer).RandFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RNG_RandFloat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RNGServer).RandFloat(ctx, req.(*RandRequestFloat))
	}
	return interceptor(ctx, in, info, handler)
}

func _RNG_HealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RNGServer).HealthCheck(&grpc.GenericServerStream[Status, Status]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RNG_HealthCheckServer = grpc.BidiStreamingServer[Status, Status]

// RNG_ServiceDesc is the grpc.ServiceDesc for RNG service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RNG_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RNG",
	HandlerType: (*RNGServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rand",
			Handler:    _RNG_Rand_Handler,
		},
		{
			MethodName: "RandFloat",
			Handler:    _RNG_RandFloat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HealthCheck",
			Handler:       _RNG_HealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/rng/rng.proto",
}
