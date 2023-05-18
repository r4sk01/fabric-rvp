// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_NOT_IMPLEMENTED          Status = 501
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	503: "SERVICE_UNAVAILABLE",
}

var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"NOT_IMPLEMENTED":          501,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type HeaderType int32

const (
	HeaderType_MESSAGE              HeaderType = 0
	HeaderType_CONFIG               HeaderType = 1
	HeaderType_CONFIG_UPDATE        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION HeaderType = 3
	HeaderType_ORDERER_TRANSACTION  HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO    HeaderType = 5
	HeaderType_CHAINCODE_PACKAGE    HeaderType = 6
	HeaderType_PEER_ADMIN_OPERATION HeaderType = 8
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",
	2: "CONFIG_UPDATE",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
	6: "CHAINCODE_PACKAGE",
	8: "PEER_ADMIN_OPERATION",
}

var HeaderType_value = map[string]int32{
	"MESSAGE":              0,
	"CONFIG":               1,
	"CONFIG_UPDATE":        2,
	"ENDORSER_TRANSACTION": 3,
	"ORDERER_TRANSACTION":  4,
	"DELIVER_SEEK_INFO":    5,
	"CHAINCODE_PACKAGE":    6,
	"PEER_ADMIN_OPERATION": 8,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}

func (HeaderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

// This enum enlists indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIG         BlockMetadataIndex = 1 // Deprecated: Do not use.
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
	BlockMetadataIndex_ORDERER             BlockMetadataIndex = 3 // Deprecated: Do not use.
	BlockMetadataIndex_COMMIT_HASH         BlockMetadataIndex = 4
	BlockMetadataIndex_LOCAL_DIGEST        BlockMetadataIndex = 5
	BlockMetadataIndex_GLOBAL_DIGEST       BlockMetadataIndex = 6
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIG",
	2: "TRANSACTIONS_FILTER",
	3: "ORDERER",
	4: "COMMIT_HASH",
	5: "LOCAL_DIGEST",
	6: "GLOBAL_DIGEST",
}

var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIG":         1,
	"TRANSACTIONS_FILTER": 2,
	"ORDERER":             3,
	"COMMIT_HASH":         4,
	"LOCAL_DIGEST":        5,
	"GLOBAL_DIGEST":       6,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}

func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

// LastConfig is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfig struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastConfig) Reset()         { *m = LastConfig{} }
func (m *LastConfig) String() string { return proto.CompactTextString(m) }
func (*LastConfig) ProtoMessage()    {}
func (*LastConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *LastConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastConfig.Unmarshal(m, b)
}
func (m *LastConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastConfig.Marshal(b, m, deterministic)
}
func (m *LastConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastConfig.Merge(m, src)
}
func (m *LastConfig) XXX_Size() int {
	return xxx_messageInfo_LastConfig.Size(m)
}
func (m *LastConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LastConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LastConfig proto.InternalMessageInfo

func (m *LastConfig) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value                []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures           []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader      []byte   `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataSignature) Reset()         { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()    {}
func (*MetadataSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *MetadataSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataSignature.Unmarshal(m, b)
}
func (m *MetadataSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataSignature.Marshal(b, m, deterministic)
}
func (m *MetadataSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataSignature.Merge(m, src)
}
func (m *MetadataSignature) XXX_Size() int {
	return xxx_messageInfo_MetadataSignature.Size(m)
}
func (m *MetadataSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataSignature proto.InternalMessageInfo

func (m *MetadataSignature) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *MetadataSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Header struct {
	ChannelHeader        []byte   `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	SignatureHeader      []byte   `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetChannelHeader() []byte {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
	// If mutual TLS is employed, this represents
	// the hash of the client's TLS certificate
	TlsCertHash          []byte   `protobuf:"bytes,8,opt,name=tls_cert_hash,json=tlsCertHash,proto3" json:"tls_cert_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelHeader) Reset()         { *m = ChannelHeader{} }
func (m *ChannelHeader) String() string { return proto.CompactTextString(m) }
func (*ChannelHeader) ProtoMessage()    {}
func (*ChannelHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *ChannelHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelHeader.Unmarshal(m, b)
}
func (m *ChannelHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelHeader.Marshal(b, m, deterministic)
}
func (m *ChannelHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelHeader.Merge(m, src)
}
func (m *ChannelHeader) XXX_Size() int {
	return xxx_messageInfo_ChannelHeader.Size(m)
}
func (m *ChannelHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelHeader proto.InternalMessageInfo

func (m *ChannelHeader) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ChannelHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ChannelHeader) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChannelHeader) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelHeader) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChannelHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ChannelHeader) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ChannelHeader) GetTlsCertHash() []byte {
	if m != nil {
		return m.TlsCertHash
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, a marshaled msp.SerializedIdentity
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureHeader) Reset()         { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()    {}
func (*SignatureHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *SignatureHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureHeader.Unmarshal(m, b)
}
func (m *SignatureHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureHeader.Marshal(b, m, deterministic)
}
func (m *SignatureHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureHeader.Merge(m, src)
}
func (m *SignatureHeader) XXX_Size() int {
	return xxx_messageInfo_SignatureHeader.Size(m)
}
func (m *SignatureHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureHeader proto.InternalMessageInfo

func (m *SignatureHeader) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *SignatureHeader) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{6}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{7}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *BlockData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Metadata             *BlockMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{8}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BlockHeader is the element of the block which forms the block chain
// The block header is hashed using the configured chain hashing algorithm
// over the ASN.1 encoding of the BlockHeader
type BlockHeader struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	PreviousHash         []byte   `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash             []byte   `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	LocalHash            []byte   `protobuf:"bytes,4,opt,name=local_hash,json=localHash,proto3" json:"local_hash,omitempty"`
	GlobalHash           []byte   `protobuf:"bytes,5,opt,name=global_hash,json=globalHash,proto3" json:"global_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{9}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (x *BlockHeader) GetLocalHash() []byte {
	if x != nil {
		return x.LocalHash
	}
	return nil
}

func (x *BlockHeader) GetGlobalHash() []byte {
	if x != nil {
		return x.GlobalHash
	}
	return nil
}

type BlockData struct {
	Data                 [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{10}
}

func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockData.Unmarshal(m, b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
}
func (m *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(m, src)
}
func (m *BlockData) XXX_Size() int {
	return xxx_messageInfo_BlockData.Size(m)
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockMetadata struct {
	Metadata             [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMetadata) Reset()         { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()    {}
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{11}
}

func (m *BlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMetadata.Unmarshal(m, b)
}
func (m *BlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMetadata.Marshal(b, m, deterministic)
}
func (m *BlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMetadata.Merge(m, src)
}
func (m *BlockMetadata) XXX_Size() int {
	return xxx_messageInfo_BlockMetadata.Size(m)
}
func (m *BlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMetadata proto.InternalMessageInfo

func (m *BlockMetadata) GetMetadata() [][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// OrdererBlockMetadata defines metadata that is set by the ordering service.
type OrdererBlockMetadata struct {
	LastConfig           *LastConfig `protobuf:"bytes,1,opt,name=last_config,json=lastConfig,proto3" json:"last_config,omitempty"`
	ConsenterMetadata    []byte      `protobuf:"bytes,2,opt,name=consenter_metadata,json=consenterMetadata,proto3" json:"consenter_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OrdererBlockMetadata) Reset()         { *m = OrdererBlockMetadata{} }
func (m *OrdererBlockMetadata) String() string { return proto.CompactTextString(m) }
func (*OrdererBlockMetadata) ProtoMessage()    {}
func (*OrdererBlockMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{12}
}

func (m *OrdererBlockMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdererBlockMetadata.Unmarshal(m, b)
}
func (m *OrdererBlockMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdererBlockMetadata.Marshal(b, m, deterministic)
}
func (m *OrdererBlockMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdererBlockMetadata.Merge(m, src)
}
func (m *OrdererBlockMetadata) XXX_Size() int {
	return xxx_messageInfo_OrdererBlockMetadata.Size(m)
}
func (m *OrdererBlockMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdererBlockMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OrdererBlockMetadata proto.InternalMessageInfo

func (m *OrdererBlockMetadata) GetLastConfig() *LastConfig {
	if m != nil {
		return m.LastConfig
	}
	return nil
}

func (m *OrdererBlockMetadata) GetConsenterMetadata() []byte {
	if m != nil {
		return m.ConsenterMetadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
	proto.RegisterType((*LastConfig)(nil), "common.LastConfig")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChannelHeader)(nil), "common.ChannelHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterType((*OrdererBlockMetadata)(nil), "common.OrdererBlockMetadata")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 1052 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x37, 0x71, 0x7e, 0x3e, 0x6f, 0x5a, 0x67, 0xd2, 0x82, 0x29, 0xac, 0xb6, 0x32, 0x2c,
	0x2a, 0x5d, 0x35, 0x15, 0xdd, 0x0b, 0x1c, 0x1d, 0x7b, 0xda, 0x5a, 0x4d, 0xec, 0x30, 0x76, 0x8a,
	0x58, 0x90, 0x46, 0x6e, 0x32, 0x4d, 0x22, 0x12, 0x3b, 0xb2, 0x27, 0x55, 0xcb, 0x95, 0x3b, 0x42,
	0x82, 0x2b, 0xff, 0x0b, 0x47, 0xfe, 0x14, 0xfe, 0x00, 0x10, 0x57, 0x64, 0x8f, 0xed, 0x26, 0x65,
	0x25, 0x4e, 0xf1, 0x7b, 0xf3, 0xf1, 0x7b, 0xdf, 0xf7, 0x23, 0x63, 0xe8, 0x8c, 0xc3, 0xe5, 0x32,
	0x0c, 0x4e, 0xc5, 0x4f, 0x77, 0x15, 0x85, 0x3c, 0x44, 0x35, 0x61, 0x1d, 0xbc, 0x9c, 0x86, 0xe1,
	0x74, 0xc1, 0x4e, 0x53, 0xef, 0xcd, 0xfa, 0xf6, 0x94, 0xcf, 0x97, 0x2c, 0xe6, 0xfe, 0x72, 0x25,
	0x40, 0x4d, 0x03, 0xe8, 0xfb, 0x31, 0x37, 0xc2, 0xe0, 0x76, 0x3e, 0x45, 0x7b, 0x50, 0x9d, 0x07,
	0x13, 0x76, 0xaf, 0x96, 0x0e, 0x4b, 0x47, 0x15, 0x22, 0x0c, 0xed, 0x5b, 0x68, 0x0c, 0x18, 0xf7,
	0x27, 0x3e, 0xf7, 0x13, 0xe2, 0xce, 0x5f, 0xac, 0x59, 0x4a, 0x3c, 0x27, 0xc2, 0x40, 0x5f, 0x02,
	0xc4, 0xf3, 0x69, 0xe0, 0xf3, 0x75, 0xc4, 0x62, 0xb5, 0x7c, 0x28, 0x1d, 0xc9, 0x67, 0x1f, 0x74,
	0x33, 0x45, 0xf9, 0xbb, 0x6e, 0x4e, 0x90, 0x0d, 0x58, 0xfb, 0x0e, 0xda, 0xff, 0x01, 0xd0, 0x67,
	0xa0, 0x14, 0x08, 0x9d, 0x31, 0x7f, 0xc2, 0xa2, 0x2c, 0xe1, 0x6e, 0xe1, 0xbf, 0x4c, 0xdd, 0xe8,
	0x23, 0x68, 0x16, 0x2e, 0xb5, 0x9c, 0x32, 0x8f, 0x0e, 0xed, 0x2d, 0xd4, 0x32, 0xee, 0x15, 0xec,
	0x8c, 0x67, 0x7e, 0x10, 0xb0, 0xc5, 0x76, 0xc0, 0x56, 0xe6, 0xcd, 0xb0, 0x77, 0x65, 0x2e, 0xbf,
	0x33, 0xb3, 0xf6, 0x63, 0x19, 0x5a, 0xc6, 0xd6, 0xcb, 0x08, 0x2a, 0xfc, 0x61, 0x25, 0x7a, 0x53,
	0x25, 0xe9, 0x33, 0x52, 0xa1, 0x7e, 0xc7, 0xa2, 0x78, 0x1e, 0x06, 0x69, 0x9c, 0x2a, 0xc9, 0x4d,
	0xf4, 0x05, 0x34, 0x8b, 0x69, 0xa8, 0xd2, 0x61, 0xe9, 0x48, 0x3e, 0x3b, 0xe8, 0x8a, 0x79, 0x75,
	0xf3, 0x79, 0x75, 0xbd, 0x9c, 0x20, 0x8f, 0x30, 0x7a, 0x01, 0x90, 0xd7, 0x32, 0x9f, 0xa8, 0x95,
	0xc3, 0xd2, 0x51, 0x93, 0x34, 0x33, 0x8f, 0x35, 0x41, 0x1d, 0xa8, 0xf2, 0xfb, 0xe4, 0xa4, 0x9a,
	0x9e, 0x54, 0xf8, 0xbd, 0x35, 0x49, 0x06, 0xc7, 0x56, 0xe1, 0x78, 0xa6, 0xd6, 0xc4, 0x68, 0x53,
	0x23, 0xe9, 0x1e, 0xbb, 0xe7, 0x2c, 0x48, 0xf5, 0xd5, 0x45, 0xf7, 0x0a, 0x07, 0xd2, 0xa0, 0xc5,
	0x17, 0x31, 0x1d, 0xb3, 0x88, 0xd3, 0x99, 0x1f, 0xcf, 0xd4, 0x46, 0x4a, 0xc8, 0x7c, 0x11, 0x1b,
	0x2c, 0xe2, 0x97, 0x7e, 0x3c, 0xd3, 0x74, 0xd8, 0x75, 0x9f, 0x8c, 0x44, 0x85, 0xfa, 0x38, 0x62,
	0x3e, 0x0f, 0xf3, 0x1e, 0xe7, 0x66, 0x22, 0x22, 0x08, 0x83, 0x71, 0x3e, 0x28, 0x61, 0x68, 0x18,
	0xea, 0x43, 0xff, 0x61, 0x11, 0xfa, 0x13, 0xf4, 0x29, 0xd4, 0x36, 0xa6, 0x23, 0x9f, 0xed, 0xe4,
	0x4b, 0x24, 0x42, 0x93, 0xec, 0x34, 0xe9, 0x74, 0xb2, 0x31, 0x59, 0x9c, 0xf4, 0x59, 0xeb, 0x41,
	0x03, 0x07, 0x77, 0x6c, 0x11, 0x8a, 0xae, 0xaf, 0x44, 0xc8, 0x5c, 0x42, 0x66, 0xfe, 0xcf, 0xbe,
	0xfc, 0x54, 0x82, 0x6a, 0x6f, 0x11, 0x8e, 0xbf, 0x47, 0xaf, 0x9f, 0x28, 0xe9, 0xe4, 0x4a, 0xd2,
	0xe3, 0x27, 0x72, 0x5e, 0x6d, 0xc8, 0x91, 0xcf, 0xda, 0x5b, 0xa8, 0xe9, 0x73, 0x5f, 0x28, 0x44,
	0x9f, 0x43, 0x63, 0x99, 0xed, 0x7a, 0x36, 0xf0, 0xfd, 0x2d, 0x34, 0xff, 0x23, 0x90, 0x02, 0xd3,
	0xa6, 0x20, 0x6f, 0x24, 0x44, 0xef, 0x41, 0x2d, 0x58, 0x2f, 0x6f, 0x32, 0x55, 0x15, 0x92, 0x59,
	0xe8, 0x63, 0x68, 0xad, 0x22, 0x76, 0x37, 0x0f, 0xd7, 0xb1, 0x98, 0x94, 0xa8, 0xec, 0x79, 0xee,
	0x4c, 0x46, 0x85, 0x3e, 0x84, 0x66, 0x12, 0x53, 0x00, 0x52, 0x0a, 0x34, 0x12, 0x47, 0x3a, 0xc7,
	0x97, 0xd0, 0x2c, 0xe4, 0x16, 0xed, 0x2d, 0x1d, 0x4a, 0x45, 0x7b, 0x5f, 0x43, 0x6b, 0x4b, 0x24,
	0x3a, 0xd8, 0xa8, 0x46, 0x80, 0x8f, 0xb2, 0x7f, 0x80, 0x3d, 0x27, 0x9a, 0xb0, 0x88, 0x45, 0xdb,
	0xef, 0xbc, 0x01, 0x79, 0xe1, 0xc7, 0x9c, 0x8e, 0xd3, 0xfb, 0x26, 0x6b, 0x2d, 0xca, 0x9b, 0xf0,
	0x78, 0x13, 0x11, 0x58, 0x3c, 0xde, 0x4a, 0x27, 0x80, 0xc6, 0x61, 0x10, 0xb3, 0x80, 0xb3, 0x88,
	0x16, 0x29, 0x45, 0x85, 0xed, 0xe2, 0x24, 0xcf, 0x71, 0xfc, 0x7b, 0x09, 0x6a, 0x2e, 0xf7, 0xf9,
	0x3a, 0x46, 0x32, 0xd4, 0x47, 0xf6, 0x95, 0xed, 0x7c, 0x6d, 0x2b, 0xcf, 0xd0, 0x73, 0xa8, 0xbb,
	0x23, 0xc3, 0xc0, 0xae, 0xab, 0xfc, 0x51, 0x42, 0x0a, 0xc8, 0x3d, 0xdd, 0xa4, 0x04, 0x7f, 0x35,
	0xc2, 0xae, 0xa7, 0xfc, 0x2c, 0xa1, 0x1d, 0x68, 0x9e, 0x3b, 0xa4, 0x67, 0x99, 0x26, 0xb6, 0x95,
	0x5f, 0x52, 0xdb, 0x76, 0x3c, 0x7a, 0xee, 0x8c, 0x6c, 0x53, 0xf9, 0x55, 0x42, 0x2f, 0x40, 0xcd,
	0x68, 0x8a, 0x6d, 0xcf, 0xf2, 0xbe, 0xa1, 0x9e, 0xe3, 0xd0, 0xbe, 0x4e, 0x2e, 0xb0, 0xf2, 0x9b,
	0x84, 0x0e, 0x60, 0xdf, 0xb2, 0x3d, 0x4c, 0x6c, 0xbd, 0x4f, 0x5d, 0x4c, 0xae, 0x31, 0xa1, 0x98,
	0x10, 0x87, 0x28, 0x7f, 0x49, 0x68, 0x0f, 0x76, 0x93, 0x50, 0xd6, 0x60, 0xd8, 0xc7, 0x03, 0x6c,
	0x7b, 0xd8, 0x54, 0xfe, 0x96, 0x90, 0x0a, 0x9d, 0x04, 0xb4, 0x0c, 0x4c, 0x47, 0xb6, 0x7e, 0xad,
	0x5b, 0x7d, 0xbd, 0xd7, 0xc7, 0xca, 0x3f, 0xd2, 0xf1, 0x9f, 0x25, 0x00, 0x31, 0x71, 0x2f, 0xb9,
	0x43, 0x64, 0xa8, 0x0f, 0xb0, 0xeb, 0xea, 0x17, 0x58, 0x79, 0x86, 0x00, 0x6a, 0x86, 0x63, 0x9f,
	0x5b, 0x17, 0x4a, 0x09, 0xb5, 0xa1, 0x25, 0x9e, 0xe9, 0x68, 0x68, 0xea, 0x1e, 0x56, 0xca, 0x48,
	0x85, 0x3d, 0x6c, 0x9b, 0x0e, 0x71, 0x31, 0xa1, 0x1e, 0xd1, 0x6d, 0x57, 0x37, 0x3c, 0xcb, 0xb1,
	0x15, 0x09, 0xbd, 0x0f, 0x1d, 0x87, 0x98, 0x98, 0x3c, 0x39, 0xa8, 0xa0, 0x7d, 0x68, 0x9b, 0xb8,
	0x6f, 0x25, 0x8a, 0x5d, 0x8c, 0xaf, 0xa8, 0x65, 0x9f, 0x3b, 0x4a, 0x35, 0x71, 0x1b, 0x97, 0xba,
	0x65, 0x1b, 0x8e, 0x89, 0xe9, 0x50, 0x37, 0xae, 0x92, 0xfc, 0xb5, 0x24, 0xc1, 0x10, 0x63, 0x42,
	0x75, 0x73, 0x60, 0xd9, 0xd4, 0x19, 0x62, 0xa2, 0xa7, 0x71, 0x1a, 0x5a, 0xa5, 0x51, 0x57, 0xea,
	0x5a, 0xa5, 0xd1, 0x54, 0x9a, 0xc7, 0x82, 0x21, 0xd8, 0x75, 0x46, 0x24, 0xa9, 0x2f, 0x95, 0x77,
	0xdc, 0xf6, 0x9c, 0x2b, 0x6c, 0x6f, 0xa6, 0x3f, 0xe6, 0x80, 0xb6, 0x56, 0xc4, 0x4a, 0x3e, 0x39,
	0x68, 0x07, 0xc0, 0xb5, 0x2e, 0x6c, 0xdd, 0x1b, 0x11, 0xec, 0x2a, 0xcf, 0x50, 0x07, 0xe4, 0xbe,
	0xee, 0x7a, 0x34, 0xaf, 0xfc, 0xa0, 0xdc, 0x28, 0x25, 0x05, 0x6d, 0x44, 0x72, 0xe9, 0xb9, 0xd5,
	0xf7, 0x30, 0x51, 0xca, 0x68, 0x17, 0xea, 0x59, 0xa5, 0x8a, 0x94, 0x92, 0xbb, 0x20, 0x1b, 0xce,
	0x60, 0x60, 0x79, 0xf4, 0x52, 0x77, 0x2f, 0x95, 0x4a, 0xef, 0x1a, 0x3e, 0x09, 0xa3, 0x69, 0x77,
	0xf6, 0xb0, 0x62, 0xd1, 0x82, 0x4d, 0xa6, 0x2c, 0xea, 0xde, 0xfa, 0x37, 0xd1, 0x7c, 0x2c, 0x6e,
	0xde, 0x38, 0xdb, 0xc8, 0xb7, 0xdd, 0xe9, 0x9c, 0xcf, 0xd6, 0x37, 0x89, 0x79, 0xba, 0x01, 0x9f,
	0x0a, 0xf8, 0x44, 0xc0, 0x27, 0xd3, 0x30, 0xfb, 0xfa, 0xde, 0xd4, 0x52, 0xcf, 0x9b, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x72, 0xa9, 0x81, 0xd5, 0x95, 0x07, 0x00, 0x00,
}
