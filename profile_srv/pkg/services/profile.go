package services

import (
	"context"
	"github.com/1senka/go-grpc-profile-svc/pkg/db"
	"github.com/1senka/go-grpc-profile-svc/pkg/pb"
	"github.com/kavenegar/kavenegar-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var smsClient *kavenegar.Kavenegar

type ClientProfile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	UserId      string             `bson:"user_id"`
	Name        string             `bson:"name"`
	BirthDate   string             `bson:"birth_date"`
	Email       string             `bson:"email"`
	Gender      string             `bson:"gender"`
}
type TherapistProfile struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	UserId      string             `bson:"user_id"`
	Name        string             `bson:"name"`
	BirthDate   string             `bson:"birth_date"`
	Email       string             `bson:"email"`
	Gender      string             `bson:"gender"`
}
type Date struct {
	Year   int32 `bson:"year"`
	Month  int32 `bson:"month"`
	Day    int32 `bson:"day"`
	Hour   int32 `bson:"hour"`
	Minute int32 `bson:"min"`
}
type FreeTime struct {
	Date        Date   `bson:"date"`
	TherapistId string `bson:"therapist_id"`
}
type Server struct {
	H db.Handler
}

func (s *Server) ClientCreateProfile(ctx context.Context, req *pb.ClientCreateProfileRequest) (*pb.ClientCreateProfileResponse, error) {
	data := &ClientProfile{}
	_ = s.H.ClientCollection.FindOne(ctx, bson.D{{"phone", req.GetPhone()}}).Decode(data)
	if data.ID != primitive.NilObjectID {
		return &pb.ClientCreateProfileResponse{
			Status: 400,
			Result: "شماره موبایل وارد شده در سیستم ثبت شده است",
		}, nil
	}
	_, error := s.H.ClientCollection.InsertOne(ctx, &ClientProfile{
		UserId:      req.GetUserId(),
		PhoneNumber: req.GetPhone(),
		Name:        req.GetName(),
		BirthDate:   req.GetBirthdate(),
		Email:       req.GetEmail(),
		Gender:      req.GetGender(),
	})
	if error != nil {
		return &pb.ClientCreateProfileResponse{
			Status: 400,
			Result: "خطا در ثبت اطلاعات",
		}, nil
	}
	return &pb.ClientCreateProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت ثبت شد",
	}, nil

}
func (s *Server) ClientGetProfile(ctx context.Context, req *pb.ClientGetProfileRequest) (*pb.ClientGetProfileResponse, error) {
	data := &ClientProfile{}
	_ = s.H.ClientCollection.FindOne(ctx, bson.D{{"user_id", req.GetUserId()}}).Decode(data)
	if data.ID == primitive.NilObjectID {
		return &pb.ClientGetProfileResponse{
			Status: 400,
			Result: "این کاربر قبلا ثبت نام نکرده است",
		}, nil
	}
	return &pb.ClientGetProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت دریافت شد",
		Profile: &pb.Profile{
			UserId:    data.UserId,
			Phone:     data.PhoneNumber,
			Name:      data.Name,
			Birthdate: data.BirthDate,
			Email:     data.Email,
			Gender:    data.Gender,
		},
	}, nil
}
func (s *Server) ClientUpdateProfile(ctx context.Context, req *pb.ClientUpdateProfileRequest) (*pb.ClientUpdateProfileResponse, error) {
	filter := bson.D{{"user_id", req.GetUserId()}}
	update := bson.D{
		{"$set", bson.D{
			{"phone_number", req.GetPhone()},
			{"user_id", req.GetUserId()},

			{"name", req.GetName()},
			{"birth_date", req.GetBirthdate()},

			{"email", req.GetEmail()},
			{"gender", req.GetGender()},
		}}}
	_, error := s.H.TherapistCollection.UpdateOne(ctx, filter, update)
	if error != nil {
		return &pb.ClientUpdateProfileResponse{
			Status: 400,
			Result: "خطا در بروزرسانی اطلاعات",
		}, nil

	}
	return &pb.ClientUpdateProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت بروزرسانی شد",
	}, nil

}
func (s *Server) TherapistUpdateProfile(ctx context.Context, req *pb.TherapistUpdateProfileRequest) (*pb.TherapistUpdateProfileResponse, error) {
	filter := bson.D{{"user_id", req.GetUserId()}}
	update := bson.D{
		{"$set", bson.D{
			{"phone_number", req.GetPhone()},
			{"user_id", req.GetUserId()},

			{"name", req.GetName()},
			{"birth_date", req.GetBirthdate()},

			{"email", req.GetEmail()},
			{"gender", req.GetGender()},
		}}}
	_, error := s.H.TherapistCollection.UpdateOne(ctx, filter, update)

	if error != nil {
		return &pb.TherapistUpdateProfileResponse{
			Status: 400,
			Result: "خطا در ثبت اطلاعات",
			Id:     req.GetUserId(),
		}, nil

	}
	return &pb.TherapistUpdateProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت ثبت شد",
		Id:     req.GetUserId(),
	}, nil
}

func (s *Server) TherapistCreateProfile(ctx context.Context, req *pb.TherapistCreateProfileRequest) (*pb.TherapistCreateProfileResponse, error) {
	data := &TherapistProfile{}
	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetUserId()}}).Decode(data)
	if data.ID != primitive.NilObjectID {
		return &pb.TherapistCreateProfileResponse{
			Status: 400,
			Result: "شماره موبایل وارد شده در سیستم ثبت شده است",
		}, nil
	}
	_, error := s.H.TherapistCollection.InsertOne(ctx, &TherapistProfile{
		UserId:      req.GetUserId(),
		PhoneNumber: req.GetPhone(),
		Name:        req.GetName(),
		BirthDate:   req.GetBirthdate(),
		Email:       req.GetEmail(),
		Gender:      req.GetGender(),
	})
	if error != nil {
		return &pb.TherapistCreateProfileResponse{
			Status: 400,
			Result: "خطا در ثبت اطلاعات",
		}, nil
	}
	return &pb.TherapistCreateProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت ثبت شد",
	}, nil

}
func (s *Server) TherapistGetProfile(ctx context.Context, req *pb.TherapistGetProfileRequest) (*pb.TherapistGetProfileResponse, error) {
	data := &TherapistProfile{}
	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
	if data.ID == primitive.NilObjectID {
		return &pb.TherapistGetProfileResponse{
			Status: 400,
			Result: "این کاربر قبلا ثبت نام نکرده است",
		}, nil
	}
	return &pb.TherapistGetProfileResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت دریافت شد",
		Profile: &pb.Profile{
			UserId:    data.UserId,
			Phone:     data.PhoneNumber,
			Name:      data.Name,
			Birthdate: data.BirthDate,
			Email:     data.Email,
			Gender:    data.Gender,
		},
	}, nil
}

//func (s *Server) GetFreeTime(ctx context.Context, req *pb.TherapistGetFreeTimeRequest) (*pb.TherapistGetFreeTimeResponse, error) {
//	data := &TherapistProfile{}
//	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
//	if data.ID == primitive.NilObjectID {
//		return &pb.TherapistGetFreeTimeResponse{
//			Status: 400,
//			Result: "این کاربر قبلا ثبت نام نکرده است",
//		}, nil
//	}
//	return &pb.TherapistGetFreeTimeResponse{
//		Status:   200,
//		Result:   "success",
//		FreeTime: data.FreeTime,
//	}, nil
//}
func (s *Server) GetTherapistFreeTime(ctx context.Context, req *pb.TherapistGetFreeTimeRequest) (*pb.TherapistGetFreeTimeResponse, error) {
	data := []*pb.FreeTime{}
	cursor, error := s.H.FreeTimeCollection.Find(ctx, bson.D{{"therapist_id", req.GetId()}})
	if error != nil {
		return &pb.TherapistGetFreeTimeResponse{
			Status:   400,
			FreeTime: []*pb.FreeTime{},
		}, nil
	}
	for cursor.Next(ctx) {
		var freeTime FreeTime
		_ = cursor.Decode(&freeTime)
		data = append(data, &pb.FreeTime{Date: &pb.Date{Year: freeTime.Date.Year, Month: freeTime.Date.Month, Day: freeTime.Date.Day, Hour: freeTime.Date.Hour, Minute: freeTime.Date.Minute},
			TherapistId: freeTime.TherapistId,
		})
	}
	return &pb.TherapistGetFreeTimeResponse{
		Status:   200,
		FreeTime: data,
	}, nil

}
func (s *Server) GetFreeTime(ctx context.Context, req *pb.GetFreeTimeRequest) (*pb.GetFreeTimeResponse, error) {
	data := []*pb.FreeTime{}
	cursor, error := s.H.FreeTimeCollection.Find(ctx, bson.D{{req.GetKeyword().Key, req.GetKeyword().Value}})
	if error != nil {
		return &pb.GetFreeTimeResponse{
			Status:    400,
			FreeTimes: []*pb.FreeTime{},
		}, nil
	}
	for cursor.Next(ctx) {
		var freeTime FreeTime
		_ = cursor.Decode(&freeTime)
		data = append(data, &pb.FreeTime{Date: &pb.Date{Year: freeTime.Date.Year, Month: freeTime.Date.Month, Day: freeTime.Date.Day, Hour: freeTime.Date.Hour, Minute: freeTime.Date.Minute},
			TherapistId: freeTime.TherapistId,
		})
	}
	return &pb.GetFreeTimeResponse{
		Status:    200,
		FreeTimes: data,
	}, nil
}
func (s *Server) SetFreeTime(ctx context.Context, req *pb.TherapistSetFreeTimeRequest) (*pb.TherapistSetFreeTimeResponse, error) {
	data := &TherapistProfile{}
	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
	if data.ID == primitive.NilObjectID {
		return &pb.TherapistSetFreeTimeResponse{
			Status: 400,
			Result: "این کاربر قبلا ثبت نام نکرده است",
		}, nil
	}
	var freeTimes []interface{}
	freeTime := req.GetFreeTime()
	for _, v := range freeTime {
		//s.H.FreeTimeCollection.InsertOne(ctx, &FreeTime{
		//	TherapistId: req.Id, Date: *v.Date,
		//})
		_ = append(freeTimes, &FreeTime{
			Date:        Date{Year: v.Year, Month: v.Month, Day: v.Day, Hour: v.Hour, Minute: v.Minute},
			TherapistId: req.Id,
		})
	}

	s.H.FreeTimeCollection.DeleteMany(ctx, bson.D{{"id", data.UserId}})
	_, error := s.H.FreeTimeCollection.InsertMany(ctx, freeTimes)
	if error != nil {
		return &pb.TherapistSetFreeTimeResponse{
			Status: 400,
			Result: "خطا در ثبت اطلاعات",
		}, nil
	}
	return &pb.TherapistSetFreeTimeResponse{
		Status: 200,
		Result: "اطلاعات با موفقیت ثبت شد",
	}, nil
}

//func (s *Server) PhoneToken(ctx context.Context, req *pb.ClientPhoneTokenRequest) (*pb.ClientPhoneTokenResponse, error) {
//	fmt.Println("ClientPhoneToken invoked %v", req)
//	var smsVerify models.SmsVerify
//
//	if result := s.H.DB.Where(&models.SmsVerify{Phone: req.Phone}).First(&smsVerify); result.Error != nil {
//		return &pb.ClientPhoneTokenResponse{
//			Result: "not found",
//			Token:  "",
//		}, nil
//	}
//
//	if smsVerify.Code != req.Code {
//		return &pb.ClientPhoneTokenResponse{
//			Result: "code not match",
//			Token:  "",
//		}, nil
//	}
//	var user models.User
//	user.Phone = req.Phone
//	user.ClientType = req.ClientType
//	s.H.DB.Create(&user)
//	token, _ := s.Jwt.GenerateToken(user)
//	return &pb.ClientPhoneTokenResponse{
//		Result: "ok",
//		Token:  token,
//	}, nil
//}
//
//func (s *Server) PhoneProof(ctx context.Context, req *pb.ClientPhoneProofRequest) (*pb.ClientPhoneProofResponse, error) {
//
//	fmt.Println("ClientPhoneProof invoked %v", req)
//	var smsVerify models.SmsVerify
//	if result := s.H.DB.Where(&models.User{Phone: req.Phone}).First(&smsVerify); result.Error == nil {
//		return &pb.ClientPhoneProofResponse{
//			Status: http.StatusConflict,
//		}, nil
//	}
//	smsVerify.Phone = req.Phone
//	smsVerify.Code = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
//	sender := "10008663"
//	receptor := []string{"09120357479"}
//	message := "کد تایید شما : " + smsVerify.Code
//	fmt.Println("moz")
//
//	if _, err := s.Sms.SmsClient.Message.Send(sender, receptor, message, nil); err != nil {
//		fmt.Println("moz")
//		switch err := err.(type) {
//		case *kavenegar.APIError:
//			fmt.Println("api")
//
//			fmt.Println(err.Error())
//		case *kavenegar.HTTPError:
//			fmt.Println("http")
//
//			fmt.Println(err.Error())
//		default:
//			fmt.Println("error")
//
//			fmt.Println(err.Error())
//		}
//	}
//	smsVerify.ExpireTime = time.Now().Add(time.Minute * 5)
//	s.H.DB.Create(&smsVerify)
//	return &pb.ClientPhoneProofResponse{
//		Status: http.StatusCreated,
//	}, nil
//}
//func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
//	claims, err := s.Jwt.ValidateToken(req.Token)
//
//	if err != nil {
//		return &pb.ValidateResponse{
//			Status: http.StatusBadRequest,
//			Error:  err.Error(),
//		}, nil
//	}
//
//	var user models.User
//
//	if result := s.H.DB.Where(&models.User{Phone: claims.Phone}).First(&user); result.Error != nil {
//		return &pb.ValidateResponse{
//			Status: http.StatusNotFound,
//			Error:  "User not found",
//		}, nil
//	}
//
//	return &pb.ValidateResponse{
//		Status: http.StatusOK,
//		UserId: user.Id,
//	}, nil
//}
