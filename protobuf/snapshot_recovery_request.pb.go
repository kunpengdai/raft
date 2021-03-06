// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snapshot_recovery_request.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SnapshotRecoveryRequest struct {
	LeaderName           string                          `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex            uint64                          `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm             uint64                          `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
	Peers                []*SnapshotRecoveryRequest_Peer `protobuf:"bytes,4,rep,name=Peers" json:"Peers,omitempty"`
	State                []byte                          `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *SnapshotRecoveryRequest) Reset()         { *m = SnapshotRecoveryRequest{} }
func (m *SnapshotRecoveryRequest) String() string { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest) ProtoMessage()    {}
func (*SnapshotRecoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_recovery_request_92f916293acea43e, []int{0}
}
func (m *SnapshotRecoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRecoveryRequest.Unmarshal(m, b)
}
func (m *SnapshotRecoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRecoveryRequest.Marshal(b, m, deterministic)
}
func (dst *SnapshotRecoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRecoveryRequest.Merge(dst, src)
}
func (m *SnapshotRecoveryRequest) XXX_Size() int {
	return xxx_messageInfo_SnapshotRecoveryRequest.Size(m)
}
func (m *SnapshotRecoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRecoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRecoveryRequest proto.InternalMessageInfo

func (m *SnapshotRecoveryRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *SnapshotRecoveryRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetPeers() []*SnapshotRecoveryRequest_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *SnapshotRecoveryRequest) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type SnapshotRecoveryRequest_Peer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	ConnectionString     string   `protobuf:"bytes,2,opt,name=ConnectionString" json:"ConnectionString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapshotRecoveryRequest_Peer) Reset()         { *m = SnapshotRecoveryRequest_Peer{} }
func (m *SnapshotRecoveryRequest_Peer) String() string { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest_Peer) ProtoMessage()    {}
func (*SnapshotRecoveryRequest_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_recovery_request_92f916293acea43e, []int{0, 0}
}
func (m *SnapshotRecoveryRequest_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotRecoveryRequest_Peer.Unmarshal(m, b)
}
func (m *SnapshotRecoveryRequest_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotRecoveryRequest_Peer.Marshal(b, m, deterministic)
}
func (dst *SnapshotRecoveryRequest_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotRecoveryRequest_Peer.Merge(dst, src)
}
func (m *SnapshotRecoveryRequest_Peer) XXX_Size() int {
	return xxx_messageInfo_SnapshotRecoveryRequest_Peer.Size(m)
}
func (m *SnapshotRecoveryRequest_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotRecoveryRequest_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotRecoveryRequest_Peer proto.InternalMessageInfo

func (m *SnapshotRecoveryRequest_Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnapshotRecoveryRequest_Peer) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

func init() {
	proto.RegisterType((*SnapshotRecoveryRequest)(nil), "protobuf.SnapshotRecoveryRequest")
	proto.RegisterType((*SnapshotRecoveryRequest_Peer)(nil), "protobuf.SnapshotRecoveryRequest.Peer")
}

func init() {
	proto.RegisterFile("snapshot_recovery_request.proto", fileDescriptor_snapshot_recovery_request_92f916293acea43e)
}

var fileDescriptor_snapshot_recovery_request_92f916293acea43e = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x76, 0x57, 0xba, 0xa3, 0x07, 0x19, 0x04, 0x43, 0x11, 0x0d, 0x1e, 0x24, 0x78,
	0xc8, 0x41, 0xaf, 0xde, 0x04, 0x41, 0x28, 0x22, 0x59, 0xef, 0x25, 0x6d, 0x47, 0xed, 0xa1, 0x49,
	0x9d, 0x4c, 0x45, 0x7f, 0x85, 0x7f, 0x59, 0x36, 0x5b, 0x75, 0x41, 0x7a, 0x4a, 0xde, 0x7b, 0xdf,
	0xc0, 0x07, 0x67, 0x29, 0xf8, 0x75, 0x7a, 0x8d, 0x32, 0x65, 0x9a, 0xc7, 0x77, 0xe2, 0xcf, 0x29,
	0xd3, 0xdb, 0x86, 0x92, 0xd8, 0x35, 0x47, 0x89, 0x38, 0xca, 0xcf, 0x6c, 0xf3, 0x7c, 0xfe, 0x35,
	0x80, 0xe3, 0x66, 0x4b, 0xbb, 0x2d, 0xec, 0x3a, 0x16, 0x4f, 0x01, 0x26, 0xe4, 0x17, 0xc4, 0x0f,
	0x7e, 0x45, 0xaa, 0xd0, 0x85, 0xa9, 0x5d, 0xaf, 0xc1, 0x13, 0xa8, 0x27, 0x3e, 0xc9, 0x7d, 0x58,
	0xd0, 0x87, 0x1a, 0xe8, 0xc2, 0x94, 0xee, 0xaf, 0xc0, 0x31, 0x8c, 0xda, 0xf0, 0x44, 0xbc, 0x52,
	0xc3, 0x3c, 0xfe, 0x66, 0xbc, 0x81, 0xea, 0x91, 0x88, 0x93, 0x2a, 0xf5, 0xd0, 0xec, 0x5f, 0x5d,
	0xd8, 0x1f, 0x1f, 0xbb, 0xc3, 0xc5, 0xb6, 0xb8, 0xeb, 0x8e, 0xf0, 0x08, 0xaa, 0x46, 0xbc, 0x90,
	0xaa, 0x74, 0x61, 0x0e, 0x5c, 0x17, 0xc6, 0x77, 0x50, 0xb6, 0x33, 0x22, 0x94, 0x3d, 0xdf, 0xfc,
	0xc7, 0x4b, 0x38, 0xbc, 0x8d, 0x21, 0xd0, 0x5c, 0x96, 0x31, 0x34, 0xc2, 0xcb, 0xf0, 0x92, 0x85,
	0x6b, 0xf7, 0xaf, 0x9f, 0xed, 0x65, 0x97, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x7e,
	0x60, 0x97, 0x45, 0x01, 0x00, 0x00,
}
