// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/answer/v1/answer.proto

package answerv1

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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublisherId int64  `protobuf:"varint,1,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	QuestionId  int64  `protobuf:"varint,2,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetPublisherId() int64 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *PublishRequest) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *PublishRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{1}
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnswerId int64 `protobuf:"varint,1,opt,name=answer_id,json=answerId,proto3" json:"answer_id,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{2}
}

func (x *DetailRequest) GetAnswerId() int64 {
	if x != nil {
		return x.AnswerId
	}
	return 0
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PublisherId int64  `protobuf:"varint,2,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	QuestionId  int64  `protobuf:"varint,3,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{3}
}

func (x *Answer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Answer) GetPublisherId() int64 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *Answer) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Answer) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer *Answer `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{4}
}

func (x *DetailResponse) GetAnswer() *Answer {
	if x != nil {
		return x.Answer
	}
	return nil
}

type ListForQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId  int64 `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	CurAnswerId int64 `protobuf:"varint,2,opt,name=cur_answer_id,json=curAnswerId,proto3" json:"cur_answer_id,omitempty"`
	Limit       int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListForQuestionRequest) Reset() {
	*x = ListForQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForQuestionRequest) ProtoMessage() {}

func (x *ListForQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForQuestionRequest.ProtoReflect.Descriptor instead.
func (*ListForQuestionRequest) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{5}
}

func (x *ListForQuestionRequest) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *ListForQuestionRequest) GetCurAnswerId() int64 {
	if x != nil {
		return x.CurAnswerId
	}
	return 0
}

func (x *ListForQuestionRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListForQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*Answer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *ListForQuestionResponse) Reset() {
	*x = ListForQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForQuestionResponse) ProtoMessage() {}

func (x *ListForQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForQuestionResponse.ProtoReflect.Descriptor instead.
func (*ListForQuestionResponse) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{6}
}

func (x *ListForQuestionResponse) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type ListForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CurAnswerId int64 `protobuf:"varint,2,opt,name=cur_answer_id,json=curAnswerId,proto3" json:"cur_answer_id,omitempty"`
	Limit       int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListForUserRequest) Reset() {
	*x = ListForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForUserRequest) ProtoMessage() {}

func (x *ListForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForUserRequest.ProtoReflect.Descriptor instead.
func (*ListForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{7}
}

func (x *ListForUserRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListForUserRequest) GetCurAnswerId() int64 {
	if x != nil {
		return x.CurAnswerId
	}
	return 0
}

func (x *ListForUserRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answers []*Answer `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *ListForUserResponse) Reset() {
	*x = ListForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListForUserResponse) ProtoMessage() {}

func (x *ListForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListForUserResponse.ProtoReflect.Descriptor instead.
func (*ListForUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{8}
}

func (x *ListForUserResponse) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type CountForQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId int64 `protobuf:"varint,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
}

func (x *CountForQuestionRequest) Reset() {
	*x = CountForQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountForQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountForQuestionRequest) ProtoMessage() {}

func (x *CountForQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountForQuestionRequest.ProtoReflect.Descriptor instead.
func (*CountForQuestionRequest) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{9}
}

func (x *CountForQuestionRequest) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

type CountForQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cnt int64 `protobuf:"varint,1,opt,name=cnt,proto3" json:"cnt,omitempty"`
}

func (x *CountForQuestionResponse) Reset() {
	*x = CountForQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_answer_v1_answer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountForQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountForQuestionResponse) ProtoMessage() {}

func (x *CountForQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_answer_v1_answer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountForQuestionResponse.ProtoReflect.Descriptor instead.
func (*CountForQuestionResponse) Descriptor() ([]byte, []int) {
	return file_proto_answer_v1_answer_proto_rawDescGZIP(), []int{10}
}

func (x *CountForQuestionResponse) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

var File_proto_answer_v1_answer_proto protoreflect.FileDescriptor

var file_proto_answer_v1_answer_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x6e, 0x0a, 0x0e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0d,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x06, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22,
	0x73, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x46, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x5f, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x42,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2c,
	0x0a, 0x18, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x32, 0x95, 0x03, 0x0a,
	0x0d, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x58, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x75, 0x78, 0x69, 0x4b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x62, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_answer_v1_answer_proto_rawDescOnce sync.Once
	file_proto_answer_v1_answer_proto_rawDescData = file_proto_answer_v1_answer_proto_rawDesc
)

func file_proto_answer_v1_answer_proto_rawDescGZIP() []byte {
	file_proto_answer_v1_answer_proto_rawDescOnce.Do(func() {
		file_proto_answer_v1_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_answer_v1_answer_proto_rawDescData)
	})
	return file_proto_answer_v1_answer_proto_rawDescData
}

var file_proto_answer_v1_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_answer_v1_answer_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),           // 0: answer.v1.PublishRequest
	(*PublishResponse)(nil),          // 1: answer.v1.PublishResponse
	(*DetailRequest)(nil),            // 2: answer.v1.DetailRequest
	(*Answer)(nil),                   // 3: answer.v1.Answer
	(*DetailResponse)(nil),           // 4: answer.v1.DetailResponse
	(*ListForQuestionRequest)(nil),   // 5: answer.v1.ListForQuestionRequest
	(*ListForQuestionResponse)(nil),  // 6: answer.v1.ListForQuestionResponse
	(*ListForUserRequest)(nil),       // 7: answer.v1.ListForUserRequest
	(*ListForUserResponse)(nil),      // 8: answer.v1.ListForUserResponse
	(*CountForQuestionRequest)(nil),  // 9: answer.v1.CountForQuestionRequest
	(*CountForQuestionResponse)(nil), // 10: answer.v1.CountForQuestionResponse
}
var file_proto_answer_v1_answer_proto_depIdxs = []int32{
	3,  // 0: answer.v1.DetailResponse.answer:type_name -> answer.v1.Answer
	3,  // 1: answer.v1.ListForQuestionResponse.answers:type_name -> answer.v1.Answer
	3,  // 2: answer.v1.ListForUserResponse.answers:type_name -> answer.v1.Answer
	0,  // 3: answer.v1.AnswerService.Publish:input_type -> answer.v1.PublishRequest
	2,  // 4: answer.v1.AnswerService.Detail:input_type -> answer.v1.DetailRequest
	5,  // 5: answer.v1.AnswerService.ListForQuestion:input_type -> answer.v1.ListForQuestionRequest
	7,  // 6: answer.v1.AnswerService.ListForUser:input_type -> answer.v1.ListForUserRequest
	9,  // 7: answer.v1.AnswerService.CountForQuestion:input_type -> answer.v1.CountForQuestionRequest
	1,  // 8: answer.v1.AnswerService.Publish:output_type -> answer.v1.PublishResponse
	4,  // 9: answer.v1.AnswerService.Detail:output_type -> answer.v1.DetailResponse
	6,  // 10: answer.v1.AnswerService.ListForQuestion:output_type -> answer.v1.ListForQuestionResponse
	8,  // 11: answer.v1.AnswerService.ListForUser:output_type -> answer.v1.ListForUserResponse
	10, // 12: answer.v1.AnswerService.CountForQuestion:output_type -> answer.v1.CountForQuestionResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_answer_v1_answer_proto_init() }
func file_proto_answer_v1_answer_proto_init() {
	if File_proto_answer_v1_answer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_answer_v1_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForQuestionRequest); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForQuestionResponse); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForUserRequest); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListForUserResponse); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountForQuestionRequest); i {
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
		file_proto_answer_v1_answer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountForQuestionResponse); i {
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
			RawDescriptor: file_proto_answer_v1_answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_answer_v1_answer_proto_goTypes,
		DependencyIndexes: file_proto_answer_v1_answer_proto_depIdxs,
		MessageInfos:      file_proto_answer_v1_answer_proto_msgTypes,
	}.Build()
	File_proto_answer_v1_answer_proto = out.File
	file_proto_answer_v1_answer_proto_rawDesc = nil
	file_proto_answer_v1_answer_proto_goTypes = nil
	file_proto_answer_v1_answer_proto_depIdxs = nil
}
