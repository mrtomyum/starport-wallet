// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/packet.proto

package types

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

type DexPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*DexPacketData_NoData
	//	*DexPacketData_SellOrderPacket
	Packet isDexPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *DexPacketData) Reset()         { *m = DexPacketData{} }
func (m *DexPacketData) String() string { return proto.CompactTextString(m) }
func (*DexPacketData) ProtoMessage()    {}
func (*DexPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e40d3eecbdb512f, []int{0}
}
func (m *DexPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DexPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DexPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DexPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DexPacketData.Merge(m, src)
}
func (m *DexPacketData) XXX_Size() int {
	return m.Size()
}
func (m *DexPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_DexPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_DexPacketData proto.InternalMessageInfo

type isDexPacketData_Packet interface {
	isDexPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DexPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type DexPacketData_SellOrderPacket struct {
	SellOrderPacket *SellOrderPacketData `protobuf:"bytes,2,opt,name=sellOrderPacket,proto3,oneof" json:"sellOrderPacket,omitempty"`
}

func (*DexPacketData_NoData) isDexPacketData_Packet()          {}
func (*DexPacketData_SellOrderPacket) isDexPacketData_Packet() {}

func (m *DexPacketData) GetPacket() isDexPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *DexPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*DexPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *DexPacketData) GetSellOrderPacket() *SellOrderPacketData {
	if x, ok := m.GetPacket().(*DexPacketData_SellOrderPacket); ok {
		return x.SellOrderPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DexPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DexPacketData_NoData)(nil),
		(*DexPacketData_SellOrderPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e40d3eecbdb512f, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// this line is used by starport scaffolding # ibc/packet/proto/message
// SellOrderPacketData defines a struct for the packet payload
type SellOrderPacketData struct {
	AmountDenom string `protobuf:"bytes,1,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	Amount      int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *SellOrderPacketData) Reset()         { *m = SellOrderPacketData{} }
func (m *SellOrderPacketData) String() string { return proto.CompactTextString(m) }
func (*SellOrderPacketData) ProtoMessage()    {}
func (*SellOrderPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e40d3eecbdb512f, []int{2}
}
func (m *SellOrderPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrderPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrderPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrderPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrderPacketData.Merge(m, src)
}
func (m *SellOrderPacketData) XXX_Size() int {
	return m.Size()
}
func (m *SellOrderPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrderPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrderPacketData proto.InternalMessageInfo

func (m *SellOrderPacketData) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *SellOrderPacketData) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SellOrderPacketData) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

// SellOrderPacketAck defines a struct for the packet acknowledgment
type SellOrderPacketAck struct {
}

func (m *SellOrderPacketAck) Reset()         { *m = SellOrderPacketAck{} }
func (m *SellOrderPacketAck) String() string { return proto.CompactTextString(m) }
func (*SellOrderPacketAck) ProtoMessage()    {}
func (*SellOrderPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e40d3eecbdb512f, []int{3}
}
func (m *SellOrderPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrderPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrderPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrderPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrderPacketAck.Merge(m, src)
}
func (m *SellOrderPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *SellOrderPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrderPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrderPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DexPacketData)(nil), "mrtomyum.wallet.dex.DexPacketData")
	proto.RegisterType((*NoData)(nil), "mrtomyum.wallet.dex.NoData")
	proto.RegisterType((*SellOrderPacketData)(nil), "mrtomyum.wallet.dex.SellOrderPacketData")
	proto.RegisterType((*SellOrderPacketAck)(nil), "mrtomyum.wallet.dex.SellOrderPacketAck")
}

func init() { proto.RegisterFile("dex/packet.proto", fileDescriptor_9e40d3eecbdb512f) }

var fileDescriptor_9e40d3eecbdb512f = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x49, 0xad, 0xd0,
	0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0x2d,
	0x2a, 0xc9, 0xcf, 0xad, 0x2c, 0xcd, 0xd5, 0x2b, 0x4f, 0xcc, 0xc9, 0x49, 0x2d, 0xd1, 0x4b, 0x49,
	0xad, 0x50, 0x5a, 0xc6, 0xc8, 0xc5, 0xeb, 0x92, 0x5a, 0x11, 0x00, 0x56, 0xe8, 0x92, 0x58, 0x92,
	0x28, 0x64, 0xca, 0xc5, 0x96, 0x97, 0x0f, 0x62, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49,
	0xeb, 0x61, 0xd1, 0xa7, 0xe7, 0x07, 0x56, 0xe2, 0xc1, 0x10, 0x04, 0x55, 0x2c, 0x14, 0xc2, 0xc5,
	0x5f, 0x9c, 0x9a, 0x93, 0xe3, 0x5f, 0x94, 0x92, 0x5a, 0x04, 0x31, 0x4d, 0x82, 0x09, 0xac, 0x5f,
	0x03, 0xab, 0xfe, 0x60, 0x54, 0xb5, 0x50, 0xc3, 0xd0, 0x8d, 0x70, 0xe2, 0xe0, 0x62, 0x83, 0xf8,
	0x41, 0x89, 0x83, 0x8b, 0x0d, 0x62, 0xa7, 0x52, 0x2a, 0x97, 0x30, 0x16, 0xdd, 0x42, 0x0a, 0x5c,
	0xdc, 0x89, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x2e, 0xa9, 0x79, 0xf9, 0xb9, 0x60, 0xc7, 0x73, 0x06,
	0x21, 0x0b, 0x09, 0x89, 0x71, 0xb1, 0x41, 0xb8, 0x60, 0x97, 0xb1, 0x06, 0x41, 0x79, 0x42, 0x22,
	0x5c, 0xac, 0x05, 0x45, 0x99, 0xc9, 0xa9, 0x12, 0xcc, 0x60, 0x61, 0x08, 0x47, 0x49, 0x84, 0x4b,
	0x08, 0xcd, 0x1a, 0xc7, 0xe4, 0x6c, 0x27, 0x87, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x49, 0x4c, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xf9,
	0x58, 0x1f, 0xe2, 0x63, 0xfd, 0x0a, 0x7d, 0x50, 0x6c, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1,
	0x81, 0x63, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x98, 0x98, 0xaa, 0xa1, 0x01, 0x00,
	0x00,
}

func (m *DexPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DexPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DexPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *DexPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DexPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *DexPacketData_SellOrderPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DexPacketData_SellOrderPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SellOrderPacket != nil {
		{
			size, err := m.SellOrderPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SellOrderPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrderPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrderPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x18
	}
	if m.Amount != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SellOrderPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrderPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrderPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DexPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *DexPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *DexPacketData_SellOrderPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SellOrderPacket != nil {
		l = m.SellOrderPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SellOrderPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovPacket(uint64(m.Amount))
	}
	if m.Price != 0 {
		n += 1 + sovPacket(uint64(m.Price))
	}
	return n
}

func (m *SellOrderPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DexPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: DexPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DexPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &DexPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SellOrderPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &DexPacketData_SellOrderPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *SellOrderPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: SellOrderPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrderPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *SellOrderPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: SellOrderPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrderPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
