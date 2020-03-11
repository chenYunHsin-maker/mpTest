// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vpn_event.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VPNEventType int32

const (
	VPNEventType_VPN_CONNECTED    VPNEventType = 0
	VPNEventType_VPN_DISCONNECTED VPNEventType = 1
	VPNEventType_VPN_REKEY        VPNEventType = 2
	VPNEventType_VPN_ERROR        VPNEventType = 3
	VPNEventType_VPN_INIT         VPNEventType = 4
)

var VPNEventType_name = map[int32]string{
	0: "VPN_CONNECTED",
	1: "VPN_DISCONNECTED",
	2: "VPN_REKEY",
	3: "VPN_ERROR",
	4: "VPN_INIT",
}
var VPNEventType_value = map[string]int32{
	"VPN_CONNECTED":    0,
	"VPN_DISCONNECTED": 1,
	"VPN_REKEY":        2,
	"VPN_ERROR":        3,
	"VPN_INIT":         4,
}

func (x VPNEventType) String() string {
	return proto.EnumName(VPNEventType_name, int32(x))
}
func (VPNEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

type PolicyMode int32

const (
	PolicyMode_PolicyModeRange  PolicyMode = 0
	PolicyMode_PolicyModeSubnet PolicyMode = 1
)

var PolicyMode_name = map[int32]string{
	0: "PolicyModeRange",
	1: "PolicyModeSubnet",
}
var PolicyMode_value = map[string]int32{
	"PolicyModeRange":  0,
	"PolicyModeSubnet": 1,
}

func (x PolicyMode) String() string {
	return proto.EnumName(PolicyMode_name, int32(x))
}
func (PolicyMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

type VPNPeerInfo struct {
	PeerAddr     string `protobuf:"bytes,1,opt,name=peer_addr,json=peerAddr" json:"peer_addr"`
	PeerIsServer bool   `protobuf:"varint,2,opt,name=peer_is_server,json=peerIsServer" json:"peer_is_server"`
}

func (m *VPNPeerInfo) Reset()                    { *m = VPNPeerInfo{} }
func (m *VPNPeerInfo) String() string            { return proto.CompactTextString(m) }
func (*VPNPeerInfo) ProtoMessage()               {}
func (*VPNPeerInfo) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *VPNPeerInfo) GetPeerAddr() string {
	if m != nil {
		return m.PeerAddr
	}
	return ""
}

func (m *VPNPeerInfo) GetPeerIsServer() bool {
	if m != nil {
		return m.PeerIsServer
	}
	return false
}

type VPNPolicy struct {
	PolicyMode PolicyMode `protobuf:"varint,1,opt,name=policy_mode,json=policyMode,enum=events.PolicyMode" json:"policy_mode"`
}

func (m *VPNPolicy) Reset()                    { *m = VPNPolicy{} }
func (m *VPNPolicy) String() string            { return proto.CompactTextString(m) }
func (*VPNPolicy) ProtoMessage()               {}
func (*VPNPolicy) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

func (m *VPNPolicy) GetPolicyMode() PolicyMode {
	if m != nil {
		return m.PolicyMode
	}
	return PolicyMode_PolicyModeRange
}

type IKE struct {
	SpiI      uint32 `protobuf:"varint,1,opt,name=spi_i,json=spiI" json:"spi_i"`
	SpiR      uint32 `protobuf:"varint,2,opt,name=spi_r,json=spiR" json:"spi_r"`
	Alg       string `protobuf:"bytes,3,opt,name=alg" json:"alg"`
	CipherKey uint64 `protobuf:"varint,4,opt,name=cipher_key,json=cipherKey" json:"cipher_key"`
}

func (m *IKE) Reset()                    { *m = IKE{} }
func (m *IKE) String() string            { return proto.CompactTextString(m) }
func (*IKE) ProtoMessage()               {}
func (*IKE) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

func (m *IKE) GetSpiI() uint32 {
	if m != nil {
		return m.SpiI
	}
	return 0
}

func (m *IKE) GetSpiR() uint32 {
	if m != nil {
		return m.SpiR
	}
	return 0
}

func (m *IKE) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *IKE) GetCipherKey() uint64 {
	if m != nil {
		return m.CipherKey
	}
	return 0
}

type IPSec struct {
	SpiI       uint32 `protobuf:"varint,1,opt,name=spi_i,json=spiI" json:"spi_i"`
	SpiO       uint32 `protobuf:"varint,2,opt,name=spi_o,json=spiO" json:"spi_o"`
	PeerHandle uint32 `protobuf:"varint,3,opt,name=peer_handle,json=peerHandle" json:"peer_handle"`
}

func (m *IPSec) Reset()                    { *m = IPSec{} }
func (m *IPSec) String() string            { return proto.CompactTextString(m) }
func (*IPSec) ProtoMessage()               {}
func (*IPSec) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{3} }

func (m *IPSec) GetSpiI() uint32 {
	if m != nil {
		return m.SpiI
	}
	return 0
}

func (m *IPSec) GetSpiO() uint32 {
	if m != nil {
		return m.SpiO
	}
	return 0
}

func (m *IPSec) GetPeerHandle() uint32 {
	if m != nil {
		return m.PeerHandle
	}
	return 0
}

type VPNConnInfo struct {
	TunnelName    string     `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName" json:"tunnel_name"`
	InterfaceName string     `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name"`
	VpnType       string     `protobuf:"bytes,3,opt,name=vpn_type,json=vpnType" json:"vpn_type"`
	LocalAddr     []byte     `protobuf:"bytes,4,opt,name=local_addr,json=localAddr,proto3" json:"local_addr"`
	LocalPort     uint32     `protobuf:"varint,5,opt,name=local_port,json=localPort" json:"local_port"`
	RemoteAddr    []byte     `protobuf:"bytes,6,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr"`
	RemotePort    uint32     `protobuf:"varint,7,opt,name=remote_port,json=remotePort" json:"remote_port"`
	LocalPolicy   *VPNPolicy `protobuf:"bytes,8,opt,name=local_policy,json=localPolicy" json:"local_policy"`
	RemotePolicy  *VPNPolicy `protobuf:"bytes,9,opt,name=remote_policy,json=remotePolicy" json:"remote_policy"`
	Ike           *IKE       `protobuf:"bytes,10,opt,name=ike" json:"ike"`
	Ipsec         *IPSec     `protobuf:"bytes,11,opt,name=ipsec" json:"ipsec"`
}

func (m *VPNConnInfo) Reset()                    { *m = VPNConnInfo{} }
func (m *VPNConnInfo) String() string            { return proto.CompactTextString(m) }
func (*VPNConnInfo) ProtoMessage()               {}
func (*VPNConnInfo) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{4} }

func (m *VPNConnInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *VPNConnInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *VPNConnInfo) GetVpnType() string {
	if m != nil {
		return m.VpnType
	}
	return ""
}

func (m *VPNConnInfo) GetLocalAddr() []byte {
	if m != nil {
		return m.LocalAddr
	}
	return nil
}

func (m *VPNConnInfo) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *VPNConnInfo) GetRemoteAddr() []byte {
	if m != nil {
		return m.RemoteAddr
	}
	return nil
}

func (m *VPNConnInfo) GetRemotePort() uint32 {
	if m != nil {
		return m.RemotePort
	}
	return 0
}

func (m *VPNConnInfo) GetLocalPolicy() *VPNPolicy {
	if m != nil {
		return m.LocalPolicy
	}
	return nil
}

func (m *VPNConnInfo) GetRemotePolicy() *VPNPolicy {
	if m != nil {
		return m.RemotePolicy
	}
	return nil
}

func (m *VPNConnInfo) GetIke() *IKE {
	if m != nil {
		return m.Ike
	}
	return nil
}

func (m *VPNConnInfo) GetIpsec() *IPSec {
	if m != nil {
		return m.Ipsec
	}
	return nil
}

type VPNEvent struct {
	EventType      VPNEventType               `protobuf:"varint,1,opt,name=event_type,json=eventType,enum=events.VPNEventType" json:"event_type"`
	ConnectionInfo *VPNConnInfo               `protobuf:"bytes,2,opt,name=connection_info,json=connectionInfo" json:"connection_info"`
	PeerInfo       *VPNPeerInfo               `protobuf:"bytes,3,opt,name=peer_info,json=peerInfo" json:"peer_info"`
	Uptime         uint64                     `protobuf:"varint,4,opt,name=uptime" json:"uptime"`
	Timeout        uint32                     `protobuf:"varint,5,opt,name=timeout" json:"timeout"`
	Msg            string                     `protobuf:"bytes,6,opt,name=msg" json:"msg"`
	DeviceSn       string                     `protobuf:"bytes,7,opt,name=device_sn,json=deviceSn" json:"device_sn"`
	Timestamp      *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=timestamp" json:"timestamp"`
}

func (m *VPNEvent) Reset()                    { *m = VPNEvent{} }
func (m *VPNEvent) String() string            { return proto.CompactTextString(m) }
func (*VPNEvent) ProtoMessage()               {}
func (*VPNEvent) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{5} }

func (m *VPNEvent) GetEventType() VPNEventType {
	if m != nil {
		return m.EventType
	}
	return VPNEventType_VPN_CONNECTED
}

func (m *VPNEvent) GetConnectionInfo() *VPNConnInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func (m *VPNEvent) GetPeerInfo() *VPNPeerInfo {
	if m != nil {
		return m.PeerInfo
	}
	return nil
}

func (m *VPNEvent) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *VPNEvent) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *VPNEvent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *VPNEvent) GetDeviceSn() string {
	if m != nil {
		return m.DeviceSn
	}
	return ""
}

func (m *VPNEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type VPNTestResult struct {
	ConnectionName string                     `protobuf:"bytes,1,opt,name=connection_name,json=connectionName" json:"connection_name"`
	DeviceSn       string                     `protobuf:"bytes,2,opt,name=device_sn,json=deviceSn" json:"device_sn"`
	Timestamp      *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp"`
}

func (m *VPNTestResult) Reset()                    { *m = VPNTestResult{} }
func (m *VPNTestResult) String() string            { return proto.CompactTextString(m) }
func (*VPNTestResult) ProtoMessage()               {}
func (*VPNTestResult) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{6} }

func (m *VPNTestResult) GetConnectionName() string {
	if m != nil {
		return m.ConnectionName
	}
	return ""
}

func (m *VPNTestResult) GetDeviceSn() string {
	if m != nil {
		return m.DeviceSn
	}
	return ""
}

func (m *VPNTestResult) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*VPNPeerInfo)(nil), "events.VPNPeerInfo")
	proto.RegisterType((*VPNPolicy)(nil), "events.VPNPolicy")
	proto.RegisterType((*IKE)(nil), "events.IKE")
	proto.RegisterType((*IPSec)(nil), "events.IPSec")
	proto.RegisterType((*VPNConnInfo)(nil), "events.VPNConnInfo")
	proto.RegisterType((*VPNEvent)(nil), "events.VPNEvent")
	proto.RegisterType((*VPNTestResult)(nil), "events.VPNTestResult")
	proto.RegisterEnum("events.VPNEventType", VPNEventType_name, VPNEventType_value)
	proto.RegisterEnum("events.PolicyMode", PolicyMode_name, PolicyMode_value)
}

func init() { proto.RegisterFile("vpn_event.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x6f, 0xda, 0x48,
	0x10, 0x3e, 0x63, 0x20, 0x78, 0x0c, 0x84, 0x6c, 0xa2, 0x93, 0x2f, 0xa7, 0x28, 0x88, 0xbb, 0xd3,
	0xa1, 0x3c, 0x90, 0x53, 0x72, 0x6a, 0xfb, 0xd0, 0x87, 0x56, 0x89, 0xa5, 0x5a, 0xa8, 0x04, 0x2d,
	0x28, 0x52, 0x9f, 0x5c, 0xc7, 0x1e, 0x88, 0x15, 0xb3, 0x6b, 0xd9, 0x06, 0x89, 0x5f, 0xd1, 0xff,
	0xd0, 0x9f, 0xd3, 0x5f, 0x55, 0xed, 0x2c, 0x06, 0xa2, 0xaa, 0x91, 0xfa, 0xe4, 0xf5, 0xb7, 0xdf,
	0xb7, 0x33, 0x3b, 0xdf, 0xec, 0xc0, 0xe1, 0x2a, 0x15, 0x3e, 0xae, 0x50, 0x14, 0x83, 0x34, 0x93,
	0x85, 0x64, 0x75, 0xfa, 0xc9, 0x4f, 0xcf, 0xe7, 0x52, 0xce, 0x13, 0xbc, 0x24, 0xf4, 0x61, 0x39,
	0xbb, 0x2c, 0xe2, 0x05, 0xe6, 0x45, 0xb0, 0x48, 0x35, 0xb1, 0x37, 0x06, 0xfb, 0x7e, 0x3c, 0x1a,
	0x23, 0x66, 0x9e, 0x98, 0x49, 0xf6, 0x27, 0x58, 0x29, 0x62, 0xe6, 0x07, 0x51, 0x94, 0x39, 0x46,
	0xd7, 0xe8, 0x5b, 0xbc, 0xa1, 0x80, 0xf7, 0x51, 0x94, 0xb1, 0xbf, 0xa1, 0x4d, 0x9b, 0x71, 0xee,
	0xe7, 0x98, 0xad, 0x30, 0x73, 0x2a, 0x5d, 0xa3, 0xdf, 0xe0, 0x4d, 0x85, 0x7a, 0xf9, 0x84, 0xb0,
	0xde, 0x3b, 0xb0, 0xd4, 0x89, 0x32, 0x89, 0xc3, 0x35, 0xbb, 0x06, 0x3b, 0xa5, 0x95, 0xbf, 0x90,
	0x11, 0xd2, 0x89, 0xed, 0x2b, 0x36, 0xd0, 0xd9, 0x0d, 0x34, 0xe9, 0xa3, 0x8c, 0x90, 0x43, 0xba,
	0x5d, 0xf7, 0x3e, 0x83, 0xe9, 0x0d, 0x5d, 0x76, 0x0c, 0xb5, 0x3c, 0x8d, 0xfd, 0x98, 0x54, 0x2d,
	0x5e, 0xcd, 0xd3, 0xd8, 0x2b, 0x41, 0x1d, 0x5a, 0x83, 0x9c, 0x75, 0xc0, 0x0c, 0x92, 0xb9, 0x63,
	0x52, 0xbe, 0x6a, 0xc9, 0xce, 0x00, 0xc2, 0x38, 0x7d, 0xc4, 0xcc, 0x7f, 0xc2, 0xb5, 0x53, 0xed,
	0x1a, 0xfd, 0x2a, 0xb7, 0x34, 0x32, 0xc4, 0x75, 0x8f, 0x43, 0xcd, 0x1b, 0x4f, 0x30, 0x7c, 0x31,
	0x86, 0xdc, 0x8b, 0x71, 0xc7, 0xce, 0xc1, 0xa6, 0xcb, 0x3f, 0x06, 0x22, 0x4a, 0x90, 0x62, 0xb5,
	0x38, 0x28, 0xe8, 0x03, 0x21, 0xbd, 0xaf, 0x26, 0x95, 0xf2, 0x46, 0x0a, 0x41, 0xa5, 0x3c, 0x07,
	0xbb, 0x58, 0x0a, 0x81, 0x89, 0x2f, 0x82, 0x05, 0x6e, 0x8a, 0x09, 0x1a, 0x1a, 0x05, 0x0b, 0x64,
	0xff, 0x40, 0x3b, 0x16, 0x05, 0x66, 0xb3, 0x20, 0x44, 0xcd, 0xa9, 0x10, 0xa7, 0xb5, 0x45, 0x89,
	0xf6, 0x07, 0x34, 0x94, 0xbb, 0xc5, 0x3a, 0xc5, 0xcd, 0x0d, 0x0f, 0x56, 0xa9, 0x98, 0xae, 0x53,
	0x54, 0xb7, 0x4c, 0x64, 0x18, 0x24, 0xda, 0x2e, 0x75, 0xcb, 0x26, 0xb7, 0x08, 0x21, 0xbf, 0xb6,
	0xdb, 0xa9, 0xcc, 0x0a, 0xa7, 0x46, 0x19, 0xeb, 0xed, 0xb1, 0xcc, 0x0a, 0x95, 0x60, 0x86, 0x0b,
	0x59, 0xa0, 0x96, 0xd7, 0x49, 0x0e, 0x1a, 0x22, 0xfd, 0x8e, 0x40, 0x07, 0x1c, 0xe8, 0x2b, 0x6b,
	0x88, 0x4e, 0xf8, 0x1f, 0x9a, 0x65, 0x00, 0x65, 0x9e, 0xd3, 0xe8, 0x1a, 0x7d, 0xfb, 0xea, 0xa8,
	0xb4, 0x77, 0xdb, 0x06, 0xdc, 0xde, 0x44, 0xa5, 0x9e, 0x78, 0x05, 0xad, 0xed, 0xb1, 0x24, 0xb3,
	0x7e, 0x26, 0x6b, 0x96, 0xb1, 0x48, 0x77, 0x06, 0x66, 0xfc, 0x84, 0x0e, 0x10, 0xdb, 0x2e, 0xd9,
	0xde, 0xd0, 0xe5, 0x0a, 0x67, 0x7f, 0x41, 0x2d, 0x4e, 0x73, 0x0c, 0x1d, 0x9b, 0x08, 0xad, 0x2d,
	0x41, 0x19, 0xcd, 0xf5, 0x5e, 0xef, 0x5b, 0x05, 0x1a, 0xf7, 0xe3, 0x91, 0xab, 0xb6, 0xd8, 0x35,
	0x00, 0x71, 0x74, 0x6d, 0x75, 0x6f, 0x9e, 0xec, 0x65, 0x41, 0x2c, 0x55, 0x68, 0x6e, 0x61, 0xb9,
	0x64, 0x6f, 0xe1, 0x30, 0x94, 0x42, 0x60, 0x58, 0xc4, 0x52, 0xf8, 0xb1, 0x98, 0xe9, 0x36, 0xb1,
	0xaf, 0x8e, 0xf7, 0x94, 0x65, 0x13, 0xf0, 0xf6, 0x8e, 0x4b, 0x4d, 0xf1, 0xdf, 0xe6, 0x7d, 0x91,
	0xce, 0xfc, 0x41, 0x57, 0xbe, 0x43, 0xfd, 0xe8, 0x48, 0xf1, 0x3b, 0xd4, 0x97, 0xa9, 0x7a, 0xb5,
	0x9b, 0x2e, 0xde, 0xfc, 0x31, 0x07, 0x0e, 0xd4, 0x57, 0x2e, 0x4b, 0x67, 0xcb, 0x5f, 0xf5, 0x1a,
	0x16, 0xf9, 0x9c, 0xfc, 0xb4, 0xb8, 0x5a, 0xaa, 0x57, 0x1d, 0xe1, 0x2a, 0x0e, 0xd1, 0xcf, 0x05,
	0xd9, 0x68, 0xf1, 0x86, 0x06, 0x26, 0x82, 0xbd, 0x01, 0x6b, 0x3b, 0x14, 0x36, 0x0e, 0x9e, 0x0e,
	0xf4, 0xd8, 0x18, 0x94, 0x63, 0x63, 0x30, 0x2d, 0x19, 0x7c, 0x47, 0xee, 0x7d, 0x31, 0xa0, 0x75,
	0x3f, 0x1e, 0x4d, 0x31, 0x2f, 0x38, 0xe6, 0xcb, 0xa4, 0x60, 0xff, 0x3e, 0x2b, 0xce, 0x5e, 0xdf,
	0xef, 0xd5, 0x81, 0x9a, 0xfa, 0x59, 0x46, 0x95, 0x97, 0x32, 0x32, 0x7f, 0x21, 0xa3, 0x8b, 0x10,
	0x9a, 0xfb, 0xbe, 0xb1, 0x23, 0x4a, 0xd0, 0xbf, 0xb9, 0x1b, 0x8d, 0xdc, 0x9b, 0xa9, 0x7b, 0xdb,
	0xf9, 0x8d, 0x9d, 0x40, 0x47, 0x41, 0xb7, 0xde, 0x64, 0x87, 0x1a, 0xac, 0x45, 0x43, 0xcb, 0xe7,
	0xee, 0xd0, 0xfd, 0xd4, 0xa9, 0x94, 0xbf, 0x2e, 0xe7, 0x77, 0xbc, 0x63, 0xb2, 0x26, 0x35, 0x8d,
	0xef, 0x8d, 0xbc, 0x69, 0xa7, 0x7a, 0xf1, 0x1a, 0x60, 0x37, 0xb8, 0xd8, 0x31, 0x1c, 0xee, 0x8d,
	0xb1, 0x40, 0xcc, 0x51, 0x07, 0xd9, 0x81, 0x93, 0xe5, 0x83, 0xc0, 0xa2, 0x63, 0x3c, 0xd4, 0x29,
	0xf9, 0xeb, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x42, 0xbd, 0x83, 0xae, 0x05, 0x00, 0x00,
}
