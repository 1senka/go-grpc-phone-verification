package routes

import (
	"fmt"
	"github.com/1senka/go-grpc-api-gateway/pkg/profile/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Date struct {
	Year   int32 `json:"year"`
	Month  int32 `json:"month"`
	Day    int32 `json:"day"`
	Hour   int32 `json:"hour"`
	Minute int32 `json:"min"`
}
type FreeTime struct {
	Date        Date   `json:"date"`
	TherapistId string `json:"therapist_id"`
}
type FreeTimes struct {
	FreeTimes []FreeTime `json:"free_times"`
}

func SetFreeTime(ctx *gin.Context, c profilepb.ProfileServiceClient) {
	body := FreeTimes{}
	user_id, e := ctx.Get("userId")
	if e == false {
		ctx.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	freeTimes := []*profilepb.Date{}
	for _, freeTime := range body.FreeTimes {
		freeTimes = append(freeTimes, &profilepb.Date{
			Year:   freeTime.Date.Year,
			Month:  freeTime.Date.Month,
			Day:    freeTime.Date.Day,
			Hour:   freeTime.Date.Hour,
			Minute: freeTime.Date.Minute,
		})
	}
	res, err := c.SetFreeTime(ctx, &profilepb.TherapistSetFreeTimeRequest{
		FreeTime: freeTimes,
		Id:       fmt.Sprintf("%v", user_id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func GetTherapistFreeTime(ctx *gin.Context, c profilepb.ProfileServiceClient) {
	user_id, e := ctx.Get("userId")
	if e == false {
		ctx.AbortWithError(http.StatusBadRequest, nil)
		return
	}
	res, err := c.GetTherapistFreeTime(ctx, &profilepb.TherapistGetFreeTimeRequest{
		Id: fmt.Sprintf("%v", user_id),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

//func CreateProfile(ctx *gin.Context, c profileprofilepb.ProfileServiceClient) {
//	body := CreateProfileRequestBody{}
//
//	if err := ctx.BindJSON(&body); err != nil {
//		ctx.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	user_id, e := ctx.Get("userId")
//	if e == false {
//		ctx.AbortWithError(http.StatusBadRequest, nil)
//		return
//	}
//
//	if body.ClientType == "client" {
//		res, err := c.ClientCreateProfile(context.Background(), &profilepb.ClientCreateProfileRequest{
//			Name:      body.Name,
//			Phone:     body.Phone,
//			Email:     body.Email,
//			Gender:    body.Gender,
//			Birthdate: body.BirthDate,
//			UserId:    fmt.Sprintf("%v", user_id),
//		})
//		if err != nil {
//			ctx.AbortWithError(http.StatusBadGateway, err)
//			return
//		}
//
//		ctx.JSON(http.StatusCreated, &res)
//	}
//	if body.ClientType == "therapist" {
//		res, err := c.TherapistCreateProfile(context.Background(), &profilepb.TherapistCreateProfileRequest{
//			Name:      body.Name,
//			Phone:     body.Phone,
//			Email:     body.Email,
//			Gender:    body.Gender,
//			Birthdate: body.BirthDate,
//			UserId:    fmt.Sprintf("%v", user_id),
//		})
//		if err != nil {
//			ctx.AbortWithError(http.StatusBadGateway, err)
//			return
//		}
//
//		ctx.JSON(http.StatusCreated, &res)
//	}
//
//}
