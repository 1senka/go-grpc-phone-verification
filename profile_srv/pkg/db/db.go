package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ClientCollection *mongo.Collection
var TherapistCollection *mongo.Collection
var FreeTimeCollection *mongo.Collection

type Handler struct {
	ClientCollection    *mongo.Collection
	TherapistCollection *mongo.Collection
	FreeTimeCollection  *mongo.Collection
}

func Init(url string) Handler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Fatalln(err)
	}
	ClientCollection = db.Database("profile").Collection("client")
	TherapistCollection = db.Database("profile").Collection("therapist")
	FreeTimeCollection = db.Database("profile").Collection("freetime")
	return Handler{ClientCollection: ClientCollection, TherapistCollection: TherapistCollection}
}
