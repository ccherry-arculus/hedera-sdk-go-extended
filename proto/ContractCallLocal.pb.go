// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ContractCallLocal.proto

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

// The log information for an event returned by a smart contract function call. One function call may return several such events.
type ContractLoginfo struct {
	ContractID           *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Bloom                []byte      `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`
	Topic                [][]byte    `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	Data                 []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ContractLoginfo) Reset()         { *m = ContractLoginfo{} }
func (m *ContractLoginfo) String() string { return proto.CompactTextString(m) }
func (*ContractLoginfo) ProtoMessage()    {}
func (*ContractLoginfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_676e8f21aa93aa2e, []int{0}
}

func (m *ContractLoginfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractLoginfo.Unmarshal(m, b)
}
func (m *ContractLoginfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractLoginfo.Marshal(b, m, deterministic)
}
func (m *ContractLoginfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractLoginfo.Merge(m, src)
}
func (m *ContractLoginfo) XXX_Size() int {
	return xxx_messageInfo_ContractLoginfo.Size(m)
}
func (m *ContractLoginfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractLoginfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractLoginfo proto.InternalMessageInfo

func (m *ContractLoginfo) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractLoginfo) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

func (m *ContractLoginfo) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *ContractLoginfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The result returned by a call to a smart contract function. This is part of the response to a ContractCallLocal query, and is in the record for a ContractCall or ContractCreateInstance transaction. The ContractCreateInstance transaction record has the results of the call to the constructor.
type ContractFunctionResult struct {
	ContractID           *ContractID        `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	ContractCallResult   []byte             `protobuf:"bytes,2,opt,name=contractCallResult,proto3" json:"contractCallResult,omitempty"`
	ErrorMessage         string             `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Bloom                []byte             `protobuf:"bytes,4,opt,name=bloom,proto3" json:"bloom,omitempty"`
	GasUsed              uint64             `protobuf:"varint,5,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	LogInfo              []*ContractLoginfo `protobuf:"bytes,6,rep,name=logInfo,proto3" json:"logInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContractFunctionResult) Reset()         { *m = ContractFunctionResult{} }
func (m *ContractFunctionResult) String() string { return proto.CompactTextString(m) }
func (*ContractFunctionResult) ProtoMessage()    {}
func (*ContractFunctionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_676e8f21aa93aa2e, []int{1}
}

func (m *ContractFunctionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractFunctionResult.Unmarshal(m, b)
}
func (m *ContractFunctionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractFunctionResult.Marshal(b, m, deterministic)
}
func (m *ContractFunctionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractFunctionResult.Merge(m, src)
}
func (m *ContractFunctionResult) XXX_Size() int {
	return xxx_messageInfo_ContractFunctionResult.Size(m)
}
func (m *ContractFunctionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractFunctionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ContractFunctionResult proto.InternalMessageInfo

func (m *ContractFunctionResult) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractFunctionResult) GetContractCallResult() []byte {
	if m != nil {
		return m.ContractCallResult
	}
	return nil
}

func (m *ContractFunctionResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *ContractFunctionResult) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

func (m *ContractFunctionResult) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ContractFunctionResult) GetLogInfo() []*ContractLoginfo {
	if m != nil {
		return m.LogInfo
	}
	return nil
}

// Call a function of the given smart contract instance, giving it functionParameters as its inputs. It will consume the entire given amount of gas.
//
// This is performed locally on the particular node that the client is communicating with. It cannot change the state of the contract instance (and so, cannot spend anything from the instance's cryptocurrency account). It will not have a consensus timestamp. It cannot generate a record or a receipt. The response will contain the output returned by the function call.  This is useful for calling getter functions, which purely read the state and don't change it. It is faster and cheaper than a normal call, because it is purely local to a single  node.
type ContractCallLocalQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Gas                  int64        `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	FunctionParameters   []byte       `protobuf:"bytes,4,opt,name=functionParameters,proto3" json:"functionParameters,omitempty"`
	MaxResultSize        int64        `protobuf:"varint,5,opt,name=maxResultSize,proto3" json:"maxResultSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractCallLocalQuery) Reset()         { *m = ContractCallLocalQuery{} }
func (m *ContractCallLocalQuery) String() string { return proto.CompactTextString(m) }
func (*ContractCallLocalQuery) ProtoMessage()    {}
func (*ContractCallLocalQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_676e8f21aa93aa2e, []int{2}
}

func (m *ContractCallLocalQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallLocalQuery.Unmarshal(m, b)
}
func (m *ContractCallLocalQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallLocalQuery.Marshal(b, m, deterministic)
}
func (m *ContractCallLocalQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallLocalQuery.Merge(m, src)
}
func (m *ContractCallLocalQuery) XXX_Size() int {
	return xxx_messageInfo_ContractCallLocalQuery.Size(m)
}
func (m *ContractCallLocalQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallLocalQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallLocalQuery proto.InternalMessageInfo

func (m *ContractCallLocalQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractCallLocalQuery) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractCallLocalQuery) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ContractCallLocalQuery) GetFunctionParameters() []byte {
	if m != nil {
		return m.FunctionParameters
	}
	return nil
}

func (m *ContractCallLocalQuery) GetMaxResultSize() int64 {
	if m != nil {
		return m.MaxResultSize
	}
	return 0
}

// Response when the client sends the node ContractCallLocalQuery
type ContractCallLocalResponse struct {
	Header               *ResponseHeader         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FunctionResult       *ContractFunctionResult `protobuf:"bytes,2,opt,name=functionResult,proto3" json:"functionResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ContractCallLocalResponse) Reset()         { *m = ContractCallLocalResponse{} }
func (m *ContractCallLocalResponse) String() string { return proto.CompactTextString(m) }
func (*ContractCallLocalResponse) ProtoMessage()    {}
func (*ContractCallLocalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_676e8f21aa93aa2e, []int{3}
}

func (m *ContractCallLocalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallLocalResponse.Unmarshal(m, b)
}
func (m *ContractCallLocalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallLocalResponse.Marshal(b, m, deterministic)
}
func (m *ContractCallLocalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallLocalResponse.Merge(m, src)
}
func (m *ContractCallLocalResponse) XXX_Size() int {
	return xxx_messageInfo_ContractCallLocalResponse.Size(m)
}
func (m *ContractCallLocalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallLocalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallLocalResponse proto.InternalMessageInfo

func (m *ContractCallLocalResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractCallLocalResponse) GetFunctionResult() *ContractFunctionResult {
	if m != nil {
		return m.FunctionResult
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractLoginfo)(nil), "proto.ContractLoginfo")
	proto.RegisterType((*ContractFunctionResult)(nil), "proto.ContractFunctionResult")
	proto.RegisterType((*ContractCallLocalQuery)(nil), "proto.ContractCallLocalQuery")
	proto.RegisterType((*ContractCallLocalResponse)(nil), "proto.ContractCallLocalResponse")
}

func init() { proto.RegisterFile("ContractCallLocal.proto", fileDescriptor_676e8f21aa93aa2e) }

var fileDescriptor_676e8f21aa93aa2e = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xe1, 0x6a, 0xd4, 0x40,
	0x10, 0xc7, 0xd9, 0xcb, 0xdd, 0x15, 0xa7, 0x55, 0xdb, 0xa5, 0xd6, 0x58, 0x10, 0x42, 0x10, 0x09,
	0x82, 0x41, 0xeb, 0x1b, 0xf4, 0x54, 0x2c, 0x54, 0xa8, 0xab, 0x3e, 0xc0, 0x74, 0x6f, 0x2e, 0x89,
	0xe4, 0xb2, 0x61, 0x77, 0x4f, 0xac, 0xdf, 0x04, 0x5f, 0xc0, 0x27, 0xf4, 0x55, 0x24, 0x9b, 0xcd,
	0x35, 0x89, 0xf7, 0x41, 0xfc, 0x94, 0x9d, 0xff, 0x4c, 0x86, 0xf9, 0xfd, 0x67, 0xe0, 0xe1, 0x42,
	0x55, 0x56, 0xa3, 0xb4, 0x0b, 0x2c, 0xcb, 0x4b, 0x25, 0xb1, 0x4c, 0x6b, 0xad, 0xac, 0xe2, 0x33,
	0xf7, 0x39, 0x3d, 0x3c, 0x47, 0x53, 0xc8, 0x4f, 0x37, 0x35, 0x99, 0x36, 0x71, 0x7a, 0xf4, 0x61,
	0x43, 0xfa, 0xe6, 0x1d, 0xe1, 0x92, 0xb4, 0x97, 0x8e, 0x05, 0x99, 0x5a, 0x55, 0x86, 0xfa, 0x6a,
	0xfc, 0x93, 0xc1, 0xfd, 0xae, 0xfb, 0xa5, 0xca, 0x8a, 0x6a, 0xa5, 0xf8, 0x4b, 0x00, 0xe9, 0xa5,
	0x8b, 0xd7, 0x21, 0x8b, 0x58, 0xb2, 0x7f, 0x76, 0xd4, 0xd6, 0xa7, 0x8b, 0x6d, 0x42, 0xf4, 0x8a,
	0xf8, 0x31, 0xcc, 0xae, 0x4b, 0xa5, 0xd6, 0xe1, 0x24, 0x62, 0xc9, 0x81, 0x68, 0x83, 0x46, 0xb5,
	0xaa, 0x2e, 0x64, 0x18, 0x44, 0x41, 0xa3, 0xba, 0x80, 0x73, 0x98, 0x2e, 0xd1, 0x62, 0x38, 0x75,
	0xa5, 0xee, 0x1d, 0xff, 0x98, 0xc0, 0x49, 0xd7, 0xfa, 0xed, 0xa6, 0x92, 0xb6, 0x50, 0x95, 0x20,
	0xb3, 0x29, 0xed, 0xff, 0x4c, 0x93, 0x02, 0x97, 0x3d, 0xc7, 0xda, 0x46, 0x7e, 0xb4, 0x1d, 0x19,
	0x1e, 0xc3, 0x01, 0x69, 0xad, 0xf4, 0x7b, 0x32, 0x06, 0x33, 0x0a, 0x83, 0x88, 0x25, 0x77, 0xc4,
	0x40, 0xbb, 0x25, 0x9c, 0xf6, 0x09, 0x43, 0xd8, 0xcb, 0xd0, 0x7c, 0x36, 0xb4, 0x0c, 0x67, 0x11,
	0x4b, 0xa6, 0xa2, 0x0b, 0xf9, 0x0b, 0xd8, 0x2b, 0x55, 0x76, 0x51, 0xad, 0x54, 0x38, 0x8f, 0x82,
	0x64, 0xff, 0xec, 0x64, 0x34, 0xb3, 0x77, 0x5b, 0x74, 0x65, 0xf1, 0x6f, 0x76, 0xeb, 0xc1, 0x76,
	0xd1, 0x6e, 0x8f, 0xfc, 0x19, 0xcc, 0x73, 0xb7, 0x35, 0xcf, 0xcf, 0x7d, 0xaf, 0xde, 0x96, 0x85,
	0xaf, 0x18, 0xf9, 0x35, 0xf9, 0x17, 0xbf, 0x0e, 0x21, 0xc8, 0xd0, 0x38, 0xec, 0x40, 0x34, 0xcf,
	0xc6, 0xc1, 0x95, 0x5f, 0xc3, 0x15, 0x6a, 0x5c, 0x93, 0x25, 0x6d, 0x3c, 0xfa, 0x8e, 0x0c, 0x7f,
	0x02, 0x77, 0xd7, 0xf8, 0xad, 0xb5, 0xf3, 0x63, 0xf1, 0x9d, 0x9c, 0x1b, 0x81, 0x18, 0x8a, 0xf1,
	0x2f, 0x06, 0x8f, 0xfe, 0x22, 0xec, 0xce, 0x92, 0x3f, 0x1f, 0x41, 0x3e, 0xf0, 0x43, 0x0f, 0xef,
	0x76, 0xcb, 0xf9, 0x06, 0xee, 0xad, 0x06, 0x97, 0xe2, 0x59, 0x1f, 0x8f, 0x58, 0x87, 0xe7, 0x24,
	0x46, 0x3f, 0x9d, 0x3f, 0x85, 0x58, 0xaa, 0x75, 0x9a, 0xd3, 0x92, 0x34, 0xe6, 0x68, 0xf2, 0x4c,
	0x63, 0x9d, 0xa7, 0x58, 0x17, 0xbe, 0xcf, 0x17, 0xfc, 0x8a, 0x57, 0xec, 0x7a, 0xee, 0xa2, 0x57,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0x36, 0x6f, 0x84, 0x8c, 0x03, 0x00, 0x00,
}
