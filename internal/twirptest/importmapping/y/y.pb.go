// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: y/y.proto

package y

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

type MsgY struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgY) Reset() {
	*x = MsgY{}
	if protoimpl.UnsafeEnabled {
		mi := &file_y_y_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgY) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgY) ProtoMessage() {}

func (x *MsgY) ProtoReflect() protoreflect.Message {
	mi := &file_y_y_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgY.ProtoReflect.Descriptor instead.
func (*MsgY) Descriptor() ([]byte, []int) {
	return file_y_y_proto_rawDescGZIP(), []int{0}
}

var File_y_y_proto protoreflect.FileDescriptor

var file_y_y_proto_rawDesc = []byte{
	0x0a, 0x09, 0x79, 0x2f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x74, 0x77, 0x69,
	0x72, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x77, 0x69, 0x72,
	0x70, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x79, 0x22, 0x06, 0x0a, 0x04, 0x4d, 0x73, 0x67, 0x59, 0x42, 0x1f, 0x5a,
	0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6c,
	0x6c, 0x2f, 0x62, 0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x2f, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_y_y_proto_rawDescOnce sync.Once
	file_y_y_proto_rawDescData = file_y_y_proto_rawDesc
)

func file_y_y_proto_rawDescGZIP() []byte {
	file_y_y_proto_rawDescOnce.Do(func() {
		file_y_y_proto_rawDescData = protoimpl.X.CompressGZIP(file_y_y_proto_rawDescData)
	})
	return file_y_y_proto_rawDescData
}

var file_y_y_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_y_y_proto_goTypes = []interface{}{
	(*MsgY)(nil), // 0: twirp.internal.twirptest.importmapping.y.MsgY
}
var file_y_y_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_y_y_proto_init() }
func file_y_y_proto_init() {
	if File_y_y_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_y_y_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgY); i {
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
			RawDescriptor: file_y_y_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_y_y_proto_goTypes,
		DependencyIndexes: file_y_y_proto_depIdxs,
		MessageInfos:      file_y_y_proto_msgTypes,
	}.Build()
	File_y_y_proto = out.File
	file_y_y_proto_rawDesc = nil
	file_y_y_proto_goTypes = nil
	file_y_y_proto_depIdxs = nil
}
