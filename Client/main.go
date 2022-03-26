package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "dewi.atop/learn/grpc/admin"
	"google.golang.org/grpc"
)

func getDataAdminByEmail(client pb.DataAdminClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Admin{Email: email}
	admin, err := client.FindAdminByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error when get admin by email", err.Error())
	}

	fmt.Println(admin)

}

func main() {
	var opts []grpc.DialOption

	//trial without TLS, Certificate
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial")
	}

	defer conn.Close()

	client := pb.NewDataAdminClient(conn)
	getDataAdminByEmail(client, "dewi@gmail.com")
	getDataAdminByEmail(client, "mark@gmail.com")
}
