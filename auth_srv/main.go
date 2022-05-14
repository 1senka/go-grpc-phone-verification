package main

import (
	"fmt"
	"log"
	"net"

	"auth_svc/pkg/config"
	"auth_svc/pkg/db"
	"auth_svc/pkg/pb"
	"auth_svc/pkg/services"
	"auth_svc/pkg/sms"
	"auth_svc/pkg/utils"

	// "auth_svc/pkg/pb"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	ss := sms.Init("YOUR_TOKEN")
	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
		Sms: ss,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
