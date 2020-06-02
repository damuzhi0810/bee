// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hive.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

type Peers struct {
	Peers []*BzzAddress `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_d635d1ead41ba02c, []int{0}
}
func (m *Peers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(m, src)
}
func (m *Peers) XXX_Size() int {
	return m.Size()
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetPeers() []*BzzAddress {
	if m != nil {
		return m.Peers
	}
	return nil
}

type BzzAddress struct {
	Underlay  []byte `protobuf:"bytes,1,opt,name=Underlay,proto3" json:"Underlay,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Overlay   []byte `protobuf:"bytes,3,opt,name=Overlay,proto3" json:"Overlay,omitempty"`
}

func (m *BzzAddress) Reset()         { *m = BzzAddress{} }
func (m *BzzAddress) String() string { return proto.CompactTextString(m) }
func (*BzzAddress) ProtoMessage()    {}
func (*BzzAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_d635d1ead41ba02c, []int{1}
}
func (m *BzzAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BzzAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BzzAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BzzAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BzzAddress.Merge(m, src)
}
func (m *BzzAddress) XXX_Size() int {
	return m.Size()
}
func (m *BzzAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_BzzAddress.DiscardUnknown(m)
}

var xxx_messageInfo_BzzAddress proto.InternalMessageInfo

func (m *BzzAddress) GetUnderlay() []byte {
	if m != nil {
		return m.Underlay
	}
	return nil
}

func (m *BzzAddress) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *BzzAddress) GetOverlay() []byte {
	if m != nil {
		return m.Overlay
	}
	return nil
}

func init() {
	proto.RegisterType((*Peers)(nil), "hive.Peers")
	proto.RegisterType((*BzzAddress)(nil), "hive.BzzAddress")
}

func init() { proto.RegisterFile("hive.proto", fileDescriptor_d635d1ead41ba02c) }

var fileDescriptor_d635d1ead41ba02c = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc8, 0x2c, 0x4b,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xf4, 0xb9, 0x58, 0x03, 0x52,
	0x53, 0x8b, 0x8a, 0x85, 0xd4, 0xb8, 0x58, 0x0b, 0x40, 0x0c, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e,
	0x23, 0x01, 0x3d, 0xb0, 0x52, 0xa7, 0xaa, 0x2a, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0x20,
	0x88, 0xb4, 0x52, 0x02, 0x17, 0x17, 0x42, 0x50, 0x48, 0x8a, 0x8b, 0x23, 0x34, 0x2f, 0x25, 0xb5,
	0x28, 0x27, 0xb1, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xce, 0x17, 0x92, 0xe1, 0xe2,
	0x0c, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x95, 0x60, 0x02, 0x4b, 0x22, 0x04, 0x84,
	0x24, 0xb8, 0xd8, 0xfd, 0xcb, 0x20, 0x1a, 0x99, 0xc1, 0x72, 0x30, 0xae, 0x93, 0xcc, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x5d, 0x6f,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x11, 0xc0, 0xe9, 0xcb, 0x00, 0x00, 0x00,
}

func (m *Peers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Peers) > 0 {
		for iNdEx := len(m.Peers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Peers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BzzAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BzzAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BzzAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Overlay) > 0 {
		i -= len(m.Overlay)
		copy(dAtA[i:], m.Overlay)
		i = encodeVarintHive(dAtA, i, uint64(len(m.Overlay)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintHive(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Underlay) > 0 {
		i -= len(m.Underlay)
		copy(dAtA[i:], m.Underlay)
		i = encodeVarintHive(dAtA, i, uint64(len(m.Underlay)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHive(dAtA []byte, offset int, v uint64) int {
	offset -= sovHive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovHive(uint64(l))
		}
	}
	return n
}

func (m *BzzAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Underlay)
	if l > 0 {
		n += 1 + l + sovHive(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovHive(uint64(l))
	}
	l = len(m.Overlay)
	if l > 0 {
		n += 1 + l + sovHive(uint64(l))
	}
	return n
}

func sovHive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHive(x uint64) (n int) {
	return sovHive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Peers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHive
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
			return fmt.Errorf("proto: Peers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &BzzAddress{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BzzAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHive
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
			return fmt.Errorf("proto: BzzAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BzzAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Underlay", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Underlay = append(m.Underlay[:0], dAtA[iNdEx:postIndex]...)
			if m.Underlay == nil {
				m.Underlay = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overlay", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Overlay = append(m.Overlay[:0], dAtA[iNdEx:postIndex]...)
			if m.Overlay == nil {
				m.Overlay = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHive
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
					return 0, ErrIntOverflowHive
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHive
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
				return 0, ErrInvalidLengthHive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHive = fmt.Errorf("proto: unexpected end of group")
)
