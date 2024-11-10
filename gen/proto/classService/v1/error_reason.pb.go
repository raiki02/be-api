// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/classService/v1/error_reason.proto

package classServicev1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_ES_AddClassFailed    ErrorReason = 0
	ErrorReason_ES_SearchClassFailed ErrorReason = 1
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "ES_AddClassFailed",
		1: "ES_SearchClassFailed",
	}
	ErrorReason_value = map[string]int32{
		"ES_AddClassFailed":    0,
		"ES_SearchClassFailed": 1,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_classService_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_proto_classService_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_classService_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_proto_classService_v1_error_reason_proto protoreflect.FileDescriptor

var file_proto_classService_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x4a, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x11, 0x45, 0x53, 0x5f, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xc2, 0x03, 0x12, 0x1e, 0x0a, 0x14,
	0x45, 0x53, 0x5f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xc3, 0x03, 0x42, 0xc9, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_classService_v1_error_reason_proto_rawDescOnce sync.Once
	file_proto_classService_v1_error_reason_proto_rawDescData = file_proto_classService_v1_error_reason_proto_rawDesc
)

func file_proto_classService_v1_error_reason_proto_rawDescGZIP() []byte {
	file_proto_classService_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_proto_classService_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_classService_v1_error_reason_proto_rawDescData)
	})
	return file_proto_classService_v1_error_reason_proto_rawDescData
}

var file_proto_classService_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_classService_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: classService.v1.ErrorReason
}
var file_proto_classService_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_classService_v1_error_reason_proto_init() }
func file_proto_classService_v1_error_reason_proto_init() {
	if File_proto_classService_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_classService_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_classService_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_proto_classService_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_proto_classService_v1_error_reason_proto_enumTypes,
	}.Build()
	File_proto_classService_v1_error_reason_proto = out.File
	file_proto_classService_v1_error_reason_proto_rawDesc = nil
	file_proto_classService_v1_error_reason_proto_goTypes = nil
	file_proto_classService_v1_error_reason_proto_depIdxs = nil
}
