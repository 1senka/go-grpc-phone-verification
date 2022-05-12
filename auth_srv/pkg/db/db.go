package db

import (
	"context"
	"log"
	"time"

	"auth_svc/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var SmsVerifyCollection *mongo.Collection
var UserCollection *mongo.Collection
type Handler struct {
	 SmsVerifyCollection *mongo.Collection
 UserCollection *mongo.Collection
}

func Init(url string) Handler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(c.DBUrl))
	if err != nil {
		log.Fatal(err)
	}
	
	if err != nil {
		log.Fatalln(err)
	}
	index := mongo.IndexModel{
		Keys: bsonx.Doc{
	   {"created_at", bsonx.Int32(1)}},
		Options:options.Index().SetExpireAfterSeconds(30),
	}
	
	SmsVerifyCollection = db.Database("auth_svc").Collection("smsVerify")
	UserCollection = db.Database("auth_svc").Collection("user")
	SmsVerifyCollection.Indexes().CreateOne(ctx ,index)
	
	return Handler{SmsVerifyCollection: SmsVerifyCollection, UserCollection: UserCollection}
}
