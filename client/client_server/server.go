package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/kavenegar/kavenegar-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"psycology_backend/client/clientpb"
	"strconv"
	"time"
)

var smsVerifyCollection *mongo.Collection
var userCollection *mongo.Collection
var smsClient *kavenegar.Kavenegar

type server struct {
}

type smsVerify struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	Code        string             `bson:"code"`
	ExpireTime  time.Time          `bson:"expire_time"`
}
type user struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PhoneNumber string             `bson:"phone_number"`
	Name        string             `bson:"name"`
	BirthDate   string             `bson:"birth_date"`
}

//func (*server) Login(ctx context.Context, req *clientpb.ClientLoginRequest) (*clientpb.ClientLoginResponse, error) {
//	fmt.Println("Login invoked %v", req)
//	name := req.GetClient().GetName()
//	n := "token" + name
//	res := &clientpb.ClientLoginResponse{
//		Token: n,
//	}
//	return res, nil
//}
func (*server) FirstCred(ctx context.Context, req *clientpb.ClientFirstCredRequest) (*clientpb.ClientFirstCredResponse, error) {
	fmt.Println("FirstCred invoked %v", req)
	Name := req.GetName()

	//_, err := userCollection.InsertOne(context.Background(), data)
	//if err != nil {
	//	fmt.Println("err in FirstCred %v", err)
	//	return nil, err
	//}
	//result := &clientpb.ClientFirstCredResponse{Result: "ok"}
	//return result, nil
	err := userCollection.FindOneAndUpdate(context.Background(), bson.D{{"phone_number", req.GetPhone()}}, bson.D{{"$set", bson.D{{"name", Name}, {"birth_date", strconv.Itoa(int(req.GetBirthday().Year)) + "-" + strconv.Itoa(int(req.GetBirthday().Month)) + "-" + strconv.Itoa(int(req.GetBirthday().Day))}}}}).Err()
	if err != nil {
		fmt.Println("err in FirstCred %v", err)
		return nil, err
	}
	return &clientpb.ClientFirstCredResponse{Result: "ok"}, nil

}
func (*server) PhoneProve(ctx context.Context, req *clientpb.ClientPhoneProofRequest) (*clientpb.ClientPhoneProofResponse, error) {
	fmt.Println("ClientPhoneProof invoked %v", req)
	code := rand.Intn(99999999-10000000) + 10000000

	data := &smsVerify{
		PhoneNumber: req.GetPhone(),
		Code:        strconv.Itoa(code),
		ExpireTime:  time.Now().Add(time.Minute * 5),
	}
	_, err := smsVerifyCollection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sender := "10008663"
	receptor := []string{"09120357479"}
	message := "کد تایید شما : " + strconv.Itoa(code)
	if _, err := smsClient.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	}
	res := &clientpb.ClientPhoneProofResponse{
		IsValid: "ok",
	}
	return res, nil
}

func (*server) PhoneToken(ctx context.Context, req *clientpb.ClientPhoneTokenRequest) (*clientpb.ClientPhoneTokenResponse, error) {
	fmt.Println("ClientPhoneToken invoked %v", req)
	data := &smsVerify{}

	err := smsVerifyCollection.FindOne(context.Background(), bson.M{"phone_number": req.GetPhone(), "code": req.GetToken()}).Decode(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if data.ExpireTime.After(time.Now()) {
		res := &clientpb.ClientPhoneTokenResponse{
			Result: "expired",
		}
		return res, nil
	}
	d := &user{
		PhoneNumber: req.GetPhone(),
		Name:        "",
		BirthDate:   "",
	}
	doesExist := &user{}
	err = userCollection.FindOne(context.Background(), bson.M{"phone_number": req.GetPhone()}).Decode(doesExist)
	if doesExist.ID != primitive.NilObjectID {
		return nil, errors.New("user already exists")
	}
	_, error := userCollection.InsertOne(context.Background(), d)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	res := &clientpb.ClientPhoneTokenResponse{
		Result: "ok",
	}
	return res, nil
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("starting server code!")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	smsVerifyCollection = client.Database("psycology").Collection("smsVerify")
	userCollection = client.Database("psycology").Collection("user")
	smsClient = kavenegar.New("556E65702B4A626279345335534B5835484574733734464238734D4D6A5143412F6A6935786C72635554303D")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	s := grpc.NewServer()
	clientpb.RegisterClientServiceServer(s, &server{})
	go func() {
		if error := s.Serve(lis); error != nil {
			log.Fatalf("Error: %v", error)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Server stopped")
	lis.Close()

}
