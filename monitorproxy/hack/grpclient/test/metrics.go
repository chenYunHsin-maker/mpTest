package main

import (
	"context"
	"fmt"

	json "github.com/golang/protobuf/jsonpb"
	pb "sdn.io/sdwan/pkg/monitorproxy/cubs/v1/metrics"
)

func ReportTrafficInfoJSON() {
	v := pb.AccumulatedTraffic{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportTrafficInfo(client pb.CubsClient, s string) error {
	var v pb.AccumulatedTraffic
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.ReportTrafficInfo(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

func ReportLinkQualityJSON() {
	v := pb.LinkQuality{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportLinkQuality(client pb.CubsClient, s string) error {
	var v pb.LinkQuality
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.ReportLinkQuality(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

func ReportLiveInfoJSON() {
	v := pb.LiveReport{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportLiveInfo(client pb.CubsClient, s string) error {
	var v pb.LiveReport
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.ReportLiveInfo(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

func ReportVPNTrafficInfoJSON() {
	v := pb.VPNTrafficInfo{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportVPNTrafficInfo(client pb.CubsClient, s string) error {
	var v pb.VPNTrafficInfo
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.ReportVPNTrafficInfo(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}

func ReportSystemLoadJSON() {
	v := pb.SystemLoad{}
	fmt.Println(exampleJSON(&v, &v))
}

func ReportSystemLoad(client pb.CubsClient, s string) error {
	var v pb.SystemLoad
	err := json.UnmarshalString(s, &v)
	if err != nil {
		return err
	}
	_, err = client.ReportSystemLoad(context.Background(), &v)
	if err != nil {
		return err
	}
	return nil
}
