package main

import (
	"flag"
	pb "github.com/kwanok/product/service/ecommerce"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.String("port", "7619", "port to listen on")

func main() {
	flag.Parse()

	con, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, pb.NewServer())

	log.Printf("Starting gRPC server on port %s", *port)

	if err = s.Serve(con); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
