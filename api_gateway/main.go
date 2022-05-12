package main

import (
	"log"

	"github.com/1senka/go-grpc-api-gateway/pkg/auth"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	auth.RegisterRoutes(r, &c)
	// you can use other services and pass auth_svc like : 
	//authSvc := *auth.RegisterRoutes(r, &c)
	//profile.RegisterRoutes(r, &c, &authSvc)


	r.Run(c.Port)
}
