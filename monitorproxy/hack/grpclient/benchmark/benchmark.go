package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"sync/atomic"
	"time"

	"github.com/golang/glog"
	"github.com/kayac/parallel-benchmark/benchmark"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	client "sdn.io/sdwan/hack/grpclient"
	pb "sdn.io/sdwan/pkg/monitorproxy/cubs/v1/metrics"
)

var (
	sites                              uint64
	durations                          uint
	reportPeriod, livePeriod           uint
	serverURL                          string
	reportSleepPeriod, liveSleepPeriod uint
	step                               uint
	liveEachToSend, reportEachToSend   uint
	liveNeedToSend, reportNeedToSend   uint32
	sumLiveSent, sumReportSend         uint32
	totalResetConn                     uint32
	cert                               credentials.TransportCredentials
)

func init() {
	flag.Uint64Var(&sites, "sites", 1, "The gRPC site numbers")
	flag.UintVar(&durations, "duration", 10, "Duration of benchmark")
	flag.UintVar(&reportPeriod, "report_period", 5, "Report period")
	flag.UintVar(&livePeriod, "live_period", 1, "Live period")
	flag.StringVar(&serverURL, "server", "", "gRPC server to connect with (ip:port)")

	flag.Set("logtostderr", "true")
}

type myWorker struct {
	ID                             int
	LiveSendCount, ReportSendCount uint32
	SleepCount                     uint
	clientConn                     *grpc.ClientConn
	cubsClient                     pb.CubsClient
}

func (w *myWorker) Setup() {
	// Setup connection
	conn := client.MustConnect(serverURL)

	// Each worker (client) owns a gRPC connection
	w.clientConn = conn
	w.cubsClient = pb.NewCubsClient(conn)
}

func (w *myWorker) Teardown() {
}

func genDeviceSn() (sn string) {
	return "S172L34180009"
}

func genInterfaceStat() []*pb.InterfaceStat {
	// Generate 2 InterfaceStat.
	wans := []string{"WAN1", "WAN2"}
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
	return intfStat
}

func genTrafficVolume() []*pb.TrafficVolume {
	// Generate 2 TrafficVolume.

	wans := []string{"WAN1", "WAN2"}
	len := len(wans)
	tv := make([]*pb.TrafficVolume, len)

	for i := 0; i < len; i++ {
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

	return tv
}

func genVPNConnInfo() []*pb.VPNConnInfo {
	// Generate 2x2 UnderlayQuality.
	vpns := []string{"sa00005WAN1-00004WAN1", "sa00005WAN2-00004WAN2", "sa00005WAN1-00004WAN2", "sa00005WAN2-00004WAN1"}
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

	return vpnConnInfos
}

func genTrafficTuple() []*pb.TrafficTuple {
	// total users = 185,000, 200 sessions/user -> 200 * 185,000 = 37,000,000 sessions
	// total devices = 4,640
	// 37,000,000 / 4,640 ~= 7,975 session/device
	len := 7975
	tts := make([]*pb.TrafficTuple, len)

	for i := 0; i < len; i++ {
		tts[i] = &pb.TrafficTuple{
			ClientIp:              []byte("wKh7KA=="),
			ClientPort:            65534,
			Proto:                 pb.L4Protocol(0),
			RemoteIp:              []byte("wKh7Jw=="),
			RemotePort:            65535,
			EstablishAt:           1524543923,
			ReportAt:              1524543923,
			Application:           1,
			ClientHostname:        "test-pc",
			ClientUsername:        "test",
			ClientOperatingSystem: "windows",
			Qos:           2,
			Status:        1,
			TrafficVolume: genTrafficVolume(),
		}
	}
	return tts
}

func (w *myWorker) Process() (subscore int) {
	// Generate mock data
	sn := genDeviceSn()

	intfStat := genInterfaceStat()
	mLiveReport := &pb.LiveReport{
		DeviceSn:   sn,
		Interfaces: intfStat,
	}

	vpnConnInfo := genVPNConnInfo()
	mVPNTrafficInfo := &pb.VPNTrafficInfo{
		DeviceSn: sn,
		ConnInfo: vpnConnInfo,
	}

	tuple := genTrafficTuple()
	mAccuTraffic := &pb.AccumulatedTraffic{
		DeviceSn:     sn,
		TrafficTuple: tuple,
	}

	ticker := time.NewTicker(time.Duration(durations) * time.Second)
	defer ticker.Stop()

	// Random start time in order to seperates different sites start reporting time.
	startTime := rand.Intn(int(reportPeriod))
	time.Sleep(time.Duration(startTime) * time.Millisecond)

	for {
		select {
		default:
			// Live //
			go func() {
				c := w.cubsClient

				_, err := c.ReportLiveInfo(context.Background(), mLiveReport)
				if err != nil {
					glog.Errorf("Worker id: %v ReportLiveInfo err: %v", w.ID, err)
					atomic.AddUint32(&totalResetConn, 1)
					return
				}
				//glog.Infof("Worker id: %v sent live", w.ID)
				atomic.AddUint32(&w.LiveSendCount, 1)
				atomic.AddUint32(&sumLiveSent, 1)
			}()

			// Periodically //
			if w.SleepCount%step == 0 {
				go func() {
					c := w.cubsClient

					_, err := c.ReportVPNTrafficInfo(context.Background(), mVPNTrafficInfo)
					if err != nil {
						glog.Errorf("Worker id: %v ReportVPNTrafficInfo err: %v", w.ID, err)
						atomic.AddUint32(&totalResetConn, 1)
						return
					}

					_, err = c.ReportTrafficInfo(context.Background(), mAccuTraffic)
					if err != nil {
						glog.Errorf("Worker id: %v ReportTrafficInfo err: %v", w.ID, err)
						atomic.AddUint32(&totalResetConn, 1)
						return
					}

					//glog.Infof("Worker id: %v sent periodically", w.ID)
					atomic.AddUint32(&w.ReportSendCount, 1)
					atomic.AddUint32(&sumReportSend, 1)
				}()
			}
			w.SleepCount++

			time.Sleep(time.Duration(liveSleepPeriod) * time.Millisecond)
		}
		select {
		case <-ticker.C:
			//log.Print("goroutine timeout")
			return 1
		default:
		}
	}
}

func main() {
	flag.Parse()

	liveSleepPeriod = livePeriod * 1000
	reportSleepPeriod = reportPeriod * 1000
	step = reportPeriod / livePeriod
	liveEachToSend = durations / livePeriod
	reportEachToSend = durations / reportPeriod
	liveNeedToSend = uint32(sites * uint64(liveEachToSend))
	reportNeedToSend = uint32(sites * uint64(reportEachToSend))

	if step == 0 {
		glog.Warningf("Invalid report period %v should larger or equal to live period %v", reportPeriod, livePeriod)
		os.Exit(1)
	}

	glog.Infof("Live goroutine sleep %d (ms) Report goroutine sleep %d (ms)", liveSleepPeriod, reportSleepPeriod)
	glog.Infof("Report send every %v steps", step)
	glog.Infof("Live each worker to sent %v ; Report each worker to sent %v", liveEachToSend, reportEachToSend)
	glog.Infof("Total needs to sent %v for live report, sent %v for periodically report", liveNeedToSend, reportNeedToSend)

	glog.Infof("Create %v workers to simulates %v sites", sites, sites)
	workers := make([]benchmark.Worker, sites)

	// Generates random seeds
	rand.Seed(time.Now().UnixNano())

	// Create #sites workers with goroutine
	for i := range workers {
		workers[i] = &myWorker{ID: i}
	}

	// Run all goroutine
	result := benchmark.Run(workers, time.Duration(durations)*time.Second)

	glog.Infof("sumLiveSent: %v ; sumReportSend: %v", sumLiveSent, sumReportSend)
	glog.Errorf("totalResetConn: %v", totalResetConn)
	if sumLiveSent != liveNeedToSend {
		glog.Errorf("sumLiveSent %v not equals to liveNeedToSend %v", sumLiveSent, liveNeedToSend)
	}
	if sumReportSend != reportNeedToSend {
		glog.Errorf("sumReportSend %v not equals to reportNeedToSend %v", sumReportSend, reportNeedToSend)
	}

	glog.Info(result)
}
