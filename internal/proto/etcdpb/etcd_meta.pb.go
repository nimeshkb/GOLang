// Code generated by protoc-gen-go. DO NOT EDIT.
// source: etcd_meta.proto

package etcdpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	schemapb "github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TenantMeta struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NumQueryNodes        uint64   `protobuf:"varint,2,opt,name=num_query_nodes,json=numQueryNodes,proto3" json:"num_query_nodes,omitempty"`
	InsertChannelIds     []string `protobuf:"bytes,3,rep,name=insert_channel_ids,json=insertChannelIds,proto3" json:"insert_channel_ids,omitempty"`
	QueryChannelId       string   `protobuf:"bytes,4,opt,name=query_channel_id,json=queryChannelId,proto3" json:"query_channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TenantMeta) Reset()         { *m = TenantMeta{} }
func (m *TenantMeta) String() string { return proto.CompactTextString(m) }
func (*TenantMeta) ProtoMessage()    {}
func (*TenantMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{0}
}

func (m *TenantMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantMeta.Unmarshal(m, b)
}
func (m *TenantMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantMeta.Marshal(b, m, deterministic)
}
func (m *TenantMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantMeta.Merge(m, src)
}
func (m *TenantMeta) XXX_Size() int {
	return xxx_messageInfo_TenantMeta.Size(m)
}
func (m *TenantMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TenantMeta proto.InternalMessageInfo

func (m *TenantMeta) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TenantMeta) GetNumQueryNodes() uint64 {
	if m != nil {
		return m.NumQueryNodes
	}
	return 0
}

func (m *TenantMeta) GetInsertChannelIds() []string {
	if m != nil {
		return m.InsertChannelIds
	}
	return nil
}

func (m *TenantMeta) GetQueryChannelId() string {
	if m != nil {
		return m.QueryChannelId
	}
	return ""
}

type ProxyMeta struct {
	Id                   uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              *commonpb.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ResultChannelIds     []string          `protobuf:"bytes,3,rep,name=result_channel_ids,json=resultChannelIds,proto3" json:"result_channel_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProxyMeta) Reset()         { *m = ProxyMeta{} }
func (m *ProxyMeta) String() string { return proto.CompactTextString(m) }
func (*ProxyMeta) ProtoMessage()    {}
func (*ProxyMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{1}
}

func (m *ProxyMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyMeta.Unmarshal(m, b)
}
func (m *ProxyMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyMeta.Marshal(b, m, deterministic)
}
func (m *ProxyMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyMeta.Merge(m, src)
}
func (m *ProxyMeta) XXX_Size() int {
	return xxx_messageInfo_ProxyMeta.Size(m)
}
func (m *ProxyMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyMeta proto.InternalMessageInfo

func (m *ProxyMeta) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProxyMeta) GetAddress() *commonpb.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ProxyMeta) GetResultChannelIds() []string {
	if m != nil {
		return m.ResultChannelIds
	}
	return nil
}

type CollectionMeta struct {
	Id                   uint64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Schema               *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime           uint64                     `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	SegmentIds           []uint64                   `protobuf:"varint,4,rep,packed,name=segment_ids,json=segmentIds,proto3" json:"segment_ids,omitempty"`
	PartitionTags        []string                   `protobuf:"bytes,5,rep,name=partition_tags,json=partitionTags,proto3" json:"partition_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionMeta) Reset()         { *m = CollectionMeta{} }
func (m *CollectionMeta) String() string { return proto.CompactTextString(m) }
func (*CollectionMeta) ProtoMessage()    {}
func (*CollectionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{2}
}

func (m *CollectionMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionMeta.Unmarshal(m, b)
}
func (m *CollectionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionMeta.Marshal(b, m, deterministic)
}
func (m *CollectionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionMeta.Merge(m, src)
}
func (m *CollectionMeta) XXX_Size() int {
	return xxx_messageInfo_CollectionMeta.Size(m)
}
func (m *CollectionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionMeta proto.InternalMessageInfo

func (m *CollectionMeta) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CollectionMeta) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionMeta) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CollectionMeta) GetSegmentIds() []uint64 {
	if m != nil {
		return m.SegmentIds
	}
	return nil
}

func (m *CollectionMeta) GetPartitionTags() []string {
	if m != nil {
		return m.PartitionTags
	}
	return nil
}

type SegmentMeta struct {
	SegmentId            uint64   `protobuf:"varint,1,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	CollectionId         uint64   `protobuf:"varint,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	PartitionTag         string   `protobuf:"bytes,3,opt,name=partition_tag,json=partitionTag,proto3" json:"partition_tag,omitempty"`
	ChannelStart         int32    `protobuf:"varint,4,opt,name=channel_start,json=channelStart,proto3" json:"channel_start,omitempty"`
	ChannelEnd           int32    `protobuf:"varint,5,opt,name=channel_end,json=channelEnd,proto3" json:"channel_end,omitempty"`
	OpenTime             uint64   `protobuf:"varint,6,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	CloseTime            uint64   `protobuf:"varint,7,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	NumRows              int64    `protobuf:"varint,8,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentMeta) Reset()         { *m = SegmentMeta{} }
func (m *SegmentMeta) String() string { return proto.CompactTextString(m) }
func (*SegmentMeta) ProtoMessage()    {}
func (*SegmentMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_975d306d62b73e88, []int{3}
}

func (m *SegmentMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentMeta.Unmarshal(m, b)
}
func (m *SegmentMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentMeta.Marshal(b, m, deterministic)
}
func (m *SegmentMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentMeta.Merge(m, src)
}
func (m *SegmentMeta) XXX_Size() int {
	return xxx_messageInfo_SegmentMeta.Size(m)
}
func (m *SegmentMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentMeta proto.InternalMessageInfo

func (m *SegmentMeta) GetSegmentId() uint64 {
	if m != nil {
		return m.SegmentId
	}
	return 0
}

func (m *SegmentMeta) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *SegmentMeta) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *SegmentMeta) GetChannelStart() int32 {
	if m != nil {
		return m.ChannelStart
	}
	return 0
}

func (m *SegmentMeta) GetChannelEnd() int32 {
	if m != nil {
		return m.ChannelEnd
	}
	return 0
}

func (m *SegmentMeta) GetOpenTime() uint64 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

func (m *SegmentMeta) GetCloseTime() uint64 {
	if m != nil {
		return m.CloseTime
	}
	return 0
}

func (m *SegmentMeta) GetNumRows() int64 {
	if m != nil {
		return m.NumRows
	}
	return 0
}

func init() {
	proto.RegisterType((*TenantMeta)(nil), "milvus.proto.etcd.TenantMeta")
	proto.RegisterType((*ProxyMeta)(nil), "milvus.proto.etcd.ProxyMeta")
	proto.RegisterType((*CollectionMeta)(nil), "milvus.proto.etcd.CollectionMeta")
	proto.RegisterType((*SegmentMeta)(nil), "milvus.proto.etcd.SegmentMeta")
}

func init() { proto.RegisterFile("etcd_meta.proto", fileDescriptor_975d306d62b73e88) }

var fileDescriptor_975d306d62b73e88 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x7e, 0x67, 0xfa, 0xb1, 0x4b, 0x4e, 0x61, 0x61, 0x45, 0x55, 0xb4, 0x28, 0x07,
	0x68, 0x24, 0x90, 0xb8, 0x81, 0x80, 0x15, 0x07, 0x0e, 0x20, 0x48, 0x7b, 0xe2, 0x12, 0xb9, 0xf1,
	0xa8, 0xb5, 0x14, 0xdb, 0xc5, 0x76, 0x58, 0x76, 0x6f, 0xfc, 0x06, 0xae, 0xfc, 0x20, 0x7e, 0x16,
	0xf2, 0x07, 0x0d, 0x95, 0x7a, 0xcc, 0x33, 0xaf, 0xc7, 0xef, 0xbc, 0xe3, 0xc0, 0x19, 0x9a, 0x8a,
	0x96, 0x1c, 0x0d, 0x59, 0xee, 0x95, 0x34, 0x32, 0xb9, 0xc7, 0x59, 0xfd, 0xbd, 0xd1, 0xfe, 0x6b,
	0x69, 0xab, 0x17, 0x93, 0x4a, 0x72, 0x2e, 0x85, 0x47, 0x17, 0x13, 0x5d, 0xed, 0x90, 0x07, 0xf9,
	0xe2, 0x77, 0x04, 0xb0, 0x46, 0x41, 0x84, 0xf9, 0x88, 0x86, 0x24, 0x33, 0xe8, 0x30, 0x9a, 0x46,
	0xf3, 0x28, 0xeb, 0x15, 0x1d, 0x46, 0x93, 0x27, 0x70, 0x26, 0x1a, 0x5e, 0x7e, 0x6b, 0x50, 0xdd,
	0x96, 0x42, 0x52, 0xd4, 0x69, 0xc7, 0x15, 0xa7, 0xa2, 0xe1, 0x5f, 0x2c, 0xfd, 0x64, 0x61, 0xf2,
	0x14, 0x12, 0x26, 0x34, 0x2a, 0x53, 0x56, 0x3b, 0x22, 0x04, 0xd6, 0x25, 0xa3, 0x3a, 0xed, 0xce,
	0xbb, 0x59, 0x5c, 0x9c, 0xfb, 0xca, 0xb5, 0x2f, 0x7c, 0xa0, 0x3a, 0xc9, 0xe0, 0xdc, 0x77, 0x6c,
	0xc5, 0x69, 0x6f, 0x1e, 0x65, 0x71, 0x31, 0x73, 0xfc, 0x20, 0x5d, 0xfc, 0x8c, 0x20, 0xfe, 0xac,
	0xe4, 0x8f, 0xdb, 0x93, 0xee, 0x5e, 0xc2, 0x90, 0x50, 0xaa, 0x50, 0x7b, 0x57, 0xe3, 0xe7, 0x0f,
	0x97, 0x47, 0xd3, 0x87, 0xb9, 0xdf, 0x7a, 0x4d, 0xf1, 0x4f, 0x6c, 0xdd, 0x2a, 0xd4, 0x4d, 0x7d,
	0xd2, 0xad, 0xaf, 0xb4, 0x6e, 0x17, 0x7f, 0x22, 0x98, 0x5d, 0xcb, 0xba, 0xc6, 0xca, 0x30, 0x29,
	0x4e, 0x1a, 0x79, 0x05, 0x03, 0x9f, 0x6a, 0xf0, 0x71, 0x75, 0xec, 0x23, 0x24, 0xde, 0x36, 0x59,
	0x39, 0x50, 0x84, 0x43, 0xc9, 0x23, 0x18, 0x57, 0x0a, 0x89, 0xc1, 0xd2, 0x30, 0x8e, 0x69, 0xd7,
	0xf5, 0x05, 0x8f, 0xd6, 0x8c, 0xa3, 0x15, 0x68, 0xdc, 0x72, 0x14, 0xc6, 0x39, 0xed, 0xcd, 0xbb,
	0x56, 0x10, 0x90, 0x4d, 0xf4, 0x0a, 0x66, 0x7b, 0xa2, 0x0c, 0xb3, 0xcd, 0x4b, 0x43, 0xb6, 0x3a,
	0xed, 0xbb, 0x69, 0xa6, 0x07, 0xba, 0x26, 0x5b, 0xbd, 0xf8, 0xd5, 0x81, 0xf1, 0xca, 0x9f, 0x72,
	0x73, 0x5c, 0x02, 0xb4, 0x7d, 0xc3, 0x3c, 0xf1, 0xa1, 0x6d, 0xf2, 0x18, 0xa6, 0xd5, 0xc1, 0xb3,
	0x55, 0xf8, 0xdd, 0x4f, 0x5a, 0xe8, 0x45, 0x47, 0x57, 0x3b, 0xfb, 0x71, 0x31, 0xf9, 0xff, 0x66,
	0xd7, 0x29, 0x44, 0xad, 0x0d, 0x51, 0xc6, 0xad, 0xbb, 0x5f, 0x4c, 0x02, 0x5c, 0x59, 0xe6, 0x62,
	0x08, 0x22, 0x14, 0x34, 0xed, 0x3b, 0x09, 0x04, 0xf4, 0x5e, 0xd0, 0xe4, 0x01, 0xc4, 0x72, 0x8f,
	0xc2, 0xa7, 0x34, 0x70, 0x5e, 0x46, 0x16, 0xb8, 0x8c, 0x2e, 0x01, 0xaa, 0x5a, 0xea, 0x90, 0xe1,
	0xd0, 0xcf, 0xe2, 0x88, 0x2b, 0xdf, 0x87, 0x91, 0x7d, 0xc9, 0x4a, 0xde, 0xe8, 0x74, 0x34, 0x8f,
	0xb2, 0x6e, 0x31, 0x14, 0x0d, 0x2f, 0xe4, 0x8d, 0x7e, 0xf7, 0xe6, 0xeb, 0xeb, 0x2d, 0x33, 0xbb,
	0x66, 0x63, 0x1f, 0x4c, 0x7e, 0xc7, 0xea, 0x9a, 0xdd, 0x19, 0xac, 0x76, 0xb9, 0x5f, 0xe2, 0x33,
	0xca, 0xb4, 0x51, 0x6c, 0xd3, 0x18, 0xa4, 0x39, 0x13, 0x06, 0x95, 0x20, 0x75, 0xee, 0x36, 0x9b,
	0xdb, 0xff, 0x6b, 0xbf, 0xd9, 0x0c, 0xdc, 0xd7, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x15, 0xf5, 0xbd, 0x8e, 0x03, 0x00, 0x00,
}