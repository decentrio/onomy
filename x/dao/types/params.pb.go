// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: onomyprotocol/dao/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// the period of blocks to withdraw the dao staking reward
	WithdrawRewardPeriod int64 `protobuf:"varint,1,opt,name=withdraw_reward_period,json=withdrawRewardPeriod,proto3" json:"withdraw_reward_period,omitempty" yaml:"withdraw_reward_period"`
	// the rate of total dao's staking coins to keep unstaked
	PoolRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=pool_rate,json=poolRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"pool_rate"`
	// the max rage of total dao's staking coins to be allowed in proposals
	MaxProposalRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=max_proposal_rate,json=maxProposalRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_proposal_rate"`
	// the max validator's commission to be staked by the dao
	MaxValCommission cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=max_val_commission,json=maxValCommission,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_val_commission"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1367893a600edc4a, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetWithdrawRewardPeriod() int64 {
	if m != nil {
		return m.WithdrawRewardPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "onomyprotocol.dao.v1.Params")
}

func init() { proto.RegisterFile("onomyprotocol/dao/v1/params.proto", fileDescriptor_1367893a600edc4a) }

var fileDescriptor_1367893a600edc4a = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x93, 0x56, 0x8a, 0x06, 0x41, 0x0d, 0x45, 0x8a, 0x62, 0xfa, 0xe3, 0xa6, 0x0b, 0xc9,
	0x50, 0xdc, 0xb9, 0x92, 0x2a, 0xae, 0x04, 0x6b, 0x16, 0x0a, 0x6e, 0xc2, 0x6d, 0x32, 0xa4, 0x83,
	0x99, 0xde, 0x61, 0x66, 0x4c, 0x93, 0xb7, 0xf0, 0xad, 0xec, 0xb2, 0x4b, 0x71, 0x51, 0xa4, 0x7d,
	0x03, 0x9f, 0x40, 0x32, 0x56, 0xa1, 0xe0, 0xa2, 0xbb, 0xcb, 0xb9, 0xe7, 0xfb, 0x36, 0xc7, 0x69,
	0xe3, 0x18, 0x79, 0x21, 0x24, 0x6a, 0x8c, 0x30, 0x25, 0x31, 0x20, 0xc9, 0x7a, 0x44, 0x80, 0x04,
	0xae, 0x7c, 0x13, 0xbb, 0xf5, 0xb5, 0x8a, 0x1f, 0x03, 0xfa, 0x59, 0xef, 0xa8, 0x9e, 0x60, 0x82,
	0x26, 0x24, 0xe5, 0xf5, 0xd3, 0xed, 0xbc, 0x55, 0x9c, 0xda, 0xc0, 0xc0, 0xee, 0xa3, 0x73, 0x38,
	0x61, 0x7a, 0x14, 0x4b, 0x98, 0x84, 0x92, 0x4e, 0x40, 0xc6, 0xa1, 0xa0, 0x92, 0x61, 0xdc, 0xb0,
	0x5b, 0x76, 0xb7, 0xda, 0x6f, 0x7f, 0xcd, 0x9b, 0x27, 0x05, 0xf0, 0xf4, 0xa2, 0xf3, 0x7f, 0xaf,
	0x13, 0xd4, 0x7f, 0x1f, 0x81, 0xc9, 0x07, 0x26, 0x76, 0x2f, 0x9d, 0x1d, 0x81, 0x98, 0x86, 0x12,
	0x34, 0x6d, 0x54, 0x5a, 0x76, 0x77, 0xb7, 0x7f, 0x3a, 0x9d, 0x37, 0xad, 0x8f, 0x79, 0xf3, 0x38,
	0x42, 0xc5, 0x51, 0xa9, 0xf8, 0xd9, 0x67, 0x48, 0x38, 0xe8, 0x91, 0x7f, 0x4b, 0x13, 0x88, 0x8a,
	0x6b, 0x1a, 0x05, 0xdb, 0x25, 0x15, 0x80, 0xa6, 0xee, 0x9d, 0x73, 0xc0, 0x21, 0x0f, 0x85, 0x44,
	0x81, 0x0a, 0x56, 0xa6, 0xea, 0xe6, 0xa6, 0x3d, 0x0e, 0xf9, 0x60, 0x05, 0x1b, 0xe1, 0xbd, 0xe3,
	0x96, 0xc2, 0x0c, 0xd2, 0x30, 0x42, 0xce, 0x99, 0x52, 0x0c, 0xc7, 0x8d, 0xad, 0xcd, 0x8d, 0xfb,
	0x1c, 0xf2, 0x07, 0x48, 0xaf, 0xfe, 0xe0, 0xfe, 0xcd, 0x74, 0xe1, 0xd9, 0xb3, 0x85, 0x67, 0x7f,
	0x2e, 0x3c, 0xfb, 0x75, 0xe9, 0x59, 0xb3, 0xa5, 0x67, 0xbd, 0x2f, 0x3d, 0xeb, 0xe9, 0x2c, 0x61,
	0x7a, 0xf4, 0x32, 0xf4, 0x23, 0xe4, 0x64, 0x7d, 0x3d, 0x49, 0x15, 0x95, 0x19, 0x25, 0xb9, 0xd9,
	0x51, 0x17, 0x82, 0xaa, 0x61, 0xcd, 0x7c, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa3,
	0x5d, 0x65, 0xe9, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxValCommission.Size()
		i -= size
		if _, err := m.MaxValCommission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MaxProposalRate.Size()
		i -= size
		if _, err := m.MaxProposalRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.PoolRate.Size()
		i -= size
		if _, err := m.PoolRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.WithdrawRewardPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawRewardPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.WithdrawRewardPeriod != 0 {
		n += 1 + sovParams(uint64(m.WithdrawRewardPeriod))
	}
	l = m.PoolRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxProposalRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxValCommission.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawRewardPeriod", wireType)
			}
			m.WithdrawRewardPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawRewardPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProposalRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxProposalRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValCommission", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxValCommission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
