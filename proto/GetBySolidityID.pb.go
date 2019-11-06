// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetBySolidityID.proto

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

// Get the IDs in the format used by transactions, given the ID in the format used by Solidity. If the Solidity ID is for a smart contract instance, then both the ContractID and associated AccountID will be returned.
type GetBySolidityIDQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SolidityID           string       `protobuf:"bytes,2,opt,name=solidityID,proto3" json:"solidityID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBySolidityIDQuery) Reset()         { *m = GetBySolidityIDQuery{} }
func (m *GetBySolidityIDQuery) String() string { return proto.CompactTextString(m) }
func (*GetBySolidityIDQuery) ProtoMessage()    {}
func (*GetBySolidityIDQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa3229a059c75ce, []int{0}
}

func (m *GetBySolidityIDQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBySolidityIDQuery.Unmarshal(m, b)
}
func (m *GetBySolidityIDQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBySolidityIDQuery.Marshal(b, m, deterministic)
}
func (m *GetBySolidityIDQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBySolidityIDQuery.Merge(m, src)
}
func (m *GetBySolidityIDQuery) XXX_Size() int {
	return xxx_messageInfo_GetBySolidityIDQuery.Size(m)
}
func (m *GetBySolidityIDQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBySolidityIDQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetBySolidityIDQuery proto.InternalMessageInfo

func (m *GetBySolidityIDQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetBySolidityIDQuery) GetSolidityID() string {
	if m != nil {
		return m.SolidityID
	}
	return ""
}

// Response when the client sends the node GetBySolidityIDQuery
type GetBySolidityIDResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID      `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	FileID               *FileID         `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	ContractID           *ContractID     `protobuf:"bytes,4,opt,name=contractID,proto3" json:"contractID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetBySolidityIDResponse) Reset()         { *m = GetBySolidityIDResponse{} }
func (m *GetBySolidityIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetBySolidityIDResponse) ProtoMessage()    {}
func (*GetBySolidityIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa3229a059c75ce, []int{1}
}

func (m *GetBySolidityIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBySolidityIDResponse.Unmarshal(m, b)
}
func (m *GetBySolidityIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBySolidityIDResponse.Marshal(b, m, deterministic)
}
func (m *GetBySolidityIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBySolidityIDResponse.Merge(m, src)
}
func (m *GetBySolidityIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetBySolidityIDResponse.Size(m)
}
func (m *GetBySolidityIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBySolidityIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBySolidityIDResponse proto.InternalMessageInfo

func (m *GetBySolidityIDResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *GetBySolidityIDResponse) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *GetBySolidityIDResponse) GetFileID() *FileID {
	if m != nil {
		return m.FileID
	}
	return nil
}

func (m *GetBySolidityIDResponse) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBySolidityIDQuery)(nil), "proto.GetBySolidityIDQuery")
	proto.RegisterType((*GetBySolidityIDResponse)(nil), "proto.GetBySolidityIDResponse")
}

func init() { proto.RegisterFile("GetBySolidityID.proto", fileDescriptor_9aa3229a059c75ce) }

var fileDescriptor_9aa3229a059c75ce = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xef, 0x4a, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0x7f, 0x06, 0xbb, 0x43, 0xd8, 0xc2, 0x86, 0x65, 0x1f, 0x64, 0x14, 0x94, 0x21,
	0x58, 0x70, 0x3e, 0x81, 0x75, 0xa8, 0xfb, 0xa6, 0xd1, 0x17, 0xb8, 0x4b, 0xaf, 0x26, 0x32, 0x9b,
	0x90, 0x64, 0x42, 0x5f, 0xd3, 0x27, 0x12, 0xb3, 0xac, 0xd6, 0x7e, 0x0a, 0x39, 0xe7, 0x77, 0xce,
	0xb9, 0x30, 0x79, 0x20, 0x5f, 0xd4, 0x2f, 0x7a, 0xa3, 0x4a, 0xe5, 0xeb, 0xd5, 0x32, 0x37, 0x56,
	0x7b, 0xcd, 0x8e, 0xc3, 0x33, 0x1d, 0x16, 0xe8, 0x94, 0x78, 0xad, 0x0d, 0xb9, 0x9d, 0x31, 0x1d,
	0x3d, 0x6f, 0xc9, 0xd6, 0x8f, 0x84, 0x25, 0xd9, 0x28, 0x8d, 0x39, 0x39, 0xa3, 0x2b, 0x47, 0x6d,
	0x35, 0x5b, 0xc3, 0xb8, 0x53, 0x1d, 0x92, 0xec, 0x12, 0x7a, 0x32, 0x70, 0x69, 0x32, 0x4b, 0xe6,
	0x83, 0x05, 0xdb, 0xf1, 0x79, 0xab, 0x97, 0x47, 0x82, 0x9d, 0x01, 0xb8, 0x26, 0x9e, 0x1e, 0xcc,
	0x92, 0x79, 0x9f, 0xb7, 0x94, 0xec, 0x3b, 0x81, 0xd3, 0xce, 0xc8, 0xfe, 0x16, 0x76, 0xd5, 0xd9,
	0x99, 0xc4, 0x9d, 0xff, 0xc7, 0x36, 0x53, 0x39, 0xf4, 0x51, 0x08, 0xbd, 0xad, 0x7c, 0x5c, 0x1a,
	0x2c, 0x86, 0x31, 0x71, 0xbb, 0xd7, 0xf9, 0x1f, 0xc2, 0xce, 0xa1, 0xf7, 0xa6, 0x36, 0xb4, 0x5a,
	0xa6, 0x87, 0x01, 0x3e, 0x89, 0xf0, 0x7d, 0x10, 0x79, 0x34, 0xd9, 0x35, 0x80, 0xd0, 0x95, 0xb7,
	0x28, 0x7e, 0x7b, 0x8f, 0x02, 0x3a, 0x8a, 0xe8, 0x5d, 0x63, 0xf0, 0x16, 0x54, 0x5c, 0x40, 0x26,
	0xf4, 0x67, 0x2e, 0xa9, 0x24, 0x8b, 0x12, 0x9d, 0x7c, 0xb7, 0x68, 0x64, 0x8e, 0x46, 0xc5, 0xdc,
	0x07, 0x7e, 0xe1, 0x53, 0xb2, 0xee, 0x85, 0xdf, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61,
	0x72, 0x47, 0x94, 0xc2, 0x01, 0x00, 0x00,
}
