// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator_public_keys.proto

package gowrapper

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Protobuf definition for the Rust struct ValidatorPublicKeys
type ValidatorPublicKeys struct {
	// Validator account address
	AccountAddress []byte `protobuf:"bytes,1,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	// Consensus public key
	ConsensusPublicKey []byte `protobuf:"bytes,2,opt,name=consensus_public_key,json=consensusPublicKey,proto3" json:"consensus_public_key,omitempty"`
	// Network signing publick key
	NetworkSigningPublicKey []byte `protobuf:"bytes,3,opt,name=network_signing_public_key,json=networkSigningPublicKey,proto3" json:"network_signing_public_key,omitempty"`
	/// Network identity publick key
	NetworkIdentityPublicKey []byte   `protobuf:"bytes,4,opt,name=network_identity_public_key,json=networkIdentityPublicKey,proto3" json:"network_identity_public_key,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ValidatorPublicKeys) Reset()         { *m = ValidatorPublicKeys{} }
func (m *ValidatorPublicKeys) String() string { return proto.CompactTextString(m) }
func (*ValidatorPublicKeys) ProtoMessage()    {}
func (*ValidatorPublicKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e933cd90ba9d127, []int{0}
}

func (m *ValidatorPublicKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorPublicKeys.Unmarshal(m, b)
}
func (m *ValidatorPublicKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorPublicKeys.Marshal(b, m, deterministic)
}
func (m *ValidatorPublicKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorPublicKeys.Merge(m, src)
}
func (m *ValidatorPublicKeys) XXX_Size() int {
	return xxx_messageInfo_ValidatorPublicKeys.Size(m)
}
func (m *ValidatorPublicKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorPublicKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorPublicKeys proto.InternalMessageInfo

func (m *ValidatorPublicKeys) GetAccountAddress() []byte {
	if m != nil {
		return m.AccountAddress
	}
	return nil
}

func (m *ValidatorPublicKeys) GetConsensusPublicKey() []byte {
	if m != nil {
		return m.ConsensusPublicKey
	}
	return nil
}

func (m *ValidatorPublicKeys) GetNetworkSigningPublicKey() []byte {
	if m != nil {
		return m.NetworkSigningPublicKey
	}
	return nil
}

func (m *ValidatorPublicKeys) GetNetworkIdentityPublicKey() []byte {
	if m != nil {
		return m.NetworkIdentityPublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorPublicKeys)(nil), "types.ValidatorPublicKeys")
}

func init() { proto.RegisterFile("validator_public_keys.proto", fileDescriptor_5e933cd90ba9d127) }

var fileDescriptor_5e933cd90ba9d127 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x71, 0xea, 0xbf, 0x43, 0x10, 0x85, 0x28, 0x58, 0xdc, 0x8b, 0x78, 0xd1, 0x83, 0xb4,
	0x82, 0x47, 0xf1, 0xa0, 0x37, 0xf1, 0x22, 0x0a, 0x1e, 0xbc, 0x94, 0x34, 0x19, 0xea, 0xb0, 0xdd,
	0x4c, 0xc8, 0x4c, 0x77, 0xe9, 0x3b, 0xfb, 0x10, 0x42, 0x36, 0x5b, 0x7b, 0xfd, 0xe5, 0xfb, 0x49,
	0x20, 0x6a, 0xb1, 0x36, 0x3d, 0x3a, 0x23, 0x14, 0x9b, 0x30, 0xb4, 0x3d, 0xda, 0x66, 0x09, 0x23,
	0x57, 0x21, 0x92, 0x90, 0x3e, 0x94, 0x31, 0x00, 0x5f, 0xff, 0x16, 0xea, 0xec, 0x6b, 0x97, 0xbd,
	0xa7, 0xea, 0x0d, 0x46, 0xd6, 0x37, 0xea, 0xd4, 0x58, 0x4b, 0x83, 0x97, 0xc6, 0x38, 0x17, 0x81,
	0xb9, 0x2c, 0xae, 0x8a, 0xdb, 0xe3, 0x8f, 0x93, 0x3c, 0x3f, 0x6f, 0x57, 0x7d, 0xaf, 0xce, 0x2d,
	0x79, 0x06, 0xcf, 0x03, 0xcf, 0x9e, 0x29, 0xf7, 0x52, 0xad, 0xa7, 0xb3, 0xe9, 0x6e, 0xfd, 0xa8,
	0x2e, 0x3d, 0xc8, 0x86, 0xe2, 0xb2, 0x61, 0xec, 0x3c, 0xfa, 0x6e, 0xee, 0xf6, 0x93, 0xbb, 0xc8,
	0xc5, 0xe7, 0x36, 0xf8, 0xc7, 0x4f, 0x6a, 0xb1, 0xc3, 0xe8, 0xc0, 0x0b, 0xca, 0x38, 0xd7, 0x07,
	0x49, 0x97, 0x39, 0x79, 0xcd, 0xc5, 0xc4, 0x5f, 0xaa, 0xef, 0xbb, 0x0e, 0xe5, 0x67, 0x68, 0x2b,
	0x4b, 0xab, 0xda, 0x92, 0x83, 0x95, 0x59, 0x43, 0x44, 0x5b, 0xf7, 0xd8, 0x46, 0x63, 0x7b, 0x04,
	0x2f, 0x75, 0x47, 0x9b, 0x68, 0x42, 0x80, 0xd8, 0x1e, 0xa5, 0xcf, 0x7a, 0xf8, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xde, 0x1c, 0x77, 0x23, 0x4b, 0x01, 0x00, 0x00,
}
