package profile

import (
	"github.com/1senka/go-grpc-api-gateway/pkg/auth"
	"github.com/1senka/go-grpc-api-gateway/pkg/config"
	"github.com/1senka/go-grpc-api-gateway/pkg/profile/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/profile")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProfile)
	routes.GET("/", svc.GetProfile)
	routes.PUT("/", svc.UpdateProfile)
	routes.GET("/freeTime", svc.GetTherapistFreeTime)
	routes.POST("/setTime", svc.SetFreeTime)

}

func (svc *ServiceClient) CreateProfile(c *gin.Context) {
	routes.CreateProfile(c, svc.Client)
}

func (svc *ServiceClient) UpdateProfile(ctx *gin.Context) {
	routes.UpdateProfile(ctx, svc.Client)
}

func (svc *ServiceClient) GetTherapistFreeTime(ctx *gin.Context) {
	routes.GetTherapistFreeTime(ctx, svc.Client)
}
func (svc *ServiceClient) SetFreeTime(ctx *gin.Context) {
	routes.SetFreeTime(ctx, svc.Client)
}
func (svc *ServiceClient) GetProfile(ctx *gin.Context) {
	routes.GetProfile(ctx, svc.Client)
}
