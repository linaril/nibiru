// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: perp/amm/v1/event.proto

package types

import (
	fmt "fmt"
	github_com_NibiruChain_nibiru_x_common_asset "github.com/NibiruChain/nibiru/x/common/asset"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ReserveSnapshotSavedEvent struct {
	Pair         github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,1,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	QuoteReserve github_com_cosmos_cosmos_sdk_types.Dec            `protobuf:"bytes,2,opt,name=quote_reserve,json=quoteReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_reserve"`
	BaseReserve  github_com_cosmos_cosmos_sdk_types.Dec            `protobuf:"bytes,3,opt,name=base_reserve,json=baseReserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_reserve"`
	// MarkPrice at the end of the block.
	// (instantaneous) markPrice := quoteReserve / baseReserve
	MarkPrice      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=mark_price,json=markPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mark_price"`
	BlockHeight    int64                                  `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	BlockTimestamp time.Time                              `protobuf:"bytes,6,opt,name=block_timestamp,json=blockTimestamp,proto3,stdtime" json:"block_timestamp"`
}

func (m *ReserveSnapshotSavedEvent) Reset()         { *m = ReserveSnapshotSavedEvent{} }
func (m *ReserveSnapshotSavedEvent) String() string { return proto.CompactTextString(m) }
func (*ReserveSnapshotSavedEvent) ProtoMessage()    {}
func (*ReserveSnapshotSavedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ed53a6cacb22fd, []int{0}
}
func (m *ReserveSnapshotSavedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReserveSnapshotSavedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReserveSnapshotSavedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReserveSnapshotSavedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveSnapshotSavedEvent.Merge(m, src)
}
func (m *ReserveSnapshotSavedEvent) XXX_Size() int {
	return m.Size()
}
func (m *ReserveSnapshotSavedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveSnapshotSavedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveSnapshotSavedEvent proto.InternalMessageInfo

func (m *ReserveSnapshotSavedEvent) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ReserveSnapshotSavedEvent) GetBlockTimestamp() time.Time {
	if m != nil {
		return m.BlockTimestamp
	}
	return time.Time{}
}

// A swap on the perp.amm represented by 'pair'.
// Amounts are negative or positive base on the perspective of the pool, i.e.
// a negative quote means the trader has gained quote and the perp.amm lost quote.
type SwapEvent struct {
	Pair github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,1,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	// delta in the quote reserves of the perp.amm
	QuoteAmount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=quote_amount,json=quoteAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quote_amount"`
	// delta in the base reserves of the perp.amm
	BaseAmount github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base_amount,json=baseAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_amount"`
}

func (m *SwapEvent) Reset()         { *m = SwapEvent{} }
func (m *SwapEvent) String() string { return proto.CompactTextString(m) }
func (*SwapEvent) ProtoMessage()    {}
func (*SwapEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ed53a6cacb22fd, []int{1}
}
func (m *SwapEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapEvent.Merge(m, src)
}
func (m *SwapEvent) XXX_Size() int {
	return m.Size()
}
func (m *SwapEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SwapEvent proto.InternalMessageInfo

type MarkPriceChangedEvent struct {
	Pair           github_com_NibiruChain_nibiru_x_common_asset.Pair `protobuf:"bytes,1,opt,name=pair,proto3,customtype=github.com/NibiruChain/nibiru/x/common/asset.Pair" json:"pair"`
	Price          github_com_cosmos_cosmos_sdk_types.Dec            `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	BlockTimestamp time.Time                                         `protobuf:"bytes,3,opt,name=block_timestamp,json=blockTimestamp,proto3,stdtime" json:"block_timestamp"`
}

func (m *MarkPriceChangedEvent) Reset()         { *m = MarkPriceChangedEvent{} }
func (m *MarkPriceChangedEvent) String() string { return proto.CompactTextString(m) }
func (*MarkPriceChangedEvent) ProtoMessage()    {}
func (*MarkPriceChangedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ed53a6cacb22fd, []int{2}
}
func (m *MarkPriceChangedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarkPriceChangedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarkPriceChangedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarkPriceChangedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkPriceChangedEvent.Merge(m, src)
}
func (m *MarkPriceChangedEvent) XXX_Size() int {
	return m.Size()
}
func (m *MarkPriceChangedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkPriceChangedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MarkPriceChangedEvent proto.InternalMessageInfo

func (m *MarkPriceChangedEvent) GetBlockTimestamp() time.Time {
	if m != nil {
		return m.BlockTimestamp
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*ReserveSnapshotSavedEvent)(nil), "nibiru.perp.amm.v1.ReserveSnapshotSavedEvent")
	proto.RegisterType((*SwapEvent)(nil), "nibiru.perp.amm.v1.SwapEvent")
	proto.RegisterType((*MarkPriceChangedEvent)(nil), "nibiru.perp.amm.v1.MarkPriceChangedEvent")
}

func init() { proto.RegisterFile("perp/amm/v1/event.proto", fileDescriptor_b5ed53a6cacb22fd) }

var fileDescriptor_b5ed53a6cacb22fd = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x6b, 0xba, 0x4d, 0xd4, 0x2d, 0x20, 0x45, 0x20, 0xb2, 0x1e, 0xd2, 0xd2, 0x03, 0xea,
	0x05, 0x5b, 0x81, 0x13, 0x47, 0xba, 0x21, 0xed, 0x52, 0xd8, 0x52, 0x4e, 0x5c, 0x22, 0x27, 0x33,
	0x89, 0xd5, 0x39, 0x0e, 0xb6, 0x13, 0xe0, 0xbf, 0xd8, 0x9f, 0xb5, 0xe3, 0x24, 0x2e, 0x88, 0xc3,
	0x40, 0xed, 0x9f, 0x81, 0x84, 0x90, 0xed, 0xa4, 0xe2, 0x80, 0x84, 0x14, 0x69, 0xa7, 0xe4, 0xfd,
	0xfa, 0xf8, 0xbd, 0xf7, 0xb5, 0x0c, 0x1f, 0x97, 0x54, 0x96, 0x98, 0x70, 0x8e, 0xeb, 0x10, 0xd3,
	0x9a, 0x16, 0x1a, 0x95, 0x52, 0x68, 0xe1, 0x79, 0x05, 0x4b, 0x98, 0xac, 0x90, 0x89, 0x23, 0xc2,
	0x39, 0xaa, 0xc3, 0x71, 0x90, 0x0a, 0xc5, 0x85, 0xc2, 0x09, 0x51, 0x14, 0xd7, 0x61, 0x42, 0x35,
	0x09, 0x71, 0x2a, 0x58, 0xe1, 0x6a, 0xc6, 0x87, 0x2e, 0x1e, 0x5b, 0x0b, 0x3b, 0xa3, 0x09, 0x3d,
	0xcc, 0x44, 0x26, 0x9c, 0xdf, 0xfc, 0x35, 0xde, 0x49, 0x26, 0x44, 0x76, 0x41, 0xb1, 0xb5, 0x92,
	0xea, 0x03, 0xd6, 0x8c, 0x53, 0xa5, 0x09, 0x2f, 0x5d, 0xc2, 0xec, 0x6b, 0x1f, 0x1e, 0x46, 0x54,
	0x51, 0x59, 0xd3, 0x55, 0x41, 0x4a, 0x95, 0x0b, 0xbd, 0x22, 0x35, 0x3d, 0x7f, 0x6d, 0x3a, 0xf5,
	0x96, 0x70, 0xaf, 0x24, 0x4c, 0xfa, 0x60, 0x0a, 0xe6, 0x83, 0xc5, 0xcb, 0xab, 0x9b, 0x49, 0xef,
	0xfb, 0xcd, 0x24, 0xcc, 0x98, 0xce, 0xab, 0x04, 0xa5, 0x82, 0xe3, 0x37, 0x76, 0x88, 0xa3, 0x9c,
	0xb0, 0x02, 0xbb, 0x81, 0xf0, 0x67, 0x9c, 0x0a, 0xce, 0x45, 0x81, 0x89, 0x52, 0x54, 0xa3, 0x53,
	0xc2, 0x64, 0x64, 0x31, 0xde, 0x0a, 0xde, 0xfb, 0x58, 0x09, 0x4d, 0x63, 0xe9, 0x4e, 0xf4, 0xef,
	0x58, 0x2e, 0x6a, 0xb8, 0x4f, 0xff, 0xe2, 0x36, 0x8b, 0x70, 0x9f, 0x67, 0xea, 0x7c, 0x8d, 0xf5,
	0x97, 0x92, 0x2a, 0x74, 0x4c, 0xd3, 0x68, 0x64, 0x21, 0x4d, 0xd7, 0xde, 0x19, 0x1c, 0x99, 0x75,
	0xed, 0x98, 0xfd, 0x4e, 0xcc, 0xa1, 0x61, 0xb4, 0xc8, 0x25, 0x84, 0x9c, 0xc8, 0x75, 0x5c, 0x4a,
	0x96, 0x52, 0x7f, 0xaf, 0x13, 0x70, 0x60, 0x08, 0xa7, 0x06, 0xe0, 0x3d, 0x81, 0xa3, 0xe4, 0x42,
	0xa4, 0xeb, 0x38, 0xa7, 0x2c, 0xcb, 0xb5, 0xbf, 0x3f, 0x05, 0xf3, 0x7e, 0x34, 0xb4, 0xbe, 0x13,
	0xeb, 0xf2, 0x96, 0xf0, 0x81, 0x4b, 0xd9, 0xe9, 0xe3, 0x1f, 0x4c, 0xc1, 0x7c, 0xf8, 0x7c, 0x8c,
	0x9c, 0x82, 0xa8, 0x55, 0x10, 0xbd, 0x6b, 0x33, 0x16, 0x77, 0x4d, 0x4b, 0x97, 0x3f, 0x26, 0x20,
	0xba, 0x6f, 0x8b, 0x77, 0x91, 0xd9, 0x6f, 0x00, 0x07, 0xab, 0x4f, 0xa4, 0xbc, 0x15, 0x15, 0xcf,
	0xa0, 0x13, 0x20, 0x26, 0x5c, 0x54, 0x85, 0xee, 0x28, 0xe2, 0xd0, 0x32, 0x5e, 0x59, 0x84, 0xf7,
	0x16, 0xda, 0xfd, 0xb7, 0xc4, 0x6e, 0x12, 0x42, 0x83, 0x70, 0xc0, 0xd9, 0x2f, 0x00, 0x1f, 0x2d,
	0x5b, 0x01, 0x8e, 0x72, 0x52, 0x64, 0xb7, 0x74, 0xa5, 0x8f, 0xe1, 0xbe, 0xbb, 0x25, 0xdd, 0xb6,
	0xe0, 0x8a, 0xff, 0x25, 0x7f, 0xbf, 0xbb, 0xfc, 0x8b, 0x93, 0xab, 0x4d, 0x00, 0xae, 0x37, 0x01,
	0xf8, 0xb9, 0x09, 0xc0, 0xe5, 0x36, 0xe8, 0x5d, 0x6f, 0x83, 0xde, 0xb7, 0x6d, 0xd0, 0x7b, 0x8f,
	0xfe, 0x37, 0xe7, 0xee, 0xb5, 0xb2, 0x3d, 0x26, 0x07, 0xf6, 0xdc, 0x17, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x26, 0xe5, 0xb0, 0xec, 0xc6, 0x04, 0x00, 0x00,
}

func (m *ReserveSnapshotSavedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReserveSnapshotSavedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReserveSnapshotSavedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BlockTimestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTimestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvent(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if m.BlockHeight != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MarkPrice.Size()
		i -= size
		if _, err := m.MarkPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BaseReserve.Size()
		i -= size
		if _, err := m.BaseReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.QuoteReserve.Size()
		i -= size
		if _, err := m.QuoteReserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *SwapEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BaseAmount.Size()
		i -= size
		if _, err := m.BaseAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.QuoteAmount.Size()
		i -= size
		if _, err := m.QuoteAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MarkPriceChangedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarkPriceChangedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarkPriceChangedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BlockTimestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTimestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEvent(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Pair.Size()
		i -= size
		if _, err := m.Pair.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReserveSnapshotSavedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pair.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.QuoteReserve.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.BaseReserve.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.MarkPrice.Size()
	n += 1 + l + sovEvent(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovEvent(uint64(m.BlockHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTimestamp)
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *SwapEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pair.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.QuoteAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.BaseAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *MarkPriceChangedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pair.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTimestamp)
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReserveSnapshotSavedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: ReserveSnapshotSavedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReserveSnapshotSavedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseReserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseReserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarkPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MarkPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BlockTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *SwapEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: SwapEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *MarkPriceChangedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: MarkPriceChangedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarkPriceChangedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BlockTimestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
