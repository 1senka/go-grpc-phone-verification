package profile

import (
	"fmt"

	"github.com/1senka/go-grpc-api-gateway/pkg/booking/pb"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.BookingServiceClient
}

func InitServiceClient(c *config.Config) pb.BookingServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProfileSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewBookingServiceClient(cc)
}
