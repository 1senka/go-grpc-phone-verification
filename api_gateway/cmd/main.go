package main

import (
	"log"

	"github.com/1senka/go-grpc-api-gateway/pkg/auth"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/1senka/go-grpc-api-gateway/pkg/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	profile.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
