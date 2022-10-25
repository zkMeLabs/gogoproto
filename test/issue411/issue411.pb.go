// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue411.proto

package issue411

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Span struct {
	TraceID              TraceID  `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3,customtype=TraceID" json:"trace_id"`
	SpanID               SpanID   `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3,customtype=SpanID" json:"span_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1ed5cde895f96f, []int{0}
}
func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Span)(nil), "issue411.Span")
}

func init() { proto.RegisterFile("issue411.proto", fileDescriptor_7e1ed5cde895f96f) }

var fileDescriptor_7e1ed5cde895f96f = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x34, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x82, 0xfa, 0x20, 0x16, 0x44, 0x5e, 0xa9, 0x80, 0x8b, 0x25, 0xb8,
	0x20, 0x31, 0x4f, 0xc8, 0x94, 0x8b, 0xa3, 0xa4, 0x28, 0x31, 0x39, 0x35, 0x3e, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xc7, 0x49, 0xea, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd9, 0x43,
	0x40, 0xe2, 0x9e, 0x2e, 0x8f, 0x10, 0xcc, 0x20, 0x76, 0xb0, 0x5a, 0xcf, 0x14, 0x21, 0x43, 0x2e,
	0xf6, 0xe2, 0x82, 0xc4, 0x3c, 0x90, 0x2e, 0x26, 0xb0, 0x2e, 0x09, 0xa8, 0x2e, 0x36, 0x90, 0xa9,
	0x60, 0x4d, 0x50, 0x56, 0x10, 0x1b, 0x48, 0xa1, 0x67, 0x4a, 0x12, 0x1b, 0xd8, 0x62, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd3, 0x65, 0xb1, 0xaa, 0x00, 0x00, 0x00,
}
