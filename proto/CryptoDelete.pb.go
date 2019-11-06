// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoDelete.proto

package proto

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

// Mark an account as deleted, moving all its current hbars to another account. It will remain in the ledger, marked as deleted, until it expires. Transfers into it a deleted account fail. But a deleted account can still have its expiration extended in the normal way.
type CryptoDeleteTransactionBody struct {
	TransferAccountID    *AccountID `protobuf:"bytes,1,opt,name=transferAccountID,proto3" json:"transferAccountID,omitempty"`
	DeleteAccountID      *AccountID `protobuf:"bytes,2,opt,name=deleteAccountID,proto3" json:"deleteAccountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CryptoDeleteTransactionBody) Reset()         { *m = CryptoDeleteTransactionBody{} }
func (m *CryptoDeleteTransactionBody) String() string { return proto.CompactTextString(m) }
func (*CryptoDeleteTransactionBody) ProtoMessage()    {}
func (*CryptoDeleteTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_600103b910f0c423, []int{0}
}

func (m *CryptoDeleteTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoDeleteTransactionBody.Unmarshal(m, b)
}
func (m *CryptoDeleteTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoDeleteTransactionBody.Marshal(b, m, deterministic)
}
func (m *CryptoDeleteTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoDeleteTransactionBody.Merge(m, src)
}
func (m *CryptoDeleteTransactionBody) XXX_Size() int {
	return xxx_messageInfo_CryptoDeleteTransactionBody.Size(m)
}
func (m *CryptoDeleteTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoDeleteTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoDeleteTransactionBody proto.InternalMessageInfo

func (m *CryptoDeleteTransactionBody) GetTransferAccountID() *AccountID {
	if m != nil {
		return m.TransferAccountID
	}
	return nil
}

func (m *CryptoDeleteTransactionBody) GetDeleteAccountID() *AccountID {
	if m != nil {
		return m.DeleteAccountID
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoDeleteTransactionBody)(nil), "proto.CryptoDeleteTransactionBody")
}

func init() { proto.RegisterFile("CryptoDelete.proto", fileDescriptor_600103b910f0c423) }

var fileDescriptor_600103b910f0c423 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x72, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0x77, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x05, 0x53, 0x52, 0x02, 0x4e, 0x89, 0xc5, 0x99, 0xc9, 0x21, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x09,
	0xa5, 0x99, 0x8c, 0x5c, 0xd2, 0xc8, 0xea, 0x43, 0x8a, 0x12, 0xf3, 0x8a, 0x13, 0x93, 0x4b, 0x32,
	0xf3, 0xf3, 0x9c, 0xf2, 0x53, 0x2a, 0x85, 0xec, 0xb8, 0x04, 0x4b, 0x40, 0x42, 0x69, 0xa9, 0x45,
	0x8e, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0x02, 0x10, 0x23, 0xf4, 0xe0, 0xe2, 0x41, 0x98, 0x4a, 0x85, 0xac, 0xb8, 0xf8, 0x53, 0xc0, 0x06,
	0x23, 0x74, 0x33, 0xe1, 0xd0, 0x8d, 0xae, 0xd0, 0x49, 0x8d, 0x4b, 0x29, 0x39, 0x3f, 0x57, 0x2f,
	0x23, 0x35, 0x25, 0xb5, 0x28, 0x31, 0x23, 0xb1, 0x38, 0x23, 0xbd, 0x28, 0xb1, 0x20, 0x43, 0x2f,
	0xb1, 0x20, 0x13, 0xaa, 0x37, 0x2b, 0xb1, 0x2c, 0x31, 0x80, 0x31, 0x89, 0x0d, 0xcc, 0x33, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x53, 0x6e, 0x31, 0x95, 0xf9, 0x00, 0x00, 0x00,
}
