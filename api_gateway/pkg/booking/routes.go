package profile

import (
	"github.com/1senka/go-grpc-api-gateway/pkg/auth"
	"github.com/1senka/go-grpc-api-gateway/pkg/booking/routes"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/profile")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.BookNewSession)

}
func (svc *ServiceClient) BookNewSession(c *gin.Context) {
	routes.BookSession(c, svc.Client)
}
