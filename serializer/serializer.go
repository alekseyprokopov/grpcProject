package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	os.WriteFile(filename, data, 0666)
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	file, err := os.ReadFile(filename)
	log.Printf("FILE: %s", string(file))

	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	if err := proto.Unmarshal(file, message); err != nil {
		return fmt.Errorf("cannot unmarshal proto file: %w", err)
	}

	return nil
}
