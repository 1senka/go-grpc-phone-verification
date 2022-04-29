package profile

import (
	"fmt"

	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client profilepb.ProfileServiceClient
}

func InitServiceClient(c *config.Config) profilepb.ProfileServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ProfileSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return profilepb.NewProfileServiceClient(cc)
}
