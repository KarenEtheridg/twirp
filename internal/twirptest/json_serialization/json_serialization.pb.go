// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: json_serialization.proto

package json_serialization

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

type Msg_FooBar int32

const (
	Msg_FOO Msg_FooBar = 0
	Msg_BAR Msg_FooBar = 1
)

// Enum value maps for Msg_FooBar.
var (
	Msg_FooBar_name = map[int32]string{
		0: "FOO",
		1: "BAR",
	}
	Msg_FooBar_value = map[string]int32{
		"FOO": 0,
		"BAR": 1,
	}
)

func (x Msg_FooBar) Enum() *Msg_FooBar {
	p := new(Msg_FooBar)
	*p = x
	return p
}

func (x Msg_FooBar) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Msg_FooBar) Descriptor() protoreflect.EnumDescriptor {
	return file_json_serialization_proto_enumTypes[0].Descriptor()
}

func (Msg_FooBar) Type() protoreflect.EnumType {
	return &file_json_serialization_proto_enumTypes[0]
}

func (x Msg_FooBar) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Msg_FooBar.Descriptor instead.
func (Msg_FooBar) EnumDescriptor() ([]byte, []int) {
	return file_json_serialization_proto_rawDescGZIP(), []int{0, 0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string     `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber int32      `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Hell       float64    `protobuf:"fixed64,3,opt,name=hell,proto3" json:"hell,omitempty"`
	Foobar     Msg_FooBar `protobuf:"varint,4,opt,name=foobar,proto3,enum=Msg_FooBar" json:"foobar,omitempty"`
	Snippets   []string   `protobuf:"bytes,5,rep,name=snippets,proto3" json:"snippets,omitempty"`
	AllEmpty   bool       `protobuf:"varint,6,opt,name=all_empty,json=allEmpty,proto3" json:"all_empty,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_serialization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_json_serialization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_json_serialization_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Msg) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Msg) GetHell() float64 {
	if x != nil {
		return x.Hell
	}
	return 0
}

func (x *Msg) GetFoobar() Msg_FooBar {
	if x != nil {
		return x.Foobar
	}
	return Msg_FOO
}

func (x *Msg) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

func (x *Msg) GetAllEmpty() bool {
	if x != nil {
		return x.AllEmpty
	}
	return false
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_json_serialization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_json_serialization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_json_serialization_proto_rawDescGZIP(), []int{1}
}

var File_json_serialization_proto protoreflect.FileDescriptor

var file_json_serialization_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x23, 0x0a,
	0x06, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x4d, 0x73, 0x67, 0x2e, 0x46, 0x6f, 0x6f, 0x42, 0x61, 0x72, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x62,
	0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x0a, 0x06, 0x46,
	0x6f, 0x6f, 0x42, 0x61, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f, 0x4f, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x01, 0x22, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x2d, 0x0a, 0x11, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x53,
	0x4f, 0x4e, 0x12, 0x04, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x04, 0x2e, 0x4d, 0x73, 0x67, 0x22, 0x00,
	0x42, 0x15, 0x5a, 0x13, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_json_serialization_proto_rawDescOnce sync.Once
	file_json_serialization_proto_rawDescData = file_json_serialization_proto_rawDesc
)

func file_json_serialization_proto_rawDescGZIP() []byte {
	file_json_serialization_proto_rawDescOnce.Do(func() {
		file_json_serialization_proto_rawDescData = protoimpl.X.CompressGZIP(file_json_serialization_proto_rawDescData)
	})
	return file_json_serialization_proto_rawDescData
}

var file_json_serialization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_json_serialization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_json_serialization_proto_goTypes = []interface{}{
	(Msg_FooBar)(0), // 0: Msg.FooBar
	(*Msg)(nil),     // 1: Msg
	(*Result)(nil),  // 2: Result
}
var file_json_serialization_proto_depIdxs = []int32{
	0, // 0: Msg.foobar:type_name -> Msg.FooBar
	1, // 1: JSONSerialization.EchoJSON:input_type -> Msg
	1, // 2: JSONSerialization.EchoJSON:output_type -> Msg
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_json_serialization_proto_init() }
func file_json_serialization_proto_init() {
	if File_json_serialization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_json_serialization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_json_serialization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_json_serialization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_json_serialization_proto_goTypes,
		DependencyIndexes: file_json_serialization_proto_depIdxs,
		EnumInfos:         file_json_serialization_proto_enumTypes,
		MessageInfos:      file_json_serialization_proto_msgTypes,
	}.Build()
	File_json_serialization_proto = out.File
	file_json_serialization_proto_rawDesc = nil
	file_json_serialization_proto_goTypes = nil
	file_json_serialization_proto_depIdxs = nil
}
