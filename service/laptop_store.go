package service

import (
	"errors"
	"grpcProject/pb"
	"sync"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
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

func (s *mapStore) Save(laptop *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	s.data[laptop.Id] = laptop
	return nil
}