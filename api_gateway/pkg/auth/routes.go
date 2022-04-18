package auth

import (
	"github.com/1senka/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/proof", svc.PhoneProof)
	routes.POST("/token", svc.PhoneToken)

	return svc
}

func (svc *ServiceClient) PhoneProof(ctx *gin.Context) {
	routes.Proof(ctx, svc.Client)
}

func (svc *ServiceClient) PhoneToken(ctx *gin.Context) {
	routes.Token(ctx, svc.Client)
}
