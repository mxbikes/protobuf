// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: mod/mod.proto

package service_Mod

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Mod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description       string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ModTypeCategoryID string                 `protobuf:"bytes,4,opt,name=ModTypeCategoryID,proto3" json:"ModTypeCategoryID,omitempty"`
	ReleaseYear       int32                  `protobuf:"varint,5,opt,name=ReleaseYear,proto3" json:"ReleaseYear,omitempty"`
	CreateAt          *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=Create_at,json=CreateAt,proto3" json:"Create_at,omitempty"`
}

func (x *Mod) Reset() {
	*x = Mod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mod) ProtoMessage() {}

func (x *Mod) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mod.ProtoReflect.Descriptor instead.
func (*Mod) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{0}
}

func (x *Mod) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Mod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mod) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Mod) GetModTypeCategoryID() string {
	if x != nil {
		return x.ModTypeCategoryID
	}
	return ""
}

func (x *Mod) GetReleaseYear() int32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *Mod) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

// SearchMod
type SearchModRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchText         string   `protobuf:"bytes,1,opt,name=SearchText,proto3" json:"SearchText,omitempty"`
	ModTypeCategoryIDs []string `protobuf:"bytes,2,rep,name=ModTypeCategoryIDs,proto3" json:"ModTypeCategoryIDs,omitempty"`
	Page               int64    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size               int64    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *SearchModRequest) Reset() {
	*x = SearchModRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchModRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchModRequest) ProtoMessage() {}

func (x *SearchModRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchModRequest.ProtoReflect.Descriptor instead.
func (*SearchModRequest) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{1}
}

func (x *SearchModRequest) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

func (x *SearchModRequest) GetModTypeCategoryIDs() []string {
	if x != nil {
		return x.ModTypeCategoryIDs
	}
	return nil
}

func (x *SearchModRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchModRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SearchModResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	Mods       []*Mod      `protobuf:"bytes,2,rep,name=mods,proto3" json:"mods,omitempty"`
}

func (x *SearchModResponse) Reset() {
	*x = SearchModResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchModResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchModResponse) ProtoMessage() {}

func (x *SearchModResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchModResponse.ProtoReflect.Descriptor instead.
func (*SearchModResponse) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{2}
}

func (x *SearchModResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *SearchModResponse) GetMods() []*Mod {
	if x != nil {
		return x.Mods
	}
	return nil
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64 `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	TotalPages int64 `protobuf:"varint,2,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Page       int64 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size       int64 `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	HasMore    bool  `protobuf:"varint,5,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *Pagination) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *Pagination) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Pagination) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// GetModByID
type GetModByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetModByIDRequest) Reset() {
	*x = GetModByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModByIDRequest) ProtoMessage() {}

func (x *GetModByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModByIDRequest.ProtoReflect.Descriptor instead.
func (*GetModByIDRequest) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{4}
}

func (x *GetModByIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetModByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mod *Mod `protobuf:"bytes,1,opt,name=Mod,proto3" json:"Mod,omitempty"`
}

func (x *GetModByIDResponse) Reset() {
	*x = GetModByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mod_mod_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModByIDResponse) ProtoMessage() {}

func (x *GetModByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mod_mod_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModByIDResponse.ProtoReflect.Descriptor instead.
func (*GetModByIDResponse) Descriptor() ([]byte, []int) {
	return file_mod_mod_proto_rawDescGZIP(), []int{5}
}

func (x *GetModByIDResponse) GetMod() *Mod {
	if x != nil {
		return x.Mod
	}
	return nil
}

var File_mod_mod_proto protoreflect.FileDescriptor

var file_mod_mod_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01,
	0x0a, 0x03, 0x4d, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x4d,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4d, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x4d, 0x6f, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x4d, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x72, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48,
	0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x03, 0x4d, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64,
	0x52, 0x03, 0x4d, 0x6f, 0x64, 0x32, 0xa7, 0x01, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f,
	0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e,
	0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6d, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x69,
	0x6b, 0x56, 0x61, 0x6e, 0x48, 0x61, 0x61, 0x72, 0x65, 0x6e, 0x2f, 0x4d, 0x78, 0x42, 0x69, 0x6b,
	0x65, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mod_mod_proto_rawDescOnce sync.Once
	file_mod_mod_proto_rawDescData = file_mod_mod_proto_rawDesc
)

func file_mod_mod_proto_rawDescGZIP() []byte {
	file_mod_mod_proto_rawDescOnce.Do(func() {
		file_mod_mod_proto_rawDescData = protoimpl.X.CompressGZIP(file_mod_mod_proto_rawDescData)
	})
	return file_mod_mod_proto_rawDescData
}

var file_mod_mod_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mod_mod_proto_goTypes = []interface{}{
	(*Mod)(nil),                   // 0: mod_service.Mod
	(*SearchModRequest)(nil),      // 1: mod_service.SearchModRequest
	(*SearchModResponse)(nil),     // 2: mod_service.SearchModResponse
	(*Pagination)(nil),            // 3: mod_service.Pagination
	(*GetModByIDRequest)(nil),     // 4: mod_service.GetModByIDRequest
	(*GetModByIDResponse)(nil),    // 5: mod_service.GetModByIDResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_mod_mod_proto_depIdxs = []int32{
	6, // 0: mod_service.Mod.Create_at:type_name -> google.protobuf.Timestamp
	3, // 1: mod_service.SearchModResponse.Pagination:type_name -> mod_service.Pagination
	0, // 2: mod_service.SearchModResponse.mods:type_name -> mod_service.Mod
	0, // 3: mod_service.GetModByIDResponse.Mod:type_name -> mod_service.Mod
	1, // 4: mod_service.ModService.SearchMod:input_type -> mod_service.SearchModRequest
	4, // 5: mod_service.ModService.GetModByID:input_type -> mod_service.GetModByIDRequest
	2, // 6: mod_service.ModService.SearchMod:output_type -> mod_service.SearchModResponse
	5, // 7: mod_service.ModService.GetModByID:output_type -> mod_service.GetModByIDResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mod_mod_proto_init() }
func file_mod_mod_proto_init() {
	if File_mod_mod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mod_mod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mod); i {
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
		file_mod_mod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchModRequest); i {
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
		file_mod_mod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchModResponse); i {
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
		file_mod_mod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_mod_mod_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModByIDRequest); i {
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
		file_mod_mod_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModByIDResponse); i {
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
			RawDescriptor: file_mod_mod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mod_mod_proto_goTypes,
		DependencyIndexes: file_mod_mod_proto_depIdxs,
		MessageInfos:      file_mod_mod_proto_msgTypes,
	}.Build()
	File_mod_mod_proto = out.File
	file_mod_mod_proto_rawDesc = nil
	file_mod_mod_proto_goTypes = nil
	file_mod_mod_proto_depIdxs = nil
}