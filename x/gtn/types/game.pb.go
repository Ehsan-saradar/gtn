// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gtn/gtn/game.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
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

type Game struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CommitmentHash string `protobuf:"bytes,2,opt,name=commitmentHash,proto3" json:"commitmentHash,omitempty"`
	Duration       uint64 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	EntryFee       uint64 `protobuf:"varint,4,opt,name=entryFee,proto3" json:"entryFee,omitempty"`
	Reward         uint64 `protobuf:"varint,5,opt,name=reward,proto3" json:"reward,omitempty"`
	Creator        string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_8285eddab4fa593f, []int{0}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Game.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return m.Size()
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Game) GetCommitmentHash() string {
	if m != nil {
		return m.CommitmentHash
	}
	return ""
}

func (m *Game) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Game) GetEntryFee() uint64 {
	if m != nil {
		return m.EntryFee
	}
	return 0
}

func (m *Game) GetReward() uint64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *Game) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Game)(nil), "gtn.gtn.Game")
}

func init() { proto.RegisterFile("gtn/gtn/game.proto", fileDescriptor_8285eddab4fa593f) }

var fileDescriptor_8285eddab4fa593f = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2f, 0xc9, 0xd3,
	0x07, 0xe3, 0xc4, 0xdc, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xf6, 0xf4, 0x92, 0x3c,
	0xbd, 0xf4, 0x92, 0x3c, 0xa5, 0x15, 0x8c, 0x5c, 0x2c, 0xee, 0x89, 0xb9, 0xa9, 0x42, 0x7c, 0x5c,
	0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x6a, 0x5c,
	0x7c, 0xc9, 0xf9, 0xb9, 0xb9, 0x99, 0x25, 0xb9, 0xa9, 0x79, 0x25, 0x1e, 0x89, 0xc5, 0x19, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x68, 0xa2, 0x42, 0x52, 0x5c, 0x1c, 0x29, 0xa5, 0x45, 0x89,
	0x25, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x60, 0xdd, 0x70, 0x3e, 0x48, 0x2e, 0x35, 0xaf, 0xa4, 0xa8,
	0xd2, 0x2d, 0x35, 0x55, 0x82, 0x05, 0x22, 0x07, 0xe3, 0x0b, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x96,
	0x27, 0x16, 0xa5, 0x48, 0xb0, 0x82, 0x65, 0xa0, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4,
	0xc4, 0x92, 0xfc, 0x22, 0x09, 0x36, 0xb0, 0x85, 0x30, 0xae, 0x93, 0xe6, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xf1, 0x83, 0x7c, 0x57, 0x01, 0xf6, 0x63, 0x49, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x97, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x13, 0x73,
	0xe3, 0xfb, 0x00, 0x00, 0x00,
}

func (m *Game) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Game) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Game) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Reward != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Reward))
		i--
		dAtA[i] = 0x28
	}
	if m.EntryFee != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.EntryFee))
		i--
		dAtA[i] = 0x20
	}
	if m.Duration != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CommitmentHash) > 0 {
		i -= len(m.CommitmentHash)
		copy(dAtA[i:], m.CommitmentHash)
		i = encodeVarintGame(dAtA, i, uint64(len(m.CommitmentHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Game) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGame(uint64(m.Id))
	}
	l = len(m.CommitmentHash)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovGame(uint64(m.Duration))
	}
	if m.EntryFee != 0 {
		n += 1 + sovGame(uint64(m.EntryFee))
	}
	if m.Reward != 0 {
		n += 1 + sovGame(uint64(m.Reward))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Game) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Game: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Game: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitmentHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommitmentHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryFee", wireType)
			}
			m.EntryFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntryFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			m.Reward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)
