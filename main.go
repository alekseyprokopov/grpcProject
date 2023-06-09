package main

import (
	"fmt"
	"grpcProject/pb"
	"grpcProject/sample"
	"grpcProject/serializer"
)

func main() {
	binFile := "./tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binFile)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binFile, laptop2)
	if err != nil {
		fmt.Printf("err: %+v", err)

	}

}
