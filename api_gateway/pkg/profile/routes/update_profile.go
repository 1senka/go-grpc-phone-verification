package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"github.com/gin-gonic/gin"
)

type UpdateProfileRequestBody struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birthDate"`
}

func UpdateProfile(ctx *gin.Context, c profilepb.ProfileServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	body := UpdateProfileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if body.ClientType == "client" {
		res, err := c.ClientUpdateProfile(context.Background(), &profilepb.ClientUpdateProfileRequest{
			UserId: string(id),
			Name:   body.Name,
			Phone:  body.Phone,
			Email:  body.Email, Birthdate: body.BirthDate, Gender: body.Gender,
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)

		}
		ctx.JSON(http.StatusCreated, &res)

	}
	if body.ClientType == "therapist" {
		res, err := c.TherapistUpdateProfile(context.Background(), &profilepb.TherapistUpdateProfileRequest{
			UserId: string(id),
			Name:   body.Name,
			Phone:  body.Phone,
			Email:  body.Email, Birthdate: body.BirthDate, Gender: body.Gender,
		})
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)

		}
		ctx.JSON(http.StatusCreated, &res)
	}
	//res, err := c.FindOne(context.Background(), &profileprofilepb.FindOneRequest{
	//	Id: int64(id),
	//})
	//
	//if err != nil {
	//	ctx.AbortWithError(http.StatusBadGateway, err)
	//	return
	//}
	//
	//ctx.JSON(http.StatusCreated, &res)
}
