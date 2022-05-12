package services

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"auth_svc/pkg/db"
	"auth_svc/pkg/models"
	"auth_svc/pkg/pb"
	"auth_svc/pkg/sms"
	"auth_svc/pkg/utils"

	"github.com/kavenegar/kavenegar-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var smsClient *kavenegar.Kavenegar

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	Sms sms.Client
}

func (s *Server) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	fmt.Printf("Verify invoked %v", req)
	var smsVerify models.SmsVerify
	err := s.H.SmsVerifyCollection.FindOne(context.Background(), bson.M{"phone": req.Phone}).Decode(&smsVerify)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.VerifyResponse{}, nil
		}
	}
	

	if smsVerify.Code != req.Code {
		return &pb.VerifyResponse{
			
		}, nil
	}
	var user models.User
	user.Id = primitive.NewObjectID()
	user.Phone = req.Phone
	_,er := s.H.UserCollection.InsertOne(context.Background(),user)
	if er != nil {
		return &pb.VerifyResponse{},nil
	}
	token, _ := s.Jwt.GenerateToken(user)
	return &pb.VerifyResponse{
		Token:  token,
	}, nil
}

func (s *Server) GetCode(ctx context.Context, req *pb.GetCodeRequest) (*pb.GetCodeResponse, error) {

	fmt.Printf("GetCode invoked %v", req)
	var smsVerify models.SmsVerify
	err := s.H.SmsVerifyCollection.FindOne(ctx, bson.M{"phone": req.Phone}).Decode(&smsVerify)
	if err != nil {
		fmt.Println(err.Error())
		if smsVerify.ID == primitive.NilObjectID {
			
			



			smsVerify.Phone = req.Phone
			// Generate an 8 digit code
			smsVerify.Code = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
			smsVerify.CreatedAt = time.Now()
			sender := "10008663"
			
			receptor := []string{req.Phone}
			message := "کد تایید شما : " + smsVerify.Code
			s.H.SmsVerifyCollection.InsertOne(ctx, smsVerify)
			if _, err := s.Sms.SmsClient.Message.Send(sender, receptor, message, nil); err != nil {
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
	return &pb.GetCodeResponse{
		Status: 200,
	},nil
		}
		return &pb.GetCodeResponse{Status: 500}, nil
	}
	return &pb.GetCodeResponse{Status: 500}, nil

	
	
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
	err = s.H.UserCollection.FindOne(ctx, bson.M{"_id": claims.Id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.ValidateResponse{
				Status: http.StatusNotFound,
				Error:  "User not found",
			}, nil
		}
	}
	

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id.String(),
		Phone:  user.Phone,
	}, nil
}
