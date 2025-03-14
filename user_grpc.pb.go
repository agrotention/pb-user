// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: user.proto

package pb_user

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
	ServiceUser_Create_FullMethodName = "/ServiceUser/Create"
)

// ServiceUserClient is the client API for ServiceUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceUserClient interface {
	Create(ctx context.Context, in *InCreate, opts ...grpc.CallOption) (*OutSuccess, error)
}

type serviceUserClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceUserClient(cc grpc.ClientConnInterface) ServiceUserClient {
	return &serviceUserClient{cc}
}

func (c *serviceUserClient) Create(ctx context.Context, in *InCreate, opts ...grpc.CallOption) (*OutSuccess, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OutSuccess)
	err := c.cc.Invoke(ctx, ServiceUser_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceUserServer is the server API for ServiceUser service.
// All implementations must embed UnimplementedServiceUserServer
// for forward compatibility.
type ServiceUserServer interface {
	Create(context.Context, *InCreate) (*OutSuccess, error)
	mustEmbedUnimplementedServiceUserServer()
}

// UnimplementedServiceUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceUserServer struct{}

func (UnimplementedServiceUserServer) Create(context.Context, *InCreate) (*OutSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServiceUserServer) mustEmbedUnimplementedServiceUserServer() {}
func (UnimplementedServiceUserServer) testEmbeddedByValue()                     {}

// UnsafeServiceUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceUserServer will
// result in compilation errors.
type UnsafeServiceUserServer interface {
	mustEmbedUnimplementedServiceUserServer()
}

func RegisterServiceUserServer(s grpc.ServiceRegistrar, srv ServiceUserServer) {
	// If the following call pancis, it indicates UnimplementedServiceUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceUser_ServiceDesc, srv)
}

func _ServiceUser_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceUserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceUser_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceUserServer).Create(ctx, req.(*InCreate))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceUser_ServiceDesc is the grpc.ServiceDesc for ServiceUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceUser",
	HandlerType: (*ServiceUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceUser_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
