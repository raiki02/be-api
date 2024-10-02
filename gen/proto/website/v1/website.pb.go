// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/website/v1/website.proto

package websitev1

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

//web
type Website struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Link        string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Website) Reset() {
	*x = Website{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Website) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Website) ProtoMessage() {}

func (x *Website) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Website.ProtoReflect.Descriptor instead.
func (*Website) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{0}
}

func (x *Website) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Website) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Website) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Website) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Website) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetWebsitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWebsitesRequest) Reset() {
	*x = GetWebsitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebsitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebsitesRequest) ProtoMessage() {}

func (x *GetWebsitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebsitesRequest.ProtoReflect.Descriptor instead.
func (*GetWebsitesRequest) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{1}
}

type GetWebsitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Websites []*Website `protobuf:"bytes,1,rep,name=websites,proto3" json:"websites,omitempty"`
}

func (x *GetWebsitesResponse) Reset() {
	*x = GetWebsitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebsitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebsitesResponse) ProtoMessage() {}

func (x *GetWebsitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebsitesResponse.ProtoReflect.Descriptor instead.
func (*GetWebsitesResponse) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{2}
}

func (x *GetWebsitesResponse) GetWebsites() []*Website {
	if x != nil {
		return x.Websites
	}
	return nil
}

type SaveWebsiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Website *Website `protobuf:"bytes,1,opt,name=website,proto3" json:"website,omitempty"`
}

func (x *SaveWebsiteRequest) Reset() {
	*x = SaveWebsiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveWebsiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveWebsiteRequest) ProtoMessage() {}

func (x *SaveWebsiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveWebsiteRequest.ProtoReflect.Descriptor instead.
func (*SaveWebsiteRequest) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{3}
}

func (x *SaveWebsiteRequest) GetWebsite() *Website {
	if x != nil {
		return x.Website
	}
	return nil
}

type SaveWebsiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Websites []*Website `protobuf:"bytes,1,rep,name=websites,proto3" json:"websites,omitempty"`
}

func (x *SaveWebsiteResponse) Reset() {
	*x = SaveWebsiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveWebsiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveWebsiteResponse) ProtoMessage() {}

func (x *SaveWebsiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveWebsiteResponse.ProtoReflect.Descriptor instead.
func (*SaveWebsiteResponse) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{4}
}

func (x *SaveWebsiteResponse) GetWebsites() []*Website {
	if x != nil {
		return x.Websites
	}
	return nil
}

type DelWebsiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelWebsiteRequest) Reset() {
	*x = DelWebsiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelWebsiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelWebsiteRequest) ProtoMessage() {}

func (x *DelWebsiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelWebsiteRequest.ProtoReflect.Descriptor instead.
func (*DelWebsiteRequest) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{5}
}

func (x *DelWebsiteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelWebsiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Websites []*Website `protobuf:"bytes,1,rep,name=websites,proto3" json:"websites,omitempty"`
}

func (x *DelWebsiteResponse) Reset() {
	*x = DelWebsiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_website_v1_website_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelWebsiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelWebsiteResponse) ProtoMessage() {}

func (x *DelWebsiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_website_v1_website_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelWebsiteResponse.ProtoReflect.Descriptor instead.
func (*DelWebsiteResponse) Descriptor() ([]byte, []int) {
	return file_proto_website_v1_website_proto_rawDescGZIP(), []int{6}
}

func (x *DelWebsiteResponse) GetWebsites() []*Website {
	if x != nil {
		return x.Websites
	}
	return nil
}

var File_proto_website_v1_website_proto protoreflect.FileDescriptor

var file_proto_website_v1_website_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73,
	0x22, 0x43, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x07, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x52, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x45, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x52,
	0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x32, 0xfd, 0x01, 0x0a, 0x0e, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b,
	0x53, 0x61, 0x76, 0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa2, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x57, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e,
	0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x0a,
	0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0b, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_website_v1_website_proto_rawDescOnce sync.Once
	file_proto_website_v1_website_proto_rawDescData = file_proto_website_v1_website_proto_rawDesc
)

func file_proto_website_v1_website_proto_rawDescGZIP() []byte {
	file_proto_website_v1_website_proto_rawDescOnce.Do(func() {
		file_proto_website_v1_website_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_website_v1_website_proto_rawDescData)
	})
	return file_proto_website_v1_website_proto_rawDescData
}

var file_proto_website_v1_website_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_website_v1_website_proto_goTypes = []interface{}{
	(*Website)(nil),             // 0: website.v1.Website
	(*GetWebsitesRequest)(nil),  // 1: website.v1.GetWebsitesRequest
	(*GetWebsitesResponse)(nil), // 2: website.v1.GetWebsitesResponse
	(*SaveWebsiteRequest)(nil),  // 3: website.v1.SaveWebsiteRequest
	(*SaveWebsiteResponse)(nil), // 4: website.v1.SaveWebsiteResponse
	(*DelWebsiteRequest)(nil),   // 5: website.v1.DelWebsiteRequest
	(*DelWebsiteResponse)(nil),  // 6: website.v1.DelWebsiteResponse
}
var file_proto_website_v1_website_proto_depIdxs = []int32{
	0, // 0: website.v1.GetWebsitesResponse.websites:type_name -> website.v1.Website
	0, // 1: website.v1.SaveWebsiteRequest.website:type_name -> website.v1.Website
	0, // 2: website.v1.SaveWebsiteResponse.websites:type_name -> website.v1.Website
	0, // 3: website.v1.DelWebsiteResponse.websites:type_name -> website.v1.Website
	1, // 4: website.v1.WebsiteService.GetWebsites:input_type -> website.v1.GetWebsitesRequest
	3, // 5: website.v1.WebsiteService.SaveWebsite:input_type -> website.v1.SaveWebsiteRequest
	5, // 6: website.v1.WebsiteService.DelWebsite:input_type -> website.v1.DelWebsiteRequest
	2, // 7: website.v1.WebsiteService.GetWebsites:output_type -> website.v1.GetWebsitesResponse
	4, // 8: website.v1.WebsiteService.SaveWebsite:output_type -> website.v1.SaveWebsiteResponse
	6, // 9: website.v1.WebsiteService.DelWebsite:output_type -> website.v1.DelWebsiteResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_website_v1_website_proto_init() }
func file_proto_website_v1_website_proto_init() {
	if File_proto_website_v1_website_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_website_v1_website_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Website); i {
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
		file_proto_website_v1_website_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebsitesRequest); i {
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
		file_proto_website_v1_website_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebsitesResponse); i {
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
		file_proto_website_v1_website_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveWebsiteRequest); i {
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
		file_proto_website_v1_website_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveWebsiteResponse); i {
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
		file_proto_website_v1_website_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelWebsiteRequest); i {
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
		file_proto_website_v1_website_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelWebsiteResponse); i {
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
			RawDescriptor: file_proto_website_v1_website_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_website_v1_website_proto_goTypes,
		DependencyIndexes: file_proto_website_v1_website_proto_depIdxs,
		MessageInfos:      file_proto_website_v1_website_proto_msgTypes,
	}.Build()
	File_proto_website_v1_website_proto = out.File
	file_proto_website_v1_website_proto_rawDesc = nil
	file_proto_website_v1_website_proto_goTypes = nil
	file_proto_website_v1_website_proto_depIdxs = nil
}
