package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcProject/pb"
	"grpcProject/sample"
	"io"
	"log"
	"time"
)

func main() {
	add := flag.String("address", "", "server address")
	flag.Parse()
	log.Printf("dial server: %s", add)

	conn, err := grpc.Dial(*add, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cant dial client")
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuGhz:   2,
		MinCpuCores: 3}
	searchLaptop(laptopClient, filter)

}

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Printf("search filter: %+v", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchReq := &pb.SearchLaptopRequest{
		Filter: filter,
	}

	stream, err := laptopClient.SearchLaptop(ctx, searchReq)
	if err != nil {
		log.Fatal("can't search laptop: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("can't receive response: ", err)
		}
		laptop := res.GetLaptop()
		log.Printf("founded laptop: %+v", laptop)
	}

}

func createLaptop(laptopClient pb.LaptopServiceClient) {
	newLaptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: newLaptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("already exist")
		} else {
			log.Fatal("can't create laptop: ", err)
		}
		return
	}

	log.Printf("created laptop: %s", res.Id)
}
