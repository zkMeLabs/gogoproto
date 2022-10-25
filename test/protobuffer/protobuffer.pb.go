// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuffer.proto

package protobuffer

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
	math_bits "math/bits"
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

type PBuffMarshal struct {
	Field1               []byte   `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2               *int32   `protobuf:"varint,2,opt,name=Field2" json:"Field2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBuffMarshal) Reset()         { *m = PBuffMarshal{} }
func (m *PBuffMarshal) String() string { return proto.CompactTextString(m) }
func (*PBuffMarshal) ProtoMessage()    {}
func (*PBuffMarshal) Descriptor() ([]byte, []int) {
	return fileDescriptor_12798c2215df1728, []int{0}
}
func (m *PBuffMarshal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBuffMarshal.Unmarshal(m, b)
}
func (m *PBuffMarshal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBuffMarshal.Marshal(b, m, deterministic)
}
func (m *PBuffMarshal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBuffMarshal.Merge(m, src)
}
func (m *PBuffMarshal) XXX_Size() int {
	return xxx_messageInfo_PBuffMarshal.Size(m)
}
func (m *PBuffMarshal) XXX_DiscardUnknown() {
	xxx_messageInfo_PBuffMarshal.DiscardUnknown(m)
}

var xxx_messageInfo_PBuffMarshal proto.InternalMessageInfo

func (m *PBuffMarshal) GetField1() []byte {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *PBuffMarshal) GetField2() int32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

type PBuffMarshaler struct {
	Field1               []byte   `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	Field2               *int32   `protobuf:"varint,2,opt,name=Field2" json:"Field2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBuffMarshaler) Reset()         { *m = PBuffMarshaler{} }
func (m *PBuffMarshaler) String() string { return proto.CompactTextString(m) }
func (*PBuffMarshaler) ProtoMessage()    {}
func (*PBuffMarshaler) Descriptor() ([]byte, []int) {
	return fileDescriptor_12798c2215df1728, []int{1}
}
func (m *PBuffMarshaler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBuffMarshaler.Unmarshal(m, b)
}
func (m *PBuffMarshaler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PBuffMarshaler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PBuffMarshaler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBuffMarshaler.Merge(m, src)
}
func (m *PBuffMarshaler) XXX_Size() int {
	return m.Size()
}
func (m *PBuffMarshaler) XXX_DiscardUnknown() {
	xxx_messageInfo_PBuffMarshaler.DiscardUnknown(m)
}

var xxx_messageInfo_PBuffMarshaler proto.InternalMessageInfo

func (m *PBuffMarshaler) GetField1() []byte {
	if m != nil {
		return m.Field1
	}
	return nil
}

func (m *PBuffMarshaler) GetField2() int32 {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return 0
}

func init() {
	proto.RegisterType((*PBuffMarshal)(nil), "protobuffer.PBuffMarshal")
	proto.RegisterType((*PBuffMarshaler)(nil), "protobuffer.PBuffMarshaler")
}

func init() { proto.RegisterFile("protobuffer.proto", fileDescriptor_12798c2215df1728) }

var fileDescriptor_12798c2215df1728 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x4b, 0x4b, 0x2d, 0xd2, 0x03, 0xb3, 0x85, 0xb8, 0x91, 0x84, 0xa4, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x02, 0xfa, 0x20, 0x16, 0x44, 0x89, 0x92, 0x07, 0x17, 0x4f, 0x80,
	0x53, 0x69, 0x5a, 0x9a, 0x6f, 0x62, 0x51, 0x71, 0x46, 0x62, 0x8e, 0x90, 0x18, 0x17, 0x9b, 0x5b,
	0x66, 0x6a, 0x4e, 0x8a, 0xa1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x94, 0x07, 0x17, 0x37,
	0x92, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x85, 0x8a, 0x1b, 0x59, 0x71, 0x74, 0x2c, 0x94, 0x67, 0x58,
	0xb0, 0x50, 0x9e, 0x41, 0xc9, 0x8b, 0x8b, 0x0f, 0xd9, 0xa4, 0xd4, 0x22, 0xb2, 0xcc, 0x62, 0x5c,
	0xb0, 0x50, 0x9e, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xeb, 0x86, 0x4b, 0xcc, 0x00, 0x00,
	0x00,
}

func (m *PBuffMarshaler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PBuffMarshaler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PBuffMarshaler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Field2 != nil {
		i = encodeVarintProtobuffer(dAtA, i, uint64(*m.Field2))
		i--
		dAtA[i] = 0x10
	}
	if m.Field1 != nil {
		i -= len(m.Field1)
		copy(dAtA[i:], m.Field1)
		i = encodeVarintProtobuffer(dAtA, i, uint64(len(m.Field1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtobuffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtobuffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PBuffMarshaler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Field1 != nil {
		l = len(m.Field1)
		n += 1 + l + sovProtobuffer(uint64(l))
	}
	if m.Field2 != nil {
		n += 1 + sovProtobuffer(uint64(*m.Field2))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtobuffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtobuffer(x uint64) (n int) {
	return sovProtobuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
