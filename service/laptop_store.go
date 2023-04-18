package service

import (
	"context"
	"errors"
	"grpcProject/pb"
	"log"
	"sync"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

func NewMapStore() LaptopStore {
	return &mapStore{
		mutex: sync.RWMutex{},
		data:  make(map[string]*pb.Laptop),
	}
}

type mapStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func (s *mapStore) Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, laptop := range s.data {
		if ctx.Err() == context.Canceled {
			log.Printf("context is cancelled")
			return errors.New("context is cancelled")
		}
		if isQualified(filter, laptop) {
			err := found(laptop)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *mapStore) Save(laptop *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	s.data[laptop.Id] = laptop
	return nil
}

func (s *mapStore) Find(id string) (*pb.Laptop, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	item, ok := s.data[id]
	if !ok {
		return nil, nil
	}

	return item, nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if filter.GetMinCpuGhz() < laptop.GetCpu().GetMinGhz() {
		return false
	}

	if filter.GetMaxPriceUsd() < laptop.GetPriceUsd() {
		return false
	}
	if filter.GetMinCpuCores() < laptop.GetCpu().GetNumberCores() {
		return false
	}

	return true
}
