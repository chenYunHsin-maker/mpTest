package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/golang/glog"

	pbevent "moniproxy/events"
	pb "moniproxy/metrics"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	ws                  string
	ca                  string
	sCert, sKey         string
	cCert, cKey         string
	port                int
	metricLog, eventLog string
)

func init() {
	ws = os.Getenv("GOPATH") + "/src/sdn.io/sdwan"
	flag.StringVar(&ca, "ca", ws+"/certs/mycerts/ca.pem", "The CA cert file")
	flag.StringVar(&sCert, "s_cert", ws+"/certs/mycerts/server.pem", "The TLS server cert file")
	flag.StringVar(&sKey, "s_key", ws+"/certs/mycerts/server-key.pem", "The TLS server key file")
	flag.StringVar(&cCert, "c_cert", ws+"/certs/mycerts/client.pem", "The TLS client cert file")
	flag.StringVar(&cKey, "c_key", ws+"/certs/mycerts/client-key.pem", "The TLS client key file")
	flag.IntVar(&port, "port", 10000, "The server port")
	flag.StringVar(&metricLog, "metric", "testdata/metrics.log", "Metric log file")
	flag.StringVar(&eventLog, "event", "testdata/events.log", "Event log file")

	flag.Parse()

	flag.Set("logtostderr", "true")
}

const (
	sn = "S172L34180009"
)

func setupGrpcServer() {
	go StartServer(sCert, sKey, ca, port, metricLog, eventLog)
}

func stopGrpcServer() {
	GracefulStopServer()
}

func setupGrpcClient() (*grpc.ClientConn, error) {
	certificate, err := tls.LoadX509KeyPair(cCert, cKey)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(ca)
	if err != nil {
		glog.Fatalf("failed to read ca cert: %s", err)
		return nil, err
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		glog.Fatal("failed to append certs")
		return nil, err
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "127.0.0.1",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	dialOption := grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), dialOption)
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
		return nil, err
	}

	return conn, nil
}

func TestReportTrafficInfo(t *testing.T) {
	// Start gRPC server
	setupGrpcServer()
	defer stopGrpcServer()

	// Start gRPC client
	conn, err := setupGrpcClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCubsClient(conn)

	// Generate mock data
	wans := []string{"WAN1", "WAN2"} // Generate 2 TrafficVolume.
	tv := make([]*pb.TrafficVolume, len(wans))
	for i := 0; i < len(wans); i++ {
		tv[i] = &pb.TrafficVolume{
			Interface: wans[i],
			Metrics: &pb.TrafficMetrics{
				TxPackets: 2,
				TxOctets:  240,
				RxPackets: 1,
				RxOctets:  120,
			},
		}
	}

	nclient := 1 // Only 1 client
	tts := make([]*pb.TrafficTuple, nclient)
	for i := 0; i < nclient; i++ {
		tts[i] = &pb.TrafficTuple{
			ClientIp:              []byte("wKh7KA=="),
			ClientPort:            0,
			Proto:                 pb.L4Protocol(0),
			RemoteIp:              []byte("wKh7Jw=="),
			RemotePort:            0,
			EstablishAt:           1524543923,
			ReportAt:              1524543923,
			Application:           0,
			ClientHostname:        "",
			ClientUsername:        "",
			ClientOperatingSystem: "",
			Qos:                   2,
			Status:                0,
			TrafficVolume:         tv,
		}
	}

	m := &pb.AccumulatedTraffic{DeviceSn: sn, TrafficTuple: tts}

	// Send to gRPC server
	_, err = client.ReportTrafficInfo(context.Background(), m)
	if err != nil {
		t.Fatalf("%v.ReportTrafficInfo(_) = _, %v: ", client, err)
	}
}

func TestReportLinkQuality(t *testing.T) {
	// Start gRPC server
	setupGrpcServer()
	defer stopGrpcServer()

	// Start gRPC client
	conn, err := setupGrpcClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCubsClient(conn)

	u := make([]*pb.PerLinkQuality, 2)
	p := make([]*pb.PeerLinkQuality, 1)
	linkQuailityPara2 := &pb.LinkQualityParameters{
		JitterMilliseconds:        6.3,
		PacketLossRate:            20,
		RoundtripTimeMilliseconds: 20.6,
	}
	linkQuaility1 := &pb.PerLinkQuality{
		LinkName:          "vti00025WAN1-00024WAN1",
		QualityParameters: nil,
		Failure:           true,
	}
	linkQuaility2 := &pb.PerLinkQuality{
		LinkName:          "vti00025WAN1-00024WAN3",
		QualityParameters: linkQuailityPara2,
		Failure:           false,
	}

	u[0] = linkQuaility1
	u[1] = linkQuaility2
	p[0] = &pb.PeerLinkQuality{
		PeerName:       "00024",
		PerLinkQuality: u,
	}
	o, _ := ptypes.TimestampProto(time.Now())
	l := &pb.LinkQuality{DeviceSn: sn, Peer: p, Timestamp: o}

	_, err = client.ReportLinkQuality(context.Background(), l)
	if err != nil {
		t.Fatalf("%v.ReportLinkQuality(_) = _, %v: ", client, err)
	}
}

func TestReportLiveInfo(t *testing.T) {
	// Start gRPC server
	setupGrpcServer()
	defer stopGrpcServer()

	// Start gRPC client
	conn, err := setupGrpcClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCubsClient(conn)

	// Generate mock data
	wans := []string{"WAN1", "WAN2"} // Generate 2 InterfaceStat.
	len := len(wans)
	intfStat := make([]*pb.InterfaceStat, len)
	for i := 0; i < len; i++ {
		intfStat[i] = &pb.InterfaceStat{
			Name: wans[i],
			Metrics: &pb.InterfaceMetrics{
				JitterMilliseconds: 20,
				PacketLost:         1,
			},
			IcmpMetrics: &pb.L4ProtocolMetrics{
				InOctets:   1,
				InPackets:  2,
				OutOctets:  10,
				OutPackets: 20,
			},
			TcpMetrics: &pb.L4ProtocolMetrics{
				InOctets:   10,
				InPackets:  20,
				OutOctets:  100,
				OutPackets: 200,
			},
			UdpMetrics: &pb.L4ProtocolMetrics{
				InOctets:   5,
				InPackets:  10,
				OutOctets:  500,
				OutPackets: 1000,
			},
			OtherMetrics: &pb.L4ProtocolMetrics{
				InOctets:   6,
				InPackets:  12,
				OutOctets:  60,
				OutPackets: 120,
			},
		}
	}

	m := &pb.LiveReport{DeviceSn: sn, Interfaces: intfStat}

	// Send to gRPC server
	_, err = client.ReportLiveInfo(context.Background(), m)
	if err != nil {
		t.Fatalf("%v.ReportLiveInfo(_) = _, %v: ", client, err)
	}
}

func TestReportVPNTrafficInfo(t *testing.T) {
	// Start gRPC server
	setupGrpcServer()
	defer stopGrpcServer()

	// Start gRPC client
	conn, err := setupGrpcClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCubsClient(conn)

	// Generate mock data
	vpns := []string{
		"sa00005WAN1-00004WAN1", "sa00005WAN2-00004WAN2",
		"sa00005WAN1-00004WAN2", "sa00005WAN2-00004WAN1"} // Generate 2x2 UnderlayQuality.
	len := len(vpns)
	vpnConnInfos := make([]*pb.VPNConnInfo, len)
	for i := 0; i < len; i++ {
		vpnConnInfos[i] = &pb.VPNConnInfo{
			ConnName: vpns[i],
			Ingress: &pb.PacketAndOctetMetrics{
				Packets: 1106,
				Octets:  42028,
			},
			Egress: &pb.PacketAndOctetMetrics{
				Packets: 1105,
				Octets:  42326,
			},
		}
	}

	m := &pb.VPNTrafficInfo{DeviceSn: sn, ConnInfo: vpnConnInfos}

	// Send to gRPC server
	_, err = client.ReportVPNTrafficInfo(context.Background(), m)
	if err != nil {
		t.Fatalf("%v.ReportVPNTrafficInfo(_) = _, %v: ", client, err)
	}
}

func TestReportIPChanged(t *testing.T) {
	// Start gRPC server
	setupGrpcServer()
	defer stopGrpcServer()

	// Start gRPC client
	conn, err := setupGrpcClient()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pbevent.NewCubsEventReportClient(conn)

	// Generate mock data
	wans := []string{"WAN1", "WAN2"} // Generate 2 InterfaceStat.
	ipv4s := [][][]byte{
		[][]byte{{192, 168, 1, 1}},
		[][]byte{{10, 0, 0, 1}},
	}
	globalIPv4s := [][][]byte{
		[][]byte{{8, 8, 8, 8}},
		[][]byte{{8, 8, 4, 4}},
	}
	len := len(wans)
	intfs := make([]*pbevent.Interface, len)
	for i := 0; i < len; i++ {
		intfs[i] = &pbevent.Interface{
			InterfaceName:  wans[i],
			Ipv4Addr:       ipv4s[i],
			State:          pbevent.InterfaceOperatingState_InterfaceStateUp,
			GlobalIpv4Addr: globalIPv4s[i],
		}
	}

	m := &pbevent.Interfaces{Interface: intfs, DeviceSn: sn}

	// Send to gRPC server
	_, err = client.IPChanged(context.Background(), m)
	if err != nil {
		t.Fatalf("%v.ReportVPNTrafficInfo(_) = _, %v: ", client, err)
	}
}
