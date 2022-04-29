package services

import (
	"context"
	"fmt"
	"github.com/1senka/go-grpc-booking-svc/pkg/db"
	"github.com/1senka/go-grpc-booking-svc/pkg/pb"
	"github.com/1senka/go-grpc-booking-svc/pkg/profilepb"
	"github.com/1senka/go-grpc-booking-svc/pkg/utils"
	ptime "github.com/yaa110/go-persian-calendar"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

//
//var smsClient *kavenegar.Kavenegar
//

type Date struct {
	Year   int32 `bson:"year"`
	Month  int32 `bson:"month"`
	Day    int32 `bson:"day"`
	Hour   int32 `bson:"hour"`
	Minute int32 `bson:"min"`
}

type Booking struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserId        string             `bson:"user_id"`
	TherapistId   string             `bson:"therapist_id"`
	BookingDate   Date               `bson:"booking_date"`
	SessionType   string             `bson:"session_type"`
	IsHotline     bool               `bson:"is_hotline"`
	TherapistName string             `bson:"therapist_name"`
	IsPaid        bool               `bson:"is_paid"`
	BookingStatus string             `bson:"booking_status"`
}

//type ClientProfile struct {
//	ID          primitive.ObjectID `bson:"_id,omitempty"`
//	PhoneNumber string             `bson:"phone_number"`
//	UserId      string             `bson:"user_id"`
//	Name        string             `bson:"name"`
//	BirthDate   string             `bson:"birth_date"`
//	Email       string             `bson:"email"`
//	Gender      string             `bson:"gender"`
//}
//type TherapistProfile struct {
//	ID          primitive.ObjectID `bson:"_id,omitempty"`
//	PhoneNumber string             `bson:"phone_number"`
//	UserId      string             `bson:"user_id"`
//	Name        string             `bson:"name"`
//	BirthDate   string             `bson:"birth_date"`
//	Email       string             `bson:"email"`
//	Gender      string             `bson:"gender"`
//}
//type Date struct {
//	Year   int32 `bson:"year"`
//	Month  int32 `bson:"month"`
//	Day    int32 `bson:"day"`
//	Hour   int32 `bson:"hour"`
//	Minute int32 `bson:"min"`
//}
//type FreeTime struct {
//	Date        Date   `bson:"date"`
//	TherapistId string `bson:"therapist_id"`
//}
type Server struct {
	H db.Handler
	C profilepb.ProfileServiceClient
}

func (s *Server) Booking(ctx context.Context, req *pb.BookingRequest) (*pb.BookingResponse, error) {
	year, month, day, hour, minute := ptime.Now().Year(), ptime.Now().Month().String(), ptime.Now().Day(), ptime.Now().Hour(), ptime.Now().Minute()
	year2, month2, day2, hour2, minute2 := ptime.Now().Add(time.Duration(time.Hour*4)).Year(), ptime.Now().Month().String(), ptime.Now().Day(), ptime.Now().Hour(), ptime.Now().Minute()

	fmt.Println(strconv.Itoa(year) + "/" + utils.GetMonthFromPersian(month) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(hour) + "/" + strconv.Itoa(minute))
	ss, err := strconv.Atoi(utils.GetMonthFromPersian(month))
	ss2, err2 := strconv.Atoi(utils.GetMonthFromPersian(month2))

	if err == nil && err2 == nil {
		startDate := &profilepb.Date{
			Year:   int32(year),
			Month:  int32(ss),
			Day:    int32(day),
			Hour:   int32(hour),
			Minute: int32(minute),
		}
		endDate := &profilepb.Date{
			Year:   int32(year2),
			Month:  int32(ss2),
			Day:    int32(day2),
			Hour:   int32(hour2),
			Minute: int32(minute2),
		}
		res, err3 := s.C.GetFreeTime(ctx, &profilepb.GetFreeTimeRequest{
			StartDate: startDate,
			EndDate:   endDate,
		})
		fmt.Println(res.FreeTimes)
		if err3 != nil {
			fmt.Println(err3.Error())
		}
	} else {
		fmt.Println("error")
	}

	//res, err := s.C.GetFreeTime(ctx, &profilepb.GetFreeTimeRequest{
	//	StartDate:
	//})
	return &pb.BookingResponse{BookingId: "", BookingStatus: "", UserId: "", SessionType: "", IsHotline: false, BookingTime: "", TherapistName: "", TherapistId: ""}, nil
}

//
//func (s *Server) ClientCreateProfile(ctx context.Context, req *profilepb.ClientCreateProfileRequest) (*profilepb.ClientCreateProfileResponse, error) {
//	data := &ClientProfile{}
//	_ = s.H.ClientCollection.FindOne(ctx, bson.D{{"phone", req.GetPhone()}}).Decode(data)
//	if data.ID != primitive.NilObjectID {
//		return &profilepb.ClientCreateProfileResponse{
//			Status: 400,
//			Result: "شماره موبایل وارد شده در سیستم ثبت شده است",
//		}, nil
//	}
//	_, error := s.H.ClientCollection.InsertOne(ctx, &ClientProfile{
//		UserId:      req.GetUserId(),
//		PhoneNumber: req.GetPhone(),
//		Name:        req.GetName(),
//		BirthDate:   req.GetBirthdate(),
//		Email:       req.GetEmail(),
//		Gender:      req.GetGender(),
//	})
//	if error != nil {
//		return &profilepb.ClientCreateProfileResponse{
//			Status: 400,
//			Result: "خطا در ثبت اطلاعات",
//		}, nil
//	}
//	return &profilepb.ClientCreateProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت ثبت شد",
//	}, nil
//
//}
//func (s *Server) ClientGetProfile(ctx context.Context, req *profilepb.ClientGetProfileRequest) (*profilepb.ClientGetProfileResponse, error) {
//	data := &ClientProfile{}
//	_ = s.H.ClientCollection.FindOne(ctx, bson.D{{"user_id", req.GetUserId()}}).Decode(data)
//	if data.ID == primitive.NilObjectID {
//		return &profilepb.ClientGetProfileResponse{
//			Status: 400,
//			Result: "این کاربر قبلا ثبت نام نکرده است",
//		}, nil
//	}
//	return &profilepb.ClientGetProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت دریافت شد",
//		Profile: &profilepb.Profile{
//			UserId:    data.UserId,
//			Phone:     data.PhoneNumber,
//			Name:      data.Name,
//			Birthdate: data.BirthDate,
//			Email:     data.Email,
//			Gender:    data.Gender,
//		},
//	}, nil
//}
//func (s *Server) ClientUpdateProfile(ctx context.Context, req *profilepb.ClientUpdateProfileRequest) (*profilepb.ClientUpdateProfileResponse, error) {
//	filter := bson.D{{"user_id", req.GetUserId()}}
//	update := bson.D{
//		{"$set", bson.D{
//			{"phone_number", req.GetPhone()},
//			{"user_id", req.GetUserId()},
//
//			{"name", req.GetName()},
//			{"birth_date", req.GetBirthdate()},
//
//			{"email", req.GetEmail()},
//			{"gender", req.GetGender()},
//		}}}
//	_, error := s.H.TherapistCollection.UpdateOne(ctx, filter, update)
//	if error != nil {
//		return &profilepb.ClientUpdateProfileResponse{
//			Status: 400,
//			Result: "خطا در بروزرسانی اطلاعات",
//		}, nil
//
//	}
//	return &profilepb.ClientUpdateProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت بروزرسانی شد",
//	}, nil
//
//}
//func (s *Server) TherapistUpdateProfile(ctx context.Context, req *profilepb.TherapistUpdateProfileRequest) (*profilepb.TherapistUpdateProfileResponse, error) {
//	filter := bson.D{{"user_id", req.GetUserId()}}
//	update := bson.D{
//		{"$set", bson.D{
//			{"phone_number", req.GetPhone()},
//			{"user_id", req.GetUserId()},
//
//			{"name", req.GetName()},
//			{"birth_date", req.GetBirthdate()},
//
//			{"email", req.GetEmail()},
//			{"gender", req.GetGender()},
//		}}}
//	_, error := s.H.TherapistCollection.UpdateOne(ctx, filter, update)
//
//	if error != nil {
//		return &profilepb.TherapistUpdateProfileResponse{
//			Status: 400,
//			Result: "خطا در ثبت اطلاعات",
//			Id:     req.GetUserId(),
//		}, nil
//
//	}
//	return &profilepb.TherapistUpdateProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت ثبت شد",
//		Id:     req.GetUserId(),
//	}, nil
//}
//
//func (s *Server) TherapistCreateProfile(ctx context.Context, req *profilepb.TherapistCreateProfileRequest) (*profilepb.TherapistCreateProfileResponse, error) {
//	data := &TherapistProfile{}
//	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetUserId()}}).Decode(data)
//	if data.ID != primitive.NilObjectID {
//		return &profilepb.TherapistCreateProfileResponse{
//			Status: 400,
//			Result: "شماره موبایل وارد شده در سیستم ثبت شده است",
//		}, nil
//	}
//	_, error := s.H.TherapistCollection.InsertOne(ctx, &TherapistProfile{
//		UserId:      req.GetUserId(),
//		PhoneNumber: req.GetPhone(),
//		Name:        req.GetName(),
//		BirthDate:   req.GetBirthdate(),
//		Email:       req.GetEmail(),
//		Gender:      req.GetGender(),
//	})
//	if error != nil {
//		return &profilepb.TherapistCreateProfileResponse{
//			Status: 400,
//			Result: "خطا در ثبت اطلاعات",
//		}, nil
//	}
//	return &profilepb.TherapistCreateProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت ثبت شد",
//	}, nil
//
//}
//func (s *Server) TherapistGetProfile(ctx context.Context, req *profilepb.TherapistGetProfileRequest) (*profilepb.TherapistGetProfileResponse, error) {
//	data := &TherapistProfile{}
//	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
//	if data.ID == primitive.NilObjectID {
//		return &profilepb.TherapistGetProfileResponse{
//			Status: 400,
//			Result: "این کاربر قبلا ثبت نام نکرده است",
//		}, nil
//	}
//	return &profilepb.TherapistGetProfileResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت دریافت شد",
//		Profile: &profilepb.Profile{
//			UserId:    data.UserId,
//			Phone:     data.PhoneNumber,
//			Name:      data.Name,
//			Birthdate: data.BirthDate,
//			Email:     data.Email,
//			Gender:    data.Gender,
//		},
//	}, nil
//}
//
////func (s *Server) GetFreeTime(ctx context.Context, req *profilepb.TherapistGetFreeTimeRequest) (*profilepb.TherapistGetFreeTimeResponse, error) {
////	data := &TherapistProfile{}
////	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
////	if data.ID == primitive.NilObjectID {
////		return &profilepb.TherapistGetFreeTimeResponse{
////			Status: 400,
////			Result: "این کاربر قبلا ثبت نام نکرده است",
////		}, nil
////	}
////	return &profilepb.TherapistGetFreeTimeResponse{
////		Status:   200,
////		Result:   "success",
////		FreeTime: data.FreeTime,
////	}, nil
////}
//func (s *Server) GetTherapistFreeTime(ctx context.Context, req *profilepb.TherapistGetFreeTimeRequest) (*profilepb.TherapistGetFreeTimeResponse, error) {
//	data := []*profilepb.FreeTime{}
//	cursor, error := s.H.FreeTimeCollection.Find(ctx, bson.D{{"therapist_id", req.GetId()}})
//	if error != nil {
//		return &profilepb.TherapistGetFreeTimeResponse{
//			Status:   400,
//			FreeTime: []*profilepb.FreeTime{},
//		}, nil
//	}
//	for cursor.Next(ctx) {
//		var freeTime FreeTime
//		_ = cursor.Decode(&freeTime)
//		data = append(data, &profilepb.FreeTime{Date: &profilepb.Date{Year: freeTime.Date.Year, Month: freeTime.Date.Month, Day: freeTime.Date.Day, Hour: freeTime.Date.Hour, Minute: freeTime.Date.Minute},
//			TherapistId: freeTime.TherapistId,
//		})
//	}
//	return &profilepb.TherapistGetFreeTimeResponse{
//		Status:   200,
//		FreeTime: data,
//	}, nil
//
//}
//func (s *Server) GetFreeTime(ctx context.Context, req *profilepb.GetFreeTimeRequest) (*profilepb.GetFreeTimeResponse, error) {
//	data := []*profilepb.FreeTime{}
//	cursor, error := s.H.FreeTimeCollection.Find(ctx, bson.D{{req.GetKeyword().Key, req.GetKeyword().Value}})
//	if error != nil {
//		return &profilepb.GetFreeTimeResponse{
//			Status:    400,
//			FreeTimes: []*profilepb.FreeTime{},
//		}, nil
//	}
//	for cursor.Next(ctx) {
//		var freeTime FreeTime
//		_ = cursor.Decode(&freeTime)
//		data = append(data, &profilepb.FreeTime{Date: &profilepb.Date{Year: freeTime.Date.Year, Month: freeTime.Date.Month, Day: freeTime.Date.Day, Hour: freeTime.Date.Hour, Minute: freeTime.Date.Minute},
//			TherapistId: freeTime.TherapistId,
//		})
//	}
//	return &profilepb.GetFreeTimeResponse{
//		Status:    200,
//		FreeTimes: data,
//	}, nil
//}
//func (s *Server) SetFreeTime(ctx context.Context, req *profilepb.TherapistSetFreeTimeRequest) (*profilepb.TherapistSetFreeTimeResponse, error) {
//	data := &TherapistProfile{}
//	_ = s.H.TherapistCollection.FindOne(ctx, bson.D{{"user_id", req.GetId()}}).Decode(data)
//	if data.ID == primitive.NilObjectID {
//		return &profilepb.TherapistSetFreeTimeResponse{
//			Status: 400,
//			Result: "این کاربر قبلا ثبت نام نکرده است",
//		}, nil
//	}
//	var freeTimes []interface{}
//	freeTime := req.GetFreeTime()
//	for _, v := range freeTime {
//		//s.H.FreeTimeCollection.InsertOne(ctx, &FreeTime{
//		//	TherapistId: req.Id, Date: *v.Date,
//		//})
//		_ = append(freeTimes, &FreeTime{
//			Date:        Date{Year: v.Year, Month: v.Month, Day: v.Day, Hour: v.Hour, Minute: v.Minute},
//			TherapistId: req.Id,
//		})
//	}
//
//	s.H.FreeTimeCollection.DeleteMany(ctx, bson.D{{"id", data.UserId}})
//	_, error := s.H.FreeTimeCollection.InsertMany(ctx, freeTimes)
//	if error != nil {
//		return &profilepb.TherapistSetFreeTimeResponse{
//			Status: 400,
//			Result: "خطا در ثبت اطلاعات",
//		}, nil
//	}
//	return &profilepb.TherapistSetFreeTimeResponse{
//		Status: 200,
//		Result: "اطلاعات با موفقیت ثبت شد",
//	}, nil
//}
//
////func (s *Server) PhoneToken(ctx context.Context, req *profilepb.ClientPhoneTokenRequest) (*profilepb.ClientPhoneTokenResponse, error) {
////	fmt.Println("ClientPhoneToken invoked %v", req)
////	var smsVerify models.SmsVerify
////
////	if result := s.H.DB.Where(&models.SmsVerify{Phone: req.Phone}).First(&smsVerify); result.Error != nil {
////		return &profilepb.ClientPhoneTokenResponse{
////			Result: "not found",
////			Token:  "",
////		}, nil
////	}
////
////	if smsVerify.Code != req.Code {
////		return &profilepb.ClientPhoneTokenResponse{
////			Result: "code not match",
////			Token:  "",
////		}, nil
////	}
////	var user models.User
////	user.Phone = req.Phone
////	user.ClientType = req.ClientType
////	s.H.DB.Create(&user)
////	token, _ := s.Jwt.GenerateToken(user)
////	return &profilepb.ClientPhoneTokenResponse{
////		Result: "ok",
////		Token:  token,
////	}, nil
////}
////
////func (s *Server) PhoneProof(ctx context.Context, req *profilepb.ClientPhoneProofRequest) (*profilepb.ClientPhoneProofResponse, error) {
////
////	fmt.Println("ClientPhoneProof invoked %v", req)
////	var smsVerify models.SmsVerify
////	if result := s.H.DB.Where(&models.User{Phone: req.Phone}).First(&smsVerify); result.Error == nil {
////		return &profilepb.ClientPhoneProofResponse{
////			Status: http.StatusConflict,
////		}, nil
////	}
////	smsVerify.Phone = req.Phone
////	smsVerify.Code = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
////	sender := "10008663"
////	receptor := []string{"09120357479"}
////	message := "کد تایید شما : " + smsVerify.Code
////	fmt.Println("moz")
////
////	if _, err := s.Sms.SmsClient.Message.Send(sender, receptor, message, nil); err != nil {
////		fmt.Println("moz")
////		switch err := err.(type) {
////		case *kavenegar.APIError:
////			fmt.Println("api")
////
////			fmt.Println(err.Error())
////		case *kavenegar.HTTPError:
////			fmt.Println("http")
////
////			fmt.Println(err.Error())
////		default:
////			fmt.Println("error")
////
////			fmt.Println(err.Error())
////		}
////	}
////	smsVerify.ExpireTime = time.Now().Add(time.Minute * 5)
////	s.H.DB.Create(&smsVerify)
////	return &profilepb.ClientPhoneProofResponse{
////		Status: http.StatusCreated,
////	}, nil
////}
////func (s *Server) ValidateToken(ctx context.Context, req *profilepb.ValidateRequest) (*profilepb.ValidateResponse, error) {
////	claims, err := s.Jwt.ValidateToken(req.Token)
////
////	if err != nil {
////		return &profilepb.ValidateResponse{
////			Status: http.StatusBadRequest,
////			Error:  err.Error(),
////		}, nil
////	}
////
////	var user models.User
////
////	if result := s.H.DB.Where(&models.User{Phone: claims.Phone}).First(&user); result.Error != nil {
////		return &profilepb.ValidateResponse{
////			Status: http.StatusNotFound,
////			Error:  "User not found",
////		}, nil
////	}
////
////	return &profilepb.ValidateResponse{
////		Status: http.StatusOK,
////		UserId: user.Id,
////	}, nil
////}
