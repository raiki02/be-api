// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/banner/v1/banner.proto

package bannerv1

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

//banner
type Banner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WebLink     string `protobuf:"bytes,2,opt,name=webLink,proto3" json:"webLink,omitempty"`
	PictureLink string `protobuf:"bytes,3,opt,name=pictureLink,proto3" json:"pictureLink,omitempty"`
}

func (x *Banner) Reset() {
	*x = Banner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Banner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Banner) ProtoMessage() {}

func (x *Banner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Banner.ProtoReflect.Descriptor instead.
func (*Banner) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{0}
}

func (x *Banner) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Banner) GetWebLink() string {
	if x != nil {
		return x.WebLink
	}
	return ""
}

func (x *Banner) GetPictureLink() string {
	if x != nil {
		return x.PictureLink
	}
	return ""
}

type GetBannersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBannersRequest) Reset() {
	*x = GetBannersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBannersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBannersRequest) ProtoMessage() {}

func (x *GetBannersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBannersRequest.ProtoReflect.Descriptor instead.
func (*GetBannersRequest) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{1}
}

type GetBannersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banners []*Banner `protobuf:"bytes,1,rep,name=banners,proto3" json:"banners,omitempty"`
}

func (x *GetBannersResponse) Reset() {
	*x = GetBannersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBannersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBannersResponse) ProtoMessage() {}

func (x *GetBannersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBannersResponse.ProtoReflect.Descriptor instead.
func (*GetBannersResponse) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{2}
}

func (x *GetBannersResponse) GetBanners() []*Banner {
	if x != nil {
		return x.Banners
	}
	return nil
}

type SaveBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WebLink     string `protobuf:"bytes,2,opt,name=webLink,proto3" json:"webLink,omitempty"`
	PictureLink string `protobuf:"bytes,3,opt,name=pictureLink,proto3" json:"pictureLink,omitempty"`
}

func (x *SaveBannerRequest) Reset() {
	*x = SaveBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBannerRequest) ProtoMessage() {}

func (x *SaveBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBannerRequest.ProtoReflect.Descriptor instead.
func (*SaveBannerRequest) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{3}
}

func (x *SaveBannerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveBannerRequest) GetWebLink() string {
	if x != nil {
		return x.WebLink
	}
	return ""
}

func (x *SaveBannerRequest) GetPictureLink() string {
	if x != nil {
		return x.PictureLink
	}
	return ""
}

type SaveBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveBannerResponse) Reset() {
	*x = SaveBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBannerResponse) ProtoMessage() {}

func (x *SaveBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBannerResponse.ProtoReflect.Descriptor instead.
func (*SaveBannerResponse) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{4}
}

type DelBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelBannerRequest) Reset() {
	*x = DelBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBannerRequest) ProtoMessage() {}

func (x *DelBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBannerRequest.ProtoReflect.Descriptor instead.
func (*DelBannerRequest) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{5}
}

func (x *DelBannerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelBannerResponse) Reset() {
	*x = DelBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_banner_v1_banner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelBannerResponse) ProtoMessage() {}

func (x *DelBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_banner_v1_banner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelBannerResponse.ProtoReflect.Descriptor instead.
func (*DelBannerResponse) Descriptor() ([]byte, []int) {
	return file_proto_banner_v1_banner_proto_rawDescGZIP(), []int{6}
}

var File_proto_banner_v1_banner_proto protoreflect.FileDescriptor

var file_proto_banner_v1_banner_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x65, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x73, 0x22, 0x5f, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x4c, 0x69,
	0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xed, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x9a, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x73, 0x79, 0x6e, 0x63, 0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42,
	0x58, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x09, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_banner_v1_banner_proto_rawDescOnce sync.Once
	file_proto_banner_v1_banner_proto_rawDescData = file_proto_banner_v1_banner_proto_rawDesc
)

func file_proto_banner_v1_banner_proto_rawDescGZIP() []byte {
	file_proto_banner_v1_banner_proto_rawDescOnce.Do(func() {
		file_proto_banner_v1_banner_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_banner_v1_banner_proto_rawDescData)
	})
	return file_proto_banner_v1_banner_proto_rawDescData
}

var file_proto_banner_v1_banner_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_banner_v1_banner_proto_goTypes = []interface{}{
	(*Banner)(nil),             // 0: banner.v1.Banner
	(*GetBannersRequest)(nil),  // 1: banner.v1.GetBannersRequest
	(*GetBannersResponse)(nil), // 2: banner.v1.GetBannersResponse
	(*SaveBannerRequest)(nil),  // 3: banner.v1.SaveBannerRequest
	(*SaveBannerResponse)(nil), // 4: banner.v1.SaveBannerResponse
	(*DelBannerRequest)(nil),   // 5: banner.v1.DelBannerRequest
	(*DelBannerResponse)(nil),  // 6: banner.v1.DelBannerResponse
}
var file_proto_banner_v1_banner_proto_depIdxs = []int32{
	0, // 0: banner.v1.GetBannersResponse.banners:type_name -> banner.v1.Banner
	1, // 1: banner.v1.BannerService.GetBanners:input_type -> banner.v1.GetBannersRequest
	3, // 2: banner.v1.BannerService.SaveBanner:input_type -> banner.v1.SaveBannerRequest
	5, // 3: banner.v1.BannerService.DelBanner:input_type -> banner.v1.DelBannerRequest
	2, // 4: banner.v1.BannerService.GetBanners:output_type -> banner.v1.GetBannersResponse
	4, // 5: banner.v1.BannerService.SaveBanner:output_type -> banner.v1.SaveBannerResponse
	6, // 6: banner.v1.BannerService.DelBanner:output_type -> banner.v1.DelBannerResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_banner_v1_banner_proto_init() }
func file_proto_banner_v1_banner_proto_init() {
	if File_proto_banner_v1_banner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_banner_v1_banner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Banner); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBannersRequest); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBannersResponse); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBannerRequest); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBannerResponse); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelBannerRequest); i {
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
		file_proto_banner_v1_banner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelBannerResponse); i {
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
			RawDescriptor: file_proto_banner_v1_banner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_banner_v1_banner_proto_goTypes,
		DependencyIndexes: file_proto_banner_v1_banner_proto_depIdxs,
		MessageInfos:      file_proto_banner_v1_banner_proto_msgTypes,
	}.Build()
	File_proto_banner_v1_banner_proto = out.File
	file_proto_banner_v1_banner_proto_rawDesc = nil
	file_proto_banner_v1_banner_proto_goTypes = nil
	file_proto_banner_v1_banner_proto_depIdxs = nil
}
