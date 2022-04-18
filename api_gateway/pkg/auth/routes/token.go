package routes

import (
	"context"
	"net/http"

	"github.com/1senka/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type TokenRequestBody struct {
	Phone      string `json:"phone"`
	Code       string `json:"code"`
	ClientType string `json:"clientType"`
}

func Token(ctx *gin.Context, c pb.AuthServiceClient) {
	b := TokenRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.PhoneToken(context.Background(), &pb.ClientPhoneTokenRequest{
		Phone: b.Phone,

		Code: b.Code,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
