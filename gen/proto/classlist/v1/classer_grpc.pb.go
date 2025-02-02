// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: proto/classlist/v1/classer.proto

package classerv1

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
	Classer_GetClass_FullMethodName                = "/classer.v1.Classer/GetClass"
	Classer_AddClass_FullMethodName                = "/classer.v1.Classer/AddClass"
	Classer_DeleteClass_FullMethodName             = "/classer.v1.Classer/DeleteClass"
	Classer_UpdateClass_FullMethodName             = "/classer.v1.Classer/UpdateClass"
	Classer_GetRecycleBinClassInfos_FullMethodName = "/classer.v1.Classer/GetRecycleBinClassInfos"
	Classer_RecoverClass_FullMethodName            = "/classer.v1.Classer/RecoverClass"
	Classer_GetAllClassInfo_FullMethodName         = "/classer.v1.Classer/GetAllClassInfo"
	Classer_GetStuIdByJxbId_FullMethodName         = "/classer.v1.Classer/GetStuIdByJxbId"
	Classer_GetSchoolDay_FullMethodName            = "/classer.v1.Classer/GetSchoolDay"
)

// ClasserClient is the client API for Classer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClasserClient interface {
	// 获取课表
	GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error)
	// 添加课程
	AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassResponse, error)
	// 删除课程
	DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error)
	// 更新课程
	UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error)
	// 获取回收站的课程(回收站的课程只能保存2个月)
	GetRecycleBinClassInfos(ctx context.Context, in *GetRecycleBinClassRequest, opts ...grpc.CallOption) (*GetRecycleBinClassResponse, error)
	// 恢复课程
	RecoverClass(ctx context.Context, in *RecoverClassRequest, opts ...grpc.CallOption) (*RecoverClassResponse, error)
	// 获取所有课程信息(为其他服务设置的)
	GetAllClassInfo(ctx context.Context, in *GetAllClassInfoRequest, opts ...grpc.CallOption) (*GetAllClassInfoResponse, error)
	// 获取教学班中的所有学生ID
	GetStuIdByJxbId(ctx context.Context, in *GetStuIdByJxbIdRequest, opts ...grpc.CallOption) (*GetStuIdByJxbIdResponse, error)
	// 获取相关日期
	GetSchoolDay(ctx context.Context, in *GetSchoolDayReq, opts ...grpc.CallOption) (*GetSchoolDayResp, error)
}

type classerClient struct {
	cc grpc.ClientConnInterface
}

func NewClasserClient(cc grpc.ClientConnInterface) ClasserClient {
	return &classerClient{cc}
}

func (c *classerClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClassResponse)
	err := c.cc.Invoke(ctx, Classer_GetClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddClassResponse)
	err := c.cc.Invoke(ctx, Classer_AddClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) DeleteClass(ctx context.Context, in *DeleteClassRequest, opts ...grpc.CallOption) (*DeleteClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteClassResponse)
	err := c.cc.Invoke(ctx, Classer_DeleteClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) UpdateClass(ctx context.Context, in *UpdateClassRequest, opts ...grpc.CallOption) (*UpdateClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateClassResponse)
	err := c.cc.Invoke(ctx, Classer_UpdateClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) GetRecycleBinClassInfos(ctx context.Context, in *GetRecycleBinClassRequest, opts ...grpc.CallOption) (*GetRecycleBinClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecycleBinClassResponse)
	err := c.cc.Invoke(ctx, Classer_GetRecycleBinClassInfos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) RecoverClass(ctx context.Context, in *RecoverClassRequest, opts ...grpc.CallOption) (*RecoverClassResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecoverClassResponse)
	err := c.cc.Invoke(ctx, Classer_RecoverClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) GetAllClassInfo(ctx context.Context, in *GetAllClassInfoRequest, opts ...grpc.CallOption) (*GetAllClassInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllClassInfoResponse)
	err := c.cc.Invoke(ctx, Classer_GetAllClassInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) GetStuIdByJxbId(ctx context.Context, in *GetStuIdByJxbIdRequest, opts ...grpc.CallOption) (*GetStuIdByJxbIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStuIdByJxbIdResponse)
	err := c.cc.Invoke(ctx, Classer_GetStuIdByJxbId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classerClient) GetSchoolDay(ctx context.Context, in *GetSchoolDayReq, opts ...grpc.CallOption) (*GetSchoolDayResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSchoolDayResp)
	err := c.cc.Invoke(ctx, Classer_GetSchoolDay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClasserServer is the server API for Classer service.
// All implementations must embed UnimplementedClasserServer
// for forward compatibility.
type ClasserServer interface {
	// 获取课表
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	// 添加课程
	AddClass(context.Context, *AddClassRequest) (*AddClassResponse, error)
	// 删除课程
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)
	// 更新课程
	UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error)
	// 获取回收站的课程(回收站的课程只能保存2个月)
	GetRecycleBinClassInfos(context.Context, *GetRecycleBinClassRequest) (*GetRecycleBinClassResponse, error)
	// 恢复课程
	RecoverClass(context.Context, *RecoverClassRequest) (*RecoverClassResponse, error)
	// 获取所有课程信息(为其他服务设置的)
	GetAllClassInfo(context.Context, *GetAllClassInfoRequest) (*GetAllClassInfoResponse, error)
	// 获取教学班中的所有学生ID
	GetStuIdByJxbId(context.Context, *GetStuIdByJxbIdRequest) (*GetStuIdByJxbIdResponse, error)
	// 获取相关日期
	GetSchoolDay(context.Context, *GetSchoolDayReq) (*GetSchoolDayResp, error)
	mustEmbedUnimplementedClasserServer()
}

// UnimplementedClasserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClasserServer struct{}

func (UnimplementedClasserServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (UnimplementedClasserServer) AddClass(context.Context, *AddClassRequest) (*AddClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClass not implemented")
}
func (UnimplementedClasserServer) DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClass not implemented")
}
func (UnimplementedClasserServer) UpdateClass(context.Context, *UpdateClassRequest) (*UpdateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClass not implemented")
}
func (UnimplementedClasserServer) GetRecycleBinClassInfos(context.Context, *GetRecycleBinClassRequest) (*GetRecycleBinClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecycleBinClassInfos not implemented")
}
func (UnimplementedClasserServer) RecoverClass(context.Context, *RecoverClassRequest) (*RecoverClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverClass not implemented")
}
func (UnimplementedClasserServer) GetAllClassInfo(context.Context, *GetAllClassInfoRequest) (*GetAllClassInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllClassInfo not implemented")
}
func (UnimplementedClasserServer) GetStuIdByJxbId(context.Context, *GetStuIdByJxbIdRequest) (*GetStuIdByJxbIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStuIdByJxbId not implemented")
}
func (UnimplementedClasserServer) GetSchoolDay(context.Context, *GetSchoolDayReq) (*GetSchoolDayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchoolDay not implemented")
}
func (UnimplementedClasserServer) mustEmbedUnimplementedClasserServer() {}
func (UnimplementedClasserServer) testEmbeddedByValue()                 {}

// UnsafeClasserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClasserServer will
// result in compilation errors.
type UnsafeClasserServer interface {
	mustEmbedUnimplementedClasserServer()
}

func RegisterClasserServer(s grpc.ServiceRegistrar, srv ClasserServer) {
	// If the following call pancis, it indicates UnimplementedClasserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Classer_ServiceDesc, srv)
}

func _Classer_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_GetClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).GetClass(ctx, req.(*GetClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_AddClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).AddClass(ctx, req.(*AddClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_DeleteClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).DeleteClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_DeleteClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).DeleteClass(ctx, req.(*DeleteClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_UpdateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).UpdateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_UpdateClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).UpdateClass(ctx, req.(*UpdateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_GetRecycleBinClassInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecycleBinClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).GetRecycleBinClassInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_GetRecycleBinClassInfos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).GetRecycleBinClassInfos(ctx, req.(*GetRecycleBinClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_RecoverClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).RecoverClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_RecoverClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).RecoverClass(ctx, req.(*RecoverClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_GetAllClassInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllClassInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).GetAllClassInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_GetAllClassInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).GetAllClassInfo(ctx, req.(*GetAllClassInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_GetStuIdByJxbId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStuIdByJxbIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).GetStuIdByJxbId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_GetStuIdByJxbId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).GetStuIdByJxbId(ctx, req.(*GetStuIdByJxbIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Classer_GetSchoolDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchoolDayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClasserServer).GetSchoolDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classer_GetSchoolDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClasserServer).GetSchoolDay(ctx, req.(*GetSchoolDayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Classer_ServiceDesc is the grpc.ServiceDesc for Classer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Classer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "classer.v1.Classer",
	HandlerType: (*ClasserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClass",
			Handler:    _Classer_GetClass_Handler,
		},
		{
			MethodName: "AddClass",
			Handler:    _Classer_AddClass_Handler,
		},
		{
			MethodName: "DeleteClass",
			Handler:    _Classer_DeleteClass_Handler,
		},
		{
			MethodName: "UpdateClass",
			Handler:    _Classer_UpdateClass_Handler,
		},
		{
			MethodName: "GetRecycleBinClassInfos",
			Handler:    _Classer_GetRecycleBinClassInfos_Handler,
		},
		{
			MethodName: "RecoverClass",
			Handler:    _Classer_RecoverClass_Handler,
		},
		{
			MethodName: "GetAllClassInfo",
			Handler:    _Classer_GetAllClassInfo_Handler,
		},
		{
			MethodName: "GetStuIdByJxbId",
			Handler:    _Classer_GetStuIdByJxbId_Handler,
		},
		{
			MethodName: "GetSchoolDay",
			Handler:    _Classer_GetSchoolDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/classlist/v1/classer.proto",
}
