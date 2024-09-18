// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/ccnu/v1/ccnu.proto

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

type Source int32

const (
	Source_GradeApi Source = 0 // 来源于成绩查询接口，这里的信息很准确
	Source_OldXkApi Source = 1 // 来源于老的教务系统接口，这里有的课程信息是不全的
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "GradeApi",
		1: "OldXkApi",
	}
	Source_value = map[string]int32{
		"GradeApi": 0,
		"OldXkApi": 1,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ccnu_v1_ccnu_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_proto_ccnu_v1_ccnu_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{0}
}

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
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[0]
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
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[1]
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
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCCNUCookieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetCCNUCookieRequest) Reset() {
	*x = GetCCNUCookieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCCNUCookieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCCNUCookieRequest) ProtoMessage() {}

func (x *GetCCNUCookieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCCNUCookieRequest.ProtoReflect.Descriptor instead.
func (*GetCCNUCookieRequest) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{2}
}

func (x *GetCCNUCookieRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *GetCCNUCookieRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetCCNUCookieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie string `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
}

func (x *GetCCNUCookieResponse) Reset() {
	*x = GetCCNUCookieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCCNUCookieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCCNUCookieResponse) ProtoMessage() {}

func (x *GetCCNUCookieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCCNUCookieResponse.ProtoReflect.Descriptor instead.
func (*GetCCNUCookieResponse) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{3}
}

func (x *GetCCNUCookieResponse) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

type CourseListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Year      string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Term      string `protobuf:"bytes,4,opt,name=term,proto3" json:"term,omitempty"`
	Source    Source `protobuf:"varint,5,opt,name=source,proto3,enum=ccnu.v1.Source" json:"source,omitempty"`
}

func (x *CourseListRequest) Reset() {
	*x = CourseListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseListRequest) ProtoMessage() {}

func (x *CourseListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseListRequest.ProtoReflect.Descriptor instead.
func (*CourseListRequest) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{4}
}

func (x *CourseListRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CourseListRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CourseListRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *CourseListRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *CourseListRequest) GetSource() Source {
	if x != nil {
		return x.Source
	}
	return Source_GradeApi
}

// 这个course只是爬出来没怎么处理的，在course服务里会处理
type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseCode string  `protobuf:"bytes,1,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Teacher    string  `protobuf:"bytes,3,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Class      string  `protobuf:"bytes,4,opt,name=class,proto3" json:"class,omitempty"`
	School     string  `protobuf:"bytes,5,opt,name=school,proto3" json:"school,omitempty"`
	Property   string  `protobuf:"bytes,6,opt,name=property,proto3" json:"property,omitempty"` // 课程性质
	Credit     float64 `protobuf:"fixed64,7,opt,name=credit,proto3" json:"credit,omitempty"`
	Year       string  `protobuf:"bytes,8,opt,name=year,proto3" json:"year,omitempty"`
	Term       string  `protobuf:"bytes,9,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{5}
}

func (x *Course) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *Course) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Course) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *Course) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *Course) GetCredit() float64 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *Course) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Course) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type CourseListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *CourseListResponse) Reset() {
	*x = CourseListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseListResponse) ProtoMessage() {}

func (x *CourseListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseListResponse.ProtoReflect.Descriptor instead.
func (*CourseListResponse) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{6}
}

func (x *CourseListResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

type GetAllGradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetAllGradesRequest) Reset() {
	*x = GetAllGradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGradesRequest) ProtoMessage() {}

func (x *GetAllGradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGradesRequest.ProtoReflect.Descriptor instead.
func (*GetAllGradesRequest) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllGradesRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *GetAllGradesRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Grade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseCode    string  `protobuf:"bytes,1,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
	CourseName    string  `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	CourseTeacher string  `protobuf:"bytes,3,opt,name=course_teacher,json=courseTeacher,proto3" json:"course_teacher,omitempty"`
	Regular       float64 `protobuf:"fixed64,4,opt,name=regular,proto3" json:"regular,omitempty"` // 平时
	Final         float64 `protobuf:"fixed64,5,opt,name=final,proto3" json:"final,omitempty"`     // 期末
	Total         float64 `protobuf:"fixed64,6,opt,name=total,proto3" json:"total,omitempty"`     // 总评
	Year          string  `protobuf:"bytes,7,opt,name=year,proto3" json:"year,omitempty"`
	Term          string  `protobuf:"bytes,8,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *Grade) Reset() {
	*x = Grade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grade) ProtoMessage() {}

func (x *Grade) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grade.ProtoReflect.Descriptor instead.
func (*Grade) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{8}
}

func (x *Grade) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

func (x *Grade) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *Grade) GetCourseTeacher() string {
	if x != nil {
		return x.CourseTeacher
	}
	return ""
}

func (x *Grade) GetRegular() float64 {
	if x != nil {
		return x.Regular
	}
	return 0
}

func (x *Grade) GetFinal() float64 {
	if x != nil {
		return x.Final
	}
	return 0
}

func (x *Grade) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Grade) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Grade) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type GetAllGradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,1,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *GetAllGradesResponse) Reset() {
	*x = GetAllGradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGradesResponse) ProtoMessage() {}

func (x *GetAllGradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGradesResponse.ProtoReflect.Descriptor instead.
func (*GetAllGradesResponse) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllGradesResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

type GetGradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Year      string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Term      string `protobuf:"bytes,4,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *GetGradesRequest) Reset() {
	*x = GetGradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGradesRequest) ProtoMessage() {}

func (x *GetGradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGradesRequest.ProtoReflect.Descriptor instead.
func (*GetGradesRequest) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{10}
}

func (x *GetGradesRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *GetGradesRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetGradesRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetGradesRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type GetGradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,1,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *GetGradesResponse) Reset() {
	*x = GetGradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGradesResponse) ProtoMessage() {}

func (x *GetGradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ccnu_v1_ccnu_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGradesResponse.ProtoReflect.Descriptor instead.
func (*GetGradesResponse) Descriptor() ([]byte, []int) {
	return file_proto_ccnu_v1_ccnu_proto_rawDescGZIP(), []int{11}
}

func (x *GetGradesResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

var File_proto_ccnu_v1_ccnu_proto protoreflect.FileDescriptor

var file_proto_ccnu_v1_ccnu_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x63, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x63, 0x6e, 0x75,
	0x2e, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x43, 0x43, 0x4e, 0x55, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x43, 0x4e, 0x55, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0x9f, 0x01,
	0x0a, 0x11, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0xe1, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x63, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x3e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22, 0x75, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x3b,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x2a, 0x24, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x72, 0x61, 0x64, 0x65, 0x41, 0x70,
	0x69, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x6c, 0x64, 0x58, 0x6b, 0x41, 0x70, 0x69, 0x10,
	0x01, 0x32, 0xed, 0x02, 0x0a, 0x0b, 0x43, 0x43, 0x4e, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x63, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x43, 0x43, 0x4e, 0x55, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x63, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x43, 0x4e, 0x55, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x63, 0x6e, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x43, 0x4e, 0x55, 0x43, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x12, 0x1c, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x63, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x8a, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x63, 0x6e, 0x75, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x43, 0x63, 0x6e, 0x75, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63,
	0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x63, 0x6e,
	0x75, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x63, 0x6e, 0x75,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43, 0x63, 0x6e, 0x75, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x43, 0x63, 0x6e, 0x75, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43, 0x63, 0x6e, 0x75, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ccnu_v1_ccnu_proto_rawDescOnce sync.Once
	file_proto_ccnu_v1_ccnu_proto_rawDescData = file_proto_ccnu_v1_ccnu_proto_rawDesc
)

func file_proto_ccnu_v1_ccnu_proto_rawDescGZIP() []byte {
	file_proto_ccnu_v1_ccnu_proto_rawDescOnce.Do(func() {
		file_proto_ccnu_v1_ccnu_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ccnu_v1_ccnu_proto_rawDescData)
	})
	return file_proto_ccnu_v1_ccnu_proto_rawDescData
}

var file_proto_ccnu_v1_ccnu_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ccnu_v1_ccnu_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_ccnu_v1_ccnu_proto_goTypes = []interface{}{
	(Source)(0),                   // 0: ccnu.v1.Source
	(*LoginRequest)(nil),          // 1: ccnu.v1.LoginRequest
	(*LoginResponse)(nil),         // 2: ccnu.v1.LoginResponse
	(*GetCCNUCookieRequest)(nil),  // 3: ccnu.v1.GetCCNUCookieRequest
	(*GetCCNUCookieResponse)(nil), // 4: ccnu.v1.GetCCNUCookieResponse
	(*CourseListRequest)(nil),     // 5: ccnu.v1.CourseListRequest
	(*Course)(nil),                // 6: ccnu.v1.Course
	(*CourseListResponse)(nil),    // 7: ccnu.v1.CourseListResponse
	(*GetAllGradesRequest)(nil),   // 8: ccnu.v1.GetAllGradesRequest
	(*Grade)(nil),                 // 9: ccnu.v1.Grade
	(*GetAllGradesResponse)(nil),  // 10: ccnu.v1.GetAllGradesResponse
	(*GetGradesRequest)(nil),      // 11: ccnu.v1.GetGradesRequest
	(*GetGradesResponse)(nil),     // 12: ccnu.v1.GetGradesResponse
}
var file_proto_ccnu_v1_ccnu_proto_depIdxs = []int32{
	0,  // 0: ccnu.v1.CourseListRequest.source:type_name -> ccnu.v1.Source
	6,  // 1: ccnu.v1.CourseListResponse.courses:type_name -> ccnu.v1.Course
	9,  // 2: ccnu.v1.GetAllGradesResponse.grades:type_name -> ccnu.v1.Grade
	9,  // 3: ccnu.v1.GetGradesResponse.grades:type_name -> ccnu.v1.Grade
	1,  // 4: ccnu.v1.CCNUService.Login:input_type -> ccnu.v1.LoginRequest
	3,  // 5: ccnu.v1.CCNUService.GetCCNUCookie:input_type -> ccnu.v1.GetCCNUCookieRequest
	5,  // 6: ccnu.v1.CCNUService.CourseList:input_type -> ccnu.v1.CourseListRequest
	8,  // 7: ccnu.v1.CCNUService.GetAllGrades:input_type -> ccnu.v1.GetAllGradesRequest
	11, // 8: ccnu.v1.CCNUService.GetGrades:input_type -> ccnu.v1.GetGradesRequest
	2,  // 9: ccnu.v1.CCNUService.Login:output_type -> ccnu.v1.LoginResponse
	4,  // 10: ccnu.v1.CCNUService.GetCCNUCookie:output_type -> ccnu.v1.GetCCNUCookieResponse
	7,  // 11: ccnu.v1.CCNUService.CourseList:output_type -> ccnu.v1.CourseListResponse
	10, // 12: ccnu.v1.CCNUService.GetAllGrades:output_type -> ccnu.v1.GetAllGradesResponse
	12, // 13: ccnu.v1.CCNUService.GetGrades:output_type -> ccnu.v1.GetGradesResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ccnu_v1_ccnu_proto_init() }
func file_proto_ccnu_v1_ccnu_proto_init() {
	if File_proto_ccnu_v1_ccnu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ccnu_v1_ccnu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCCNUCookieRequest); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCCNUCookieResponse); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseListRequest); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseListResponse); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGradesRequest); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grade); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGradesResponse); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGradesRequest); i {
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
		file_proto_ccnu_v1_ccnu_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGradesResponse); i {
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
			RawDescriptor: file_proto_ccnu_v1_ccnu_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ccnu_v1_ccnu_proto_goTypes,
		DependencyIndexes: file_proto_ccnu_v1_ccnu_proto_depIdxs,
		EnumInfos:         file_proto_ccnu_v1_ccnu_proto_enumTypes,
		MessageInfos:      file_proto_ccnu_v1_ccnu_proto_msgTypes,
	}.Build()
	File_proto_ccnu_v1_ccnu_proto = out.File
	file_proto_ccnu_v1_ccnu_proto_rawDesc = nil
	file_proto_ccnu_v1_ccnu_proto_goTypes = nil
	file_proto_ccnu_v1_ccnu_proto_depIdxs = nil
}
