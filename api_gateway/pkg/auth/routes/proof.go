package routes

import (
	"context"
	"github.com/1senka/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProofRequestBody struct {
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
}

func Proof(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ProofRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.PhoneProof(context.Background(), &pb.ClientPhoneProofRequest{
		Phone:      body.Phone,
		ClientType: body.ClientType,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
