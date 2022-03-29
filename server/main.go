package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	pb "dewi.atop/learn/grpc/admin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type dataAdminServer struct {
	pb.UnimplementedDataAdminServer
	mu     sync.Mutex
	admins []*pb.Admin
}

func (d *dataAdminServer) FindAdminByEmail(ctx context.Context, admin *pb.Admin) (*pb.Admin, error) {
	fmt.Println("Incoming request")

	for _, v := range d.admins {
		if v.Email == admin.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataAdminServer) loadData() {
	data, err := ioutil.ReadFile("data/datas.json")
	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.admins); err != nil {
		log.Fatalln("error in unmarshall data json", err.Error())
	}
}

func newServer() *dataAdminServer {
	s := dataAdminServer{}
	s.loadData()
	return &s
}

func main() {
	cert, err := tls.LoadX509KeyPair("srv.crt", "srv.key")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDataAdminServer(grpcServer, newServer())
	reflection.Register(grpcServer)

	log.Printf("server listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve grpc", err.Error())
	}
}
