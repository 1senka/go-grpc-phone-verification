package main

import (
	"fmt"

	//"github.com/1senka/go-grpc-profile-svc/pkg/sms"
	"github.com/1senka/go-grpc-booking-svc/pkg/config"
	"github.com/1senka/go-grpc-booking-svc/pkg/db"
	"github.com/1senka/go-grpc-booking-svc/pkg/pb"

	"log"
	"net"

	"github.com/1senka/go-grpc-booking-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	//ss := sms.Init("556E65702B4A626279345335534B5835484574733734464238734D4D6A5143412F6A6935786C72635554303D")
	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	conn, err := grpc.Dial(c.ProfileUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("created client ", c)
	fmt.Println("Booking Svc on", c.Port)

	s := services.Server{
		H: h,
		//Sms: ss,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBookingServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
