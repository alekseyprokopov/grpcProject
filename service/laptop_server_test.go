package service_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcProject/pb"
	"grpcProject/sample"
	"grpcProject/service"
	"testing"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid ID"

	laptopDuplicateId := sample.NewLaptop()
	storeDuplicate := service.NewMapStore()
	err := storeDuplicate.Save(laptopDuplicateId)
	require.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewMapStore(),
			code:   codes.OK,
		},
		{
			name:   "success_without_id",
			laptop: laptopNoID,
			store:  service.NewMapStore(),
			code:   codes.OK,
		},
		{
			name:   "success_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewMapStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "duplicate_id",
			laptop: laptopDuplicateId,
			store:  storeDuplicate,
			code:   codes.AlreadyExists,
		},
	}

	for _, item := range testCases {
		t.Run(item.name, func(t *testing.T) {

			req := &pb.CreateLaptopRequest{
				Laptop: item.laptop,
			}

			server := service.NewLaptopServer(item.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if item.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(item.laptop.Id) > 0 {
					require.Equal(t, res.Id, item.laptop.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, item.code, st.Code())
			}

		})
	}

}
