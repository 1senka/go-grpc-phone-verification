package booking

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

	routes := r.Group("/booking")
	routes.Use(a.AuthRequired)
	routes.POST("/book", svc.BookNewSession)
	routes.GET("/freeTime", svc.GetTherapistFreeTime)
	routes.POST("/setTime", svc.SetFreeTime)
}
func (svc *ServiceClient) BookNewSession(c *gin.Context) {
	routes.BookSession(c, svc.Client)
}
func (svc *ServiceClient) GetTherapistFreeTime(ctx *gin.Context) {
	routes.GetTherapistFreeTime(ctx, svc.Client)
}
func (svc *ServiceClient) SetFreeTime(ctx *gin.Context) {
	routes.SetFreeTime(ctx, svc.Client)
}
