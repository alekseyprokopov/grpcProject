package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcProject/pb"
	"grpcProject/sample"
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
