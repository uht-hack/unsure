// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uht.proto

package uhtpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/luno/reflex/reflexpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RoundStatus int32

const (
	RoundStatus_ROUND_STATUS_UNKNOWN   RoundStatus = 0
	RoundStatus_ROUND_STATUS_JOIN      RoundStatus = 1
	RoundStatus_ROUND_STATUS_JOINED    RoundStatus = 2
	RoundStatus_ROUND_STATUS_COLLECT   RoundStatus = 3
	RoundStatus_ROUND_STATUS_COLLECTED RoundStatus = 4
	RoundStatus_ROUND_STATUS_SUBMIT    RoundStatus = 5
	RoundStatus_ROUND_STATUS_SUBMITTED RoundStatus = 6
	RoundStatus_ROUND_STATUS_SUCCESS   RoundStatus = 7
	RoundStatus_ROUND_STATUS_FAILED    RoundStatus = 8
	RoundStatus_ROUND_STATUS_SENTINEL  RoundStatus = 9
)

var RoundStatus_name = map[int32]string{
	0: "ROUND_STATUS_UNKNOWN",
	1: "ROUND_STATUS_JOIN",
	2: "ROUND_STATUS_JOINED",
	3: "ROUND_STATUS_COLLECT",
	4: "ROUND_STATUS_COLLECTED",
	5: "ROUND_STATUS_SUBMIT",
	6: "ROUND_STATUS_SUBMITTED",
	7: "ROUND_STATUS_SUCCESS",
	8: "ROUND_STATUS_FAILED",
	9: "ROUND_STATUS_SENTINEL",
}
var RoundStatus_value = map[string]int32{
	"ROUND_STATUS_UNKNOWN":   0,
	"ROUND_STATUS_JOIN":      1,
	"ROUND_STATUS_JOINED":    2,
	"ROUND_STATUS_COLLECT":   3,
	"ROUND_STATUS_COLLECTED": 4,
	"ROUND_STATUS_SUBMIT":    5,
	"ROUND_STATUS_SUBMITTED": 6,
	"ROUND_STATUS_SUCCESS":   7,
	"ROUND_STATUS_FAILED":    8,
	"ROUND_STATUS_SENTINEL":  9,
}

func (x RoundStatus) String() string {
	return proto.EnumName(RoundStatus_name, int32(x))
}
func (RoundStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_uht_810ebf75a3bcb9d5, []int{0}
}

type Round struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MatchId              int64                `protobuf:"varint,2,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	Index                int64                `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Team                 string               `protobuf:"bytes,4,opt,name=team,proto3" json:"team,omitempty"`
	Status               RoundStatus          `protobuf:"varint,5,opt,name=status,proto3,enum=uhtpb.RoundStatus" json:"status,omitempty"`
	State                *RoundState          `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Error                string               `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Round) Reset()         { *m = Round{} }
func (m *Round) String() string { return proto.CompactTextString(m) }
func (*Round) ProtoMessage()    {}
func (*Round) Descriptor() ([]byte, []int) {
	return fileDescriptor_uht_810ebf75a3bcb9d5, []int{0}
}
func (m *Round) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Round.Unmarshal(m, b)
}
func (m *Round) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Round.Marshal(b, m, deterministic)
}
func (dst *Round) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Round.Merge(dst, src)
}
func (m *Round) XXX_Size() int {
	return xxx_messageInfo_Round.Size(m)
}
func (m *Round) XXX_DiscardUnknown() {
	xxx_messageInfo_Round.DiscardUnknown(m)
}

var xxx_messageInfo_Round proto.InternalMessageInfo

func (m *Round) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Round) GetMatchId() int64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *Round) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Round) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Round) GetStatus() RoundStatus {
	if m != nil {
		return m.Status
	}
	return RoundStatus_ROUND_STATUS_UNKNOWN
}

func (m *Round) GetState() *RoundState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Round) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Round) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Round) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type RoundState struct {
	Players              []*RoundPlayerState `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RoundState) Reset()         { *m = RoundState{} }
func (m *RoundState) String() string { return proto.CompactTextString(m) }
func (*RoundState) ProtoMessage()    {}
func (*RoundState) Descriptor() ([]byte, []int) {
	return fileDescriptor_uht_810ebf75a3bcb9d5, []int{1}
}
func (m *RoundState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundState.Unmarshal(m, b)
}
func (m *RoundState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundState.Marshal(b, m, deterministic)
}
func (dst *RoundState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundState.Merge(dst, src)
}
func (m *RoundState) XXX_Size() int {
	return xxx_messageInfo_RoundState.Size(m)
}
func (m *RoundState) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundState.DiscardUnknown(m)
}

var xxx_messageInfo_RoundState proto.InternalMessageInfo

func (m *RoundState) GetPlayers() []*RoundPlayerState {
	if m != nil {
		return m.Players
	}
	return nil
}

type RoundPlayerState struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rank                 int32            `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Parts                map[string]int32 `protobuf:"bytes,3,rep,name=parts,proto3" json:"parts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Included             bool             `protobuf:"varint,4,opt,name=included,proto3" json:"included,omitempty"`
	Collected            bool             `protobuf:"varint,5,opt,name=collected,proto3" json:"collected,omitempty"`
	Submitted            bool             `protobuf:"varint,6,opt,name=submitted,proto3" json:"submitted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RoundPlayerState) Reset()         { *m = RoundPlayerState{} }
func (m *RoundPlayerState) String() string { return proto.CompactTextString(m) }
func (*RoundPlayerState) ProtoMessage()    {}
func (*RoundPlayerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_uht_810ebf75a3bcb9d5, []int{2}
}
func (m *RoundPlayerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundPlayerState.Unmarshal(m, b)
}
func (m *RoundPlayerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundPlayerState.Marshal(b, m, deterministic)
}
func (dst *RoundPlayerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundPlayerState.Merge(dst, src)
}
func (m *RoundPlayerState) XXX_Size() int {
	return xxx_messageInfo_RoundPlayerState.Size(m)
}
func (m *RoundPlayerState) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundPlayerState.DiscardUnknown(m)
}

var xxx_messageInfo_RoundPlayerState proto.InternalMessageInfo

func (m *RoundPlayerState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoundPlayerState) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RoundPlayerState) GetParts() map[string]int32 {
	if m != nil {
		return m.Parts
	}
	return nil
}

func (m *RoundPlayerState) GetIncluded() bool {
	if m != nil {
		return m.Included
	}
	return false
}

func (m *RoundPlayerState) GetCollected() bool {
	if m != nil {
		return m.Collected
	}
	return false
}

func (m *RoundPlayerState) GetSubmitted() bool {
	if m != nil {
		return m.Submitted
	}
	return false
}

func init() {
	proto.RegisterType((*Round)(nil), "uhtpb.Round")
	proto.RegisterType((*RoundState)(nil), "uhtpb.RoundState")
	proto.RegisterType((*RoundPlayerState)(nil), "uhtpb.RoundPlayerState")
	proto.RegisterMapType((map[string]int32)(nil), "uhtpb.RoundPlayerState.PartsEntry")
	proto.RegisterEnum("uhtpb.RoundStatus", RoundStatus_name, RoundStatus_value)
}

func init() { proto.RegisterFile("uht.proto", fileDescriptor_uht_810ebf75a3bcb9d5) }

var fileDescriptor_uht_810ebf75a3bcb9d5 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0x9d, 0xd8, 0x53, 0xa9, 0x72, 0x97, 0x96, 0xba, 0x16, 0x12, 0x51, 0x2e, 0x44,
	0x15, 0x72, 0x44, 0xb8, 0x04, 0x2e, 0x28, 0x24, 0x46, 0x0a, 0x04, 0xa7, 0x5a, 0x3b, 0xe2, 0x18,
	0xf9, 0x63, 0x9b, 0x58, 0xf5, 0x97, 0xec, 0x35, 0x6a, 0xfe, 0x00, 0x67, 0xfe, 0x10, 0xff, 0x0d,
	0x79, 0xd7, 0x69, 0xda, 0x1a, 0xc4, 0xc9, 0x33, 0xf3, 0xde, 0x9b, 0xe7, 0x99, 0x59, 0x50, 0xab,
	0x2d, 0x35, 0xf3, 0x22, 0xa3, 0x19, 0x92, 0xab, 0x2d, 0xcd, 0x7d, 0xe3, 0xd5, 0x26, 0xcb, 0x36,
	0x31, 0x19, 0xb2, 0xa2, 0x5f, 0xdd, 0x0c, 0x69, 0x94, 0x90, 0x92, 0x7a, 0x49, 0xce, 0x79, 0xc6,
	0x9b, 0x4d, 0x44, 0xb7, 0x95, 0x6f, 0x06, 0x59, 0x32, 0x8c, 0xab, 0x34, 0x1b, 0x16, 0xe4, 0x26,
	0x26, 0x77, 0xcd, 0x27, 0xf7, 0x9b, 0x80, 0xb3, 0xfb, 0xbf, 0x45, 0x90, 0x71, 0x56, 0xa5, 0x21,
	0x3a, 0x01, 0x31, 0x0a, 0x75, 0xa1, 0x27, 0x0c, 0x24, 0x2c, 0x46, 0x21, 0xba, 0x04, 0x25, 0xf1,
	0x68, 0xb0, 0x5d, 0x47, 0xa1, 0x2e, 0xb2, 0x6a, 0x97, 0xe5, 0xf3, 0x10, 0x9d, 0x81, 0x1c, 0xa5,
	0x21, 0xb9, 0xd3, 0x25, 0x56, 0xe7, 0x09, 0x42, 0x70, 0x44, 0x89, 0x97, 0xe8, 0x47, 0x3d, 0x61,
	0xa0, 0x62, 0x16, 0xa3, 0x2b, 0xe8, 0x94, 0xd4, 0xa3, 0x55, 0xa9, 0xcb, 0x3d, 0x61, 0x70, 0x32,
	0x42, 0x26, 0x9b, 0xc2, 0x64, 0x96, 0x0e, 0x43, 0x70, 0xc3, 0x40, 0xaf, 0x41, 0xae, 0x23, 0xa2,
	0x77, 0x7a, 0xc2, 0xe0, 0x78, 0x74, 0xfa, 0x94, 0x4a, 0x30, 0xc7, 0x6b, 0x7b, 0x52, 0x14, 0x59,
	0xa1, 0x77, 0x99, 0x13, 0x4f, 0xd0, 0x7b, 0x80, 0xa0, 0x20, 0x1e, 0x25, 0xe1, 0xda, 0xa3, 0xba,
	0xc2, 0x7a, 0x18, 0x26, 0xdf, 0x96, 0xb9, 0xdf, 0x96, 0xe9, 0xee, 0xb7, 0x85, 0xd5, 0x86, 0x3d,
	0xa1, 0xb5, 0xb4, 0xca, 0xc3, 0xbd, 0x54, 0xfd, 0xbf, 0xb4, 0x61, 0x4f, 0x68, 0xff, 0x23, 0xc0,
	0xe1, 0x07, 0xd1, 0x5b, 0xe8, 0xe6, 0xb1, 0xb7, 0x23, 0x45, 0xa9, 0x0b, 0x3d, 0x69, 0x70, 0x3c,
	0xba, 0x78, 0x38, 0xc4, 0x35, 0x83, 0xf8, 0x28, 0x7b, 0x5e, 0xff, 0xa7, 0x08, 0xda, 0x53, 0xb4,
	0x5e, 0x65, 0xea, 0x25, 0x84, 0x5d, 0x43, 0xc5, 0x2c, 0xae, 0x6b, 0x85, 0x97, 0xde, 0xb2, 0x5b,
	0xc8, 0x98, 0xc5, 0x68, 0x0c, 0x72, 0xee, 0x15, 0xb4, 0xd4, 0x25, 0xe6, 0xd6, 0xff, 0x87, 0x9b,
	0x79, 0x5d, 0x93, 0xac, 0x94, 0x16, 0x3b, 0xcc, 0x05, 0xc8, 0x00, 0x25, 0x4a, 0x83, 0xb8, 0x0a,
	0x49, 0xc8, 0x0e, 0xa6, 0xe0, 0xfb, 0x1c, 0xbd, 0x04, 0x35, 0xc8, 0xe2, 0x98, 0x04, 0x94, 0x84,
	0xec, 0x6e, 0x0a, 0x3e, 0x14, 0x6a, 0xb4, 0xac, 0xfc, 0x24, 0xa2, 0x35, 0xda, 0xe1, 0xe8, 0x7d,
	0xc1, 0x18, 0x03, 0x1c, 0xcc, 0x90, 0x06, 0xd2, 0x2d, 0xd9, 0x35, 0x63, 0xd4, 0x61, 0x7d, 0xbb,
	0x1f, 0x5e, 0x5c, 0x91, 0x66, 0x0c, 0x9e, 0x7c, 0x10, 0xc7, 0xc2, 0xd5, 0x2f, 0x11, 0x8e, 0x1f,
	0x3c, 0x0b, 0xa4, 0xc3, 0x19, 0x5e, 0xae, 0xec, 0xd9, 0xda, 0x71, 0x27, 0xee, 0xca, 0x59, 0xaf,
	0xec, 0xaf, 0xf6, 0xf2, 0xbb, 0xad, 0x3d, 0x43, 0xe7, 0x70, 0xfa, 0x08, 0xf9, 0xb2, 0x9c, 0xdb,
	0x9a, 0x80, 0x2e, 0xe0, 0x79, 0xab, 0x6c, 0xcd, 0x34, 0xb1, 0xd5, 0x69, 0xba, 0x5c, 0x2c, 0xac,
	0xa9, 0xab, 0x49, 0xc8, 0x80, 0x17, 0x7f, 0x43, 0xac, 0x99, 0x76, 0xd4, 0x6a, 0xe7, 0xac, 0x3e,
	0x7d, 0x9b, 0xbb, 0x9a, 0xdc, 0x12, 0x71, 0xa0, 0x16, 0x75, 0x5a, 0x56, 0xce, 0x6a, 0x3a, 0xb5,
	0x1c, 0x47, 0xeb, 0xb6, 0xda, 0x7d, 0x9e, 0xcc, 0x17, 0xd6, 0x4c, 0x53, 0xd0, 0x25, 0x9c, 0x3f,
	0x96, 0x58, 0xb6, 0x3b, 0xb7, 0xad, 0x85, 0xa6, 0xfa, 0x1d, 0xf6, 0xf6, 0xde, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xda, 0x5a, 0xa8, 0x7e, 0x06, 0x04, 0x00, 0x00,
}
