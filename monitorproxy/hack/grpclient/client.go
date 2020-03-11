package client

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"os"

	"github.com/golang/glog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	ws   string
	ca   string
	cert string
	key  string
	port int
)

func init() {
	ws = os.Getenv("GOPATH") + "/src/sdn.io/sdwan"
	flag.StringVar(&ca, "ca", ws+"/certs/mycerts/ca.pem", "The CA cert file")
	flag.StringVar(&cert, "cert", ws+"/certs/mycerts/client.pem", "The TLS client cert file")
	flag.StringVar(&key, "key", ws+"/certs/mycerts/client-key.pem", "The TLS client key file")
	flag.IntVar(&port, "port", 10000, "The server port")
}

// MustConnect connects to the target gRPC server
func MustConnect(target string) *grpc.ClientConn {
	certificate, err := tls.LoadX509KeyPair(cert, key)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(ca)
	if err != nil {
		glog.Fatalf("failed to read ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		glog.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		//ServerName:   "localhost",
		ServerName:   "S172L34180009",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	dialOption := grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial(target, dialOption)
	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}
	return conn
}
