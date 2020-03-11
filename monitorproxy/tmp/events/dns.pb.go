// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DnsAnswerEntry struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type"`
	Cls  string `protobuf:"bytes,3,opt,name=cls" json:"cls"`
	Addr []byte `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr"`
	Fqdn string `protobuf:"bytes,5,opt,name=fqdn" json:"fqdn"`
	Ttl  uint32 `protobuf:"varint,6,opt,name=ttl" json:"ttl"`
}

func (m *DnsAnswerEntry) Reset()                    { *m = DnsAnswerEntry{} }
func (m *DnsAnswerEntry) String() string            { return proto.CompactTextString(m) }
func (*DnsAnswerEntry) ProtoMessage()               {}
func (*DnsAnswerEntry) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *DnsAnswerEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsAnswerEntry) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DnsAnswerEntry) GetCls() string {
	if m != nil {
		return m.Cls
	}
	return ""
}

func (m *DnsAnswerEntry) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *DnsAnswerEntry) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DnsAnswerEntry) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type DnsAnswer struct {
	DeviceSn  string                     `protobuf:"bytes,1,opt,name=device_sn,json=deviceSn" json:"device_sn"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp"`
	Entry     []*DnsAnswerEntry          `protobuf:"bytes,3,rep,name=entry" json:"entry"`
}

func (m *DnsAnswer) Reset()                    { *m = DnsAnswer{} }
func (m *DnsAnswer) String() string            { return proto.CompactTextString(m) }
func (*DnsAnswer) ProtoMessage()               {}
func (*DnsAnswer) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *DnsAnswer) GetDeviceSn() string {
	if m != nil {
		return m.DeviceSn
	}
	return ""
}

func (m *DnsAnswer) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DnsAnswer) GetEntry() []*DnsAnswerEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsAnswerEntry)(nil), "events.DnsAnswerEntry")
	proto.RegisterType((*DnsAnswer)(nil), "events.DnsAnswer")
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xdb, 0x2e, 0x26, 0x55, 0x91, 0x1c, 0x24, 0xd4, 0x83, 0x4b, 0x4f, 0x7b, 0x90,
	0x14, 0xea, 0xc5, 0xab, 0xa0, 0x2f, 0x10, 0xbd, 0xcb, 0xb6, 0x99, 0x96, 0xc2, 0xee, 0xec, 0x9a,
	0x8c, 0x95, 0xde, 0x7d, 0x00, 0x1f, 0x59, 0x26, 0xb1, 0x95, 0xde, 0x3e, 0x3e, 0xfe, 0x09, 0xff,
	0x1f, 0x25, 0x3d, 0x46, 0x3b, 0x84, 0x9e, 0x7a, 0x5d, 0xc2, 0x0e, 0x90, 0xe2, 0xf4, 0x6e, 0xd3,
	0xf7, 0x9b, 0x16, 0xe6, 0xc9, 0x2e, 0x3f, 0xd7, 0x73, 0xda, 0x76, 0x10, 0xa9, 0xe9, 0x86, 0x1c,
	0x9c, 0x7d, 0x0b, 0x75, 0xf5, 0x8c, 0xf1, 0x09, 0xe3, 0x17, 0x84, 0x17, 0xa4, 0xb0, 0xd7, 0x5a,
	0x8d, 0xb0, 0xe9, 0xc0, 0x88, 0x4a, 0xd4, 0xd2, 0x25, 0x66, 0x47, 0xfb, 0x01, 0xcc, 0x59, 0x76,
	0xcc, 0xfa, 0x5a, 0x15, 0xab, 0x36, 0x9a, 0x22, 0x29, 0x46, 0x4e, 0x35, 0xde, 0x07, 0x33, 0xaa,
	0x44, 0x7d, 0xe1, 0x12, 0xb3, 0x5b, 0x7f, 0x78, 0x34, 0xe3, 0x7c, 0xc9, 0xcc, 0x97, 0x44, 0xad,
	0x29, 0x2b, 0x51, 0x5f, 0x3a, 0xc6, 0xd9, 0x8f, 0x50, 0xf2, 0x58, 0x43, 0xdf, 0x2a, 0xe9, 0x61,
	0xb7, 0x5d, 0xc1, 0x7b, 0xc4, 0xbf, 0x1a, 0xe7, 0x59, 0xbc, 0xa2, 0x7e, 0x54, 0xf2, 0x38, 0x22,
	0xf5, 0x99, 0x2c, 0xa6, 0x36, 0xcf, 0xb4, 0x87, 0x99, 0xf6, 0xed, 0x90, 0x70, 0xff, 0x61, 0x7d,
	0xaf, 0xc6, 0xc0, 0x0b, 0x4d, 0x51, 0x15, 0xf5, 0x64, 0x71, 0x63, 0xf3, 0x27, 0xd9, 0xd3, 0xfd,
	0x2e, 0x87, 0x96, 0x65, 0x7a, 0xec, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa0, 0x6a, 0x32,
	0x56, 0x01, 0x00, 0x00,
}
