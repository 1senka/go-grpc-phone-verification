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
	routes.POST("/getcode", svc.GetCode)
	routes.POST("/verify", svc.Verify)

	return svc
}

func (svc *ServiceClient) GetCode(ctx *gin.Context) {
	routes.GetCode(ctx, svc.Client)
}

func (svc *ServiceClient) Verify(ctx *gin.Context) {
	routes.Verify(ctx, svc.Client)
}
