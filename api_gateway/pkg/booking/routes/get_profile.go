package routes

import (
	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Profile struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	ClientType string `json:"clientType"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birthDate"`
}
type GetProfileRequestBody struct {
	ClientType string `json:"clientType"`
}

func GetProfile(ctx *gin.Context, c pb.ProfileServiceClient) {
	body := GetProfileRequestBody{}
	user_id, e := ctx.Get("userId")
	if e == false {
		ctx.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if body.ClientType == "client" {
		resp, err := c.ClientGetProfile(ctx, &pb.ClientGetProfileRequest{UserId: user_id.(string)})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
	if body.ClientType == "therapist" {
		resp, err := c.TherapistGetProfile(ctx, &pb.TherapistGetProfileRequest{Id: user_id.(string)})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}

}