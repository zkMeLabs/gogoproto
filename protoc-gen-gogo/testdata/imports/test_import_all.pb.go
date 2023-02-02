// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_import_all.proto

package imports

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	fmt1 "github.com/cosmos/gogoproto/protoc-gen-gogo/testdata/imports/fmt"
	test_a_1 "github.com/cosmos/gogoproto/protoc-gen-gogo/testdata/imports/test_a_1"
	test_a_2 "github.com/cosmos/gogoproto/protoc-gen-gogo/testdata/imports/test_a_2"
	test_b_1 "github.com/cosmos/gogoproto/protoc-gen-gogo/testdata/imports/test_b_1"
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

type All struct {
	Am1                  *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1,proto3" json:"am1,omitempty"`
	Am2                  *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2,proto3" json:"am2,omitempty"`
	Am3                  *test_a_2.M3 `protobuf:"bytes,3,opt,name=am3,proto3" json:"am3,omitempty"`
	Am4                  *test_a_2.M4 `protobuf:"bytes,4,opt,name=am4,proto3" json:"am4,omitempty"`
	Bm1                  *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1,proto3" json:"bm1,omitempty"`
	Bm2                  *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2,proto3" json:"bm2,omitempty"`
	Fmt                  *fmt1.M      `protobuf:"bytes,7,opt,name=fmt,proto3" json:"fmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_324466f0afc16f77, []int{0}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetAm1() *test_a_1.M1 {
	if m != nil {
		return m.Am1
	}
	return nil
}

func (m *All) GetAm2() *test_a_1.M2 {
	if m != nil {
		return m.Am2
	}
	return nil
}

func (m *All) GetAm3() *test_a_2.M3 {
	if m != nil {
		return m.Am3
	}
	return nil
}

func (m *All) GetAm4() *test_a_2.M4 {
	if m != nil {
		return m.Am4
	}
	return nil
}

func (m *All) GetBm1() *test_b_1.M1 {
	if m != nil {
		return m.Bm1
	}
	return nil
}

func (m *All) GetBm2() *test_b_1.M2 {
	if m != nil {
		return m.Bm2
	}
	return nil
}

func (m *All) GetFmt() *fmt1.M {
	if m != nil {
		return m.Fmt
	}
	return nil
}

func init() {
	proto.RegisterType((*All)(nil), "imports.All")
}

func init() { proto.RegisterFile("imports/test_import_all.proto", fileDescriptor_324466f0afc16f77) }

var fileDescriptor_324466f0afc16f77 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0x71, 0xb6, 0xea, 0x06, 0x15, 0x3c, 0x54, 0x0f, 0x51, 0x10, 0x86, 0x07, 0xf1, 0xb2,
	0x84, 0xbc, 0xf4, 0x28, 0x82, 0xde, 0x7b, 0xd9, 0xd1, 0x4b, 0x49, 0xea, 0x5a, 0x07, 0x7d, 0x66,
	0xac, 0xcf, 0xcf, 0xe2, 0xd7, 0x95, 0x97, 0xb4, 0xa2, 0xd0, 0x5c, 0x0a, 0xcd, 0xff, 0x47, 0xf2,
	0x78, 0xf9, 0xdd, 0x01, 0x8f, 0xfe, 0x44, 0x83, 0xa2, 0xfd, 0x40, 0x75, 0xfc, 0xa9, 0x6d, 0xdf,
	0xcb, 0xe3, 0xc9, 0x93, 0x2f, 0xd6, 0x63, 0xbe, 0xbd, 0xf9, 0xe7, 0x6c, 0xad, 0x15, 0xea, 0x68,
	0xe6, 0x12, 0x24, 0x12, 0x28, 0x34, 0xe9, 0x54, 0xce, 0x26, 0x97, 0x7e, 0xcb, 0xfd, 0x7d, 0xeb,
	0x6a, 0x4a, 0x2d, 0x92, 0xc2, 0x78, 0x78, 0xff, 0xbd, 0xcc, 0xb3, 0x97, 0xbe, 0x2f, 0x1e, 0xf2,
	0xcc, 0xa2, 0x16, 0x8b, 0xcd, 0xe2, 0xf1, 0x02, 0xae, 0xe5, 0x48, 0xe5, 0x34, 0xb1, 0xac, 0xf4,
	0x8e, 0x41, 0x74, 0x20, 0x96, 0x49, 0x07, 0xec, 0x20, 0x3a, 0x23, 0xb2, 0x79, 0x07, 0xb2, 0x32,
	0xec, 0x4c, 0x74, 0xa5, 0x38, 0x4b, 0xba, 0x92, 0x5d, 0xc9, 0xce, 0xa1, 0x16, 0xe7, 0x73, 0xce,
	0x8d, 0xf3, 0xb9, 0x38, 0x9f, 0x43, 0x10, 0xab, 0xa4, 0x03, 0x76, 0x50, 0x6c, 0xf2, 0xac, 0x45,
	0x12, 0xeb, 0xe0, 0x2e, 0x7f, 0x5d, 0x8b, 0x24, 0xab, 0x1d, 0xa7, 0xd7, 0xe7, 0xb7, 0xa7, 0xee,
	0x40, 0x1f, 0x5f, 0x4e, 0x36, 0x1e, 0x55, 0xe3, 0x07, 0xf4, 0x83, 0xea, 0x7c, 0xe7, 0xc3, 0xe2,
	0x54, 0xf8, 0x36, 0xdb, 0x6e, 0xff, 0xb9, 0xe5, 0xc3, 0xb0, 0xef, 0x77, 0x4b, 0x56, 0x8d, 0x57,
	0xb9, 0x55, 0x10, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xac, 0xfe, 0xc9, 0xa7, 0x41, 0x02, 0x00,
	0x00,
}
