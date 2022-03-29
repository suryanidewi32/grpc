1. install mod
    go mod init example.com/helloworld

2. type proto structure

3. install grpc
    go get google.golang.org/grpc

4. run  
    protoc --go_out=paths=source_relative:. proto/helloworld.proto (to make extension .pb.go)
    protoc --go-grpc_out=paths=source_relative:. proto/helloworld.proto (to make extension _grpc.pb.go)

5. make server and client file

6. get certificate and key 
    step ca certificate localhost srv.crt srv.key

7. Request a copy of your CA root certificate
    step ca root ca.crt

8. grpcurll install
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

9. run server
    go run server/main.go

10. Run grpc client
     grpcurl -d '{"name": "bob"}' localhost:5443 helloworld.Greeter.SayHello
