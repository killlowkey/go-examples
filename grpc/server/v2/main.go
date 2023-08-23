package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-example/api/v2"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(req *pb.HelloRequest, greeterSayHelloServer pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", req.GetName())
	for i := 0; i < 10; i++ {
		err := greeterSayHelloServer.Send(&pb.HelloReply{
			Code:    200,
			Message: fmt.Sprintf("hello %s, %d", req.GetName(), i),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("v2: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
