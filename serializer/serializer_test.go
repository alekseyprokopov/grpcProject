package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"grpcProject/pb"
	"grpcProject/sample"
	"log"
	"testing"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binFile, laptop2)
	log.Printf("%+v", laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	jsonFile := "../tmp/json"

	err = WriteProtobufToJSONfile(laptop1, jsonFile)
	require.NoError(t, err)
}
