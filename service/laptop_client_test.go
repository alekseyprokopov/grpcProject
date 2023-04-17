package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"grpcProject/pb"
	"grpcProject/sample"
	"grpcProject/serializer"
	"grpcProject/service"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	server, addr := startTestLaptopServer(t)
	client := newTestLaptopClient(t, addr)

	newLaptop := sample.NewLaptop()
	expectedID := newLaptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: newLaptop,
	}
	//test client
	res, err := client.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	//test server
	other, err := server.Store.Find(res.Id)
	//serializer.WriteProtobufToJSONfile(other, "../tmp/json")

	require.NoError(t, err)
	require.NotNil(t, other)

	json1, _ := serializer.ProtobufToJSON(newLaptop)
	json2, _ := serializer.ProtobufToJSON(other)
	require.Equal(t, json2, json1)

}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	server := service.NewLaptopServer(service.NewMapStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8080")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return server, listener.Addr().String()

}

func newTestLaptopClient(t *testing.T, addr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
