package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	//"sdn.io/sdwan/cmd/cubs/monitorproxy/apiclient"

	"monitorproxy/grpc"
)

var (
	podname   string
	buildTime string
	commitID  string
	gitBranch string
	gittag    string
	gitstatus string

	ws                            string
	ca                            string
	cert, key                     string
	port                          int
	metricLog, eventLog, alertLog string
)

func showversion() {
	podname, _ = os.Hostname()
	versionData := "Podname : " + podname + " BuildTime :" + buildTime + " CommitID : " + commitID + " GitBranch : " + gitBranch + " Gittag :" + gittag + " GitStatus :" + gitstatus + "\n"
	glog.Infof("Gitversion: %v", versionData)
}

func init() {
	//ws = os.Getenv("GOPATH") + "/src/sdn.io/sdwan"
	ws = "/home/zyxel/vicky/zyxelProjects/monitorproxy"
	flag.StringVar(&cert, "cert", ws+"/certs/mycerts/server.pem", "The TLS cert file")
	flag.StringVar(&key, "key", ws+"/certs/mycerts/server-key.pem", "The TLS key file")
	flag.StringVar(&ca, "ca", ws+"/certs/mycerts/ca.pem", "The CA cert file")
	flag.IntVar(&port, "port", 10001, "The server port")
	flag.StringVar(&metricLog, "metric", "grpc/testdata/metrics_c.log", "Metric log file")
	flag.StringVar(&eventLog, "event", "grpc/testdata/events_c.log", "Event log file")
	flag.StringVar(&alertLog, "alert", "grpc/testdata/alerts_c.log", "Alert log file")

	flag.Set("logtostderr", "true")
}

func setupSignalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		glog.Info("GracefulStopServer")
		grpc.GracefulStopServer()
		grpc.StopCronJob()
		os.Exit(1)
	}()
}

func setupSigusr1forDumpStack() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
}

func dumpStacks() {

	buf := make([]byte, 1638400)
	buf = buf[:runtime.Stack(buf, true)]
	glog.Infof("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", string(buf[:]))
}
func main() {

	flag.Parse()
	showversion()

	//setupSigusr1forDumpStack()
	//setupSignalHandler()
	_, err := grpc.StartClient(cert, key, ca, port)
	if err != nil {
		fmt.Println("deviceClent:", err)
	}
	//grpc.StartClient(cert, key, ca, port, metricLog, eventLog)
}
