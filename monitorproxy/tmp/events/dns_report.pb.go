// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns_report.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DNSReport struct {
	Success   bool                       `protobuf:"varint,1,opt,name=success" json:"success"`
	TaskId    string                     `protobuf:"bytes,2,opt,name=task_id,json=taskId" json:"task_id"`
	Answer    string                     `protobuf:"bytes,3,opt,name=answer" json:"answer"`
	Msg       string                     `protobuf:"bytes,4,opt,name=msg" json:"msg"`
	DeviceSn  string                     `protobuf:"bytes,5,opt,name=device_sn,json=deviceSn" json:"device_sn"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=timestamp" json:"timestamp"`
}

func (m *DNSReport) Reset()                    { *m = DNSReport{} }
func (m *DNSReport) String() string            { return proto.CompactTextString(m) }
func (*DNSReport) ProtoMessage()               {}
func (*DNSReport) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *DNSReport) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DNSReport) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DNSReport) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *DNSReport) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DNSReport) GetDeviceSn() string {
	if m != nil {
		return m.DeviceSn
	}
	return ""
}

func (m *DNSReport) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*DNSReport)(nil), "events.DNSReport")
}

func init() { proto.RegisterFile("dns_report.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x65, 0x16, 0xbc, 0xeb, 0xa1, 0x59, 0xb9, 0x00, 0x6b, 0x29, 0x88, 0xa8, 0x52, 0x79,
	0x25, 0x68, 0x38, 0x00, 0x0d, 0x0d, 0x85, 0x43, 0x1f, 0x25, 0xf1, 0x10, 0x45, 0x10, 0x3b, 0xf2,
	0x38, 0xe1, 0x78, 0x5c, 0x0d, 0xc5, 0x56, 0xa0, 0xf3, 0x7f, 0x7e, 0x23, 0x3d, 0x38, 0x5a, 0x47,
	0x75, 0xc0, 0xc9, 0x87, 0xa8, 0xa7, 0xe0, 0xa3, 0x97, 0x1c, 0x17, 0x74, 0x91, 0x4e, 0xf7, 0xbd,
	0xf7, 0xfd, 0x17, 0x9e, 0x13, 0x6d, 0xe7, 0x8f, 0x73, 0x1c, 0x46, 0xa4, 0xd8, 0x8c, 0x53, 0x16,
	0x1f, 0x7e, 0x18, 0x88, 0x97, 0xb7, 0xca, 0xa4, 0x63, 0xa9, 0x60, 0x4f, 0x73, 0xd7, 0x21, 0x91,
	0x62, 0x05, 0x2b, 0x0f, 0x66, 0x9b, 0xf2, 0x16, 0xf6, 0xb1, 0xa1, 0xcf, 0x7a, 0xb0, 0xea, 0xa2,
	0x60, 0xa5, 0x30, 0x7c, 0x9d, 0xaf, 0x56, 0xde, 0x00, 0x6f, 0x1c, 0x7d, 0x63, 0x50, 0xbb, 0xcc,
	0xf3, 0x92, 0x47, 0xd8, 0x8d, 0xd4, 0xab, 0xcb, 0x04, 0xd7, 0xa7, 0xbc, 0x03, 0x61, 0x71, 0x19,
	0x3a, 0xac, 0xc9, 0xa9, 0xab, 0xc4, 0x0f, 0x19, 0x54, 0x4e, 0x3e, 0x83, 0xf8, 0x4b, 0x53, 0xbc,
	0x60, 0xe5, 0xf5, 0xe3, 0x49, 0xe7, 0x78, 0xbd, 0xc5, 0xeb, 0xf7, 0xcd, 0x30, 0xff, 0x72, 0xcb,
	0xd3, 0xf7, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xdc, 0xaf, 0xd7, 0x05, 0x01, 0x00,
	0x00,
}
