// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: modImage/modImage.proto

package modImage

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

type ModImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Bucket string `protobuf:"bytes,2,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Url    string `protobuf:"bytes,4,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *ModImage) Reset() {
	*x = ModImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modImage_modImage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModImage) ProtoMessage() {}

func (x *ModImage) ProtoReflect() protoreflect.Message {
	mi := &file_modImage_modImage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModImage.ProtoReflect.Descriptor instead.
func (*ModImage) Descriptor() ([]byte, []int) {
	return file_modImage_modImage_proto_rawDescGZIP(), []int{0}
}

func (x *ModImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModImage) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *ModImage) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modImage_modImage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_modImage_modImage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_modImage_modImage_proto_rawDescGZIP(), []int{1}
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// GetModImagesByModID
type GetModImagesByModIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModID string `protobuf:"bytes,1,opt,name=ModID,proto3" json:"ModID,omitempty"`
}

func (x *GetModImagesByModIDRequest) Reset() {
	*x = GetModImagesByModIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modImage_modImage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModImagesByModIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModImagesByModIDRequest) ProtoMessage() {}

func (x *GetModImagesByModIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modImage_modImage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModImagesByModIDRequest.ProtoReflect.Descriptor instead.
func (*GetModImagesByModIDRequest) Descriptor() ([]byte, []int) {
	return file_modImage_modImage_proto_rawDescGZIP(), []int{2}
}

func (x *GetModImagesByModIDRequest) GetModID() string {
	if x != nil {
		return x.ModID
	}
	return ""
}

type GetModImagesByModIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModImages []*ModImage `protobuf:"bytes,2,rep,name=ModImages,proto3" json:"ModImages,omitempty"`
}

func (x *GetModImagesByModIDResponse) Reset() {
	*x = GetModImagesByModIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modImage_modImage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModImagesByModIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModImagesByModIDResponse) ProtoMessage() {}

func (x *GetModImagesByModIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modImage_modImage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModImagesByModIDResponse.ProtoReflect.Descriptor instead.
func (*GetModImagesByModIDResponse) Descriptor() ([]byte, []int) {
	return file_modImage_modImage_proto_rawDescGZIP(), []int{3}
}

func (x *GetModImagesByModIDResponse) GetModImages() []*ModImage {
	if x != nil {
		return x.ModImages
	}
	return nil
}

var File_modImage_modImage_proto protoreflect.FileDescriptor

var file_modImage_modImage_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x6f, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x4d,
	0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x72, 0x6c, 0x22, 0x2d, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x6f, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x32, 0x85, 0x01, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x2c, 0x2e, 0x6d,
	0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f,
	0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x6f, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x78, 0x62, 0x69, 0x6b, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modImage_modImage_proto_rawDescOnce sync.Once
	file_modImage_modImage_proto_rawDescData = file_modImage_modImage_proto_rawDesc
)

func file_modImage_modImage_proto_rawDescGZIP() []byte {
	file_modImage_modImage_proto_rawDescOnce.Do(func() {
		file_modImage_modImage_proto_rawDescData = protoimpl.X.CompressGZIP(file_modImage_modImage_proto_rawDescData)
	})
	return file_modImage_modImage_proto_rawDescData
}

var file_modImage_modImage_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_modImage_modImage_proto_goTypes = []interface{}{
	(*ModImage)(nil),                    // 0: modImage_service.ModImage
	(*Tag)(nil),                         // 1: modImage_service.Tag
	(*GetModImagesByModIDRequest)(nil),  // 2: modImage_service.GetModImagesByModIDRequest
	(*GetModImagesByModIDResponse)(nil), // 3: modImage_service.GetModImagesByModIDResponse
}
var file_modImage_modImage_proto_depIdxs = []int32{
	0, // 0: modImage_service.GetModImagesByModIDResponse.ModImages:type_name -> modImage_service.ModImage
	2, // 1: modImage_service.ModImageService.GetModImagesByModID:input_type -> modImage_service.GetModImagesByModIDRequest
	3, // 2: modImage_service.ModImageService.GetModImagesByModID:output_type -> modImage_service.GetModImagesByModIDResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modImage_modImage_proto_init() }
func file_modImage_modImage_proto_init() {
	if File_modImage_modImage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modImage_modImage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModImage); i {
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
		file_modImage_modImage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_modImage_modImage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModImagesByModIDRequest); i {
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
		file_modImage_modImage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModImagesByModIDResponse); i {
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
			RawDescriptor: file_modImage_modImage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modImage_modImage_proto_goTypes,
		DependencyIndexes: file_modImage_modImage_proto_depIdxs,
		MessageInfos:      file_modImage_modImage_proto_msgTypes,
	}.Build()
	File_modImage_modImage_proto = out.File
	file_modImage_modImage_proto_rawDesc = nil
	file_modImage_modImage_proto_goTypes = nil
	file_modImage_modImage_proto_depIdxs = nil
}
