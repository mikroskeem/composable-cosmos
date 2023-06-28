// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: centauri/transfermiddleware/v1beta1/genesis.proto

package types

import (
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

// GenesisState defines the module various parameters when first
// initialized
type GenesisState struct {
	TokenInfos []ParachainIBCTokenInfo `protobuf:"bytes,1,rep,name=token_infos,json=tokenInfos,proto3" json:"token_infos"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bba390f1246595d3, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetTokenInfos() []ParachainIBCTokenInfo {
	if m != nil {
		return m.TokenInfos
	}
	return nil
}

// ParachainIBCTokenInfo represents information about transferable IBC tokens from Parachain.
type ParachainIBCTokenInfo struct {
	// ibc_denom is the denomination of the ibced token transferred from the dotsama chain.
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty" yaml:"ibc_denom"`
	// channel_id is source channel in IBC connection from centauri chain
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	// native denom is new native minted denom in centauri chain.
	NativeDenom string `protobuf:"bytes,3,opt,name=native_denom,json=nativeDenom,proto3" json:"native_denom,omitempty" yaml:"native_denom"`
	// asset id is the id of the asset on Picasso
	AssetId string `protobuf:"bytes,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *ParachainIBCTokenInfo) Reset()         { *m = ParachainIBCTokenInfo{} }
func (m *ParachainIBCTokenInfo) String() string { return proto.CompactTextString(m) }
func (*ParachainIBCTokenInfo) ProtoMessage()    {}
func (*ParachainIBCTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bba390f1246595d3, []int{1}
}
func (m *ParachainIBCTokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParachainIBCTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParachainIBCTokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParachainIBCTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParachainIBCTokenInfo.Merge(m, src)
}
func (m *ParachainIBCTokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *ParachainIBCTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ParachainIBCTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ParachainIBCTokenInfo proto.InternalMessageInfo

func (m *ParachainIBCTokenInfo) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetNativeDenom() string {
	if m != nil {
		return m.NativeDenom
	}
	return ""
}

func (m *ParachainIBCTokenInfo) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "centauri.transfermiddleware.v1beta1.GenesisState")
	proto.RegisterType((*ParachainIBCTokenInfo)(nil), "centauri.transfermiddleware.v1beta1.ParachainIBCTokenInfo")
}

func init() {
	proto.RegisterFile("centauri/transfermiddleware/v1beta1/genesis.proto", fileDescriptor_bba390f1246595d3)
}

var fileDescriptor_bba390f1246595d3 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xea, 0x40,
	0x14, 0x85, 0xdb, 0x07, 0x79, 0x0f, 0x06, 0x92, 0xa7, 0x05, 0x22, 0x61, 0xd1, 0x92, 0x71, 0xc3,
	0xaa, 0x4d, 0x95, 0x15, 0xcb, 0x6a, 0x62, 0xba, 0x33, 0xd5, 0x95, 0x1b, 0x32, 0x6d, 0x2f, 0x30,
	0x11, 0x66, 0xb0, 0x33, 0xa2, 0xfc, 0x0b, 0x7f, 0x16, 0x4b, 0x96, 0xae, 0x1a, 0x03, 0x5b, 0x57,
	0xfd, 0x05, 0xa6, 0x1d, 0x50, 0x13, 0x59, 0xb8, 0x3b, 0xf7, 0x9e, 0xfb, 0xdd, 0xb3, 0x38, 0xc8,
	0x8d, 0x80, 0x49, 0xf2, 0x98, 0x50, 0x47, 0x26, 0x84, 0x89, 0x11, 0x24, 0x33, 0x1a, 0xc7, 0x53,
	0x78, 0x22, 0x09, 0x38, 0x0b, 0x37, 0x04, 0x49, 0x5c, 0x67, 0x0c, 0x0c, 0x04, 0x15, 0xf6, 0x3c,
	0xe1, 0x92, 0x1b, 0xa7, 0x7b, 0xc4, 0xfe, 0x89, 0xd8, 0x3b, 0xa4, 0xd3, 0x1c, 0xf3, 0x31, 0x2f,
	0xee, 0x9d, 0x5c, 0x29, 0x14, 0x3f, 0xa0, 0xfa, 0x95, 0xfa, 0x75, 0x23, 0x89, 0x04, 0x83, 0xa0,
	0x9a, 0xe4, 0xf7, 0xc0, 0x86, 0x94, 0x8d, 0xb8, 0x68, 0xeb, 0xdd, 0x52, 0xaf, 0x76, 0x36, 0xb0,
	0x7f, 0x11, 0x60, 0x5f, 0x93, 0x84, 0x44, 0x13, 0x42, 0x99, 0xef, 0x5d, 0xdc, 0xe6, 0x3f, 0x7c,
	0x36, 0xe2, 0x5e, 0x79, 0x95, 0x5a, 0x5a, 0x80, 0xe4, 0x7e, 0x21, 0xf0, 0xbb, 0x8e, 0x5a, 0x07,
	0x6f, 0x0d, 0x17, 0x55, 0x69, 0x18, 0x0d, 0x63, 0x60, 0x7c, 0xd6, 0xd6, 0xbb, 0x7a, 0xaf, 0xea,
	0x35, 0xb3, 0xd4, 0x3a, 0x5a, 0x92, 0xd9, 0x74, 0x80, 0x3f, 0x2d, 0x1c, 0x54, 0x68, 0x18, 0x5d,
	0xe6, 0xd2, 0xe8, 0x23, 0x14, 0x4d, 0x08, 0x63, 0x30, 0x1d, 0xd2, 0xb8, 0xfd, 0xa7, 0x60, 0x5a,
	0x59, 0x6a, 0x1d, 0x2b, 0xe6, 0xcb, 0xc3, 0x41, 0x75, 0x37, 0xf8, 0xb1, 0x31, 0x40, 0x75, 0x46,
	0x24, 0x5d, 0xc0, 0x2e, 0xab, 0x54, 0x70, 0x27, 0x59, 0x6a, 0x35, 0x14, 0xf7, 0xdd, 0xc5, 0x41,
	0x4d, 0x8d, 0x2a, 0xd1, 0x46, 0x15, 0x22, 0x04, 0xc8, 0x3c, 0xaf, 0x5c, 0x70, 0x8d, 0x2c, 0xb5,
	0xfe, 0x2b, 0x6e, 0xef, 0xe0, 0xe0, 0x5f, 0x21, 0xfd, 0xd8, 0xeb, 0xaf, 0x36, 0xa6, 0xbe, 0xde,
	0x98, 0xfa, 0xdb, 0xc6, 0xd4, 0x5f, 0xb6, 0xa6, 0xb6, 0xde, 0x9a, 0xda, 0xeb, 0xd6, 0xd4, 0xee,
	0x3a, 0xcf, 0x87, 0x2a, 0x96, 0xcb, 0x39, 0x88, 0xf0, 0x6f, 0x51, 0xcf, 0xf9, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x03, 0xda, 0x49, 0x0e, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenInfos) > 0 {
		for iNdEx := len(m.TokenInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ParachainIBCTokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParachainIBCTokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParachainIBCTokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NativeDenom) > 0 {
		i -= len(m.NativeDenom)
		copy(dAtA[i:], m.NativeDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.NativeDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenInfos) > 0 {
		for _, e := range m.TokenInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ParachainIBCTokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.NativeDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenInfos = append(m.TokenInfos, ParachainIBCTokenInfo{})
			if err := m.TokenInfos[len(m.TokenInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ParachainIBCTokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ParachainIBCTokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParachainIBCTokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
