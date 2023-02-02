// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3extension/proto3ext.proto

package proto3extension

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptor "github.com/cosmos/gogoproto/protoc-gen-gogo/descriptor"
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

var E_Primary = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51234,
	Name:          "proto3extension.primary",
	Tag:           "varint,51234,opt,name=primary",
	Filename:      "proto3extension/proto3ext.proto",
}

var E_Index = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51235,
	Name:          "proto3extension.index",
	Tag:           "varint,51235,opt,name=index",
	Filename:      "proto3extension/proto3ext.proto",
}

func init() {
	proto.RegisterExtension(E_Primary)
	proto.RegisterExtension(E_Index)
}

func init() { proto.RegisterFile("proto3extension/proto3ext.proto", fileDescriptor_b0231eb4ca239c4e) }

var fileDescriptor_b0231eb4ca239c4e = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x4e, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x87, 0xf3, 0xf5, 0xc0,
	0x2c, 0x21, 0x7e, 0x34, 0x05, 0x52, 0x0a, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0x10, 0x85, 0x49,
	0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0x2d, 0x56,
	0x96, 0x5c, 0xec, 0x05, 0x45, 0x99, 0xb9, 0x89, 0x45, 0x95, 0x42, 0xb2, 0x7a, 0x10, 0xd5, 0x7a,
	0x30, 0xd5, 0x7a, 0x6e, 0x99, 0xa9, 0x39, 0x29, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12,
	0x8b, 0x26, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xd4, 0x5b, 0x99, 0x72, 0xb1, 0x66, 0xe6,
	0xa5, 0xa4, 0x56, 0x10, 0xd2, 0xb8, 0x18, 0xaa, 0x11, 0xa2, 0x3a, 0x89, 0x0d, 0xe2, 0x48, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x43, 0xc3, 0x57, 0xce, 0x00, 0x00, 0x00,
}
