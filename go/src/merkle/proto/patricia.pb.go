// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patricia.proto

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

type Node struct {
	Hash                 string            `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Next                 []string          `protobuf:"bytes,2,rep,name=Next,proto3" json:"Next,omitempty"`
	Val                  string            `protobuf:"bytes,3,opt,name=Val,proto3" json:"Val,omitempty"`
	Count                int32             `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
	EncodedPaths         map[string]string `protobuf:"bytes,5,rep,name=EncodedPaths,proto3" json:"EncodedPaths,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_patricia_76c2e0332967e273, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Node) GetNext() []string {
	if m != nil {
		return m.Next
	}
	return nil
}

func (m *Node) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func (m *Node) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Node) GetEncodedPaths() map[string]string {
	if m != nil {
		return m.EncodedPaths
	}
	return nil
}

type Tree struct {
	Root                 *Node            `protobuf:"bytes,1,opt,name=Root,proto3" json:"Root,omitempty"`
	Ht                   map[string]*Node `protobuf:"bytes,2,rep,name=Ht,proto3" json:"Ht,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastRadixCompression int64            `protobuf:"varint,3,opt,name=LastRadixCompression,proto3" json:"LastRadixCompression,omitempty"`
	BatchSize            int64            `protobuf:"varint,4,opt,name=BatchSize,proto3" json:"BatchSize,omitempty"`
	Zipped               bool             `protobuf:"varint,5,opt,name=Zipped,proto3" json:"Zipped,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_patricia_76c2e0332967e273, []int{1}
}
func (m *Tree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tree.Unmarshal(m, b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
}
func (dst *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(dst, src)
}
func (m *Tree) XXX_Size() int {
	return xxx_messageInfo_Tree.Size(m)
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetRoot() *Node {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Tree) GetHt() map[string]*Node {
	if m != nil {
		return m.Ht
	}
	return nil
}

func (m *Tree) GetLastRadixCompression() int64 {
	if m != nil {
		return m.LastRadixCompression
	}
	return 0
}

func (m *Tree) GetBatchSize() int64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *Tree) GetZipped() bool {
	if m != nil {
		return m.Zipped
	}
	return false
}

func init() {
	proto.RegisterType((*Node)(nil), "pb.Node")
	proto.RegisterMapType((map[string]string)(nil), "pb.Node.EncodedPathsEntry")
	proto.RegisterType((*Tree)(nil), "pb.Tree")
	proto.RegisterMapType((map[string]*Node)(nil), "pb.Tree.HtEntry")
}

func init() { proto.RegisterFile("patricia.proto", fileDescriptor_patricia_76c2e0332967e273) }

var fileDescriptor_patricia_76c2e0332967e273 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0x33, 0x31,
	0x14, 0x24, 0xc9, 0x6e, 0xbf, 0xf6, 0xf5, 0x43, 0x6a, 0x28, 0x12, 0x4a, 0x91, 0xd0, 0xd3, 0x9e,
	0xf6, 0xb0, 0x5e, 0xc4, 0x83, 0x05, 0x4b, 0x61, 0x0f, 0x52, 0x24, 0x8a, 0x07, 0x6f, 0x69, 0x37,
	0xd0, 0x60, 0xdd, 0x84, 0xdd, 0x54, 0x5a, 0xff, 0xa8, 0xbf, 0xc3, 0x7f, 0x20, 0xc9, 0x56, 0xba,
	0x62, 0x6f, 0xf3, 0xe6, 0x4d, 0x98, 0x99, 0x17, 0x38, 0xb3, 0xd2, 0x55, 0x7a, 0xa5, 0x65, 0x6a,
	0x2b, 0xe3, 0x0c, 0xc5, 0x76, 0x39, 0xf9, 0x44, 0x10, 0x2d, 0x4c, 0xa1, 0x28, 0x85, 0x28, 0x97,
	0xf5, 0x9a, 0x21, 0x8e, 0x92, 0x9e, 0x08, 0xd8, 0x73, 0x0b, 0xb5, 0x73, 0x0c, 0x73, 0xe2, 0x39,
	0x8f, 0xe9, 0x00, 0xc8, 0xb3, 0xdc, 0x30, 0x12, 0x64, 0x1e, 0xd2, 0x21, 0xc4, 0x33, 0xb3, 0x2d,
	0x1d, 0x8b, 0x38, 0x4a, 0x62, 0xd1, 0x0c, 0xf4, 0x16, 0xfe, 0xcf, 0xcb, 0x95, 0x29, 0x54, 0xf1,
	0x20, 0xdd, 0xba, 0x66, 0x31, 0x27, 0x49, 0x3f, 0x1b, 0xa5, 0x76, 0x99, 0x7a, 0xbf, 0xb4, 0xbd,
	0x9c, 0x97, 0xae, 0xda, 0x8b, 0x5f, 0xfa, 0xd1, 0x14, 0xce, 0xff, 0x48, 0xbc, 0xf9, 0xab, 0xda,
	0x1f, 0x32, 0x7a, 0xe8, 0xcd, 0xdf, 0xe5, 0x66, 0xab, 0x18, 0x0e, 0x5c, 0x33, 0xdc, 0xe0, 0x6b,
	0x34, 0xf9, 0x42, 0x10, 0x3d, 0x55, 0x4a, 0xd1, 0x31, 0x44, 0xc2, 0x18, 0x17, 0x5e, 0xf5, 0xb3,
	0xee, 0x4f, 0x02, 0x11, 0x58, 0xca, 0x01, 0xe7, 0x4d, 0xc3, 0x7e, 0x36, 0xf0, 0x3b, 0xff, 0x26,
	0xcd, 0x5d, 0x93, 0x09, 0xe7, 0x8e, 0x66, 0x30, 0xbc, 0x97, 0xb5, 0x13, 0xb2, 0xd0, 0xbb, 0x99,
	0x79, 0xb3, 0x95, 0xaa, 0x6b, 0x6d, 0xca, 0x70, 0x02, 0x22, 0x4e, 0xee, 0xe8, 0x18, 0x7a, 0x77,
	0xd2, 0xad, 0xd6, 0x8f, 0xfa, 0x43, 0x85, 0xbb, 0x10, 0x71, 0x24, 0xe8, 0x05, 0x74, 0x5e, 0xb4,
	0xb5, 0xaa, 0x60, 0x31, 0x47, 0x49, 0x57, 0x1c, 0xa6, 0xd1, 0x14, 0xfe, 0x1d, 0x8c, 0x4f, 0x34,
	0xbd, 0x6c, 0x37, 0x6d, 0xf7, 0x38, 0x76, 0x5e, 0x76, 0xc2, 0xc7, 0x5e, 0x7d, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xec, 0x9a, 0x0e, 0x2b, 0xea, 0x01, 0x00, 0x00,
}
