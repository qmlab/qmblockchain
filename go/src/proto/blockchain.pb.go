// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package pb

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

type Transaction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Recipient            string   `protobuf:"bytes,3,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Val                  float64  `protobuf:"fixed64,4,opt,name=Val,proto3" json:"Val,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_d4304b0d40850632, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Transaction) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Transaction) GetVal() float64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Block struct {
	Index                int32    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	PrevHash             string   `protobuf:"bytes,3,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Proof                int64    `protobuf:"varint,4,opt,name=Proof,proto3" json:"Proof,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Txs                  *Tree    `protobuf:"bytes,6,opt,name=Txs,proto3" json:"Txs,omitempty"`
	Balances             *Tree    `protobuf:"bytes,7,opt,name=Balances,proto3" json:"Balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_d4304b0d40850632, []int{1}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetProof() int64 {
	if m != nil {
		return m.Proof
	}
	return 0
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetTxs() *Tree {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *Block) GetBalances() *Tree {
	if m != nil {
		return m.Balances
	}
	return nil
}

type User struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_d4304b0d40850632, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Chain struct {
	Blocks               []*Block       `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	OpenTxs              []*Transaction `protobuf:"bytes,2,rep,name=OpenTxs,proto3" json:"OpenTxs,omitempty"`
	Difficulty           int32          `protobuf:"varint,3,opt,name=Difficulty,proto3" json:"Difficulty,omitempty"`
	Usr                  *User          `protobuf:"bytes,4,opt,name=Usr,proto3" json:"Usr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Chain) Reset()         { *m = Chain{} }
func (m *Chain) String() string { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()    {}
func (*Chain) Descriptor() ([]byte, []int) {
	return fileDescriptor_blockchain_d4304b0d40850632, []int{3}
}
func (m *Chain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chain.Unmarshal(m, b)
}
func (m *Chain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chain.Marshal(b, m, deterministic)
}
func (dst *Chain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chain.Merge(dst, src)
}
func (m *Chain) XXX_Size() int {
	return xxx_messageInfo_Chain.Size(m)
}
func (m *Chain) XXX_DiscardUnknown() {
	xxx_messageInfo_Chain.DiscardUnknown(m)
}

var xxx_messageInfo_Chain proto.InternalMessageInfo

func (m *Chain) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Chain) GetOpenTxs() []*Transaction {
	if m != nil {
		return m.OpenTxs
	}
	return nil
}

func (m *Chain) GetDifficulty() int32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Chain) GetUsr() *User {
	if m != nil {
		return m.Usr
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Chain)(nil), "pb.Chain")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_blockchain_d4304b0d40850632) }

var fileDescriptor_blockchain_d4304b0d40850632 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x41, 0xc9, 0x92, 0xed, 0x33, 0xe0, 0x1a, 0x44, 0x51, 0x10, 0x42, 0x51, 0xa8, 0x42,
	0x07, 0x75, 0xf1, 0xe0, 0x3e, 0x41, 0xdd, 0x0e, 0xf5, 0x54, 0x83, 0x91, 0xb3, 0x53, 0x12, 0x0d,
	0x13, 0x91, 0x29, 0x81, 0xa4, 0x03, 0xe7, 0x31, 0x32, 0xe7, 0x71, 0xf2, 0x62, 0x01, 0x4f, 0x8a,
	0x1d, 0x64, 0xc9, 0x76, 0xff, 0x7f, 0x07, 0xde, 0x77, 0x77, 0x84, 0x45, 0xd9, 0xb4, 0xd5, 0x5d,
	0x75, 0x10, 0x4a, 0x2f, 0x3b, 0xd3, 0xba, 0x96, 0x06, 0x5d, 0x99, 0xcc, 0x3b, 0xe1, 0x8c, 0xaa,
	0x94, 0xe8, 0xbd, 0xec, 0x89, 0xc0, 0xac, 0x30, 0x42, 0x5b, 0x51, 0x39, 0xd5, 0x6a, 0x3a, 0x87,
	0x60, 0x53, 0x33, 0x92, 0x92, 0x7c, 0xca, 0x83, 0x4d, 0x4d, 0xbf, 0x40, 0x7c, 0x23, 0x75, 0x2d,
	0x0d, 0x0b, 0xd0, 0x1b, 0x14, 0xfd, 0x0a, 0x53, 0x2e, 0x2b, 0xd5, 0x29, 0xa9, 0x1d, 0x0b, 0x31,
	0x75, 0x35, 0xe8, 0x02, 0xc2, 0x5b, 0xd1, 0xb0, 0x51, 0x4a, 0x72, 0xc2, 0x7d, 0xe8, 0xeb, 0x0b,
	0x75, 0x94, 0xd6, 0x89, 0x63, 0xc7, 0xa2, 0x94, 0xe4, 0x21, 0xbf, 0x1a, 0xd8, 0xc5, 0x09, 0x77,
	0xb2, 0x2c, 0x1e, 0xba, 0xa0, 0xca, 0x9e, 0x09, 0x44, 0x6b, 0x3f, 0x06, 0xfd, 0x0c, 0xd1, 0x46,
	0xd7, 0xf2, 0x8c, 0x68, 0x11, 0xef, 0x05, 0xa5, 0x30, 0xfa, 0x27, 0xec, 0x61, 0x60, 0xc3, 0x98,
	0x26, 0x30, 0xd9, 0x1a, 0x79, 0x8f, 0x7e, 0x0f, 0x76, 0xd1, 0xfe, 0x95, 0xad, 0x69, 0xdb, 0x3d,
	0x92, 0x85, 0xbc, 0x17, 0x1f, 0xb0, 0x25, 0x10, 0x16, 0xe7, 0x1e, 0x6c, 0xb6, 0x9a, 0x2c, 0xbb,
	0x72, 0x59, 0x18, 0x29, 0xb9, 0x37, 0xe9, 0x0f, 0x98, 0xac, 0x45, 0x23, 0x74, 0x25, 0x2d, 0x1b,
	0xbf, 0x2b, 0xb8, 0x64, 0xb2, 0x04, 0x46, 0x3b, 0x2b, 0x8d, 0xa7, 0xfd, 0x5d, 0xd7, 0x66, 0xd8,
	0x2e, 0xc6, 0xd9, 0x23, 0x81, 0xe8, 0x8f, 0xbf, 0x11, 0xfd, 0x0e, 0x31, 0x8e, 0x6a, 0x19, 0x49,
	0xc3, 0x7c, 0xb6, 0x9a, 0xfa, 0x97, 0xd0, 0xe1, 0x43, 0x82, 0xfe, 0x84, 0xf1, 0xff, 0x4e, 0x6a,
	0x8f, 0x13, 0x60, 0xcd, 0xa7, 0xbe, 0xdb, 0xe5, 0x7c, 0xfc, 0x35, 0x4f, 0xbf, 0x01, 0xfc, 0x55,
	0xfb, 0xbd, 0xaa, 0x4e, 0x8d, 0x7b, 0xc0, 0x3d, 0x44, 0xfc, 0x8d, 0xe3, 0xa7, 0xda, 0x59, 0x83,
	0x7b, 0x18, 0xa0, 0x3d, 0x22, 0xf7, 0x66, 0x19, 0xe3, 0xd7, 0xf8, 0xf5, 0x12, 0x00, 0x00, 0xff,
	0xff, 0xbc, 0xd0, 0x2f, 0x1a, 0x42, 0x02, 0x00, 0x00,
}
