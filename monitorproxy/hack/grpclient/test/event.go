package main

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	json "github.com/golang/protobuf/jsonpb"

	pb "sdn.io/sdwan/pkg/monitorproxy/cubs/v1/events"
)

// IPChangedJSON func return formatted JSON
func IPChangedJSON() {
	var v pb.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// IPChanged send the IPChanged to grpc server
func IPChanged(client pb.CubsEventReportClient, s string) error {
	var v pb.Interfaces
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
func UserLoginJSON() {
	var v pb.UserInfo
	fmt.Println(exampleJSON(&v, &v))
}

// UserLogin send the UserLogin to grpc server
func UserLogin(client pb.CubsEventReportClient, s string) error {
	var v pb.UserInfo
	err := json.UnmarshalString(s, &v)
	if err != nil {
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
	var v pb.UserInfo
	fmt.Println(exampleJSON(&v, &v))
}

// UserLogout send the UserLogout to grpc server
func UserLogout(client pb.CubsEventReportClient, s string) error {
	var v pb.UserInfo
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
	var v pb.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// InterfaceStatusChange send the InterfaceStatusChange to grpc server
func InterfaceStatusChange(client pb.CubsEventReportClient, s string) error {
	var v pb.Interfaces
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
	var v pb.Interfaces
	fmt.Println(exampleJSON(&v, &v))
}

// LinkStateChange send the LinkStateChange to grpc server
func LinkStateChange(client pb.CubsEventReportClient, s string) error {
	var v pb.Interfaces
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
	var v pb.Ports
	fmt.Println(exampleJSON(&v, &v))
}

// PortStateChange send the PortStateChange to grpc server
func PortStateChange(client pb.CubsEventReportClient, s string) error {
	var v pb.Ports
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
	var v pb.VPNEvent
	fmt.Println(exampleJSON(&v, &v))
}

// OnVPNEvent send the VPNEvent to grpc server
func OnVPNEvent(client pb.CubsEventReportClient, s string) error {
	var v pb.VPNEvent
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
	var v pb.DnsAnswer
	fmt.Println(exampleJSON(&v, &v))
}

// ReportDNSAnswer send the DNSAnswer to grpc server
func ReportDNSAnswer(client pb.CubsEventReportClient, s string) error {
	var v pb.DnsAnswer
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
