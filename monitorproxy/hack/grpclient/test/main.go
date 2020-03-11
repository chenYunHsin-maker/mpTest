package main

import (
	"flag"
	"strings"

	"github.com/golang/glog"

	client "sdn.io/sdwan/hack/grpclient"
	pbevent "sdn.io/sdwan/pkg/monitorproxy/cubs/v1/events"
	pb "sdn.io/sdwan/pkg/monitorproxy/cubs/v1/metrics"
)

var (
	serverURL = flag.String("server", "", "gRPC server to connect with (ip:port)")
	object    = flag.String("object", "", "json string of object")
	action    = flag.String("action", "", "action to be done")
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	if *object == "" {
		switch *action {
		case "ReportTrafficInfo":
			ReportTrafficInfoJSON()
		case "ReportLinkQuality":
			ReportLinkQualityJSON()
		case "ReportLiveInfo":
			ReportLiveInfoJSON()
		case "ReportVPNTrafficInfo":
			ReportVPNTrafficInfoJSON()
		case "IPChanged":
			IPChangedJSON()
		case "LinkStateChange":
			LinkStateChangeJSON()
		case "PortStateChange":
			PortStateChangeJSON()
		case "InterfaceStatusChange":
			InterfaceStatusChangeJSON()
		case "OnVPNEvent":
			OnVPNEventJSON()
		case "UserLogin":
			UserLoginJSON()
		case "UserLogout":
			UserLogoutJSON()
		case "ReportSystemLoad":
			ReportSystemLoadJSON()
		case "ReportDNSAnswer":
			ReportDNSAnswerJSON()
		default:
			glog.Errorf("unknown action %s", *action)
		}
		return
	}

	pb.MustRegister()

	conn := client.MustConnect(*serverURL)
	defer conn.Close()

	cubsClient := pb.NewCubsClient(conn)
	cubsEventClient := pbevent.NewCubsEventReportClient(conn)

	var err error
	s := strings.Replace(*object, "'", "\"", -1)

	switch *action {
	case "ReportTrafficInfo":
		err = ReportTrafficInfo(cubsClient, s)
	case "ReportLinkQuality":
		err = ReportLinkQuality(cubsClient, s)
	case "ReportLiveInfo":
		err = ReportLiveInfo(cubsClient, s)
	case "ReportVPNTrafficInfo":
		err = ReportVPNTrafficInfo(cubsClient, s)
	case "IPChanged":
		err = IPChanged(cubsEventClient, s)
	case "LinkStateChange":
		err = LinkStateChange(cubsEventClient, s)
	case "PortStateChange":
		err = PortStateChange(cubsEventClient, s)
	case "InterfaceStatusChange":
		err = InterfaceStatusChange(cubsEventClient, s)
	case "OnVPNEvent":
		err = OnVPNEvent(cubsEventClient, s)
	case "UserLogin":
		err = UserLogin(cubsEventClient, s)
	case "UserLogout":
		err = UserLogout(cubsEventClient, s)
	case "ReportSystemLoad":
		err = ReportSystemLoad(cubsClient, s)
	case "ReportDNSAnswer":
		err = ReportDNSAnswer(cubsEventClient, s)
	default:
		glog.Errorf("unknown action %s", *action)
		return
	}

	if err != nil {
		glog.Errorln(err.Error())
	} else {
		glog.Infoln("msg sent.")
	}
}
