package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"github.com/gin-gonic/gin"
)

type CreateProfileRequestBody struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birthDate"`
	userId     string `json:"userId"`
}

func CreateProfile(ctx *gin.Context, c profilepb.ProfileServiceClient) {
	body := CreateProfileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user_id, e := ctx.Get("userId")
	fmt.Println(user_id)
	if e == false {
		ctx.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	fmt.Println(body)

	if body.ClientType == "client" {
		res, err := c.ClientCreateProfile(context.Background(), &profilepb.ClientCreateProfileRequest{
			Name:      body.Name,
			Phone:     body.Phone,
			Email:     body.Email,
			Gender:    body.Gender,
			Birthdate: body.BirthDate,
			UserId:    fmt.Sprintf("%v", user_id),
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}

		ctx.JSON(http.StatusCreated, &res)
	}
	if body.ClientType == "therapist" {
		res, err := c.TherapistCreateProfile(context.Background(), &profilepb.TherapistCreateProfileRequest{
			Name:      body.Name,
			Phone:     body.Phone,
			Email:     body.Email,
			Gender:    body.Gender,
			Birthdate: body.BirthDate,
			UserId:    fmt.Sprintf("%v", user_id),
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}

		ctx.JSON(http.StatusCreated, &res)
	}

}
