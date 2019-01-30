// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcproto/stats.proto

// Copyright (c) 2018 Donovan Solms / MiningHQ. Licensed under the MIT license.
// See LICENSE in the root of this repository for details.

package rpcproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CPUStats contains the CPU stats
type CPUStats struct {
	ThreadsHashrate      []float64 `protobuf:"fixed64,1,rep,packed,name=ThreadsHashrate,proto3" json:"ThreadsHashrate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CPUStats) Reset()         { *m = CPUStats{} }
func (m *CPUStats) String() string { return proto.CompactTextString(m) }
func (*CPUStats) ProtoMessage()    {}
func (*CPUStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_f515951261dd90e2, []int{0}
}

func (m *CPUStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPUStats.Unmarshal(m, b)
}
func (m *CPUStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPUStats.Marshal(b, m, deterministic)
}
func (m *CPUStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUStats.Merge(m, src)
}
func (m *CPUStats) XXX_Size() int {
	return xxx_messageInfo_CPUStats.Size(m)
}
func (m *CPUStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUStats.DiscardUnknown(m)
}

var xxx_messageInfo_CPUStats proto.InternalMessageInfo

func (m *CPUStats) GetThreadsHashrate() []float64 {
	if m != nil {
		return m.ThreadsHashrate
	}
	return nil
}

// CPUStats contains all the stats for a miner
type MinerStats struct {
	// Key identifies this miner's configuration on MiningHQ
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// Hashrate is the current miner hashrate
	Hashrate float64 `protobuf:"fixed64,2,opt,name=Hashrate,proto3" json:"Hashrate,omitempty"`
	// MaxHashrate is the highest hashrate achieved since start
	MaxHashrate float64 `protobuf:"fixed64,3,opt,name=MaxHashrate,proto3" json:"MaxHashrate,omitempty"`
	// TotalHashes is the total number of hashes submitted
	TotalHashes uint64 `protobuf:"varint,4,opt,name=TotalHashes,proto3" json:"TotalHashes,omitempty"`
	// CurrentDifficulty is the current pool given difficulty
	CurrentDifficulty uint64 `protobuf:"varint,5,opt,name=CurrentDifficulty,proto3" json:"CurrentDifficulty,omitempty"`
	// TotalShares is the total amount of shares produced since start
	TotalShares uint32 `protobuf:"varint,6,opt,name=TotalShares,proto3" json:"TotalShares,omitempty"`
	// AcceptedShares is the number of pool accepted shares since start
	AcceptedShares uint32 `protobuf:"varint,7,opt,name=AcceptedShares,proto3" json:"AcceptedShares,omitempty"`
	// RejectedShares is the number of pool rejected shares
	RejectedShares uint32 `protobuf:"varint,8,opt,name=RejectedShares,proto3" json:"RejectedShares,omitempty"`
	// CPUs hold the CPU stats of the miner
	CPUs                 []*CPUStats `protobuf:"bytes,9,rep,name=CPUs,proto3" json:"CPUs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MinerStats) Reset()         { *m = MinerStats{} }
func (m *MinerStats) String() string { return proto.CompactTextString(m) }
func (*MinerStats) ProtoMessage()    {}
func (*MinerStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_f515951261dd90e2, []int{1}
}

func (m *MinerStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinerStats.Unmarshal(m, b)
}
func (m *MinerStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinerStats.Marshal(b, m, deterministic)
}
func (m *MinerStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinerStats.Merge(m, src)
}
func (m *MinerStats) XXX_Size() int {
	return xxx_messageInfo_MinerStats.Size(m)
}
func (m *MinerStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MinerStats.DiscardUnknown(m)
}

var xxx_messageInfo_MinerStats proto.InternalMessageInfo

func (m *MinerStats) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MinerStats) GetHashrate() float64 {
	if m != nil {
		return m.Hashrate
	}
	return 0
}

func (m *MinerStats) GetMaxHashrate() float64 {
	if m != nil {
		return m.MaxHashrate
	}
	return 0
}

func (m *MinerStats) GetTotalHashes() uint64 {
	if m != nil {
		return m.TotalHashes
	}
	return 0
}

func (m *MinerStats) GetCurrentDifficulty() uint64 {
	if m != nil {
		return m.CurrentDifficulty
	}
	return 0
}

func (m *MinerStats) GetTotalShares() uint32 {
	if m != nil {
		return m.TotalShares
	}
	return 0
}

func (m *MinerStats) GetAcceptedShares() uint32 {
	if m != nil {
		return m.AcceptedShares
	}
	return 0
}

func (m *MinerStats) GetRejectedShares() uint32 {
	if m != nil {
		return m.RejectedShares
	}
	return 0
}

func (m *MinerStats) GetCPUs() []*CPUStats {
	if m != nil {
		return m.CPUs
	}
	return nil
}

// StatsRequest requests stats from the miner
type StatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f515951261dd90e2, []int{2}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

// StateResponse is sent in response to a StatsRequest and periodically
type StatsResponse struct {
	// Stats contains the miner's stats
	Stats []*MinerStats `protobuf:"bytes,1,rep,name=Stats,proto3" json:"Stats,omitempty"`
	// MinerVersions returns the versions of the miner assigned
	MinerVersions        []string `protobuf:"bytes,2,rep,name=MinerVersions,proto3" json:"MinerVersions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsResponse) Reset()         { *m = StatsResponse{} }
func (m *StatsResponse) String() string { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()    {}
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f515951261dd90e2, []int{3}
}

func (m *StatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsResponse.Unmarshal(m, b)
}
func (m *StatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsResponse.Marshal(b, m, deterministic)
}
func (m *StatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsResponse.Merge(m, src)
}
func (m *StatsResponse) XXX_Size() int {
	return xxx_messageInfo_StatsResponse.Size(m)
}
func (m *StatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsResponse proto.InternalMessageInfo

func (m *StatsResponse) GetStats() []*MinerStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *StatsResponse) GetMinerVersions() []string {
	if m != nil {
		return m.MinerVersions
	}
	return nil
}

func init() {
	proto.RegisterType((*CPUStats)(nil), "rpcproto.CPUStats")
	proto.RegisterType((*MinerStats)(nil), "rpcproto.MinerStats")
	proto.RegisterType((*StatsRequest)(nil), "rpcproto.StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "rpcproto.StatsResponse")
}

func init() { proto.RegisterFile("rpcproto/stats.proto", fileDescriptor_f515951261dd90e2) }

var fileDescriptor_f515951261dd90e2 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xdf, 0x4a, 0x02, 0x41,
	0x14, 0xc6, 0x19, 0x57, 0x6d, 0x3d, 0xa6, 0xd5, 0xe0, 0xc5, 0xd0, 0xd5, 0xb0, 0x84, 0x0c, 0x11,
	0x06, 0xd6, 0x0b, 0x84, 0x5d, 0x04, 0x21, 0xc8, 0xa8, 0xdd, 0x4f, 0xeb, 0x11, 0x37, 0x64, 0x77,
	0x9b, 0x33, 0x42, 0xbe, 0x60, 0xcf, 0x15, 0x3b, 0xab, 0xeb, 0x9f, 0xee, 0xe6, 0xfc, 0xbe, 0xdf,
	0x39, 0x17, 0xdf, 0x40, 0xcf, 0xe6, 0x71, 0x6e, 0x33, 0x97, 0x3d, 0x92, 0x33, 0x8e, 0x06, 0xfe,
	0xcd, 0xc3, 0x3d, 0x8d, 0x9e, 0x21, 0x1c, 0x4d, 0xe6, 0xd3, 0x22, 0xe3, 0x0a, 0xae, 0x66, 0x2b,
	0x8b, 0x66, 0x41, 0x6f, 0x86, 0x56, 0xd6, 0x38, 0x14, 0x4c, 0x06, 0x8a, 0xe9, 0x73, 0x1c, 0xfd,
	0xd6, 0x00, 0xc6, 0x49, 0x8a, 0xb6, 0x5c, 0xbc, 0x86, 0xe0, 0x1d, 0xb7, 0x82, 0x49, 0xa6, 0x5a,
	0xba, 0x78, 0xf2, 0x5b, 0x08, 0xab, 0x1b, 0x35, 0xc9, 0x14, 0xd3, 0xd5, 0xcc, 0x25, 0xb4, 0xc7,
	0xe6, 0xa7, 0x8a, 0x03, 0x1f, 0x1f, 0xa3, 0xc2, 0x98, 0x65, 0xce, 0xac, 0x0b, 0x80, 0x24, 0xea,
	0x92, 0xa9, 0xba, 0x3e, 0x46, 0xfc, 0x01, 0x6e, 0x46, 0x1b, 0x6b, 0x31, 0x75, 0xaf, 0xc9, 0x72,
	0x99, 0xc4, 0x9b, 0xb5, 0xdb, 0x8a, 0x86, 0xf7, 0xfe, 0x07, 0xd5, 0xbd, 0xe9, 0xca, 0x58, 0x24,
	0xd1, 0x94, 0x4c, 0x75, 0xf4, 0x31, 0xe2, 0x7d, 0xe8, 0xbe, 0xc4, 0x31, 0xe6, 0x0e, 0x17, 0x3b,
	0xe9, 0xc2, 0x4b, 0x67, 0xb4, 0xf0, 0x34, 0x7e, 0x61, 0x7c, 0xf0, 0xc2, 0xd2, 0x3b, 0xa5, 0xbc,
	0x0f, 0xf5, 0xd1, 0x64, 0x4e, 0xa2, 0x25, 0x03, 0xd5, 0x1e, 0xf2, 0xc1, 0xbe, 0xef, 0xc1, 0xbe,
	0x6c, 0xed, 0xf3, 0xa8, 0x0b, 0x97, 0xe5, 0x88, 0xdf, 0x1b, 0x24, 0x17, 0x19, 0xe8, 0xec, 0x66,
	0xca, 0xb3, 0x94, 0x90, 0xdf, 0x43, 0xc3, 0x03, 0xff, 0x13, 0xed, 0x61, 0xef, 0x70, 0xe9, 0xd0,
	0xbf, 0x2e, 0x15, 0x7e, 0x07, 0x1d, 0x0f, 0x3f, 0xd0, 0x52, 0x92, 0xa5, 0x24, 0x6a, 0x32, 0x50,
	0x2d, 0x7d, 0x0a, 0x3f, 0x9b, 0x7e, 0xfd, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x15, 0x1d,
	0xe2, 0x1a, 0x02, 0x00, 0x00,
}
