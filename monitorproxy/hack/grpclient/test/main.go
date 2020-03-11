package main

import (
	"context"
	"flag"
	"fmt"
	pbevent "monitorproxy/events"
	client "monitorproxy/hack/grpclient"
	pb "monitorproxy/metrics"
	"strconv"
	"strings"

	"github.com/golang/glog"
	json "github.com/golang/protobuf/jsonpb"

	//m "metrics.go"
	//"fmt"

	"github.com/gogo/protobuf/proto"

	"reflect"
)

var (
	serverURL = flag.String("server", "", "gRPC server to connect with (ip:port)")
	object    = flag.String("object", "", "json string of object")
	action    = flag.String("action", "", "action to be done")
)

// IPChangedJSON func return formatted JSON
func IPChangedJSON() {
	var v pbevent.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// IPChanged send the IPChanged to grpc server
func IPChanged(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.Interfaces
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.IPChanged(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// UserLoginJSON func return formatted JSON
func UserLoginJSON() string {
	var v pbevent.UserInfo
	fmt.Println(exampleJSON(&v, &v))
	return exampleJSON(&v, &v)
}

// UserLogin send the UserLogin to grpc server
func UserLogin(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.UserInfo
	//err := json.UnmarshalString(s, &v)
	err := json.UnmarshalString(s, &v)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return err
	}
	_, err = client.UserLogin(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// UserLogoutJSON func return formatted JSON
func UserLogoutJSON() {
	var v pbevent.UserInfo
	fmt.Println(exampleJSON(&v, &v))
}

// UserLogout send the UserLogout to grpc server
func UserLogout(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.UserInfo
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.UserLogout(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// InterfaceStatusChangeJSON func return formatted JSON
func InterfaceStatusChangeJSON() {
	var v pbevent.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// InterfaceStatusChange send the InterfaceStatusChange to grpc server
func InterfaceStatusChange(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.Interfaces
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.InterfaceStatusChange(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// LinkStateChangeJSON func return formatted JSON
func LinkStateChangeJSON() {
	var v pbevent.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// LinkStateChange send the LinkStateChange to grpc server
func LinkStateChange(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.Interfaces
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.LinkStateChange(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// PortStateChangeJSON func return formatted JSON
func PortStateChangeJSON() {
	var v pbevent.Ports
	fmt.Println(exampleJSON(&v, &v))
}

// PortStateChange send the PortStateChange to grpc server
func PortStateChange(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.Ports
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.PortStateChange(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// OnVPNEventJSON func return formatted JSON
func OnVPNEventJSON() {
	var v pbevent.VPNEvent
	fmt.Println(exampleJSON(&v, &v))
}

// OnVPNEvent send the VPNEvent to grpc server
func OnVPNEvent(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.VPNEvent
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.OnVPNEvent(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

// ReportDNSAnswerJSON func return formatted JSON
func ReportDNSAnswerJSON() {
	var v pbevent.DnsAnswer
	fmt.Println(exampleJSON(&v, &v))
}

// ReportDNSAnswer send the DNSAnswer to grpc server
func ReportDNSAnswer(client pbevent.CubsEventReportClient, s string) error {
	var v pbevent.DnsAnswer
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	stream, err := client.ReportDNSAnswer(context.Background())
	if err != nil {
		return err
	}

	ctx := stream.Context()

	if err := stream.Send(&v); err != nil {
		glog.Errorf("can not send %v", err)
	}

	glog.Infof("%v sent", &v)

	if err := stream.CloseSend(); err != nil {
		glog.Errorf("Close send error %v", err)
	}

	glog.Info("Stream closed.")

	ctx.Done()

	if err := ctx.Err(); err != nil {
		glog.Error(err)
	}

	return nil
}

func encodeElem(v reflect.Value) {
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			vf := v.Field(i)
			if vf.Kind() == reflect.Slice {
				vf.Set(reflect.MakeSlice(vf.Type(), 1, 1))
				vf.SetLen(1)
				if strings.Contains(vf.Type().String(), "*") {
					vf.Index(0).Set(reflect.New(vf.Index(0).Type().Elem()))
					encodeElem(vf.Index(0).Elem())
				} else {
					encodeElem(vf.Index(0))
				}
			} else if vf.Kind() == reflect.Struct {
				encodeElem(vf)
			} else if vf.Kind() == reflect.Ptr {
				vf.Set(reflect.New(vf.Type().Elem()))
				encodeElem(vf.Elem())
			}
		}
	}
}

func exampleJSON(v interface{}, pb proto.Message) string {
	vv := reflect.ValueOf(v)
	ve := vv.Elem()
	encodeElem(ve)

	m := &json.Marshaler{EmitDefaults: true, Indent: "   ", OrigName: true}
	js, _ := m.MarshalToString(pb)
	fmt.Println(js)
	fmt.Println()
	m = &json.Marshaler{EmitDefaults: true, OrigName: true}
	js, _ = m.MarshalToString(pb)
	return fmt.Sprintf("\"%s\"", strings.Replace(string(js), "\"", "'", -1))
}
func ReportTrafficInfoJSON() {
	v := pb.AccumulatedTraffic{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportLinkQualityJSON() {
	v := pb.LinkQuality{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportLiveInfoJSON() {
	v := pb.LiveReport{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportVPNTrafficInfoJSON() {
	v := pb.VPNTrafficInfo{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportSystemLoadJSON() {
	v := pb.SystemLoad{}
	fmt.Println(exampleJSON(&v, &v))
}

func init() {
	flag.Set("logtostderr", "true")
}
func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
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
		//return
	}

	//pb.MustRegister()

	conn := client.MustConnect(*serverURL)
	defer conn.Close()

	//cubsClient := pb.NewCubsClient(conn)
	cubsEventClient := pbevent.NewCubsEventReportClient(conn)

	var err error
	//s := UserLoginJSON()

	//fmt.Println("s before:", s)
	s := "{\"login_name\":\"\",\"login_time\":\"0\",\"ipv4_addr\":[null],\"ipv6_addr\":[null],\"device_sn\":\"\",\"timestamp\":\"1970-01-01T00:00:00Z\",\"user_group\":\"\"}"
	//s = strings.Replace(s, "'", "\"", -1)
	fmt.Println("s after:", s)

	switch *action {

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
		//fmt.Println("do userlogin")
		err = UserLogin(cubsEventClient, s)
	case "UserLogout":
		err = UserLogout(cubsEventClient, s)

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
