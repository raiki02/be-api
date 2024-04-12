// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/tag/v1/tag.proto

package tagv1

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

// 因为标签是固定的，写死，所以枚举，规范来讲应该都大写吧，但是难读...
type AssessmentTag int32

const (
	AssessmentTag_UnknownAssessment     AssessmentTag = 0 // 未知考核
	AssessmentTag_OpenBookExamination   AssessmentTag = 1 // 开卷考试
	AssessmentTag_ClosedBookExamination AssessmentTag = 2 // 闭卷考试
	AssessmentTag_ThesisExamination     AssessmentTag = 3 // 论文考核
	AssessmentTag_GroupReporting        AssessmentTag = 4 // 小组汇报
	AssessmentTag_NoAssessment          AssessmentTag = 5 // 无考核
)

// Enum value maps for AssessmentTag.
var (
	AssessmentTag_name = map[int32]string{
		0: "UnknownAssessment",
		1: "OpenBookExamination",
		2: "ClosedBookExamination",
		3: "ThesisExamination",
		4: "GroupReporting",
		5: "NoAssessment",
	}
	AssessmentTag_value = map[string]int32{
		"UnknownAssessment":     0,
		"OpenBookExamination":   1,
		"ClosedBookExamination": 2,
		"ThesisExamination":     3,
		"GroupReporting":        4,
		"NoAssessment":          5,
	}
)

func (x AssessmentTag) Enum() *AssessmentTag {
	p := new(AssessmentTag)
	*p = x
	return p
}

func (x AssessmentTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssessmentTag) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tag_v1_tag_proto_enumTypes[0].Descriptor()
}

func (AssessmentTag) Type() protoreflect.EnumType {
	return &file_proto_tag_v1_tag_proto_enumTypes[0]
}

func (x AssessmentTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssessmentTag.Descriptor instead.
func (AssessmentTag) EnumDescriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{0}
}

type Biz int32

const (
	Biz_Unknown Biz = 0
	Biz_Course  Biz = 1
)

// Enum value maps for Biz.
var (
	Biz_name = map[int32]string{
		0: "Unknown",
		1: "Course",
	}
	Biz_value = map[string]int32{
		"Unknown": 0,
		"Course":  1,
	}
)

func (x Biz) Enum() *Biz {
	p := new(Biz)
	*p = x
	return p
}

func (x Biz) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Biz) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tag_v1_tag_proto_enumTypes[1].Descriptor()
}

func (Biz) Type() protoreflect.EnumType {
	return &file_proto_tag_v1_tag_proto_enumTypes[1]
}

func (x Biz) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Biz.Descriptor instead.
func (Biz) EnumDescriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{1}
}

type FeatureTag int32

const (
	FeatureTag_UnknownFeature               FeatureTag = 0 // 未知特点
	FeatureTag_EasyToLearn                  FeatureTag = 1 // 课程简单易学
	FeatureTag_RichInContent                FeatureTag = 2 // 课程干货满满
	FeatureTag_Challenging                  FeatureTag = 3 // 课程很有挑战
	FeatureTag_RigorousAndResponsible       FeatureTag = 4 // 老师严谨负责
	FeatureTag_KindAndEasygoing             FeatureTag = 5 // 老师温柔随和
	FeatureTag_Humorous                     FeatureTag = 6 // 老师风趣幽默
	FeatureTag_LessHomework                 FeatureTag = 7 // 平时作业少
	FeatureTag_KeyPointsForFinal            FeatureTag = 8 // 期末划重点
	FeatureTag_ComprehensiveOnlineMaterials FeatureTag = 9 // 云课堂资料全
)

// Enum value maps for FeatureTag.
var (
	FeatureTag_name = map[int32]string{
		0: "UnknownFeature",
		1: "EasyToLearn",
		2: "RichInContent",
		3: "Challenging",
		4: "RigorousAndResponsible",
		5: "KindAndEasygoing",
		6: "Humorous",
		7: "LessHomework",
		8: "KeyPointsForFinal",
		9: "ComprehensiveOnlineMaterials",
	}
	FeatureTag_value = map[string]int32{
		"UnknownFeature":               0,
		"EasyToLearn":                  1,
		"RichInContent":                2,
		"Challenging":                  3,
		"RigorousAndResponsible":       4,
		"KindAndEasygoing":             5,
		"Humorous":                     6,
		"LessHomework":                 7,
		"KeyPointsForFinal":            8,
		"ComprehensiveOnlineMaterials": 9,
	}
)

func (x FeatureTag) Enum() *FeatureTag {
	p := new(FeatureTag)
	*p = x
	return p
}

func (x FeatureTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureTag) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tag_v1_tag_proto_enumTypes[2].Descriptor()
}

func (FeatureTag) Type() protoreflect.EnumType {
	return &file_proto_tag_v1_tag_proto_enumTypes[2]
}

func (x FeatureTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureTag.Descriptor instead.
func (FeatureTag) EnumDescriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{2}
}

type AttachAssessmentTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaggerId int64           `protobuf:"varint,1,opt,name=tagger_id,json=taggerId,proto3" json:"tagger_id,omitempty"`
	Biz      Biz             `protobuf:"varint,2,opt,name=biz,proto3,enum=tag.v1.Biz" json:"biz,omitempty"`
	BizId    int64           `protobuf:"varint,3,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	Tags     []AssessmentTag `protobuf:"varint,4,rep,packed,name=tags,proto3,enum=tag.v1.AssessmentTag" json:"tags,omitempty"`
}

func (x *AttachAssessmentTagsRequest) Reset() {
	*x = AttachAssessmentTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachAssessmentTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachAssessmentTagsRequest) ProtoMessage() {}

func (x *AttachAssessmentTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachAssessmentTagsRequest.ProtoReflect.Descriptor instead.
func (*AttachAssessmentTagsRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *AttachAssessmentTagsRequest) GetTaggerId() int64 {
	if x != nil {
		return x.TaggerId
	}
	return 0
}

func (x *AttachAssessmentTagsRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Unknown
}

func (x *AttachAssessmentTagsRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *AttachAssessmentTagsRequest) GetTags() []AssessmentTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AttachAssessmentTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttachAssessmentTagsResponse) Reset() {
	*x = AttachAssessmentTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachAssessmentTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachAssessmentTagsResponse) ProtoMessage() {}

func (x *AttachAssessmentTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachAssessmentTagsResponse.ProtoReflect.Descriptor instead.
func (*AttachAssessmentTagsResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{1}
}

type AttachFeatureTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaggerId int64        `protobuf:"varint,1,opt,name=tagger_id,json=taggerId,proto3" json:"tagger_id,omitempty"`
	Biz      Biz          `protobuf:"varint,2,opt,name=biz,proto3,enum=tag.v1.Biz" json:"biz,omitempty"`
	BizId    int64        `protobuf:"varint,3,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	Tags     []FeatureTag `protobuf:"varint,4,rep,packed,name=tags,proto3,enum=tag.v1.FeatureTag" json:"tags,omitempty"`
}

func (x *AttachFeatureTagsRequest) Reset() {
	*x = AttachFeatureTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_v1_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachFeatureTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachFeatureTagsRequest) ProtoMessage() {}

func (x *AttachFeatureTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_v1_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachFeatureTagsRequest.ProtoReflect.Descriptor instead.
func (*AttachFeatureTagsRequest) Descriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{2}
}

func (x *AttachFeatureTagsRequest) GetTaggerId() int64 {
	if x != nil {
		return x.TaggerId
	}
	return 0
}

func (x *AttachFeatureTagsRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Unknown
}

func (x *AttachFeatureTagsRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *AttachFeatureTagsRequest) GetTags() []FeatureTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AttachFeatureTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttachFeatureTagsResponse) Reset() {
	*x = AttachFeatureTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tag_v1_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachFeatureTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachFeatureTagsResponse) ProtoMessage() {}

func (x *AttachFeatureTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tag_v1_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachFeatureTagsResponse.ProtoReflect.Descriptor instead.
func (*AttachFeatureTagsResponse) Descriptor() ([]byte, []int) {
	return file_proto_tag_v1_tag_proto_rawDescGZIP(), []int{3}
}

var File_proto_tag_v1_tag_proto protoreflect.FileDescriptor

var file_proto_tag_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x22, 0x9b, 0x01, 0x0a, 0x1b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x41, 0x73, 0x73, 0x65, 0x73,
	0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x74, 0x61, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06,
	0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69,
	0x7a, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x73,
	0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x1e,
	0x0a, 0x1c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95,
	0x01, 0x0a, 0x18, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x67,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2a, 0x97, 0x01, 0x0a, 0x0d, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x4f, 0x70, 0x65, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x54, 0x68, 0x65, 0x73, 0x69, 0x73, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x6f, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x05, 0x2a, 0x1e, 0x0a,
	0x03, 0x42, 0x69, 0x7a, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x10, 0x01, 0x2a, 0xe0, 0x01,
	0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x0e,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x61, 0x73, 0x79, 0x54, 0x6f, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x69, 0x63, 0x68, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x69, 0x67, 0x6f, 0x72, 0x6f, 0x75,
	0x73, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x10,
	0x04, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x69, 0x6e, 0x64, 0x41, 0x6e, 0x64, 0x45, 0x61, 0x73, 0x79,
	0x67, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x75, 0x6d, 0x6f, 0x72,
	0x6f, 0x75, 0x73, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x08, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x68, 0x65, 0x6e, 0x73, 0x69, 0x76, 0x65, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x10, 0x09,
	0x32, 0xc9, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x61, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74,
	0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x41, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x61, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x63, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x54, 0x61, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61,
	0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58,
	0xaa, 0x02, 0x06, 0x54, 0x61, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x54, 0x61, 0x67, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x12, 0x54, 0x61, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x54, 0x61, 0x67, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tag_v1_tag_proto_rawDescOnce sync.Once
	file_proto_tag_v1_tag_proto_rawDescData = file_proto_tag_v1_tag_proto_rawDesc
)

func file_proto_tag_v1_tag_proto_rawDescGZIP() []byte {
	file_proto_tag_v1_tag_proto_rawDescOnce.Do(func() {
		file_proto_tag_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tag_v1_tag_proto_rawDescData)
	})
	return file_proto_tag_v1_tag_proto_rawDescData
}

var file_proto_tag_v1_tag_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_tag_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_tag_v1_tag_proto_goTypes = []interface{}{
	(AssessmentTag)(0),                   // 0: tag.v1.AssessmentTag
	(Biz)(0),                             // 1: tag.v1.Biz
	(FeatureTag)(0),                      // 2: tag.v1.FeatureTag
	(*AttachAssessmentTagsRequest)(nil),  // 3: tag.v1.AttachAssessmentTagsRequest
	(*AttachAssessmentTagsResponse)(nil), // 4: tag.v1.AttachAssessmentTagsResponse
	(*AttachFeatureTagsRequest)(nil),     // 5: tag.v1.AttachFeatureTagsRequest
	(*AttachFeatureTagsResponse)(nil),    // 6: tag.v1.AttachFeatureTagsResponse
}
var file_proto_tag_v1_tag_proto_depIdxs = []int32{
	1, // 0: tag.v1.AttachAssessmentTagsRequest.biz:type_name -> tag.v1.Biz
	0, // 1: tag.v1.AttachAssessmentTagsRequest.tags:type_name -> tag.v1.AssessmentTag
	1, // 2: tag.v1.AttachFeatureTagsRequest.biz:type_name -> tag.v1.Biz
	2, // 3: tag.v1.AttachFeatureTagsRequest.tags:type_name -> tag.v1.FeatureTag
	3, // 4: tag.v1.TagService.AttachAssessmentTags:input_type -> tag.v1.AttachAssessmentTagsRequest
	5, // 5: tag.v1.TagService.AttachFeatureTags:input_type -> tag.v1.AttachFeatureTagsRequest
	4, // 6: tag.v1.TagService.AttachAssessmentTags:output_type -> tag.v1.AttachAssessmentTagsResponse
	6, // 7: tag.v1.TagService.AttachFeatureTags:output_type -> tag.v1.AttachFeatureTagsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_tag_v1_tag_proto_init() }
func file_proto_tag_v1_tag_proto_init() {
	if File_proto_tag_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tag_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachAssessmentTagsRequest); i {
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
		file_proto_tag_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachAssessmentTagsResponse); i {
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
		file_proto_tag_v1_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachFeatureTagsRequest); i {
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
		file_proto_tag_v1_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachFeatureTagsResponse); i {
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
			RawDescriptor: file_proto_tag_v1_tag_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tag_v1_tag_proto_goTypes,
		DependencyIndexes: file_proto_tag_v1_tag_proto_depIdxs,
		EnumInfos:         file_proto_tag_v1_tag_proto_enumTypes,
		MessageInfos:      file_proto_tag_v1_tag_proto_msgTypes,
	}.Build()
	File_proto_tag_v1_tag_proto = out.File
	file_proto_tag_v1_tag_proto_rawDesc = nil
	file_proto_tag_v1_tag_proto_goTypes = nil
	file_proto_tag_v1_tag_proto_depIdxs = nil
}
