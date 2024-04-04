// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/ccnu/ccnu.proto

package ccnuv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_ccnu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_ccnu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_ccnu_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_ccnu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_ccnu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_ccnu_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_ccnu_ccnu_proto protoreflect.FileDescriptor

var file_proto_ccnu_ccnu_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x63, 0x63, 0x6e,
	0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31,
	0x22, 0x49, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x45, 0x0a, 0x0b, 0x43, 0x43, 0x4e, 0x55, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15,
	0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6b, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x43, 0x63,
	0x6e, 0x75, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x63, 0x6e, 0x75, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x63, 0x6e, 0x75, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x07, 0x43, 0x63, 0x6e, 0x75, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x63, 0x6e, 0x75,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x08, 0x43, 0x63, 0x6e, 0x75, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_ccnu_ccnu_proto_rawDescOnce sync.Once
	file_proto_ccnu_ccnu_proto_rawDescData = file_proto_ccnu_ccnu_proto_rawDesc
)

func file_proto_ccnu_ccnu_proto_rawDescGZIP() []byte {
	file_proto_ccnu_ccnu_proto_rawDescOnce.Do(func() {
		file_proto_ccnu_ccnu_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ccnu_ccnu_proto_rawDescData)
	})
	return file_proto_ccnu_ccnu_proto_rawDescData
}

var file_proto_ccnu_ccnu_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ccnu_ccnu_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),  // 0: ccnu.v1.LoginRequest
	(*LoginResponse)(nil), // 1: ccnu.v1.LoginResponse
}
var file_proto_ccnu_ccnu_proto_depIdxs = []int32{
	0, // 0: ccnu.v1.CCNUService.Login:input_type -> ccnu.v1.LoginRequest
	1, // 1: ccnu.v1.CCNUService.Login:output_type -> ccnu.v1.LoginResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ccnu_ccnu_proto_init() }
func file_proto_ccnu_ccnu_proto_init() {
	if File_proto_ccnu_ccnu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ccnu_ccnu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ccnu_ccnu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ccnu_ccnu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ccnu_ccnu_proto_goTypes,
		DependencyIndexes: file_proto_ccnu_ccnu_proto_depIdxs,
		MessageInfos:      file_proto_ccnu_ccnu_proto_msgTypes,
	}.Build()
	File_proto_ccnu_ccnu_proto = out.File
	file_proto_ccnu_ccnu_proto_rawDesc = nil
	file_proto_ccnu_ccnu_proto_goTypes = nil
	file_proto_ccnu_ccnu_proto_depIdxs = nil
}
