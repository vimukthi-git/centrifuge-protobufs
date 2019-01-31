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
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{0}
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
	Signatures []*Signature `protobuf:"bytes,6,rep,name=signatures" json:"signatures,omitempty"`
	// Document a serialized document
	EmbeddedData      *any.Any `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData" json:"embedded_data,omitempty"`
	EmbeddedDataSalts *any.Any `protobuf:"bytes,14,opt,name=embedded_data_salts,json=embeddedDataSalts" json:"embedded_data_salts,omitempty"`
	// CoreDocumentSalts is inlined
	CoredocumentSalts *CoreDocumentSalts `protobuf:"bytes,15,opt,name=coredocument_salts,json=coredocumentSalts" json:"coredocument_salts,omitempty"`
	// Collaborators involved in the document consensus
	Collaborators [][]byte `protobuf:"bytes,17,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	// list of roles
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	// read rules
	ReadRules            []*ReadRule `protobuf:"bytes,19,rep,name=read_rules,json=readRules" json:"read_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CoreDocument) Reset()         { *m = CoreDocument{} }
func (m *CoreDocument) String() string { return proto.CompactTextString(m) }
func (*CoreDocument) ProtoMessage()    {}
func (*CoreDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{0}
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

func (m *CoreDocument) GetEmbeddedDataSalts() *any.Any {
	if m != nil {
		return m.EmbeddedDataSalts
	}
	return nil
}

func (m *CoreDocument) GetCoredocumentSalts() *CoreDocumentSalts {
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
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{1}
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

type AccessTokenSalts struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Granter              []byte   `protobuf:"bytes,3,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee              []byte   `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	RoleIdentifier       []byte   `protobuf:"bytes,5,opt,name=role_identifier,json=roleIdentifier,proto3" json:"role_identifier,omitempty"`
	DocumentIdentifier   []byte   `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	Signature            []byte   `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Key                  []byte   `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenSalts) Reset()         { *m = AccessTokenSalts{} }
func (m *AccessTokenSalts) String() string { return proto.CompactTextString(m) }
func (*AccessTokenSalts) ProtoMessage()    {}
func (*AccessTokenSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{2}
}
func (m *AccessTokenSalts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenSalts.Unmarshal(m, b)
}
func (m *AccessTokenSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenSalts.Marshal(b, m, deterministic)
}
func (dst *AccessTokenSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenSalts.Merge(dst, src)
}
func (m *AccessTokenSalts) XXX_Size() int {
	return xxx_messageInfo_AccessTokenSalts.Size(m)
}
func (m *AccessTokenSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenSalts.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenSalts proto.InternalMessageInfo

func (m *AccessTokenSalts) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *AccessTokenSalts) GetGranter() []byte {
	if m != nil {
		return m.Granter
	}
	return nil
}

func (m *AccessTokenSalts) GetGrantee() []byte {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *AccessTokenSalts) GetRoleIdentifier() []byte {
	if m != nil {
		return m.RoleIdentifier
	}
	return nil
}

func (m *AccessTokenSalts) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *AccessTokenSalts) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AccessTokenSalts) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type CoreDocumentSalts struct {
	DocumentIdentifier   []byte            `protobuf:"bytes,9,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	PreviousVersion      []byte            `protobuf:"bytes,1,opt,name=previous_version,json=previousVersion,proto3" json:"previous_version,omitempty"`
	CurrentVersion       []byte            `protobuf:"bytes,3,opt,name=current_version,json=currentVersion,proto3" json:"current_version,omitempty"`
	NextVersion          []byte            `protobuf:"bytes,4,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	PreviousRoot         []byte            `protobuf:"bytes,2,opt,name=previous_root,json=previousRoot,proto3" json:"previous_root,omitempty"`
	Collaborators        [][]byte          `protobuf:"bytes,6,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	CollaboratorsLength  []byte            `protobuf:"bytes,7,opt,name=collaborators_length,json=collaboratorsLength,proto3" json:"collaborators_length,omitempty"`
	Roles                []*SaltedRole     `protobuf:"bytes,8,rep,name=roles" json:"roles,omitempty"`
	RolesLength          []byte            `protobuf:"bytes,10,opt,name=roles_length,json=rolesLength,proto3" json:"roles_length,omitempty"`
	ReadRules            []*SaltedReadRule `protobuf:"bytes,11,rep,name=read_rules,json=readRules" json:"read_rules,omitempty"`
	ReadRulesLength      []byte            `protobuf:"bytes,12,opt,name=read_rules_length,json=readRulesLength,proto3" json:"read_rules_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CoreDocumentSalts) Reset()         { *m = CoreDocumentSalts{} }
func (m *CoreDocumentSalts) String() string { return proto.CompactTextString(m) }
func (*CoreDocumentSalts) ProtoMessage()    {}
func (*CoreDocumentSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{3}
}
func (m *CoreDocumentSalts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreDocumentSalts.Unmarshal(m, b)
}
func (m *CoreDocumentSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreDocumentSalts.Marshal(b, m, deterministic)
}
func (dst *CoreDocumentSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreDocumentSalts.Merge(dst, src)
}
func (m *CoreDocumentSalts) XXX_Size() int {
	return xxx_messageInfo_CoreDocumentSalts.Size(m)
}
func (m *CoreDocumentSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreDocumentSalts.DiscardUnknown(m)
}

var xxx_messageInfo_CoreDocumentSalts proto.InternalMessageInfo

func (m *CoreDocumentSalts) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CoreDocumentSalts) GetPreviousVersion() []byte {
	if m != nil {
		return m.PreviousVersion
	}
	return nil
}

func (m *CoreDocumentSalts) GetCurrentVersion() []byte {
	if m != nil {
		return m.CurrentVersion
	}
	return nil
}

func (m *CoreDocumentSalts) GetNextVersion() []byte {
	if m != nil {
		return m.NextVersion
	}
	return nil
}

func (m *CoreDocumentSalts) GetPreviousRoot() []byte {
	if m != nil {
		return m.PreviousRoot
	}
	return nil
}

func (m *CoreDocumentSalts) GetCollaborators() [][]byte {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *CoreDocumentSalts) GetCollaboratorsLength() []byte {
	if m != nil {
		return m.CollaboratorsLength
	}
	return nil
}

func (m *CoreDocumentSalts) GetRoles() []*SaltedRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *CoreDocumentSalts) GetRolesLength() []byte {
	if m != nil {
		return m.RolesLength
	}
	return nil
}

func (m *CoreDocumentSalts) GetReadRules() []*SaltedReadRule {
	if m != nil {
		return m.ReadRules
	}
	return nil
}

func (m *CoreDocumentSalts) GetReadRulesLength() []byte {
	if m != nil {
		return m.ReadRulesLength
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
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{4}
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
	AccessTokens         []*AccessToken `protobuf:"bytes,5,rep,name=access_tokens,json=accessTokens" json:"access_tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{5}
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

type SaltedRole struct {
	RoleKey              []byte              `protobuf:"bytes,1,opt,name=role_key,json=roleKey,proto3" json:"role_key,omitempty"`
	Collaborators        [][]byte            `protobuf:"bytes,2,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	CollaboratorsLength  []byte              `protobuf:"bytes,3,opt,name=collaborators_length,json=collaboratorsLength,proto3" json:"collaborators_length,omitempty"`
	Nfts                 [][]byte            `protobuf:"bytes,4,rep,name=nfts,proto3" json:"nfts,omitempty"`
	NftsLength           []byte              `protobuf:"bytes,5,opt,name=nfts_length,json=nftsLength,proto3" json:"nfts_length,omitempty"`
	AccessTokens         []*AccessTokenSalts `protobuf:"bytes,6,rep,name=access_tokens,json=accessTokens" json:"access_tokens,omitempty"`
	AccessTokensLength   []byte              `protobuf:"bytes,7,opt,name=access_tokens_length,json=accessTokensLength,proto3" json:"access_tokens_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SaltedRole) Reset()         { *m = SaltedRole{} }
func (m *SaltedRole) String() string { return proto.CompactTextString(m) }
func (*SaltedRole) ProtoMessage()    {}
func (*SaltedRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{6}
}
func (m *SaltedRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaltedRole.Unmarshal(m, b)
}
func (m *SaltedRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaltedRole.Marshal(b, m, deterministic)
}
func (dst *SaltedRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaltedRole.Merge(dst, src)
}
func (m *SaltedRole) XXX_Size() int {
	return xxx_messageInfo_SaltedRole.Size(m)
}
func (m *SaltedRole) XXX_DiscardUnknown() {
	xxx_messageInfo_SaltedRole.DiscardUnknown(m)
}

var xxx_messageInfo_SaltedRole proto.InternalMessageInfo

func (m *SaltedRole) GetRoleKey() []byte {
	if m != nil {
		return m.RoleKey
	}
	return nil
}

func (m *SaltedRole) GetCollaborators() [][]byte {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *SaltedRole) GetCollaboratorsLength() []byte {
	if m != nil {
		return m.CollaboratorsLength
	}
	return nil
}

func (m *SaltedRole) GetNfts() [][]byte {
	if m != nil {
		return m.Nfts
	}
	return nil
}

func (m *SaltedRole) GetNftsLength() []byte {
	if m != nil {
		return m.NftsLength
	}
	return nil
}

func (m *SaltedRole) GetAccessTokens() []*AccessTokenSalts {
	if m != nil {
		return m.AccessTokens
	}
	return nil
}

func (m *SaltedRole) GetAccessTokensLength() []byte {
	if m != nil {
		return m.AccessTokensLength
	}
	return nil
}

type ReadRule struct {
	Roles                [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Action               Action   `protobuf:"varint,4,opt,name=action,enum=coredocument.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRule) Reset()         { *m = ReadRule{} }
func (m *ReadRule) String() string { return proto.CompactTextString(m) }
func (*ReadRule) ProtoMessage()    {}
func (*ReadRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{7}
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

type SaltedReadRule struct {
	Roles                [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	RolesLength          []byte   `protobuf:"bytes,3,opt,name=roles_length,json=rolesLength,proto3" json:"roles_length,omitempty"`
	Action               []byte   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaltedReadRule) Reset()         { *m = SaltedReadRule{} }
func (m *SaltedReadRule) String() string { return proto.CompactTextString(m) }
func (*SaltedReadRule) ProtoMessage()    {}
func (*SaltedReadRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_c963b2bc2a7a56c8, []int{8}
}
func (m *SaltedReadRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaltedReadRule.Unmarshal(m, b)
}
func (m *SaltedReadRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaltedReadRule.Marshal(b, m, deterministic)
}
func (dst *SaltedReadRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaltedReadRule.Merge(dst, src)
}
func (m *SaltedReadRule) XXX_Size() int {
	return xxx_messageInfo_SaltedReadRule.Size(m)
}
func (m *SaltedReadRule) XXX_DiscardUnknown() {
	xxx_messageInfo_SaltedReadRule.DiscardUnknown(m)
}

var xxx_messageInfo_SaltedReadRule proto.InternalMessageInfo

func (m *SaltedReadRule) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *SaltedReadRule) GetRolesLength() []byte {
	if m != nil {
		return m.RolesLength
	}
	return nil
}

func (m *SaltedReadRule) GetAction() []byte {
	if m != nil {
		return m.Action
	}
	return nil
}

func init() {
	proto.RegisterType((*CoreDocument)(nil), "coredocument.CoreDocument")
	proto.RegisterType((*AccessToken)(nil), "coredocument.AccessToken")
	proto.RegisterType((*AccessTokenSalts)(nil), "coredocument.AccessTokenSalts")
	proto.RegisterType((*CoreDocumentSalts)(nil), "coredocument.CoreDocumentSalts")
	proto.RegisterType((*Signature)(nil), "coredocument.Signature")
	proto.RegisterType((*Role)(nil), "coredocument.Role")
	proto.RegisterType((*SaltedRole)(nil), "coredocument.SaltedRole")
	proto.RegisterType((*ReadRule)(nil), "coredocument.ReadRule")
	proto.RegisterType((*SaltedReadRule)(nil), "coredocument.SaltedReadRule")
	proto.RegisterEnum("coredocument.Action", Action_name, Action_value)
}

func init() {
	proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_coredocument_c963b2bc2a7a56c8)
}

var fileDescriptor_coredocument_c963b2bc2a7a56c8 = []byte{
	// 988 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x4d, 0x93, 0x26, 0x2f, 0x4e, 0x9a, 0x4e, 0xa3, 0xc5, 0x2d, 0xcb, 0x36, 0x1b, 0x90,
	0x36, 0x54, 0x90, 0x40, 0x11, 0x02, 0x09, 0x84, 0x48, 0x5a, 0x84, 0xa2, 0x5d, 0x95, 0x95, 0x5b,
	0xf6, 0xc0, 0xc5, 0x72, 0xec, 0x69, 0xb0, 0xea, 0x78, 0xa2, 0xf1, 0x64, 0x45, 0xfe, 0x18, 0x04,
	0x47, 0x0e, 0x1c, 0xe0, 0xca, 0xad, 0x7f, 0x0f, 0x17, 0x2e, 0xbd, 0xa3, 0xf9, 0xb2, 0xc7, 0xf9,
	0x10, 0x5a, 0x09, 0x2e, 0x7b, 0x8a, 0xe7, 0xbd, 0xdf, 0xfb, 0xcd, 0xbc, 0x8f, 0xdf, 0x0b, 0x9c,
	0x04, 0x84, 0xe2, 0x90, 0x04, 0x8b, 0x19, 0x4e, 0xd8, 0xc0, 0x3c, 0xf4, 0xe7, 0x94, 0x30, 0x82,
	0x6c, 0xd3, 0x76, 0x7c, 0x34, 0x25, 0x64, 0x1a, 0xe3, 0x81, 0xf0, 0x4d, 0x16, 0x37, 0x03, 0x3f,
	0x59, 0x4a, 0xe0, 0xf1, 0xc9, 0xaa, 0x8b, 0x45, 0x33, 0x9c, 0x32, 0x7f, 0x36, 0x57, 0x80, 0x27,
	0x73, 0x8a, 0x83, 0x28, 0xc5, 0x1f, 0xcc, 0x29, 0x21, 0x37, 0xe9, 0x20, 0xff, 0x61, 0x44, 0x1e,
	0x24, 0xb0, 0xfb, 0x53, 0x05, 0xec, 0x73, 0x42, 0xf1, 0x85, 0xba, 0x15, 0x0d, 0xe0, 0x50, 0xbf,
	0xc0, 0x8b, 0x42, 0x9c, 0xb0, 0xe8, 0x26, 0xc2, 0xd4, 0xa9, 0x75, 0xac, 0x9e, 0xed, 0x22, 0xed,
	0x1a, 0x67, 0x1e, 0xf4, 0x1e, 0xb4, 0xe6, 0x14, 0xbf, 0x8c, 0xc8, 0x22, 0xf5, 0x5e, 0x62, 0x9a,
	0x46, 0x24, 0x71, 0x5a, 0x02, 0xbd, 0xaf, 0xed, 0x2f, 0xa4, 0x19, 0x3d, 0x81, 0xfd, 0x60, 0x41,
	0x29, 0xa7, 0xd6, 0xc8, 0x92, 0x40, 0x36, 0x95, 0x59, 0x03, 0x1f, 0x83, 0x9d, 0xe0, 0x1f, 0x73,
	0xd4, 0xae, 0x40, 0xd5, 0xb9, 0x4d, 0x43, 0x4e, 0xa1, 0x91, 0xbd, 0x93, 0x12, 0xc2, 0x9c, 0x3d,
	0x8e, 0x19, 0x95, 0x7f, 0xb9, 0xbb, 0x07, 0xcb, 0xb5, 0xb5, 0xcf, 0x25, 0x84, 0xa1, 0x1e, 0xd8,
	0x69, 0x34, 0x4d, 0xa2, 0x64, 0x2a, 0xa1, 0x60, 0x42, 0xeb, 0xca, 0x25, 0x90, 0xa7, 0xd0, 0xc8,
	0x92, 0x11, 0xd0, 0x1d, 0x09, 0xfd, 0x55, 0xb2, 0x6a, 0x9f, 0xc0, 0x76, 0xa1, 0x16, 0xfa, 0xcc,
	0x97, 0xb8, 0xb2, 0x49, 0x59, 0xe5, 0x76, 0x81, 0xf9, 0x12, 0x80, 0xd3, 0xfb, 0x6c, 0x41, 0x71,
	0xea, 0x54, 0x3a, 0xa5, 0x5e, 0xfd, 0xec, 0xcd, 0x7e, 0xa1, 0xf5, 0x57, 0xda, 0xaf, 0xa3, 0x8d,
	0x08, 0xf4, 0x15, 0x34, 0xf0, 0x6c, 0x82, 0xc3, 0x10, 0x87, 0x1e, 0x27, 0x75, 0x1a, 0x1d, 0xab,
	0x57, 0x3f, 0x6b, 0xf7, 0xe5, 0x00, 0xf4, 0xf5, 0x00, 0xf4, 0x87, 0xc9, 0x32, 0xcb, 0x5d, 0x47,
	0x5c, 0xf8, 0xcc, 0x47, 0x4f, 0xe1, 0xb0, 0xc0, 0xe0, 0xa5, 0x7e, 0xcc, 0x52, 0xa7, 0xf9, 0xef,
	0x3c, 0x07, 0x26, 0xcf, 0x15, 0x8f, 0x42, 0xdf, 0x01, 0x32, 0xdf, 0xae, 0xb8, 0xf6, 0x05, 0xd7,
	0x49, 0x31, 0x2d, 0x73, 0xa8, 0x44, 0x70, 0x46, 0x6b, 0xc2, 0x24, 0xed, 0xbb, 0xd0, 0x08, 0x48,
	0x1c, 0xfb, 0x13, 0x42, 0x7d, 0x46, 0x68, 0xea, 0x1c, 0x74, 0x4a, 0x3d, 0xdb, 0x2d, 0x1a, 0xd1,
	0x17, 0x50, 0xa6, 0x24, 0xc6, 0xa9, 0x63, 0x89, 0x32, 0xa2, 0xe2, 0x7d, 0x2e, 0x89, 0xf1, 0x08,
	0xfd, 0x7e, 0x77, 0x0f, 0x9d, 0x3f, 0xef, 0xee, 0xa1, 0xca, 0xa1, 0xde, 0x2d, 0x5e, 0xba, 0x32,
	0x08, 0x7d, 0x02, 0x40, 0xb1, 0x1f, 0x7a, 0x74, 0xc1, 0x29, 0x0e, 0x05, 0xc5, 0x83, 0x15, 0x0a,
	0xec, 0x87, 0xee, 0x22, 0xc6, 0x6e, 0x8d, 0xaa, 0xaf, 0xb4, 0xfb, 0x97, 0x05, 0xf5, 0x61, 0x10,
	0xe0, 0x34, 0xbd, 0x26, 0xb7, 0x38, 0x41, 0x8f, 0x00, 0x0c, 0x55, 0x58, 0x62, 0x2e, 0x0d, 0x0b,
	0x72, 0x60, 0x6f, 0x4a, 0xfd, 0x84, 0x61, 0xaa, 0x46, 0x5b, 0x1f, 0x73, 0x0f, 0x56, 0xe3, 0xac,
	0x8f, 0x5c, 0x16, 0xe2, 0xb5, 0x06, 0x71, 0x59, 0xca, 0x82, 0x9b, 0x0d, 0xa9, 0x6d, 0xd1, 0xe6,
	0xce, 0x56, 0x6d, 0x3e, 0x84, 0x5a, 0x36, 0x4c, 0x4e, 0x45, 0xc0, 0x72, 0x03, 0x6a, 0x41, 0xe9,
	0x16, 0x2f, 0xa5, 0x70, 0x5c, 0xfe, 0xd9, 0xfd, 0xdb, 0x82, 0x96, 0x91, 0xad, 0xec, 0xce, 0xeb,
	0x9d, 0xf2, 0x7d, 0x09, 0x0e, 0xd6, 0x66, 0xf5, 0xbf, 0xd9, 0x82, 0xd6, 0xff, 0xbf, 0x05, 0xdf,
	0xd9, 0xb8, 0xaf, 0x56, 0x16, 0xd5, 0x9a, 0xbc, 0x2a, 0x9b, 0xe4, 0xf5, 0x11, 0xb4, 0x0b, 0x06,
	0x2f, 0xc6, 0xc9, 0x94, 0xfd, 0xa0, 0x6a, 0x75, 0x58, 0xf0, 0x3d, 0x13, 0x2e, 0xd4, 0xd7, 0x8a,
	0xac, 0x0a, 0x39, 0x39, 0x2b, 0x8b, 0xcd, 0x8f, 0x19, 0x0e, 0xb9, 0x2e, 0xb5, 0x06, 0x1f, 0x83,
	0x2d, 0x3e, 0x34, 0x35, 0xc8, 0x84, 0x84, 0x4d, 0x51, 0x7e, 0x5e, 0x90, 0x69, 0x5d, 0xf0, 0x3e,
	0xdc, 0xc8, 0xbb, 0x2e, 0x56, 0x74, 0x0a, 0x07, 0x79, 0xb0, 0xbe, 0xc4, 0x96, 0x5d, 0xc8, 0x50,
	0xf2, 0xa2, 0xee, 0xcf, 0x16, 0xd4, 0xb2, 0xd5, 0x8b, 0xde, 0x82, 0x1a, 0x6f, 0x25, 0x5b, 0x7a,
	0x51, 0xa8, 0xfa, 0x56, 0x95, 0x86, 0x71, 0x88, 0xde, 0x06, 0x98, 0x2f, 0x26, 0x71, 0x14, 0xf0,
	0x7d, 0xa2, 0x2a, 0x5c, 0x93, 0x96, 0xa7, 0x78, 0x59, 0x9c, 0xb8, 0xd2, 0xea, 0xc4, 0x7d, 0x06,
	0xb5, 0xec, 0xcf, 0x59, 0x74, 0xb0, 0x7e, 0x76, 0xbc, 0xb6, 0x75, 0xaf, 0x35, 0xc2, 0xcd, 0xc1,
	0xdd, 0x3f, 0x2c, 0xd8, 0xe5, 0xd5, 0x43, 0x47, 0xf9, 0x36, 0x53, 0x6f, 0xdb, 0xe3, 0x67, 0x7e,
	0xf7, 0x5a, 0x6b, 0x4b, 0x9b, 0x5a, 0x8b, 0x60, 0x37, 0xb9, 0x61, 0xa9, 0xb3, 0x2b, 0x9c, 0xe2,
	0x1b, 0x5d, 0x43, 0xc3, 0x17, 0x4a, 0xf7, 0x18, 0x97, 0x7a, 0xea, 0x94, 0x45, 0xad, 0x8f, 0x8a,
	0xb5, 0x36, 0x96, 0xc1, 0xa8, 0x9d, 0x2d, 0x57, 0x43, 0xfd, 0xae, 0xed, 0xe7, 0x90, 0xb4, 0xfb,
	0xdb, 0x0e, 0x40, 0xde, 0xf7, 0x57, 0x7a, 0xf9, 0xce, 0xab, 0x0c, 0x65, 0x69, 0xfb, 0x50, 0x6e,
	0x4a, 0xf6, 0x04, 0xea, 0xfc, 0x57, 0x47, 0xcb, 0x55, 0x03, 0xdc, 0xa4, 0x82, 0xce, 0x57, 0xab,
	0x21, 0xff, 0xaa, 0x1f, 0x6d, 0xad, 0x86, 0x58, 0x13, 0xc5, 0xe4, 0xd1, 0x87, 0xd0, 0x2e, 0x90,
	0x14, 0x15, 0x84, 0x4c, 0xac, 0x1a, 0xc2, 0x4b, 0xa8, 0xea, 0x39, 0x46, 0x6d, 0x2d, 0x26, 0x59,
	0x08, 0x25, 0x99, 0xf7, 0xa1, 0xe2, 0x07, 0x4c, 0xab, 0xbf, 0x79, 0xd6, 0x5e, 0x7d, 0x11, 0xf7,
	0xb9, 0x0a, 0xd3, 0xf5, 0xa1, 0x59, 0x54, 0xc7, 0x16, 0xd6, 0x55, 0x21, 0x96, 0xd6, 0x85, 0xf8,
	0xa0, 0x70, 0xb1, 0xad, 0xaf, 0x38, 0x3d, 0x87, 0x8a, 0xbc, 0x14, 0x21, 0x68, 0x0e, 0xcf, 0xaf,
	0xc7, 0xdf, 0x5e, 0x7a, 0xe3, 0xcb, 0x17, 0xc3, 0x67, 0xe3, 0x8b, 0xd6, 0x1b, 0x68, 0x1f, 0xea,
	0xca, 0xe6, 0x7e, 0x3d, 0xbc, 0x68, 0x59, 0xa8, 0x0d, 0x2d, 0xc3, 0xe0, 0x5d, 0x8d, 0xbf, 0xb9,
	0x6c, 0xed, 0x8c, 0x3e, 0x85, 0x56, 0x40, 0x66, 0x85, 0x54, 0x46, 0x62, 0x0b, 0xeb, 0xd3, 0x73,
	0xae, 0x8c, 0xe7, 0xd6, 0xf7, 0x4d, 0x13, 0x32, 0x9f, 0x4c, 0x2a, 0x42, 0x32, 0x1f, 0xff, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xa2, 0x7e, 0xba, 0x4b, 0x0b, 0x00, 0x00,
}
