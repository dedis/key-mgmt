// Code generated by protoc-gen-go.
// source: c2s.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	c2s.proto
	util.proto

It has these top-level messages:
	Registration
	CommitmentReq
	KeyLookup
	RegistrationResp
	AuthPath
	Hash
	Commitment
	ServerResp
	CompleteRootNode
	WitnessedCommitment
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// ok to use same format for getComm and getWitness calls
type CommitmentReq_CommitmentType int32

const (
	CommitmentReq_SELF    CommitmentReq_CommitmentType = 0
	CommitmentReq_WITNESS CommitmentReq_CommitmentType = 1
)

var CommitmentReq_CommitmentType_name = map[int32]string{
	0: "SELF",
	1: "WITNESS",
}
var CommitmentReq_CommitmentType_value = map[string]int32{
	"SELF":    0,
	"WITNESS": 1,
}

func (x CommitmentReq_CommitmentType) Enum() *CommitmentReq_CommitmentType {
	p := new(CommitmentReq_CommitmentType)
	*p = x
	return p
}
func (x CommitmentReq_CommitmentType) String() string {
	return proto.EnumName(CommitmentReq_CommitmentType_name, int32(x))
}
func (x *CommitmentReq_CommitmentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommitmentReq_CommitmentType_value, data, "CommitmentReq_CommitmentType")
	if err != nil {
		return err
	}
	*x = CommitmentReq_CommitmentType(value)
	return nil
}
func (CommitmentReq_CommitmentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

// indicates if the hash is for the left or right subtree
type AuthPath_PrunedChild int32

const (
	AuthPath_LEFT  AuthPath_PrunedChild = 0
	AuthPath_RIGHT AuthPath_PrunedChild = 1
)

var AuthPath_PrunedChild_name = map[int32]string{
	0: "LEFT",
	1: "RIGHT",
}
var AuthPath_PrunedChild_value = map[string]int32{
	"LEFT":  0,
	"RIGHT": 1,
}

func (x AuthPath_PrunedChild) Enum() *AuthPath_PrunedChild {
	p := new(AuthPath_PrunedChild)
	*p = x
	return p
}
func (x AuthPath_PrunedChild) String() string {
	return proto.EnumName(AuthPath_PrunedChild_name, int32(x))
}
func (x *AuthPath_PrunedChild) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthPath_PrunedChild_value, data, "AuthPath_PrunedChild")
	if err != nil {
		return err
	}
	*x = AuthPath_PrunedChild(value)
	return nil
}
func (AuthPath_PrunedChild) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Registration struct {
	// server must make sure that these two fields are specified
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Publickey        *string `protobuf:"bytes,2,opt,name=publickey" json:"publickey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Registration) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Registration) GetPublickey() string {
	if m != nil && m.Publickey != nil {
		return *m.Publickey
	}
	return ""
}

type CommitmentReq struct {
	// server checks that type, and epoch are specified
	Type *CommitmentReq_CommitmentType `protobuf:"varint,1,opt,name=type,enum=common.CommitmentReq_CommitmentType" json:"type,omitempty"`
	// epoch  should <= than current epoch
	Epoch *uint64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// provider MUST be specified if type is WITNESS
	Provider         *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CommitmentReq) Reset()                    { *m = CommitmentReq{} }
func (m *CommitmentReq) String() string            { return proto.CompactTextString(m) }
func (*CommitmentReq) ProtoMessage()               {}
func (*CommitmentReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommitmentReq) GetType() CommitmentReq_CommitmentType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return CommitmentReq_SELF
}

func (m *CommitmentReq) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

func (m *CommitmentReq) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

type KeyLookup struct {
	// server checks that name and epoch are specified
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Epoch            *uint64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyLookup) Reset()                    { *m = KeyLookup{} }
func (m *KeyLookup) String() string            { return proto.CompactTextString(m) }
func (*KeyLookup) ProtoMessage()               {}
func (*KeyLookup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyLookup) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *KeyLookup) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

type RegistrationResp struct {
	// client checks that initial epoch and epoch interval are specified
	InitEpoch        *uint64 `protobuf:"varint,1,opt,name=init_epoch" json:"init_epoch,omitempty"`
	EpochInterval    *uint32 `protobuf:"varint,2,opt,name=epoch_interval" json:"epoch_interval,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegistrationResp) Reset()                    { *m = RegistrationResp{} }
func (m *RegistrationResp) String() string            { return proto.CompactTextString(m) }
func (*RegistrationResp) ProtoMessage()               {}
func (*RegistrationResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegistrationResp) GetInitEpoch() uint64 {
	if m != nil && m.InitEpoch != nil {
		return *m.InitEpoch
	}
	return 0
}

func (m *RegistrationResp) GetEpochInterval() uint32 {
	if m != nil && m.EpochInterval != nil {
		return *m.EpochInterval
	}
	return 0
}

type AuthPath struct {
	Leaf             *AuthPath_UserLeafNode   `protobuf:"bytes,1,opt,name=leaf" json:"leaf,omitempty"`
	Interior         []*AuthPath_InteriorNode `protobuf:"bytes,2,rep,name=interior" json:"interior,omitempty"`
	Root             *AuthPath_RootNode       `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *AuthPath) Reset()                    { *m = AuthPath{} }
func (m *AuthPath) String() string            { return proto.CompactTextString(m) }
func (*AuthPath) ProtoMessage()               {}
func (*AuthPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthPath) GetLeaf() *AuthPath_UserLeafNode {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *AuthPath) GetInterior() []*AuthPath_InteriorNode {
	if m != nil {
		return m.Interior
	}
	return nil
}

func (m *AuthPath) GetRoot() *AuthPath_RootNode {
	if m != nil {
		return m.Root
	}
	return nil
}

// auth path consists of user leaf node, possibly interior nodes, and root node
type AuthPath_UserLeafNode struct {
	// client does not assume server has not sent malformed leaf node
	Name                    *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Publickey               *string  `protobuf:"bytes,2,opt,name=publickey" json:"publickey,omitempty"`
	EpochAdded              *uint64  `protobuf:"varint,3,opt,name=epoch_added" json:"epoch_added,omitempty"`
	AllowsUnsignedKeychange *bool    `protobuf:"varint,4,opt,name=allows_unsigned_keychange" json:"allows_unsigned_keychange,omitempty"`
	AllowsPublicLookup      *bool    `protobuf:"varint,5,opt,name=allows_public_lookup" json:"allows_public_lookup,omitempty"`
	LookupIndex             []uint32 `protobuf:"fixed32,6,rep,name=lookup_index" json:"lookup_index,omitempty"`
	// repeated fixed32 signature = 7;
	Intlevels        *uint32 `protobuf:"varint,7,opt,name=intlevels" json:"intlevels,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthPath_UserLeafNode) Reset()                    { *m = AuthPath_UserLeafNode{} }
func (m *AuthPath_UserLeafNode) String() string            { return proto.CompactTextString(m) }
func (*AuthPath_UserLeafNode) ProtoMessage()               {}
func (*AuthPath_UserLeafNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *AuthPath_UserLeafNode) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AuthPath_UserLeafNode) GetPublickey() string {
	if m != nil && m.Publickey != nil {
		return *m.Publickey
	}
	return ""
}

func (m *AuthPath_UserLeafNode) GetEpochAdded() uint64 {
	if m != nil && m.EpochAdded != nil {
		return *m.EpochAdded
	}
	return 0
}

func (m *AuthPath_UserLeafNode) GetAllowsUnsignedKeychange() bool {
	if m != nil && m.AllowsUnsignedKeychange != nil {
		return *m.AllowsUnsignedKeychange
	}
	return false
}

func (m *AuthPath_UserLeafNode) GetAllowsPublicLookup() bool {
	if m != nil && m.AllowsPublicLookup != nil {
		return *m.AllowsPublicLookup
	}
	return false
}

func (m *AuthPath_UserLeafNode) GetLookupIndex() []uint32 {
	if m != nil {
		return m.LookupIndex
	}
	return nil
}

func (m *AuthPath_UserLeafNode) GetIntlevels() uint32 {
	if m != nil && m.Intlevels != nil {
		return *m.Intlevels
	}
	return 0
}

type AuthPath_InteriorNode struct {
	// client needs to check that both of these fields are set for each interior node
	Prunedchild      *AuthPath_PrunedChild `protobuf:"varint,1,opt,name=prunedchild,enum=common.AuthPath_PrunedChild" json:"prunedchild,omitempty"`
	Subtree          *Hash                 `protobuf:"bytes,2,opt,name=subtree" json:"subtree,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *AuthPath_InteriorNode) Reset()                    { *m = AuthPath_InteriorNode{} }
func (m *AuthPath_InteriorNode) String() string            { return proto.CompactTextString(m) }
func (*AuthPath_InteriorNode) ProtoMessage()               {}
func (*AuthPath_InteriorNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *AuthPath_InteriorNode) GetPrunedchild() AuthPath_PrunedChild {
	if m != nil && m.Prunedchild != nil {
		return *m.Prunedchild
	}
	return AuthPath_LEFT
}

func (m *AuthPath_InteriorNode) GetSubtree() *Hash {
	if m != nil {
		return m.Subtree
	}
	return nil
}

type AuthPath_RootNode struct {
	// client does not assume server has not sent malformed root node
	Prunedchild      *AuthPath_PrunedChild `protobuf:"varint,1,opt,name=prunedchild,enum=common.AuthPath_PrunedChild" json:"prunedchild,omitempty"`
	Subtree          *Hash                 `protobuf:"bytes,2,opt,name=subtree" json:"subtree,omitempty"`
	Prev             *Hash                 `protobuf:"bytes,3,opt,name=prev" json:"prev,omitempty"`
	Epoch            *uint64               `protobuf:"varint,4,opt,name=epoch" json:"epoch,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *AuthPath_RootNode) Reset()                    { *m = AuthPath_RootNode{} }
func (m *AuthPath_RootNode) String() string            { return proto.CompactTextString(m) }
func (*AuthPath_RootNode) ProtoMessage()               {}
func (*AuthPath_RootNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 2} }

func (m *AuthPath_RootNode) GetPrunedchild() AuthPath_PrunedChild {
	if m != nil && m.Prunedchild != nil {
		return *m.Prunedchild
	}
	return AuthPath_LEFT
}

func (m *AuthPath_RootNode) GetSubtree() *Hash {
	if m != nil {
		return m.Subtree
	}
	return nil
}

func (m *AuthPath_RootNode) GetPrev() *Hash {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *AuthPath_RootNode) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*Registration)(nil), "common.Registration")
	proto.RegisterType((*CommitmentReq)(nil), "common.CommitmentReq")
	proto.RegisterType((*KeyLookup)(nil), "common.KeyLookup")
	proto.RegisterType((*RegistrationResp)(nil), "common.RegistrationResp")
	proto.RegisterType((*AuthPath)(nil), "common.AuthPath")
	proto.RegisterType((*AuthPath_UserLeafNode)(nil), "common.AuthPath.UserLeafNode")
	proto.RegisterType((*AuthPath_InteriorNode)(nil), "common.AuthPath.InteriorNode")
	proto.RegisterType((*AuthPath_RootNode)(nil), "common.AuthPath.RootNode")
	proto.RegisterEnum("common.CommitmentReq_CommitmentType", CommitmentReq_CommitmentType_name, CommitmentReq_CommitmentType_value)
	proto.RegisterEnum("common.AuthPath_PrunedChild", AuthPath_PrunedChild_name, AuthPath_PrunedChild_value)
}

var fileDescriptor0 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0x5b, 0xb6, 0x26, 0x6f, 0xba, 0xa9, 0xb2, 0x26, 0x91, 0x55, 0x9b, 0x18, 0x11, 0x5f,
	0x17, 0x2a, 0xa8, 0x84, 0x04, 0x1c, 0x10, 0xd3, 0xd4, 0xb1, 0x6a, 0xd5, 0x34, 0x79, 0x9d, 0xe0,
	0x56, 0x65, 0x89, 0x69, 0xac, 0xa5, 0x71, 0x48, 0x9c, 0x42, 0xcf, 0xfc, 0x01, 0x7e, 0x04, 0x27,
	0x7e, 0x25, 0xce, 0xeb, 0x64, 0x4d, 0xb5, 0x5e, 0xb8, 0x70, 0x8b, 0x9f, 0x0f, 0x3f, 0xce, 0xf3,
	0xda, 0x60, 0x07, 0xfd, 0xbc, 0x97, 0x66, 0x42, 0x0a, 0xb2, 0x1d, 0x88, 0xd9, 0x4c, 0x24, 0x5d,
	0x28, 0x24, 0x8f, 0x35, 0xe6, 0x7d, 0x84, 0x36, 0x65, 0x53, 0x9e, 0xcb, 0xcc, 0x97, 0x5c, 0x24,
	0x84, 0x80, 0x99, 0xf8, 0x33, 0xe6, 0x1a, 0x47, 0xc6, 0x0b, 0x9b, 0xe2, 0x37, 0x39, 0x00, 0x3b,
	0x2d, 0x6e, 0x62, 0x1e, 0xdc, 0xb2, 0x85, 0xbb, 0x81, 0xc4, 0x12, 0xf0, 0x7e, 0x1b, 0xb0, 0x73,
	0xa2, 0x36, 0xe6, 0x72, 0xc6, 0x12, 0x49, 0xd9, 0x37, 0xf2, 0x16, 0x4c, 0xb9, 0x48, 0xf5, 0x1e,
	0xbb, 0xfd, 0x27, 0x3d, 0x1d, 0xdb, 0x5b, 0x11, 0x35, 0x56, 0x63, 0xa5, 0xa5, 0xe8, 0x20, 0x7b,
	0xb0, 0xc5, 0x52, 0x11, 0x44, 0x98, 0x62, 0x52, 0xbd, 0x20, 0x5d, 0xb0, 0xd4, 0x61, 0xe7, 0x3c,
	0x64, 0x99, 0xbb, 0x89, 0xf1, 0x77, 0x6b, 0xef, 0x39, 0xec, 0xae, 0xee, 0x44, 0x2c, 0x30, 0xaf,
	0x06, 0xa3, 0xd3, 0xce, 0x03, 0xe2, 0x40, 0xeb, 0xf3, 0x70, 0x7c, 0x31, 0xb8, 0xba, 0xea, 0x18,
	0xde, 0x1b, 0xb0, 0xcf, 0xd9, 0x62, 0x24, 0xc4, 0x6d, 0x91, 0xae, 0xfd, 0xcb, 0xb5, 0xd9, 0xde,
	0x17, 0xe8, 0x34, 0xfb, 0xa1, 0x2c, 0x4f, 0xc9, 0x21, 0x00, 0x4f, 0xb8, 0x9c, 0x68, 0xb9, 0x81,
	0x72, 0xbb, 0x44, 0x06, 0x78, 0xdc, 0xa7, 0xb0, 0x8b, 0xcc, 0x84, 0x27, 0x92, 0x65, 0x73, 0x3f,
	0xc6, 0x1d, 0x77, 0xe8, 0x0e, 0xa2, 0xc3, 0x0a, 0xf4, 0x7e, 0x6e, 0x83, 0x75, 0x5c, 0xc8, 0xe8,
	0xd2, 0x97, 0x11, 0x79, 0x0d, 0x66, 0xcc, 0xfc, 0xaf, 0xb8, 0x99, 0xd3, 0x3f, 0xac, 0x2b, 0xab,
	0xf9, 0xde, 0x75, 0xce, 0xb2, 0x91, 0x12, 0x5c, 0x88, 0x50, 0x75, 0x55, 0x4a, 0xc9, 0x3b, 0xb0,
	0x30, 0x80, 0x8b, 0x4c, 0x05, 0x6c, 0xae, 0xb5, 0x0d, 0x2b, 0x01, 0xda, 0xee, 0xe4, 0xe4, 0x25,
	0x98, 0x99, 0x10, 0x12, 0xcb, 0x74, 0xfa, 0xfb, 0xf7, 0x6c, 0x54, 0x91, 0x3a, 0xa9, 0x94, 0x75,
	0x7f, 0x6d, 0x40, 0xbb, 0x79, 0x80, 0x7f, 0xbf, 0x24, 0xe4, 0x11, 0x38, 0xba, 0x13, 0x3f, 0x0c,
	0x59, 0x88, 0xc1, 0x26, 0x05, 0x84, 0x8e, 0x4b, 0x84, 0xbc, 0x87, 0x7d, 0x3f, 0x8e, 0xc5, 0xf7,
	0x7c, 0x52, 0x24, 0x39, 0x9f, 0x26, 0x2c, 0x9c, 0x28, 0x5f, 0x10, 0xf9, 0xc9, 0x94, 0xb9, 0xa6,
	0x92, 0x5b, 0xf4, 0xa1, 0x16, 0x5c, 0x57, 0xfc, 0x79, 0x4d, 0x93, 0x57, 0xb0, 0x57, 0x79, 0x75,
	0xe0, 0x24, 0xc6, 0x29, 0xbb, 0x5b, 0x68, 0x23, 0x9a, 0xbb, 0x44, 0xaa, 0x9a, 0xff, 0x63, 0x68,
	0x6b, 0x8d, 0x9a, 0x51, 0xc8, 0x7e, 0xb8, 0xdb, 0xaa, 0xbf, 0x16, 0x75, 0x34, 0x36, 0x2c, 0xa1,
	0xf2, 0x7f, 0x54, 0x5f, 0x31, 0x9b, 0xb3, 0x38, 0x77, 0x5b, 0x38, 0xc0, 0x25, 0xd0, 0x9d, 0x43,
	0xbb, 0xd9, 0x2d, 0xf9, 0x00, 0x4e, 0x9a, 0x15, 0xea, 0x54, 0x41, 0xc4, 0xe3, 0xb0, 0xba, 0xf9,
	0x07, 0xf7, 0x8a, 0xbd, 0x44, 0xcd, 0x49, 0xa9, 0xa1, 0x4d, 0x03, 0x79, 0x06, 0xad, 0xbc, 0xb8,
	0x91, 0x19, 0x63, 0xd8, 0x9d, 0xd3, 0x6f, 0xd7, 0xde, 0x33, 0x3f, 0x8f, 0x68, 0x4d, 0x76, 0xff,
	0x18, 0x60, 0xd5, 0xd3, 0xf9, 0x5f, 0xa1, 0xe4, 0x08, 0xcc, 0x34, 0x63, 0xf3, 0xea, 0xba, 0xac,
	0x8a, 0x90, 0x59, 0xbe, 0x1d, 0xb3, 0xf9, 0x76, 0x3c, 0x70, 0x1a, 0xd9, 0xe5, 0xc3, 0x1c, 0x0d,
	0x4e, 0xc7, 0xea, 0x61, 0xda, 0xb0, 0x45, 0x87, 0x9f, 0xce, 0xc6, 0x1d, 0xe3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd2, 0x7b, 0x34, 0x85, 0x9f, 0x04, 0x00, 0x00,
}