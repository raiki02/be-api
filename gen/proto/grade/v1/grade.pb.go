// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/111/v1/111.proto

package gradev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求体
type GradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"` //学号
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`                    //密码
}

func (x *GradeRequest) Reset() {
	*x = GradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_v1_grade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradeRequest) ProtoMessage() {}

func (x *GradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_v1_grade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradeRequest.ProtoReflect.Descriptor instead.
func (*GradeRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_v1_grade_proto_rawDescGZIP(), []int{0}
}

func (x *GradeRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *GradeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 响应体
type GradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 响应码
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`                  // 描述
	Class       []*Class `protobuf:"bytes,3,rep,name=class,proto3" json:"class,omitempty"`                              // 课程信息
}

func (x *GradeResponse) Reset() {
	*x = GradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_v1_grade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradeResponse) ProtoMessage() {}

func (x *GradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_v1_grade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradeResponse.ProtoReflect.Descriptor instead.
func (*GradeResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_v1_grade_proto_rawDescGZIP(), []int{1}
}

func (x *GradeResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GradeResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GradeResponse) GetClass() []*Class {
	if x != nil {
		return x.Class
	}
	return nil
}

//成绩结构体
type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course              string `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`                                                        //课程名
	Credit              string `protobuf:"bytes,2,opt,name=credit,proto3" json:"credit,omitempty"`                                                        //学分
	Grade               string `protobuf:"bytes,3,opt,name=111,proto3" json:"111,omitempty"`                                                          //总成绩
	RegularGrade        string `protobuf:"bytes,4,opt,name=regular_grade,json=regularGrade,proto3" json:"regular_grade,omitempty"`                        //平时成绩
	RegularGradePercent string `protobuf:"bytes,5,opt,name=regular_grade_percent,json=regularGradePercent,proto3" json:"regular_grade_percent,omitempty"` //平时成绩占比
	FinalGrade          string `protobuf:"bytes,6,opt,name=final_grade,json=finalGrade,proto3" json:"final_grade,omitempty"`                              //期末成绩
	FinalGradePercent   string `protobuf:"bytes,7,opt,name=final_grade_percent,json=finalGradePercent,proto3" json:"final_grade_percent,omitempty"`       //期末成绩占比
	Kcxzmc              string `protobuf:"bytes,8,opt,name=kcxzmc,proto3" json:"kcxzmc,omitempty"`                                                        //课程性质名称 比如专业主干课程/通识必修课
	Kclbmc              string `protobuf:"bytes,9,opt,name=Kclbmc,proto3" json:"Kclbmc,omitempty"`                                                        //课程类别名称，比如专业课/公共课
	Kcbj                string `protobuf:"bytes,10,opt,name=kcbj,proto3" json:"kcbj,omitempty"`                                                           //课程标记，比如主修/辅修
	Xnm                 string `protobuf:"bytes,11,opt,name=xnm,proto3" json:"xnm,omitempty"`                                                             //学年
	Xqm                 string `protobuf:"bytes,12,opt,name=xqm,proto3" json:"xqm,omitempty"`                                                             //学期名
	JxbId               string `protobuf:"bytes,13,opt,name=jxbId,proto3" json:"jxbId,omitempty"`                                                         //课程Id
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_v1_grade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_v1_grade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_proto_grade_v1_grade_proto_rawDescGZIP(), []int{2}
}

func (x *Class) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Class) GetCredit() string {
	if x != nil {
		return x.Credit
	}
	return ""
}

func (x *Class) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *Class) GetRegularGrade() string {
	if x != nil {
		return x.RegularGrade
	}
	return ""
}

func (x *Class) GetRegularGradePercent() string {
	if x != nil {
		return x.RegularGradePercent
	}
	return ""
}

func (x *Class) GetFinalGrade() string {
	if x != nil {
		return x.FinalGrade
	}
	return ""
}

func (x *Class) GetFinalGradePercent() string {
	if x != nil {
		return x.FinalGradePercent
	}
	return ""
}

func (x *Class) GetKcxzmc() string {
	if x != nil {
		return x.Kcxzmc
	}
	return ""
}

func (x *Class) GetKclbmc() string {
	if x != nil {
		return x.Kclbmc
	}
	return ""
}

func (x *Class) GetKcbj() string {
	if x != nil {
		return x.Kcbj
	}
	return ""
}

func (x *Class) GetXnm() string {
	if x != nil {
		return x.Xnm
	}
	return ""
}

func (x *Class) GetXqm() string {
	if x != nil {
		return x.Xqm
	}
	return ""
}

func (x *Class) GetJxbId() string {
	if x != nil {
		return x.JxbId
	}
	return ""
}

var File_proto_grade_v1_grade_proto protoreflect.FileDescriptor

var file_proto_grade_v1_grade_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x79, 0x0a, 0x0d,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xf5, 0x02, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x15,
	0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x47, 0x72, 0x61, 0x64, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x63, 0x78, 0x7a, 0x6d, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6b, 0x63, 0x78, 0x7a, 0x6d, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x63, 0x6c,
	0x62, 0x6d, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4b, 0x63, 0x6c, 0x62, 0x6d,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x63, 0x62, 0x6a, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x63, 0x62, 0x6a, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x6e, 0x6d, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x78, 0x6e, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x71, 0x6d, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x71, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x78, 0x62,
	0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x78, 0x62, 0x49, 0x64, 0x32,
	0x41, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x92, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x47, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73,
	0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x67, 0x72, 0x61, 0x64, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x47, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x47, 0x72, 0x61, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grade_v1_grade_proto_rawDescOnce sync.Once
	file_proto_grade_v1_grade_proto_rawDescData = file_proto_grade_v1_grade_proto_rawDesc
)

func file_proto_grade_v1_grade_proto_rawDescGZIP() []byte {
	file_proto_grade_v1_grade_proto_rawDescOnce.Do(func() {
		file_proto_grade_v1_grade_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grade_v1_grade_proto_rawDescData)
	})
	return file_proto_grade_v1_grade_proto_rawDescData
}

var file_proto_grade_v1_grade_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_grade_v1_grade_proto_goTypes = []interface{}{
	(*GradeRequest)(nil),  // 0: 111.v1.GradeRequest
	(*GradeResponse)(nil), // 1: 111.v1.GradeResponse
	(*Class)(nil),         // 2: 111.v1.Class
}
var file_proto_grade_v1_grade_proto_depIdxs = []int32{
	2, // 0: 111.v1.GradeResponse.class:type_name -> 111.v1.Class
	0, // 1: 111.v1.Grade.Grade:input_type -> 111.v1.GradeRequest
	1, // 2: 111.v1.Grade.Grade:output_type -> 111.v1.GradeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_grade_v1_grade_proto_init() }
func file_proto_grade_v1_grade_proto_init() {
	if File_proto_grade_v1_grade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grade_v1_grade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradeRequest); i {
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
		file_proto_grade_v1_grade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradeResponse); i {
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
		file_proto_grade_v1_grade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
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
			RawDescriptor: file_proto_grade_v1_grade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grade_v1_grade_proto_goTypes,
		DependencyIndexes: file_proto_grade_v1_grade_proto_depIdxs,
		MessageInfos:      file_proto_grade_v1_grade_proto_msgTypes,
	}.Build()
	File_proto_grade_v1_grade_proto = out.File
	file_proto_grade_v1_grade_proto_rawDesc = nil
	file_proto_grade_v1_grade_proto_goTypes = nil
	file_proto_grade_v1_grade_proto_depIdxs = nil
}
