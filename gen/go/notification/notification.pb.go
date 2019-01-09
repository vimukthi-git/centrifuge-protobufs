// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification/notification.proto

package notificationpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// NotificationMessage wraps a single CoreDocument to be notified to upstream services
type NotificationMessage struct {
	EventType            uint32               `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	CentrifugeId         string               `protobuf:"bytes,6,opt,name=centrifuge_id,json=centrifugeId,proto3" json:"centrifuge_id,omitempty"`
	Recorded             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=recorded,proto3" json:"recorded,omitempty"`
	DocumentType         string               `protobuf:"bytes,4,opt,name=document_type,json=documentType,proto3" json:"document_type,omitempty"`
	DocumentId           string               `protobuf:"bytes,7,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NotificationMessage) Reset()         { *m = NotificationMessage{} }
func (m *NotificationMessage) String() string { return proto.CompactTextString(m) }
func (*NotificationMessage) ProtoMessage()    {}
func (*NotificationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_0edff2f7bf97892c, []int{0}
}
func (m *NotificationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationMessage.Unmarshal(m, b)
}
func (m *NotificationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationMessage.Marshal(b, m, deterministic)
}
func (dst *NotificationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationMessage.Merge(dst, src)
}
func (m *NotificationMessage) XXX_Size() int {
	return xxx_messageInfo_NotificationMessage.Size(m)
}
func (m *NotificationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationMessage proto.InternalMessageInfo

func (m *NotificationMessage) GetEventType() uint32 {
	if m != nil {
		return m.EventType
	}
	return 0
}

func (m *NotificationMessage) GetCentrifugeId() string {
	if m != nil {
		return m.CentrifugeId
	}
	return ""
}

func (m *NotificationMessage) GetRecorded() *timestamp.Timestamp {
	if m != nil {
		return m.Recorded
	}
	return nil
}

func (m *NotificationMessage) GetDocumentType() string {
	if m != nil {
		return m.DocumentType
	}
	return ""
}

func (m *NotificationMessage) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func init() {
	proto.RegisterType((*NotificationMessage)(nil), "notification.NotificationMessage")
}

func init() {
	proto.RegisterFile("notification/notification.proto", fileDescriptor_notification_0edff2f7bf97892c)
}

var fileDescriptor_notification_0edff2f7bf97892c = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x09, 0xca, 0x74, 0x6f, 0x9b, 0x68, 0xbd, 0x94, 0x81, 0xb4, 0xe8, 0xa5, 0xa7, 0x14,
	0x14, 0xf4, 0xbe, 0xdb, 0x0e, 0xca, 0x28, 0x3b, 0x79, 0x19, 0x6d, 0xf2, 0x5a, 0x02, 0xb6, 0x2f,
	0xb4, 0xaf, 0xc2, 0x7e, 0xa9, 0x7f, 0x47, 0x9a, 0xd2, 0x2e, 0xc7, 0x7c, 0xf9, 0xf2, 0xf2, 0x25,
	0x10, 0x35, 0xc4, 0xa6, 0x34, 0x2a, 0x67, 0x43, 0x4d, 0xea, 0x2f, 0xa4, 0x6d, 0x89, 0x29, 0x58,
	0xfb, 0x6c, 0x1b, 0x55, 0x44, 0xd5, 0x0f, 0xa6, 0x6e, 0xaf, 0xe8, 0xcb, 0x94, 0x4d, 0x8d, 0x1d,
	0xe7, 0xb5, 0x1d, 0xf5, 0xe7, 0x3f, 0x01, 0x8f, 0x5f, 0xde, 0x89, 0x4f, 0xec, 0xba, 0xbc, 0xc2,
	0xe0, 0x09, 0x00, 0x7f, 0xb1, 0xe1, 0x13, 0x9f, 0x2d, 0x86, 0x22, 0x16, 0xc9, 0x26, 0x5b, 0x3a,
	0x72, 0x3c, 0x5b, 0x0c, 0x5e, 0x60, 0xa3, 0xb0, 0xe1, 0xd6, 0x94, 0x7d, 0x85, 0x27, 0xa3, 0xc3,
	0x45, 0x2c, 0x92, 0x65, 0xb6, 0xbe, 0xc0, 0xbd, 0x0e, 0xde, 0xe1, 0xb6, 0x45, 0x45, 0xad, 0x46,
	0x1d, 0x5e, 0xc5, 0x22, 0x59, 0xbd, 0x6e, 0xe5, 0xd8, 0x23, 0xa7, 0x1e, 0x79, 0x9c, 0x7a, 0xb2,
	0xd9, 0x1d, 0x86, 0x6b, 0x52, 0x7d, 0x3d, 0x5f, 0x7f, 0x3d, 0x0e, 0x9f, 0xa0, 0x2b, 0x88, 0x60,
	0x35, 0x4b, 0x46, 0x87, 0x37, 0x4e, 0x81, 0x09, 0xed, 0xf5, 0xee, 0x03, 0xee, 0x15, 0xd5, 0xd2,
	0xff, 0x8e, 0xdd, 0x83, 0xff, 0xd4, 0xc3, 0xd0, 0x70, 0x10, 0xdf, 0x77, 0xbe, 0x62, 0x8b, 0x62,
	0xe1, 0xe2, 0xde, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xf2, 0x79, 0x35, 0x6b, 0x01, 0x00,
	0x00,
}
