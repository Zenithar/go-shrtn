// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shrtn/link/v1/link.proto

package linkv1

import (
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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
	// Hash contains the url hash used to resolve the url from public side.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash"`
	// Url is the explicit url stored matching the hash.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url"`
	// Description is the link description.
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
	return m.Unmarshal(b)
}

func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Link.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}

func (m *Link) XXX_Size() int {
	return m.Size()
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
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xce, 0x28, 0x2a,
	0xc9, 0xd3, 0xcf, 0xc9, 0xcc, 0xcb, 0xd6, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x60, 0x19, 0x3d, 0xb0, 0x48, 0x99, 0xa1, 0x94, 0x63, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x6a, 0x5e, 0x59, 0x7e, 0x65, 0x41, 0x51, 0x7e, 0x45,
	0xa5, 0x3e, 0x58, 0x6d, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62,
	0x49, 0xaa, 0x3e, 0x06, 0x03, 0x62, 0xa2, 0x94, 0x2e, 0x92, 0x11, 0xe9, 0xf9, 0xe9, 0xf9, 0x10,
	0xcd, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x94, 0x2b, 0xd5, 0x70, 0xb1, 0xf8,
	0x64, 0xe6, 0x65, 0x0b, 0xc9, 0x70, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x3a, 0x71, 0xbc, 0xba, 0x27, 0x0f, 0xe6, 0x07, 0x81, 0x49, 0x21, 0x49, 0x2e, 0xe6, 0xd2,
	0xa2, 0x1c, 0x09, 0x26, 0xb0, 0x24, 0xfb, 0xab, 0x7b, 0xf2, 0x20, 0x6e, 0x10, 0x88, 0x10, 0xb2,
	0xe6, 0xe2, 0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x06,
	0x2b, 0x91, 0x7c, 0x75, 0x4f, 0x5e, 0x14, 0x49, 0x58, 0x27, 0x3f, 0x37, 0xb3, 0x24, 0x35, 0xb7,
	0xa0, 0xa4, 0x32, 0x08, 0x59, 0xb5, 0x53, 0xf2, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0xc8, 0x25, 0x95, 0x5f, 0x94, 0xae, 0x57, 0x95, 0x9a, 0x97, 0x59, 0x92,
	0x91, 0x58, 0xa4, 0x87, 0x12, 0x32, 0x4e, 0x9c, 0x20, 0x57, 0x06, 0x80, 0x9c, 0x1c, 0xc0, 0x18,
	0xc5, 0x06, 0x12, 0x2d, 0x33, 0x5c, 0xc4, 0xc4, 0x1c, 0xec, 0x13, 0xb1, 0x8a, 0x89, 0x37, 0x18,
	0xac, 0x14, 0xa4, 0x42, 0x2f, 0xcc, 0xf0, 0x14, 0x94, 0x1f, 0x03, 0xe2, 0xc7, 0x84, 0x19, 0x26,
	0xb1, 0x81, 0x7d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0xda, 0x2b, 0x42, 0x86, 0x01,
	0x00, 0x00,
}

func (m *Link) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Link) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLink(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLink(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLink(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLink(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func (m *Link) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovLink(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLink(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func sozLink(x uint64) (n int) {
	return sovLink(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Link) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLink
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Link: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Link: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLink
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLink
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLink
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLink(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLink
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLink
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipLink(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLink
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLink
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLink
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLink
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLink
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLink
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLink(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLink
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLink = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLink   = fmt.Errorf("proto: integer overflow")
)
