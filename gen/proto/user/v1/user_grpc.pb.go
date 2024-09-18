// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/user/v1/user.proto

package userv1

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
	UserService_FindOrCreateByStudentId_FullMethodName = "/user.v1.UserService/FindOrCreateByStudentId"
	UserService_UpdateNonSensitiveInfo_FullMethodName  = "/user.v1.UserService/UpdateNonSensitiveInfo"
	UserService_GetCookie_FullMethodName               = "/user.v1.UserService/GetCookie"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindOrCreateByStudentId(ctx context.Context, in *FindOrCreateByStudentIdRequest, opts ...grpc.CallOption) (*FindOrCreateByStudentIdResponse, error)
	UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error)
	GetCookie(ctx context.Context, in *GetCookieRequest, opts ...grpc.CallOption) (*GetCookieResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOrCreateByStudentId(ctx context.Context, in *FindOrCreateByStudentIdRequest, opts ...grpc.CallOption) (*FindOrCreateByStudentIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindOrCreateByStudentIdResponse)
	err := c.cc.Invoke(ctx, UserService_FindOrCreateByStudentId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNonSensitiveInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateNonSensitiveInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCookie(ctx context.Context, in *GetCookieRequest, opts ...grpc.CallOption) (*GetCookieResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCookieResponse)
	err := c.cc.Invoke(ctx, UserService_GetCookie_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	FindOrCreateByStudentId(context.Context, *FindOrCreateByStudentIdRequest) (*FindOrCreateByStudentIdResponse, error)
	UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error)
	GetCookie(context.Context, *GetCookieRequest) (*GetCookieResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) FindOrCreateByStudentId(context.Context, *FindOrCreateByStudentIdRequest) (*FindOrCreateByStudentIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrCreateByStudentId not implemented")
}
func (UnimplementedUserServiceServer) UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNonSensitiveInfo not implemented")
}
func (UnimplementedUserServiceServer) GetCookie(context.Context, *GetCookieRequest) (*GetCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCookie not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindOrCreateByStudentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrCreateByStudentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOrCreateByStudentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindOrCreateByStudentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOrCreateByStudentId(ctx, req.(*FindOrCreateByStudentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateNonSensitiveInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNonSensitiveInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateNonSensitiveInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateNonSensitiveInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateNonSensitiveInfo(ctx, req.(*UpdateNonSensitiveInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCookie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCookie(ctx, req.(*GetCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOrCreateByStudentId",
			Handler:    _UserService_FindOrCreateByStudentId_Handler,
		},
		{
			MethodName: "UpdateNonSensitiveInfo",
			Handler:    _UserService_UpdateNonSensitiveInfo_Handler,
		},
		{
			MethodName: "GetCookie",
			Handler:    _UserService_GetCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/v1/user.proto",
}
