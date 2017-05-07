// Code generated by protoc-gen-go.
// source: pb/basic.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/basic.proto

It has these top-level messages:
	Struct
	String
	Bytes
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Control
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Struct struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Struct) Reset()                    { *m = Struct{} }
func (m *Struct) String() string            { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()               {}
func (*Struct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Struct) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type String struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Bytes struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Bytes) Reset()                    { *m = Bytes{} }
func (m *Bytes) String() string            { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()               {}
func (*Bytes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Bytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Bool struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Int struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int) Reset()                    { *m = Int{} }
func (m *Int) String() string            { return proto.CompactTextString(m) }
func (*Int) ProtoMessage()               {}
func (*Int) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Int) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int8 struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int8) Reset()                    { *m = Int8{} }
func (m *Int8) String() string            { return proto.CompactTextString(m) }
func (*Int8) ProtoMessage()               {}
func (*Int8) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Int8) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int16 struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int16) Reset()                    { *m = Int16{} }
func (m *Int16) String() string            { return proto.CompactTextString(m) }
func (*Int16) ProtoMessage()               {}
func (*Int16) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Int16) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32 struct {
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int32) Reset()                    { *m = Int32{} }
func (m *Int32) String() string            { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()               {}
func (*Int32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64 struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int64) Reset()                    { *m = Int64{} }
func (m *Int64) String() string            { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()               {}
func (*Int64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Uint) Reset()                    { *m = Uint{} }
func (m *Uint) String() string            { return proto.CompactTextString(m) }
func (*Uint) ProtoMessage()               {}
func (*Uint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Uint) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint8 struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Uint8) Reset()                    { *m = Uint8{} }
func (m *Uint8) String() string            { return proto.CompactTextString(m) }
func (*Uint8) ProtoMessage()               {}
func (*Uint8) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Uint8) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint16 struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Uint16) Reset()                    { *m = Uint16{} }
func (m *Uint16) String() string            { return proto.CompactTextString(m) }
func (*Uint16) ProtoMessage()               {}
func (*Uint16) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Uint16) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint32 struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Uint32) Reset()                    { *m = Uint32{} }
func (m *Uint32) String() string            { return proto.CompactTextString(m) }
func (*Uint32) ProtoMessage()               {}
func (*Uint32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Uint32) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Uint64 struct {
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Uint64) Reset()                    { *m = Uint64{} }
func (m *Uint64) String() string            { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()               {}
func (*Uint64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Uint64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Float32 struct {
	Value float32 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
}

func (m *Float32) Reset()                    { *m = Float32{} }
func (m *Float32) String() string            { return proto.CompactTextString(m) }
func (*Float32) ProtoMessage()               {}
func (*Float32) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Float32) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Float64 struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *Float64) Reset()                    { *m = Float64{} }
func (m *Float64) String() string            { return proto.CompactTextString(m) }
func (*Float64) ProtoMessage()               {}
func (*Float64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Float64) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Control struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Control) Reset()                    { *m = Control{} }
func (m *Control) String() string            { return proto.CompactTextString(m) }
func (*Control) ProtoMessage()               {}
func (*Control) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Control) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Struct)(nil), "pb.Struct")
	proto.RegisterType((*String)(nil), "pb.String")
	proto.RegisterType((*Bytes)(nil), "pb.Bytes")
	proto.RegisterType((*Bool)(nil), "pb.Bool")
	proto.RegisterType((*Int)(nil), "pb.Int")
	proto.RegisterType((*Int8)(nil), "pb.Int8")
	proto.RegisterType((*Int16)(nil), "pb.Int16")
	proto.RegisterType((*Int32)(nil), "pb.Int32")
	proto.RegisterType((*Int64)(nil), "pb.Int64")
	proto.RegisterType((*Uint)(nil), "pb.Uint")
	proto.RegisterType((*Uint8)(nil), "pb.Uint8")
	proto.RegisterType((*Uint16)(nil), "pb.Uint16")
	proto.RegisterType((*Uint32)(nil), "pb.Uint32")
	proto.RegisterType((*Uint64)(nil), "pb.Uint64")
	proto.RegisterType((*Float32)(nil), "pb.Float32")
	proto.RegisterType((*Float64)(nil), "pb.Float64")
	proto.RegisterType((*Control)(nil), "pb.Control")
}

func init() { proto.RegisterFile("pb/basic.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x4f,
	0x4a, 0x2c, 0xce, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92,
	0xe3, 0x62, 0x0b, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x70, 0xa0, 0xf2, 0x99, 0x79, 0xe9, 0xa8,
	0xf2, 0x9c, 0x30, 0x79, 0x59, 0x2e, 0x56, 0xa7, 0xca, 0x92, 0xd4, 0x62, 0x54, 0x69, 0x1e, 0x98,
	0xb4, 0x0c, 0x17, 0x8b, 0x53, 0x7e, 0x7e, 0x0e, 0x0e, 0xc3, 0xa5, 0xb9, 0x98, 0x3d, 0xf3, 0xd0,
	0x6c, 0x66, 0x45, 0xd2, 0xea, 0x99, 0x57, 0x62, 0x81, 0x43, 0x56, 0x96, 0x8b, 0xd5, 0x33, 0xaf,
	0xc4, 0xd0, 0x0c, 0xaf, 0xb4, 0xb1, 0x11, 0x5e, 0x69, 0x33, 0x13, 0x54, 0x69, 0x66, 0x24, 0xab,
	0x43, 0x33, 0xd1, 0x1d, 0xc6, 0x8b, 0xa4, 0x19, 0x24, 0x6b, 0x81, 0x43, 0x5a, 0x8e, 0x8b, 0x0d,
	0x24, 0x8d, 0xee, 0x34, 0x74, 0x79, 0x74, 0xb7, 0xa1, 0xcb, 0xa3, 0x3b, 0x8e, 0x05, 0x26, 0x2f,
	0xcf, 0xc5, 0xee, 0x96, 0x93, 0x9f, 0x88, 0x61, 0x00, 0x13, 0xba, 0x02, 0x74, 0x13, 0x18, 0x91,
	0x14, 0x38, 0xe7, 0xe7, 0x95, 0x14, 0xa1, 0xc7, 0x0b, 0xcc, 0x09, 0x49, 0x6c, 0xe0, 0xf4, 0x61,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x9e, 0x8d, 0x30, 0x31, 0x02, 0x00, 0x00,
}