package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	pb "dewi.atop/learn/grpc/admin"
	"google.golang.org/grpc"
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
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataAdminServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve grpc", err.Error())
	}
}
