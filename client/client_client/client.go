package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"psycology_backend/client/clientpb"
)

func main() {
	fmt.Println("Hello, world")
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := clientpb.NewClientServiceClient(conn)
	fmt.Println("created client ", c)
	doUnary(c)

}
func doUnary(c clientpb.ClientServiceClient) {
	res, errr := c.PhoneToken(context.Background(), &clientpb.ClientPhoneTokenRequest{
		Phone: "09120357479",
		Token: "48498095",
	})
	if errr != nil {
		panic(errr)
	}
	fmt.Println("token result ", res)
}
