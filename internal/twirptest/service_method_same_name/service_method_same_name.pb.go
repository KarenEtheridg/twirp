// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_method_same_name.proto

package service_method_same_name

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a03bb974d2dd0e1, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg)(nil), "Msg")
}

func init() { proto.RegisterFile("service_method_same_name.proto", fileDescriptor_6a03bb974d2dd0e1) }

var fileDescriptor_6a03bb974d2dd0e1 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x89, 0x2f, 0x4e, 0xcc, 0x4d, 0x8d,
	0xcf, 0x4b, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe5, 0x62, 0xf6, 0x2d,
	0x4e, 0x37, 0x92, 0xe1, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17, 0x12, 0x81, 0xd2, 0x2c, 0x7a, 0xbe,
	0xc5, 0xe9, 0x52, 0x60, 0x52, 0x89, 0xc1, 0x49, 0x2a, 0x4a, 0x02, 0x97, 0x31, 0x49, 0x6c, 0x60,
	0x73, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xcf, 0xa1, 0x77, 0x69, 0x00, 0x00, 0x00,
}
