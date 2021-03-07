// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package types

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

type ResponseEnvelope struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseEnvelope) Reset()         { *m = ResponseEnvelope{} }
func (m *ResponseEnvelope) String() string { return proto.CompactTextString(m) }
func (*ResponseEnvelope) ProtoMessage()    {}
func (*ResponseEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{0}
}

func (m *ResponseEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseEnvelope.Unmarshal(m, b)
}
func (m *ResponseEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseEnvelope.Marshal(b, m, deterministic)
}
func (m *ResponseEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseEnvelope.Merge(m, src)
}
func (m *ResponseEnvelope) XXX_Size() int {
	return xxx_messageInfo_ResponseEnvelope.Size(m)
}
func (m *ResponseEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseEnvelope proto.InternalMessageInfo

func (m *ResponseEnvelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ResponseEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Payload struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Response             []byte          `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{1}
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

func (m *Payload) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

type ResponseHeader struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{2}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

type GetDBStatusResponse struct {
	Exist                bool     `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDBStatusResponse) Reset()         { *m = GetDBStatusResponse{} }
func (m *GetDBStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetDBStatusResponse) ProtoMessage()    {}
func (*GetDBStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{3}
}

func (m *GetDBStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDBStatusResponse.Unmarshal(m, b)
}
func (m *GetDBStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDBStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetDBStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDBStatusResponse.Merge(m, src)
}
func (m *GetDBStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetDBStatusResponse.Size(m)
}
func (m *GetDBStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDBStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDBStatusResponse proto.InternalMessageInfo

func (m *GetDBStatusResponse) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

type GetDataResponse struct {
	Value                []byte    `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{4}
}

func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (m *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(m, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetDataResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type GetUserResponse struct {
	User                 *User     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{5}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type GetConfigResponse struct {
	Config               *ClusterConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Metadata             *Metadata      `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetConfigResponse) Reset()         { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()    {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{6}
}

func (m *GetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigResponse.Unmarshal(m, b)
}
func (m *GetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigResponse.Marshal(b, m, deterministic)
}
func (m *GetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigResponse.Merge(m, src)
}
func (m *GetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetConfigResponse.Size(m)
}
func (m *GetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigResponse proto.InternalMessageInfo

func (m *GetConfigResponse) GetConfig() *ClusterConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetConfigResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type GetNodeConfigResponse struct {
	NodeConfig           *NodeConfig `protobuf:"bytes,1,opt,name=node_config,json=nodeConfig,proto3" json:"node_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetNodeConfigResponse) Reset()         { *m = GetNodeConfigResponse{} }
func (m *GetNodeConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeConfigResponse) ProtoMessage()    {}
func (*GetNodeConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{7}
}

func (m *GetNodeConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeConfigResponse.Unmarshal(m, b)
}
func (m *GetNodeConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeConfigResponse.Marshal(b, m, deterministic)
}
func (m *GetNodeConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeConfigResponse.Merge(m, src)
}
func (m *GetNodeConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetNodeConfigResponse.Size(m)
}
func (m *GetNodeConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeConfigResponse proto.InternalMessageInfo

func (m *GetNodeConfigResponse) GetNodeConfig() *NodeConfig {
	if m != nil {
		return m.NodeConfig
	}
	return nil
}

type GetBlockResponse struct {
	BlockHeader          *BlockHeader `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetBlockResponse) Reset()         { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()    {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{8}
}

func (m *GetBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockResponse.Unmarshal(m, b)
}
func (m *GetBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponse.Merge(m, src)
}
func (m *GetBlockResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockResponse.Size(m)
}
func (m *GetBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponse proto.InternalMessageInfo

func (m *GetBlockResponse) GetBlockHeader() *BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

type GetLedgerPathResponse struct {
	BlockHeaders         []*BlockHeader `protobuf:"bytes,1,rep,name=block_headers,json=blockHeaders,proto3" json:"block_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetLedgerPathResponse) Reset()         { *m = GetLedgerPathResponse{} }
func (m *GetLedgerPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetLedgerPathResponse) ProtoMessage()    {}
func (*GetLedgerPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{9}
}

func (m *GetLedgerPathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLedgerPathResponse.Unmarshal(m, b)
}
func (m *GetLedgerPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLedgerPathResponse.Marshal(b, m, deterministic)
}
func (m *GetLedgerPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLedgerPathResponse.Merge(m, src)
}
func (m *GetLedgerPathResponse) XXX_Size() int {
	return xxx_messageInfo_GetLedgerPathResponse.Size(m)
}
func (m *GetLedgerPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLedgerPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLedgerPathResponse proto.InternalMessageInfo

func (m *GetLedgerPathResponse) GetBlockHeaders() []*BlockHeader {
	if m != nil {
		return m.BlockHeaders
	}
	return nil
}

type GetTxProofResponse struct {
	Hashes               [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxProofResponse) Reset()         { *m = GetTxProofResponse{} }
func (m *GetTxProofResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxProofResponse) ProtoMessage()    {}
func (*GetTxProofResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{10}
}

func (m *GetTxProofResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxProofResponse.Unmarshal(m, b)
}
func (m *GetTxProofResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxProofResponse.Marshal(b, m, deterministic)
}
func (m *GetTxProofResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxProofResponse.Merge(m, src)
}
func (m *GetTxProofResponse) XXX_Size() int {
	return xxx_messageInfo_GetTxProofResponse.Size(m)
}
func (m *GetTxProofResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxProofResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxProofResponse proto.InternalMessageInfo

func (m *GetTxProofResponse) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type GetHistoricalDataResponse struct {
	Values               []*ValueWithMetadata `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetHistoricalDataResponse) Reset()         { *m = GetHistoricalDataResponse{} }
func (m *GetHistoricalDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetHistoricalDataResponse) ProtoMessage()    {}
func (*GetHistoricalDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{11}
}

func (m *GetHistoricalDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoricalDataResponse.Unmarshal(m, b)
}
func (m *GetHistoricalDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoricalDataResponse.Marshal(b, m, deterministic)
}
func (m *GetHistoricalDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoricalDataResponse.Merge(m, src)
}
func (m *GetHistoricalDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetHistoricalDataResponse.Size(m)
}
func (m *GetHistoricalDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoricalDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoricalDataResponse proto.InternalMessageInfo

func (m *GetHistoricalDataResponse) GetValues() []*ValueWithMetadata {
	if m != nil {
		return m.Values
	}
	return nil
}

type GetDataReadersResponse struct {
	ReadBy               map[string]uint32 `protobuf:"bytes,1,rep,name=read_by,json=readBy,proto3" json:"read_by,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetDataReadersResponse) Reset()         { *m = GetDataReadersResponse{} }
func (m *GetDataReadersResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataReadersResponse) ProtoMessage()    {}
func (*GetDataReadersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{12}
}

func (m *GetDataReadersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataReadersResponse.Unmarshal(m, b)
}
func (m *GetDataReadersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataReadersResponse.Marshal(b, m, deterministic)
}
func (m *GetDataReadersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataReadersResponse.Merge(m, src)
}
func (m *GetDataReadersResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataReadersResponse.Size(m)
}
func (m *GetDataReadersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataReadersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataReadersResponse proto.InternalMessageInfo

func (m *GetDataReadersResponse) GetReadBy() map[string]uint32 {
	if m != nil {
		return m.ReadBy
	}
	return nil
}

type GetDataWritersResponse struct {
	WrittenBy            map[string]uint32 `protobuf:"bytes,1,rep,name=written_by,json=writtenBy,proto3" json:"written_by,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetDataWritersResponse) Reset()         { *m = GetDataWritersResponse{} }
func (m *GetDataWritersResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataWritersResponse) ProtoMessage()    {}
func (*GetDataWritersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{13}
}

func (m *GetDataWritersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataWritersResponse.Unmarshal(m, b)
}
func (m *GetDataWritersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataWritersResponse.Marshal(b, m, deterministic)
}
func (m *GetDataWritersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataWritersResponse.Merge(m, src)
}
func (m *GetDataWritersResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataWritersResponse.Size(m)
}
func (m *GetDataWritersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataWritersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataWritersResponse proto.InternalMessageInfo

func (m *GetDataWritersResponse) GetWrittenBy() map[string]uint32 {
	if m != nil {
		return m.WrittenBy
	}
	return nil
}

type GetDataProvenanceResponse struct {
	KVs                  []*KVWithMetadata `protobuf:"bytes,1,rep,name=KVs,proto3" json:"KVs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetDataProvenanceResponse) Reset()         { *m = GetDataProvenanceResponse{} }
func (m *GetDataProvenanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataProvenanceResponse) ProtoMessage()    {}
func (*GetDataProvenanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{14}
}

func (m *GetDataProvenanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataProvenanceResponse.Unmarshal(m, b)
}
func (m *GetDataProvenanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataProvenanceResponse.Marshal(b, m, deterministic)
}
func (m *GetDataProvenanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataProvenanceResponse.Merge(m, src)
}
func (m *GetDataProvenanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataProvenanceResponse.Size(m)
}
func (m *GetDataProvenanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataProvenanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataProvenanceResponse proto.InternalMessageInfo

func (m *GetDataProvenanceResponse) GetKVs() []*KVWithMetadata {
	if m != nil {
		return m.KVs
	}
	return nil
}

type GetTxIDsSubmittedByResponse struct {
	TxIDs                []string `protobuf:"bytes,1,rep,name=txIDs,proto3" json:"txIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxIDsSubmittedByResponse) Reset()         { *m = GetTxIDsSubmittedByResponse{} }
func (m *GetTxIDsSubmittedByResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxIDsSubmittedByResponse) ProtoMessage()    {}
func (*GetTxIDsSubmittedByResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{15}
}

func (m *GetTxIDsSubmittedByResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxIDsSubmittedByResponse.Unmarshal(m, b)
}
func (m *GetTxIDsSubmittedByResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxIDsSubmittedByResponse.Marshal(b, m, deterministic)
}
func (m *GetTxIDsSubmittedByResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxIDsSubmittedByResponse.Merge(m, src)
}
func (m *GetTxIDsSubmittedByResponse) XXX_Size() int {
	return xxx_messageInfo_GetTxIDsSubmittedByResponse.Size(m)
}
func (m *GetTxIDsSubmittedByResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxIDsSubmittedByResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxIDsSubmittedByResponse proto.InternalMessageInfo

func (m *GetTxIDsSubmittedByResponse) GetTxIDs() []string {
	if m != nil {
		return m.TxIDs
	}
	return nil
}

type TxResponse struct {
	Receipt              *TxReceipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TxResponse) Reset()         { *m = TxResponse{} }
func (m *TxResponse) String() string { return proto.CompactTextString(m) }
func (*TxResponse) ProtoMessage()    {}
func (*TxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fbc901015fa5021, []int{16}
}

func (m *TxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxResponse.Unmarshal(m, b)
}
func (m *TxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxResponse.Marshal(b, m, deterministic)
}
func (m *TxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponse.Merge(m, src)
}
func (m *TxResponse) XXX_Size() int {
	return xxx_messageInfo_TxResponse.Size(m)
}
func (m *TxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponse proto.InternalMessageInfo

func (m *TxResponse) GetReceipt() *TxReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseEnvelope)(nil), "types.ResponseEnvelope")
	proto.RegisterType((*Payload)(nil), "types.Payload")
	proto.RegisterType((*ResponseHeader)(nil), "types.ResponseHeader")
	proto.RegisterType((*GetDBStatusResponse)(nil), "types.GetDBStatusResponse")
	proto.RegisterType((*GetDataResponse)(nil), "types.GetDataResponse")
	proto.RegisterType((*GetUserResponse)(nil), "types.GetUserResponse")
	proto.RegisterType((*GetConfigResponse)(nil), "types.GetConfigResponse")
	proto.RegisterType((*GetNodeConfigResponse)(nil), "types.GetNodeConfigResponse")
	proto.RegisterType((*GetBlockResponse)(nil), "types.GetBlockResponse")
	proto.RegisterType((*GetLedgerPathResponse)(nil), "types.GetLedgerPathResponse")
	proto.RegisterType((*GetTxProofResponse)(nil), "types.GetTxProofResponse")
	proto.RegisterType((*GetHistoricalDataResponse)(nil), "types.GetHistoricalDataResponse")
	proto.RegisterType((*GetDataReadersResponse)(nil), "types.GetDataReadersResponse")
	proto.RegisterMapType((map[string]uint32)(nil), "types.GetDataReadersResponse.ReadByEntry")
	proto.RegisterType((*GetDataWritersResponse)(nil), "types.GetDataWritersResponse")
	proto.RegisterMapType((map[string]uint32)(nil), "types.GetDataWritersResponse.WrittenByEntry")
	proto.RegisterType((*GetDataProvenanceResponse)(nil), "types.GetDataProvenanceResponse")
	proto.RegisterType((*GetTxIDsSubmittedByResponse)(nil), "types.GetTxIDsSubmittedByResponse")
	proto.RegisterType((*TxResponse)(nil), "types.TxResponse")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_0fbc901015fa5021) }

var fileDescriptor_0fbc901015fa5021 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x4f, 0x1a, 0x41,
	0x14, 0x0d, 0x5a, 0x51, 0x2e, 0xa8, 0xb8, 0x7e, 0x84, 0x62, 0x93, 0x9a, 0x7d, 0xa9, 0xad, 0x74,
	0x69, 0x34, 0x4d, 0x6d, 0xd3, 0x27, 0xc4, 0xa0, 0xa5, 0x36, 0x64, 0xa5, 0x9a, 0xf4, 0x85, 0xcc,
	0xee, 0x5e, 0x61, 0x23, 0xcc, 0x90, 0x99, 0xbb, 0x28, 0xbf, 0xa4, 0x3f, 0xa0, 0x7f, 0xb4, 0xd9,
	0xd9, 0xd9, 0x05, 0x4c, 0xfb, 0xe0, 0xdb, 0xde, 0xb9, 0x67, 0xce, 0x39, 0x33, 0xf7, 0xec, 0xc0,
	0x86, 0x44, 0x35, 0x16, 0x5c, 0xa1, 0x33, 0x96, 0x82, 0x84, 0xb5, 0x42, 0xd3, 0x31, 0xaa, 0xea,
	0xb6, 0x2f, 0xf8, 0x5d, 0xd8, 0x8f, 0x24, 0xa3, 0x50, 0xf0, 0xa4, 0x57, 0xdd, 0xf7, 0x86, 0xc2,
	0xbf, 0xef, 0x31, 0x1e, 0xf4, 0x48, 0x32, 0xae, 0x98, 0x3f, 0x6b, 0xda, 0xdf, 0xa0, 0xec, 0x1a,
	0xaa, 0x73, 0x3e, 0xc1, 0xa1, 0x18, 0xa3, 0x55, 0x81, 0xd5, 0x31, 0x9b, 0x0e, 0x05, 0x0b, 0x2a,
	0xb9, 0x83, 0xdc, 0x61, 0xc9, 0x4d, 0x4b, 0xeb, 0x15, 0x14, 0x54, 0xd8, 0xe7, 0x8c, 0x22, 0x89,
	0x95, 0x25, 0xdd, 0x9b, 0x2d, 0xd8, 0x5d, 0x58, 0xed, 0x18, 0xe0, 0x7b, 0xc8, 0x0f, 0x90, 0x05,
	0x28, 0x35, 0x43, 0xf1, 0x78, 0xd7, 0xd1, 0x06, 0x9d, 0x54, 0xeb, 0x42, 0x37, 0x5d, 0x03, 0xb2,
	0xaa, 0xb0, 0x96, 0x1e, 0xc8, 0xd0, 0x66, 0xb5, 0x7d, 0x08, 0x1b, 0x8b, 0xbb, 0xac, 0x3d, 0xc8,
	0x73, 0x11, 0xe0, 0x65, 0x53, 0x93, 0x17, 0x5c, 0x53, 0xd9, 0x47, 0xb0, 0xdd, 0x42, 0x6a, 0x36,
	0xae, 0x89, 0x51, 0xa4, 0xd2, 0x4d, 0xd6, 0x0e, 0xac, 0xe0, 0x63, 0xa8, 0x48, 0xa3, 0xd7, 0xdc,
	0xa4, 0xb0, 0xbb, 0xb0, 0x19, 0x83, 0x19, 0xb1, 0x79, 0xe0, 0x84, 0x0d, 0x23, 0x34, 0xa7, 0x4e,
	0x0a, 0xeb, 0x08, 0xd6, 0x46, 0x48, 0x2c, 0x60, 0xc4, 0xb4, 0xb7, 0xe2, 0xf1, 0xa6, 0x39, 0xcc,
	0x95, 0x59, 0x76, 0x33, 0x80, 0xdd, 0xd3, 0xac, 0x3f, 0x15, 0xca, 0x8c, 0xf5, 0x35, 0xbc, 0x88,
	0x54, 0x76, 0x11, 0x45, 0xb3, 0x57, 0x43, 0x74, 0xe3, 0x79, 0x02, 0x1c, 0xb6, 0x5a, 0x48, 0x67,
	0x7a, 0xcc, 0x99, 0x44, 0x0d, 0xf2, 0xc9, 0xe0, 0x8d, 0xc8, 0x8e, 0xd9, 0x7f, 0x36, 0x8c, 0x14,
	0xa1, 0x34, 0x68, 0x83, 0x79, 0x9e, 0x5e, 0x1b, 0x76, 0x5b, 0x48, 0x3f, 0x44, 0x80, 0x4f, 0x34,
	0x8f, 0xa1, 0x18, 0x5f, 0x7b, 0x6f, 0x41, 0x78, 0xcb, 0x10, 0xcd, 0xe1, 0x81, 0x67, 0xdf, 0xf6,
	0x25, 0x94, 0x5b, 0x48, 0x8d, 0x38, 0x8e, 0x19, 0xcf, 0x47, 0x28, 0x25, 0xf9, 0x5c, 0xc8, 0x8b,
	0x65, 0x88, 0x34, 0xd6, 0x84, 0xa5, 0xe8, 0xcd, 0x0a, 0xbb, 0xa3, 0x7d, 0x7d, 0xc7, 0xa0, 0x8f,
	0xb2, 0xc3, 0x68, 0x90, 0xf1, 0x7d, 0x82, 0xf5, 0x79, 0x3e, 0x55, 0xc9, 0x1d, 0x2c, 0xff, 0x87,
	0xb0, 0x34, 0x47, 0xa8, 0xec, 0x1a, 0x58, 0x2d, 0xa4, 0xee, 0x63, 0x47, 0x0a, 0x71, 0x97, 0xd1,
	0xed, 0x41, 0x7e, 0xc0, 0xd4, 0x00, 0x13, 0x9e, 0x92, 0x6b, 0x2a, 0xfb, 0x0a, 0x5e, 0xb6, 0x90,
	0x2e, 0x42, 0x45, 0x42, 0x86, 0x3e, 0x1b, 0x2e, 0x04, 0xe9, 0x03, 0xe4, 0x75, 0x76, 0x52, 0xf1,
	0x8a, 0x11, 0xbf, 0x89, 0x17, 0x6f, 0x43, 0x1a, 0x64, 0x17, 0x6d, 0x70, 0xf6, 0xef, 0x1c, 0xec,
	0x65, 0x71, 0xd4, 0x7e, 0x32, 0xb2, 0x06, 0xac, 0x4a, 0x64, 0x41, 0xcf, 0x9b, 0x1a, 0xb6, 0xb7,
	0x86, 0xed, 0xdf, 0x78, 0x27, 0xae, 0x1b, 0xd3, 0x73, 0x4e, 0x72, 0xea, 0xe6, 0xa5, 0x2e, 0xaa,
	0x9f, 0xa1, 0x38, 0xb7, 0x6c, 0x95, 0x61, 0xf9, 0x1e, 0xa7, 0xe6, 0xef, 0x89, 0x3f, 0x67, 0xd1,
	0x8f, 0x03, 0xb1, 0x6e, 0xa2, 0xff, 0x65, 0xe9, 0x34, 0x67, 0xff, 0x99, 0x39, 0xbb, 0x95, 0x21,
	0xcd, 0x3b, 0x6b, 0x03, 0x3c, 0xc8, 0x90, 0x08, 0xf9, 0xcc, 0x5c, 0x6d, 0xd1, 0xdc, 0x93, 0x2d,
	0xce, 0x6d, 0x82, 0x4f, 0xfd, 0x15, 0x1e, 0xd2, 0xba, 0xfa, 0x15, 0x36, 0x16, 0x9b, 0xcf, 0x72,
	0xd9, 0xd4, 0xe3, 0x88, 0x15, 0x3b, 0x52, 0x4c, 0x90, 0x33, 0xee, 0x63, 0xe6, 0xf3, 0x0d, 0x2c,
	0xb7, 0x6f, 0xd2, 0x59, 0xa4, 0x2f, 0x51, 0xfb, 0x66, 0x61, 0x10, 0x31, 0xc2, 0x3e, 0x81, 0x7d,
	0x1d, 0x81, 0xcb, 0xa6, 0xba, 0x8e, 0xbc, 0x51, 0xec, 0x26, 0x68, 0x4c, 0xe7, 0xdf, 0x07, 0x8a,
	0x7b, 0x9a, 0xa9, 0xe0, 0x26, 0x85, 0x7d, 0x0a, 0xd0, 0x7d, 0xcc, 0x30, 0xef, 0xe2, 0x69, 0xf9,
	0x18, 0x8e, 0xc9, 0x24, 0xb9, 0x6c, 0xf4, 0x62, 0x8c, 0x5e, 0x77, 0x53, 0x40, 0xc3, 0xf9, 0x55,
	0xeb, 0x87, 0x34, 0x88, 0x3c, 0x27, 0xf4, 0x46, 0x8e, 0x2f, 0x46, 0x75, 0x1d, 0x48, 0x7f, 0xc0,
	0x42, 0x1e, 0x78, 0x75, 0x85, 0x72, 0x82, 0xb2, 0x3e, 0xbe, 0xef, 0xd7, 0x35, 0x85, 0x97, 0xd7,
	0x4f, 0xf6, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x64, 0x08, 0xe4, 0xfd, 0x05, 0x00,
	0x00,
}