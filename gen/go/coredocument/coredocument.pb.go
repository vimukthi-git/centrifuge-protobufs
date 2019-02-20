// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coredocument/coredocument.proto

package coredocumentpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/centrifuge/precise-proofs/proofs/proto"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Action defines the set of actions a collaborator can/have per document.
type Action int32

const (
	Action_ACTION_INVALID Action = 0
	// read represents just reading the doc/fields
	Action_ACTION_READ Action = 1
	// read_sign represents reading as well the sign the documents. We will pick this one when requesting the signatures.
	Action_ACTION_READ_SIGN Action = 2
)

var Action_name = map[int32]string{
	0: "ACTION_INVALID",
	1: "ACTION_READ",
	2: "ACTION_READ_SIGN",
}
var Action_value = map[string]int32{
	"ACTION_INVALID":   0,
	"ACTION_READ":      1,
	"ACTION_READ_SIGN": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{0}
}

// `CoreDocument` is a document that can be sent to different nodes and anchored
// on chain. It handles all the generic features native Centrifuge Documents support:
//
// * Merkle Roots for the document data
// * Signatures on document data
// * Access Control
type CoreDocument struct {
	// # Identifiers
	// CoreDocument has two kinds of identifiers, the `document_identifier` is assigned
	// once per document and stays the same for the lifetime of the document.
	// document_identifier is the first ID ever used to anchor the document on chain and
	// is used internally to store all future versions. The `previous_version`, `current_version`, and the
	// `next_version` refer only to a particular version.
	//
	// 32 byte value
	DocumentIdentifier []byte `protobuf:"bytes,9,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// previous_version refers to the previous state of the document.
	// 32 byte value
	PreviousVersion []byte `protobuf:"bytes,16,opt,name=previous_version,json=previousVersion,proto3" json:"previous_version,omitempty"`
	// current_version is the version used to refer to the current state of the document.
	// 32 byte value
	CurrentVersion []byte `protobuf:"bytes,3,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	// next_version is the version that is going to be used for the next version if any
	// party wants to update the state.
	NextVersion []byte `protobuf:"bytes,4,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	// # Merkle roots
	// document_root the root of all roots. It's signature_root along with a list of all signatures
	DocumentRoot []byte `protobuf:"bytes,7,opt,name=document_root,json=documentRoot,proto3" json:"document_root,omitempty"`
	// signing_root is the Merkle root calculated from all fields on the document that need
	// to be signed. This is all fields except for the signatures themselves and the `document_root`.
	SigningRoot []byte `protobuf:"bytes,10,opt,name=signing_root,json=signingRoot,proto3" json:"signing_root,omitempty"`
	// previous_root is the `document_root` of the previous version of the document
	PreviousRoot []byte `protobuf:"bytes,2,opt,name=previous_root,json=previousRoot,proto3" json:"previous_root,omitempty"`
	// data_root is the Merkle root calculated for the `embedded_data` & `embedded_salts` document (such as an invoice).
	DataRoot []byte `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	// Signatures
	// Signatures of the signature_root by centrifuge identities. This array should be sorted by the Centrifuge ID.
	Signatures []*Signature `protobuf:"bytes,6,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// Document a serialized document
	EmbeddedData      *any.Any        `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData,proto3" json:"embedded_data,omitempty"`
	EmbeddedDataSalts []*DocumentSalt `protobuf:"bytes,14,rep,name=embedded_data_salts,json=embeddedDataSalts,proto3" json:"embedded_data_salts,omitempty"`
	// CoreDocumentSalts is inlined
	CoredocumentSalts []*DocumentSalt `protobuf:"bytes,15,rep,name=coredocument_salts,json=coredocumentSalts,proto3" json:"coredocument_salts,omitempty"`
	// Collaborators involved in the document consensus
	Collaborators [][]byte `protobuf:"bytes,17,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	// list of roles
	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// read rules
	ReadRules []*ReadRule `protobuf:"bytes,19,rep,name=read_rules,json=readRules,proto3" json:"read_rules,omitempty"`
	// nft list for uniqueness check
	Nfts                 []*NFT   `protobuf:"bytes,20,rep,name=nfts,proto3" json:"nfts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoreDocument) Reset()         { *m = CoreDocument{} }
func (m *CoreDocument) String() string { return proto.CompactTextString(m) }
func (*CoreDocument) ProtoMessage()    {}
func (*CoreDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{0}
}
func (m *CoreDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreDocument.Unmarshal(m, b)
}
func (m *CoreDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreDocument.Marshal(b, m, deterministic)
}
func (dst *CoreDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreDocument.Merge(dst, src)
}
func (m *CoreDocument) XXX_Size() int {
	return xxx_messageInfo_CoreDocument.Size(m)
}
func (m *CoreDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreDocument.DiscardUnknown(m)
}

var xxx_messageInfo_CoreDocument proto.InternalMessageInfo

func (m *CoreDocument) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CoreDocument) GetPreviousVersion() []byte {
	if m != nil {
		return m.PreviousVersion
	}
	return nil
}

func (m *CoreDocument) GetCurrentVersion() []byte {
	if m != nil {
		return m.CurrentVersion
	}
	return nil
}

func (m *CoreDocument) GetNextVersion() []byte {
	if m != nil {
		return m.NextVersion
	}
	return nil
}

func (m *CoreDocument) GetDocumentRoot() []byte {
	if m != nil {
		return m.DocumentRoot
	}
	return nil
}

func (m *CoreDocument) GetSigningRoot() []byte {
	if m != nil {
		return m.SigningRoot
	}
	return nil
}

func (m *CoreDocument) GetPreviousRoot() []byte {
	if m != nil {
		return m.PreviousRoot
	}
	return nil
}

func (m *CoreDocument) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

func (m *CoreDocument) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *CoreDocument) GetEmbeddedData() *any.Any {
	if m != nil {
		return m.EmbeddedData
	}
	return nil
}

func (m *CoreDocument) GetEmbeddedDataSalts() []*DocumentSalt {
	if m != nil {
		return m.EmbeddedDataSalts
	}
	return nil
}

func (m *CoreDocument) GetCoredocumentSalts() []*DocumentSalt {
	if m != nil {
		return m.CoredocumentSalts
	}
	return nil
}

func (m *CoreDocument) GetCollaborators() [][]byte {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *CoreDocument) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *CoreDocument) GetReadRules() []*ReadRule {
	if m != nil {
		return m.ReadRules
	}
	return nil
}

func (m *CoreDocument) GetNfts() []*NFT {
	if m != nil {
		return m.Nfts
	}
	return nil
}

type DocumentSalt struct {
	Compact              []byte   `protobuf:"bytes,1,opt,name=compact,proto3" json:"compact,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentSalt) Reset()         { *m = DocumentSalt{} }
func (m *DocumentSalt) String() string { return proto.CompactTextString(m) }
func (*DocumentSalt) ProtoMessage()    {}
func (*DocumentSalt) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{1}
}
func (m *DocumentSalt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentSalt.Unmarshal(m, b)
}
func (m *DocumentSalt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentSalt.Marshal(b, m, deterministic)
}
func (dst *DocumentSalt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentSalt.Merge(dst, src)
}
func (m *DocumentSalt) XXX_Size() int {
	return xxx_messageInfo_DocumentSalt.Size(m)
}
func (m *DocumentSalt) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentSalt.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentSalt proto.InternalMessageInfo

func (m *DocumentSalt) GetCompact() []byte {
	if m != nil {
		return m.Compact
	}
	return nil
}

func (m *DocumentSalt) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type AccessToken struct {
	// The identifier is an internal 256bit word
	Identifier []byte `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// The identity granting access to the document
	Granter []byte `protobuf:"bytes,3,opt,name=granter,proto3" json:"granter,omitempty"`
	// The identity being granted access to the document
	Grantee []byte `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Role identifier is the identifier on the read rule that this token should be mapped to
	RoleIdentifier []byte `protobuf:"bytes,5,opt,name=role_identifier,json=roleIdentifier,proto3" json:"role_identifier,omitempty"`
	// Original identifier of the document
	DocumentIdentifier []byte `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// Cryptographic signature that an access token is valid
	Signature []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	// The public key of the signed message
	Key                  []byte   `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{2}
}
func (m *AccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessToken.Unmarshal(m, b)
}
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessToken.Marshal(b, m, deterministic)
}
func (dst *AccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessToken.Merge(dst, src)
}
func (m *AccessToken) XXX_Size() int {
	return xxx_messageInfo_AccessToken.Size(m)
}
func (m *AccessToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessToken.DiscardUnknown(m)
}

var xxx_messageInfo_AccessToken proto.InternalMessageInfo

func (m *AccessToken) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *AccessToken) GetGranter() []byte {
	if m != nil {
		return m.Granter
	}
	return nil
}

func (m *AccessToken) GetGrantee() []byte {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *AccessToken) GetRoleIdentifier() []byte {
	if m != nil {
		return m.RoleIdentifier
	}
	return nil
}

func (m *AccessToken) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *AccessToken) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AccessToken) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// Signature contains the entity ID, public key used and signature)
type Signature struct {
	// `entity_id` is the CentrifugeID of the entity signing the document.
	EntityId []byte `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// `public_key` is the public key of the `entity` used for signing the `CoreDocument`
	PublicKey            []byte               `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature            []byte               `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{3}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetEntityId() []byte {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *Signature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// Roles holds a list of collaborators, NFTs, and/or access tokens.
type Role struct {
	// Index of the position of the Role within the coredocument
	RoleKey []byte `protobuf:"bytes,1,opt,name=role_key,json=roleKey,proto3" json:"role_key,omitempty"`
	// collaborators holds the list of document collaborators
	Collaborators [][]byte `protobuf:"bytes,3,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	// nfts is a list of registry address/tokenID pairs.
	// For easier verification in merkle proofs, the values are simply concatenated with the first 22 bytes being the NFT registry and the remaining 32 bytes the tokenID.
	Nfts [][]byte `protobuf:"bytes,4,rep,name=nfts,proto3" json:"nfts,omitempty"`
	// AccessTokens which have been added to this CoreDocument
	AccessTokens         []*AccessToken `protobuf:"bytes,5,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{4}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetRoleKey() []byte {
	if m != nil {
		return m.RoleKey
	}
	return nil
}

func (m *Role) GetCollaborators() [][]byte {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *Role) GetNfts() [][]byte {
	if m != nil {
		return m.Nfts
	}
	return nil
}

func (m *Role) GetAccessTokens() []*AccessToken {
	if m != nil {
		return m.AccessTokens
	}
	return nil
}

type ReadRule struct {
	Roles                [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Action               Action   `protobuf:"varint,4,opt,name=action,proto3,enum=coredocument.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRule) Reset()         { *m = ReadRule{} }
func (m *ReadRule) String() string { return proto.CompactTextString(m) }
func (*ReadRule) ProtoMessage()    {}
func (*ReadRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{5}
}
func (m *ReadRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRule.Unmarshal(m, b)
}
func (m *ReadRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRule.Marshal(b, m, deterministic)
}
func (dst *ReadRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRule.Merge(dst, src)
}
func (m *ReadRule) XXX_Size() int {
	return xxx_messageInfo_ReadRule.Size(m)
}
func (m *ReadRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRule.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRule proto.InternalMessageInfo

func (m *ReadRule) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *ReadRule) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_ACTION_INVALID
}

type NFT struct {
	RegistryId           []byte   `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	TokenId              []byte   `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NFT) Reset()         { *m = NFT{} }
func (m *NFT) String() string { return proto.CompactTextString(m) }
func (*NFT) ProtoMessage()    {}
func (*NFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_bfdc360a8c395a0b, []int{6}
}
func (m *NFT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NFT.Unmarshal(m, b)
}
func (m *NFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NFT.Marshal(b, m, deterministic)
}
func (dst *NFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFT.Merge(dst, src)
}
func (m *NFT) XXX_Size() int {
	return xxx_messageInfo_NFT.Size(m)
}
func (m *NFT) XXX_DiscardUnknown() {
	xxx_messageInfo_NFT.DiscardUnknown(m)
}

var xxx_messageInfo_NFT proto.InternalMessageInfo

func (m *NFT) GetRegistryId() []byte {
	if m != nil {
		return m.RegistryId
	}
	return nil
}

func (m *NFT) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func init() {
	proto.RegisterType((*CoreDocument)(nil), "coredocument.CoreDocument")
	proto.RegisterType((*DocumentSalt)(nil), "coredocument.DocumentSalt")
	proto.RegisterType((*AccessToken)(nil), "coredocument.AccessToken")
	proto.RegisterType((*Signature)(nil), "coredocument.Signature")
	proto.RegisterType((*Role)(nil), "coredocument.Role")
	proto.RegisterType((*ReadRule)(nil), "coredocument.ReadRule")
	proto.RegisterType((*NFT)(nil), "coredocument.NFT")
	proto.RegisterEnum("coredocument.Action", Action_name, Action_value)
}

func init() {
	proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_coredocument_bfdc360a8c395a0b)
}

var fileDescriptor_coredocument_bfdc360a8c395a0b = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xeb, 0x44,
	0x14, 0xc5, 0x4d, 0x93, 0x36, 0x37, 0x4e, 0x9a, 0x4e, 0x03, 0xb8, 0xe5, 0xa3, 0x21, 0x42, 0x7a,
	0xa1, 0x82, 0x44, 0x2a, 0x42, 0xb0, 0x40, 0x4f, 0x24, 0x2d, 0x0f, 0x45, 0xa0, 0x50, 0xdc, 0xe8,
	0x2d, 0xd8, 0x58, 0x13, 0x7b, 0x1a, 0x59, 0x75, 0x3c, 0xd6, 0xcc, 0xb8, 0x22, 0xbf, 0x06, 0x96,
	0x2c, 0x61, 0xcb, 0xae, 0xff, 0x83, 0x7f, 0xc0, 0xba, 0x7b, 0x34, 0x33, 0x1e, 0x67, 0x1c, 0xa8,
	0xf4, 0x56, 0xc9, 0x3d, 0xf7, 0xdc, 0x73, 0xc7, 0x73, 0xcf, 0x1d, 0x38, 0x0f, 0x29, 0x23, 0x11,
	0x0d, 0xf3, 0x35, 0x49, 0xc5, 0xd8, 0x0e, 0x46, 0x19, 0xa3, 0x82, 0x22, 0xd7, 0xc6, 0xce, 0x4e,
	0x57, 0x94, 0xae, 0x12, 0x32, 0x56, 0xb9, 0x65, 0x7e, 0x37, 0xc6, 0xe9, 0x46, 0x13, 0xcf, 0xce,
	0x77, 0x53, 0x22, 0x5e, 0x13, 0x2e, 0xf0, 0x3a, 0x2b, 0x08, 0x2f, 0x32, 0x46, 0xc2, 0x98, 0x93,
	0xcf, 0x32, 0x46, 0xe9, 0x1d, 0x1f, 0x6f, 0x7f, 0x04, 0xd5, 0x81, 0x26, 0x0e, 0xfe, 0x6e, 0x80,
	0x7b, 0x45, 0x19, 0xb9, 0x2e, 0xba, 0xa2, 0x31, 0x9c, 0x98, 0x13, 0x04, 0x71, 0x44, 0x52, 0x11,
	0xdf, 0xc5, 0x84, 0x79, 0xcd, 0xbe, 0x33, 0x74, 0x7d, 0x64, 0x52, 0xb3, 0x32, 0x83, 0x3e, 0x81,
	0x6e, 0xc6, 0xc8, 0x43, 0x4c, 0x73, 0x1e, 0x3c, 0x10, 0xc6, 0x63, 0x9a, 0x7a, 0x5d, 0xc5, 0x3e,
	0x32, 0xf8, 0x6b, 0x0d, 0xa3, 0x17, 0x70, 0x14, 0xe6, 0x8c, 0x49, 0x69, 0xc3, 0xac, 0x29, 0x66,
	0xa7, 0x80, 0x0d, 0xf1, 0x23, 0x70, 0x53, 0xf2, 0xcb, 0x96, 0xb5, 0xaf, 0x58, 0x2d, 0x89, 0x19,
	0xca, 0x05, 0xb4, 0xcb, 0x73, 0x32, 0x4a, 0x85, 0x77, 0x20, 0x39, 0xd3, 0xfa, 0x6f, 0x8f, 0x4f,
	0xe0, 0xf8, 0xae, 0xc9, 0xf9, 0x94, 0x0a, 0x34, 0x04, 0x97, 0xc7, 0xab, 0x34, 0x4e, 0x57, 0x9a,
	0x0a, 0x36, 0xb5, 0x55, 0xa4, 0x14, 0xf3, 0x02, 0xda, 0xe5, 0xc7, 0x28, 0xea, 0x9e, 0xa6, 0xfe,
	0xae, 0x55, 0x4d, 0x4e, 0x71, 0x07, 0xd0, 0x8c, 0xb0, 0xc0, 0x9a, 0x57, 0xb7, 0x25, 0x0f, 0x25,
	0xae, 0x38, 0x2f, 0x01, 0xa4, 0x3c, 0x16, 0x39, 0x23, 0xdc, 0x6b, 0xf4, 0x6b, 0xc3, 0xd6, 0xe5,
	0xbb, 0xa3, 0xca, 0xe8, 0x6f, 0x4d, 0xde, 0x54, 0x5b, 0x15, 0xe8, 0x1b, 0x68, 0x93, 0xf5, 0x92,
	0x44, 0x11, 0x89, 0x02, 0x29, 0xea, 0xb5, 0xfb, 0xce, 0xb0, 0x75, 0xd9, 0x1b, 0x69, 0x03, 0x8c,
	0x8c, 0x01, 0x46, 0x93, 0x74, 0x53, 0x7e, 0xbb, 0xa9, 0xb8, 0xc6, 0x02, 0xa3, 0x9f, 0xe0, 0xa4,
	0xa2, 0x10, 0x70, 0x9c, 0x08, 0xee, 0x75, 0xd4, 0x51, 0xce, 0xaa, 0x47, 0x31, 0x26, 0xb8, 0xc5,
	0x89, 0x30, 0x6a, 0xc7, 0xb6, 0x9a, 0x4c, 0x70, 0x74, 0x03, 0xc8, 0x2e, 0x2b, 0x14, 0x8f, 0xde,
	0x58, 0xd1, 0x66, 0x68, 0xc5, 0x8f, 0xa1, 0x1d, 0xd2, 0x24, 0xc1, 0x4b, 0xca, 0xb0, 0xa0, 0x8c,
	0x7b, 0xc7, 0xfd, 0xda, 0xd0, 0xf5, 0xab, 0x20, 0xfa, 0x1a, 0xea, 0x8c, 0x26, 0x84, 0x7b, 0x8e,
	0x6a, 0x85, 0xaa, 0xad, 0x7c, 0x9a, 0x90, 0x29, 0xfa, 0xe3, 0xf1, 0x09, 0xfa, 0x7f, 0x3d, 0x3e,
	0xc1, 0xa1, 0xa4, 0x06, 0xf7, 0x64, 0xe3, 0xeb, 0x22, 0xf4, 0x05, 0x00, 0x23, 0x38, 0x0a, 0x58,
	0x2e, 0x25, 0x4e, 0x94, 0xc4, 0x3b, 0x3b, 0x12, 0x04, 0x47, 0x7e, 0x9e, 0x10, 0xbf, 0xc9, 0x8a,
	0x7f, 0xb2, 0xe9, 0x7e, 0x7a, 0x27, 0xb8, 0xd7, 0x53, 0x05, 0xc7, 0xd5, 0x82, 0xf9, 0xab, 0xc5,
	0xf4, 0xed, 0xb2, 0x65, 0x8b, 0x91, 0x55, 0xcc, 0x05, 0xdb, 0x04, 0x71, 0xe4, 0xab, 0xaa, 0xc1,
	0x4b, 0x70, 0xed, 0x2b, 0x40, 0x1e, 0x1c, 0x84, 0x74, 0x9d, 0xe1, 0x50, 0x78, 0x8e, 0xf2, 0xb4,
	0x09, 0x51, 0x0f, 0xea, 0x0f, 0x38, 0xc9, 0x89, 0x76, 0x9c, 0xaf, 0x83, 0xc1, 0x3f, 0x0e, 0xb4,
	0x26, 0x61, 0x48, 0x38, 0x5f, 0xd0, 0x7b, 0x92, 0xa2, 0x0f, 0x01, 0xac, 0xa5, 0xd4, 0x12, 0x16,
	0x22, 0xf5, 0x57, 0x0c, 0xa7, 0x82, 0xb0, 0x62, 0xb3, 0x4c, 0xb8, 0xcd, 0x90, 0x62, 0x9b, 0x4c,
	0x28, 0xb7, 0x52, 0xdd, 0x95, 0x25, 0x5c, 0xd7, 0x5b, 0x29, 0x61, 0x6b, 0xd3, 0x9f, 0x79, 0x1a,
	0xf6, 0x9e, 0x7d, 0x1a, 0xde, 0x87, 0x66, 0xe9, 0x65, 0xaf, 0xa1, 0x68, 0x5b, 0x00, 0x75, 0xa1,
	0x76, 0x4f, 0x36, 0x7a, 0x6f, 0x7d, 0xf9, 0x77, 0xf0, 0xab, 0x03, 0xcd, 0x72, 0x1d, 0xd0, 0x7b,
	0xd0, 0x94, 0x4a, 0x42, 0x5e, 0x67, 0xf1, 0xa9, 0x87, 0x1a, 0x98, 0x45, 0xe8, 0x03, 0x80, 0x2c,
	0x5f, 0x26, 0x71, 0x28, 0x47, 0x5c, 0x1c, 0xa1, 0xa9, 0x91, 0xef, 0xc9, 0xa6, 0xda, 0xb9, 0xb6,
	0xdb, 0xf9, 0x2b, 0x68, 0x96, 0x0f, 0xa6, 0xba, 0x0d, 0xe9, 0xdb, 0xdd, 0x8d, 0x5a, 0x18, 0x86,
	0xbf, 0x25, 0x0f, 0xfe, 0x74, 0x60, 0x5f, 0x1a, 0x0d, 0x9d, 0x6e, 0x0d, 0x66, 0x26, 0x29, 0x63,
	0xd9, 0xfb, 0x3f, 0x66, 0xae, 0xfd, 0x9f, 0x99, 0x51, 0xe1, 0xab, 0x7d, 0x95, 0x54, 0xff, 0xd1,
	0x02, 0xda, 0x58, 0x0d, 0x3b, 0x10, 0x72, 0xda, 0xdc, 0xab, 0x2b, 0xd3, 0x9d, 0x56, 0x4d, 0x67,
	0xf9, 0x61, 0xda, 0x2b, 0xcd, 0x67, 0xb9, 0xc0, 0x77, 0xf1, 0x96, 0xc2, 0x07, 0x73, 0x38, 0x34,
	0xc6, 0x96, 0x2e, 0xd3, 0x2b, 0xb4, 0xa7, 0xda, 0x16, 0xab, 0xf1, 0x29, 0x34, 0x70, 0x28, 0xcc,
	0x43, 0xdb, 0xb9, 0xec, 0xed, 0x36, 0x94, 0x39, 0xbf, 0xe0, 0x0c, 0x26, 0x50, 0x9b, 0xbf, 0x5a,
	0xa0, 0xf3, 0x8a, 0xdf, 0x8d, 0x17, 0x0d, 0x34, 0x8b, 0xe4, 0x15, 0xa9, 0xcf, 0x90, 0x59, 0x3d,
	0xa0, 0x03, 0x15, 0xcf, 0xa2, 0x8b, 0x2b, 0x68, 0x68, 0x51, 0x84, 0xa0, 0x33, 0xb9, 0x5a, 0xcc,
	0x7e, 0x9c, 0x07, 0xb3, 0xf9, 0xeb, 0xc9, 0x0f, 0xb3, 0xeb, 0xee, 0x5b, 0xe8, 0x08, 0x5a, 0x05,
	0xe6, 0x7f, 0x3b, 0xb9, 0xee, 0x3a, 0xa8, 0x07, 0x5d, 0x0b, 0x08, 0x6e, 0x67, 0xdf, 0xcd, 0xbb,
	0x7b, 0xd3, 0x2f, 0xa1, 0x1b, 0xd2, 0x75, 0xe5, 0xa8, 0xd3, 0xe3, 0x2b, 0x2b, 0xba, 0x91, 0xa3,
	0xbc, 0x71, 0x7e, 0xee, 0xd8, 0x94, 0x6c, 0xb9, 0x6c, 0xa8, 0x19, 0x7f, 0xfe, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xd2, 0x46, 0x43, 0x90, 0x07, 0x00, 0x00,
}
