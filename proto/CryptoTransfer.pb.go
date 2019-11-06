// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoTransfer.proto

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

// An account, and the amount that it sends or receives during a cryptocurrency transfer.
type AccountAmount struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Amount               int64      `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountAmount) Reset()         { *m = AccountAmount{} }
func (m *AccountAmount) String() string { return proto.CompactTextString(m) }
func (*AccountAmount) ProtoMessage()    {}
func (*AccountAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_1807c2e11ee8417d, []int{0}
}

func (m *AccountAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAmount.Unmarshal(m, b)
}
func (m *AccountAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAmount.Marshal(b, m, deterministic)
}
func (m *AccountAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAmount.Merge(m, src)
}
func (m *AccountAmount) XXX_Size() int {
	return xxx_messageInfo_AccountAmount.Size(m)
}
func (m *AccountAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAmount.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAmount proto.InternalMessageInfo

func (m *AccountAmount) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *AccountAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// A list of accounts and amounts to transfer out of each account (negative) or into it (positive).
type TransferList struct {
	AccountAmounts       []*AccountAmount `protobuf:"bytes,1,rep,name=accountAmounts,proto3" json:"accountAmounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransferList) Reset()         { *m = TransferList{} }
func (m *TransferList) String() string { return proto.CompactTextString(m) }
func (*TransferList) ProtoMessage()    {}
func (*TransferList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1807c2e11ee8417d, []int{1}
}

func (m *TransferList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferList.Unmarshal(m, b)
}
func (m *TransferList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferList.Marshal(b, m, deterministic)
}
func (m *TransferList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferList.Merge(m, src)
}
func (m *TransferList) XXX_Size() int {
	return xxx_messageInfo_TransferList.Size(m)
}
func (m *TransferList) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferList.DiscardUnknown(m)
}

var xxx_messageInfo_TransferList proto.InternalMessageInfo

func (m *TransferList) GetAccountAmounts() []*AccountAmount {
	if m != nil {
		return m.AccountAmounts
	}
	return nil
}

// Transfer cryptocurrency from some accounts to other accounts. The accounts list can contain up to 10 accounts. The amounts list must be the same length as the accounts list. Each negative amount is withdrawn from the corresponding account (a sender), and each positive one is added to the corresponding account (a receiver). The amounts list must sum to zero. Each amount is a number of tinyBars (there are 100,000,000 tinyBars in one Hbar). If any sender account fails to have sufficient hbars to do the withdrawal, then the entire transaction fails, and none of those transfers occur, though the transaction fee is still charged. This transaction must be signed by the keys for all the sending accounts, and for any receiving accounts that have receiverSigRequired == true. The signatures are in the same order as the accounts, skipping those accounts that don't need a signature.
type CryptoTransferTransactionBody struct {
	Transfers            *TransferList `protobuf:"bytes,1,opt,name=transfers,proto3" json:"transfers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CryptoTransferTransactionBody) Reset()         { *m = CryptoTransferTransactionBody{} }
func (m *CryptoTransferTransactionBody) String() string { return proto.CompactTextString(m) }
func (*CryptoTransferTransactionBody) ProtoMessage()    {}
func (*CryptoTransferTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_1807c2e11ee8417d, []int{2}
}

func (m *CryptoTransferTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoTransferTransactionBody.Unmarshal(m, b)
}
func (m *CryptoTransferTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoTransferTransactionBody.Marshal(b, m, deterministic)
}
func (m *CryptoTransferTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoTransferTransactionBody.Merge(m, src)
}
func (m *CryptoTransferTransactionBody) XXX_Size() int {
	return xxx_messageInfo_CryptoTransferTransactionBody.Size(m)
}
func (m *CryptoTransferTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoTransferTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoTransferTransactionBody proto.InternalMessageInfo

func (m *CryptoTransferTransactionBody) GetTransfers() *TransferList {
	if m != nil {
		return m.Transfers
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountAmount)(nil), "proto.AccountAmount")
	proto.RegisterType((*TransferList)(nil), "proto.TransferList")
	proto.RegisterType((*CryptoTransferTransactionBody)(nil), "proto.CryptoTransferTransactionBody")
}

func init() { proto.RegisterFile("CryptoTransfer.proto", fileDescriptor_1807c2e11ee8417d) }

var fileDescriptor_1807c2e11ee8417d = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0x62, 0xa1, 0x53, 0x95, 0x12, 0x8b, 0x2c, 0x82, 0xb0, 0xe4, 0x20, 0x7b, 0x0a,
	0x58, 0xaf, 0x5e, 0x76, 0xf5, 0x22, 0xf4, 0x20, 0xa1, 0xe0, 0x79, 0x4c, 0xa3, 0x89, 0xd0, 0x4d,
	0x48, 0xa2, 0xb0, 0xff, 0x5e, 0x4c, 0x52, 0xba, 0xeb, 0x25, 0xc3, 0xcc, 0xfb, 0xe6, 0xe5, 0x0d,
	0xac, 0x9e, 0xfc, 0xe0, 0xa2, 0xdd, 0x7a, 0xec, 0xc3, 0x87, 0xf2, 0xdc, 0x79, 0x1b, 0x2d, 0x3d,
	0x4b, 0xe5, 0x66, 0xd9, 0x61, 0x30, 0x72, 0x3b, 0x38, 0x15, 0xb2, 0xc0, 0xde, 0xe0, 0xa2, 0x95,
	0xd2, 0x7e, 0xf7, 0xb1, 0xdd, 0xff, 0xbd, 0x94, 0xc3, 0x1c, 0xf3, 0xe0, 0xe5, 0xb9, 0x22, 0x35,
	0x69, 0x16, 0xeb, 0x65, 0x66, 0x79, 0x7b, 0x98, 0x8b, 0x23, 0x42, 0xaf, 0x61, 0x86, 0x69, 0xb3,
	0x3a, 0xa9, 0x49, 0x43, 0x45, 0xe9, 0xd8, 0x06, 0xce, 0x0f, 0x19, 0x36, 0x26, 0x44, 0xfa, 0x08,
	0x97, 0x38, 0xfe, 0x28, 0x54, 0xa4, 0x3e, 0x6d, 0x16, 0xeb, 0xd5, 0xd4, 0x3c, 0x8b, 0xe2, 0x1f,
	0xcb, 0x04, 0xdc, 0x4e, 0xef, 0x4a, 0x15, 0x65, 0x34, 0xb6, 0xef, 0xec, 0x6e, 0xa0, 0xf7, 0x30,
	0x8f, 0x45, 0x0a, 0x25, 0xf6, 0x55, 0x71, 0x1e, 0xc7, 0x10, 0x47, 0xaa, 0xbb, 0x03, 0x26, 0xed,
	0x9e, 0x6b, 0xb5, 0x53, 0x1e, 0x35, 0x06, 0xfd, 0xe9, 0xd1, 0x69, 0x8e, 0xce, 0x94, 0xc5, 0x2f,
	0xfc, 0xc1, 0x57, 0xf2, 0x3e, 0x4b, 0xdd, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x0a,
	0x9e, 0x8d, 0x5a, 0x01, 0x00, 0x00,
}
