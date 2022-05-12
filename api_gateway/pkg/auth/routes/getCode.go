package routes

import (
	"context"
	"net/http"

	"github.com/1senka/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type GetCodeRequestBody struct {
	Phone      string `json:"phone"`
}

func GetCode(ctx *gin.Context, c pb.AuthServiceClient) {
	body := GetCodeRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetCode(context.Background(), &pb.GetCodeRequest{
		Phone:      body.Phone,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Status == 500{
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(int(res.Status), &res)
}
