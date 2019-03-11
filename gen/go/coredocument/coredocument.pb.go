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
	// read_sign represents reading as well the sign the documents. We will pick this one when requesting the signatures.
	Action_ACTION_READ_SIGN Action = 0
	// read represents just reading the doc/fields
	Action_ACTION_READ Action = 1
)

var Action_name = map[int32]string{
	0: "ACTION_READ_SIGN",
	1: "ACTION_READ",
}
var Action_value = map[string]int32{
	"ACTION_READ_SIGN": 0,
	"ACTION_READ":      1,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}
func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{0}
}

type FieldMatchType int32

const (
	FieldMatchType_FIELD_MATCH_TYPE_PREFIX FieldMatchType = 0
	FieldMatchType_FIELD_MATCH_TYPE_EXACT  FieldMatchType = 1
)

var FieldMatchType_name = map[int32]string{
	0: "FIELD_MATCH_TYPE_PREFIX",
	1: "FIELD_MATCH_TYPE_EXACT",
}
var FieldMatchType_value = map[string]int32{
	"FIELD_MATCH_TYPE_PREFIX": 0,
	"FIELD_MATCH_TYPE_EXACT":  1,
}

func (x FieldMatchType) String() string {
	return proto.EnumName(FieldMatchType_name, int32(x))
}
func (FieldMatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{1}
}

type TransitionAction int32

const (
	TransitionAction_TRANSITION_ACTION_EDIT TransitionAction = 0
)

var TransitionAction_name = map[int32]string{
	0: "TRANSITION_ACTION_EDIT",
}
var TransitionAction_value = map[string]int32{
	"TRANSITION_ACTION_EDIT": 0,
}

func (x TransitionAction) String() string {
	return proto.EnumName(TransitionAction_name, int32(x))
}
func (TransitionAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{2}
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
	// current_preimage is the sha256 pre-image of the current_version. It prevents current state commitment id(anchor id) from getting exposed.
	// 32 byte value
	CurrentPreimage []byte `protobuf:"bytes,23,opt,name=current_preimage,json=currentPreimage,proto3" json:"current_preimage,omitempty"`
	// next_version is the version that is going to be used for the next version if any
	// party wants to update the state.
	NextVersion []byte `protobuf:"bytes,4,opt,name=next_version,json=nextVersion,proto3" json:"next_version,omitempty"`
	// next_preimage is the sha256 pre-image of the next_version. It prevents next state commitment id(anchor id) from getting exposed.
	NextPreimage []byte `protobuf:"bytes,22,opt,name=next_preimage,json=nextPreimage,proto3" json:"next_preimage,omitempty"`
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
	// Signatures of the signature_root by centrifuge identities. This array should be sorted by the Centrifuge ID
	Signatures []*Signature `protobuf:"bytes,6,rep,name=signatures" json:"signatures,omitempty"`
	// Document a serialized document
	EmbeddedData      *any.Any        `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData" json:"embedded_data,omitempty"`
	EmbeddedDataSalts []*DocumentSalt `protobuf:"bytes,14,rep,name=embedded_data_salts,json=embeddedDataSalts" json:"embedded_data_salts,omitempty"`
	// CoreDocumentSalts is inlined
	CoredocumentSalts []*DocumentSalt `protobuf:"bytes,15,rep,name=coredocument_salts,json=coredocumentSalts" json:"coredocument_salts,omitempty"`
	// list of roles
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	// read rules
	ReadRules []*ReadRule `protobuf:"bytes,19,rep,name=read_rules,json=readRules" json:"read_rules,omitempty"`
	// transition rules
	TransitionRules []*TransitionRule `protobuf:"bytes,24,rep,name=transition_rules,json=transitionRules" json:"transition_rules,omitempty"`
	// nft list for uniqueness check
	Nfts []*NFT `protobuf:"bytes,20,rep,name=nfts" json:"nfts,omitempty"`
	// AccessTokens which have been added to this CoreDocument
	AccessTokens []*AccessToken `protobuf:"bytes,21,rep,name=access_tokens,json=accessTokens" json:"access_tokens,omitempty"`
	// latest update log
	Update               *UpdateLog `protobuf:"bytes,25,opt,name=update" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CoreDocument) Reset()         { *m = CoreDocument{} }
func (m *CoreDocument) String() string { return proto.CompactTextString(m) }
func (*CoreDocument) ProtoMessage()    {}
func (*CoreDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{0}
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

func (m *CoreDocument) GetCurrentPreimage() []byte {
	if m != nil {
		return m.CurrentPreimage
	}
	return nil
}

func (m *CoreDocument) GetNextVersion() []byte {
	if m != nil {
		return m.NextVersion
	}
	return nil
}

func (m *CoreDocument) GetNextPreimage() []byte {
	if m != nil {
		return m.NextPreimage
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

func (m *CoreDocument) GetTransitionRules() []*TransitionRule {
	if m != nil {
		return m.TransitionRules
	}
	return nil
}

func (m *CoreDocument) GetNfts() []*NFT {
	if m != nil {
		return m.Nfts
	}
	return nil
}

func (m *CoreDocument) GetAccessTokens() []*AccessToken {
	if m != nil {
		return m.AccessTokens
	}
	return nil
}

func (m *CoreDocument) GetUpdate() *UpdateLog {
	if m != nil {
		return m.Update
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
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{1}
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
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{2}
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
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{3}
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
	// role key which is used to identify the group internally and map the role to rules
	RoleKey []byte `protobuf:"bytes,1,opt,name=role_key,json=roleKey,proto3" json:"role_key,omitempty"`
	// collaborators holds the list of document collaborators
	Collaborators [][]byte `protobuf:"bytes,3,rep,name=collaborators,proto3" json:"collaborators,omitempty"`
	// nfts is a list of registry address/tokenID pairs.
	// For easier verification in merkle proofs, the values are simply concatenated with the first 22 bytes being the NFT registry and the remaining 32 bytes the tokenID.
	Nfts                 [][]byte `protobuf:"bytes,4,rep,name=nfts,proto3" json:"nfts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{4}
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
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{5}
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
	return Action_ACTION_READ_SIGN
}

type TransitionRule struct {
	// rule key, to help track of the rule
	RuleKey []byte `protobuf:"bytes,1,opt,name=rule_key,json=ruleKey,proto3" json:"rule_key,omitempty"`
	// Indicates which roles can make changes or read the fields specified:
	// this list holds role keys correlated to those in the 'roles' field of the CoreDocument
	Roles [][]byte `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	// prefix or exact
	MatchType FieldMatchType `protobuf:"varint,3,opt,name=match_type,json=matchType,enum=coredocument.FieldMatchType" json:"match_type,omitempty"`
	// compact property of the field
	Field []byte `protobuf:"bytes,4,opt,name=field,proto3" json:"field,omitempty"`
	// what kind of action this rule allows
	Action               TransitionAction `protobuf:"varint,5,opt,name=action,enum=coredocument.TransitionAction" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransitionRule) Reset()         { *m = TransitionRule{} }
func (m *TransitionRule) String() string { return proto.CompactTextString(m) }
func (*TransitionRule) ProtoMessage()    {}
func (*TransitionRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{6}
}
func (m *TransitionRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionRule.Unmarshal(m, b)
}
func (m *TransitionRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionRule.Marshal(b, m, deterministic)
}
func (dst *TransitionRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionRule.Merge(dst, src)
}
func (m *TransitionRule) XXX_Size() int {
	return xxx_messageInfo_TransitionRule.Size(m)
}
func (m *TransitionRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionRule.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionRule proto.InternalMessageInfo

func (m *TransitionRule) GetRuleKey() []byte {
	if m != nil {
		return m.RuleKey
	}
	return nil
}

func (m *TransitionRule) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *TransitionRule) GetMatchType() FieldMatchType {
	if m != nil {
		return m.MatchType
	}
	return FieldMatchType_FIELD_MATCH_TYPE_PREFIX
}

func (m *TransitionRule) GetField() []byte {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *TransitionRule) GetAction() TransitionAction {
	if m != nil {
		return m.Action
	}
	return TransitionAction_TRANSITION_ACTION_EDIT
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
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{7}
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

type UpdateLog struct {
	// author of the update
	EntityId []byte `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// where this update originated from, currently a libp2p peer id string
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	// timestamp of the update
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateLog) Reset()         { *m = UpdateLog{} }
func (m *UpdateLog) String() string { return proto.CompactTextString(m) }
func (*UpdateLog) ProtoMessage()    {}
func (*UpdateLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_d9142c99b8a5cf51, []int{8}
}
func (m *UpdateLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateLog.Unmarshal(m, b)
}
func (m *UpdateLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateLog.Marshal(b, m, deterministic)
}
func (dst *UpdateLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateLog.Merge(dst, src)
}
func (m *UpdateLog) XXX_Size() int {
	return xxx_messageInfo_UpdateLog.Size(m)
}
func (m *UpdateLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateLog.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateLog proto.InternalMessageInfo

func (m *UpdateLog) GetEntityId() []byte {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *UpdateLog) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UpdateLog) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
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
	proto.RegisterType((*TransitionRule)(nil), "coredocument.TransitionRule")
	proto.RegisterType((*NFT)(nil), "coredocument.NFT")
	proto.RegisterType((*UpdateLog)(nil), "coredocument.UpdateLog")
	proto.RegisterEnum("coredocument.Action", Action_name, Action_value)
	proto.RegisterEnum("coredocument.FieldMatchType", FieldMatchType_name, FieldMatchType_value)
	proto.RegisterEnum("coredocument.TransitionAction", TransitionAction_name, TransitionAction_value)
}

func init() {
	proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_coredocument_d9142c99b8a5cf51)
}

var fileDescriptor_coredocument_d9142c99b8a5cf51 = []byte{
	// 1094 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0xde, 0x34, 0x4d, 0xda, 0xbc, 0xfc, 0xa8, 0x3b, 0xcd, 0xb6, 0x6e, 0x77, 0xd9, 0x96, 0x80,
	0xb4, 0xdd, 0x0a, 0x12, 0xa9, 0x88, 0x1f, 0x12, 0x68, 0x45, 0xd2, 0xa6, 0x8b, 0xc5, 0x6e, 0x08,
	0xae, 0x41, 0x0b, 0x1c, 0xac, 0x89, 0x3d, 0x0d, 0x56, 0x13, 0x8f, 0x35, 0x1e, 0x57, 0x44, 0xfc,
	0x31, 0x70, 0xe4, 0xc8, 0x99, 0x5b, 0xef, 0x9c, 0xf9, 0x27, 0x38, 0xf7, 0x8e, 0x66, 0xc6, 0xe3,
	0xd8, 0x81, 0x22, 0x38, 0xc5, 0xef, 0xbd, 0xef, 0x7d, 0xf3, 0xde, 0xcc, 0xfb, 0x5e, 0xe0, 0xd0,
	0xa3, 0x8c, 0xf8, 0xd4, 0x4b, 0xe6, 0x24, 0xe4, 0xbd, 0xbc, 0xd1, 0x8d, 0x18, 0xe5, 0x14, 0x35,
	0xf2, 0xbe, 0x83, 0xfd, 0x29, 0xa5, 0xd3, 0x19, 0xe9, 0xc9, 0xd8, 0x24, 0xb9, 0xea, 0xe1, 0x70,
	0xa1, 0x80, 0x07, 0x87, 0xab, 0x21, 0x1e, 0xcc, 0x49, 0xcc, 0xf1, 0x3c, 0x4a, 0x01, 0x4f, 0x23,
	0x46, 0xbc, 0x20, 0x26, 0xef, 0x46, 0x8c, 0xd2, 0xab, 0xb8, 0xb7, 0xfc, 0xe1, 0x54, 0x19, 0x0a,
	0xd8, 0xf9, 0x63, 0x13, 0x1a, 0x67, 0x94, 0x91, 0xf3, 0xf4, 0x54, 0xd4, 0x83, 0x1d, 0x5d, 0x81,
	0x1b, 0xf8, 0x24, 0xe4, 0xc1, 0x55, 0x40, 0x98, 0x59, 0x3b, 0x2a, 0x1d, 0x37, 0x6c, 0xa4, 0x43,
	0x56, 0x16, 0x41, 0xcf, 0xc0, 0x88, 0x18, 0xb9, 0x09, 0x68, 0x12, 0xbb, 0x37, 0x84, 0xc5, 0x01,
	0x0d, 0x4d, 0x43, 0xa2, 0xb7, 0xb4, 0xff, 0x6b, 0xe5, 0x46, 0x4f, 0x61, 0xcb, 0x4b, 0x18, 0x13,
	0xd4, 0x1a, 0x59, 0x96, 0xc8, 0x56, 0xea, 0xd6, 0xc0, 0x67, 0x60, 0x68, 0x60, 0xc4, 0x48, 0x30,
	0xc7, 0x53, 0x62, 0xee, 0x29, 0xce, 0xd4, 0x3f, 0x4e, 0xdd, 0xe8, 0x4d, 0x68, 0x84, 0xe4, 0x87,
	0x25, 0xe1, 0xba, 0x84, 0xd5, 0x85, 0x4f, 0xb3, 0xbd, 0x05, 0x4d, 0x09, 0xc9, 0xa8, 0x76, 0x25,
	0x46, 0xe6, 0x65, 0x3c, 0x27, 0xd0, 0xcc, 0xfa, 0x66, 0x94, 0x72, 0x73, 0x43, 0x80, 0x06, 0x95,
	0x9f, 0x6f, 0xef, 0xa0, 0x64, 0x37, 0x74, 0xcc, 0xa6, 0x94, 0xa3, 0x63, 0x68, 0xc4, 0xc1, 0x34,
	0x0c, 0xc2, 0xa9, 0x82, 0x42, 0x1e, 0x5a, 0x4f, 0x43, 0x12, 0x79, 0x02, 0xcd, 0xec, 0x72, 0x24,
	0x74, 0x4d, 0x41, 0x7f, 0x51, 0xac, 0x3a, 0x26, 0xb1, 0x1d, 0xa8, 0xf9, 0x98, 0x63, 0x85, 0xab,
	0xe4, 0x29, 0x37, 0x85, 0x5f, 0x62, 0x9e, 0x03, 0x08, 0x7a, 0xcc, 0x13, 0x46, 0x62, 0xb3, 0x7a,
	0x54, 0x3e, 0xae, 0x9f, 0xee, 0x75, 0x0b, 0xa3, 0x74, 0xa9, 0xe3, 0x3a, 0x3b, 0x97, 0x81, 0x3e,
	0x85, 0x26, 0x99, 0x4f, 0x88, 0xef, 0x13, 0xdf, 0x15, 0xa4, 0x66, 0xf3, 0xa8, 0x74, 0x5c, 0x3f,
	0x6d, 0x77, 0xd5, 0x40, 0x75, 0xf5, 0x40, 0x75, 0xfb, 0xe1, 0x22, 0xeb, 0x5d, 0x67, 0x9c, 0x63,
	0x8e, 0xd1, 0x97, 0xb0, 0x53, 0x60, 0x70, 0x63, 0x3c, 0xe3, 0xb1, 0xd9, 0x92, 0xa5, 0x1c, 0x14,
	0x4b, 0xd1, 0x43, 0x75, 0x89, 0x67, 0x5c, 0xb3, 0x6d, 0xe7, 0xd9, 0x44, 0x20, 0x46, 0x63, 0x40,
	0xf9, 0xb4, 0x94, 0x71, 0xeb, 0x3f, 0x33, 0xe6, 0x11, 0x8a, 0xf1, 0x13, 0xa8, 0x30, 0x3a, 0x23,
	0xb1, 0x59, 0x92, 0x24, 0xa8, 0x48, 0x62, 0xd3, 0x19, 0x19, 0xa0, 0x5f, 0x6f, 0xef, 0xe0, 0xe8,
	0xb7, 0xdb, 0x3b, 0xd8, 0x14, 0x50, 0xf7, 0x9a, 0x2c, 0x6c, 0x95, 0x84, 0xde, 0x07, 0x60, 0x04,
	0xfb, 0x2e, 0x4b, 0x04, 0xc5, 0x8e, 0xa4, 0xd8, 0x5d, 0xa1, 0x20, 0xd8, 0xb7, 0x93, 0x19, 0xb1,
	0x6b, 0x2c, 0xfd, 0x8a, 0xd1, 0x0b, 0x30, 0x38, 0xc3, 0x61, 0x1c, 0xf0, 0x80, 0x86, 0x69, 0xb2,
	0x29, 0x93, 0x1f, 0x17, 0x93, 0x9d, 0x0c, 0x25, 0x29, 0xb6, 0x78, 0xc1, 0x16, 0xd5, 0xaf, 0x87,
	0x57, 0x3c, 0x36, 0xdb, 0x32, 0x79, 0xbb, 0x98, 0x3c, 0xba, 0x70, 0x06, 0x0f, 0xb3, 0xda, 0xeb,
	0x8c, 0x4c, 0x83, 0x98, 0xb3, 0x85, 0x1b, 0xf8, 0xb6, 0xcc, 0x42, 0x0e, 0x34, 0xb1, 0xe7, 0x91,
	0x38, 0x76, 0x39, 0xbd, 0x26, 0x61, 0x6c, 0x3e, 0x94, 0x34, 0xfb, 0x45, 0x9a, 0xbe, 0x84, 0x38,
	0x02, 0x31, 0x68, 0x67, 0x74, 0xb0, 0x54, 0xbc, 0xdd, 0xc0, 0x4b, 0x48, 0x8c, 0x7a, 0x50, 0x4d,
	0x22, 0x1f, 0x73, 0x62, 0xee, 0xcb, 0x89, 0x59, 0x19, 0xba, 0xaf, 0x64, 0xec, 0x25, 0x9d, 0xda,
	0x29, 0xac, 0xf3, 0x1c, 0x1a, 0xf9, 0xc7, 0x42, 0x26, 0x6c, 0x78, 0x74, 0x1e, 0x61, 0x8f, 0x9b,
	0x25, 0x29, 0x3f, 0x6d, 0xa2, 0x36, 0x54, 0x6e, 0xf0, 0x2c, 0x21, 0x4a, 0x1b, 0xb6, 0x32, 0x3a,
	0x7f, 0x96, 0xa0, 0x9e, 0x2b, 0x12, 0x3d, 0xc9, 0x17, 0x97, 0x52, 0xe4, 0x3c, 0x82, 0x7f, 0xca,
	0x70, 0xc8, 0x09, 0x4b, 0x77, 0x8a, 0x36, 0x97, 0x11, 0x92, 0x2e, 0x07, 0x6d, 0x8a, 0x7d, 0x24,
	0xdf, 0x3e, 0x47, 0x5c, 0x51, 0xfb, 0x48, 0xb8, 0x73, 0x3b, 0xee, 0x9e, 0xa5, 0xb8, 0x76, 0xef,
	0x52, 0x7c, 0x0c, 0xb5, 0x4c, 0x75, 0x66, 0x55, 0xc2, 0x96, 0x0e, 0x64, 0x40, 0xf9, 0x9a, 0x2c,
	0xd4, 0x86, 0xb1, 0xc5, 0x67, 0xe7, 0xa7, 0x12, 0xd4, 0x32, 0xe1, 0xa2, 0x47, 0x50, 0x13, 0x4c,
	0x5c, 0xbc, 0x6a, 0xda, 0xea, 0xa6, 0x72, 0x58, 0x3e, 0x7a, 0x03, 0x20, 0x4a, 0x26, 0xb3, 0xc0,
	0x13, 0x23, 0x9b, 0x96, 0x50, 0x53, 0x9e, 0xcf, 0xc9, 0xa2, 0x78, 0x72, 0x79, 0xf5, 0xe4, 0x8f,
	0xa0, 0x96, 0xfd, 0x55, 0xc8, 0xdb, 0x10, 0x0a, 0x5b, 0xd5, 0xbe, 0xa3, 0x11, 0xf6, 0x12, 0xdc,
	0xf9, 0x0e, 0xd6, 0x85, 0x6e, 0xd0, 0xfe, 0x52, 0x2f, 0xfa, 0x21, 0x85, 0x2d, 0x8e, 0x7e, 0x1b,
	0x9a, 0x1e, 0x9d, 0xcd, 0xf0, 0x84, 0x32, 0xcc, 0x29, 0x8b, 0xcd, 0xf2, 0x51, 0xf9, 0xb8, 0x61,
	0x17, 0x9d, 0x08, 0xa5, 0xd3, 0xbd, 0x2e, 0x83, 0xf2, 0xbb, 0x33, 0x82, 0x4d, 0xad, 0x28, 0x31,
	0x0e, 0x4a, 0xbb, 0x6b, 0x12, 0x90, 0x6a, 0xf2, 0x1d, 0xa8, 0x62, 0x8f, 0xeb, 0x05, 0xdf, 0x3a,
	0x6d, 0xaf, 0x8e, 0xb3, 0x94, 0x4f, 0x8a, 0xe9, 0xfc, 0x5e, 0x82, 0x56, 0x51, 0x65, 0xb2, 0xee,
	0x64, 0xa5, 0xee, 0x44, 0xd5, 0xfd, 0xcf, 0x27, 0x7e, 0x0c, 0x30, 0xc7, 0xdc, 0xfb, 0xde, 0xe5,
	0x8b, 0x48, 0xdd, 0x64, 0x6b, 0x55, 0xc8, 0x17, 0x01, 0x99, 0xf9, 0xaf, 0x04, 0xc8, 0x59, 0x44,
	0xc4, 0xae, 0xcd, 0xf5, 0xa7, 0xa0, 0xbc, 0x12, 0xc1, 0x74, 0xe2, 0x94, 0x81, 0x3e, 0xc8, 0x9a,
	0xa8, 0x48, 0xba, 0x27, 0xf7, 0xed, 0x85, 0x95, 0x76, 0xfa, 0x50, 0x1e, 0x5d, 0x38, 0xe8, 0xb0,
	0x20, 0x77, 0xad, 0x01, 0xed, 0xb2, 0x7c, 0xd1, 0xa3, 0xd4, 0xbc, 0x88, 0xaa, 0xc1, 0xd8, 0x90,
	0xb6, 0xe5, 0x77, 0x7e, 0x84, 0x5a, 0xa6, 0xd1, 0x7f, 0x9f, 0xaf, 0x3d, 0xd8, 0x08, 0xa9, 0x4f,
	0x34, 0x47, 0xcd, 0xae, 0x0a, 0xd3, 0xf2, 0x8b, 0xb3, 0x53, 0xfe, 0x1f, 0xb3, 0x73, 0xd2, 0x83,
	0xaa, 0xea, 0x08, 0xb5, 0xc1, 0xe8, 0x9f, 0x39, 0xd6, 0x17, 0x23, 0xd7, 0x1e, 0xf6, 0xcf, 0xdd,
	0x4b, 0xeb, 0xc5, 0xc8, 0x78, 0x80, 0xb6, 0xa0, 0x9e, 0xf3, 0x1a, 0xa5, 0x13, 0x0b, 0x5a, 0xc5,
	0xbb, 0x45, 0x8f, 0x60, 0xef, 0xc2, 0x1a, 0xbe, 0x3c, 0x77, 0x5f, 0xf5, 0x9d, 0xb3, 0xcf, 0x5c,
	0xe7, 0x9b, 0xf1, 0xd0, 0x1d, 0xdb, 0xc3, 0x0b, 0xeb, 0xb5, 0xf1, 0x00, 0x1d, 0xc0, 0xee, 0xdf,
	0x82, 0xc3, 0xd7, 0xfd, 0x33, 0xc7, 0x28, 0x9d, 0x74, 0xc1, 0x58, 0xbd, 0x57, 0x81, 0x77, 0xec,
	0xfe, 0xe8, 0xd2, 0x92, 0x67, 0xa6, 0x47, 0x0f, 0xcf, 0x2d, 0xc7, 0x78, 0x30, 0xf8, 0x10, 0x0c,
	0x8f, 0xce, 0x0b, 0x0f, 0x33, 0xd8, 0x3e, 0xcb, 0x59, 0x63, 0xd1, 0xea, 0xb8, 0xf4, 0x6d, 0x2b,
	0x0f, 0x89, 0x26, 0x93, 0xaa, 0xbc, 0x83, 0xf7, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x93,
	0xd4, 0xdf, 0xe6, 0x09, 0x00, 0x00,
}
