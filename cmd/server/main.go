package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcProject/pb"
	"grpcProject/service"
	"log"
	"net"
)

func main() {
	port := flag.String("port", "", "the server port")
	flag.Parse()
	log.Printf("server started on port: %s", *port)
	laptopServer := service.NewLaptopServer(service.NewMapStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	addr := fmt.Sprintf("0.0.0.0:%s", *port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("can't start server: ", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("can't start server: ", err)
	}
}
