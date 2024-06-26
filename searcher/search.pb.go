// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: search.proto

package searcher

import (
	ent "gar/ent"
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

type DocId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DocId) Reset() {
	*x = DocId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocId) ProtoMessage() {}

func (x *DocId) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocId.ProtoReflect.Descriptor instead.
func (*DocId) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *DocId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AffectedCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *AffectedCount) Reset() {
	*x = AffectedCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedCount) ProtoMessage() {}

func (x *AffectedCount) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedCount.ProtoReflect.Descriptor instead.
func (*AffectedCount) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *AffectedCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   *ent.TermQuery `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	OnFlag  uint64         `protobuf:"varint,2,opt,name=OnFlag,proto3" json:"OnFlag,omitempty"`
	OffFlag uint64         `protobuf:"varint,3,opt,name=OffFlag,proto3" json:"OffFlag,omitempty"`
	OrFlags []uint64       `protobuf:"varint,4,rep,packed,name=OrFlags,proto3" json:"OrFlags,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
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
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchRequest) GetQuery() *ent.TermQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SearchRequest) GetOnFlag() uint64 {
	if x != nil {
		return x.OnFlag
	}
	return 0
}

func (x *SearchRequest) GetOffFlag() uint64 {
	if x != nil {
		return x.OffFlag
	}
	return 0
}

func (x *SearchRequest) GetOrFlags() []uint64 {
	if x != nil {
		return x.OrFlags
	}
	return nil
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*ent.Document `protobuf:"bytes,1,rep,name=Documents,proto3" json:"Documents,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchResult) GetDocuments() []*ent.Document {
	if x != nil {
		return x.Documents
	}
	return nil
}

type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x0d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x72,
	0x6d, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a,
	0x05, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x81, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4f, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x4f, 0x66, 0x66, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x4f, 0x66, 0x66, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x4f, 0x72, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x22, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2b, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x0e,
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xe7,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x44, 0x6f, 0x63, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38,
	0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_search_proto_goTypes = []interface{}{
	(*DocId)(nil),         // 0: searcher.DocId
	(*AffectedCount)(nil), // 1: searcher.AffectedCount
	(*SearchRequest)(nil), // 2: searcher.SearchRequest
	(*SearchResult)(nil),  // 3: searcher.SearchResult
	(*CountRequest)(nil),  // 4: searcher.CountRequest
	(*ent.TermQuery)(nil), // 5: ent.TermQuery
	(*ent.Document)(nil),  // 6: ent.Document
}
var file_search_proto_depIdxs = []int32{
	5, // 0: searcher.SearchRequest.Query:type_name -> ent.TermQuery
	6, // 1: searcher.SearchResult.Documents:type_name -> ent.Document
	2, // 2: searcher.SearchService.Search:input_type -> searcher.SearchRequest
	0, // 3: searcher.SearchService.Delete:input_type -> searcher.DocId
	6, // 4: searcher.SearchService.Add:input_type -> ent.Document
	4, // 5: searcher.SearchService.Count:input_type -> searcher.CountRequest
	3, // 6: searcher.SearchService.Search:output_type -> searcher.SearchResult
	1, // 7: searcher.SearchService.Delete:output_type -> searcher.AffectedCount
	1, // 8: searcher.SearchService.Add:output_type -> searcher.AffectedCount
	1, // 9: searcher.SearchService.Count:output_type -> searcher.AffectedCount
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocId); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedCount); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
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
		file_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRequest); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
