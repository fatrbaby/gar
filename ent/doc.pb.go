// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: doc.proto

package ent

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

type Keyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=Field,proto3" json:"Field,omitempty"`
	Word  string `protobuf:"bytes,2,opt,name=Word,proto3" json:"Word,omitempty"`
}

func (x *Keyword) Reset() {
	*x = Keyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyword) ProtoMessage() {}

func (x *Keyword) ProtoReflect() protoreflect.Message {
	mi := &file_doc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyword.ProtoReflect.Descriptor instead.
func (*Keyword) Descriptor() ([]byte, []int) {
	return file_doc_proto_rawDescGZIP(), []int{0}
}

func (x *Keyword) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Keyword) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uid      uint64     `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Features uint64     `protobuf:"varint,3,opt,name=Features,proto3" json:"Features,omitempty"`
	Keywords []*Keyword `protobuf:"bytes,4,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
	Bytes    []byte     `protobuf:"bytes,5,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_doc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_doc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_doc_proto_rawDescGZIP(), []int{1}
}

func (x *Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Document) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Document) GetFeatures() uint64 {
	if x != nil {
		return x.Features
	}
	return 0
}

func (x *Document) GetKeywords() []*Keyword {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Document) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_doc_proto protoreflect.FileDescriptor

var file_doc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x6e, 0x74,
	0x22, 0x33, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x12, 0x28, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_doc_proto_rawDescOnce sync.Once
	file_doc_proto_rawDescData = file_doc_proto_rawDesc
)

func file_doc_proto_rawDescGZIP() []byte {
	file_doc_proto_rawDescOnce.Do(func() {
		file_doc_proto_rawDescData = protoimpl.X.CompressGZIP(file_doc_proto_rawDescData)
	})
	return file_doc_proto_rawDescData
}

var file_doc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_doc_proto_goTypes = []interface{}{
	(*Keyword)(nil),  // 0: ent.Keyword
	(*Document)(nil), // 1: ent.Document
}
var file_doc_proto_depIdxs = []int32{
	0, // 0: ent.Document.Keywords:type_name -> ent.Keyword
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_doc_proto_init() }
func file_doc_proto_init() {
	if File_doc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_doc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyword); i {
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
		file_doc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
			RawDescriptor: file_doc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_doc_proto_goTypes,
		DependencyIndexes: file_doc_proto_depIdxs,
		MessageInfos:      file_doc_proto_msgTypes,
	}.Build()
	File_doc_proto = out.File
	file_doc_proto_rawDesc = nil
	file_doc_proto_goTypes = nil
	file_doc_proto_depIdxs = nil
}
