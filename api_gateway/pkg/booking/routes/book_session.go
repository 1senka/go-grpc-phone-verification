package routes

import (
	"context"
	"fmt"
	"github.com/1senka/go-grpc-api-gateway/pkg/booking/pb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookSessionRequestBody struct {
	SessionType string `json:"session_type"`
	IsHotline   bool   `json:"is_hotline"`
}

func BookSession(ctx *gin.Context, c pb.BookingServiceClient) {
	var body BookSessionRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(body)
	req := &pb.BookingRequest{
		SessionType: body.SessionType,
		IsHotline:   body.IsHotline,
	}
	res, err := c.Booking(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
