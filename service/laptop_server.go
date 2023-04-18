package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcProject/pb"
	"log"
	"time"
)

type LaptopServer struct {
	pb.LaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(Store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: Store,
	}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid ID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "can't generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	if err := s.Store.Save(laptop); err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "already exists :%v", err)
	}
	//heavy operations:
	time.Sleep(6 * time.Second)
	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	log.Printf("saved laptop with id :%s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func (s *LaptopServer) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter %+v", filter)

	err := s.Store.Search(filter, func(laptop *pb.Laptop) error {
		res := &pb.SearchLaptopResponse{
			Laptop: laptop,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		log.Printf("sent laptop with id: %s", laptop.Id)
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %w", err)
	}

	return nil
}
