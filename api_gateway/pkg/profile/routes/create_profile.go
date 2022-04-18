package routes

import (
	"context"
	"net/http"

	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"github.com/gin-gonic/gin"
)

type CreatefileRequestBody struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
	Email      string `json:"email"`
	Gender     int32  `json:"gender"`
	BirthDate  string `json:"birthDate"`
}

func CreateProfile(ctx *gin.Context, c pb.ProfileServiceClient) {
	body := CreatefileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProfile(context.Background(), &pb.CreateProfileRequest{
		Name:       body.Name,
		Email:      body.Email,
		ClientType: body.ClientType,
		Phone:      body.Phone,
		Gender:     body.Gender,
		BirthDate:  body.BirthDate,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
