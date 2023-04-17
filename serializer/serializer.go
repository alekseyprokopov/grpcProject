package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"os"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
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

func WriteProtobufToJSONfile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot ProtobufToJSON file: %w", err)
	}

	if err := os.WriteFile(filename, []byte(data), 0666); err != nil {
		return err
	}
	return nil
}
