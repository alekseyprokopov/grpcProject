package sample

import (
	"github.com/google/uuid"
	"grpcProject/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewRam() *pb.Memory {
	return &pb.Memory{
		Value: uint64(randInt(4, 32)),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func NewCPU() *pb.CPU {
	brands := []string{"intel", "amd"}
	names := map[string][]string{
		"inter": {"Intel Core M", "Intel Core i3", "Intel Core i5", "Intel Core i7"},
		"amd":   {"AMD Ryzen 7", "AMD Ryzen 5", "AMD Ryzen 3", "AMD APU", "AMD FX"}}
	minGhz := randFloat(1, 1.5)
	maxGhz := randFloat(minGhz, 3)
	return &pb.CPU{
		Brand:  randomString(brands...),
		Name:   randomString(names[randomString(brands...)]...),
		MinGhz: minGhz,
		MaxGhz: maxGhz}

}

func NewGPU() *pb.GPU {
	brands := []string{"nvidia", "amd"}
	names := map[string][]string{
		"nvidia": {
			"GeForce RTX 4080",
			"GeForce RTX 4080 12GB",
			"GeForce RTX 4070 Ti",
			"GeForce RTX 4070",
			"GeForce RTX 3090 Ti"},
		"amd": {
			"AMD Radeon RX 7900 XTX",
			"AMD Radeon RX 7900 XT",
			"AMD Radeon RX 6700",
			"AMD Radeon RX 6650 XT",
			"AMD Radeon RX 6750 XT"}}
	minGhz := randFloat(1, 1.5)
	maxGhz := randFloat(minGhz, 3)
	return &pb.GPU{
		Brand:  randomString(brands...),
		Name:   randomString(names[randomString(brands...)]...),
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(randInt(1, 4)),
			Unit:  pb.Memory_GIGABYTE,
		}}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randInt(1, 4)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randInt(1, 4)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewLaptop() *pb.Laptop {
	brands := []string{"dell", "acer", "lenovo"}
	names := map[string][]string{
		"dell":   {"Alienware", "Latitude", "Precision", "Vostro"},
		"acer":   {"Nitro AN517-41-R11Z", "Aspire 3 A317-53-58UL", "Extensa 15 EX 215-31-C6FV"},
		"lenovo": {"Legion 5", "Yoga Slim 7", "IdeaPad 5 Gen 7"},
	}
	return &pb.Laptop{
		Id:          uuid.New().String(),
		Brand:       randomString(brands...),
		Name:        randomString(names[randomString(brands...)]...),
		Cpu:         NewCPU(),
		Gpus:        []*pb.GPU{NewGPU(), NewGPU()},
		Storages:    []*pb.Storage{NewSSD(), NewHDD()},
		Keyboard:    NewKeyboard(),
		Ram:         NewRam(),
		PriceUsd:    float32(randFloat(100, 500)),
		ReleaseYear: uint32(randInt(2010, 2020)),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: float32(randFloat(1.0, 3.0)),
		},
	}
}
