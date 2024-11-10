// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/classService/v1/classService.proto

package classServicev1

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

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//搜索关键词,匹配的是课程名称和教师姓名
	SearchKeyWords string `protobuf:"bytes,1,opt,name=searchKeyWords,proto3" json:"searchKeyWords,omitempty"`
	Year           string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Semester       string `protobuf:"bytes,3,opt,name=semester,proto3" json:"semester,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_classService_v1_classService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_classService_v1_classService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_classService_v1_classService_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetSearchKeyWords() string {
	if x != nil {
		return x.SearchKeyWords
	}
	return ""
}

func (x *SearchRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *SearchRequest) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//课程信息
	ClassInfos []*ClassInfo `protobuf:"bytes,1,rep,name=class_infos,json=classInfos,proto3" json:"class_infos,omitempty"`
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_classService_v1_classService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_classService_v1_classService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReply.ProtoReflect.Descriptor instead.
func (*SearchReply) Descriptor() ([]byte, []int) {
	return file_proto_classService_v1_classService_proto_rawDescGZIP(), []int{1}
}

func (x *SearchReply) GetClassInfos() []*ClassInfo {
	if x != nil {
		return x.ClassInfos
	}
	return nil
}

type AddClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//学号
	StuId string `protobuf:"bytes,1,opt,name=stu_id,json=stuId,proto3" json:"stu_id,omitempty"`
	//课程名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//第几节 '形如 "1-3","1-1"'
	DurClass string `protobuf:"bytes,3,opt,name=dur_class,json=durClass,proto3" json:"dur_class,omitempty"`
	//地点
	Where string `protobuf:"bytes,4,opt,name=where,proto3" json:"where,omitempty"`
	//教师
	Teacher string `protobuf:"bytes,5,opt,name=teacher,proto3" json:"teacher,omitempty"`
	//哪些周
	Weeks int64 `protobuf:"varint,6,opt,name=weeks,proto3" json:"weeks,omitempty"`
	// 学期
	Semester string `protobuf:"bytes,7,opt,name=semester,proto3" json:"semester,omitempty"`
	//学年
	Year string `protobuf:"bytes,8,opt,name=year,proto3" json:"year,omitempty"`
	//星期几
	Day int64 `protobuf:"varint,9,opt,name=day,proto3" json:"day,omitempty"`
	//学分
	Credit *float64 `protobuf:"fixed64,10,opt,name=credit,proto3,oneof" json:"credit,omitempty"`
}

func (x *AddClassRequest) Reset() {
	*x = AddClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_classService_v1_classService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClassRequest) ProtoMessage() {}

func (x *AddClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_classService_v1_classService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClassRequest.ProtoReflect.Descriptor instead.
func (*AddClassRequest) Descriptor() ([]byte, []int) {
	return file_proto_classService_v1_classService_proto_rawDescGZIP(), []int{2}
}

func (x *AddClassRequest) GetStuId() string {
	if x != nil {
		return x.StuId
	}
	return ""
}

func (x *AddClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddClassRequest) GetDurClass() string {
	if x != nil {
		return x.DurClass
	}
	return ""
}

func (x *AddClassRequest) GetWhere() string {
	if x != nil {
		return x.Where
	}
	return ""
}

func (x *AddClassRequest) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *AddClassRequest) GetWeeks() int64 {
	if x != nil {
		return x.Weeks
	}
	return 0
}

func (x *AddClassRequest) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *AddClassRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *AddClassRequest) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *AddClassRequest) GetCredit() float64 {
	if x != nil && x.Credit != nil {
		return *x.Credit
	}
	return 0
}

type AddClassReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//添加的课程ID
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddClassReply) Reset() {
	*x = AddClassReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_classService_v1_classService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClassReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClassReply) ProtoMessage() {}

func (x *AddClassReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_classService_v1_classService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClassReply.ProtoReflect.Descriptor instead.
func (*AddClassReply) Descriptor() ([]byte, []int) {
	return file_proto_classService_v1_classService_proto_rawDescGZIP(), []int{3}
}

func (x *AddClassReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddClassReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ClassInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//星期几
	Day int64 `protobuf:"varint,1,opt,name=day,proto3" json:"day,omitempty"`
	//任课教师
	Teacher string `protobuf:"bytes,2,opt,name=teacher,proto3" json:"teacher,omitempty"`
	//上课地点
	Where string `protobuf:"bytes,3,opt,name=where,proto3" json:"where,omitempty"`
	//上课是第几节（如1-2,3,4）
	ClassWhen string `protobuf:"bytes,4,opt,name=class_when,json=classWhen,proto3" json:"class_when,omitempty"`
	//上课的周数(文字描述,如1-9周)
	WeekDuration string `protobuf:"bytes,5,opt,name=week_duration,json=weekDuration,proto3" json:"week_duration,omitempty"`
	//课程名称
	Classname string `protobuf:"bytes,6,opt,name=classname,proto3" json:"classname,omitempty"`
	//学分
	Credit float64 `protobuf:"fixed64,7,opt,name=credit,proto3" json:"credit,omitempty"`
	//哪些周 这个是一个64位的数字,如果有第一周,那么该数的二进制从右往左的第一位为1,以此类推
	//比如该数的二进制是000000101,就代表第一周和第三周有课.
	Weeks int64 `protobuf:"varint,8,opt,name=weeks,proto3" json:"weeks,omitempty"`
	//学期 "1"代表第一学期，"2"代表第二学期，"3"代表第三学期
	Semester string `protobuf:"bytes,9,opt,name=semester,proto3" json:"semester,omitempty"`
	//学年  "2024" 代表"2024-2025学年"
	Year string `protobuf:"bytes,10,opt,name=year,proto3" json:"year,omitempty"`
	//课程唯一标识id
	Id string `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClassInfo) Reset() {
	*x = ClassInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_classService_v1_classService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassInfo) ProtoMessage() {}

func (x *ClassInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_classService_v1_classService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassInfo.ProtoReflect.Descriptor instead.
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return file_proto_classService_v1_classService_proto_rawDescGZIP(), []int{4}
}

func (x *ClassInfo) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *ClassInfo) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *ClassInfo) GetWhere() string {
	if x != nil {
		return x.Where
	}
	return ""
}

func (x *ClassInfo) GetClassWhen() string {
	if x != nil {
		return x.ClassWhen
	}
	return ""
}

func (x *ClassInfo) GetWeekDuration() string {
	if x != nil {
		return x.WeekDuration
	}
	return ""
}

func (x *ClassInfo) GetClassname() string {
	if x != nil {
		return x.Classname
	}
	return ""
}

func (x *ClassInfo) GetCredit() float64 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *ClassInfo) GetWeeks() int64 {
	if x != nil {
		return x.Weeks
	}
	return 0
}

func (x *ClassInfo) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *ClassInfo) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ClassInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_classService_v1_classService_proto protoreflect.FileDescriptor

var file_proto_classService_v1_classService_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x4a, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x74, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x75, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x75,
	0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x1b, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x22, 0x31, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x9d, 0x02, 0x0a, 0x09, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x57, 0x68, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x65,
	0x6b, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x77, 0x65, 0x65, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa9, 0x01, 0x0a, 0x0c, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0xca, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x11,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x10, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_classService_v1_classService_proto_rawDescOnce sync.Once
	file_proto_classService_v1_classService_proto_rawDescData = file_proto_classService_v1_classService_proto_rawDesc
)

func file_proto_classService_v1_classService_proto_rawDescGZIP() []byte {
	file_proto_classService_v1_classService_proto_rawDescOnce.Do(func() {
		file_proto_classService_v1_classService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_classService_v1_classService_proto_rawDescData)
	})
	return file_proto_classService_v1_classService_proto_rawDescData
}

var file_proto_classService_v1_classService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_classService_v1_classService_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),   // 0: classService.v1.SearchRequest
	(*SearchReply)(nil),     // 1: classService.v1.SearchReply
	(*AddClassRequest)(nil), // 2: classService.v1.AddClassRequest
	(*AddClassReply)(nil),   // 3: classService.v1.AddClassReply
	(*ClassInfo)(nil),       // 4: classService.v1.ClassInfo
}
var file_proto_classService_v1_classService_proto_depIdxs = []int32{
	4, // 0: classService.v1.SearchReply.class_infos:type_name -> classService.v1.ClassInfo
	0, // 1: classService.v1.ClassService.SearchClass:input_type -> classService.v1.SearchRequest
	2, // 2: classService.v1.ClassService.AddClass:input_type -> classService.v1.AddClassRequest
	1, // 3: classService.v1.ClassService.SearchClass:output_type -> classService.v1.SearchReply
	3, // 4: classService.v1.ClassService.AddClass:output_type -> classService.v1.AddClassReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_classService_v1_classService_proto_init() }
func file_proto_classService_v1_classService_proto_init() {
	if File_proto_classService_v1_classService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_classService_v1_classService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_proto_classService_v1_classService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReply); i {
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
		file_proto_classService_v1_classService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClassRequest); i {
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
		file_proto_classService_v1_classService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClassReply); i {
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
		file_proto_classService_v1_classService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassInfo); i {
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
	file_proto_classService_v1_classService_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_classService_v1_classService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_classService_v1_classService_proto_goTypes,
		DependencyIndexes: file_proto_classService_v1_classService_proto_depIdxs,
		MessageInfos:      file_proto_classService_v1_classService_proto_msgTypes,
	}.Build()
	File_proto_classService_v1_classService_proto = out.File
	file_proto_classService_v1_classService_proto_rawDesc = nil
	file_proto_classService_v1_classService_proto_goTypes = nil
	file_proto_classService_v1_classService_proto_depIdxs = nil
}
