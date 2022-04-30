package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ClientCollection *mongo.Collection
var TherapistCollection *mongo.Collection

type Handler struct {
	ClientCollection    *mongo.Collection
	TherapistCollection *mongo.Collection
}

func Init(url string) Handler {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	/*
	   List databases
	*/
	fmt.Println(client.ListDatabaseNames(ctx, bson.M{}))
	ClientCollection = client.Database("psycology").Collection("client")
	TherapistCollection = client.Database("psycology").Collection("therapist")
	return Handler{ClientCollection: ClientCollection, TherapistCollection: TherapistCollection}
}
