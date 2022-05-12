package routes

import (
	"context"
	"net/http"

	"github.com/1senka/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type VerifyRequestBody struct {
	Phone      string `json:"phone"`
	Code       string `json:"code"`
}

func Verify(ctx *gin.Context, c pb.AuthServiceClient) {
	b := VerifyRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Verify(context.Background(), &pb.VerifyRequest{
		Phone: b.Phone,

		Code: b.Code,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if (res.Token == "") {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
