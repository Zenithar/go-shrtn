// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shrtn/link/v1/link.proto

package linkv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Link is a shortened URL.
type Link struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d89dfa300aa123, []int{0}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}

func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}

func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}

func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}

func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Link)(nil), "shrtn.link.v1.Link")
}

func init() { proto.RegisterFile("shrtn/link/v1/link.proto", fileDescriptor_44d89dfa300aa123) }

var fileDescriptor_44d89dfa300aa123 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xce, 0x28, 0x2a,
	0xc9, 0xd3, 0xcf, 0xc9, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x60, 0x19, 0x3d, 0xb0, 0x48, 0x99, 0xa1, 0x92, 0x1f, 0x17, 0x8b, 0x4f, 0x66,
	0x5e, 0xb6, 0x90, 0x10, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x5c, 0x5a, 0x94, 0x23, 0xc1, 0x04, 0x16, 0x02, 0x31, 0x85,
	0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98,
	0xc1, 0x32, 0xc8, 0x42, 0x4e, 0xe1, 0x5c, 0x52, 0xf9, 0x45, 0xe9, 0x7a, 0x55, 0xa9, 0x79, 0x99,
	0x25, 0x19, 0x89, 0x45, 0x7a, 0x28, 0xb6, 0x39, 0x71, 0x82, 0xec, 0x0a, 0x00, 0xb9, 0x23, 0x80,
	0x31, 0x8a, 0x0d, 0x24, 0x5a, 0x66, 0xb8, 0x88, 0x89, 0x39, 0xd8, 0x27, 0x62, 0x15, 0x13, 0x6f,
	0x30, 0x58, 0x29, 0x48, 0x85, 0x5e, 0x98, 0xe1, 0x29, 0x28, 0x3f, 0x06, 0xc4, 0x8f, 0x09, 0x33,
	0x4c, 0x62, 0x03, 0x3b, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x7e, 0x71, 0x5e, 0xda,
	0x00, 0x00, 0x00,
}
