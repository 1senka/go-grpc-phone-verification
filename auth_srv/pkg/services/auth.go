package services

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/1senka/go-grpc-auth-svc/pkg/db"
	"github.com/1senka/go-grpc-auth-svc/pkg/models"
	"github.com/1senka/go-grpc-auth-svc/pkg/pb"
	"github.com/1senka/go-grpc-auth-svc/pkg/sms"
	"github.com/1senka/go-grpc-auth-svc/pkg/utils"
	"github.com/kavenegar/kavenegar-go"
)

var smsClient *kavenegar.Kavenegar

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	Sms sms.Client
}

func (s *Server) PhoneToken(ctx context.Context, req *pb.ClientPhoneTokenRequest) (*pb.ClientPhoneTokenResponse, error) {
	fmt.Println("ClientPhoneToken invoked %v", req)
	var smsVerify models.SmsVerify

	if result := s.H.DB.Where(&models.SmsVerify{Phone: req.Phone}).First(&smsVerify); result.Error != nil {
		return &pb.ClientPhoneTokenResponse{
			Result: "not found",
			Token:  "",
		}, nil
	}

	if smsVerify.Code != req.Code {
		return &pb.ClientPhoneTokenResponse{
			Result: "code not match",
			Token:  "",
		}, nil
	}
	var user models.User
	user.Phone = req.Phone
	user.ClientType = req.ClientType
	s.H.DB.Create(&user)
	token, _ := s.Jwt.GenerateToken(user)
	return &pb.ClientPhoneTokenResponse{
		Result: "ok",
		Token:  token,
	}, nil
}

func (s *Server) PhoneProof(ctx context.Context, req *pb.ClientPhoneProofRequest) (*pb.ClientPhoneProofResponse, error) {

	fmt.Println("ClientPhoneProof invoked %v", req)
	var smsVerify models.SmsVerify
	if result := s.H.DB.Where(&models.User{Phone: req.Phone}).First(&smsVerify); result.Error == nil {
		return &pb.ClientPhoneProofResponse{
			Status: http.StatusConflict,
		}, nil
	}
	smsVerify.Phone = req.Phone
	smsVerify.Code = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
	sender := "10008663"
	receptor := []string{req.Phone}
	message := "کد تایید شما : " + smsVerify.Code
	fmt.Println("moz")

	if _, err := s.Sms.SmsClient.Message.Send(sender, receptor, message, nil); err != nil {
		fmt.Println("moz")
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println("api")

			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println("http")

			fmt.Println(err.Error())
		default:
			fmt.Println("error")

			fmt.Println(err.Error())
		}
	}
	smsVerify.ExpireTime = time.Now().Add(time.Minute * 5)
	s.H.DB.Create(&smsVerify)
	return &pb.ClientPhoneProofResponse{
		Status: http.StatusCreated,
	}, nil
}
func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Phone: claims.Phone}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
		Phone:  user.Phone,
	}, nil
}
