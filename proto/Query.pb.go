// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Query.proto

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

// A single query, which is sent from the client to the node. This includes all possible queries. Each Query should not have more than 50 levels.
type Query struct {
	// Types that are valid to be assigned to Query:
	//	*Query_GetByKey
	//	*Query_GetBySolidityID
	//	*Query_ContractCallLocal
	//	*Query_ContractGetInfo
	//	*Query_ContractGetBytecode
	//	*Query_ContractGetRecords
	//	*Query_CryptogetAccountBalance
	//	*Query_CryptoGetAccountRecords
	//	*Query_CryptoGetInfo
	//	*Query_CryptoGetClaim
	//	*Query_CryptoGetProxyStakers
	//	*Query_FileGetContents
	//	*Query_FileGetInfo
	//	*Query_TransactionGetReceipt
	//	*Query_TransactionGetRecord
	//	*Query_TransactionGetFastRecord
	Query                isQuery_Query `protobuf_oneof:"query"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_12565a3f476cd53f, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

type isQuery_Query interface {
	isQuery_Query()
}

type Query_GetByKey struct {
	GetByKey *GetByKeyQuery `protobuf:"bytes,1,opt,name=getByKey,proto3,oneof"`
}

type Query_GetBySolidityID struct {
	GetBySolidityID *GetBySolidityIDQuery `protobuf:"bytes,2,opt,name=getBySolidityID,proto3,oneof"`
}

type Query_ContractCallLocal struct {
	ContractCallLocal *ContractCallLocalQuery `protobuf:"bytes,3,opt,name=contractCallLocal,proto3,oneof"`
}

type Query_ContractGetInfo struct {
	ContractGetInfo *ContractGetInfoQuery `protobuf:"bytes,4,opt,name=contractGetInfo,proto3,oneof"`
}

type Query_ContractGetBytecode struct {
	ContractGetBytecode *ContractGetBytecodeQuery `protobuf:"bytes,5,opt,name=contractGetBytecode,proto3,oneof"`
}

type Query_ContractGetRecords struct {
	ContractGetRecords *ContractGetRecordsQuery `protobuf:"bytes,6,opt,name=ContractGetRecords,proto3,oneof"`
}

type Query_CryptogetAccountBalance struct {
	CryptogetAccountBalance *CryptoGetAccountBalanceQuery `protobuf:"bytes,7,opt,name=cryptogetAccountBalance,proto3,oneof"`
}

type Query_CryptoGetAccountRecords struct {
	CryptoGetAccountRecords *CryptoGetAccountRecordsQuery `protobuf:"bytes,8,opt,name=cryptoGetAccountRecords,proto3,oneof"`
}

type Query_CryptoGetInfo struct {
	CryptoGetInfo *CryptoGetInfoQuery `protobuf:"bytes,9,opt,name=cryptoGetInfo,proto3,oneof"`
}

type Query_CryptoGetClaim struct {
	CryptoGetClaim *CryptoGetClaimQuery `protobuf:"bytes,10,opt,name=cryptoGetClaim,proto3,oneof"`
}

type Query_CryptoGetProxyStakers struct {
	CryptoGetProxyStakers *CryptoGetStakersQuery `protobuf:"bytes,11,opt,name=cryptoGetProxyStakers,proto3,oneof"`
}

type Query_FileGetContents struct {
	FileGetContents *FileGetContentsQuery `protobuf:"bytes,12,opt,name=fileGetContents,proto3,oneof"`
}

type Query_FileGetInfo struct {
	FileGetInfo *FileGetInfoQuery `protobuf:"bytes,13,opt,name=fileGetInfo,proto3,oneof"`
}

type Query_TransactionGetReceipt struct {
	TransactionGetReceipt *TransactionGetReceiptQuery `protobuf:"bytes,14,opt,name=transactionGetReceipt,proto3,oneof"`
}

type Query_TransactionGetRecord struct {
	TransactionGetRecord *TransactionGetRecordQuery `protobuf:"bytes,15,opt,name=transactionGetRecord,proto3,oneof"`
}

type Query_TransactionGetFastRecord struct {
	TransactionGetFastRecord *TransactionGetFastRecordQuery `protobuf:"bytes,16,opt,name=transactionGetFastRecord,proto3,oneof"`
}

func (*Query_GetByKey) isQuery_Query() {}

func (*Query_GetBySolidityID) isQuery_Query() {}

func (*Query_ContractCallLocal) isQuery_Query() {}

func (*Query_ContractGetInfo) isQuery_Query() {}

func (*Query_ContractGetBytecode) isQuery_Query() {}

func (*Query_ContractGetRecords) isQuery_Query() {}

func (*Query_CryptogetAccountBalance) isQuery_Query() {}

func (*Query_CryptoGetAccountRecords) isQuery_Query() {}

func (*Query_CryptoGetInfo) isQuery_Query() {}

func (*Query_CryptoGetClaim) isQuery_Query() {}

func (*Query_CryptoGetProxyStakers) isQuery_Query() {}

func (*Query_FileGetContents) isQuery_Query() {}

func (*Query_FileGetInfo) isQuery_Query() {}

func (*Query_TransactionGetReceipt) isQuery_Query() {}

func (*Query_TransactionGetRecord) isQuery_Query() {}

func (*Query_TransactionGetFastRecord) isQuery_Query() {}

func (m *Query) GetQuery() isQuery_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Query) GetGetByKey() *GetByKeyQuery {
	if x, ok := m.GetQuery().(*Query_GetByKey); ok {
		return x.GetByKey
	}
	return nil
}

func (m *Query) GetGetBySolidityID() *GetBySolidityIDQuery {
	if x, ok := m.GetQuery().(*Query_GetBySolidityID); ok {
		return x.GetBySolidityID
	}
	return nil
}

func (m *Query) GetContractCallLocal() *ContractCallLocalQuery {
	if x, ok := m.GetQuery().(*Query_ContractCallLocal); ok {
		return x.ContractCallLocal
	}
	return nil
}

func (m *Query) GetContractGetInfo() *ContractGetInfoQuery {
	if x, ok := m.GetQuery().(*Query_ContractGetInfo); ok {
		return x.ContractGetInfo
	}
	return nil
}

func (m *Query) GetContractGetBytecode() *ContractGetBytecodeQuery {
	if x, ok := m.GetQuery().(*Query_ContractGetBytecode); ok {
		return x.ContractGetBytecode
	}
	return nil
}

func (m *Query) GetContractGetRecords() *ContractGetRecordsQuery {
	if x, ok := m.GetQuery().(*Query_ContractGetRecords); ok {
		return x.ContractGetRecords
	}
	return nil
}

func (m *Query) GetCryptogetAccountBalance() *CryptoGetAccountBalanceQuery {
	if x, ok := m.GetQuery().(*Query_CryptogetAccountBalance); ok {
		return x.CryptogetAccountBalance
	}
	return nil
}

func (m *Query) GetCryptoGetAccountRecords() *CryptoGetAccountRecordsQuery {
	if x, ok := m.GetQuery().(*Query_CryptoGetAccountRecords); ok {
		return x.CryptoGetAccountRecords
	}
	return nil
}

func (m *Query) GetCryptoGetInfo() *CryptoGetInfoQuery {
	if x, ok := m.GetQuery().(*Query_CryptoGetInfo); ok {
		return x.CryptoGetInfo
	}
	return nil
}

func (m *Query) GetCryptoGetClaim() *CryptoGetClaimQuery {
	if x, ok := m.GetQuery().(*Query_CryptoGetClaim); ok {
		return x.CryptoGetClaim
	}
	return nil
}

func (m *Query) GetCryptoGetProxyStakers() *CryptoGetStakersQuery {
	if x, ok := m.GetQuery().(*Query_CryptoGetProxyStakers); ok {
		return x.CryptoGetProxyStakers
	}
	return nil
}

func (m *Query) GetFileGetContents() *FileGetContentsQuery {
	if x, ok := m.GetQuery().(*Query_FileGetContents); ok {
		return x.FileGetContents
	}
	return nil
}

func (m *Query) GetFileGetInfo() *FileGetInfoQuery {
	if x, ok := m.GetQuery().(*Query_FileGetInfo); ok {
		return x.FileGetInfo
	}
	return nil
}

func (m *Query) GetTransactionGetReceipt() *TransactionGetReceiptQuery {
	if x, ok := m.GetQuery().(*Query_TransactionGetReceipt); ok {
		return x.TransactionGetReceipt
	}
	return nil
}

func (m *Query) GetTransactionGetRecord() *TransactionGetRecordQuery {
	if x, ok := m.GetQuery().(*Query_TransactionGetRecord); ok {
		return x.TransactionGetRecord
	}
	return nil
}

func (m *Query) GetTransactionGetFastRecord() *TransactionGetFastRecordQuery {
	if x, ok := m.GetQuery().(*Query_TransactionGetFastRecord); ok {
		return x.TransactionGetFastRecord
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Query) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Query_GetByKey)(nil),
		(*Query_GetBySolidityID)(nil),
		(*Query_ContractCallLocal)(nil),
		(*Query_ContractGetInfo)(nil),
		(*Query_ContractGetBytecode)(nil),
		(*Query_ContractGetRecords)(nil),
		(*Query_CryptogetAccountBalance)(nil),
		(*Query_CryptoGetAccountRecords)(nil),
		(*Query_CryptoGetInfo)(nil),
		(*Query_CryptoGetClaim)(nil),
		(*Query_CryptoGetProxyStakers)(nil),
		(*Query_FileGetContents)(nil),
		(*Query_FileGetInfo)(nil),
		(*Query_TransactionGetReceipt)(nil),
		(*Query_TransactionGetRecord)(nil),
		(*Query_TransactionGetFastRecord)(nil),
	}
}

func init() {
	proto.RegisterType((*Query)(nil), "proto.Query")
}

func init() { proto.RegisterFile("Query.proto", fileDescriptor_12565a3f476cd53f) }

var fileDescriptor_12565a3f476cd53f = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5f, 0x73, 0x12, 0x3d,
	0x14, 0xc6, 0x5f, 0x5e, 0xa5, 0xad, 0x07, 0x4b, 0x6d, 0x0a, 0x92, 0x52, 0x8b, 0xf5, 0xcf, 0x85,
	0x57, 0x5c, 0xd4, 0x4b, 0xaf, 0xba, 0x74, 0x8a, 0x1d, 0x75, 0x06, 0x4b, 0xc7, 0x19, 0xaf, 0x9c,
	0x34, 0x1b, 0x60, 0xc7, 0xed, 0x06, 0x43, 0x3a, 0xe3, 0x7e, 0x05, 0x3f, 0xb5, 0xc3, 0x21, 0x59,
	0x36, 0xd9, 0xac, 0x57, 0xc0, 0x79, 0xce, 0xf3, 0xdb, 0x2c, 0xe7, 0xc9, 0x81, 0xd6, 0xd7, 0x07,
	0xa1, 0xf2, 0xe1, 0x52, 0x49, 0x2d, 0x49, 0x13, 0x3f, 0xfa, 0xed, 0xb1, 0xd0, 0x51, 0xfe, 0x49,
	0x98, 0x72, 0xbf, 0x8b, 0xbf, 0xa7, 0x32, 0x4d, 0xe2, 0x44, 0xe7, 0xd7, 0x97, 0xa6, 0xdc, 0x1b,
	0xc9, 0x4c, 0x2b, 0xc6, 0xf5, 0x88, 0xa5, 0xe9, 0x67, 0xc9, 0x59, 0x6a, 0xfb, 0xad, 0x30, 0x16,
	0xfa, 0x3a, 0x9b, 0x49, 0x53, 0x3e, 0x2e, 0x95, 0xa3, 0x5c, 0x0b, 0x2e, 0x63, 0x61, 0x24, 0x5a,
	0x92, 0x6e, 0x04, 0x97, 0x2a, 0x5e, 0x19, 0xe5, 0x74, 0xa4, 0xf2, 0xa5, 0x96, 0x63, 0xa1, 0x2f,
	0x38, 0x97, 0x0f, 0x99, 0x8e, 0x58, 0xca, 0x32, 0x2e, 0xea, 0x64, 0xd7, 0x7d, 0x54, 0xc8, 0xa5,
	0x73, 0x74, 0x8a, 0xe2, 0x28, 0x65, 0xc9, 0xbd, 0xa9, 0x3e, 0x2f, 0xaa, 0x53, 0xcd, 0x7e, 0x0a,
	0x65, 0x11, 0xdd, 0xab, 0x24, 0x15, 0xeb, 0x5e, 0x99, 0x69, 0x91, 0x69, 0x5b, 0x3e, 0x34, 0xe5,
	0x12, 0xf7, 0xe4, 0x56, 0xb1, 0x6c, 0xc5, 0xb8, 0x4e, 0x64, 0xb6, 0x79, 0x0f, 0x91, 0x2c, 0xb5,
	0x11, 0xfb, 0x15, 0x51, 0xaa, 0xd8, 0x68, 0x03, 0x57, 0xbb, 0x62, 0x2b, 0x47, 0x7f, 0xfd, 0x07,
	0xa0, 0x89, 0x63, 0x22, 0xe7, 0xb0, 0x37, 0x37, 0xb3, 0xa1, 0x8d, 0xb3, 0xc6, 0xbb, 0xd6, 0x79,
	0x67, 0xd3, 0x33, 0xb4, 0x23, 0xc3, 0xbe, 0x8f, 0xff, 0xdd, 0x14, 0x7d, 0x64, 0x0c, 0x07, 0x73,
	0x77, 0x7e, 0xf4, 0x7f, 0xb4, 0x9e, 0x94, 0xad, 0x5b, 0xd5, 0x12, 0x7c, 0x17, 0xf9, 0x02, 0x87,
	0xdc, 0x9f, 0x38, 0x7d, 0x84, 0xa8, 0x53, 0x83, 0xaa, 0x24, 0xc2, 0xc2, 0xaa, 0xce, 0xf5, 0xb9,
	0xb8, 0x9b, 0x13, 0xfa, 0xd8, 0x39, 0x97, 0x97, 0xa2, 0xe2, 0x5c, 0x9e, 0x8b, 0x4c, 0xe1, 0x88,
	0x57, 0x93, 0x45, 0x9b, 0x08, 0x7b, 0x59, 0x85, 0xd9, 0x0e, 0x0b, 0x0c, 0xb9, 0xc9, 0x04, 0x48,
	0x35, 0x93, 0x74, 0x07, 0x99, 0x83, 0x2a, 0xd3, 0x34, 0x58, 0x64, 0xc0, 0x4b, 0x7e, 0x40, 0x8f,
	0x63, 0xc4, 0xe6, 0x7e, 0x96, 0xe9, 0x2e, 0x62, 0xdf, 0x58, 0x6c, 0x38, 0xf1, 0x96, 0x5d, 0x47,
	0xd9, 0x3e, 0xa0, 0x72, 0x1b, 0xe8, 0xde, 0x3f, 0x1f, 0xe0, 0x1d, 0xbe, 0x8e, 0x42, 0x2e, 0x60,
	0x9f, 0x97, 0xef, 0x13, 0x7d, 0x82, 0xd8, 0x63, 0x1f, 0x5b, 0x9e, 0x96, 0xeb, 0x20, 0x97, 0xd0,
	0xe6, 0xce, 0xed, 0xa3, 0x80, 0x8c, 0xbe, 0xcf, 0x40, 0xd1, 0x42, 0x3c, 0x0f, 0xb9, 0x85, 0x6e,
	0x51, 0x99, 0x28, 0xf9, 0x3b, 0x37, 0x57, 0x96, 0xb6, 0x10, 0xf6, 0xc2, 0x87, 0x19, 0xd9, 0xe2,
	0xc2, 0xe6, 0x75, 0x20, 0x67, 0xee, 0x5d, 0xa7, 0x4f, 0x9d, 0x40, 0x7a, 0x9b, 0xa0, 0x08, 0xa4,
	0xe7, 0x22, 0x1f, 0xa0, 0x35, 0xdb, 0x6e, 0x07, 0xba, 0x8f, 0x90, 0x9e, 0x0b, 0x29, 0xff, 0x47,
	0xe5, 0x6e, 0xf2, 0x1d, 0xba, 0x3a, 0xb4, 0x47, 0x68, 0x1b, 0x31, 0xaf, 0x0c, 0x26, 0xb8, 0x6b,
	0x8a, 0x17, 0x0c, 0x12, 0xc8, 0x37, 0xe8, 0xe8, 0xc0, 0x16, 0xa2, 0x07, 0x48, 0x3e, 0xab, 0x23,
	0x4b, 0x15, 0x5b, 0x70, 0xd0, 0x4f, 0xee, 0x80, 0xea, 0x9a, 0x0d, 0x46, 0x9f, 0x21, 0xfb, 0x6d,
	0x90, 0xbd, 0x6d, 0xb3, 0xfc, 0x5a, 0x4e, 0xb4, 0x0b, 0xcd, 0x5f, 0xeb, 0xa6, 0x68, 0x00, 0x7d,
	0x2e, 0xef, 0x87, 0x0b, 0x11, 0x0b, 0xc5, 0x86, 0x0b, 0xb6, 0x5a, 0xcc, 0x15, 0x5b, 0x2e, 0x36,
	0x0f, 0x98, 0x34, 0xee, 0x76, 0xf0, 0xcb, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x86,
	0x63, 0x63, 0xd9, 0x06, 0x00, 0x00,
}